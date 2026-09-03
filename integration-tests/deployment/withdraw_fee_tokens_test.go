package deployment

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	bnmERC20Bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
	"github.com/stretchr/testify/require"

	evmadaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evmadaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	evmonrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	soladaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	solcommon "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	solstate "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	soltokens "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

// withdrawFixture is the per-chain state a withdrawal test needs: where the fees sit and where
// they should end up.
type withdrawFixture struct {
	selector      uint64
	onRamp        common.Address
	feeAggregator common.Address
	tokens        []*bnmERC20Bindings.BurnMintERC20
}

func onRampAddressForTest(t *testing.T, env *cldf_deployment.Environment, selector uint64) common.Address {
	t.Helper()
	ref, err := datastore_utils.FindAndFormatRef(
		env.DataStore,
		datastore.AddressRef{
			ChainSelector: selector,
			Type:          datastore.ContractType(evmonrampops.ContractType),
			Version:       evmonrampops.Version,
		},
		selector,
		datastore_utils.FullRef,
	)
	require.NoError(t, err)
	require.True(t, common.IsHexAddress(ref.Address))
	return common.HexToAddress(ref.Address)
}

// fundOnRampWithFeeToken deploys a fresh burn/mint token and mints `amount` of it directly to the
// OnRamp, standing in for fees that accumulated there from ccipSend calls.
func fundOnRampWithFeeToken(t *testing.T, env *cldf_deployment.Environment, selector uint64, onRamp common.Address, amount *big.Int) *bnmERC20Bindings.BurnMintERC20 {
	t.Helper()

	chain, ok := env.BlockChains.EVMChains()[selector]
	require.True(t, ok)

	token := DeployBurnMintTokenEVM(t, env, selector, "")

	opts := *chain.DeployerKey
	opts.Context = t.Context()

	grantTx, err := token.GrantMintAndBurnRoles(&opts, chain.DeployerKey.From)
	require.NoError(t, err)
	_, err = chain.Confirm(grantTx)
	require.NoError(t, err)

	mintTx, err := token.Mint(&opts, onRamp, amount)
	require.NoError(t, err)
	_, err = chain.Confirm(mintTx)
	require.NoError(t, err)

	balance, err := token.BalanceOf(&bind.CallOpts{Context: t.Context()}, onRamp)
	require.NoError(t, err)
	require.Equal(t, 0, balance.Cmp(amount), "OnRamp should hold the minted fee tokens before withdrawal")

	return token
}

