package chainfee

import (
	"math/big"
	"math/rand"
	"sort"
	"testing"
	"time"

	"golang.org/x/exp/maps"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	plugincommon2 "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	reader2 "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_processor_Observation(t *testing.T) {
	fourHoursAgo := time.Now().Add(-4 * time.Hour).UTC().Truncate(time.Hour)

	testCases := []struct {
		name                         string
		supportedChains              []ccipocr3.ChainSelector
		chainFeeComponents           map[ccipocr3.ChainSelector]types.ChainFeeComponents
		nativeTokenPrices            map[ccipocr3.ChainSelector]ccipocr3.BigInt
		existingChainFeePriceUpdates map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig
		fChain                       map[ccipocr3.ChainSelector]int
		expectedChainFeePriceUpdates map[ccipocr3.ChainSelector]Update

		dstChain ccipocr3.ChainSelector

		expErr   bool
		emptyObs bool
	}{
		{
			name:            "two chains excluding dest",
			supportedChains: []ccipocr3.ChainSelector{1, 2, 3},
			dstChain:        3,
			chainFeeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(10),
					DataAvailabilityFee: big.NewInt(20),
				},
				2: {
					ExecutionFee:        big.NewInt(100),
					DataAvailabilityFee: big.NewInt(200),
				},
			},
			nativeTokenPrices: map[ccipocr3.ChainSelector]ccipocr3.BigInt{
				1: ccipocr3.NewBigIntFromInt64(1000),
				2: ccipocr3.NewBigIntFromInt64(2000),
			},
			existingChainFeePriceUpdates: map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig{
				1: {
					Timestamp: fourHoursAgo,
					Value: ccipocr3.NewBigInt(FeeComponentsToPackedFee(ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(1234),
						DataAvFeePriceUSD:    big.NewInt(4321),
					})),
				},
				2: {
					Timestamp: fourHoursAgo,
					Value: ccipocr3.NewBigInt(FeeComponentsToPackedFee(ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(12340),
						DataAvFeePriceUSD:    big.NewInt(43210),
					})),
				},
			},
			expectedChainFeePriceUpdates: map[ccipocr3.ChainSelector]Update{
				1: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(1234),
						DataAvFeePriceUSD:    big.NewInt(4321),
					},
					Timestamp: fourHoursAgo,
				},
				2: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(12340),
						DataAvFeePriceUSD:    big.NewInt(43210),
					},
					Timestamp: fourHoursAgo,
				},
			},
			fChain: map[ccipocr3.ChainSelector]int{
				1: 1,
				2: 2,
				3: 1,
			},
			expErr: false,
		},
		{
			name:            "only dest chain",
			supportedChains: []ccipocr3.ChainSelector{1},
			dstChain:        1,
			emptyObs:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs := plugincommon.NewMockChainSupport(t)
			ccipReader := reader.NewMockCCIPReader(t)
			homeChain := reader2.NewMockHomeChain(t)
			oracleID := commontypes.OracleID(rand.Int() % 255)
			lggr := logger.Test(t)
			ctx := tests.Context(t)

			p := &processor{
				lggr:            lggr,
				chainSupport:    cs,
				destChain:       tc.dstChain,
				ccipReader:      ccipReader,
				oracleID:        oracleID,
				homeChain:       homeChain,
				metricsReporter: plugincommon2.NoopReporter{},
				obs:             newBaseObserver(ccipReader, tc.dstChain, oracleID, cs),
			}

			supportedSet := mapset.NewSet(tc.supportedChains...)
			cs.EXPECT().DestChain().Return(tc.dstChain).Maybe()
			cs.EXPECT().SupportedChains(oracleID).
				Return(supportedSet, nil).Maybe()

			supportedSet.Remove(tc.dstChain)
			slicesWithoutDst := supportedSet.ToSlice()
			sort.Slice(slicesWithoutDst, func(i, j int) bool { return slicesWithoutDst[i] < slicesWithoutDst[j] })

			if len(slicesWithoutDst) == 0 {
				slicesWithoutDst = []ccipocr3.ChainSelector(nil)
			}

			ccipReader.EXPECT().GetChainsFeeComponents(ctx, slicesWithoutDst).
				Return(tc.chainFeeComponents).Maybe()

			ccipReader.EXPECT().GetWrappedNativeTokenPriceUSD(ctx, slicesWithoutDst).
				Return(tc.nativeTokenPrices).Maybe()

			ccipReader.EXPECT().GetChainFeePriceUpdate(ctx, slicesWithoutDst).
				Return(tc.existingChainFeePriceUpdates).Maybe()

			homeChain.EXPECT().GetFChain().Return(tc.fChain, nil).Maybe()

			tStart := time.Now()
			obs, err := p.Observation(ctx, Outcome{}, Query{})
			tEnd := time.Now()
			if tc.expErr {
				require.Error(t, err)
				return
			}
			if tc.emptyObs {
				require.Empty(t, obs)
				return
			}

			require.NoError(t, err)
			require.GreaterOrEqual(t, obs.TimestampNow.UnixNano(), tStart.UnixNano())
			require.LessOrEqual(t, obs.TimestampNow.UnixNano(), tEnd.UnixNano())
			require.Equal(t, tc.chainFeeComponents, obs.FeeComponents)
			require.ElementsMatch(t, slicesWithoutDst, maps.Keys(obs.FeeComponents))
			require.Equal(t, tc.nativeTokenPrices, obs.NativeTokenPrices)
			require.ElementsMatch(t, slicesWithoutDst, maps.Keys(obs.NativeTokenPrices))
			require.Equal(t, tc.expectedChainFeePriceUpdates, obs.ChainFeeUpdates)
			require.ElementsMatch(t, slicesWithoutDst, maps.Keys(obs.ChainFeeUpdates))
			require.Equal(t, tc.fChain, obs.FChain)
		})
	}
}

