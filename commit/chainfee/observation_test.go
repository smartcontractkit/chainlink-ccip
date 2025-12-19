package chainfee

import (
	"context"
	"maps"
	"math/big"
	"math/rand"
	"slices"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	plugincommon2 "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	reader2 "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	reader3 "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
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
			expErr:   false,
			emptyObs: false,
		},
		{
			name:            "only dest chain",
			supportedChains: []ccipocr3.ChainSelector{1},
			dstChain:        1,
			fChain: map[ccipocr3.ChainSelector]int{
				1: 1,
				2: 2,
				3: 1,
			},
			emptyObs: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs := plugincommon.NewMockChainSupport(t)
			ccipReader := reader.NewMockCCIPReader(t)
			homeChain := reader2.NewMockHomeChain(t)
			oracleID := commontypes.OracleID(rand.Int() % 255)
			lggr := logger.Test(t)
			ctx := t.Context()

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
			cs.EXPECT().SupportedChains(oracleID).Return(supportedSet, nil).Maybe()

			supportedSet.Remove(tc.dstChain)
			slicesWithoutDst := supportedSet.ToSlice()
			sort.Slice(slicesWithoutDst, func(i, j int) bool { return slicesWithoutDst[i] < slicesWithoutDst[j] })

			if len(slicesWithoutDst) == 0 {
				slicesWithoutDst = []ccipocr3.ChainSelector(nil)
			}

			cs.EXPECT().KnownSourceChainsSlice().Return(slicesWithoutDst, nil).Maybe()
			srcChainsCfg := make(map[ccipocr3.ChainSelector]reader3.StaticSourceChainConfig, len(slicesWithoutDst))
			for _, chain := range slicesWithoutDst {
				srcChainsCfg[chain] = reader3.StaticSourceChainConfig{
					IsEnabled: true,
				}
			}
			ccipReader.EXPECT().GetOffRampSourceChainsConfig(mock.Anything, slicesWithoutDst).
				Return(srcChainsCfg, nil).Maybe()

			ccipReader.EXPECT().GetChainsFeeComponents(mock.Anything, slicesWithoutDst).
				Return(tc.chainFeeComponents).Maybe()

			ccipReader.EXPECT().GetWrappedNativeTokenPriceUSD(mock.Anything, slicesWithoutDst).
				Return(tc.nativeTokenPrices).Maybe()

			ccipReader.EXPECT().GetChainFeePriceUpdate(mock.Anything, slicesWithoutDst).
				Return(tc.existingChainFeePriceUpdates).Maybe()

			cs.EXPECT().SupportsDestChain(oracleID).Return(true, nil).Maybe()

			homeChain.EXPECT().GetFChain().Return(tc.fChain, nil).Maybe()

			tStart := time.Now()
			obs, err := p.Observation(ctx, Outcome{}, Query{})
			tEnd := time.Now()
			if tc.expErr {
				require.Error(t, err)
				return
			}
			if tc.emptyObs {
				require.Equal(t, tc.fChain, obs.FChain)
				require.NotEqual(t, time.Time{}, obs.TimestampNow)
				return
			}

			require.NoError(t, err)
			require.GreaterOrEqual(t, obs.TimestampNow.UnixNano(), tStart.UnixNano())
			require.LessOrEqual(t, obs.TimestampNow.UnixNano(), tEnd.UnixNano())
			require.Equal(t, tc.chainFeeComponents, obs.FeeComponents)
			require.ElementsMatch(t, slicesWithoutDst, slices.Collect(maps.Keys(obs.FeeComponents)))
			require.Equal(t, tc.nativeTokenPrices, obs.NativeTokenPrices)
			require.ElementsMatch(t, slicesWithoutDst, slices.Collect(maps.Keys(obs.NativeTokenPrices)))
			require.Equal(t, tc.expectedChainFeePriceUpdates, obs.ChainFeeUpdates)
			require.ElementsMatch(t, slicesWithoutDst, slices.Collect(maps.Keys(obs.ChainFeeUpdates)))
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
			ctx := t.Context()

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
			cs.EXPECT().SupportsDestChain(oracleID).Return(true, nil).Maybe()

			supportedSet.Remove(tc.dstChain)
			slicesWithoutDst := supportedSet.ToSlice()
			sort.Slice(slicesWithoutDst, func(i, j int) bool { return slicesWithoutDst[i] < slicesWithoutDst[j] })

			cs.EXPECT().KnownSourceChainsSlice().Return(slicesWithoutDst, nil).Maybe()
			srcChainsCfg := make(map[ccipocr3.ChainSelector]reader3.StaticSourceChainConfig, len(slicesWithoutDst))
			for _, chain := range slicesWithoutDst {
				srcChainsCfg[chain] = reader3.StaticSourceChainConfig{
					IsEnabled: true,
				}
			}
			ccipReader.EXPECT().GetOffRampSourceChainsConfig(mock.Anything, slicesWithoutDst).
				Return(srcChainsCfg, nil).Maybe()

			ccipReader.EXPECT().GetChainsFeeComponents(mock.Anything, slicesWithoutDst).
				Return(tc.chainFeeComponents).Maybe()

			ccipReader.EXPECT().GetWrappedNativeTokenPriceUSD(mock.Anything, slicesWithoutDst).
				Return(tc.nativeTokenPrices).Maybe()

			ccipReader.EXPECT().GetChainFeePriceUpdate(mock.Anything, slicesWithoutDst).
				Return(tc.existingChainFeePriceUpdates).Maybe()

			homeChain.EXPECT().GetFChain().Return(tc.fChain, nil).Maybe()

			obs, err := p.Observation(ctx, Outcome{}, Query{})
			require.NoError(t, err)
			if tc.expUniqueChains == 0 {
				require.Equal(t, tc.fChain, obs.FChain)
				require.NotEqual(t, time.Time{}, obs.TimestampNow)
				return
			}

			require.True(t, tc.expUniqueChains == len(obs.FeeComponents))
			require.True(t, tc.expUniqueChains == len(obs.NativeTokenPrices))
			require.ElementsMatch(t,
				slices.Collect(maps.Keys(obs.FeeComponents)),
				slices.Collect(maps.Keys(obs.NativeTokenPrices)),
			)
		})
	}
}