// TestWithdrawFeeTokensEVM_V1_6_0 exercises the changeset across two chains at once, with two fee
// tokens on each, which is the shape operators actually run: one proposal, many lanes.
func TestWithdrawFeeTokensEVM_V1_6_0(t *testing.T) {
	selA := chainsel.TEST_90000001.Selector
	selB := chainsel.TEST_90000002.Selector

	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{selA, selB}),
	)
	require.NoError(t, err)

	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	deployRegistry := deployapi.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deployapi.MCMSVersion, &evmadaptersV1_0_0.EVMDeployer{})

	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadaptersV1_0_0.EVMMCMSReader{})

	feeAggAdapter := evmadaptersV1_6_0.NewFeeAggregatorAdapter(&evmAdapter)

	SeedUltraFastCurseMCMS(t, env)
	output, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			selA: NewDefaultDeploymentConfigForEVM(cciputils.Version_1_6_0),
			selB: NewDefaultDeploymentConfigForEVM(cciputils.Version_1_6_0),
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	DeployMCMS(t, env, selA, []string{cciputils.CLLQualifier})
	DeployMCMS(t, env, selB, []string{cciputils.CLLQualifier})

	// Point each chain at a fresh fee aggregator so its starting balance is unambiguously zero.
	feeAggregators := map[uint64]common.Address{
		selA: common.HexToAddress("0x00000000000000000000000000000000000000A1"),
		selB: common.HexToAddress("0x00000000000000000000000000000000000000B2"),
	}
	_, err = fees.SetFeeAggregator().Apply(*env, fees.SetFeeAggregatorInput{
		Version: cciputils.Version_1_6_0,
		MCMS:    mcms.Input{},
		Args: []fees.FeeAggregatorForChain{
			{ChainSelector: selA, FeeAggregator: feeAggregators[selA].Hex()},
			{ChainSelector: selB, FeeAggregator: feeAggregators[selB].Hex()},
		},
	})
	require.NoError(t, err)

	amount := big.NewInt(1e18)

	fixtures := make([]withdrawFixture, 0, 2)
	for _, sel := range []uint64{selA, selB} {
		onRamp := onRampAddressForTest(t, env, sel)

		onchainFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, sel)
		require.NoError(t, err)
		require.Equal(t, feeAggregators[sel], common.HexToAddress(onchainFeeAgg))

		fixtures = append(fixtures, withdrawFixture{
			selector:      sel,
			onRamp:        onRamp,
			feeAggregator: feeAggregators[sel],
			tokens: []*bnmERC20Bindings.BurnMintERC20{
				fundOnRampWithFeeToken(t, env, sel, onRamp, amount),
				fundOnRampWithFeeToken(t, env, sel, onRamp, amount),
			},
		})
	}

	args := make([]fees.WithdrawFeeTokensForChain, 0, len(fixtures))
	for _, f := range fixtures {
		feeTokens := make([]fees.FeeTokenWithdrawal, 0, len(f.tokens))
		for _, tok := range f.tokens {
			feeTokens = append(feeTokens, fees.FeeTokenWithdrawal{Token: tok.Address().Hex()})
		}
		args = append(args, fees.WithdrawFeeTokensForChain{
			ChainSelector: f.selector,
			FeeTokens:     feeTokens,
		})
	}

	_, err = fees.WithdrawFeeTokens().Apply(*env, fees.WithdrawFeeTokensInput{
		Version: cciputils.Version_1_6_0,
		MCMS:    mcms.Input{},
		Args:    args,
	})
	require.NoError(t, err)

	for _, f := range fixtures {
		for i, tok := range f.tokens {
			onRampBalance, err := tok.BalanceOf(&bind.CallOpts{Context: t.Context()}, f.onRamp)
			require.NoError(t, err)
			require.Equal(t, 0, onRampBalance.Cmp(big.NewInt(0)),
				"chain %d token %d: OnRamp balance should be swept to zero", f.selector, i)

			aggBalance, err := tok.BalanceOf(&bind.CallOpts{Context: t.Context()}, f.feeAggregator)
			require.NoError(t, err)
			require.Equal(t, 0, aggBalance.Cmp(amount),
				"chain %d token %d: fee aggregator should have received the full balance", f.selector, i)
		}
	}

	// The deployed 1.6.0 OnRamp rejects a zero fee aggregator in setDynamicConfig, so the adapter's
	// zero-address guard cannot fire on this version -- the invariant is enforced on-chain. (The 2.0
	// Proxy and OnRamp do NOT validate it, which is where that guard earns its keep.) If a future
	// 1.6.x drops this validation, this test fails and the 1.6 guard becomes load-bearing.
	t.Run("OnRamp itself refuses a zero fee aggregator", func(t *testing.T) {
		_, err := fees.SetFeeAggregator().Apply(*env, fees.SetFeeAggregatorInput{
			Version: cciputils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.FeeAggregatorForChain{
				{ChainSelector: selA, FeeAggregator: (common.Address{}).Hex()},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "InvalidConfig")
	})
}

func TestWithdrawFeeTokens_ValidationErrors(t *testing.T) {
	evmSel := chainsel.TEST_90000001.Selector

	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{evmSel}),
	)
	require.NoError(t, err)

	linkAddr := "0x779877A7B0D9E8603169DdbD7836e478b4624789"
	cs := fees.WithdrawFeeTokens()

	t.Run("nil version", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: nil,
			Args: []fees.WithdrawFeeTokensForChain{
				{ChainSelector: evmSel, FeeTokens: []fees.FeeTokenWithdrawal{{Token: linkAddr}}},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "version is required")
	})

	t.Run("empty args", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			Args:    []fees.WithdrawFeeTokensForChain{},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "at least one chain must be specified")
	})

	t.Run("no fee tokens for a chain", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			Args: []fees.WithdrawFeeTokensForChain{
				{ChainSelector: evmSel, FeeTokens: nil},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "at least one fee token is required")
	})

	t.Run("empty fee token address", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			Args: []fees.WithdrawFeeTokensForChain{
				{ChainSelector: evmSel, FeeTokens: []fees.FeeTokenWithdrawal{{Token: ""}}},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "fee token address is required")
	})

	t.Run("duplicate chain selector", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			Args: []fees.WithdrawFeeTokensForChain{
				{ChainSelector: evmSel, FeeTokens: []fees.FeeTokenWithdrawal{{Token: linkAddr}}},
				{ChainSelector: evmSel, FeeTokens: []fees.FeeTokenWithdrawal{{Token: linkAddr}}},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "duplicate chain selector")
	})

	t.Run("duplicate fee token within a chain", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			Args: []fees.WithdrawFeeTokensForChain{
				{ChainSelector: evmSel, FeeTokens: []fees.FeeTokenWithdrawal{
					{Token: linkAddr},
					{Token: linkAddr},
				}},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "duplicate fee token")
	})

	t.Run("non-positive amount", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			Args: []fees.WithdrawFeeTokensForChain{
				{ChainSelector: evmSel, FeeTokens: []fees.FeeTokenWithdrawal{
					{Token: linkAddr, Amount: big.NewInt(0)},
				}},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "amount must be positive")
	})

	t.Run("multiple chains with multiple tokens is valid", func(t *testing.T) {
		err := cs.VerifyPreconditions(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			Args: []fees.WithdrawFeeTokensForChain{
				{ChainSelector: evmSel, FeeTokens: []fees.FeeTokenWithdrawal{
					{Token: linkAddr},
					{Token: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"},
				}},
			},
		})
		require.NoError(t, err)
	})
}

