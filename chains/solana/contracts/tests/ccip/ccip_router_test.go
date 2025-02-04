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
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	ag_binary "github.com/gagliardetto/binary"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_ccip_receiver"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

const MaxCU = 1_400_000 // this is SVM's hard max Compute Unit limit

func TestCCIPRouter(t *testing.T) {
	t.Parallel()

	ccip_router.SetProgramID(config.CcipRouterProgram)
	test_ccip_receiver.SetProgramID(config.CcipLogicReceiver)
	token_pool.SetProgramID(config.CcipTokenPoolProgram)
	fee_quoter.SetProgramID(config.FeeQuoterProgram)

	ctx := tests.Context(t)

	user, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	anotherUser, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	tokenlessUser, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	legacyAdmin, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	ccipAdmin, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	token0PoolAdmin, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	token1PoolAdmin, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	feeAggregator, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)

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

	// token addresses
	token0, gerr := tokens.NewTokenPool(config.Token2022Program)
	require.NoError(t, gerr)
	token1, gerr := tokens.NewTokenPool(config.Token2022Program)
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
	emptyAddress := [64]byte{}

	validSourceChainConfig := ccip_router.SourceChainConfig{
		OnRamp:    [2][64]byte{onRampAddress, emptyAddress},
		IsEnabled: true,
	}
	validFqDestChainConfig := fee_quoter.DestChainConfig{
		IsEnabled: true,

		// minimal valid config
		DefaultTxGasLimit:           200000,
		MaxPerMsgGasLimit:           3000000,
		MaxDataBytes:                30000,
		MaxNumberOfTokensPerMsg:     5,
		DefaultTokenDestGasOverhead: 50000,
		// bytes4(keccak256("CCIP ChainFamilySelector EVM"))
		ChainFamilySelector: [4]uint8{40, 18, 213, 44},

		DefaultTokenFeeUsdcents: 50,
		NetworkFeeUsdcents:      50,
	}
	// Small enough to fit in u160, big enough to not fall in the precompile space.
	validReceiverAddress := [32]byte{}
	validReceiverAddress[12] = 1
	emptyEVMExtraArgsV2 := testutils.MustSerializeExtraArgs(t, struct{}{}, ccip.EVMExtraArgsV2Tag)

	var ccipSendLookupTable map[solana.PublicKey]solana.PublicKeySlice
	var commitLookupTable map[solana.PublicKey]solana.PublicKeySlice

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
				feeAggregator),
				solanaGoClient,
				t)
		})

		t.Run("receiver", func(t *testing.T) {
			instruction, ixErr := test_ccip_receiver.NewInitializeInstruction(
				config.ReceiverTargetAccountPDA,
				config.ReceiverExternalExecutionConfigPDA,
				user.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, ixErr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
		})

		t.Run("token", func(t *testing.T) {
			ixs, ixErr := tokens.CreateToken(ctx, token0.Program, token0.Mint.PublicKey(), token0PoolAdmin.PublicKey(), 0, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, ixErr)

			ixsAnotherToken, anotherTokenErr := tokens.CreateToken(ctx, token1.Program, token1.Mint.PublicKey(), token1PoolAdmin.PublicKey(), 0, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, anotherTokenErr)

			// mint tokens to user
			ixAta0, addr0, ataErr := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint.PublicKey(), user.PublicKey(), token0PoolAdmin.PublicKey())
			require.NoError(t, ataErr)
			ixMintTo0, mintErr := tokens.MintTo(10000000, token0.Program, token0.Mint.PublicKey(), addr0, token0PoolAdmin.PublicKey())
			require.NoError(t, mintErr)
			ixAta1, addr1, ataErr := tokens.CreateAssociatedTokenAccount(token1.Program, token1.Mint.PublicKey(), user.PublicKey(), token0PoolAdmin.PublicKey())
			require.NoError(t, ataErr)
			ixMintTo1, mintErr := tokens.MintTo(10000000, token1.Program, token1.Mint.PublicKey(), addr1, token1PoolAdmin.PublicKey())
			require.NoError(t, mintErr)

			// create ATA for receiver (receiver program address)
			ixAtaReceiver0, recAddr0, recErr := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint.PublicKey(), config.ReceiverExternalExecutionConfigPDA, token0PoolAdmin.PublicKey())
			require.NoError(t, recErr)
			ixAtaReceiver1, recAddr1, recErr := tokens.CreateAssociatedTokenAccount(token1.Program, token1.Mint.PublicKey(), config.ReceiverExternalExecutionConfigPDA, token0PoolAdmin.PublicKey())
			require.NoError(t, recErr)

			token0.User[user.PublicKey()] = addr0
			token0.User[config.ReceiverExternalExecutionConfigPDA] = recAddr0
			token1.User[user.PublicKey()] = addr1
			token1.User[config.ReceiverExternalExecutionConfigPDA] = recAddr1
			ixs = append(ixs, ixAta0, ixMintTo0, ixAtaReceiver0)
			ixs = append(ixs, ixsAnotherToken...)
			ixs = append(ixs, ixAta1, ixMintTo1, ixAtaReceiver1)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, ixs, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token0.Mint, token1.Mint, token1PoolAdmin))
		})

		t.Run("token-pool", func(t *testing.T) {
			token0.AdditionalAccounts = append(token0.AdditionalAccounts, solana.MemoProgramID) // add test additional accounts in pool interactions

			ixInit0, err := token_pool.NewInitializeInstruction(
				token_pool.BurnAndMint_PoolType,
				config.ExternalTokenPoolsSignerPDA,
				token0.PoolConfig,
				token0.Mint.PublicKey(),
				token0.PoolSigner,
				token0PoolAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			ixInit1, err := token_pool.NewInitializeInstruction(
				token_pool.BurnAndMint_PoolType,
				config.ExternalTokenPoolsSignerPDA,
				token1.PoolConfig,
				token1.Mint.PublicKey(),
				token1.PoolSigner,
				token1PoolAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			ixAta0, addr0, err := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint.PublicKey(), token0.PoolSigner, token0PoolAdmin.PublicKey())
			require.NoError(t, err)
			token0.PoolTokenAccount = addr0
			token0.User[token0.PoolSigner] = token0.PoolTokenAccount
			ixAta1, addr1, err := tokens.CreateAssociatedTokenAccount(token1.Program, token1.Mint.PublicKey(), token1.PoolSigner, token0PoolAdmin.PublicKey())
			require.NoError(t, err)
			token1.PoolTokenAccount = addr1
			token1.User[token1.PoolSigner] = token1.PoolTokenAccount

			ixAuth, err := tokens.SetTokenMintAuthority(token0.Program, token0.PoolSigner, token0.Mint.PublicKey(), token0PoolAdmin.PublicKey())
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixInit0, ixInit1, ixAta0, ixAta1, ixAuth}, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token1PoolAdmin))

			// Lookup Table for Tokens
			require.NoError(t, token0.SetupLookupTable(ctx, solanaGoClient, token0PoolAdmin))
			token0Entries := token0.ToTokenPoolEntries()
			require.NoError(t, token1.SetupLookupTable(ctx, solanaGoClient, token1PoolAdmin))
			token1Entries := token1.ToTokenPoolEntries()

			// Verify Lookup tables where correctly initialized
			lookupTableEntries0, err := common.GetAddressLookupTable(ctx, solanaGoClient, token0.PoolLookupTable)
			require.NoError(t, err)
			require.Equal(t, len(token0Entries), len(lookupTableEntries0))
			require.Equal(t, token0Entries, lookupTableEntries0)

			lookupTableEntries1, err := common.GetAddressLookupTable(ctx, solanaGoClient, token1.PoolLookupTable)
			require.NoError(t, err)
			require.Equal(t, len(token1Entries), len(lookupTableEntries1))
			require.Equal(t, token1Entries, lookupTableEntries1)
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

			// Create link22 token, managed by "admin" (not "ccipAdmin" who manages CCIP).
			// Random-generated key, but fixing it adds determinism to tests to make it easier to debug.
			mintPrivK := solana.MustPrivateKeyFromBase58("32YVeJArcWWWV96fztfkRQhohyFz5Hwno93AeGVrN4g2LuFyvwznrNd9A6tbvaTU6BuyBsynwJEMLre8vSy3CrVU")

			mintPubK := mintPrivK.PublicKey()
			ixToken, terr := tokens.CreateToken(ctx, config.Token2022Program, mintPubK, legacyAdmin.PublicKey(), 9, solanaGoClient, config.DefaultCommitment)
			require.NoError(t, terr)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, ixToken, legacyAdmin, config.DefaultCommitment, common.AddSigners(mintPrivK))

			link22PDA, _, aerr := state.FindFqBillingTokenConfigPDA(mintPubK, config.FeeQuoterProgram)
			require.NoError(t, aerr)
			link22EvmConfigPDA, _, puerr := state.FindFqPerChainPerTokenConfigPDA(config.EvmChainSelector, mintPubK, config.FeeQuoterProgram)
			require.NoError(t, puerr)
			link22Receiver, _, rerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, mintPubK, config.BillingSignerPDA)
			require.NoError(t, rerr)
			link22UserATA, _, uerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, mintPubK, user.PublicKey())
			require.NoError(t, uerr)
			link22AnotherUserATA, _, auerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, mintPubK, anotherUser.PublicKey())
			require.NoError(t, auerr)
			link22TokenlessUserATA, _, tuerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, mintPubK, tokenlessUser.PublicKey())
			require.NoError(t, tuerr)
			link22FeeAggregatorATA, _, fuerr := tokens.FindAssociatedTokenAddress(config.Token2022Program, mintPubK, feeAggregator.PublicKey())
			require.NoError(t, fuerr)

			// persist the link22 billing config for later use
			link22.program = config.Token2022Program
			link22.mint = mintPubK
			link22.fqBillingConfigPDA = link22PDA
			link22.userATA = link22UserATA
			link22.anotherUserATA = link22AnotherUserATA
			link22.tokenlessUserATA = link22TokenlessUserATA
			link22.billingATA = link22Receiver
			link22.feeAggregatorATA = link22FeeAggregatorATA
			link22.fqEvmConfigPDA = link22EvmConfigPDA
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
				config.ExternalTokenPoolsSignerPDA,
			}
			lookupTableAddr, err := common.SetupLookupTable(ctx, solanaGoClient, legacyAdmin, lookupEntries)
			require.NoError(t, err)

			ccipSendLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
				lookupTableAddr: lookupEntries,
			}
		})

		t.Run("Commit price updates address lookup table", func(t *testing.T) {
			// Create single Address Lookup Table, to be used in all commit tests.
			// Create it early in the test suite (a "setup" step) to let it warm up with more than enough time,
			// as otherwise it can slow down tests  for ~20 seconds.

			lookupEntries := []solana.PublicKey{
				// static accounts that are always needed
				ccip_router.ProgramID,
				config.RouterConfigPDA,
				config.RouterStatePDA,
				config.EvmSourceChainStatePDA, // for checking the seq numbers
				solana.SystemProgramID,
				solana.SysVarInstructionsPubkey,
				config.ExternalExecutionConfigPDA,
				config.ExternalTokenPoolsSignerPDA,
				config.BillingSignerPDA,
				config.FeeQuoterProgram,
				config.FqConfigPDA,

				// remaining_accounts that are only sometimes needed
				wsol.fqBillingConfigPDA,
				link22.fqBillingConfigPDA,
				config.EvmDestChainStatePDA, // to update prices
				config.SvmDestChainStatePDA,
			}
			lookupTableAddr, err := common.SetupLookupTable(ctx, solanaGoClient, legacyAdmin, lookupEntries)
			require.NoError(t, err)

			commitLookupTable = map[solana.PublicKey]solana.PublicKeySlice{
				lookupTableAddr: lookupEntries,
			}
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

		t.Run("Router is initialized", func(t *testing.T) {
			invalidSVMChainSelector := uint64(17)
			defaultMaxFeeJuelsPerMsg := bin.Uint128{Lo: 300000000, Hi: 0, Endianness: nil}

			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipRouterProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)

			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			instruction, err := ccip_router.NewInitializeInstruction(
				invalidSVMChainSelector,
				config.EnableExecutionAfter,
				// fee aggregator address, will be changed in later test
				anotherUser.PublicKey(),
				config.FeeQuoterProgram,
				link22.mint,
				defaultMaxFeeJuelsPerMsg,
				config.RouterConfigPDA,
				config.RouterStatePDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.CcipRouterProgram,
				programData.Address,
				config.ExternalExecutionConfigPDA,
				config.ExternalTokenPoolsSignerPDA,
			).ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			// Fetch account data
			var configAccount ccip_router.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
			if err != nil {
				require.NoError(t, err, "failed to get account info")
			}
			require.Equal(t, uint64(17), configAccount.SvmChainSelector)
			require.Equal(t, config.FeeQuoterProgram, configAccount.FeeQuoter)
		})

		t.Run("FeeQuoter is initialized", func(t *testing.T) {
			defaultMaxFeeJuelsPerMsg := bin.Uint128{Lo: 300000000, Hi: 0, Endianness: nil}

			// get program data account
			data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.FeeQuoterProgram, &rpc.GetAccountInfoOpts{
				Commitment: config.DefaultCommitment,
			})
			require.NoError(t, err)

			// Decode program data
			var programData ProgramData
			require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

			ix, err := fee_quoter.NewInitializeInstruction(
				link22.mint,
				defaultMaxFeeJuelsPerMsg,
				config.CcipRouterProgram,
				config.BillingSignerPDA, // TODO fix offramp_signer address
				// config solana.PublicKey, authority solana.PublicKey, systemProgram solana.PublicKey, program solana.PublicKey, programData solana.PublicKey
				config.FqConfigPDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
				config.FeeQuoterProgram,
				programData.Address,
			).ValidateAndBuild()
			require.NoError(t, err)

			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			// Fetch account data
			var fqConfig fee_quoter.Config
			require.NoError(t, common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &fqConfig))

			require.Equal(t, link22.mint, fqConfig.LinkTokenMint)
			require.Equal(t, defaultMaxFeeJuelsPerMsg, fqConfig.MaxFeeJuelsPerMsg)
			require.Equal(t, legacyAdmin.PublicKey(), fqConfig.Owner)
			require.True(t, fqConfig.ProposedOwner.IsZero())
			require.Equal(t, config.CcipRouterProgram, fqConfig.Onramp)
			require.Equal(t, config.BillingSignerPDA, fqConfig.OfframpSigner) // TODO fix this
		})

		t.Run("When admin updates the solana chain selector it's updated", func(t *testing.T) {
			instruction, err := ccip_router.NewUpdateSvmChainSelectorInstruction(
				config.SvmChainSelector,
				config.RouterConfigPDA,
				legacyAdmin.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)
			result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
			require.NotNil(t, result)

			var configAccount ccip_router.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
			require.NoError(t, err, "failed to get account info")
			require.Equal(t, config.SvmChainSelector, configAccount.SvmChainSelector)
		})

		type InvalidChainBillingInputTest struct {
			Name         string
			Selector     uint64
			Conf         fee_quoter.DestChainConfig
			SkipOnUpdate bool
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
			},
			{
				Name:         "Zero DestChainSelector",
				Selector:     0,
				Conf:         validFqDestChainConfig,
				SkipOnUpdate: true, // as the 0-selector is invalid, the config account can never be initialized
			},
			{
				Name:     "Zero ChainFamilySelector",
				Selector: config.EvmChainSelector,
				Conf: fee_quoter.DestChainConfig{
					DefaultTxGasLimit:   validFqDestChainConfig.DefaultTxGasLimit,
					MaxPerMsgGasLimit:   validFqDestChainConfig.MaxPerMsgGasLimit,
					ChainFamilySelector: [4]uint8{0, 0, 0, 0},
				},
			},
			{
				Name:     "DefaultTxGasLimit > MaxPerMsgGasLimit",
				Selector: config.EvmChainSelector,
				Conf: fee_quoter.DestChainConfig{
					DefaultTxGasLimit:   100,
					MaxPerMsgGasLimit:   1,
					ChainFamilySelector: validFqDestChainConfig.ChainFamilySelector,
				},
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
					result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidInputs_CcipRouterError.String()}) // TODO change error type
					require.NotNil(t, result)
				})
			}
		})

		t.Run("When an unauthorized user tries to add a chain selector, it fails", func(t *testing.T) {
			t.Run("CCIP Router", func(t *testing.T) {
				instruction, err := ccip_router.NewAddChainSelectorInstruction(
					config.EvmChainSelector,
					validSourceChainConfig,
					ccip_router.DestChainConfig{},
					config.EvmSourceChainStatePDA,
					config.EvmDestChainStatePDA,
					config.RouterConfigPDA,
					user.PublicKey(), // not an admin
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()}) // TODO fix error type to fee_quoter
				require.NotNil(t, result)
			})
		})

		t.Run("When admin adds a chain selector it's added on the list", func(t *testing.T) {
			t.Run("CCIP Router", func(t *testing.T) {
				instruction, err := ccip_router.NewAddChainSelectorInstruction(
					config.EvmChainSelector,
					validSourceChainConfig,
					ccip_router.DestChainConfig{},
					config.EvmSourceChainStatePDA,
					config.EvmDestChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var sourceChainStateAccount ccip_router.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmSourceChainStatePDA, config.DefaultCommitment, &sourceChainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, uint64(1), sourceChainStateAccount.State.MinSeqNr)
				require.Equal(t, true, sourceChainStateAccount.Config.IsEnabled)
				require.Equal(t, [2][64]byte{config.OnRampAddressPadded, emptyAddress}, sourceChainStateAccount.Config.OnRamp)

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
		})

		t.Run("When admin adds another chain selector it's also added on the list", func(t *testing.T) {
			// Using another chain, solana as an example (which allows SVM -> SVM messages)
			// Regardless of whether we allow SVM -> SVM in mainnet, it's easy to use for tests here

			// the router is the SVM onramp
			var paddedCcipRouterProgram [64]byte
			copy(paddedCcipRouterProgram[:], config.CcipRouterProgram[:])
			onRampConfig := [2][64]byte{paddedCcipRouterProgram, emptyAddress}

			t.Run("CCIP Router", func(t *testing.T) {
				instruction, err := ccip_router.NewAddChainSelectorInstruction(
					config.SvmChainSelector,
					ccip_router.SourceChainConfig{
						OnRamp:    onRampConfig, // the source on ramp address must be padded, as this value is an array of 64 bytes
						IsEnabled: true,
					},
					ccip_router.DestChainConfig{},
					config.SvmSourceChainStatePDA,
					config.SvmDestChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var sourceChainStateAccount ccip_router.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.SvmSourceChainStatePDA, config.DefaultCommitment, &sourceChainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, uint64(1), sourceChainStateAccount.State.MinSeqNr)
				require.Equal(t, true, sourceChainStateAccount.Config.IsEnabled)
				require.Equal(t, [2][64]byte{paddedCcipRouterProgram, emptyAddress}, sourceChainStateAccount.Config.OnRamp)

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
						DefaultTxGasLimit:   1,
						MaxPerMsgGasLimit:   100,
						ChainFamilySelector: [4]uint8{3, 2, 1, 0},
						EnforceOutOfOrder:   true,
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
		})

		t.Run("When a non-admin tries to disable the chain selector, it fails", func(t *testing.T) {
			t.Run("Source", func(t *testing.T) {
				ix, err := ccip_router.NewDisableSourceChainSelectorInstruction(
					config.EvmChainSelector,
					config.EvmSourceChainStatePDA,
					config.RouterConfigPDA,
					user.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
				require.NotNil(t, result)
			})
		})

		t.Run("When an admin disables the chain selector, it is no longer enabled", func(t *testing.T) {
			t.Run("Source", func(t *testing.T) {
				var initial ccip_router.SourceChain
				err := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmSourceChainStatePDA, config.DefaultCommitment, &initial)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, true, initial.Config.IsEnabled)

				ix, err := ccip_router.NewDisableSourceChainSelectorInstruction(
					config.EvmChainSelector,
					config.EvmSourceChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, legacyAdmin, config.DefaultCommitment)

				var final ccip_router.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmSourceChainStatePDA, config.DefaultCommitment, &final)
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
					result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidInputs_CcipRouterError.String()})
					require.NotNil(t, result)
				})
			}
		})

		t.Run("When an unauthorized user tries to update the chain state config, it fails", func(t *testing.T) {
			t.Run("Source", func(t *testing.T) {
				instruction, err := ccip_router.NewUpdateSourceChainConfigInstruction(
					config.EvmChainSelector,
					validSourceChainConfig,
					config.EvmSourceChainStatePDA,
					config.RouterConfigPDA,
					user.PublicKey(), // unauthorized
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
				require.NotNil(t, result)
			})
		})

		t.Run("When an admin updates the chain state config, it is configured", func(t *testing.T) {
			var initialSource ccip_router.SourceChain
			serr := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmSourceChainStatePDA, config.DefaultCommitment, &initialSource)
			require.NoError(t, serr, "failed to get account info")

			t.Run("Source", func(t *testing.T) {
				updated := initialSource.Config
				updated.IsEnabled = true
				require.NotEqual(t, initialSource.Config, updated) // at this point, onchain is disabled and we'll re-enable it

				instruction, err := ccip_router.NewUpdateSourceChainConfigInstruction(
					config.EvmChainSelector,
					updated,
					config.EvmSourceChainStatePDA,
					config.RouterConfigPDA,
					legacyAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				var final ccip_router.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmSourceChainStatePDA, config.DefaultCommitment, &final)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, updated, final.Config)
			})

			var initialDest fee_quoter.DestChain
			derr := common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqEvmDestChainPDA, config.DefaultCommitment, &initialDest)
			require.NoError(t, derr, "failed to get account info")

			t.Run("Fee Quoter: Dest", func(t *testing.T) {
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
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidInputs_CcipRouterError.String()})
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
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})
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
			result = testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidInputs_CcipRouterError.String()}) // TODO change error type to FQ
			require.NotNil(t, result)

			// Validate proposed set to 0-address
			var configAccount fee_quoter.Config
			err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.FqConfigPDA, config.DefaultCommitment, &configAccount)
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

		t.Run("setup:funding_and_approvals", func(t *testing.T) {
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

		t.Run("FeeQuoter: Billing Token Config", func(t *testing.T) {
			pools := []tokens.TokenPool{token0, token1}

			for i, token := range pools {
				t.Run(fmt.Sprintf("token%d", i), func(t *testing.T) {
					t.Run("Pre-condition: Does not support token by default", func(t *testing.T) {
						tokenBillingPDA := getFqTokenConfigPDA(token.Mint.PublicKey())
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
							Mint:    token.Mint.PublicKey(),
							UsdPerToken: fee_quoter.TimestampedPackedU224{
								Timestamp: validTimestamp,
								Value:     value,
							},
							PremiumMultiplierWeiPerEth: 1,
						}

						tokenBillingPDA := getFqTokenConfigPDA(token.Mint.PublicKey())
						tokenReceiver, _, ferr := tokens.FindAssociatedTokenAddress(token.Program, token.Mint.PublicKey(), config.BillingSignerPDA)
						require.NoError(t, ferr)

						ixConfig, cerr := fee_quoter.NewAddBillingTokenConfigInstruction(
							tokenConfig,
							config.FqConfigPDA,
							tokenBillingPDA,
							token.Program,
							token.Mint.PublicKey(),
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
						tokenBillingPDA := getFqTokenConfigPDA(token.Mint.PublicKey())
						var initial fee_quoter.BillingTokenConfigWrapper
						ierr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &initial)
						require.NoError(t, ierr)

						tokenConfig := initial.Config
						tokenConfig.PremiumMultiplierWeiPerEth = initial.Config.PremiumMultiplierWeiPerEth*2 + 1 // updating something valid

						ixConfig, cerr := fee_quoter.NewUpdateBillingTokenConfigInstruction(tokenConfig, config.FqConfigPDA, tokenBillingPDA, legacyAdmin.PublicKey()).ValidateAndBuild() // wrong admin
						require.NoError(t, cerr)
						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ixConfig}, legacyAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})

						var final fee_quoter.BillingTokenConfigWrapper
						ferr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &final)
						require.NoError(t, ferr)

						require.Equal(t, initial.Config, final.Config) // it was not updated, same values as initial
					})

					t.Run("When admin updates token it is updated", func(t *testing.T) {
						tokenBillingPDA := getFqTokenConfigPDA(token.Mint.PublicKey())
						var initial fee_quoter.BillingTokenConfigWrapper
						ierr := common.GetAccountDataBorshInto(ctx, solanaGoClient, tokenBillingPDA, config.DefaultCommitment, &initial)
						require.NoError(t, ierr)

						tokenConfig := initial.Config
						tokenConfig.PremiumMultiplierWeiPerEth = initial.Config.PremiumMultiplierWeiPerEth*2 + 1 // updating something else

						ixConfig, cerr := fee_quoter.NewUpdateBillingTokenConfigInstruction(tokenConfig, config.FqConfigPDA, tokenBillingPDA, ccipAdmin.PublicKey()).ValidateAndBuild()
						require.NoError(t, cerr)
						testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixConfig}, ccipAdmin, config.DefaultCommitment)

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

	t.Run("Config SetOcrConfig", func(t *testing.T) {
		t.Run("Successfully configures commit & execute DON ocr config for maximum signers and transmitters", func(t *testing.T) {
			// Check owner permissions
			instruction, err := ccip_router.NewSetOcrConfigInstruction(
				0,
				ccip_router.Ocr3ConfigInfo{},
				[][20]byte{},
				[]solana.PublicKey{},
				config.RouterConfigPDA,
				config.RouterStatePDA,
				user.PublicKey(),
			).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip_router.Unauthorized_CcipRouterError.String()})

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
					instruction, err := ccip_router.NewSetOcrConfigInstruction(
						uint8(v.plugin),
						ccip_router.Ocr3ConfigInfo{
							ConfigDigest:                   config.ConfigDigest,
							F:                              config.OcrF,
							IsSignatureVerificationEnabled: v.verifySig,
						},
						v.signers,
						v.transmitters,
						config.RouterConfigPDA,
						config.RouterStatePDA,
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
					var configAccount ccip_router.Config
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
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
				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(100),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					transmitterPubKeys,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidInputs_CcipRouterError.String()})
			})

			t.Run("It rejects F = 0", func(t *testing.T) {
				t.Parallel()
				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            0,
					},
					signerAddresses,
					transmitterPubKeys,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorInvalidConfigFMustBePositive.String()})
			})

			t.Run("It rejects too many transmitters", func(t *testing.T) {
				t.Parallel()
				invalidTransmitters := make([]solana.PublicKey, config.MaxOracles+1)
				for i := range invalidTransmitters {
					invalidTransmitters[i] = getTransmitter().PublicKey()
				}
				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					invalidTransmitters,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorInvalidConfigTooManyTransmitters.String()})
			})

			t.Run("It rejects too many signers", func(t *testing.T) {
				t.Parallel()
				invalidSigners := make([][20]byte, config.MaxOracles+1)
				for i := range invalidSigners {
					invalidSigners[i] = signerAddresses[0]
				}

				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					invalidSigners,
					transmitterPubKeys,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorInvalidConfigTooManySigners.String()})
			})

			t.Run("It rejects too high of F for signers", func(t *testing.T) {
				t.Parallel()
				invalidSigners := make([][20]byte, 1)
				invalidSigners[0] = signerAddresses[0]

				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					invalidSigners,
					transmitterPubKeys,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorInvalidConfigFIsTooHigh.String()})
			})

			t.Run("It rejects duplicate transmitters", func(t *testing.T) {
				t.Parallel()
				transmitter := getTransmitter().PublicKey()

				invalidTransmitters := make([]solana.PublicKey, 2)
				for i := range invalidTransmitters {
					invalidTransmitters[i] = transmitter
				}
				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					invalidTransmitters,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorInvalidConfigRepeatedOracle.String()})
			})

			t.Run("It rejects duplicate signers", func(t *testing.T) {
				t.Parallel()
				repeatedSignerAddresses := [][20]byte{}
				for range signers {
					repeatedSignerAddresses = append(repeatedSignerAddresses, signers[0].Address)
				}
				oneTransmitter := []solana.PublicKey{transmitterPubKeys[0]}

				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					repeatedSignerAddresses,
					oneTransmitter,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorInvalidConfigRepeatedOracle.String()})
			})

			t.Run("It rejects zero transmitter address", func(t *testing.T) {
				t.Parallel()
				invalidTransmitterPubKeys := []solana.PublicKey{transmitterPubKeys[0], common.ZeroAddress}

				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					signerAddresses,
					invalidTransmitterPubKeys,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorOracleCannotBeZeroAddress.String()})
			})

			t.Run("It rejects zero signer address", func(t *testing.T) {
				t.Parallel()
				invalidSignerAddresses := [][20]byte{{}}
				for _, v := range signers[1:] {
					invalidSignerAddresses = append(invalidSignerAddresses, v.Address)
				}
				instruction, err := ccip_router.NewSetOcrConfigInstruction(
					uint8(testutils.OcrCommitPlugin),
					ccip_router.Ocr3ConfigInfo{
						ConfigDigest: config.ConfigDigest,
						F:            config.OcrF,
					},
					invalidSignerAddresses,
					transmitterPubKeys,
					config.RouterConfigPDA,
					config.RouterStatePDA,
					ccipAdmin.PublicKey(),
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorOracleCannotBeZeroAddress.String()})
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
						token0.Mint.PublicKey(),
						user.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When transmitter wants to set up the token admin registry, it fails", func(t *testing.T) {
					transmitter := getTransmitter()
					instruction, err := ccip_router.NewCcipAdminProposeAdministratorInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						transmitter.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When admin wants to set up the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewCcipAdminProposeAdministratorInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						ccipAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token0.Mint.PublicKey(),
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When ccip admin wants to accept the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When proposed admin wants to accept the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
					token0.Mint.PublicKey(),
					ccipAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.InvalidTokenAdminRegistryProposedAdmin_CcipRouterError.String()})
			})

			t.Run("set pool", func(t *testing.T) {
				t.Run("When any user wants to set up the pool, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token0.PoolLookupTable,
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When transmitter wants to set up the pool, it fails", func(t *testing.T) {
					transmitter := getTransmitter()
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token0.PoolLookupTable,
						transmitter.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When admin wants to set up the pool, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token0.PoolLookupTable,
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When setting pool to incorrect addresses in lookup table, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token1.PoolLookupTable, // accounts do not match the expected mint related accounts
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When Token Pool Admin wants to set up the pool, it succeeds", func(t *testing.T) {
					base := ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token0.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					)

					base.AccountMetaSlice = append(base.AccountMetaSlice, solana.Meta(token0.PoolLookupTable))
					instruction, err := base.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token0.Mint.PublicKey(),
						solana.PublicKey{},
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token0.Mint.PublicKey(),
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
						token0.Mint.PublicKey(),
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When admin wants to transfer the token admin registry, it succeeds and permissions stay with no changes", func(t *testing.T) {
					instruction, err := ccip_router.NewTransferAdminRoleTokenAdminRegistryInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token0.Mint.PublicKey(),
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
						token0.Mint.PublicKey(),
						token0.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When new admin accepts the token admin registry, it succeeds and permissions are updated", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token0.Mint.PublicKey(),
						token0.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})

					// new one can make changes now
					instruction, err = ccip_router.NewSetPoolInstruction(
						token0.WritableIndexes,
						config.RouterConfigPDA,
						token0.AdminRegistryPDA,
						token0.Mint.PublicKey(),
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
						token1.Mint.PublicKey(),
						user.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When transmitter wants to set up the token admin registry, it fails", func(t *testing.T) {
					transmitter := getTransmitter()
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						transmitter.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When ccip admin wants to set up the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						ccipAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When token mint_authority wants to set up the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewOwnerProposeAdministratorInstruction(
						token1PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						token1PoolAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
					err = common.GetAccountDataBorshInto(ctx, solanaGoClient, token1.AdminRegistryPDA, config.DefaultCommitment, &tokenAdminRegistry)
					require.NoError(t, err)
					require.Equal(t, uint8(1), tokenAdminRegistry.Version)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.Administrator)
					require.Equal(t, token1PoolAdmin.PublicKey(), tokenAdminRegistry.PendingAdministrator)
					require.Equal(t, solana.PublicKey{}, tokenAdminRegistry.LookupTable)
				})
			})

			t.Run("accept token admin registry as token admin", func(t *testing.T) {
				t.Run("When any user wants to accept the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						user.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When ccip admin wants to accept the token admin registry, it fails", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						ccipAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When proposed admin wants to accept the token admin registry, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
					token1.Mint.PublicKey(),
					token1PoolAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment, []string{ccip_router.InvalidTokenAdminRegistryProposedAdmin_CcipRouterError.String()})
			})

			t.Run("set pool", func(t *testing.T) {
				t.Run("When Mint Authority wants to set up the pool, it succeeds", func(t *testing.T) {
					instruction, err := ccip_router.NewSetPoolInstruction(
						token1.WritableIndexes,
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						token1.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					// transfer mint authority to pool once admin registry is set
					ixAuth, err := tokens.SetTokenMintAuthority(token1.Program, token1.PoolSigner, token1.Mint.PublicKey(), token1PoolAdmin.PublicKey()) // TODO: Check this
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction, ixAuth}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token1.Mint.PublicKey(),
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})
				t.Run("When mint authority wants to transfer the token admin registry, it succeeds and permissions stay with no changes", func(t *testing.T) {
					instruction, err := ccip_router.NewTransferAdminRoleTokenAdminRegistryInstruction(
						token0PoolAdmin.PublicKey(),
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token1.Mint.PublicKey(),
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
						token1.Mint.PublicKey(),
						token1.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
				})

				t.Run("When new admin accepts the token admin registry, it succeeds and permissions are updated", func(t *testing.T) {
					instruction, err := ccip_router.NewAcceptAdminRoleTokenAdminRegistryInstruction(
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)

					// Validate Token Pool Registry PDA
					tokenAdminRegistry := ccip_router.TokenAdminRegistry{}
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
						token1.Mint.PublicKey(),
						token1.PoolLookupTable,
						token1PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token1PoolAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})

					// new one can make changes now
					instruction, err = ccip_router.NewSetPoolInstruction(
						token1.WritableIndexes,
						config.RouterConfigPDA,
						token1.AdminRegistryPDA,
						token1.Mint.PublicKey(),
						token1.PoolLookupTable,
						token0PoolAdmin.PublicKey(),
					).ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, token0PoolAdmin, config.DefaultCommitment)
				})
			})
		})
	})

	//////////////////////////////
	// Token Pool Config Tests //
	/////////////////////////////
	t.Run("Token Pool Configuration", func(t *testing.T) {
		t.Run("RemoteConfig", func(t *testing.T) {
			ix0, err := token_pool.NewInitChainRemoteConfigInstruction(config.EvmChainSelector, token0.Mint.PublicKey(), token_pool.RemoteConfig{
				// PoolAddresses: []token_pool.RemoteAddress{{Address: []byte{1, 2, 3}}},
				TokenAddress: token_pool.RemoteAddress{Address: []byte{1, 2, 3}},
			}, token0.PoolConfig, token0.Chain[config.EvmChainSelector], token0PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix1, err := token_pool.NewInitChainRemoteConfigInstruction(config.EvmChainSelector, token1.Mint.PublicKey(), token_pool.RemoteConfig{
				PoolAddresses: []token_pool.RemoteAddress{{Address: []byte{4, 5, 6}}},
				TokenAddress:  token_pool.RemoteAddress{Address: []byte{4, 5, 6}},
			}, token1.PoolConfig, token1.Chain[config.EvmChainSelector], token1PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix0, ix1}, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token1PoolAdmin))
		})

		t.Run("AppendRemotePools", func(t *testing.T) {
			ix, err := token_pool.NewAppendRemotePoolAddressesInstruction(config.EvmChainSelector, token0.Mint.PublicKey(), []token_pool.RemoteAddress{{Address: []byte{1, 2, 3}}},
				token0.PoolConfig, token0.Chain[config.EvmChainSelector], token0PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, token0PoolAdmin, config.DefaultCommitment)
		})

		t.Run("RateLimit", func(t *testing.T) {
			ix0, err := token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, token0.Mint.PublicKey(), token_pool.RateLimitConfig{}, token_pool.RateLimitConfig{}, token0.PoolConfig, token0.Chain[config.EvmChainSelector], token0PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix1, err := token_pool.NewSetChainRateLimitInstruction(config.EvmChainSelector, token1.Mint.PublicKey(), token_pool.RateLimitConfig{}, token_pool.RateLimitConfig{}, token1.PoolConfig, token1.Chain[config.EvmChainSelector], token1PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix0, ix1}, token0PoolAdmin, config.DefaultCommitment, common.AddSigners(token1PoolAdmin))
		})

		t.Run("Billing", func(t *testing.T) {
			ix0, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token0.Mint.PublicKey(), fee_quoter.TokenTransferFeeConfig{}, config.FqConfigPDA, token0.Billing[config.EvmChainSelector], ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			ix1, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token1.Mint.PublicKey(), fee_quoter.TokenTransferFeeConfig{}, config.FqConfigPDA, token1.Billing[config.EvmChainSelector], ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix0, ix1}, ccipAdmin, config.DefaultCommitment)
		})

		// validate permissions for setting config
		t.Run("Permissions", func(t *testing.T) {
			t.Parallel()
			t.Run("Billing can only be set by CCIP admin", func(t *testing.T) {
				ix, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token0.Mint.PublicKey(), fee_quoter.TokenTransferFeeConfig{}, config.FqConfigPDA, token0.Billing[config.EvmChainSelector], token1PoolAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, token1PoolAdmin, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
			})
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				TokenAmounts: []fee_quoter.SVMTokenAmount{{Token: token0.Mint.PublicKey(), Amount: 1}},
				ExtraArgs:    emptyEVMExtraArgsV2,
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
			token0BillingConfigPda := getFqTokenConfigPDA(token0.Mint.PublicKey())
			token0PerChainPerConfigPda := getFqPerChainPerTokenConfigBillingPDA(token0.Mint.PublicKey())
			ix, err := fee_quoter.NewSetTokenTransferFeeConfigInstruction(config.EvmChainSelector, token0.Mint.PublicKey(), billing, config.FqConfigPDA, token0PerChainPerConfigPda, ccipAdmin.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment)

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
					ExtraArgs: emptyEVMExtraArgsV2,
				}

				raw := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, message, config.FqConfigPDA, config.FqEvmDestChainPDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA)
				instruction, err := raw.ValidateAndBuild()
				require.NoError(t, err)

				result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: InvalidEVMAddress"})
				require.NotNil(t, result)
			}
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()

			require.NoError(t, err)
			result := testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: AccountNotInitialized"})
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
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
			require.Equal(t, bin.Uint128{Lo: uint64(validFqDestChainConfig.DefaultTxGasLimit), Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).GasLimit) // default gas limit
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).AllowOutOfOrderExecution)                                                    // default OOO Execution
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
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.EVMExtraArgsV2{
					GasLimit:                 bin.Uint128{Lo: 99, Hi: 0},
					AllowOutOfOrderExecution: trueValue,
				}, ccip.EVMExtraArgsV2Tag),
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
				config.ExternalTokenPoolsSignerPDA,
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
			require.Equal(t, bin.Uint128{Lo: 99, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).GasLimit) // check it's overwritten
			require.Equal(t, true, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).AllowOutOfOrderExecution)       // check it's overwritten
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
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.EVMExtraArgsV2{
					GasLimit: bin.Uint128{Lo: 99, Hi: 0},
				}, ccip.EVMExtraArgsV2Tag),
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
				config.ExternalTokenPoolsSignerPDA,
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
			require.Equal(t, bin.Uint128{Lo: 99, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).GasLimit) // check it's overwritten
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).AllowOutOfOrderExecution)      // check it's default value
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
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.EVMExtraArgsV2{
					AllowOutOfOrderExecution: true,
					GasLimit:                 bin.Uint128{Lo: 5000, Hi: 0},
				}, ccip.EVMExtraArgsV2Tag),
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
				config.ExternalTokenPoolsSignerPDA,
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
			require.Equal(t, bin.Uint128{Lo: 5000, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).GasLimit) // default gas limit
			require.Equal(t, true, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).AllowOutOfOrderExecution)         // check it's overwritten
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
				ExtraArgs: testutils.MustSerializeExtraArgs(t, fee_quoter.EVMExtraArgsV2{
					GasLimit: bin.Uint128{Lo: 0, Hi: 0},
				}, ccip.EVMExtraArgsV2Tag),
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
				config.ExternalTokenPoolsSignerPDA,
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
			require.Equal(t, bin.Uint128{Lo: 0, Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).GasLimit)
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).AllowOutOfOrderExecution)
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
			)

			// do NOT mark the user ATA as writable

			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndFailWithRPCError(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.InvalidInputsAtaWritable_CcipRouterError.String()})
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
												ExtraArgs: emptyEVMExtraArgsV2,
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
											config.ExternalTokenPoolsSignerPDA,
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
			)
			raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
			instruction, err := raw.ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, anotherUser, config.DefaultCommitment, []string{ccip_router.InvalidInputs_CcipRouterError.String()})
		})

		t.Run("When another user sending a Valid CCIP Message Emits CCIPMessageSent", func(t *testing.T) {
			destinationChainSelector := config.EvmChainSelector
			destinationChainStatePDA := config.EvmDestChainStatePDA
			message := ccip_router.SVM2AnyMessage{
				FeeToken:  link22.mint,
				Receiver:  validReceiverAddress[:],
				Data:      []byte{4, 5, 6},
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
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
			require.Equal(t, bin.Uint128{Lo: uint64(validFqDestChainConfig.DefaultTxGasLimit), Hi: 0}, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).GasLimit)
			require.Equal(t, false, testutils.MustDeserializeExtraArgs(t, &fee_quoter.EVMExtraArgsV2{}, ccipMessageSentEvent.Message.ExtraArgs, ccip.EVMExtraArgsV2Tag).AllowOutOfOrderExecution)
			require.Equal(t, uint64(15), ccipMessageSentEvent.Message.Header.SourceChainSelector)
			require.Equal(t, uint64(21), ccipMessageSentEvent.Message.Header.DestChainSelector)
			require.Equal(t, uint64(6), ccipMessageSentEvent.Message.Header.SequenceNumber)
			require.Equal(t, uint64(1), ccipMessageSentEvent.Message.Header.Nonce)
		})

		t.Run("token happy path", func(t *testing.T) {
			t.Run("single token", func(t *testing.T) {
				_, initSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint.PublicKey(), config.DefaultCommitment)
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
							Token:  token0.Mint.PublicKey(),
							Amount: 1,
						},
					},
					ExtraArgs: emptyEVMExtraArgsV2,
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
					config.ExternalTokenPoolsSignerPDA,
				)
				base.GetFeeTokenUserAssociatedAccountAccount().WRITE()

				tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, userTokenAccount)
				require.NoError(t, err)
				base.AccountMetaSlice = append(base.AccountMetaSlice, tokenMetas...)

				ix, err := base.ValidateAndBuild()
				require.NoError(t, err)

				ixApprove, err := tokens.TokenApproveChecked(1, 0, token0.Program, userTokenAccount, token0.Mint.PublicKey(), config.ExternalTokenPoolsSignerPDA, user.PublicKey(), nil)
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
				// Since they have been configured in the test setup to differ by a factor of 10, so does the token amount and its value in juels
				require.Equal(t, tokens.ToLittleEndianU256(3633302), ccipMessageSentEvent.Message.FeeValueJuels.LeBytes)
				require.Equal(t, token0.PoolConfig, ta.SourcePoolAddress)
				require.Equal(t, []byte{1, 2, 3}, ta.DestTokenAddress)
				require.Equal(t, 0, len(ta.ExtraData))
				require.Equal(t, tokens.ToLittleEndianU256(1), ta.Amount.LeBytes)
				require.Equal(t, 32, len(ta.DestExecData))
				expectedDestExecData := make([]byte, 32)
				// Token0 billing had DestGasOverhead set to 100 during setup
				binary.BigEndian.PutUint64(expectedDestExecData[24:], 100)
				require.Equal(t, expectedDestExecData, ta.DestExecData)

				// check pool event
				poolEvent := tokens.EventBurnLock{}
				require.NoError(t, common.ParseEvent(result.Meta.LogMessages, "Burned", &poolEvent, config.PrintEvents))
				require.Equal(t, config.ExternalTokenPoolsSignerPDA, poolEvent.Sender)
				require.Equal(t, uint64(1), poolEvent.Amount)

				// check balances
				_, currSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint.PublicKey(), config.DefaultCommitment)
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
							Token:  token0.Mint.PublicKey(),
							Amount: 1,
						},
						{
							Token:  token1.Mint.PublicKey(),
							Amount: 2,
						},
					},
					ExtraArgs: emptyEVMExtraArgsV2,
				}

				userTokenAccount0, ok := token0.User[user.PublicKey()]
				require.True(t, ok)
				userTokenAccount1, ok := token1.User[user.PublicKey()]
				require.True(t, ok)

				base := ccip_router.NewCcipSendInstruction(
					destinationChainSelector,
					message,
					[]byte{0, 13}, // starting indices for accounts
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
					config.ExternalTokenPoolsSignerPDA,
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

				ixApprove0, err := tokens.TokenApproveChecked(1, 0, token0.Program, userTokenAccount0, token0.Mint.PublicKey(), config.ExternalTokenPoolsSignerPDA, user.PublicKey(), nil)
				require.NoError(t, err)
				ixApprove1, err := tokens.TokenApproveChecked(2, 0, token1.Program, userTokenAccount1, token1.Mint.PublicKey(), config.ExternalTokenPoolsSignerPDA, user.PublicKey(), nil)
				require.NoError(t, err)

				result := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ixApprove0, ixApprove1, ix}, user, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(400_000))
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
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
						Token:  token0.Mint.PublicKey(),
						Amount: 1,
					},
				},
				ExtraArgs: emptyEVMExtraArgsV2,
			}

			userTokenAccount, ok := token0.User[user.PublicKey()]
			require.True(t, ok)

			inputs := []struct {
				name        string
				index       uint
				replaceWith *solana.AccountMeta // default to zero address
				errorStr    ccip_router.CcipRouterError
			}{
				{
					name:     "incorrect user token account",
					index:    0,
					errorStr: ccip_router.InvalidInputsTokenAccounts_CcipRouterError,
				},
				{
					name:     "invalid billing config",
					index:    1,
					errorStr: ccip_router.InvalidInputsConfigAccounts_CcipRouterError,
				},
				{
					name:     "invalid token pool chain config",
					index:    2,
					errorStr: ccip_router.InvalidInputsConfigAccounts_CcipRouterError,
				},
				{
					name:     "pool config is not owned by pool program",
					index:    6,
					errorStr: ccip_router.InvalidInputsPoolAccounts_CcipRouterError,
				},
				{
					name:        "is pool config but for wrong token",
					index:       6,
					replaceWith: solana.Meta(token1.PoolConfig),
					errorStr:    ccip_router.InvalidInputsPoolAccounts_CcipRouterError,
				},
				{
					name:        "is pool config but missing write permissions",
					index:       6,
					replaceWith: solana.Meta(token0.PoolConfig),
					errorStr:    ccip_router.InvalidInputsLookupTableAccountWritable_CcipRouterError,
				},
				{
					name:        "is pool lookup table but has write permissions",
					index:       3,
					replaceWith: solana.Meta(token0.PoolLookupTable).WRITE(),
					errorStr:    ccip_router.InvalidInputsLookupTableAccountWritable_CcipRouterError,
				},
				{
					name:     "incorrect pool signer",
					index:    8,
					errorStr: ccip_router.InvalidInputsPoolAccounts_CcipRouterError,
				},
				{
					name:     "invalid token program",
					index:    9,
					errorStr: ccip_router.InvalidInputsTokenAccounts_CcipRouterError,
				},
				{
					name:     "incorrect pool token account",
					index:    7,
					errorStr: ccip_router.InvalidInputsTokenAccounts_CcipRouterError,
				},
				{
					name:        "incorrect token pool lookup table",
					index:       3,
					replaceWith: solana.Meta(token1.PoolLookupTable),
					errorStr:    ccip_router.InvalidInputsLookupTableAccounts_CcipRouterError,
				},
				{
					name:     "invalid fee token config",
					index:    11,
					errorStr: ccip_router.InvalidInputsConfigAccounts_CcipRouterError,
				},
				{
					name:     "extra accounts not in lookup table",
					index:    1_000, // large number to indicate append
					errorStr: ccip_router.InvalidInputsLookupTableAccounts_CcipRouterError,
				},
				{
					name:     "remaining accounts mismatch",
					index:    12, // only works with token0
					errorStr: ccip_router.InvalidInputsLookupTableAccounts_CcipRouterError,
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
						config.ExternalTokenPoolsSignerPDA,
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

					testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, addressTables, []string{in.errorStr.String()})
				})
			}
		})

		decodeGetFeeResult := func(t *testing.T) func([]byte) fee_quoter.GetFeeResult {
			return func(bytes []byte) fee_quoter.GetFeeResult {
				result := fee_quoter.GetFeeResult{}
				decoder := ag_binary.NewBinDecoder(bytes)
				err := result.UnmarshalWithDecoder(decoder)
				require.NoError(t, err)
				return result
			}
		}

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
						ExtraArgs: emptyEVMExtraArgsV2,
					}
					fqMsg := toFqMsg(message)
					rawGetFeeIx := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, fqMsg, config.FqConfigPDA, config.FqEvmDestChainPDA, token.fqBillingConfigPDA, link22.fqBillingConfigPDA)
					ix, err := rawGetFeeIx.ValidateAndBuild()
					require.NoError(t, err)

					feeResult := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment)
					require.NotNil(t, feeResult)
					fmt.Println(feeResult.Meta.LogMessages)
					fee, _ := common.ExtractTypedReturnValue(ctx, feeResult.Meta.LogMessages, config.FeeQuoterProgram.String(), decodeGetFeeResult(t))
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
						config.ExternalTokenPoolsSignerPDA,
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
				ExtraArgs: emptyEVMExtraArgsV2,
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
				config.ExternalTokenPoolsSignerPDA,
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
				ExtraArgs: emptyEVMExtraArgsV2,
			}
			fqMsg := toFqMsg(message)

			// getFee
			rawGetFeeIx := fee_quoter.NewGetFeeInstruction(config.EvmChainSelector, fqMsg, config.FqConfigPDA, config.FqEvmDestChainPDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA)
			ix, err := rawGetFeeIx.ValidateAndBuild()
			require.NoError(t, err)

			feeResult := testutils.SimulateTransaction(ctx, t, solanaGoClient, []solana.Instruction{ix}, user)
			require.NotNil(t, feeResult)
			var fee fee_quoter.GetFeeResult
			fee, err = common.ExtractTypedReturnValue(ctx, feeResult.Value.Logs, config.FeeQuoterProgram.String(), decodeGetFeeResult(t))
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
				config.ExternalTokenPoolsSignerPDA,
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
		})
	})

	///////////////////////////
	// Withdraw billed funds //
	////////////////////.......
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

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, user, config.DefaultCommitment, []string{ccip_router.Unauthorized_CcipRouterError.String()})
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

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.InvalidInputs_CcipRouterError.String()})
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

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{"Error Code: ConstraintTokenOwner. Error Number: 2015"})
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

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.InvalidInputs_CcipRouterError.String()})
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

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.InsufficientFunds_CcipRouterError.String()})
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

			testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{ix}, ccipAdmin, config.DefaultCommitment, []string{ccip_router.InsufficientFunds_CcipRouterError.String()})
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
					PriceUpdates            ccip_router.PriceUpdates
					RemainingAccounts       []solana.PublicKey
					RunEventValidations     func(t *testing.T, tx *rpc.GetTransactionResult)
					RunStateValidations     func(t *testing.T)
					ReportContext           *[2][32]byte
					PriceSequenceComparator Comparator
				}{
					{
						Name:              "No price updates",
						PriceUpdates:      ccip_router.PriceUpdates{},
						RemainingAccounts: []solana.PublicKey{},
						RunEventValidations: func(t *testing.T, tx *rpc.GetTransactionResult) {
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerTokenUpdated", nil, config.PrintEvents), "event not found")
							require.ErrorContains(t, common.ParseEvent(tx.Meta.LogMessages, "UsdPerUnitGasUpdated", nil, config.PrintEvents), "event not found")
						},
						RunStateValidations:     func(t *testing.T) {},
						PriceSequenceComparator: Greater, // it is a newer commit but with no price update
					},
					{
						Name: "Single token price update",
						PriceUpdates: ccip_router.PriceUpdates{
							TokenPriceUpdates: []ccip_router.TokenPriceUpdate{{
								SourceToken: wsol.mint,
								UsdPerToken: common.To28BytesBE(1),
							}},
						},
						RemainingAccounts: []solana.PublicKey{config.RouterStatePDA, wsol.fqBillingConfigPDA},
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
						PriceUpdates: ccip_router.PriceUpdates{
							GasPriceUpdates: []ccip_router.GasPriceUpdate{{
								DestChainSelector: config.EvmChainSelector,
								UsdPerUnitGas:     common.To28BytesBE(1),
							}},
						},
						RemainingAccounts: []solana.PublicKey{config.RouterStatePDA, config.FqEvmDestChainPDA},
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
						PriceUpdates: ccip_router.PriceUpdates{
							GasPriceUpdates: []ccip_router.GasPriceUpdate{{
								DestChainSelector: config.SvmChainSelector,
								UsdPerUnitGas:     common.To28BytesBE(2),
							}},
						},
						RemainingAccounts: []solana.PublicKey{config.RouterStatePDA, config.FqSvmDestChainPDA},
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
						PriceUpdates: ccip_router.PriceUpdates{
							TokenPriceUpdates: []ccip_router.TokenPriceUpdate{
								{SourceToken: wsol.mint, UsdPerToken: common.To28BytesBE(3)},
								{SourceToken: link22.mint, UsdPerToken: common.To28BytesBE(4)},
							},
							GasPriceUpdates: []ccip_router.GasPriceUpdate{
								{DestChainSelector: config.EvmChainSelector, UsdPerUnitGas: common.To28BytesBE(5)},
								{DestChainSelector: config.SvmChainSelector, UsdPerUnitGas: common.To28BytesBE(6)},
							},
						},
						RemainingAccounts: []solana.PublicKey{config.RouterStatePDA, wsol.fqBillingConfigPDA, link22.fqBillingConfigPDA, config.FqEvmDestChainPDA, config.FqSvmDestChainPDA},
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
						PriceUpdates: ccip_router.PriceUpdates{
							TokenPriceUpdates: []ccip_router.TokenPriceUpdate{
								{SourceToken: wsol.mint, UsdPerToken: common.To28BytesBE(1)},
							},
							GasPriceUpdates: []ccip_router.GasPriceUpdate{
								{DestChainSelector: config.EvmChainSelector, UsdPerUnitGas: common.To28BytesBE(1)},
							},
						},
						RemainingAccounts: []solana.PublicKey{config.RouterStatePDA, wsol.fqBillingConfigPDA, config.EvmDestChainStatePDA},
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
						rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
						require.NoError(t, err)

						minV := currentMinSeqNr
						maxV := currentMinSeqNr + sequenceLength - 1

						currentMinSeqNr = maxV + 1 // advance the outer sequence counter

						report := ccip_router.CommitInput{
							MerkleRoot: ccip_router.MerkleRoot{
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

						raw := ccip_router.NewCommitInstruction(
							reportContext,
							testutils.MustMarshalBorsh(t, report),
							sigs.Rs,
							sigs.Ss,
							sigs.RawVs,
							config.RouterConfigPDA,
							config.EvmSourceChainStatePDA,
							rootPDA,
							transmitter.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.BillingSignerPDA,
							config.FeeQuoterProgram,
							config.FqConfigPDA,
						)

						for _, pubkey := range testcase.RemainingAccounts {
							raw.AccountMetaSlice.Append(solana.Meta(pubkey).WRITE())
						}

						instruction, err := raw.ValidateAndBuild()
						require.NoError(t, err)
						tx := testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, rpc.CommitmentConfirmed, commitLookupTable, common.AddComputeUnitLimit(MaxCU))

						commitEvent := ccip.EventCommitReportAccepted{}
						require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &commitEvent, config.PrintEvents))
						require.Equal(t, config.EvmChainSelector, commitEvent.Report.SourceChainSelector)
						require.Equal(t, root, commitEvent.Report.MerkleRoot)
						require.Equal(t, minV, commitEvent.Report.MinSeqNr)
						require.Equal(t, maxV, commitEvent.Report.MaxSeqNr)

						transmittedEvent := ccip.EventTransmitted{}
						require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Transmitted", &transmittedEvent, config.PrintEvents))
						require.Equal(t, config.ConfigDigest, transmittedEvent.ConfigDigest)
						require.Equal(t, uint8(testutils.OcrCommitPlugin), transmittedEvent.OcrPluginType)
						require.Equal(t, reportSequence, transmittedEvent.SequenceNumber)

						var chainStateAccount ccip_router.SourceChain
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmSourceChainStatePDA, config.DefaultCommitment, &chainStateAccount)
						require.NoError(t, err, "failed to get account info")
						require.Equal(t, currentMinSeqNr, chainStateAccount.State.MinSeqNr) // state now holds the "advanced outer" sequence number, which is the minimum for the next report
						// Do not check dest chain config, as it may have been updated by other tests in ccip onramp

						var rootAccount ccip_router.CommitReport
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
						require.NoError(t, err, "failed to get account info")
						require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)

						var globalState ccip_router.GlobalState
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterStatePDA, config.DefaultCommitment, &globalState)
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
			})

			t.Run("Edge cases", func(t *testing.T) {
				t.Run("When committing a report with an invalid source chain selector it fails", func(t *testing.T) {
					t.Parallel()
					sourceChainSelector := uint64(34)
					sourceChainStatePDA, err := state.FindSourceChainStatePDA(sourceChainSelector, config.CcipRouterProgram)
					require.NoError(t, err)
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, sourceChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
					rootPDA, err := state.FindCommitReportPDA(sourceChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr + 4

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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
					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						sourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: AccountNotInitialized"})
				})

				t.Run("When committing a report with an invalid interval it fails", func(t *testing.T) {
					t.Parallel()
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr - 2 // max lower than min

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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
					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidSequenceInterval_CcipRouterError.String()})
				})

				t.Run("When committing a report with an interval size bigger than supported it fails", func(t *testing.T) {
					t.Parallel()
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, []solana.PublicKey{})
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr + 65 // max - min > 64

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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
					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidSequenceInterval_CcipRouterError.String()})
				})

				t.Run("When committing a report with a zero merkle root it fails", func(t *testing.T) {
					t.Parallel()
					root := [32]byte{}
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr // max = min

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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
					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidProof_CcipRouterError.String()})
				})

				t.Run("When committing a report with a repeated merkle root, it fails", func(t *testing.T) {
					t.Parallel()
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{1, 2, 3, 1}, msgAccounts) // repeated root
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := currentMinSeqNr
					maxV := currentMinSeqNr + 4

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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
					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment,
						[]string{"Allocate: account Address", "already in use", "failed: custom program error: 0x0"})
				})

				t.Run("When committing a report with an invalid min interval, it fails", func(t *testing.T) {
					t.Parallel()
					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := uint64(8) // this is lower than expected
					maxV := uint64(10)

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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
					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip_router.InvalidSequenceInterval_CcipRouterError.String()})
				})

				t.Run("Invalid price updates", func(t *testing.T) {
					randomToken := solana.MustPublicKeyFromBase58("AGDpGy7auzgKT8zt6qhfHFm1rDwvqQGGTYxuYn7MtydQ") // just some non-existing token

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
							Name:             "with a price update for a token that doesn't exist",
							Tokens:           []solana.PublicKey{randomToken},
							AccountMetaSlice: solana.AccountMetaSlice{solana.Meta(getFqTokenConfigPDA(randomToken)).WRITE()},
							ExpectedError:    "AccountNotInitialized",
						},
						{
							Name:              "with a gas price update for a chain that doesn't exist",
							GasChainSelectors: []uint64{randomChain},
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(randomChainPDA).WRITE()},
							ExpectedError:     "AccountNotInitialized",
						},
						{
							Name:             "with a non-writable billing token config account",
							Tokens:           []solana.PublicKey{wsol.mint},
							AccountMetaSlice: solana.AccountMetaSlice{solana.Meta(wsol.fqBillingConfigPDA)}, // not writable
							ExpectedError:    ccip_router.InvalidInputs_CcipRouterError.String(),
						},
						{
							// when the message source chain is the same as the chain whose gas is updated, then the same chain state is passed
							// in twice, in which case the resulting permissions are the sum of both instances. As only one is manually constructed here,
							// the other one is always writable (handled by the auto-generated code).
							Name:              "with a non-writable chain state account (different from the message source chain)",
							GasChainSelectors: []uint64{config.SvmChainSelector},                                 // the message source chain is EVM
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(config.SvmDestChainStatePDA)}, // not writable
							ExpectedError:     ccip_router.InvalidInputs_CcipRouterError.String(),
						},
						{
							Name:             "with the wrong billing token config account for a valid token",
							Tokens:           []solana.PublicKey{wsol.mint},
							AccountMetaSlice: solana.AccountMetaSlice{solana.Meta(link22.fqBillingConfigPDA).WRITE()}, // mismatch token
							ExpectedError:    ccip_router.InvalidInputs_CcipRouterError.String(),
						},
						{
							Name:              "with the wrong chain state account for a valid gas update",
							GasChainSelectors: []uint64{config.SvmChainSelector},
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(config.EvmDestChainStatePDA).WRITE()}, // mismatch chain
							ExpectedError:     ccip_router.InvalidInputs_CcipRouterError.String(),
						},
						{
							Name:              "with too few accounts",
							Tokens:            []solana.PublicKey{wsol.mint},
							GasChainSelectors: []uint64{config.EvmChainSelector},
							AccountMetaSlice:  solana.AccountMetaSlice{solana.Meta(wsol.fqBillingConfigPDA).WRITE()}, // missing chain state account
							ExpectedError:     ccip_router.InvalidInputs_CcipRouterError.String(),
						},
						// TODO right now I'm allowing sending too many remaining_accounts, but if we want to be restrictive with that we can add a test here
					}

					msgAccounts := []solana.PublicKey{}
					_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{1, 2, 3}, msgAccounts)
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					for _, testcase := range testcases {
						t.Run(testcase.Name, func(t *testing.T) {
							t.Parallel()

							priceUpdates := ccip_router.PriceUpdates{
								TokenPriceUpdates: make([]ccip_router.TokenPriceUpdate, len(testcase.Tokens)),
								GasPriceUpdates:   make([]ccip_router.GasPriceUpdate, len(testcase.GasChainSelectors)),
							}
							for i, token := range testcase.Tokens {
								priceUpdates.TokenPriceUpdates[i] = ccip_router.TokenPriceUpdate{
									SourceToken: token,
									UsdPerToken: common.To28BytesBE(uint64(i)),
								}
							}
							for i, chainSelector := range testcase.GasChainSelectors {
								priceUpdates.GasPriceUpdates[i] = ccip_router.GasPriceUpdate{
									DestChainSelector: chainSelector,
									UsdPerUnitGas:     common.To28BytesBE(uint64(i)),
								}
							}

							transmitter := getTransmitter()

							report := ccip_router.CommitInput{
								MerkleRoot: ccip_router.MerkleRoot{
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

							raw := ccip_router.NewCommitInstruction(
								reportContext,
								testutils.MustMarshalBorsh(t, report),
								sigs.Rs,
								sigs.Ss,
								sigs.RawVs,
								config.RouterConfigPDA,
								config.EvmSourceChainStatePDA,
								rootPDA,
								transmitter.PublicKey(),
								solana.SystemProgramID,
								solana.SysVarInstructionsPubkey,
								config.BillingSignerPDA,
								config.FeeQuoterProgram,
								config.FqConfigPDA,
							)

							raw.AccountMetaSlice.Append(solana.Meta(config.RouterStatePDA).WRITE())
							for _, meta := range testcase.AccountMetaSlice {
								raw.AccountMetaSlice.Append(meta)
							}

							instruction, err := raw.ValidateAndBuild()
							require.NoError(t, err)
							testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, rpc.CommitmentConfirmed, commitLookupTable, []string{testcase.ExpectedError}, common.AddComputeUnitLimit(MaxCU))
						})
					}
				})
			})

			t.Run("When committing a report with the exact next interval, it succeeds", func(t *testing.T) {
				msgAccounts := []solana.PublicKey{}
				_, root := testutils.MakeAnyToSVMMessage(t, config.CcipTokenReceiver, config.EvmChainSelector, config.SvmChainSelector, []byte{4, 5, 6}, msgAccounts)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				minV := currentMinSeqNr
				maxV := currentMinSeqNr + 4

				currentMinSeqNr = maxV + 1 // advance the outer sequence counter as this will succeed

				report := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
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
				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, report),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				commitEvent := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &commitEvent, config.PrintEvents))
				require.Equal(t, config.EvmChainSelector, commitEvent.Report.SourceChainSelector)
				require.Equal(t, root, commitEvent.Report.MerkleRoot)
				require.Equal(t, minV, commitEvent.Report.MinSeqNr)
				require.Equal(t, maxV, commitEvent.Report.MaxSeqNr)

				transmittedEvent := ccip.EventTransmitted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Transmitted", &transmittedEvent, config.PrintEvents))
				require.Equal(t, config.ConfigDigest, transmittedEvent.ConfigDigest)
				require.Equal(t, uint8(testutils.OcrCommitPlugin), transmittedEvent.OcrPluginType)
				require.Equal(t, ccip.ReportSequence(), transmittedEvent.SequenceNumber)

				var chainStateAccount ccip_router.SourceChain
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.EvmSourceChainStatePDA, config.DefaultCommitment, &chainStateAccount)
				require.NoError(t, err, "failed to get account info")
				require.Equal(t, currentMinSeqNr, chainStateAccount.State.MinSeqNr)
				// Do not check dest chain config, as it may have been updated by other tests in ccip onramp

				var rootAccount ccip_router.CommitReport
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
				require.NoError(t, err, "failed to get account info")
				require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
			})

			t.Run("Ocr3Base::Transmit edge cases", func(t *testing.T) {
				t.Run("It rejects mismatch config digest", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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

					instruction, err := ccip_router.NewCommitInstruction(
						emptyReportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorConfigDigestMismatch.String()})
				})

				t.Run("It rejects unauthorized transmitter", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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

					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						user.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorUnauthorizedTransmitter.String()})
				})

				t.Run("It rejects incorrect signature count", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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

					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorWrongNumberOfSignatures.String()})
				})

				t.Run("It rejects invalid signature", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
							SourceChainSelector: config.EvmChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            minV,
							MaxSeqNr:            maxV,
							MerkleRoot:          root,
						},
					}
					sigs := ccip.Signatures{}
					transmitter := getTransmitter()

					instruction, err := ccip_router.NewCommitInstruction(
						ccip.NextCommitReportContext(),
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorWrongNumberOfSignatures.String()})
				})

				t.Run("It rejects unauthorized signer", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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

					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorUnauthorizedSigner.String()})
				})

				t.Run("It rejects duplicate signatures", func(t *testing.T) {
					t.Parallel()
					msg, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					minV := msg.Header.SequenceNumber
					maxV := msg.Header.SequenceNumber

					report := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
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

					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, report),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip.Ocr3ErrorNonUniqueSignatures.String()}, common.AddComputeUnitLimit(210_000))
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

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(210_000)) // signature verification compute unit amounts can vary depending on sorting
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
				executionEvent := ccip.EventExecutionStateChanged{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ExecutionStateChanged", &executionEvent, config.PrintEvents))

				require.NoError(t, err)
				require.NotNil(t, executionEvent)
				require.Equal(t, config.EvmChainSelector, executionEvent.SourceChainSelector)
				require.Equal(t, sequenceNumber, executionEvent.SequenceNumber)
				require.Equal(t, hex.EncodeToString(message.Header.MessageId[:]), hex.EncodeToString(executionEvent.MessageID[:]))
				require.Equal(t, hex.EncodeToString(root[:]), hex.EncodeToString(executionEvent.MessageHash[:]))
				require.Equal(t, ccip_router.Success_MessageExecutionState, executionEvent.State)

				var rootAccount ccip_router.CommitReport
				err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
				require.NoError(t, err, "failed to get account info")
				require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
				require.Equal(t, bin.Uint128{Lo: 2, Hi: 0}, rootAccount.ExecutionStates)
				require.Equal(t, sequenceNumber, rootAccount.MinMsgNr)
				require.Equal(t, sequenceNumber, rootAccount.MaxMsgNr)
			})

			t.Run("When executing a report with not matching source chain selector in message, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(210_000)) // signature verification compute unit amounts can vary depending on sorting
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				message.Header.SourceChainSelector = 89

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Message: Source chain selector not supported."})
			})

			t.Run("When executing a report with unsupported source chain selector account, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				unsupportedChainSelector := uint64(34)
				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID})
				sequenceNumber := message.Header.SequenceNumber

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(210_000)) // signature verification compute unit amounts can vary depending on sorting
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				unsupportedSourceChainStatePDA, err := state.FindSourceChainStatePDA(unsupportedChainSelector, config.CcipRouterProgram)
				require.NoError(t, err)
				unsupportedDestChainStatePDA, err := state.FindDestChainStatePDA(unsupportedChainSelector, config.CcipRouterProgram)
				require.NoError(t, err)
				message.Header.SourceChainSelector = unsupportedChainSelector
				message.Header.SequenceNumber = 1

				instruction, err = ccip_router.NewAddChainSelectorInstruction(
					unsupportedChainSelector,
					validSourceChainConfig,
					ccip_router.DestChainConfig{},
					unsupportedSourceChainStatePDA,
					unsupportedDestChainStatePDA,
					config.RouterConfigPDA,
					ccipAdmin.PublicKey(),
					solana.SystemProgramID,
				).ValidateAndBuild()
				require.NoError(t, err)
				result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
				require.NotNil(t, result)

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: unsupportedChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					unsupportedSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"AnchorError caused by account: commit_report. Error Code: ConstraintSeeds. Error Number: 2006. Error Message: A seeds constraint was violated."})
			})

			t.Run("When executing a report with incorrect solana chain selector, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message.Header.DestChainSelector = 89 // invalid dest chain selector
				sequenceNumber := message.Header.SequenceNumber
				hash, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)

				require.NoError(t, err)
				root := [32]byte(hash)

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Error Code: " + ccip_router.UnsupportedDestinationChainSelector_CcipRouterError.String()})
			})

			t.Run("When executing a report with nonexisting PDA for the Merkle Root, it fails", func(t *testing.T) {
				transmitter := getTransmitter()

				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
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

				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
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

				root := [32]byte(ccip.MerkleFrom([][]byte{hash1[:], hash2[:]}))

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            message1.Header.SequenceNumber,
						MaxSeqNr:            message2.Header.SequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				executionReport1 := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message2, // execute out of order
					Root:                root,
					Proofs:              [][32]uint8{hash1},
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport1),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
				executionEvent := ccip.EventExecutionStateChanged{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ExecutionStateChanged", &executionEvent, config.PrintEvents))

				executionReport2 := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message1,
					Root:                root,
					Proofs:              [][32]uint8{[32]byte(hash2)},
				}
				raw = ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport2),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
				executionEvent = ccip.EventExecutionStateChanged{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ExecutionStateChanged", &executionEvent, config.PrintEvents))

				require.NoError(t, err)
				require.NotNil(t, executionEvent)
				require.Equal(t, config.EvmChainSelector, executionEvent.SourceChainSelector)
				require.Equal(t, message1.Header.SequenceNumber, executionEvent.SequenceNumber)
				require.Equal(t, hex.EncodeToString(message1.Header.MessageId[:]), hex.EncodeToString(executionEvent.MessageID[:]))
				require.Equal(t, hex.EncodeToString(hash1[:]), hex.EncodeToString(executionEvent.MessageHash[:]))

				require.Equal(t, ccip_router.Success_MessageExecutionState, executionEvent.State)

				var rootAccount ccip_router.CommitReport
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

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipInvalidReceiverProgram, false, false),
					solana.NewAccountMeta(stubAccountPDA, false, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)

				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				// failed ccipReceiver - init account requires mutable authority
				// ccipSigner is not a mutable account
				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"writable privilege escalated", "Cross-program invocation with unauthorized signer or writable account"})
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

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(210_000)) // signature verification compute unit amounts can vary depending on sorting
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)

				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
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
					_, initSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint.PublicKey(), config.DefaultCommitment)
					require.NoError(t, err)
					_, initBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)

					transmitter := getTransmitter()

					sourceChainSelector := config.EvmChainSelector
					msgAccounts := []solana.PublicKey{config.CcipLogicReceiver, config.ReceiverExternalExecutionConfigPDA, config.ReceiverTargetAccountPDA, solana.SystemProgramID}
					message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
					message.TokenReceiver = config.ReceiverExternalExecutionConfigPDA
					message.TokenAmounts = []ccip_router.Any2SVMTokenTransfer{{
						SourcePoolAddress: []byte{1, 2, 3},
						DestTokenAddress:  token0.Mint.PublicKey(),
						Amount:            ccip_router.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1)},
					}}
					rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
					event := ccip.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

					executionReport := ccip_router.ExecutionReportSingleChain{
						SourceChainSelector: sourceChainSelector,
						Message:             message,
						OffchainTokenData:   [][]byte{{}},
						Root:                root,
						Proofs:              [][32]uint8{},
					}
					raw := ccip_router.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{4},
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						config.ExternalExecutionConfigPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.ExternalTokenPoolsSignerPDA,
					)
					raw.AccountMetaSlice = append(
						raw.AccountMetaSlice,
						solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
						solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
						solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
						solana.NewAccountMeta(solana.SystemProgramID, false, false),
					)

					tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas...)
					instruction, err = raw.ValidateAndBuild()
					require.NoError(t, err)

					tx = testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(300_000))
					executionEvent := ccip.EventExecutionStateChanged{}
					require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ExecutionStateChanged", &executionEvent, config.PrintEvents))
					require.Equal(t, ccip_router.Success_MessageExecutionState, executionEvent.State)

					mintEvent := tokens.EventMintRelease{}
					require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "Minted", &mintEvent, config.PrintEvents))
					require.Equal(t, config.ReceiverExternalExecutionConfigPDA, mintEvent.Recipient)
					require.Equal(t, token0.PoolSigner, mintEvent.Sender)
					require.Equal(t, uint64(1), mintEvent.Amount)

					_, finalSupply, err := tokens.TokenSupply(ctx, solanaGoClient, token0.Mint.PublicKey(), config.DefaultCommitment)
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
					message.ExtraArgs = ccip_router.Any2SVMRampExtraArgs{}
					message.Data = []byte{}
					message.TokenReceiver = config.ReceiverExternalExecutionConfigPDA
					message.TokenAmounts = []ccip_router.Any2SVMTokenTransfer{{
						SourcePoolAddress: []byte{1, 2, 3},
						DestTokenAddress:  token0.Mint.PublicKey(),
						Amount:            ccip_router.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1)},
					}, {
						SourcePoolAddress: []byte{4, 5, 6},
						DestTokenAddress:  token1.Mint.PublicKey(),
						Amount:            ccip_router.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(2)},
					}}
					rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
					require.NoError(t, err)

					root := [32]byte(rootBytes)
					sequenceNumber := message.Header.SequenceNumber

					commitReport := ccip_router.CommitInput{
						MerkleRoot: ccip_router.MerkleRoot{
							SourceChainSelector: sourceChainSelector,
							OnRampAddress:       config.OnRampAddress,
							MinSeqNr:            sequenceNumber,
							MaxSeqNr:            sequenceNumber,
							MerkleRoot:          root,
						},
					}
					sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
					require.NoError(t, err)
					rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
					require.NoError(t, err)

					instruction, err := ccip_router.NewCommitInstruction(
						reportContext,
						testutils.MustMarshalBorsh(t, commitReport),
						sigs.Rs,
						sigs.Ss,
						sigs.RawVs,
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.BillingSignerPDA,
						config.FeeQuoterProgram,
						config.FqConfigPDA,
					).ValidateAndBuild()
					require.NoError(t, err)
					tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
					event := ccip.EventCommitReportAccepted{}
					require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

					executionReport := ccip_router.ExecutionReportSingleChain{
						SourceChainSelector: sourceChainSelector,
						Message:             message,
						OffchainTokenData:   [][]byte{{}, {}},
						Root:                root,
						Proofs:              [][32]uint8{},
					}
					raw := ccip_router.NewExecuteInstruction(
						testutils.MustMarshalBorsh(t, executionReport),
						reportContext,
						[]byte{0, 13},
						config.RouterConfigPDA,
						config.EvmSourceChainStatePDA,
						rootPDA,
						config.ExternalExecutionConfigPDA,
						transmitter.PublicKey(),
						solana.SystemProgramID,
						solana.SysVarInstructionsPubkey,
						config.ExternalTokenPoolsSignerPDA,
					)

					tokenMetas0, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas0...)
					tokenMetas1, addressTables1, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token1, token1.User[config.ReceiverExternalExecutionConfigPDA])
					require.NoError(t, err)
					raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas1...)
					maps.Copy(addressTables, addressTables1)
					maps.Copy(addressTables, commitLookupTable) // commonly used ccip addresses - required otherwise tx is too large

					instruction, err = raw.ValidateAndBuild()
					require.NoError(t, err)

					testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables, common.AddComputeUnitLimit(300_000))

					// validate amounts
					_, finalBal0, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, 1, finalBal0-initBal0)
					_, finalBal1, err := tokens.TokenBalance(ctx, solanaGoClient, token1.User[config.ReceiverExternalExecutionConfigPDA], config.DefaultCommitment)
					require.NoError(t, err)
					require.Equal(t, 2, finalBal1-initBal1)
				})
			})

			t.Run("OffRamp Manual Execution: when executing a non-committed report, it fails", func(t *testing.T) {
				message, root := testutils.CreateNextMessage(ctx, solanaGoClient, t, []solana.PublicKey{})
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: config.EvmChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}

				raw := ccip_router.NewManuallyExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					user.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
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

				root := [32]byte(ccip.MerkleFrom([][]byte{hash1, hash2}))

				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: config.EvmChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            message1.Header.SequenceNumber,
						MaxSeqNr:            message2.Header.SequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(210_000))
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				t.Run("Before elapsed time", func(t *testing.T) {
					t.Run("When user manually executing before the period of time has passed, it fails", func(t *testing.T) {
						executionReport := ccip_router.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message1,
							Root:                root,
							Proofs:              [][32]uint8{[32]byte(hash2)},
						}

						raw := ccip_router.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.RouterConfigPDA,
							config.EvmSourceChainStatePDA,
							rootPDA,
							config.ExternalExecutionConfigPDA,
							user.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.ExternalTokenPoolsSignerPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						fmt.Printf("User: %s\n", user.PublicKey().String())
						fmt.Printf("Transmitter: %s\n", transmitter.PublicKey().String())

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment, []string{ccip_router.ManualExecutionNotAllowed_CcipRouterError.String()})
					})

					t.Run("When transmitter manually executing before the period of time has passed, it fails", func(t *testing.T) {
						executionReport := ccip_router.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message1,
							Root:                root,
							Proofs:              [][32]uint8{[32]byte(hash2)},
						}

						raw := ccip_router.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.RouterConfigPDA,
							config.EvmSourceChainStatePDA,
							rootPDA,
							config.ExternalExecutionConfigPDA,
							transmitter.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.ExternalTokenPoolsSignerPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{ccip_router.ManualExecutionNotAllowed_CcipRouterError.String()})
					})
				})

				t.Run("Given the period of time has passed", func(t *testing.T) {
					instruction, err = ccip_router.NewUpdateEnableManualExecutionAfterInstruction(
						-1,
						config.RouterConfigPDA,
						ccipAdmin.PublicKey(),
						solana.SystemProgramID,
					).ValidateAndBuild()
					require.NoError(t, err)
					result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
					require.NotNil(t, result)

					t.Run("When user manually executing after the period of time has passed, it succeeds", func(t *testing.T) {
						executionReport := ccip_router.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message1,
							Root:                root,
							Proofs:              [][32]uint8{[32]byte(hash2)},
						}

						raw := ccip_router.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.RouterConfigPDA,
							config.EvmSourceChainStatePDA,
							rootPDA,
							config.ExternalExecutionConfigPDA,
							user.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.ExternalTokenPoolsSignerPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
						executionEvent := ccip.EventExecutionStateChanged{}
						require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ExecutionStateChanged", &executionEvent, config.PrintEvents))

						require.NoError(t, err)
						require.NotNil(t, executionEvent)
						require.Equal(t, config.EvmChainSelector, executionEvent.SourceChainSelector)
						require.Equal(t, message1.Header.SequenceNumber, executionEvent.SequenceNumber)
						require.Equal(t, hex.EncodeToString(message1.Header.MessageId[:]), hex.EncodeToString(executionEvent.MessageID[:]))
						require.Equal(t, hex.EncodeToString(hash1[:]), hex.EncodeToString(executionEvent.MessageHash[:]))
						require.Equal(t, ccip_router.Success_MessageExecutionState, executionEvent.State)

						var rootAccount ccip_router.CommitReport
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
						require.NoError(t, err, "failed to get account info")
						require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
						require.Equal(t, bin.Uint128{Lo: 2, Hi: 0}, rootAccount.ExecutionStates)
						require.Equal(t, commitReport.MerkleRoot.MinSeqNr, rootAccount.MinMsgNr)
						require.Equal(t, commitReport.MerkleRoot.MaxSeqNr, rootAccount.MaxMsgNr)
					})

					t.Run("When transmitter executing after the period of time has passed, it succeeds", func(t *testing.T) {
						executionReport := ccip_router.ExecutionReportSingleChain{
							SourceChainSelector: config.EvmChainSelector,
							Message:             message2,
							Root:                root,
							Proofs:              [][32]uint8{[32]byte(hash1)},
						}

						raw := ccip_router.NewManuallyExecuteInstruction(
							testutils.MustMarshalBorsh(t, executionReport),
							[]byte{},
							config.RouterConfigPDA,
							config.EvmSourceChainStatePDA,
							rootPDA,
							config.ExternalExecutionConfigPDA,
							transmitter.PublicKey(),
							solana.SystemProgramID,
							solana.SysVarInstructionsPubkey,
							config.ExternalTokenPoolsSignerPDA,
						)
						raw.AccountMetaSlice = append(
							raw.AccountMetaSlice,
							solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
							solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
							solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
							solana.NewAccountMeta(solana.SystemProgramID, false, false),
						)
						instruction, err = raw.ValidateAndBuild()
						require.NoError(t, err)

						tx = testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment)
						executionEvent := ccip.EventExecutionStateChanged{}
						require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "ExecutionStateChanged", &executionEvent, config.PrintEvents))

						require.NoError(t, err)
						require.NotNil(t, executionEvent)
						require.Equal(t, config.EvmChainSelector, executionEvent.SourceChainSelector)
						require.Equal(t, message2.Header.SequenceNumber, executionEvent.SequenceNumber)
						require.Equal(t, hex.EncodeToString(message2.Header.MessageId[:]), hex.EncodeToString(executionEvent.MessageID[:]))
						require.Equal(t, hex.EncodeToString(hash2[:]), hex.EncodeToString(executionEvent.MessageHash[:]))
						require.Equal(t, ccip_router.Success_MessageExecutionState, executionEvent.State)

						var rootAccount ccip_router.CommitReport
						err = common.GetAccountDataBorshInto(ctx, solanaGoClient, rootPDA, config.DefaultCommitment, &rootAccount)
						require.NoError(t, err, "failed to get account info")
						require.NotEqual(t, bin.Uint128{Lo: 0, Hi: 0}, rootAccount.Timestamp)
						require.Equal(t, bin.Uint128{Lo: 10, Hi: 0}, rootAccount.ExecutionStates)
						require.Equal(t, commitReport.MerkleRoot.MinSeqNr, rootAccount.MinMsgNr)
						require.Equal(t, commitReport.MerkleRoot.MaxSeqNr, rootAccount.MaxMsgNr)
					})
				})
			})

			// solana re-entry is limited by a simple self-recursion and a limited depth
			// https://defisec.info/solana_top_vulnerabilities
			// note: simple recursion execute -> ccipSend is currently not possible as the router does not implement the ccipReceive method signature
			t.Run("failed reentrancy A (execute) -> B (ccipReceive) -> A (ccipSend)", func(t *testing.T) {
				transmitter := getTransmitter()
				receiverContractEvmPDA, err := state.FindNoncePDA(config.EvmChainSelector, config.ReceiverExternalExecutionConfigPDA, config.CcipRouterProgram)
				require.NoError(t, err)

				msgAccounts := []solana.PublicKey{
					config.CcipLogicReceiver,
					config.ReceiverExternalExecutionConfigPDA,
					config.ReceiverTargetAccountPDA,
					solana.SystemProgramID,
					config.CcipRouterProgram,
					config.RouterConfigPDA,
					config.ReceiverExternalExecutionConfigPDA,
					config.EvmSourceChainStatePDA,
					receiverContractEvmPDA,
					solana.SystemProgramID}

				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)

				// To make the message go through the validations we need to specify the correct bitmap in the order
				// of the remaining accounts (writable accounts at positions 1, 5, 6 and 7, i.e. ReceiverTargetAccountPDA,
				// ReceiverExternalExecutionConfigPDA, EVMSourceChainStatePDA and receiverContractEVMPDA)
				message.ExtraArgs.IsWritableBitmap = ccip.GenerateBitMapForIndexes([]int{0, 1, 5, 6, 7})
				hash, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)
				root := [32]byte(hash)

				sourceChainSelector := config.EvmChainSelector
				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            message.Header.SequenceNumber,
						MaxSeqNr:            message.Header.SequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, _ := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)

				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					Root:                root,
					Proofs:              [][32]uint8{}, // single leaf merkle tree
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{},
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)
				raw.AccountMetaSlice = append(
					raw.AccountMetaSlice,
					solana.NewAccountMeta(config.CcipLogicReceiver, false, false),
					// accounts for base CPI call
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.ReceiverTargetAccountPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),

					// accounts for receiver -> router re-entrant CPI call
					solana.NewAccountMeta(config.CcipRouterProgram, false, false),
					solana.NewAccountMeta(config.RouterConfigPDA, false, false),
					solana.NewAccountMeta(config.ReceiverExternalExecutionConfigPDA, true, false),
					solana.NewAccountMeta(config.EvmSourceChainStatePDA, true, false),
					solana.NewAccountMeta(receiverContractEvmPDA, true, false),
					solana.NewAccountMeta(solana.SystemProgramID, false, false),
				)
				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWith(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, []string{"Cross-program invocation reentrancy not allowed for this instruction"})
			})

			t.Run("uninitialized token account can be manually executed", func(t *testing.T) {
				// create new token receiver + find address (does not actually create account, just instruction)
				receiver, err := solana.NewRandomPrivateKey()
				require.NoError(t, err)
				ixATA, ata, err := tokens.CreateAssociatedTokenAccount(token0.Program, token0.Mint.PublicKey(), receiver.PublicKey(), legacyAdmin.PublicKey())
				require.NoError(t, err)
				token0.User[receiver.PublicKey()] = ata

				// create commit report ---------------------
				transmitter := getTransmitter()
				sourceChainSelector := config.EvmChainSelector
				msgAccounts := []solana.PublicKey{}
				message, _ := testutils.CreateNextMessage(ctx, solanaGoClient, t, msgAccounts)
				message.TokenAmounts = []ccip_router.Any2SVMTokenTransfer{{
					SourcePoolAddress: []byte{1, 2, 3},
					DestTokenAddress:  token0.Mint.PublicKey(),
					Amount:            ccip_router.CrossChainAmount{LeBytes: tokens.ToLittleEndianU256(1)},
				}}
				message.TokenReceiver = receiver.PublicKey()
				rootBytes, err := ccip.HashAnyToSVMMessage(message, config.OnRampAddress, msgAccounts)
				require.NoError(t, err)

				root := [32]byte(rootBytes)
				sequenceNumber := message.Header.SequenceNumber
				commitReport := ccip_router.CommitInput{
					MerkleRoot: ccip_router.MerkleRoot{
						SourceChainSelector: sourceChainSelector,
						OnRampAddress:       config.OnRampAddress,
						MinSeqNr:            sequenceNumber,
						MaxSeqNr:            sequenceNumber,
						MerkleRoot:          root,
					},
				}
				sigs, err := ccip.SignCommitReport(reportContext, commitReport, signers)
				require.NoError(t, err)
				rootPDA, err := state.FindCommitReportPDA(config.EvmChainSelector, root, config.CcipRouterProgram)
				require.NoError(t, err)
				instruction, err := ccip_router.NewCommitInstruction(
					reportContext,
					testutils.MustMarshalBorsh(t, commitReport),
					sigs.Rs,
					sigs.Ss,
					sigs.RawVs,
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.BillingSignerPDA,
					config.FeeQuoterProgram,
					config.FqConfigPDA,
				).ValidateAndBuild()
				require.NoError(t, err)
				tx := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, common.AddComputeUnitLimit(300_000))
				event := ccip.EventCommitReportAccepted{}
				require.NoError(t, common.ParseEvent(tx.Meta.LogMessages, "CommitReportAccepted", &event, config.PrintEvents))

				// try to execute report ----------------------
				// should fail because token account does not exist
				executionReport := ccip_router.ExecutionReportSingleChain{
					SourceChainSelector: sourceChainSelector,
					Message:             message,
					OffchainTokenData:   [][]byte{{}},
					Root:                root,
					Proofs:              [][32]uint8{},
				}
				raw := ccip_router.NewExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					reportContext,
					[]byte{0}, // only token transfer message
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					transmitter.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)

				tokenMetas, addressTables, err := tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[receiver.PublicKey()])
				require.NoError(t, err)
				raw.AccountMetaSlice = append(raw.AccountMetaSlice, tokenMetas...)

				instruction, err = raw.ValidateAndBuild()
				require.NoError(t, err)

				testutils.SendAndFailWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, transmitter, config.DefaultCommitment, addressTables, []string{"AccountNotInitialized"})

				// create associated token account for user --------------------
				testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{ixATA}, legacyAdmin, config.DefaultCommitment)
				_, initBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[receiver.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 0, initBal)

				// manual re-execution is successful -----------------------------------
				// NOTE: expects re-execution time to be instantaneous
				rawManual := ccip_router.NewManuallyExecuteInstruction(
					testutils.MustMarshalBorsh(t, executionReport),
					[]byte{0}, // only token transfer message
					config.RouterConfigPDA,
					config.EvmSourceChainStatePDA,
					rootPDA,
					config.ExternalExecutionConfigPDA,
					legacyAdmin.PublicKey(),
					solana.SystemProgramID,
					solana.SysVarInstructionsPubkey,
					config.ExternalTokenPoolsSignerPDA,
				)

				tokenMetas, addressTables, err = tokens.ParseTokenLookupTable(ctx, solanaGoClient, token0, token0.User[receiver.PublicKey()])
				require.NoError(t, err)
				rawManual.AccountMetaSlice = append(rawManual.AccountMetaSlice, tokenMetas...)
				instruction, err = rawManual.ValidateAndBuild()
				require.NoError(t, err)
				testutils.SendAndConfirmWithLookupTables(ctx, t, solanaGoClient, []solana.Instruction{instruction}, legacyAdmin, config.DefaultCommitment, addressTables)

				_, finalBal, err := tokens.TokenBalance(ctx, solanaGoClient, token0.User[receiver.PublicKey()], config.DefaultCommitment)
				require.NoError(t, err)
				require.Equal(t, 1, finalBal-initBal)
			})
		})
	})
}
