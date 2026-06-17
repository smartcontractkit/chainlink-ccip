package hooks

import (
	"context"
	"encoding/csv"
	"errors"
	"math/big"
	"os"
	"testing"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evmchain "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

type stubLaneSanityProvider struct {
	stubPostProposalProvider
	applySenderPrivateKeyFn       func(ctx context.Context, lggr logger.Logger, env *cldf.Environment, senderKey string) error
	availableTransferTokensFn     func(env cldf.Environment, source, dest uint64) (map[string]string, error)
	encodeReceiverAddressFn       func(env cldf.Environment, destSel uint64, receiverAddress string) ([]byte, error)
	mockReceiverAddressFn         func(env cldf.Environment, chainSel uint64) ([]byte, error)
	fundAndApproveTransferTokenFn func(ctx context.Context, env cldf.Environment, srcSel uint64, tokenAddress string) (*big.Int, error)
	getMessageFeeFn               func(ctx context.Context, env cldf.Environment, srcSel, destSel uint64, msg any) (string, error)
}

func (s *stubLaneSanityProvider) ApplySenderPrivateKey(
	ctx context.Context,
	lggr logger.Logger,
	env *cldf.Environment,
	senderKey string,
) error {
	if s.applySenderPrivateKeyFn != nil {
		return s.applySenderPrivateKeyFn(ctx, lggr, env, senderKey)
	}
	return nil
}

func (s *stubLaneSanityProvider) FeeTokenName(env cldf.Environment, source uint64, tokenAddr string) (string, error) {
	return "MOCK" + tokenAddr, nil
}

func (s *stubLaneSanityProvider) AvailableTransferTokens(
	env cldf.Environment,
	source, dest uint64,
) (map[string]string, error) {
	if s.availableTransferTokensFn != nil {
		return s.availableTransferTokensFn(env, source, dest)
	}
	return nil, nil
}

func (s *stubLaneSanityProvider) EncodeReceiverAddress(
	env cldf.Environment,
	destSel uint64,
	receiverAddress string,
) ([]byte, error) {
	if s.encodeReceiverAddressFn != nil {
		return s.encodeReceiverAddressFn(env, destSel, receiverAddress)
	}
	return []byte(receiverAddress), nil
}

func (s *stubLaneSanityProvider) MockReceiverAddress(
	env cldf.Environment,
	chainSel uint64,
) ([]byte, error) {
	if s.mockReceiverAddressFn != nil {
		return s.mockReceiverAddressFn(env, chainSel)
	}
	return nil, nil
}

func (s *stubLaneSanityProvider) FundAndApproveTransferToken(
	ctx context.Context,
	env cldf.Environment,
	srcSel uint64,
	tokenAddress string,
) (*big.Int, error) {
	if s.fundAndApproveTransferTokenFn != nil {
		return s.fundAndApproveTransferTokenFn(ctx, env, srcSel, tokenAddress)
	}
	return big.NewInt(1), nil
}

func (s *stubLaneSanityProvider) GetMessageFee(
	ctx context.Context,
	env cldf.Environment,
	srcSel, destSel uint64,
	msg any,
) (string, error) {
	if s.getMessageFeeFn != nil {
		return s.getMessageFeeFn(ctx, env, srcSel, destSel, msg)
	}
	return "", nil
}

func TestPostProposalLaneSanityRegistry_RegisterFirstWinsAndReset(t *testing.T) {
	ResetPostProposalLaneSanityRegistryForTest()

	reg := GetPostProposalLaneSanityRegistry()
	first := &stubLaneSanityProvider{}
	second := &stubLaneSanityProvider{}
	reg.Register(chain_selectors.FamilyEVM, first)
	reg.Register(chain_selectors.FamilyEVM, second)

	got, ok := reg.Get(chain_selectors.FamilyEVM)
	require.True(t, ok)
	require.Same(t, first, got)

	ResetPostProposalLaneSanityRegistryForTest()
	_, ok = GetPostProposalLaneSanityRegistry().Get(chain_selectors.FamilyEVM)
	require.False(t, ok)
}

func TestExpandBidirectionalRequests(t *testing.T) {
	t.Run("nil for empty input", func(t *testing.T) {
		require.Nil(t, expandBidirectionalRequests(nil))
	})

	t.Run("expands each pair both ways", func(t *testing.T) {
		got := expandBidirectionalRequests([]LaneSanityChainPair{
			{ChainA: 10, ChainB: 20},
			{ChainA: 30, ChainB: 40},
		})
		require.Len(t, got, 4)
		require.Equal(t, laneSanityCheckRequest{src: 10, dest: 20}, got[0])
		require.Equal(t, laneSanityCheckRequest{src: 20, dest: 10}, got[1])
		require.Equal(t, laneSanityCheckRequest{src: 30, dest: 40}, got[2])
		require.Equal(t, laneSanityCheckRequest{src: 40, dest: 30}, got[3])
	})
}

func TestApplySenderPrivateKeyForRequests(t *testing.T) {
	ResetPostProposalLaneSanityRegistryForTest()
	t.Cleanup(ResetPostProposalLaneSanityRegistryForTest)

	eth := chain_selectors.ETHEREUM_MAINNET.Selector
	poly := chain_selectors.POLYGON_MAINNET.Selector

	var lastKey string
	var callCount int
	GetPostProposalLaneSanityRegistry().Register(chain_selectors.FamilyEVM, &stubLaneSanityProvider{
		applySenderPrivateKeyFn: func(_ context.Context, _ logger.Logger, env *cldf.Environment, senderKey string) error {
			lastKey = senderKey
			callCount++
			return nil
		},
	})

	requests := []laneSanityCheckRequest{
		{src: eth, dest: poly},
		{src: poly, dest: eth},
	}

	env, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{eth, poly}))
	require.NoError(t, err)

	t.Run("empty env var leaves env unchanged", func(t *testing.T) {
		lastKey = ""
		callCount = 0
		t.Setenv(laneSanityPrivateKeyEnv, "")

		env := cldf.Environment{}
		err := applySenderPrivateKeyForRequests(t.Context(), logger.Test(t), &env, requests)
		require.NoError(t, err)
		require.Empty(t, lastKey)
		require.Zero(t, callCount)
	})

	t.Run("delegates once per family", func(t *testing.T) {
		lastKey = ""
		callCount = 0
		t.Setenv(laneSanityPrivateKeyEnv, "0xdeadbeef")
		err = applySenderPrivateKeyForRequests(t.Context(), logger.Test(t), env, requests)
		require.NoError(t, err)
		require.Equal(t, "0xdeadbeef", lastKey)
		require.Equal(t, 1, callCount)
		require.Equal(t, 2, len(env.BlockChains.EVMChains()))
	})

	t.Run("missing provider errors", func(t *testing.T) {
		ResetPostProposalLaneSanityRegistryForTest()
		t.Setenv(laneSanityPrivateKeyEnv, "0xabc")

		err := applySenderPrivateKeyForRequests(t.Context(), logger.Test(t), &cldf.Environment{}, requests)
		require.Error(t, err)
		require.ErrorContains(t, err, "no provider for family")
	})

	t.Run("provider error propagates", func(t *testing.T) {
		ResetPostProposalLaneSanityRegistryForTest()
		GetPostProposalLaneSanityRegistry().Register(chain_selectors.FamilyEVM, &stubLaneSanityProvider{
			applySenderPrivateKeyFn: func(_ context.Context, _ logger.Logger, env *cldf.Environment, _ string) error {
				return errors.New("bad key")
			},
		})
		t.Setenv(laneSanityPrivateKeyEnv, "0xabc")

		err := applySenderPrivateKeyForRequests(t.Context(), logger.Test(t), env, requests)
		require.Error(t, err)
		require.ErrorContains(t, err, "bad key")
	})
}