func Test_processor_Observation_PreventsOverlappingOps(t *testing.T) {
	cs := plugincommon.NewMockChainSupport(t)
	ccipReader := reader.NewMockCCIPReader(t)
	homeChain := reader2.NewMockHomeChain(t)
	oracleID := commontypes.OracleID(rand.Int() % 255)
	lggr := logger.Test(t)
	ctx := t.Context()

	// 1. Setup Processor
	p := &processor{
		lggr:            lggr,
		chainSupport:    cs,
		destChain:       1,
		ccipReader:      ccipReader,
		oracleID:        oracleID,
		homeChain:       homeChain,
		metricsReporter: plugincommon2.NoopReporter{},
		obs:             newBaseObserver(ccipReader, 1, oracleID, cs),
		// Set a short timeout so the first Observation call returns quickly even if op hangs
		cfg: pluginconfig.CommitOffchainConfig{
			ChainFeeAsyncObserverSyncTimeout: 100 * time.Millisecond,
		},
	}

	// 2. Setup Mocks
	supportedSet := mapset.NewSet(ccipocr3.ChainSelector(2))
	cs.EXPECT().DestChain().Return(ccipocr3.ChainSelector(1)).Maybe()
	cs.EXPECT().SupportedChains(oracleID).Return(supportedSet, nil).Maybe()
	cs.EXPECT().KnownSourceChainsSlice().Return([]ccipocr3.ChainSelector{2}, nil).Maybe()
	cs.EXPECT().SupportsDestChain(oracleID).Return(true, nil).Maybe()

	srcChainsCfg := map[ccipocr3.ChainSelector]reader3.StaticSourceChainConfig{
		2: {IsEnabled: true},
	}
	ccipReader.EXPECT().GetOffRampSourceChainsConfig(mock.Anything, mock.Anything).Return(srcChainsCfg, nil).Maybe()

	// Mocks for other operations (return immediately/empty)
	ccipReader.EXPECT().GetWrappedNativeTokenPriceUSD(mock.Anything, mock.Anything).Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{}).Maybe()
	ccipReader.EXPECT().GetChainFeePriceUpdate(mock.Anything, mock.Anything).Return(map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig{}).Maybe()
	homeChain.EXPECT().GetFChain().Return(map[ccipocr3.ChainSelector]int{}, nil).Maybe()

	// 3. Setup the HANGING Mock
	// This should only be called ONCE. If overlap protection fails, it might be called twice.
	opStarted := make(chan struct{})
	ccipReader.EXPECT().GetChainsFeeComponents(mock.Anything, mock.Anything).
		Run(func(_ context.Context, _ []ccipocr3.ChainSelector) {
			close(opStarted)
			// Simulate hang
			time.Sleep(1 * time.Hour)
		}).
		Return(map[ccipocr3.ChainSelector]types.ChainFeeComponents{}).
		Once()

	// 4. Run First Observation (will hang then timeout)
	go func() {
		_, _ = p.Observation(ctx, Outcome{}, Query{})
	}()

	// Wait for the operation to start
	select {
	case <-opStarted:
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for op to start")
	}

	// Wait for Observation to timeout and return (approx 100ms)
	time.Sleep(200 * time.Millisecond)

	// 5. Run Second Observation
	// This should SKIP GetChainsFeeComponents because it's still running.
	// If it doesn't skip, the mock will fail (because .Once() was specified) or it will hang again.
	obs, err := p.Observation(ctx, Outcome{}, Query{})
	require.NoError(t, err)

	// Verify that we got a result (even if empty fee components) and didn't block
	require.NotNil(t, obs)
}
