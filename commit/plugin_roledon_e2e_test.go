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

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	clrand "github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readerinternal "github.com/smartcontractkit/chainlink-ccip/internal/reader"
	mockinternalreader "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	mockreader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	ccipreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestPlugin_RoleDonE2E_NoPrevOutcome(t *testing.T) {
	ctx := tests.Context(t)

	s := newRoleDonTestSetup(t, 2, 7, 1)
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
			deps.ccipReader.EXPECT().DiscoverContracts(mock.Anything, mock.Anything).Return(nil, nil)
			deps.ccipReader.EXPECT().Sync(mock.Anything, mock.Anything).Return(nil)
		}

		// Source Chain Expectations - Makes sure only oracles that support specific source chains are reading them.
		{
			if len(oracleSourceChains) > 0 {
				deps.ccipReader.EXPECT().LatestMsgSeqNum(mock.Anything, mock.Anything).
					RunAndReturn(func(ctx context.Context, ch cciptypes.ChainSelector) (cciptypes.SeqNum, error) {
						// should only be called if oracle supports the target source chain
						require.True(t, oracleChains.Contains(ch))
						return cciptypes.SeqNum(55), nil
					})
			}
		}

		// Dest Chain Expectations - Makes sure only oracles that support the destination chain are reading it.
		{
			if oracleChains.Contains(s.destChain) {
				deps.ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).Return(ccipreader.CurseInfo{}, nil)

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

		p := NewPlugin(
			plugintypes.DonID(999),
			s.oracleIDToPeerID,
			pluginconfig.CommitOffchainConfig{
				TokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
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
		p.contractsInitialized.Store(true)

		plugins = append(plugins, p)
	}

	runner := testhelpers.NewOCR3Runner(plugins, s.oracles, ocr3types.Outcome{})
	res, err := runner.RunRound(ctx)
	t.Log(res, err)
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
//	   F Chain      : map[ChainSelector(3577778157919314504):1 ChainSelector(16235373811196386733):1 ChainSelector(16244020411108056671):1]
//	   Chain Oracles: map[ChainSelector(3577778157919314504):[1 2 3] ChainSelector(16235373811196386733):[7 6 4] ChainSelector(16244020411108056671):[5 4 1]]
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
		s.chainOracles[chainSel] = getRandomPermutation(s.oracles, numRequiredOracles+rand.Intn(1))
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