func TestResolveTransferTokenChoice(t *testing.T) {
	available := map[string]string{
		"LINK":  "0xlink",
		"USDC":  "0xusdc",
		"Token": "0xtoken",
	}
	names := []string{"LINK", "Token", "USDC"}

	t.Run("numeric choice", func(t *testing.T) {
		name, addr, ok := resolveTransferTokenChoice("2", names, available)
		require.True(t, ok)
		require.Equal(t, "Token", name)
		require.Equal(t, "0xtoken", addr)
	})

	t.Run("out of range numeric choice", func(t *testing.T) {
		_, _, ok := resolveTransferTokenChoice("99", names, available)
		require.False(t, ok)
	})

	t.Run("case insensitive name", func(t *testing.T) {
		name, addr, ok := resolveTransferTokenChoice("link", names, available)
		require.True(t, ok)
		require.Equal(t, "LINK", name)
		require.Equal(t, "0xlink", addr)
	})

	t.Run("unknown choice", func(t *testing.T) {
		_, _, ok := resolveTransferTokenChoice("WETH", names, available)
		require.False(t, ok)
	})
}

func TestCCIPExplorerURL(t *testing.T) {
	require.Equal(t, "https://ccip.chain.link/msg/0xabc", ccipExplorerURL("0xabc"))
	require.Equal(t, "https://ccip.chain.link/msg/0xABC", ccipExplorerURL("ABC"))
}

