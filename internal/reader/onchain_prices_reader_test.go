package reader

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	chainreadermocks "github.com/smartcontractkit/chainlink-ccip/mocks/cl-common/chainreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	EthAddr = ocr2types.Account("0x2e03388D351BF87CF2409EFf18C45Df59775Fbb2")
	OpAddr  = ocr2types.Account("0x3e03388D351BF87CF2409EFf18C45Df59775Fbb2")
	ArbAddr = ocr2types.Account("0x4e03388D351BF87CF2409EFf18C45Df59775Fbb2")
)

var (
	EthPrice   = big.NewInt(1).Mul(big.NewInt(7), big.NewInt(1e18))
	OpPrice    = big.NewInt(1).Mul(big.NewInt(6), big.NewInt(1e18))
	ArbPrice   = big.NewInt(1).Mul(big.NewInt(5), big.NewInt(1e18))
	OnlyPrice  = big.NewInt(1).Mul(big.NewInt(5), big.NewInt(1e18))
	Decimals18 = uint8(18)
)

func TestOnchainTokenPricesReader_GetTokenPricesUSD(t *testing.T) {
	testCases := []struct {
		name           string
		inputTokens    []ocr2types.Account
		priceSources   map[ocr2types.Account]pluginconfig.ArbitrumPriceSource
		tokenDecimals  map[ocr2types.Account]uint8
		want           []*big.Int
		getChainReader func(t *testing.T) *chainreadermocks.MockChainReader
		wantErr        bool
	}{
		{
			name: "On-chain one price",
			// No need to put sources as we're mocking the reader
			priceSources: map[ocr2types.Account]pluginconfig.ArbitrumPriceSource{},
			tokenDecimals: map[ocr2types.Account]uint8{
				ArbAddr: Decimals18,
				OpAddr:  Decimals18,
				EthAddr: Decimals18,
			},
			inputTokens: []ocr2types.Account{ArbAddr},
			want:        []*big.Int{ArbPrice},
			getChainReader: func(t *testing.T) *chainreadermocks.MockChainReader {
				chainReader := chainreadermocks.NewMockChainReader(t)
				// expect a single decimals() call.
				chainReader.
					EXPECT().
					GetLatestValue(
						mock.Anything,
						consts.ContractNamePriceAggregator,
						consts.MethodNameGetDecimals,
						primitives.Unconfirmed,
						mock.Anything,
						mock.Anything).
					Run(func(
						ctx context.Context,
						contractName,
						method string,
						confidenceLevel primitives.ConfidenceLevel,
						params,
						returnVal interface{}) {
						returnValUint8, ok := returnVal.(*uint8)
						if !ok {
							panic("returnVal is not a *uint8")
						}
						*returnValUint8 = Decimals18
					}).
					Return(nil)
				// expect a single getLatestRoundData() call.
				chainReader.
					EXPECT().
					GetLatestValue(
						mock.Anything,
						consts.ContractNamePriceAggregator,
						consts.MethodNameGetLatestRoundData,
						primitives.Unconfirmed,
						mock.Anything,
						mock.Anything).
					Run(func(
						ctx context.Context,
						contractName,
						method string,
						confidenceLevel primitives.ConfidenceLevel,
						params,
						returnVal interface{}) {
						returnValLatestRoundData := returnVal.(*LatestRoundData)
						if returnValLatestRoundData == nil {
							panic("returnVal is nil")
						}
						returnValLatestRoundData.Answer = big.NewInt(ArbPrice.Int64())
					}).Return(nil).Once()
				return chainReader
			},
		},
		{
			name:         "Missing price should error",
			priceSources: map[ocr2types.Account]pluginconfig.ArbitrumPriceSource{},
			inputTokens:  []ocr2types.Account{ArbAddr},
			getChainReader: func(t *testing.T) *chainreadermocks.MockChainReader {
				chainReader := chainreadermocks.NewMockChainReader(t)
				// expect a single getLatestRoundData() call that will error
				chainReader.
					EXPECT().
					GetLatestValue(
						mock.Anything,
						consts.ContractNamePriceAggregator,
						consts.MethodNameGetLatestRoundData,
						primitives.Unconfirmed,
						mock.Anything,
						mock.Anything).
					Return(fmt.Errorf("some error")).Once()
				// no decimals() call since the above call errors.
				return chainReader
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contractReader := tc.getChainReader(t)

			tokenPricesReader := OnchainTokenPricesReader{
				ContractReader: contractReader,
				PriceSources:   tc.priceSources,
				TokenDecimals:  tc.tokenDecimals,
			}
			ctx := tests.Context(t)
			result, err := tokenPricesReader.GetTokenPricesUSD(ctx, tc.inputTokens)

			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.want, result)
		})
	}

}

func TestPriceService_calculateUsdPer1e18TokenAmount(t *testing.T) {
	testCases := []struct {
		name       string
		price      *big.Int
		decimal    uint8
		wantResult *big.Int
	}{
		{
			name:       "18-decimal token, $6.5 per token",
			price:      big.NewInt(65e17),
			decimal:    18,
			wantResult: big.NewInt(65e17),
		},
		{
			name:       "6-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    6,
			wantResult: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e12)), // 1e30
		},
		{
			name:       "0-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    0,
			wantResult: new(big.Int).Mul(big.NewInt(1e18), big.NewInt(1e18)), // 1e36
		},
		{
			name:       "36-decimal token, $1 per token",
			price:      big.NewInt(1e18),
			decimal:    36,
			wantResult: big.NewInt(1),
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := calculateUsdPer1e18TokenAmount(tt.price, tt.decimal)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}
