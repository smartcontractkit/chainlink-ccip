package commit

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"slices"
	"sort"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	clrand "github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readerinternal "github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/ccipocr3"
	mockinternalreader "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	mockreader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	ccipreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func TestPlugin_RoleDonE2E_NoPrevOutcome(t *testing.T) {
	ctx := t.Context()

	s := newRoleDonTestSetup(t, 3, 7, 1)
	const numUnderstaffedChains = 1
	for i := 0; i < numUnderstaffedChains; i++ {
		s.fChain[s.sourceChains[i]] = 2 // <--- we set fChain to 2, that requires at least 2f+1=5 oracles for this chain
		// we set a maximum of 2*fChain+2=4 originally within newRoleDonTestSetup, we expect no results for the chain
	}
	t.Logf("Running test with Role DON Setup:\n%s", s)

	allChainsConfig := make(map[cciptypes.ChainSelector]readerinternal.ChainConfig)
	for _, ch := range append(s.sourceChains, s.destChain) {
		allChainsConfig[ch] = readerinternal.ChainConfig{
			FChain:         s.fChain[ch],
			SupportedNodes: s.oracleIDsToPeerIDsSet(s.chainOracles[ch]),
			Config:         chainconfig.ChainConfig{},
		}
	}

	plugins := make([]ocr3types.ReportingPlugin[[]byte], 0, len(s.oracles))
	for _, oracleID := range s.oracles {
		deps := s.oracleDependencies[oracleID]

		oracleChains := s.getChainsOfOracle(oracleID)
		oracleSourceChainsSet := oracleChains.Clone()
		oracleSourceChainsSet.Remove(s.destChain)
		var oracleSourceChains []cciptypes.ChainSelector
		if oracleSourceChainsSet.Cardinality() > 0 {
			oracleSourceChains = oracleSourceChainsSet.ToSlice()
			sort.Slice(oracleSourceChains, func(i, j int) bool { return oracleSourceChains[i] < oracleSourceChains[j] })
		}

		// Home Chain Expectations - Every oracle should be able to read
		{
			deps.homeChainReader.EXPECT().GetFChain().Return(s.fChain, nil)
			deps.homeChainReader.EXPECT().GetAllChainConfigs().Return(allChainsConfig, nil)
			deps.homeChainReader.EXPECT().GetSupportedChainsForPeer(mock.Anything).
				RunAndReturn(func(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error) {
					supportedChainsOfOracle := s.getChainsOfOracle(s.peerIDToOracleID[id])
					return supportedChainsOfOracle, nil
				})
			deps.homeChainReader.EXPECT().GetChainConfig(mock.Anything).
				RunAndReturn(func(ch cciptypes.ChainSelector) (readerinternal.ChainConfig, error) {
					return allChainsConfig[ch], nil
				})
			deps.homeChainReader.EXPECT().GetKnownCCIPChains().Return(mapset.NewSet(append(s.sourceChains, s.destChain)...), nil)
		}

		// Discovery and Sync - Out of scope of this test case
		{
			deps.ccipReader.EXPECT().DiscoverContracts(mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			deps.ccipReader.EXPECT().Sync(mock.Anything, mock.Anything).Return(nil)
		}

		// Source Chain Expectations - Makes sure only oracles that support specific source chains are reading them.
		{
			if len(oracleSourceChains) > 0 {
				deps.ccipReader.EXPECT().LatestMsgSeqNum(mock.Anything, mock.Anything).
					RunAndReturn(func(ctx context.Context, ch cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
						// should only be called if oracle supports the target source chain
						require.True(t, oracleChains.Contains(ch))
						return cciptypes.SeqNum(66), nil // <--- new msgs
					})
			}
		}

		// Dest Chain Expectations - Makes sure only oracles that support the destination chain are reading it.
		{
			if oracleChains.Contains(s.destChain) {
				deps.ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).Return(cciptypes.CurseInfo{}, nil)

				nextSeqNums := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
				for _, ch := range s.sourceChains {
					nextSeqNums[ch] = cciptypes.SeqNum(56)
				}
				deps.ccipReader.EXPECT().NextSeqNum(mock.Anything, s.sourceChains).Return(nextSeqNums, nil)

				deps.ccipReader.EXPECT().GetRMNRemoteConfig(mock.Anything).Return(cciptypes.RemoteConfig{FSign: 1234}, nil)

				sourceChainsCfg := make(map[cciptypes.ChainSelector]ccipreader.StaticSourceChainConfig)
				for _, ch := range s.sourceChains {
					sourceChainsCfg[ch] = ccipreader.StaticSourceChainConfig{
						Router:                    clrand.RandomBytes(32),
						IsEnabled:                 true,
						IsRMNVerificationDisabled: true,
						OnRamp:                    clrand.RandomBytes(32),
					}
				}
				deps.ccipReader.EXPECT().GetOffRampSourceChainsConfig(mock.Anything, s.sourceChains).Return(sourceChainsCfg, nil)

				deps.priceReader.EXPECT().GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, s.destChain).Return(nil, nil)

				deps.ccipReader.EXPECT().GetChainFeePriceUpdate(mock.Anything, s.sourceChains).Return(nil)
			}
		}

		// Pricing Related Expectations - Makes sure that it reads only chains it supports. related to pricing.
		{
			deps.ccipReader.EXPECT().GetWrappedNativeTokenPriceUSD(mock.Anything, oracleSourceChains).Return(nil)
			if len(oracleSourceChains) > 0 {
				deps.ccipReader.EXPECT().GetChainsFeeComponents(mock.Anything, oracleSourceChains).Return(nil)
			}
		}

		p := s.newRoleDonTestPlugin(oracleID, false)
		plugins = append(plugins, p)
	}

	runner := testhelpers.NewOCR3Runner(plugins, s.oracles, ocr3types.Outcome{})
	res, err := runner.RunRound(ctx)
	require.NoError(t, err)

	o, err := ocrTypCodec.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Equal(t, len(s.sourceChains)-numUnderstaffedChains, len(o.MerkleRootOutcome.RangesSelectedForReport))
}