func TestBuildLaneSanityFailureSummary(t *testing.T) {
	failed := map[laneFailureKey]map[string]struct{}{
		{srcSel: 2, destSel: 3}: {"native": {}, "PTT:0xtok": {}},
		{srcSel: 1, destSel: 2}: {"LINK": {}},
	}
	got := buildLaneSanityFailureSummary(failed)
	require.Contains(t, got, "src=1 dest=2 scenarios=[LINK]")
	require.Contains(t, got, "src=2 dest=3 scenarios=[PTT:0xtok,native]")
}

func TestResolveLaneSanityReceiver(t *testing.T) {
	eth := chain_selectors.ETHEREUM_MAINNET.Selector
	destAdapter := &stubTestAdapter{receiver: []byte("default-receiver")}
	provider := &stubLaneSanityProvider{
		encodeReceiverAddressFn: func(_ cldf.Environment, _ uint64, receiverAddress string) ([]byte, error) {
			return []byte("encoded:" + receiverAddress), nil
		},
	}

	t.Run("empty address uses adapter default", func(t *testing.T) {
		got, err := resolveLaneSanityReceiver(logger.Test(t), provider, cldf.Environment{}, destAdapter, eth, "")
		require.NoError(t, err)
		require.Equal(t, []byte("default-receiver"), got)
	})

	t.Run("explicit address is encoded", func(t *testing.T) {
		got, err := resolveLaneSanityReceiver(logger.Test(t), provider, cldf.Environment{}, destAdapter, eth, "0xabc")
		require.NoError(t, err)
		require.Equal(t, []byte("encoded:0xabc"), got)
	})

	t.Run("encode error propagates", func(t *testing.T) {
		badProvider := &stubLaneSanityProvider{
			encodeReceiverAddressFn: func(_ cldf.Environment, _ uint64, _ string) ([]byte, error) {
				return nil, errors.New("invalid receiver")
			},
		}
		_, err := resolveLaneSanityReceiver(logger.Test(t), badProvider, cldf.Environment{}, destAdapter, eth, "0xbad")
		require.Error(t, err)
		require.ErrorContains(t, err, "invalid receiver")
	})
}

func TestFeeTokenLabels(t *testing.T) {
	require.Equal(t, []string{"native", "0xtoken"}, feeTokenLabels([]string{"", "0xtoken"}))
}

func TestRunLaneSanityChecks_EmptyRequests(t *testing.T) {
	require.NoError(t, RunLaneSanityChecks(t.Context(), logger.Test(t), cldf.Environment{}, nil, ""))
}

func TestRunLaneSanityChecks_SkipsUnregisteredFamily(t *testing.T) {
	ResetPostProposalLaneSanityRegistryForTest()
	t.Cleanup(ResetPostProposalLaneSanityRegistryForTest)

	err := RunLaneSanityChecks(t.Context(), logger.Test(t), cldf.Environment{}, []LaneSanityChainPair{
		{ChainA: chain_selectors.ETHEREUM_MAINNET.Selector, ChainB: chain_selectors.POLYGON_MAINNET.Selector},
	}, "")
	require.NoError(t, err)
}

func TestRunLaneSanityTokenChecksForPair_SkipsPTTWithoutMockReceiver(t *testing.T) {
	ResetPostProposalLaneSanityRegistryForTest()
	t.Cleanup(func() {
		transferTokenSelector = promptTransferTokenSelection
	})

	eth := chain_selectors.ETHEREUM_MAINNET.Selector
	poly := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()
	adapters := map[uint64]*stubTestAdapter{
		eth:  {selector: eth, family: chain_selectors.FamilyEVM},
		poly: {selector: poly, family: chain_selectors.FamilyEVM},
	}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, adapters)

	var sendCount int
	provider := &stubLaneSanityProvider{
		stubPostProposalProvider: stubPostProposalProvider{
			adapterVersionForLaneFn: func(_ cldf.Environment, _, _ uint64) (*semver.Version, error) {
				return version, nil
			},
		},
		availableTransferTokensFn: func(_ cldf.Environment, _, _ uint64) (map[string]string, error) {
			return map[string]string{"LINK": "0xlink"}, nil
		},
		mockReceiverAddressFn: func(_ cldf.Environment, _ uint64) ([]byte, error) {
			return nil, nil
		},
		fundAndApproveTransferTokenFn: func(_ context.Context, _ cldf.Environment, _ uint64, _ string) (*big.Int, error) {
			return big.NewInt(1), nil
		},
	}
	adapters[eth].sendMessageFn = func(_ context.Context, _ uint64, _ any) (uint64, string, error) {
		sendCount++
		return 1, "transfer-id", nil
	}

	transferTokenSelector = func(_ map[string]string, _, _ uint64) ([]string, error) {
		return []string{"0xlink"}, nil
	}

	env := cldf.Environment{
		BlockChains: chain.NewBlockChains(map[uint64]chain.BlockChain{
			eth:  evmchain.Chain{Selector: eth},
			poly: evmchain.Chain{Selector: poly},
		}),
	}

	err := runLaneSanityTokenChecksForPair(
		t.Context(), logger.Test(t), env, chain_selectors.FamilyEVM, provider, eth, poly,
		[]string{"0xlink"}, "", &laneSanityResultCollector{},
	)
	require.NoError(t, err)
	require.Equal(t, 1, sendCount)
}

