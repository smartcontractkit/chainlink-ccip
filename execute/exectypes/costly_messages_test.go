package exectypes

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestWaitBoostedFee(t *testing.T) {
	tests := []struct {
		name                     string
		sendTimeDiff             time.Duration
		fee                      *big.Int
		diff                     *big.Int
		relativeBoostPerWaitHour float64
	}{
		{
			"wait 10s",
			time.Second * 10,
			big.NewInt(6e18), // Fee:   6    LINK

			big.NewInt(1166666666665984), // Boost: 0.01 LINK
			0.07,
		},
		{
			"wait 5m",
			time.Minute * 5,
			big.NewInt(6e18),  // Fee:   6    LINK
			big.NewInt(35e15), // Boost: 0.35 LINK
			0.07,
		},
		{
			"wait 7m",
			time.Minute * 7,
			big.NewInt(6e18),  // Fee:   6    LINK
			big.NewInt(49e15), // Boost: 0.49 LINK
			0.07,
		},
		{
			"wait 12m",
			time.Minute * 12,
			big.NewInt(6e18),  // Fee:   6    LINK
			big.NewInt(84e15), // Boost: 0.84 LINK
			0.07,
		},
		{
			"wait 25m",
			time.Minute * 25,
			big.NewInt(6e18),               // Fee:   6 LINK
			big.NewInt(174999999999998976), // Boost: 1.75 LINK
			0.07,
		},
		{
			"wait 1h",
			time.Hour * 1,
			big.NewInt(6e18),   // Fee:   6 LINK
			big.NewInt(420e15), // Boost: 4.2 LINK
			0.07,
		},
		{
			"wait 5h",
			time.Hour * 5,
			big.NewInt(6e18),                // Fee:   6 LINK
			big.NewInt(2100000000000001024), // Boost: 21LINK
			0.07,
		},
		{
			"wait 24h",
			time.Hour * 24,
			big.NewInt(6e18), // Fee:   6 LINK
			big.NewInt(0).Mul(big.NewInt(10), big.NewInt(1008e15)), // Boost: 100LINK
			0.07,
		},
		{
			"high boost wait 10s",
			time.Second * 10,
			big.NewInt(5e18),
			big.NewInt(9722222222222336), // 1e16
			0.7,
		},
		{
			"high boost wait 5m",
			time.Minute * 5,
			big.NewInt(5e18),
			big.NewInt(291666666666667008), // 1e18
			0.7,
		},
		{
			"high boost wait 25m",
			time.Minute * 25,
			big.NewInt(5e18),
			big.NewInt(1458333333333334016), // 1e19
			0.7,
		},
		{
			"high boost wait 5h",
			time.Hour * 5,
			big.NewInt(5e18),
			big.NewInt(0).Mul(big.NewInt(10), big.NewInt(175e16)), // 1e20
			0.7,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			boosted := waitBoostedFee(tc.sendTimeDiff, tc.fee, tc.relativeBoostPerWaitHour)
			diff := big.NewInt(0).Sub(boosted, tc.fee)
			assert.Equal(t, diff, tc.diff)
		})
	}
}

func TestCcipMessageFeeE18USDCalculator_MessageFeeE18USD(t *testing.T) {
	tests := []struct {
		name                     string
		messages                 []cciptypes.Message
		messageTimeStamps        map[cciptypes.Bytes32]time.Time
		linkPrice                cciptypes.BigInt
		relativeBoostPerWaitHour float64
		want                     map[cciptypes.Bytes32]cciptypes.BigInt
		wantErr                  assert.ErrorAssertionFunc
	}{
		{
			name:                     "happy path",
			messages:                 []cciptypes.Message{},
			messageTimeStamps:        map[cciptypes.Bytes32]time.Time{},
			linkPrice:                cciptypes.NewBigIntFromInt64(100),
			relativeBoostPerWaitHour: 0.5,
			want:                     map[cciptypes.Bytes32]cciptypes.BigInt{},
			wantErr:                  assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			lggr := logger.Test(t)

			mockReader := readerpkg_mock.NewMockCCIPReader(t)
			mockReader.EXPECT().LinkPriceUSD(ctx).Return(tt.linkPrice, nil)

			calculator := &CcipMessageFeeE18USDCalculator{
				lggr:                     lggr,
				ccipReader:               mockReader,
				RelativeBoostPerWaitHour: tt.relativeBoostPerWaitHour,
			}

			got, err := calculator.MessageFeeE18USD(ctx, tt.messages, tt.messageTimeStamps)
			if !tt.wantErr(t, err, "getMessageTimestampMap(...)") {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
