package tokenprice

import (
	"context"
	"errors"
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	common_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
		tokenA: cciptypes.NewTimestampedBig(bi100.Int64(), timestamp),
		tokenB: cciptypes.NewTimestampedBig(bi200.Int64(), timestamp),
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
						tokenA: cciptypes.NewTimestampedBig(bi100.Int64(), timestamp),
						tokenB: cciptypes.NewTimestampedBig(bi200.Int64(), timestamp),
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
				homeChain := readermock.NewMockHomeChain(t)
				homeChain.EXPECT().GetFChain().Return(nil, errors.New("failed to get FChain"))

				chainSupport := common_mock.NewMockChainSupport(t)
				tokenPriceReader := readerpkg_mock.NewMockPriceReader(t)

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
			expObs: Observation{},
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
	TokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
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