func TestFormatDataForCSV(t *testing.T) {
	require.Equal(t, "", formatDataForCSV(nil))
	require.Equal(t, "lane-sanity-check", formatDataForCSV([]byte("lane-sanity-check")))
	require.Equal(t, "0xdeadbeef", formatDataForCSV([]byte{0xde, 0xad, 0xbe, 0xef}))
}

func TestChainName(t *testing.T) {
	require.Equal(t, chain_selectors.ETHEREUM_MAINNET.Name, chainName(chain_selectors.ETHEREUM_MAINNET.Selector))
	require.Equal(t, "chain-999", chainName(999))
}

func TestWriteLaneSanityResultsCSV(t *testing.T) {
	path := t.TempDir() + "/results.csv"
	t.Setenv(laneSanityCSVOutputEnv, path)

	records := []laneSanityCSVRecord{
		{
			SourceChain:   "ethereum-mainnet",
			DestChain:     "polygon-mainnet",
			FeeToken:      "native",
			TransferToken: "",
			Data:          "lane-sanity-check",
			Fee:           "100",
			ExplorerLink:  "https://ccip.chain.link/msg/0xabc",
		},
		{
			SourceChain:   "ethereum-mainnet",
			DestChain:     "polygon-mainnet",
			FeeToken:      "native",
			TransferToken: "0xlink",
			Data:          "",
			Fee:           "200",
			ExplorerLink:  "https://ccip.chain.link/msg/0xdef",
		},
		{
			SourceChain:   "ethereum-mainnet",
			DestChain:     "polygon-mainnet",
			FeeToken:      "0xusdc",
			TransferToken: "",
			Data:          "lane-sanity-check",
			Error:         "insufficient funds",
		},
	}
	require.NoError(t, writeLaneSanityResultsCSV(logger.Test(t), records))

	f, err := os.Open(path)
	require.NoError(t, err)
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	require.NoError(t, err)
	require.Len(t, rows, 4)
	require.Equal(t, []string{
		"source", "destination", "fee_token", "transfer_token", "data", "fee", "explorer_link", "error",
	}, rows[0])
	require.Equal(t, []string{
		"ethereum-mainnet", "polygon-mainnet", "native", "", "lane-sanity-check", "100", "https://ccip.chain.link/msg/0xabc", "",
	}, rows[1])
	require.Equal(t, []string{
		"ethereum-mainnet", "polygon-mainnet", "native", "0xlink", "", "200", "https://ccip.chain.link/msg/0xdef", "",
	}, rows[2])
	require.Equal(t, []string{
		"ethereum-mainnet", "polygon-mainnet", "0xusdc", "", "lane-sanity-check", "", "", "insufficient funds",
	}, rows[3])
}

func TestWriteLaneSanityResultsCSV_EmptyRecords(t *testing.T) {
	require.NoError(t, writeLaneSanityResultsCSV(logger.Test(t), nil))
}

