package tokenprice

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	common_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_Observation(t *testing.T) {
	fChains := map[cciptypes.ChainSelector]int{
		feedChainSel: f,
		destChainSel: f,
	}
	timestamp := time.Now().UTC()
	feedTokenPrices := cciptypes.TokenPriceMap{
		tokenA: cciptypes.NewBigInt(bi100),
		tokenB: cciptypes.NewBigInt(bi200),
	}
	feeQuoterTokenUpdates := map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
		tokenA: cciptypes.NewTimestampedBig(100, timestamp), // bi100 is big.NewInt(100)
		tokenB: cciptypes.NewTimestampedBig(200, timestamp), // bi200 is big.NewInt(200)
	}
	oracleID := commontypes.OracleID(1)
	lggr := logger.Test(t)

	testCases := []struct {
		name         string
		getProcessor func(t *testing.T) plugincommon.PluginProcessor[Query, Observation, Outcome]
		expObs       Observation
		expErr       error
	}{
		{
			name: "Successful observation",
			getProcessor: func(t *testing.T) plugincommon.PluginProcessor[Query, Observation, Outcome] {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportedChains(mock.Anything).Return(
					mapset.NewSet(feedChainSel, destChainSel), nil,
				)
				chainSupport.EXPECT().SupportsDestChain(mock.Anything).Return(true, nil).Maybe()

				tokenPriceReader := readerpkg_mock.NewMockPriceReader(t)
				tokenPriceReader.EXPECT().GetFeedPricesUSD(mock.Anything, mock.MatchedBy(
					func(tokens []cciptypes.UnknownEncodedAddress) bool {
						expectedTokens := mapset.NewSet(tokenA, tokenB)
						actualTokens := mapset.NewSet(tokens...)
						return expectedTokens.Equal(actualTokens)
					})).
					Return(cciptypes.TokenPriceMap{
						tokenA: cciptypes.NewBigInt(bi100),
						tokenB: cciptypes.NewBigInt(bi200)}, nil)

				tokenPriceReader.EXPECT().GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, mock.Anything).Return(
					map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
						tokenA: cciptypes.NewTimestampedBig(100, timestamp),
						tokenB: cciptypes.NewTimestampedBig(200, timestamp),
					},
					nil,
				)

				homeChain := readermock.NewMockHomeChain(t)
				homeChain.EXPECT().GetFChain().Return(
					map[cciptypes.ChainSelector]int{destChainSel: f, feedChainSel: f},
					nil,
				)

				return NewProcessor(
					oracleID,
					lggr,
					defaultCfg,
					destChainSel,
					chainSupport,
					tokenPriceReader,
					homeChain,
					f,
					plugincommon.NoopReporter{},
				)
			},
			expObs: Observation{
				FeedTokenPrices:       feedTokenPrices,
				FeeQuoterTokenUpdates: feeQuoterTokenUpdates,
				FChain:                fChains,
				Timestamp:             time.Now().UTC(),
			},
		},
		{
			name: "Failed to get FDestChain",
			getProcessor: func(t *testing.T) plugincommon.PluginProcessor[Query, Observation, Outcome] {
				chainSupport := common_mock.NewMockChainSupport(t)
				chainSupport.EXPECT().SupportedChains(mock.Anything).Return(
					mapset.NewSet(feedChainSel, destChainSel), nil,
				)
				chainSupport.EXPECT().SupportsDestChain(mock.Anything).Return(true, nil).Maybe()

				tokenPriceReader := readerpkg_mock.NewMockPriceReader(t)
				tokenPriceReader.EXPECT().GetFeedPricesUSD(mock.Anything, mock.MatchedBy(
					func(tokens []cciptypes.UnknownEncodedAddress) bool {
						expectedTokens := mapset.NewSet(tokenA, tokenB)
						actualTokens := mapset.NewSet(tokens...)
						return expectedTokens.Equal(actualTokens)
					})).
					Return(cciptypes.TokenPriceMap{
						tokenA: cciptypes.NewBigInt(bi100),
						tokenB: cciptypes.NewBigInt(bi200)}, nil)

				tokenPriceReader.EXPECT().GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, mock.Anything).Return(
					map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{
						tokenA: cciptypes.NewTimestampedBig(100, timestamp),
						tokenB: cciptypes.NewTimestampedBig(200, timestamp),
					},
					nil,
				)

				homeChain := readermock.NewMockHomeChain(t)
				homeChain.EXPECT().GetFChain().Return(nil,
					fmt.Errorf("some unexpected error getting fChain"), // <----------------
				)

				return NewProcessor(
					oracleID,
					lggr,
					defaultCfg,
					destChainSel,
					chainSupport,
					tokenPriceReader,
					homeChain,
					f,
					plugincommon.NoopReporter{},
				)
			},
			expObs: Observation{
				FeedTokenPrices:       feedTokenPrices,
				FeeQuoterTokenUpdates: feeQuoterTokenUpdates,
				FChain:                map[cciptypes.ChainSelector]int{},
				Timestamp:             time.Now().UTC(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			p := tc.getProcessor(t)

			actualObs, err := p.Observation(ctx, Outcome{}, Query{})
			if tc.expErr != nil {
				require.Error(t, err)
				assert.Equal(t, tc.expErr.Error(), err.Error())
				assert.Equal(t, Observation{}, actualObs)
			} else {
				require.NoError(t, err)
				// No need to check timestamp
				actualObs.Timestamp = tc.expObs.Timestamp
				assert.Equal(t, tc.expObs, actualObs)
			}
		})
	}
}

