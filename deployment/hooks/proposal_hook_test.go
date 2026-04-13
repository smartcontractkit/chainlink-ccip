package hooks

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"os"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evmchain "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_changeset "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/changeset"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

type stubPostProposalProvider struct {
	skipSendFn              func(env cldf_changeset.ProposalHookEnv) bool
	preSendValidationFn     func(env cldf.Environment, srcSel uint64) error
	supportedFeeTokensFn    func(env cldf.Environment, srcSel uint64, forkContext cldf_changeset.ForkContext) ([]string, error)
	supportedDestinationsFn func(env cldf.Environment, srcSel uint64) ([]uint64, error)
	adapterVersionForLaneFn func(env cldf.Environment, srcSel, destSel uint64) (*semver.Version, error)
}

func (s *stubPostProposalProvider) SkipSend(env cldf_changeset.ProposalHookEnv) bool {
	if s.skipSendFn != nil {
		return s.skipSendFn(env)
	}
	return false
}

func (s *stubPostProposalProvider) PreSendValidation(env cldf.Environment, srcSel uint64) error {
	if s.preSendValidationFn != nil {
		return s.preSendValidationFn(env, srcSel)
	}
	return nil
}

func (s *stubPostProposalProvider) SupportedFeeTokens(
	env cldf.Environment,
	srcSel uint64,
	forkContext cldf_changeset.ForkContext,
) ([]string, error) {
	if s.supportedFeeTokensFn != nil {
		return s.supportedFeeTokensFn(env, srcSel, forkContext)
	}
	return []string{""}, nil
}

func (s *stubPostProposalProvider) SupportedDestinations(env cldf.Environment, srcSel uint64) ([]uint64, error) {
	if s.supportedDestinationsFn != nil {
		return s.supportedDestinationsFn(env, srcSel)
	}
	return nil, nil
}

func (s *stubPostProposalProvider) AdapterVersionForLane(env cldf.Environment, srcSel, destSel uint64) (*semver.Version, error) {
	if s.adapterVersionForLaneFn != nil {
		return s.adapterVersionForLaneFn(env, srcSel, destSel)
	}
	return nil, nil
}

type stubTestAdapter struct {
	selector                  uint64
	family                    string
	receiver                  []byte
	buildMessageFn            func(testadapters.MessageComponents) (any, error)
	sendMessageFn             func(ctx context.Context, destChainSelector uint64, msg any) (uint64, string, error)
	getExtraArgsFn            func(receiver []byte, sourceFamily string, opts ...testadapters.ExtraArgOpt) ([]byte, error)
	builtMessages             []testadapters.MessageComponents
	lastExtraArgsSourceFamily string
	lastExtraArgsReceiver     []byte
}

func (s *stubTestAdapter) ChainSelector() uint64 {
	return s.selector
}

func (s *stubTestAdapter) Family() string {
	if s.family == "" {
		return chain_selectors.FamilyEVM
	}
	return s.family
}

func (s *stubTestAdapter) BuildMessage(components testadapters.MessageComponents) (any, error) {
	s.builtMessages = append(s.builtMessages, components)
	if s.buildMessageFn != nil {
		return s.buildMessageFn(components)
	}
	return "mock-msg", nil
}

func (s *stubTestAdapter) SendMessage(ctx context.Context, destChainSelector uint64, msg any) (uint64, string, error) {
	if s.sendMessageFn != nil {
		return s.sendMessageFn(ctx, destChainSelector, msg)
	}
	return 1, "msg-id", nil
}

func (s *stubTestAdapter) CCIPReceiver() []byte {
	if s.receiver != nil {
		return s.receiver
	}
	return []byte("receiver")
}

func (s *stubTestAdapter) EOAReceiver(t *testing.T) []byte {
	return nil
}

func (s *stubTestAdapter) InvalidAddresses() [][]byte {
	return nil
}

func (s *stubTestAdapter) SetReceiverRejectAll(ctx context.Context, t *testing.T, rejectAll bool) error {
	return nil
}

func (s *stubTestAdapter) NativeFeeToken() string {
	return ""
}

func (s *stubTestAdapter) GetExtraArgs(receiver []byte, sourceFamily string, opts ...testadapters.ExtraArgOpt) ([]byte, error) {
	s.lastExtraArgsSourceFamily = sourceFamily
	s.lastExtraArgsReceiver = receiver
	if s.getExtraArgsFn != nil {
		return s.getExtraArgsFn(receiver, sourceFamily, opts...)
	}
	return []byte("extra"), nil
}