func Test_unique_chain_filter_in_Observation(t *testing.T) {
	fourHoursAgo := time.Now().Add(-4 * time.Hour).UTC().Truncate(time.Hour)

	testCases := []struct {
		name                         string
		supportedChains              []ccipocr3.ChainSelector
		chainFeeComponents           map[ccipocr3.ChainSelector]types.ChainFeeComponents
		nativeTokenPrices            map[ccipocr3.ChainSelector]ccipocr3.BigInt
		existingChainFeePriceUpdates map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig
		fChain                       map[ccipocr3.ChainSelector]int
		dstChain                     ccipocr3.ChainSelector
		expUniqueChains              int
	}{
		{
			name:            "unique chains intersection",
			supportedChains: []ccipocr3.ChainSelector{1, 2, 3},
			dstChain:        3,
			chainFeeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(10),
					DataAvailabilityFee: big.NewInt(20),
				},
				2: {
					ExecutionFee:        big.NewInt(100),
					DataAvailabilityFee: big.NewInt(200),
				},
			},
			nativeTokenPrices: map[ccipocr3.ChainSelector]ccipocr3.BigInt{
				1: ccipocr3.NewBigIntFromInt64(1000),
				2: ccipocr3.NewBigIntFromInt64(2000),
			},
			existingChainFeePriceUpdates: map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig{
				1: {
					Timestamp: fourHoursAgo,
					Value: ccipocr3.NewBigInt(FeeComponentsToPackedFee(ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(1234),
						DataAvFeePriceUSD:    big.NewInt(4321),
					})),
				},
				2: {
					Timestamp: fourHoursAgo,
					Value: ccipocr3.NewBigInt(FeeComponentsToPackedFee(ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(12340),
						DataAvFeePriceUSD:    big.NewInt(43210),
					})),
				},
			},
			fChain: map[ccipocr3.ChainSelector]int{
				1: 1,
				2: 2,
				3: 1,
			},
			expUniqueChains: 2,
		},
		{
			name:            "only one unique chain between fee components and native token prices",
			supportedChains: []ccipocr3.ChainSelector{1, 2, 3},
			dstChain:        3,
			chainFeeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(10),
					DataAvailabilityFee: big.NewInt(20),
				},
				2: {
					ExecutionFee:        big.NewInt(100),
					DataAvailabilityFee: big.NewInt(200),
				},
			},
			nativeTokenPrices: map[ccipocr3.ChainSelector]ccipocr3.BigInt{
				1: ccipocr3.NewBigIntFromInt64(1000),
			},
			existingChainFeePriceUpdates: map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig{
				1: {
					Timestamp: fourHoursAgo,
					Value: ccipocr3.NewBigInt(FeeComponentsToPackedFee(ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(1234),
						DataAvFeePriceUSD:    big.NewInt(4321),
					})),
				},
				3: {
					Timestamp: fourHoursAgo,
					Value: ccipocr3.NewBigInt(FeeComponentsToPackedFee(ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(1234),
						DataAvFeePriceUSD:    big.NewInt(4321),
					})),
				},
			},
			fChain: map[ccipocr3.ChainSelector]int{
				1: 1,
				2: 2,
				3: 1,
			},
			expUniqueChains: 1,
		},
		{
			name:            "zero unique chains between fee components and native token prices",
			supportedChains: []ccipocr3.ChainSelector{1, 2, 3},
			dstChain:        3,
			chainFeeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(10),
					DataAvailabilityFee: big.NewInt(20),
				},
				2: {
					ExecutionFee:        big.NewInt(100),
					DataAvailabilityFee: big.NewInt(200),
				},
			},
			nativeTokenPrices: map[ccipocr3.ChainSelector]ccipocr3.BigInt{
				3: ccipocr3.NewBigIntFromInt64(1000),
			},
			existingChainFeePriceUpdates: map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig{
				3: {
					Timestamp: fourHoursAgo,
					Value: ccipocr3.NewBigInt(FeeComponentsToPackedFee(ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(1234),
						DataAvFeePriceUSD:    big.NewInt(4321),
					})),
				},
			},
			fChain: map[ccipocr3.ChainSelector]int{
				1: 1,
				2: 2,
				3: 1,
			},
			expUniqueChains: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs := plugincommon.NewMockChainSupport(t)
			ccipReader := reader.NewMockCCIPReader(t)
			homeChain := reader2.NewMockHomeChain(t)
			oracleID := commontypes.OracleID(rand.Int() % 255)
			lggr := logger.Test(t)
			ctx := tests.Context(t)

			p := &processor{
				lggr:            lggr,
				chainSupport:    cs,
				destChain:       tc.dstChain,
				ccipReader:      ccipReader,
				oracleID:        oracleID,
				homeChain:       homeChain,
				metricsReporter: plugincommon2.NoopReporter{},
				obs:             newBaseObserver(ccipReader, tc.dstChain, oracleID, cs),
			}

			supportedSet := mapset.NewSet(tc.supportedChains...)
			cs.EXPECT().DestChain().Return(tc.dstChain).Maybe()
			cs.EXPECT().SupportedChains(oracleID).
				Return(supportedSet, nil).Maybe()

			supportedSet.Remove(tc.dstChain)
			slicesWithoutDst := supportedSet.ToSlice()
			sort.Slice(slicesWithoutDst, func(i, j int) bool { return slicesWithoutDst[i] < slicesWithoutDst[j] })

			ccipReader.EXPECT().GetChainsFeeComponents(ctx, slicesWithoutDst).
				Return(tc.chainFeeComponents).Maybe()

			ccipReader.EXPECT().GetWrappedNativeTokenPriceUSD(ctx, slicesWithoutDst).
				Return(tc.nativeTokenPrices).Maybe()

			ccipReader.EXPECT().GetChainFeePriceUpdate(ctx, slicesWithoutDst).
				Return(tc.existingChainFeePriceUpdates).Maybe()

			homeChain.EXPECT().GetFChain().Return(tc.fChain, nil).Maybe()

			obs, err := p.Observation(ctx, Outcome{}, Query{})
			require.NoError(t, err)
			if tc.expUniqueChains == 0 {
				require.Empty(t, obs)
				return
			}

			require.True(t, tc.expUniqueChains == len(maps.Keys(obs.FeeComponents)))
			require.True(t, tc.expUniqueChains == len(maps.Keys(obs.NativeTokenPrices)))
			require.ElementsMatch(t, maps.Keys(obs.FeeComponents), maps.Keys(obs.NativeTokenPrices))
		})
	}
}