// solTokenBalance reads an SPL token account balance in base units.
func solTokenBalance(t *testing.T, chain cldf_solana.Chain, account solana.PublicKey) uint64 {
	t.Helper()
	_, balance, err := soltokens.TokenBalance(t.Context(), chain.Client, account, cldf_solana.SolDefaultCommitment)
	require.NoError(t, err)
	require.GreaterOrEqual(t, balance, 0)
	return uint64(balance)
}

// solanaFeeToken is a mint plus the two token accounts a withdrawal moves value between.
type solanaFeeToken struct {
	mint     solana.PublicKey
	accumATA solana.PublicKey // owned by the router's fee billing signer PDA; where fees pile up
	aggATA   solana.PublicKey // owned by the fee aggregator; where they should land
}

// setupSolanaFeeToken mints a fresh SPL token straight into the router's billing accumulator,
// standing in for fees that accrued there from ccipSend calls, and creates the fee aggregator's
// ATA so the router has somewhere to send them.
func setupSolanaFeeToken(
	t *testing.T,
	chain cldf_solana.Chain,
	billingSigner solana.PublicKey,
	feeAggregator solana.PublicKey,
	amount uint64,
	createAggATA bool,
) solanaFeeToken {
	t.Helper()

	mintPriv, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)
	mint := mintPriv.PublicKey()

	createMintIxs, err := soltokens.CreateToken(
		t.Context(),
		solana.TokenProgramID,
		mint,
		chain.DeployerKey.PublicKey(),
		9,
		chain.Client,
		cldf_solana.SolDefaultCommitment,
	)
	require.NoError(t, err)
	require.NoError(t, chain.Confirm(createMintIxs, solcommon.AddSigners(mintPriv)))

	accumIx, accumATA, err := soltokens.CreateAssociatedTokenAccount(
		solana.TokenProgramID, mint, billingSigner, chain.DeployerKey.PublicKey())
	require.NoError(t, err)
	ixs := []solana.Instruction{accumIx}

	aggATA, _, err := soltokens.FindAssociatedTokenAddress(solana.TokenProgramID, mint, feeAggregator)
	require.NoError(t, err)
	if createAggATA {
		aggIx, _, err := soltokens.CreateAssociatedTokenAccount(
			solana.TokenProgramID, mint, feeAggregator, chain.DeployerKey.PublicKey())
		require.NoError(t, err)
		ixs = append(ixs, aggIx)
	}
	require.NoError(t, chain.Confirm(ixs))

	mintToIx, err := soltokens.MintTo(amount, solana.TokenProgramID, mint, accumATA, chain.DeployerKey.PublicKey())
	require.NoError(t, err)
	require.NoError(t, chain.Confirm([]solana.Instruction{mintToIx}))

	require.Equal(t, amount, solTokenBalance(t, chain, accumATA), "billing accumulator should be funded")
	if createAggATA {
		require.Equal(t, uint64(0), solTokenBalance(t, chain, aggATA), "fee aggregator should start empty")
	}

	return solanaFeeToken{mint: mint, accumATA: accumATA, aggATA: aggATA}
}