func (s *stubTestAdapter) LowGasLimit() *big.Int {
	return big.NewInt(0)
}

func (s *stubTestAdapter) GetInboundNonce(ctx context.Context, sender []byte, srcSel uint64) (uint64, error) {
	return 0, nil
}

func (s *stubTestAdapter) ValidateCommit(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNumRange ccipocr3.SeqNumRange) {
}

func (s *stubTestAdapter) ValidateExecSucceeds(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNrs []uint64) map[uint64]int {
	return nil
}

func (s *stubTestAdapter) ValidateExecFails(t *testing.T, sourceSelector uint64, startBlock *uint64, seqNrs []uint64) {
}

func (s *stubTestAdapter) AllowRouterToWithdrawTokens(ctx context.Context, tokenAddress string, amount *big.Int) error {
	return nil
}

func (s *stubTestAdapter) GetTokenBalance(ctx context.Context, tokenAddress string, ownerAddress []byte) (*big.Int, error) {
	return big.NewInt(0), nil
}

func (s *stubTestAdapter) GetTokenExpansionConfig() (*tokensapi.TokenExpansionInputPerChain, error) {
	return nil, nil
}

func (s *stubTestAdapter) GetRegistryAddress() (string, error) {
	return "", nil
}

func (s *stubTestAdapter) CurrentBlock(t *testing.T) uint64 {
	return 0
}

var adapterVersionCounter atomic.Uint64

func newUniqueAdapterVersion() *semver.Version {
	return semver.MustParse(fmt.Sprintf("99.0.%d", adapterVersionCounter.Add(1)))
}

func registerFactoryForVersion(t *testing.T, family string, version *semver.Version, adapters map[uint64]*stubTestAdapter) {
	t.Helper()
	testadapters.GetTestAdapterRegistry().RegisterTestAdapter(
		family,
		version,
		func(_ *cldf.Environment, selector uint64) testadapters.TestAdapter {
			if ad, ok := adapters[selector]; ok {
				return ad
			}
			return &stubTestAdapter{selector: selector, family: family}
		},
	)
}

func writeMinimalEnvDatastore(t *testing.T, dom domain.Domain, env string) {
	t.Helper()
	envDir := dom.EnvDir(env)
	require.NoError(t, os.MkdirAll(envDir.DataStoreDirPath(), 0o755))
	require.NoError(t, os.WriteFile(envDir.AddressRefsFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ChainMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.ContractMetadataFilePath(), []byte("[]"), 0o600))
	require.NoError(t, os.WriteFile(envDir.EnvMetadataFilePath(), []byte("null"), 0o600))
}

func successfulTimelockReport(selector uint64) cldf_changeset.MCMSTimelockExecuteReport {
	return cldf_changeset.MCMSTimelockExecuteReport{
		Type:   cldf_changeset.MCMSTimelockExecuteReportType,
		Status: "SUCCESS",
		Input: cldf_changeset.MCMSTimelockExecuteReportInput{
			ChainSelector: selector,
		},
	}
}

func newTestHookEnvAndDataStore(t *testing.T) (cldf_changeset.ProposalHookEnv, datastore.DataStore) {
	ds := datastore.NewMemoryDataStore()
	return cldf_changeset.ProposalHookEnv{
		Name:   "test-env",
		Logger: logger.Test(t),
	}, ds.Seal()
}

func runWithTestHookEnv(
	t *testing.T,
	family string,
	provider PostProposalCCIPSend,
	srcSelectors []uint64,
) error {
	t.Helper()
	hookEnv, ds := newTestHookEnvAndDataStore(t)
	return runPostProposalCCIPSends(t.Context(), logger.Test(t), hookEnv, ds, family, provider, srcSelectors)
}

func TestPostProposalCCIPSendRegistry_RegisterFirstWinsAndReset(t *testing.T) {
	ResetPostProposalCCIPSendRegistryForTest()

	reg := GetPostProposalCCIPSendRegistry()
	first := &stubPostProposalProvider{}
	second := &stubPostProposalProvider{}
	reg.Register(chain_selectors.FamilyEVM, first)
	reg.Register(chain_selectors.FamilyEVM, second)

	got, ok := reg.Get(chain_selectors.FamilyEVM)
	require.True(t, ok)
	require.Same(t, first, got)

	ResetPostProposalCCIPSendRegistryForTest()
	_, ok = GetPostProposalCCIPSendRegistry().Get(chain_selectors.FamilyEVM)
	require.False(t, ok)
}