func TestPlugin_RoleDonE2E_RangesAndPricesSelectedPreviously(t *testing.T) {
	ctx := t.Context()

	s := newRoleDonTestSetup(t, 4, 7, 1)
	chainsWithMsgs := []cciptypes.ChainSelector{s.sourceChains[0], s.sourceChains[1]}
	t.Logf("Running test with Role DON Setup:\n%s", s)

	allChainsConfig := make(map[cciptypes.ChainSelector]readerinternal.ChainConfig)
	for _, ch := range append(s.sourceChains, s.destChain) {
		allChainsConfig[ch] = readerinternal.ChainConfig{
			FChain:         s.fChain[ch],
			SupportedNodes: s.oracleIDsToPeerIDsSet(s.chainOracles[ch]),
			Config:         chainconfig.ChainConfig{},
		}
	}

	plugins := make([]ocr3types.ReportingPlugin[[]byte], 0, len(s.oracles))
	for _, oracleID := range s.oracles {
		deps := s.oracleDependencies[oracleID]

		oracleChains := s.getChainsOfOracle(oracleID)
		oracleSourceChainsSet := oracleChains.Clone()
		oracleSourceChainsSet.Remove(s.destChain)
		var oracleSourceChains []cciptypes.ChainSelector
		if oracleSourceChainsSet.Cardinality() > 0 {
			oracleSourceChains = oracleSourceChainsSet.ToSlice()
			sort.Slice(oracleSourceChains, func(i, j int) bool { return oracleSourceChains[i] < oracleSourceChains[j] })
		}

		// Home Chain Expectations - Every oracle should be able to read
		{
			deps.homeChainReader.EXPECT().GetFChain().Return(s.fChain, nil)
			deps.homeChainReader.EXPECT().GetAllChainConfigs().Return(allChainsConfig, nil)
			deps.homeChainReader.EXPECT().GetSupportedChainsForPeer(mock.Anything).
				RunAndReturn(func(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error) {
					supportedChainsOfOracle := s.getChainsOfOracle(s.peerIDToOracleID[id])
					return supportedChainsOfOracle, nil
				})
			deps.homeChainReader.EXPECT().GetChainConfig(mock.Anything).
				RunAndReturn(func(ch cciptypes.ChainSelector) (readerinternal.ChainConfig, error) {
					return allChainsConfig[ch], nil
				})
		}

		// Discovery and Sync - Out of scope of this test case
		{
			deps.ccipReader.EXPECT().DiscoverContracts(mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			deps.ccipReader.EXPECT().Sync(mock.Anything, mock.Anything).Return(nil)
		}

		// Source Chain Expectations - Makes sure only oracles that support specific source chains are reading them.
		{
			if len(oracleSourceChains) > 0 && mapset.NewSet(oracleSourceChains...).ContainsAny(chainsWithMsgs...) {
				deps.ccipReader.EXPECT().MsgsBetweenSeqNums(mock.Anything, mock.Anything, mock.Anything).
					RunAndReturn(
						func(ctx context.Context, selector cciptypes.ChainSelector, numRange cciptypes.SeqNumRange,
						) ([]cciptypes.Message, error) {
							require.True(t, oracleChains.Contains(selector))
							// merkle roots consensus and computation is out of scope for this test
							return []cciptypes.Message{}, nil
						})
			}
		}

		// Dest Chain Expectations - Makes sure only oracles that support the destination chain are reading it.
		{
			if oracleChains.Contains(s.destChain) {
				deps.ccipReader.EXPECT().GetLatestPriceSeqNr(mock.Anything).Return(53, nil) // <-- still inflight (less than 54)
			}
		}

		p := s.newRoleDonTestPlugin(oracleID, false)
		plugins = append(plugins, p)
	}

	require.True(t, len(s.sourceChains) >= 2, "this test requires at least two chains")
	prevOutcome := committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType:             merkleroot.ReportIntervalsSelected,
			RangesSelectedForReport: []plugintypes.ChainRange{},
		},
		TokenPriceOutcome: tokenprice.Outcome{
			TokenPrices: map[cciptypes.UnknownEncodedAddress]cciptypes.BigInt{
				"0x123": cciptypes.NewBigIntFromInt64(9999),
			},
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: []cciptypes.GasPriceChain{
				{
					ChainSel: s.sourceChains[0],
					GasPrice: cciptypes.NewBigIntFromInt64(1999),
				},
				{
					ChainSel: s.sourceChains[1],
					GasPrice: cciptypes.NewBigIntFromInt64(2999),
				},
			},
		},
		MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 54, RemainingPriceChecks: 10},
	}
	for i, ch := range chainsWithMsgs {
		prevOutcome.MerkleRootOutcome.RangesSelectedForReport = append(
			prevOutcome.MerkleRootOutcome.RangesSelectedForReport,
			plugintypes.ChainRange{
				ChainSel:    ch,
				SeqNumRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(10+i), cciptypes.SeqNum(11+i)),
			})
	}

	b, err := ocrTypCodec.EncodeOutcome(prevOutcome)
	require.NoError(t, err)
	runner := testhelpers.NewOCR3Runner(plugins, s.oracles, b)
	res, err := runner.RunRound(ctx)
	require.NoError(t, err)

	o, err := ocrTypCodec.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Equal(t, committypes.MainOutcome{InflightPriceOcrSequenceNumber: 54, RemainingPriceChecks: 9}, o.MainOutcome)
}

