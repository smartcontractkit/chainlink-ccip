package reader

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	ArbAddr           = ocr2types.Account("0xa100000000000000000000000000000000000000")
	ArbAggregatorAddr = ocr2types.Account("0xa2000000000000000000000000000000000000000")

	EthAddr           = ocr2types.Account("0xe100000000000000000000000000000000000000")
	EthAggregatorAddr = ocr2types.Account("0xe200000000000000000000000000000000000000")
)

var (
	EthPrice   = big.NewInt(1).Mul(big.NewInt(7), big.NewInt(1e18))
	ArbPrice   = big.NewInt(1).Mul(big.NewInt(5), big.NewInt(1e18))
	Decimals18 = uint8(18)

	ArbInfo = pluginconfig.TokenInfo{
		AggregatorAddress: string(ArbAggregatorAddr),
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
	EthInfo = pluginconfig.TokenInfo{
		AggregatorAddress: string(EthAggregatorAddr),
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
)

func TestOnchainTokenPricesReader_GetTokenPricesUSD(t *testing.T) {
	testCases := []struct {
		name          string
		inputTokens   []ocr2types.Account
		tokenInfo     map[ocr2types.Account]pluginconfig.TokenInfo
		mockPrices    map[ocr2types.Account]*big.Int
		want          []*big.Int
		errorAccounts []ocr2types.Account
		wantErr       bool
	}{
		{
			name: "On-chain one price",
			tokenInfo: map[ocr2types.Account]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []ocr2types.Account{ArbAddr},
			mockPrices:  map[ocr2types.Account]*big.Int{ArbAddr: ArbPrice},
			want:        []*big.Int{ArbPrice},
		},
		{
			name: "On-chain multiple prices",
			tokenInfo: map[ocr2types.Account]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
			},
			inputTokens: []ocr2types.Account{ArbAddr, EthAddr},
			mockPrices:  map[ocr2types.Account]*big.Int{ArbAddr: ArbPrice, EthAddr: EthPrice},
			want:        []*big.Int{ArbPrice, EthPrice},
		},
		{
			name: "Missing price should error",
			tokenInfo: map[ocr2types.Account]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
			},
			inputTokens:   []ocr2types.Account{ArbAddr, EthAddr},
			mockPrices:    map[ocr2types.Account]*big.Int{ArbAddr: ArbPrice},
			errorAccounts: []ocr2types.Account{EthAddr},
			want:          nil,
			wantErr:       true,
		},
	}

	for _, tc := range testCases {
		contractReader := createMockReader(t, tc.mockPrices, tc.errorAccounts, tc.tokenInfo)
		feedChain := cciptypes.ChainSelector(1)
		tokenPricesReader := priceReader{
			chainReaders: map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				feedChain: contractReader,
			},
			tokenInfo: tc.tokenInfo,
			feedChain: feedChain,
		}
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()
			result, err := tokenPricesReader.GetFeedPricesUSD(ctx, tc.inputTokens)

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
	t *testing.T,
	mockPrices map[ocr2types.Account]*big.Int,
	errorAccounts []ocr2types.Account,
	tokenInfo map[ocr2types.Account]pluginconfig.TokenInfo,
) *readermock.MockContractReaderFacade {
	reader := readermock.NewMockContractReaderFacade(t)

	for token, price := range mockPrices {
		info := tokenInfo[token]
		boundContract := commontypes.BoundContract{
			Address: info.AggregatorAddress,
			Name:    consts.ContractNamePriceAggregator,
		}

		identifier := boundContract.ReadIdentifier(consts.MethodNameGetLatestRoundData)
		reader.On("GetLatestValue",
			mock.Anything,
			identifier,
			primitives.Unconfirmed,
			nil,
			mock.Anything).Run(
			func(args mock.Arguments) {
				arg := args.Get(4).(*LatestRoundData)
				arg.Answer = big.NewInt(price.Int64())
			}).Return(nil).Once()

		reader.On("GetLatestValue",
			mock.Anything,
			boundContract.ReadIdentifier(consts.MethodNameGetDecimals),
			primitives.Unconfirmed,
			nil,
			mock.Anything).Run(
			func(args mock.Arguments) {
				arg := args.Get(4).(*uint8)
				*arg = info.Decimals
			}).Return(nil)
	}

	for _, account := range errorAccounts {
		info := tokenInfo[account]
		boundContract := commontypes.BoundContract{
			Address: info.AggregatorAddress,
			Name:    consts.ContractNamePriceAggregator,
		}
		reader.On("GetLatestValue",
			mock.Anything,
			boundContract.ReadIdentifier(consts.MethodNameGetLatestRoundData),
			primitives.Unconfirmed,
			nil,
			mock.Anything).Return(fmt.Errorf("error")).Once()
	}

	return reader
}