func TestGlobalPostProposalCCIPSendHook_Metadata(t *testing.T) {
	h := GlobalPostProposalCCIPSendHook(domain.NewDomain(t.TempDir(), "test"))
	require.Equal(t, PostProposalCCIPSendHookName, h.Name)
	require.Equal(t, cldf_changeset.Abort, h.FailurePolicy)
	require.Equal(t, 5*time.Minute, h.Timeout)
	require.NotNil(t, h.Func)
}

func TestGroupSelectorsByFamily_SkipsInvalidSelectors(t *testing.T) {
	eth := chain_selectors.ETHEREUM_MAINNET.Selector
	poly := chain_selectors.POLYGON_MAINNET.Selector
	sol := chain_selectors.SOLANA_MAINNET.Selector

	grouped := groupSelectorsByFamily(logger.Test(t), []uint64{eth, 0, poly, sol})
	require.Len(t, grouped, 2)
	require.ElementsMatch(t, []uint64{eth, poly}, grouped[chain_selectors.FamilyEVM])
	require.Equal(t, []uint64{sol}, grouped[chain_selectors.FamilySolana])
}

func TestVerifyCCIPSend_DataStoreError(t *testing.T) {
	ResetPostProposalCCIPSendRegistryForTest()
	dom := domain.NewDomain(t.TempDir(), "test")
	fn := verifyCCIPSend(dom)

	err := fn(t.Context(), cldf_changeset.PostProposalHookParams{
		Env: cldf_changeset.ProposalHookEnv{
			Name:   "missing-env",
			Logger: logger.Test(t),
			BlockChains: chain.NewBlockChains(map[uint64]chain.BlockChain{
				chain_selectors.ETHEREUM_MAINNET.Selector: evmchain.Chain{},
			}),
		},
		Reports: []cldf_changeset.MCMSTimelockExecuteReport{
			successfulTimelockReport(chain_selectors.ETHEREUM_MAINNET.Selector),
		},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, `verify-ccip-send: datastore for env "missing-env"`)
}

func TestVerifyCCIPSend_NoRegisteredProvider_Skips(t *testing.T) {
	ResetPostProposalCCIPSendRegistryForTest()

	dom := domain.NewDomain(t.TempDir(), "test")
	writeMinimalEnvDatastore(t, dom, "staging")

	fn := verifyCCIPSend(dom)
	err := fn(t.Context(), cldf_changeset.PostProposalHookParams{
		Env: cldf_changeset.ProposalHookEnv{
			Name:   "staging",
			Logger: logger.Test(t),
		},
		Reports: []cldf_changeset.MCMSTimelockExecuteReport{
			successfulTimelockReport(chain_selectors.ETHEREUM_MAINNET.Selector),
		},
	})
	require.NoError(t, err)
}

func TestVerifyCCIPSend_AggregatesErrorsByFamily(t *testing.T) {
	ResetPostProposalCCIPSendRegistryForTest()
	reg := GetPostProposalCCIPSendRegistry()
	reg.Register(chain_selectors.FamilyEVM, &stubPostProposalProvider{
		preSendValidationFn: func(env cldf.Environment, srcSel uint64) error {
			return errors.New("evm pre-send")
		},
	})
	reg.Register(chain_selectors.FamilySolana, &stubPostProposalProvider{
		preSendValidationFn: func(env cldf.Environment, srcSel uint64) error {
			return errors.New("solana pre-send")
		},
	})

	dom := domain.NewDomain(t.TempDir(), "test")
	writeMinimalEnvDatastore(t, dom, "staging")

	fn := verifyCCIPSend(dom)
	err := fn(t.Context(), cldf_changeset.PostProposalHookParams{
		Env: cldf_changeset.ProposalHookEnv{
			Name:   "staging",
			Logger: logger.Test(t),
			BlockChains: chain.NewBlockChains(map[uint64]chain.BlockChain{
				chain_selectors.ETHEREUM_MAINNET.Selector: evmchain.Chain{},
				chain_selectors.SOLANA_MAINNET.Selector:   solana.Chain{},
			}),
		},
		Reports: []cldf_changeset.MCMSTimelockExecuteReport{
			successfulTimelockReport(chain_selectors.ETHEREUM_MAINNET.Selector),
			successfulTimelockReport(chain_selectors.SOLANA_MAINNET.Selector),
		},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "family evm")
	require.ErrorContains(t, err, "evm pre-send")
	require.ErrorContains(t, err, "family solana")
	require.ErrorContains(t, err, "solana pre-send")
}

func TestRunPostProposalCCIPSends_SkipSendSkips(t *testing.T) {
	provider := &stubPostProposalProvider{
		skipSendFn: func(env cldf_changeset.ProposalHookEnv) bool {
			return true
		},
	}
	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{1})
	require.NoError(t, err)
}

func TestRunPostProposalCCIPSends_PreSendAndDiscoveryErrors(t *testing.T) {
	preErr := errors.New("pre failed")
	destsErr := errors.New("dests failed")
	feeErr := errors.New("fee tokens failed")

	provider := &stubPostProposalProvider{
		preSendValidationFn: func(env cldf.Environment, srcSel uint64) error {
			if srcSel == 1 {
				return preErr
			}
			return nil
		},
		supportedDestinationsFn: func(env cldf.Environment, srcSel uint64) ([]uint64, error) {
			switch srcSel {
			case 2:
				return nil, destsErr
			case 3:
				return nil, nil
			case 4:
				return []uint64{chain_selectors.POLYGON_MAINNET.Selector}, nil
			default:
				return nil, nil
			}
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSel uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			if srcSel == 4 {
				return nil, feeErr
			}
			return []string{""}, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{1, 2, 3, 4})
	require.Error(t, err)
	require.ErrorContains(t, err, preErr.Error())
	require.ErrorContains(t, err, destsErr.Error())
	require.ErrorContains(t, err, feeErr.Error())
}

func TestRunPostProposalCCIPSends_AdapterVersionError(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.POLYGON_MAINNET.Selector
	adapterVersionErr := errors.New("adapter version failed")

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{""}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return nil, adapterVersionErr
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.Error(t, err)
	require.ErrorContains(t, err, "adapter version src")
	require.ErrorContains(t, err, adapterVersionErr.Error())
}

func TestRunPostProposalCCIPSends_MissingSourceAdapter(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{""}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.Error(t, err)
	require.ErrorContains(t, err, "no test adapter for family evm version")
	require.ErrorContains(t, err, version.String())
}

func TestRunPostProposalCCIPSends_InvalidDestinationSelector(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	version := newUniqueAdapterVersion()

	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, map[uint64]*stubTestAdapter{
		srcSel: {selector: srcSel, family: chain_selectors.FamilyEVM},
	})

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{0}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{""}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.Error(t, err)
	require.ErrorContains(t, err, "dest selector 0")
}

func TestRunPostProposalCCIPSends_MissingCrossFamilyAdapter(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.SOLANA_MAINNET.Selector
	version := newUniqueAdapterVersion()

	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, map[uint64]*stubTestAdapter{
		srcSel: {selector: srcSel, family: chain_selectors.FamilyEVM},
	})

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{""}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.Error(t, err)
	require.ErrorContains(t, err, "no test adapter for dest family solana version")
	require.ErrorContains(t, err, version.String())
}