var defaultCfg = pluginconfig.CommitOffchainConfig{
	TokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
		tokenA: {
			Decimals:          18,
			AggregatorAddress: "0x1111111111111111111111Ff18C45Df59775Fbb2",
			DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)},
		},
		tokenB: {
			Decimals:          18,
			AggregatorAddress: "0x2222222222222222222222Ff18C45Df59775Fbb2",
			DeviationPPB:      cciptypes.BigInt{Int: big.NewInt(1)}},
	},
	PriceFeedChainSelector: feedChainSel,
	// Have this disabled for testing purposes
	TokenPriceAsyncObserverDisabled: true,
}

func Test_processor_Observation_PreventsOverlappingOps(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)
	oracleID := commontypes.OracleID(1)

	// 1. Setup Processor
	chainSupport := common_mock.NewMockChainSupport(t)
	tokenPriceReader := readerpkg_mock.NewMockPriceReader(t)
	homeChain := readermock.NewMockHomeChain(t)

	cfg := defaultCfg
	cfg.TokenPriceAsyncObserverSyncTimeout = *commonconfig.MustNewDuration(100 * time.Millisecond)

	p := NewProcessor(
		oracleID,
		lggr,
		cfg,
		destChainSel,
		chainSupport,
		tokenPriceReader,
		homeChain,
		f,
		plugincommon.NoopReporter{},
	)

	// 2. Setup Mocks
	// Common mocks for both calls
	chainSupport.EXPECT().SupportedChains(mock.Anything).Return(
		mapset.NewSet(feedChainSel, destChainSel), nil,
	).Maybe()
	chainSupport.EXPECT().SupportsDestChain(mock.Anything).Return(true, nil).Maybe()

	// Mocks for other operations (return immediately/empty)
	// We need these because Observation calls other operations too.
	// FeedTokenPrices is the one we will make hang.
	tokenPriceReader.
		EXPECT().
		GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, mock.Anything).
		Return(
			map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{}, nil,
		).Maybe()
	homeChain.EXPECT().GetFChain().Return(
		map[cciptypes.ChainSelector]int{}, nil,
	).Maybe()

	// 3. Setup the HANGING Mock for FeedTokenPrices
	// This should only be called ONCE. If overlap protection fails, it might be called twice.
	opStarted := make(chan struct{})
	tokenPriceReader.EXPECT().GetFeedPricesUSD(mock.Anything, mock.Anything).
		Run(func(_ context.Context, _ []cciptypes.UnknownEncodedAddress) {
			close(opStarted)
			// Simulate hang
			time.Sleep(1 * time.Hour)
		}).
		Return(cciptypes.TokenPriceMap{}, nil).
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
	// This should SKIP GetFeedPricesUSD because it's still running.
	// If it doesn't skip, the mock will fail (because .Once() was specified)
	// or it will hang again (if we didn't mock Once).
	obs, err := p.Observation(ctx, Outcome{}, Query{})
	require.NoError(t, err)

	// Verify that we got a result (even if empty) and didn't block
	// Note: obs.FeedTokenPrices will be empty because the second call skipped fetching it.
	require.Empty(t, obs.FeedTokenPrices)
}