// TestWithdrawFeeTokensSolana_V1_6_0 covers the Solana path end to end, including the per-token
// amount that only Solana can honour: a partial withdrawal first, then a full sweep of what's left.
func TestWithdrawFeeTokensSolana_V1_6_0(t *testing.T) {
	solSel := chainsel.SOLANA_DEVNET.Selector

	programsPath, ds, err := PreloadSolanaEnvironment(t, solSel)
	require.NoError(t, err)

	env, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{solSel}, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	solAdapter := solseqV1_6_0.SolanaAdapter{}

	deployRegistry := deployapi.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deployapi.MCMSVersion, &solAdapter)

	SeedUltraFastCurseMCMS(t, env)
	output, err := deployapi.DeployContracts(deployRegistry).Apply(*env, deployapi.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deployapi.ContractDeploymentConfigPerChain{
			solSel: NewDefaultDeploymentConfigForSolana(cciputils.Version_1_6_0),
		},
	})
	require.NoError(t, err)
	require.NoError(t, output.DataStore.Merge(env.DataStore))
	env.DataStore = output.DataStore.Seal()

	solChain, ok := env.BlockChains.SolanaChains()[solSel]
	require.True(t, ok)

	// Point the router at a fresh fee aggregator so its ATAs start empty.
	feeAggregator := solana.NewWallet()
	_, err = fees.SetFeeAggregator().Apply(*env, fees.SetFeeAggregatorInput{
		Version: cciputils.Version_1_6_0,
		MCMS:    mcms.Input{},
		Args: []fees.FeeAggregatorForChain{
			{ChainSelector: solSel, FeeAggregator: feeAggregator.PublicKey().String()},
		},
	})
	require.NoError(t, err)

	feeAggAdapter := soladaptersV1_6_0.NewFeeAggregatorAdapter(&solAdapter)
	onchainFeeAgg, err := feeAggAdapter.GetFeeAggregator(*env, solSel)
	require.NoError(t, err)
	require.Equal(t, feeAggregator.PublicKey().String(), onchainFeeAgg)

	routerAddr, err := solAdapter.GetRouterAddress(env.DataStore, solSel)
	require.NoError(t, err)
	routerPubkey := solana.PublicKeyFromBytes(routerAddr)

	billingSigner, _, err := solstate.FindFeeBillingSignerPDA(routerPubkey)
	require.NoError(t, err)

	const (
		accumulated = uint64(1_000_000)
		partial     = uint64(250_000)
	)

	// Two fee tokens, so the sequence has to iterate rather than handle a single mint.
	tokenA := setupSolanaFeeToken(t, solChain, billingSigner, feeAggregator.PublicKey(), accumulated, true)
	tokenB := setupSolanaFeeToken(t, solChain, billingSigner, feeAggregator.PublicKey(), accumulated, true)

	t.Run("partial withdrawal honours the per-token amount", func(t *testing.T) {
		_, err := fees.WithdrawFeeTokens().Apply(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.WithdrawFeeTokensForChain{
				{
					ChainSelector: solSel,
					FeeTokens: []fees.FeeTokenWithdrawal{
						{Token: tokenA.mint.String(), Amount: new(big.Int).SetUint64(partial)},
						{Token: tokenB.mint.String(), Amount: new(big.Int).SetUint64(partial)},
					},
				},
			},
		})
		require.NoError(t, err)

		for name, tok := range map[string]solanaFeeToken{"tokenA": tokenA, "tokenB": tokenB} {
			require.Equal(t, accumulated-partial, solTokenBalance(t, solChain, tok.accumATA),
				"%s: only the requested amount should leave the accumulator", name)
			require.Equal(t, partial, solTokenBalance(t, solChain, tok.aggATA),
				"%s: fee aggregator should have received exactly the requested amount", name)
		}
	})

	t.Run("omitting the amount sweeps the remaining balance", func(t *testing.T) {
		_, err := fees.WithdrawFeeTokens().Apply(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.WithdrawFeeTokensForChain{
				{
					ChainSelector: solSel,
					FeeTokens: []fees.FeeTokenWithdrawal{
						{Token: tokenA.mint.String()},
						{Token: tokenB.mint.String()},
					},
				},
			},
		})
		require.NoError(t, err)

		for name, tok := range map[string]solanaFeeToken{"tokenA": tokenA, "tokenB": tokenB} {
			require.Equal(t, uint64(0), solTokenBalance(t, solChain, tok.accumATA),
				"%s: accumulator should be fully swept", name)
			require.Equal(t, accumulated, solTokenBalance(t, solChain, tok.aggATA),
				"%s: fee aggregator should hold everything that accrued", name)
		}
	})

	// The router will not create the recipient account, and ATA creation is permissionless, so the
	// sequence creates it with the deployer key rather than making the operator do it beforehand.
	t.Run("creates the fee aggregator ATA when it is missing", func(t *testing.T) {
		tok := setupSolanaFeeToken(t, solChain, billingSigner, feeAggregator.PublicKey(), accumulated, false)

		_, err := solChain.Client.GetAccountInfoWithOpts(t.Context(), tok.aggATA, &rpc.GetAccountInfoOpts{
			Commitment: cldf_solana.SolDefaultCommitment,
		})
		require.ErrorIs(t, err, rpc.ErrNotFound, "fee aggregator ATA should not exist yet")

		_, err = fees.WithdrawFeeTokens().Apply(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.WithdrawFeeTokensForChain{
				{
					ChainSelector: solSel,
					FeeTokens:     []fees.FeeTokenWithdrawal{{Token: tok.mint.String()}},
				},
			},
		})
		require.NoError(t, err)

		require.Equal(t, uint64(0), solTokenBalance(t, solChain, tok.accumATA),
			"accumulator should be swept")
		require.Equal(t, accumulated, solTokenBalance(t, solChain, tok.aggATA),
			"fee aggregator ATA should have been created and funded")
	})

	t.Run("unknown mint fails before any transaction is sent", func(t *testing.T) {
		_, err := fees.WithdrawFeeTokens().Apply(*env, fees.WithdrawFeeTokensInput{
			Version: cciputils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.WithdrawFeeTokensForChain{
				{
					ChainSelector: solSel,
					FeeTokens: []fees.FeeTokenWithdrawal{
						{Token: solana.NewWallet().PublicKey().String()},
					},
				},
			},
		})
		require.Error(t, err)
		require.Contains(t, err.Error(), "not found")
	})
}