func TestRunPostProposalCCIPSends_ExtraArgsError(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()
	extraArgsErr := errors.New("extra args failed")

	srcAdapter := &stubTestAdapter{selector: srcSel, family: chain_selectors.FamilyEVM}
	destAdapter := &stubTestAdapter{
		selector: destSel,
		family:   chain_selectors.FamilyEVM,
		getExtraArgsFn: func(receiver []byte, sourceFamily string, opts ...testadapters.ExtraArgOpt) ([]byte, error) {
			return nil, extraArgsErr
		},
	}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, map[uint64]*stubTestAdapter{
		srcSel:  srcAdapter,
		destSel: destAdapter,
	})

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{""}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.Error(t, err)
	require.ErrorContains(t, err, "extra args for src")
	require.ErrorContains(t, err, extraArgsErr.Error())
}

func TestRunPostProposalCCIPSends_BuildMessageError(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()
	buildErr := errors.New("build failed")

	srcAdapter := &stubTestAdapter{
		selector: srcSel,
		family:   chain_selectors.FamilyEVM,
		buildMessageFn: func(components testadapters.MessageComponents) (any, error) {
			return nil, buildErr
		},
	}
	destAdapter := &stubTestAdapter{selector: destSel, family: chain_selectors.FamilyEVM}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, map[uint64]*stubTestAdapter{
		srcSel:  srcAdapter,
		destSel: destAdapter,
	})

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{""}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.Error(t, err)
	require.ErrorContains(t, err, "build message src")
	require.ErrorContains(t, err, buildErr.Error())
}

