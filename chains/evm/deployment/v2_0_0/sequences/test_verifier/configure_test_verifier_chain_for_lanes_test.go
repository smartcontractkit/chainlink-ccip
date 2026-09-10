package test_verifier

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	evmds "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/verifier_test_helper"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/versioned_verifier_resolver"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

type testRemoteVerifierChain struct{}

var _ adapters.RemoteTestVerifierChain = (*testRemoteVerifierChain)(nil)

func (a *testRemoteVerifierChain) TokenAddress(
	d datastore.DataStore,
	_ cldf_chain.BlockChains,
	chainSelector uint64,
) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(
		d,
		datastore.AddressRef{
			Type:      datastore.ContractType("BurnMintERC20WithDrip"),
			Qualifier: "TESTVTR",
		},
		chainSelector,
		evmds.ToEVMAddressBytes,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"resolve TESTVTR token on chain %d: %w",
			chainSelector,
			err,
		)
	}

	return addr, nil
}

func (a *testRemoteVerifierChain) PoolAddress(
	d datastore.DataStore,
	_ cldf_chain.BlockChains,
	chainSelector uint64,
) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(
		d,
		datastore.AddressRef{
			Type:      datastore.ContractType(burn_mint_token_pool.ContractType),
			Version:   burn_mint_token_pool.Version,
			Qualifier: "TESTVTR",
		},
		chainSelector,
		evmds.ToEVMAddressBytes,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"resolve TESTVTR pool on chain %d: %w",
			chainSelector,
			err,
		)
	}

	return addr, nil
}

func (a *testRemoteVerifierChain) VerifierAddress(
	d datastore.DataStore,
	chainSelector uint64,
) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(
		d,
		datastore.AddressRef{
			Type:    datastore.ContractType(verifier_test_helper.ContractType),
			Version: verifier_test_helper.Version,
		},
		chainSelector,
		evmds.ToEVMAddressBytes,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"resolve test verifier on chain %d: %w",
			chainSelector,
			err,
		)
	}

	return addr, nil
}

func (a *testRemoteVerifierChain) VerifierResolverAddress(
	d datastore.DataStore,
	chainSelector uint64,
) ([]byte, error) {
	addr, err := datastore_utils.FindAndFormatRef(
		d,
		datastore.AddressRef{
			Type:      datastore.ContractType(versioned_verifier_resolver.TestVerifierResolverType),
			Version:   versioned_verifier_resolver.Version,
			Qualifier: TestVerifierResolverQualifier,
		},
		chainSelector,
		evmds.ToEVMAddressBytes,
	)
	if err != nil {
		return nil, fmt.Errorf(
			"resolve test verifier resolver on chain %d: %w",
			chainSelector,
			err,
		)
	}

	return addr, nil
}

