package reader

import (
	"context"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

const (
	EthAcc = ocr2types.Account("ETH")
	OpAcc  = ocr2types.Account("OP")
	ArbAcc = ocr2types.Account("ARB")
)

var (
	EthPrice = big.NewInt(100)
	OpPrice  = big.NewInt(10)
	ArbPrice = big.NewInt(1)
)

func TestOnchainTokenPricesReader_GetTokenPricesUSD(t *testing.T) {
	t.Skip("Skipping until we have a full price reader implementation")
	testCases := []struct {
		name          string
		inputTokens   []ocr2types.Account
		priceSources  map[ocr2types.Account]pluginconfig.ArbitrumPriceSource
		mockPrices    map[ocr2types.Account]*big.Int
		want          []*big.Int
		errorAccounts []ocr2types.Account
		wantErr       bool
	}{
		{
			name:         "On-chain price",
			priceSources: map[ocr2types.Account]pluginconfig.ArbitrumPriceSource{},
			inputTokens:  []ocr2types.Account{ArbAcc, OpAcc, EthAcc},
			mockPrices:   map[ocr2types.Account]*big.Int{OpAcc: OpPrice, ArbAcc: ArbPrice, EthAcc: EthPrice},
			want:         []*big.Int{ArbPrice, OpPrice, EthPrice},
		},
		{
			name:          "Missing price should error",
			priceSources:  map[ocr2types.Account]pluginconfig.ArbitrumPriceSource{},
			inputTokens:   []ocr2types.Account{ArbAcc, OpAcc, EthAcc},
			mockPrices:    map[ocr2types.Account]*big.Int{OpAcc: OpPrice, ArbAcc: ArbPrice},
			errorAccounts: []ocr2types.Account{EthAcc},
			want:          nil,
			wantErr:       true,
		},
	}

	for _, tc := range testCases {
		contractReader := createMockReader(tc.mockPrices, tc.errorAccounts, tc.priceSources)
		tokenPricesReader := OnchainTokenPricesReader{
			ContractReader: contractReader,
		}
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
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

func createMockReader(
	mockPrices map[ocr2types.Account]*big.Int,
	errorAccounts []ocr2types.Account,
	priceSources map[ocr2types.Account]pluginconfig.ArbitrumPriceSource,
) *mocks.ContractReaderMock {
	reader := mocks.NewContractReaderMock()
	println(errorAccounts)
	println(priceSources)
	// TODO: Create a list of bound contracts from priceSources and return the price given in mockPrices
	for _, price := range mockPrices {
		price := price
		reader.On("GetLatestValue",
			mock.Anything,
			consts.ContractNamePriceAggregator,
			consts.MethodNameGetLatestRoundData,
			mock.Anything,
			nil,
			mock.Anything).Run(
			func(args mock.Arguments) {
				arg := args.Get(5).(*big.Int)
				arg.Set(price)
			}).Return(nil)
	}
	return reader
}