func TestPlugin_RoleDonE2E_Discovery(t *testing.T) {
	ctx := t.Context()

	s := newRoleDonTestSetup(t, 3, 7, 1)
	t.Logf("Running test with Role DON Setup:\n%s", s)

	allChainsConfig := make(map[cciptypes.ChainSelector]readerinternal.ChainConfig)
	for _, ch := range append(s.sourceChains, s.destChain) {
		allChainsConfig[ch] = readerinternal.ChainConfig{
			FChain:         s.fChain[ch],
			SupportedNodes: s.oracleIDsToPeerIDsSet(s.chainOracles[ch]),
			Config:         chainconfig.ChainConfig{},
		}
	}

	plugins := make([]ocr3types.ReportingPlugin[[]byte], 0, len(s.oracles))
	for _, oracleID := range s.oracles {
		deps := s.oracleDependencies[oracleID]

		oracleChains := s.getChainsOfOracle(oracleID)
		oracleSourceChainsSet := oracleChains.Clone()
		oracleSourceChainsSet.Remove(s.destChain)
		var oracleSourceChains []cciptypes.ChainSelector
		if oracleSourceChainsSet.Cardinality() > 0 {
			oracleSourceChains = oracleSourceChainsSet.ToSlice()
			sort.Slice(oracleSourceChains, func(i, j int) bool { return oracleSourceChains[i] < oracleSourceChains[j] })
		}

		// Home Chain Expectations - Every oracle should be able to read
		{
			deps.homeChainReader.EXPECT().GetFChain().Return(s.fChain, nil)
			deps.homeChainReader.EXPECT().GetAllChainConfigs().Return(allChainsConfig, nil)

			deps.homeChainReader.EXPECT().GetSupportedChainsForPeer(mock.Anything).
				RunAndReturn(func(id libocrtypes.PeerID) (mapset.Set[cciptypes.ChainSelector], error) {
					supportedChainsOfOracle := s.getChainsOfOracle(s.peerIDToOracleID[id])
					return supportedChainsOfOracle, nil
				}).Maybe() // once by the leader

			deps.homeChainReader.EXPECT().GetChainConfig(mock.Anything).
				RunAndReturn(func(ch cciptypes.ChainSelector) (readerinternal.ChainConfig, error) {
					return allChainsConfig[ch], nil
				})
		}

		// Discovery and Sync
		{
			deps.ccipReader.EXPECT().DiscoverContracts(mock.Anything, mock.Anything, mock.Anything).
				RunAndReturn(func(ctx context.Context, supportedChains, selectors []cciptypes.ChainSelector,
				) (ccipreader.ContractAddresses, error) {
					addrs := ccipreader.ContractAddresses{}

					// the following contracts can only be discovered by dest supporting oracles
					if oracleChains.Contains(s.destChain) {
						addrs = map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
							consts.ContractNameOffRamp:      {s.destChain: []byte("offramp")},
							consts.ContractNameNonceManager: {s.destChain: []byte("nonceManager")},
							consts.ContractNameRMNRemote:    {s.destChain: []byte("rmnRemote")},
							consts.ContractNameRouter:       {s.destChain: []byte("router")},
							consts.ContractNameFeeQuoter:    {s.destChain: []byte("feeQuoter")},
							consts.ContractNameOnRamp:       {s.sourceChains[0]: []byte("onramp")},
						}
					}

					if oracleChains.Contains(s.sourceChains[0]) {
						if len(addrs) == 0 {
							addrs = map[string]map[cciptypes.ChainSelector]cciptypes.UnknownAddress{
								consts.ContractNameFeeQuoter: {},
								consts.ContractNameRouter:    {},
							}
						}
						addrs[consts.ContractNameFeeQuoter][s.sourceChains[0]] = cciptypes.UnknownAddress("feeQuoter")
						addrs[consts.ContractNameRouter][s.sourceChains[0]] = cciptypes.UnknownAddress("router")
					}

					return addrs, nil
				})

			deps.ccipReader.EXPECT().Sync(mock.Anything, mock.Anything).
				RunAndReturn(func(ctx context.Context, addresses ccipreader.ContractAddresses) error {
					require.Equal(t, ccipreader.ContractAddresses{
						// offramp is synced once and cannot change afterwards
						//consts.ContractNameOffRamp:      {s.destChain: []byte("offramp")},
						consts.ContractNameOnRamp:       {s.sourceChains[0]: []byte("onramp")},
						consts.ContractNameNonceManager: {s.destChain: []byte("nonceManager")},
						consts.ContractNameRMNRemote:    {s.destChain: []byte("rmnRemote")},
						consts.ContractNameRouter:       {s.sourceChains[0]: []byte("router"), s.destChain: []byte("router")},
						consts.ContractNameFeeQuoter:    {s.sourceChains[0]: []byte("feeQuoter"), s.destChain: []byte("feeQuoter")},
					}, addresses)
					return nil
				})
		}

		p := s.newRoleDonTestPlugin(oracleID, true)
		plugins = append(plugins, p)
	}
	runner := testhelpers.NewOCR3Runner(plugins, s.oracles, ocr3types.Outcome{})
	res, err := runner.RunRound(ctx)
	require.NoError(t, err)

	_, err = ocrTypCodec.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
}