func TestConfigureTestVerifierChainForLanes(t *testing.T) {
	const (
		chainA = uint64(5009297550715157269)
		chainB = uint64(4949039107694359620)
	)

	e, err := environment.New(
		t.Context(),
		environment.WithEVMSimulated(t, []uint64{chainA, chainB}),
	)
	require.NoError(t, err)

	baseA := setupTestVerifierBaseChain(t, e, chainA)
	baseB := setupTestVerifierBaseChain(t, e, chainB)
	baseDS := mergeTestVerifierDataStores(t, baseA, baseB)

	allowedSender := e.BlockChains.EVMChains()[chainA].DeployerKey.From
	differentSender := common.HexToAddress("0x1234567890123456789012345678901234567890")

	deploy := func(chainSelector uint64) []datastore.AddressRef {
		report, err := operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			DeployTestVerifierChain,
			adapters.DeployTestVerifierChainDeps{
				BlockChains: e.BlockChains,
				DataStore:   baseDS,
			},
			adapters.DeployTestVerifierChainInput{
				ChainSelector: chainSelector,
				TokenName:     "TESTVTR",
				TokenSymbol:   "TESTVTR",
				TokenDecimals: 18,
				AllowedSenders: []string{
					allowedSender.Hex(),
				},
			},
		)
		require.NoError(t, err)

		return report.Output.Addresses
	}

	addressesA := deploy(chainA)
	addressesB := deploy(chainB)

	ds := mergeTestVerifierAddresses(t, baseDS, addressesA)
	ds = mergeTestVerifierAddresses(t, ds, addressesB)

	remoteAdapter := &testRemoteVerifierChain{}

	cfg := adapters.RemoteTestVerifierChainConfig{
		FeeUSDCents:        123,
		GasForVerification: 250_000,
		PayloadSizeBytes:   512,
	}

	input := adapters.ConfigureTestVerifierForLanesInput{
		ChainSelector: chainA,
		AllowedSenders: []string{
			allowedSender.Hex(),
		},
		RemoteChains: map[uint64]adapters.RemoteTestVerifierChainConfig{
			chainB: cfg,
		},
	}

	deps := adapters.ConfigureTestVerifierForLanesDeps{
		BlockChains: e.BlockChains,
		DataStore:   ds,
		RemoteChains: map[uint64]adapters.RemoteTestVerifierChain{
			chainB: remoteAdapter,
		},
	}

	_, err = operations.ExecuteSequence(
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		ConfigureTestVerifierChainForLanes,
		deps,
		input,
	)
	require.NoError(t, err)

	chain := e.BlockChains.EVMChains()[chainA]

	poolRef := requireAddressRef(
		t,
		ds,
		chainA,
		datastore.ContractType(burn_mint_token_pool.ContractType),
		burn_mint_token_pool.Version,
		"TESTVTR",
	)
	poolAddr := common.HexToAddress(poolRef.Address)

	remotePoolRef := requireAddressRef(
		t,
		ds,
		chainB,
		datastore.ContractType(burn_mint_token_pool.ContractType),
		burn_mint_token_pool.Version,
		"TESTVTR",
	)
	remotePoolAddr := common.HexToAddress(remotePoolRef.Address)

	verifierRef := requireAddressRef(
		t,
		ds,
		chainA,
		datastore.ContractType(verifier_test_helper.ContractType),
		verifier_test_helper.Version,
		"",
	)
	verifierAddr := common.HexToAddress(verifierRef.Address)

	resolverRef := requireAddressRef(
		t,
		ds,
		chainA,
		datastore.ContractType(versioned_verifier_resolver.TestVerifierResolverType),
		versioned_verifier_resolver.Version,
		TestVerifierResolverQualifier,
	)

	testRouterRef := requireAddressRef(
		t,
		ds,
		chainA,
		datastore.ContractType(router.TestRouterContractType),
		router.Version,
		"",
	)

	t.Run("adds remote chain to token pool", func(t *testing.T) {
		supportedReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			token_pool.GetSupportedChains,
			chain,
			contract_utils.FunctionInput[struct{}]{
				ChainSelector: chainA,
				Address:       poolAddr,
			},
		)
		require.NoError(t, err)

		require.True(
			t,
			containsSelector(supportedReport.Output, chainB),
			"TESTVTR pool must support remote chain",
		)
	})

	t.Run("configures expected remote pool", func(t *testing.T) {
		remotePoolsReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			token_pool.GetRemotePools,
			chain,
			contract_utils.FunctionInput[uint64]{
				ChainSelector: chainA,
				Address:       poolAddr,
				Args:          chainB,
			},
		)
		require.NoError(t, err)

		expected := common.LeftPadBytes(remotePoolAddr.Bytes(), 32)

		found := false
		for _, remotePool := range remotePoolsReport.Output {
			if string(remotePool) == string(expected) {
				found = true
				break
			}
		}

		require.True(
			t,
			found,
			"expected remote TESTVTR pool to be configured",
		)
	})

	t.Run("configures verifier remote chain and allowlist", func(t *testing.T) {
		configReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			verifier_test_helper.GetRemoteChainConfig,
			chain,
			contract_utils.FunctionInput[uint64]{
				ChainSelector: chainA,
				Address:       verifierAddr,
				Args:          chainB,
			},
		)
		require.NoError(t, err)

		require.Equal(
			t,
			common.HexToAddress(testRouterRef.Address),
			configReport.Output.RemoteChainConfig.Router,
		)
		require.Equal(
			t,
			chainB,
			configReport.Output.RemoteChainConfig.RemoteChainSelector,
		)
		require.True(t, configReport.Output.RemoteChainConfig.AllowlistEnabled)
		require.Equal(
			t,
			cfg.FeeUSDCents,
			configReport.Output.RemoteChainConfig.FeeUSDCents,
		)
		require.Equal(
			t,
			cfg.GasForVerification,
			configReport.Output.RemoteChainConfig.GasForVerification,
		)
		require.Equal(
			t,
			cfg.PayloadSizeBytes,
			configReport.Output.RemoteChainConfig.PayloadSizeBytes,
		)
		require.Contains(
			t,
			configReport.Output.AllowedSendersList,
			allowedSender,
		)
	})

	t.Run("configure allowlist can remove senders", func(t *testing.T) {
		// Remove the allowed sender from the allowlist.
		input := adapters.ConfigureTestVerifierForLanesInput{
			ChainSelector:  chainA,
			AllowedSenders: []string{
				// No allowed senders.
			},
			RemoteChains: map[uint64]adapters.RemoteTestVerifierChainConfig{
				chainB: cfg,
			},
		}

		_, err := operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			ConfigureTestVerifierChainForLanes,
			deps,
			input,
		)
		require.NoError(t, err)

		configReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			verifier_test_helper.GetRemoteChainConfig,
			chain,
			contract_utils.FunctionInput[uint64]{
				ChainSelector: chainA,
				Address:       verifierAddr,
				Args:          chainB,
			},
		)
		require.NoError(t, err)

		require.Empty(
			t,
			configReport.Output.AllowedSendersList,
			"allowlist should be empty after removing all senders",
		)
	})

	t.Run("configure allowlist can add senders", func(t *testing.T) {
		// Add the allowed sender back to the allowlist.
		input := adapters.ConfigureTestVerifierForLanesInput{
			ChainSelector: chainA,
			AllowedSenders: []string{
				differentSender.Hex(),
			},
			RemoteChains: map[uint64]adapters.RemoteTestVerifierChainConfig{
				chainB: cfg,
			},
		}

		_, err := operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			ConfigureTestVerifierChainForLanes,
			deps,
			input,
		)
		require.NoError(t, err)

		configReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			verifier_test_helper.GetRemoteChainConfig,
			chain,
			contract_utils.FunctionInput[uint64]{
				ChainSelector: chainA,
				Address:       verifierAddr,
				Args:          chainB,
			},
		)
		require.NoError(t, err)

		require.Contains(
			t,
			configReport.Output.AllowedSendersList,
			differentSender,
		)
	})

	t.Run("configures resolver outbound implementation", func(t *testing.T) {
		outboundReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			versioned_verifier_resolver.GetAllOutboundImplementations,
			chain,
			contract_utils.FunctionInput[any]{
				ChainSelector: chainA,
				Address:       common.HexToAddress(resolverRef.Address),
			},
		)
		require.NoError(t, err)

		found := false
		for _, impl := range outboundReport.Output {
			if impl.DestChainSelector == chainB &&
				impl.Verifier == verifierAddr {
				found = true
				break
			}
		}

		require.True(
			t,
			found,
			"resolver must route outbound verification for chain B through VerifierTestHelper",
		)
	})

	t.Run("rerun succeeds and keeps a single remote pool", func(t *testing.T) {
		_, err := operations.ExecuteSequence(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			ConfigureTestVerifierChainForLanes,
			deps,
			input,
		)
		require.NoError(
			t,
			err,
			"lane configuration should be safe to rerun",
		)

		supportedReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			token_pool.GetSupportedChains,
			chain,
			contract_utils.FunctionInput[struct{}]{
				ChainSelector: chainA,
				Address:       poolAddr,
			},
		)
		require.NoError(t, err)
		require.True(
			t,
			containsSelector(supportedReport.Output, chainB),
		)

		remotePoolsReport, err := operations.ExecuteOperation(
			testsetup.BundleWithFreshReporter(e.OperationsBundle),
			token_pool.GetRemotePools,
			chain,
			contract_utils.FunctionInput[uint64]{
				ChainSelector: chainA,
				Address:       poolAddr,
				Args:          chainB,
			},
		)
		require.NoError(t, err)

		expected := common.LeftPadBytes(remotePoolAddr.Bytes(), 32)

		count := 0
		for _, remotePool := range remotePoolsReport.Output {
			if string(remotePool) == string(expected) {
				count++
			}
		}

		require.Equal(
			t,
			1,
			count,
			"rerun must not duplicate the configured remote pool",
		)
	})
}