func TestRunLaneSanityTokenChecksForPair_WritesCSVRecord(t *testing.T) {
	ResetPostProposalLaneSanityRegistryForTest()
	t.Cleanup(func() {
		transferTokenSelector = promptTransferTokenSelection
	})

	eth := chain_selectors.ETHEREUM_MAINNET.Selector
	poly := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()
	adapters := map[uint64]*stubTestAdapter{
		eth:  {selector: eth, family: chain_selectors.FamilyEVM},
		poly: {selector: poly, family: chain_selectors.FamilyEVM},
	}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, adapters)

	csvPath := t.TempDir() + "/lane-sanity.csv"
	t.Setenv(laneSanityCSVOutputEnv, csvPath)

	provider := &stubLaneSanityProvider{
		stubPostProposalProvider: stubPostProposalProvider{
			adapterVersionForLaneFn: func(_ cldf.Environment, _, _ uint64) (*semver.Version, error) {
				return version, nil
			},
		},
		availableTransferTokensFn: func(_ cldf.Environment, _, _ uint64) (map[string]string, error) {
			return map[string]string{"LINK": "0xlink"}, nil
		},
		mockReceiverAddressFn: func(_ cldf.Environment, _ uint64) ([]byte, error) {
			return nil, nil
		},
		fundAndApproveTransferTokenFn: func(_ context.Context, _ cldf.Environment, _ uint64, _ string) (*big.Int, error) {
			return big.NewInt(1), nil
		},
		getMessageFeeFn: func(_ context.Context, _ cldf.Environment, _, _ uint64, _ any) (string, error) {
			return "42", nil
		},
	}
	adapters[eth].sendMessageFn = func(_ context.Context, _ uint64, _ any) (uint64, string, error) {
		return 1, "transfer-id", nil
	}

	transferTokenSelector = func(_ map[string]string, _, _ uint64) ([]string, error) {
		return []string{"0xlink"}, nil
	}

	env := cldf.Environment{
		BlockChains: chain.NewBlockChains(map[uint64]chain.BlockChain{
			eth:  evmchain.Chain{Selector: eth},
			poly: evmchain.Chain{Selector: poly},
		}),
	}

	collector := &laneSanityResultCollector{}
	require.NoError(t, runLaneSanityTokenChecksForPair(
		t.Context(), logger.Test(t), env, chain_selectors.FamilyEVM, provider, eth, poly,
		[]string{"0xlink"}, "", collector,
	))
	require.NoError(t, writeLaneSanityResultsCSV(logger.Test(t), collector.snapshot()))

	rows, err := csv.NewReader(mustOpen(t, csvPath)).ReadAll()
	require.NoError(t, err)
	require.Len(t, rows, 2)
	require.Equal(t, chain_selectors.ETHEREUM_MAINNET.Name, rows[1][0])
	require.Equal(t, chain_selectors.POLYGON_MAINNET.Name, rows[1][1])
	require.Equal(t, "native", rows[1][2])
	require.Equal(t, "0xlink", rows[1][3])
	require.Equal(t, "42", rows[1][5])
	require.Equal(t, "https://ccip.chain.link/msg/0xtransfer-id", rows[1][6])
	require.Empty(t, rows[1][7])
}

func TestRunTokenTransferScenario_RecordsFailureInCSV(t *testing.T) {
	eth := chain_selectors.ETHEREUM_MAINNET.Selector
	poly := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()
	adapters := map[uint64]*stubTestAdapter{
		eth: {selector: eth, family: chain_selectors.FamilyEVM},
	}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, adapters)

	sendErr := errors.New("send rejected")
	provider := &stubLaneSanityProvider{
		fundAndApproveTransferTokenFn: func(_ context.Context, _ cldf.Environment, _ uint64, _ string) (*big.Int, error) {
			return big.NewInt(1), nil
		},
	}
	adapters[eth].sendMessageFn = func(_ context.Context, _ uint64, _ any) (uint64, string, error) {
		return 0, "", sendErr
	}

	env := cldf.Environment{
		BlockChains: chain.NewBlockChains(map[uint64]chain.BlockChain{
			eth: evmchain.Chain{Selector: eth},
		}),
	}
	collector := &laneSanityResultCollector{}
	err := runTokenTransferScenario(
		t.Context(), logger.Test(t), env, adapters[eth], eth, poly,
		"0xlink", []byte("receiver"), []byte("extra"), nil, provider, collector,
	)
	require.ErrorIs(t, err, sendErr)

	records := collector.snapshot()
	require.Len(t, records, 1)
	require.Equal(t, chain_selectors.ETHEREUM_MAINNET.Name, records[0].SourceChain)
	require.Equal(t, chain_selectors.POLYGON_MAINNET.Name, records[0].DestChain)
	require.Equal(t, "0xlink", records[0].TransferToken)
	require.Equal(t, sendErr.Error(), records[0].Error)
	require.Empty(t, records[0].ExplorerLink)
}

func mustOpen(t *testing.T, path string) *os.File {
	t.Helper()
	f, err := os.Open(path)
	require.NoError(t, err)
	t.Cleanup(func() { _ = f.Close() })
	return f
}
