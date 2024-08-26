package commitrmnocb

import (
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/assert"
)

func Test_SelectTokens(t *testing.T) {
	lggr := logger.Test(t)

	TokenA := types.Account("tokenA")
	TokenB := types.Account("tokenB")

	defaultFeedPrices := map[types.Account]cciptypes.BigInt{
		TokenA: BigInt(10),
		TokenB: BigInt(2),
	}

	defaultRegistryPrices := map[types.Account]cciptypes.BigInt{
		TokenA: BigInt(5), // Deviating
		TokenB: BigInt(2),
	}

	defaultTokenInfo := map[types.Account]pluginconfig.TokenInfo{
		TokenA: {DeviationPPB: BigInt(1)},
		TokenB: {DeviationPPB: BigInt(1)},
	}

	tests := []struct {
		name               string
		consensusTimestamp time.Time
		lastPricesUpdate   time.Time
		updateFrequency    commonconfig.Duration
		feedPrices         map[types.Account]cciptypes.BigInt
		registryPrices     map[types.Account]cciptypes.BigInt
		tokenInfo          map[types.Account]pluginconfig.TokenInfo
		expectedOutput     map[types.Account]cciptypes.BigInt
	}{
		{
			name:               "Update all tokens when time passed is greater than update frequency",
			consensusTimestamp: time.Now(),
			lastPricesUpdate:   time.Now().Add(-10 * time.Minute),
			updateFrequency:    *commonconfig.MustNewDuration(5 * time.Minute),
			feedPrices:         defaultFeedPrices,
			registryPrices:     defaultRegistryPrices,
			tokenInfo:          defaultTokenInfo,
			expectedOutput:     defaultFeedPrices,
		},
		{
			name:               "Update only deviated tokens",
			consensusTimestamp: time.Now(),
			lastPricesUpdate:   time.Now().Add(-2 * time.Minute),
			updateFrequency:    *commonconfig.MustNewDuration(5 * time.Minute),
			feedPrices:         defaultFeedPrices,
			registryPrices:     defaultRegistryPrices,
			tokenInfo:          defaultTokenInfo,
			expectedOutput: map[types.Account]cciptypes.BigInt{
				TokenA: defaultFeedPrices[TokenA], // tokenA deviates, tokenB does not
			},
		},
		{
			name:               "No updates when no deviation and time passed is less than update frequency",
			consensusTimestamp: time.Now(),
			lastPricesUpdate:   time.Now().Add(-2 * time.Minute),
			updateFrequency:    *commonconfig.MustNewDuration(5 * time.Minute),
			feedPrices:         defaultFeedPrices,
			registryPrices: map[types.Account]cciptypes.BigInt{
				TokenA: defaultFeedPrices[TokenA],
				TokenB: defaultFeedPrices[TokenB],
			},
			tokenInfo:      defaultTokenInfo,
			expectedOutput: map[types.Account]cciptypes.BigInt{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := selectTokens(lggr, tt.feedPrices, tt.registryPrices, tt.consensusTimestamp, tt.lastPricesUpdate, tt.updateFrequency, tt.tokenInfo)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
func TestDeviates(t *testing.T) {
	type args struct {
		x1  *big.Int
		x2  *big.Int
		ppb int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "base case",
			args: args{x1: big.NewInt(1e9), x2: big.NewInt(2e9), ppb: 1},
			want: true,
		},
		{
			name: "x1 is zero and x1 neq x2",
			args: args{x1: big.NewInt(0), x2: big.NewInt(1), ppb: 999},
			want: true,
		},
		{
			name: "x2 is zero and x1 neq x2",
			args: args{x1: big.NewInt(1), x2: big.NewInt(0), ppb: 999},
			want: true,
		},
		{
			name: "x1 and x2 are both zero",
			args: args{x1: big.NewInt(0), x2: big.NewInt(0), ppb: 999},
			want: false,
		},
		{
			name: "deviates when ppb is 0",
			args: args{x1: big.NewInt(0), x2: big.NewInt(1), ppb: 0},
			want: true,
		},
		{
			name: "does not deviate when x1 eq x2",
			args: args{x1: big.NewInt(5), x2: big.NewInt(5), ppb: 1},
			want: false,
		},
		{
			name: "does not deviate with high ppb when x2 is greater",
			args: args{x1: big.NewInt(5), x2: big.NewInt(10), ppb: 2e9},
			want: false,
		},
		{
			name: "does not deviate with high ppb when x1 is greater",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 2e9},
			want: false,
		},
		{
			name: "deviates with low ppb when x2 is greater",
			args: args{x1: big.NewInt(5), x2: big.NewInt(10), ppb: 9e8},
			want: true,
		},
		{
			name: "deviates with low ppb when x1 is greater",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 9e8},
			want: true,
		},
		{
			name: "near deviation limit but deviates",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 1e9 - 1},
			want: true,
		},
		{
			name: "near deviation limit but deviates",
			args: args{x1: big.NewInt(9e18), x2: big.NewInt(4e18), ppb: 1e9 - 1},
			want: true,
		},
		{
			name: "at deviation limit but does not deviate",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 1e9},
			want: false,
		},
		{
			name: "near deviation limit but does not deviate",
			args: args{x1: big.NewInt(10), x2: big.NewInt(5), ppb: 1e9 + 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Deviates(tt.args.x1, tt.args.x2, tt.args.ppb), "Deviates(%v, %v, %v)", tt.args.x1, tt.args.x2, tt.args.ppb)
		})
	}
}

func BigInt(i int64) cciptypes.BigInt {
	return cciptypes.BigInt{Int: big.NewInt(i)}
}