func (s roleDonTestSetup) newRoleDonTestPlugin(oracleID commontypes.OracleID, initContracts bool) *Plugin {
	deps := s.oracleDependencies[oracleID]
	p := NewPlugin(
		plugintypes.DonID(999),
		s.oracleIDToPeerID,
		pluginconfig.CommitOffchainConfig{
			TokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				"0x01": {
					AggregatorAddress: "0x02",
					DeviationPPB:      cciptypes.NewBigIntFromInt64(123),
					Decimals:          12,
				},
			},
			PriceFeedChainSelector:          145,
			NewMsgScanBatchSize:             50,
			MaxMerkleTreeSize:               256,
			InflightPriceCheckRetries:       1,
			MerkleRootAsyncObserverDisabled: true,
			ChainFeeAsyncObserverDisabled:   true,
			TokenPriceAsyncObserverDisabled: true,
			DonBreakingChangesVersion:       pluginconfig.DonBreakingChangesVersion1RoleDonSupport,
		},
		s.destChain,
		deps.ccipReader,
		deps.priceReader,
		deps.reportCodec,
		deps.msgHasher,
		deps.lggr,
		deps.homeChainReader,
		deps.rmnHomeReader,
		nil,
		nil,
		ocr3types.ReportingPluginConfig{
			OracleID: oracleID,
			N:        len(s.oracles),
			F:        s.fRoleDon,
		},
		&metrics.Noop{},
		deps.addressCodec,
		deps.reportBuilder,
	)
	p.contractsInitialized.Store(!initContracts)
	return p
}

