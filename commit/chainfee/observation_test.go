package chainfee

import (
	"math/big"
	"math/rand"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
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
		existingChainFeePriceUpdates map[ccipocr3.ChainSelector]plugintypes.TimestampedBig
		fChain                       map[ccipocr3.ChainSelector]int
		expectedChainFeePriceUpdates map[ccipocr3.ChainSelector]Update

		expErr bool
	}{
		{
			name:            "two chains",
			supportedChains: []ccipocr3.ChainSelector{1},
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
			existingChainFeePriceUpdates: map[ccipocr3.ChainSelector]plugintypes.TimestampedBig{
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
			},
			expErr: false,
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
				ccipReader:      ccipReader,
				oracleID:        oracleID,
				homeChain:       homeChain,
				metricsReporter: NoopMetrics{},
			}

			cs.EXPECT().SupportedChains(oracleID).
				Return(mapset.NewSet(tc.supportedChains...), nil)

			ccipReader.EXPECT().GetChainsFeeComponents(ctx, tc.supportedChains).
				Return(tc.chainFeeComponents)

			ccipReader.EXPECT().GetWrappedNativeTokenPriceUSD(ctx, tc.supportedChains).
				Return(tc.nativeTokenPrices)

			ccipReader.EXPECT().GetChainFeePriceUpdate(ctx, tc.supportedChains).
				Return(tc.existingChainFeePriceUpdates)

			homeChain.EXPECT().GetFChain().Return(tc.fChain, nil)

			tStart := time.Now()
			obs, err := p.Observation(ctx, Outcome{}, Query{})
			tEnd := time.Now()
			if tc.expErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.GreaterOrEqual(t, obs.TimestampNow.UnixNano(), tStart.UnixNano())
			require.LessOrEqual(t, obs.TimestampNow.UnixNano(), tEnd.UnixNano())
			require.Equal(t, tc.chainFeeComponents, obs.FeeComponents)
			require.Equal(t, tc.nativeTokenPrices, obs.NativeTokenPrices)
			require.Equal(t, tc.expectedChainFeePriceUpdates, obs.ChainFeeUpdates)
			require.Equal(t, tc.fChain, obs.FChain)
		})
	}
}
