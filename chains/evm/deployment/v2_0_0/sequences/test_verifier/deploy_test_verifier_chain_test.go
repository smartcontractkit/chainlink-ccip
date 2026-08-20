package test_verifier

import (
	"math/big"
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/verifier_test_helper"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

func TestDeployTestVerifierChain(t *testing.T) {
	chainSelector := chainsel.TEST_90000001.Selector

	e, err := environment.New(
		t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainSelector}),
	)
	require.NoError(t, err)

	baseDS := setupTestVerifierBaseChain(t, e, chainSelector)

	chain := e.BlockChains.EVMChains()[chainSelector]
	preMintAccount := chain.DeployerKey.From
	preMintAmount := big.NewInt(1_000_000_000_000_000_000)

	input := adapters.DeployTestVerifierChainInput{
		ChainSelector: chainSelector,
		TokenName:     "TESTVTR",
		TokenSymbol:   "TESTVTR",
		TokenDecimals: 18,
		PreMintAccounts: map[string]string{
			preMintAccount.Hex(): preMintAmount.String(),
		},
		AllowedSenders: []string{
			preMintAccount.Hex(),
		},
	}

	report, err := operations.ExecuteSequence(
		e.OperationsBundle,
		DeployTestVerifierChain,
		adapters.DeployTestVerifierChainDeps{
			BlockChains: e.BlockChains,
			DataStore:   baseDS,
		},
		input,
	)
	require.NoError(t, err)
	require.NotEmpty(t, report.Output.Addresses)

	ds := mergeTestVerifierAddresses(t, baseDS, report.Output.Addresses)

	tokenRef := requireAddressRef(
		t,
		ds,
		chainSelector,
		datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
		burn_mint_erc20_with_drip.Version,
		"TESTVTR",
	)
	poolRef := requireAddressRef(
		t,
		ds,
		chainSelector,
		datastore.ContractType(burn_mint_token_pool.ContractType),
		burn_mint_token_pool.Version,
		"TESTVTR",
	)
	hooksRef := requireAddressRef(
		t,
		ds,
		chainSelector,
		datastore.ContractType(advanced_pool_hooks.ContractType),
		advanced_pool_hooks.Version,
		"TESTVTR",
	)
	verifierRef := requireAddressRef(
		t,
		ds,
		chainSelector,
		datastore.ContractType(verifier_test_helper.ContractType),
		verifier_test_helper.Version,
		"",
	)
	resolverRef := requireAddressRef(
		t,
		ds,
		chainSelector,
		datastore.ContractType(versioned_verifier_resolver.TestVerifierResolverType),
		versioned_verifier_resolver.Version,
		TestVerifierResolverQualifier,
	)

	tokenAddr := common.HexToAddress(tokenRef.Address)
	poolAddr := common.HexToAddress(poolRef.Address)
	hooksAddr := common.HexToAddress(hooksRef.Address)
	verifierAddr := common.HexToAddress(verifierRef.Address)
	resolverAddr := common.HexToAddress(resolverRef.Address)

	t.Run("premints requested token amount", func(t *testing.T) {
		balanceReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			erc20.BalanceOf,
			chain,
			contract_utils.FunctionInput[common.Address]{
				ChainSelector: chainSelector,
				Address:       tokenAddr,
				Args:          preMintAccount,
			},
		)
		require.NoError(t, err)
		require.Zero(t, preMintAmount.Cmp(balanceReport.Output))
	})

	t.Run("authorizes token pool on advanced hooks", func(t *testing.T) {
		callersReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			advanced_pool_hooks.GetAllAuthorizedCallers,
			chain,
			contract_utils.FunctionInput[struct{}]{
				ChainSelector: chainSelector,
				Address:       hooksAddr,
			},
		)
		require.NoError(t, err)
		require.True(
			t,
			slices.Contains(callersReport.Output, poolAddr),
			"token pool must be an authorized AdvancedPoolHooks caller",
		)
	})

	t.Run("wires resolver inbound implementation to test verifier", func(t *testing.T) {
		inboundReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			versioned_verifier_resolver.GetAllInboundImplementations,
			chain,
			contract_utils.FunctionInput[any]{
				ChainSelector: chainSelector,
				Address:       resolverAddr,
			},
		)
		require.NoError(t, err)

		found := false
		for _, impl := range inboundReport.Output {
			if impl.Version == TestVerifierVersionTag &&
				impl.Verifier == verifierAddr {
				found = true
				break
			}
		}

		require.True(
			t,
			found,
			"resolver must point TestVerifierVersionTag at VerifierTestHelper",
		)
	})

	t.Run("rerun reuses deployments and does not mint again", func(t *testing.T) {
		rerun, err := operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			DeployTestVerifierChain,
			adapters.DeployTestVerifierChainDeps{
				BlockChains: e.BlockChains,
				DataStore:   ds,
			},
			input,
		)
		require.NoError(t, err)

		rerunDS := mergeTestVerifierAddresses(
			t,
			ds,
			rerun.Output.Addresses,
		)

		rerunTokenRef := requireAddressRef(
			t,
			rerunDS,
			chainSelector,
			datastore.ContractType(burn_mint_erc20_with_drip.ContractType),
			burn_mint_erc20_with_drip.Version,
			"TESTVTR",
		)
		rerunPoolRef := requireAddressRef(
			t,
			rerunDS,
			chainSelector,
			datastore.ContractType(burn_mint_token_pool.ContractType),
			burn_mint_token_pool.Version,
			"TESTVTR",
		)
		rerunVerifierRef := requireAddressRef(
			t,
			rerunDS,
			chainSelector,
			datastore.ContractType(verifier_test_helper.ContractType),
			verifier_test_helper.Version,
			"",
		)
		rerunResolverRef := requireAddressRef(
			t,
			rerunDS,
			chainSelector,
			datastore.ContractType(versioned_verifier_resolver.TestVerifierResolverType),
			versioned_verifier_resolver.Version,
			TestVerifierResolverQualifier,
		)

		require.Equal(t, tokenRef.Address, rerunTokenRef.Address)
		require.Equal(t, poolRef.Address, rerunPoolRef.Address)
		require.Equal(t, verifierRef.Address, rerunVerifierRef.Address)
		require.Equal(t, resolverRef.Address, rerunResolverRef.Address)

		balanceReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			erc20.BalanceOf,
			chain,
			contract_utils.FunctionInput[common.Address]{
				ChainSelector: chainSelector,
				Address:       tokenAddr,
				Args:          preMintAccount,
			},
		)
		require.NoError(t, err)
		require.Zero(
			t,
			preMintAmount.Cmp(balanceReport.Output),
			"rerunning deployment must not increase pre-minted balance",
		)
	})
}