type roleDonTestSetup struct {
	sourceChains       []cciptypes.ChainSelector
	destChain          cciptypes.ChainSelector
	oracles            []commontypes.OracleID
	fRoleDon           int
	fChain             map[cciptypes.ChainSelector]int
	chainOracles       map[cciptypes.ChainSelector][]commontypes.OracleID
	oracleIDToPeerID   map[commontypes.OracleID]libocrtypes.PeerID
	peerIDToOracleID   map[libocrtypes.PeerID]commontypes.OracleID
	oracleDependencies map[commontypes.OracleID]oracleMockDependencySet
}

func (s roleDonTestSetup) String() string {
	return fmt.Sprintf(`Role DON Test Setup
Source Chains: %v
Dest Chain   : %v
Oracles      : %v
F Role DON   : %v
F Chain      : %v
Chain Oracles: %v
`, s.sourceChains, s.destChain, s.oracles, s.fRoleDon, s.fChain, s.chainOracles)
}

type oracleMockDependencySet struct {
	ccipReader      *mockreader.MockCCIPReader
	priceReader     *mockreader.MockPriceReader
	reportCodec     *ccipocr3.MockCommitPluginCodec
	msgHasher       *mocks.MessageHasher
	lggr            logger.Logger
	homeChainReader *mockinternalreader.MockHomeChain
	rmnHomeReader   *mockreader.MockRMNHome
	addressCodec    *ccipocr3.MockAddressCodec
	reportBuilder   builder.ReportBuilderFunc
}