func TestRunPostProposalCCIPSends_SendMessageError(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()
	sendErr := errors.New("send failed")

	srcAdapter := &stubTestAdapter{
		selector: srcSel,
		family:   chain_selectors.FamilyEVM,
		sendMessageFn: func(ctx context.Context, destChainSelector uint64, msg any) (uint64, string, error) {
			return 0, "", sendErr
		},
	}
	destAdapter := &stubTestAdapter{selector: destSel, family: chain_selectors.FamilyEVM}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, map[uint64]*stubTestAdapter{
		srcSel:  srcAdapter,
		destSel: destAdapter,
	})

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{""}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.Error(t, err)
	require.ErrorContains(t, err, "CCIP send from")
	require.ErrorContains(t, err, sendErr.Error())
}

func TestRunPostProposalCCIPSends_UsesNativeFeeWhenProviderReturnsNone(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.POLYGON_MAINNET.Selector
	version := newUniqueAdapterVersion()

	srcAdapter := &stubTestAdapter{selector: srcSel, family: chain_selectors.FamilyEVM}
	destAdapter := &stubTestAdapter{selector: destSel, family: chain_selectors.FamilyEVM}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, map[uint64]*stubTestAdapter{
		srcSel:  srcAdapter,
		destSel: destAdapter,
	})

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return nil, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.NoError(t, err)
	require.Len(t, srcAdapter.builtMessages, 1)
	require.Equal(t, "", srcAdapter.builtMessages[0].FeeToken)
}

func TestRunPostProposalCCIPSends_CrossFamilyHappyPath(t *testing.T) {
	srcSel := chain_selectors.ETHEREUM_MAINNET.Selector
	destSel := chain_selectors.SOLANA_MAINNET.Selector
	version := newUniqueAdapterVersion()

	srcAdapter := &stubTestAdapter{selector: srcSel, family: chain_selectors.FamilyEVM}
	destAdapter := &stubTestAdapter{
		selector: destSel,
		family:   chain_selectors.FamilySolana,
		receiver: []byte("solana-receiver"),
		getExtraArgsFn: func(receiver []byte, sourceFamily string, opts ...testadapters.ExtraArgOpt) ([]byte, error) {
			return []byte("solana-extra"), nil
		},
	}
	registerFactoryForVersion(t, chain_selectors.FamilyEVM, version, map[uint64]*stubTestAdapter{
		srcSel: srcAdapter,
	})
	registerFactoryForVersion(t, chain_selectors.FamilySolana, version, map[uint64]*stubTestAdapter{
		destSel: destAdapter,
	})

	provider := &stubPostProposalProvider{
		supportedDestinationsFn: func(env cldf.Environment, srcSelArg uint64) ([]uint64, error) {
			return []uint64{destSel}, nil
		},
		supportedFeeTokensFn: func(env cldf.Environment, srcSelArg uint64, forkContext cldf_changeset.ForkContext) ([]string, error) {
			return []string{"fee-token"}, nil
		},
		adapterVersionForLaneFn: func(env cldf.Environment, srcSelArg, destSelArg uint64) (*semver.Version, error) {
			return version, nil
		},
	}

	err := runWithTestHookEnv(t, chain_selectors.FamilyEVM, provider, []uint64{srcSel})
	require.NoError(t, err)
	require.Len(t, srcAdapter.builtMessages, 1)
	require.Equal(t, destSel, srcAdapter.builtMessages[0].DestChainSelector)
	require.Equal(t, []byte("solana-receiver"), srcAdapter.builtMessages[0].Receiver)
	require.Equal(t, []byte("solana-extra"), srcAdapter.builtMessages[0].ExtraArgs)
	require.Equal(t, "fee-token", srcAdapter.builtMessages[0].FeeToken)
	require.Equal(t, chain_selectors.FamilyEVM, destAdapter.lastExtraArgsSourceFamily)
}
