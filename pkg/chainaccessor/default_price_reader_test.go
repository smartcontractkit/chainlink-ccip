package chainaccessor

import (
	"context"
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

const (
	ArbAddr           = cciptypes.UnknownEncodedAddress("0xa100000000000000000000000000000000000000")
	ArbAggregatorAddr = cciptypes.UnknownEncodedAddress("0xa2000000000000000000000000000000000000000")

	EthAddr           = cciptypes.UnknownEncodedAddress("0xe100000000000000000000000000000000000000")
	EthAggregatorAddr = cciptypes.UnknownEncodedAddress("0xe200000000000000000000000000000000000000")

	BtcAddr          = cciptypes.UnknownEncodedAddress("0xb100000000000000000000000000000000000000")
	BtcAgregatorAddr = cciptypes.UnknownEncodedAddress("0xb200000000000000000000000000000000000000")
)

var (
	EthPrice   = big.NewInt(1).Mul(big.NewInt(7), big.NewInt(1e18))
	ArbPrice   = big.NewInt(1).Mul(big.NewInt(5), big.NewInt(1e18))
	Decimals18 = uint8(18)

	ArbInfo = cciptypes.TokenInfo{
		AggregatorAddress: ArbAggregatorAddr,
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
	EthInfo = cciptypes.TokenInfo{
		AggregatorAddress: EthAggregatorAddr,
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
	BtcInfo = cciptypes.TokenInfo{
		AggregatorAddress: BtcAgregatorAddr,
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
)

func TestDefaultAccessor_GetFeedPricesUSD(t *testing.T) {
	testCases := []struct {
		name          string
		inputTokens   []cciptypes.UnknownEncodedAddress
		tokenInfo     map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo
		mockPrices    map[cciptypes.UnknownEncodedAddress]*big.Int
		want          cciptypes.TokenPriceMap
		errorAccounts []cciptypes.UnknownEncodedAddress
		wantErr       bool
	}{
		{
			name: "On-chain one price",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
			want:        cciptypes.TokenPriceMap{ArbAddr: cciptypes.NewBigInt(ArbPrice)},
		},
		{
			name: "On-chain multiple prices",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr, EthAddr},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice, EthAddr: EthPrice},
			want:        cciptypes.TokenPriceMap{ArbAddr: cciptypes.NewBigInt(ArbPrice), EthAddr: cciptypes.NewBigInt(EthPrice)},
		},
		{
			name: "Missing price doesn't fail, return available prices",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
			},
			inputTokens:   []cciptypes.UnknownEncodedAddress{ArbAddr, EthAddr},
			mockPrices:    map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
			errorAccounts: []cciptypes.UnknownEncodedAddress{EthAddr},
			want:          cciptypes.TokenPriceMap{ArbAddr: cciptypes.NewBigInt(ArbPrice)},
		},
		{
			name: "Empty input tokens list",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{},
			want:        cciptypes.TokenPriceMap{},
		},
		{
			name: "Repeated token in input",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr, ArbAddr},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
			want:        cciptypes.TokenPriceMap{ArbAddr: cciptypes.NewBigInt(ArbPrice)},
		},
		{
			name: "Zero price should be discarded",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: big.NewInt(0)},
			want:        cciptypes.TokenPriceMap{},
		},
		{
			name: "Multiple error accounts",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
				BtcAddr: BtcInfo,
			},
			inputTokens:   []cciptypes.UnknownEncodedAddress{ArbAddr, EthAddr, BtcAddr},
			mockPrices:    map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
			errorAccounts: []cciptypes.UnknownEncodedAddress{EthAddr, BtcAddr},
			want:          cciptypes.TokenPriceMap{ArbAddr: cciptypes.NewBigInt(ArbPrice)},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			contractReader := createMockReader(t, tc.mockPrices, tc.errorAccounts, tc.tokenInfo)
			accessor := &DefaultAccessor{
				lggr:           logger.Test(t),
				contractReader: contractReader,
			}

			ctx := context.Background()
			result, err := accessor.GetFeedPricesUSD(ctx, tc.inputTokens, tc.tokenInfo)

			if tc.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.want, result)
		})
	}
}

func Test_calculateUsdPer1e18TokenAmount(t *testing.T) {
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
	mockPrices map[cciptypes.UnknownEncodedAddress]*big.Int,
	errorAccounts []cciptypes.UnknownEncodedAddress,
	tokenInfo map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo,
) *readermock.MockExtended {
	reader := readermock.NewMockExtended(t)

	// Create the expected batch request and results
	expectedRequest := make(commontypes.BatchGetLatestValuesRequest)
	expectedResults := make(commontypes.BatchGetLatestValuesResult)

	// Handle successful cases
	for token, price := range mockPrices {
		info := tokenInfo[token]
		boundContract := commontypes.BoundContract{
			Address: string(info.AggregatorAddress),
			Name:    consts.ContractNamePriceAggregator,
		}

		// Add to expected request
		if _, exists := expectedRequest[boundContract]; !exists {
			expectedRequest[boundContract] = make(commontypes.ContractBatch, 0)
		}
		expectedRequest[boundContract] = append(expectedRequest[boundContract],
			commontypes.BatchRead{
				ReadName:  consts.MethodNameGetLatestRoundData,
				Params:    nil,
				ReturnVal: &LatestRoundData{},
			},
			commontypes.BatchRead{
				ReadName:  consts.MethodNameGetDecimals,
				Params:    nil,
				ReturnVal: new(uint8),
			},
		)

		// Create results
		results := make(commontypes.ContractBatchResults, 2)
		// Price result
		priceResult := commontypes.BatchReadResult{ReadName: consts.MethodNameGetLatestRoundData}
		priceResult.SetResult(&LatestRoundData{Answer: new(big.Int).Set(price)}, nil)
		results[0] = priceResult

		// Decimals result
		decimalsResult := commontypes.BatchReadResult{ReadName: consts.MethodNameGetDecimals}
		decimals := info.Decimals
		decimalsResult.SetResult(&decimals, nil)
		results[1] = decimalsResult

		expectedResults[boundContract] = results
	}

	// Handle error cases
	for _, account := range errorAccounts {
		info := tokenInfo[account]
		boundContract := commontypes.BoundContract{
			Address: string(info.AggregatorAddress),
			Name:    consts.ContractNamePriceAggregator,
		}

		results := make(commontypes.ContractBatchResults, 2)
		// Price result with error
		priceResult := commontypes.BatchReadResult{ReadName: consts.MethodNameGetLatestRoundData}
		priceResult.SetResult(nil, fmt.Errorf("error"))
		results[0] = priceResult

		// Decimals result
		decimalsResult := commontypes.BatchReadResult{ReadName: consts.MethodNameGetDecimals}
		decimalsResult.SetResult(nil, nil)
		results[1] = decimalsResult

		expectedResults[boundContract] = results
	}

	// Set up the mock expectation for BatchGetLatestValues
	reader.On("BatchGetLatestValues",
		mock.Anything,
		mock.MatchedBy(func(req commontypes.BatchGetLatestValuesRequest) bool {
			// Validate request structure
			for boundContract, batch := range req {
				// Verify contract has exactly two reads (price and decimals)
				if len(batch) != 2 {
					return false
				}
				// Verify read names
				if batch[0].ReadName != consts.MethodNameGetLatestRoundData ||
					batch[1].ReadName != consts.MethodNameGetDecimals {
					return false
				}
				// Verify contract exists in our expected results
				if _, exists := expectedResults[boundContract]; !exists {
					return false
				}
			}
			return true
		}),
	).Return(expectedResults, nil).Once()

	return reader
}