// This function will create a random role DON test setup based on the provided parameters.
// For example:
//
//	numSourceChains=2
//	numOracles=7
//	   Source Chains: [ChainSelector(3577778157919314504) ChainSelector(16235373811196386733)]
//	   Dest Chain   : ChainSelector(16244020411108056671)
//	   Oracles      : [5 3 7 2 4 1 6]
//	   F Role DON   : 3
//	   F Chain      : map[3577778157919314504:1 16235373811196386733:1 16244020411108056671:1]
//	   Chain Oracles: map[3577778157919314504:[1 2 3] 16235373811196386733:[7 6 4] 16244020411108056671:[5 4 1]]
func newRoleDonTestSetup(t *testing.T, numSourceChains, numOracles, fChain int) roleDonTestSetup {
	s := roleDonTestSetup{}

	if numSourceChains > len(chainsel.ALL)-1 {
		t.Fatal("too many source chains")
	}
	s.sourceChains = make([]cciptypes.ChainSelector, numSourceChains)
	for i := 0; i < numSourceChains; i++ {
		s.sourceChains[i] = cciptypes.ChainSelector(chainsel.ALL[i].Selector)
	}
	sort.Slice(s.sourceChains, func(i, j int) bool { return s.sourceChains[i] < s.sourceChains[j] })

	s.destChain = cciptypes.ChainSelector(chainsel.ALL[len(chainsel.ALL)-1].Selector)

	s.oracles = make([]commontypes.OracleID, numOracles)
	s.oracleIDToPeerID = make(map[commontypes.OracleID]libocrtypes.PeerID, len(s.oracles))
	s.peerIDToOracleID = make(map[libocrtypes.PeerID]commontypes.OracleID, len(s.oracles))
	s.oracleDependencies = make(map[commontypes.OracleID]oracleMockDependencySet)
	for i := range s.oracles {
		s.oracles[i] = commontypes.OracleID(i + 1)
		peerID := libocrtypes.PeerID{byte(s.oracles[i])}
		s.oracleIDToPeerID[s.oracles[i]] = peerID
		s.peerIDToOracleID[peerID] = s.oracles[i]

		reportBuilder, err := builder.NewReportBuilder(false, 0, 0)
		require.NoError(t, err)
		s.oracleDependencies[s.oracles[i]] = oracleMockDependencySet{
			ccipReader:      mockreader.NewMockCCIPReader(t),
			priceReader:     mockreader.NewMockPriceReader(t),
			reportCodec:     ccipocr3.NewMockCommitPluginCodec(t),
			msgHasher:       mocks.NewMessageHasher(),
			lggr:            logger.Test(t),
			homeChainReader: mockinternalreader.NewMockHomeChain(t),
			rmnHomeReader:   mockreader.NewMockRMNHome(t),
			addressCodec:    ccipocr3.NewMockAddressCodec(t),
			reportBuilder:   reportBuilder,
		}
	}

	s.fChain = make(map[cciptypes.ChainSelector]int)
	for _, ch := range append(s.sourceChains, s.destChain) {
		s.fChain[ch] = fChain
	}

	s.fRoleDon = int(math.Floor((float64(len(s.oracles)) - 1.0) / 2.0))

	s.chainOracles = map[cciptypes.ChainSelector][]commontypes.OracleID{}
	for chainSel, f := range s.fChain {
		numRequiredOracles := 2*f + 1
		s.chainOracles[chainSel] = getRandomPermutation(s.oracles, numRequiredOracles+rand.Intn(2))
	}

	return s
}

// creates a copy of the provided slice and returns a random permutation limited to lim items.
// example: getRandomPermutation({1,2,3,4,5}, 2) = {4, 1}
func getRandomPermutation[T any](sl []T, lim int) []T {
	var cp []T
	for i := 0; i < lim; i++ {
		cp = append(cp, sl[i])
	}
	rand.Shuffle(len(sl), func(i, j int) { sl[i], sl[j] = sl[j], sl[i] })
	return cp[:lim]
}

// returns the peerIDs for the provided oracles as a set.
func (s roleDonTestSetup) oracleIDsToPeerIDsSet(oracleIDs []commontypes.OracleID) mapset.Set[libocrtypes.PeerID] {
	pids := make([]libocrtypes.PeerID, len(oracleIDs))
	for i, oracleID := range oracleIDs {
		pids[i] = s.oracleIDToPeerID[oracleID]
	}
	return mapset.NewSet(pids...)
}

// returns the chains that the provided oracle supports
func (s roleDonTestSetup) getChainsOfOracle(oracleID commontypes.OracleID) mapset.Set[cciptypes.ChainSelector] {
	chains := make([]cciptypes.ChainSelector, 0)
	for chainSel, oracles := range s.chainOracles {
		if slices.Contains(oracles, oracleID) {
			chains = append(chains, chainSel)
		}
	}
	return mapset.NewSet(chains...)
}
