package reader

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	typeconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	reader_mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
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

	ArbInfo = pluginconfig.TokenInfo{
		AggregatorAddress: ArbAggregatorAddr,
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
	EthInfo = pluginconfig.TokenInfo{
		AggregatorAddress: EthAggregatorAddr,
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
	BtcInfo = pluginconfig.TokenInfo{
		AggregatorAddress: BtcAgregatorAddr,
		DeviationPPB:      cciptypes.NewBigInt(big.NewInt(1e5)),
		Decimals:          Decimals18,
	}
)

func setupTestEnvironment(
	t *testing.T,
	lggr logger.Logger,
	testChain cciptypes.ChainSelector,
	feeQuoterAddr []byte,
	cr *reader_mocks.MockContractReaderFacade,
) (*ccipChainReader, PriceReader) {
	crs := make(map[cciptypes.ChainSelector]contractreader.Extended)
	crs[testChain] = contractreader.NewExtendedContractReader(cr)

	ccipReader := &ccipChainReader{
		lggr:            lggr,
		contractReaders: crs,
		contractWriters: nil,
		destChain:       testChain,
	}

	contracts := ContractAddresses{
		consts.ContractNameFeeQuoter: {
			testChain: feeQuoterAddr,
		},
	}
	require.NoError(t, ccipReader.Sync(tests.Context(t), contracts))

	tokenInfo := map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
		ArbAddr: ArbInfo,
		EthAddr: EthInfo,
		BtcAddr: BtcInfo,
	}

	pr := NewPriceReader(
		lggr,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			testChain: cr,
		},
		tokenInfo,
		ccipReader,
		testChain,
	)

	return ccipReader, pr
}

func TestPriceReader_GetFeeQuoterTokenUpdates(t *testing.T) {
	const testChain = cciptypes.ChainSelector(5)
	lggr := logger.Test(t)
	feeQuoterAddr := []byte{0x4}

	testCases := []struct {
		name        string
		inputTokens []cciptypes.UnknownEncodedAddress
		setup       func(*reader_mocks.MockContractReaderFacade)
		want        map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig
		wantErr     bool
	}{
		{
			name:        "success - single token",
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr},
			setup: func(cr *reader_mocks.MockContractReaderFacade) {
				// Mock GetLatestValue for token prices
				cr.On("GetLatestValue",
					mock.Anything,
					mock.MatchedBy(func(s string) bool {
						return strings.Contains(s, consts.MethodNameFeeQuoterGetTokenPrices)
					}),
					primitives.Unconfirmed,
					mock.MatchedBy(func(params map[string]any) bool {
						tokens, ok := params["tokens"].([][]byte)
						return ok && len(tokens) == 1
					}),
					mock.AnythingOfType("*[]plugintypes.TimestampedUnixBig"),
				).Run(func(args mock.Arguments) {
					updates := args[4].(*[]plugintypes.TimestampedUnixBig)
					*updates = []plugintypes.TimestampedUnixBig{{
						Timestamp: 1000,
						Value:     big.NewInt(5e18),
					}}
				}).Return(nil)

				// Mock Bind for FeeQuoter contract
				cr.On("Bind", mock.Anything, mock.MatchedBy(func(contracts []types.BoundContract) bool {
					return len(contracts) == 1 &&
						contracts[0].Name == consts.ContractNameFeeQuoter &&
						contracts[0].Address == typeconv.AddressBytesToString(feeQuoterAddr, uint64(testChain))
				})).Return(nil)
			},
			want: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
				ArbAddr: {
					Timestamp: time.Unix(1000, 0),
					Value:     cciptypes.NewBigInt(big.NewInt(5e18)),
				},
			},
		},
		{
			name:        "success - multiple tokens",
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr, EthAddr},
			setup: func(cr *reader_mocks.MockContractReaderFacade) {
				cr.On("GetLatestValue",
					mock.Anything,
					mock.MatchedBy(func(s string) bool {
						return strings.Contains(s, consts.MethodNameFeeQuoterGetTokenPrices)
					}),
					primitives.Unconfirmed,
					mock.MatchedBy(func(params map[string]any) bool {
						tokens, ok := params["tokens"].([][]byte)
						return ok && len(tokens) == 2
					}),
					mock.AnythingOfType("*[]plugintypes.TimestampedUnixBig"),
				).Run(func(args mock.Arguments) {
					updates := args[4].(*[]plugintypes.TimestampedUnixBig)
					*updates = []plugintypes.TimestampedUnixBig{
						{
							Timestamp: 2000,
							Value:     big.NewInt(1e18), // ArbAddr price
						},
						{
							Timestamp: 2001,
							Value:     big.NewInt(2e18), // EthAddr price
						},
					}
				}).Return(nil)

				cr.On("Bind", mock.Anything, mock.MatchedBy(func(contracts []types.BoundContract) bool {
					return len(contracts) == 1 &&
						contracts[0].Name == consts.ContractNameFeeQuoter &&
						contracts[0].Address == typeconv.AddressBytesToString(feeQuoterAddr, uint64(testChain))
				})).Return(nil)
			},
			want: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
				ArbAddr: {
					Timestamp: time.Unix(2000, 0),
					Value:     cciptypes.NewBigInt(big.NewInt(1e18)),
				},
				EthAddr: {
					Timestamp: time.Unix(2001, 0),
					Value:     cciptypes.NewBigInt(big.NewInt(2e18)),
				},
			},
		},
		{
			name:        "no tokens provided",
			inputTokens: []cciptypes.UnknownEncodedAddress{},
			setup: func(cr *reader_mocks.MockContractReaderFacade) {
				// If no tokens are provided, no GetLatestValue call should happen.
				// Just mock Bind as it's required by Sync.
				cr.On("Bind", mock.Anything, mock.Anything).Return(nil).Maybe()
			},
			want:    map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{},
			wantErr: false,
		},
		{
			name:        "GetLatestValue returns error",
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr},
			setup: func(cr *reader_mocks.MockContractReaderFacade) {
				cr.On("GetLatestValue",
					mock.Anything,
					mock.MatchedBy(func(s string) bool {
						return strings.Contains(s, consts.MethodNameFeeQuoterGetTokenPrices)
					}),
					primitives.Unconfirmed,
					mock.Anything,
					mock.Anything,
				).Return(fmt.Errorf("some error"))

				cr.On("Bind", mock.Anything, mock.Anything).Return(nil)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name:        "cache usage - token already cached",
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr},
			setup: func(cr *reader_mocks.MockContractReaderFacade) {
				// First call: return a value
				cr.On("GetLatestValue",
					mock.Anything,
					mock.MatchedBy(func(s string) bool {
						return strings.Contains(s, consts.MethodNameFeeQuoterGetTokenPrices)
					}),
					primitives.Unconfirmed,
					mock.MatchedBy(func(params map[string]any) bool {
						tokens, ok := params["tokens"].([][]byte)
						return ok && len(tokens) == 1
					}),
					mock.AnythingOfType("*[]plugintypes.TimestampedUnixBig"),
				).Once().Run(func(args mock.Arguments) {
					updates := args[4].(*[]plugintypes.TimestampedUnixBig)
					*updates = []plugintypes.TimestampedUnixBig{{
						Timestamp: 1500,
						Value:     big.NewInt(7e18),
					}}
				}).Return(nil)

				// Mock Bind for FeeQuoter contract
				cr.On("Bind", mock.Anything, mock.MatchedBy(func(contracts []types.BoundContract) bool {
					return len(contracts) == 1 &&
						contracts[0].Name == consts.ContractNameFeeQuoter &&
						contracts[0].Address == typeconv.AddressBytesToString(feeQuoterAddr, uint64(testChain))
				})).Return(nil)
			},
			want: map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{
				ArbAddr: {
					Timestamp: time.Unix(1500, 0),
					Value:     cciptypes.NewBigInt(big.NewInt(7e18)),
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cr := reader_mocks.NewMockContractReaderFacade(t)
			tc.setup(cr)

			_, pr := setupTestEnvironment(t, lggr, testChain, feeQuoterAddr, cr)

			got, err := pr.GetFeeQuoterTokenUpdates(context.Background(), tc.inputTokens, testChain)

			if tc.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.want, got)

			// If this is the cache test case, let's call again to ensure no new GetLatestValue calls are made
			if tc.name == "cache usage - token already cached" {
				// No additional setup here means no GetLatestValue call is expected this time
				got2, err2 := pr.GetFeeQuoterTokenUpdates(context.Background(), tc.inputTokens, testChain)
				require.NoError(t, err2)
				// Should return the same cached result
				assert.Equal(t, tc.want, got2)
			}
		})
	}
}

func TestOnchainTokenPricesReader_GetTokenPricesUSD(t *testing.T) {
	testCases := []struct {
		name          string
		inputTokens   []cciptypes.UnknownEncodedAddress
		tokenInfo     map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo
		mockPrices    map[cciptypes.UnknownEncodedAddress]*big.Int
		want          []*big.Int
		errorAccounts []cciptypes.UnknownEncodedAddress
		wantErr       bool
	}{
		// {
		// 	name: "On-chain one price",
		// 	tokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
		// 		ArbAddr: ArbInfo,
		// 	},
		// 	inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr},
		// 	mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
		// 	want:        []*big.Int{ArbPrice},
		// },
		{
			name: "On-chain multiple prices",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr, EthAddr},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice, EthAddr: EthPrice},
			want:        []*big.Int{ArbPrice, EthPrice},
		},
		{
			name: "Missing price should error",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
			},
			inputTokens:   []cciptypes.UnknownEncodedAddress{ArbAddr, EthAddr},
			mockPrices:    map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
			errorAccounts: []cciptypes.UnknownEncodedAddress{EthAddr},
			want:          nil,
			wantErr:       true,
		},
		{
			name: "Empty input tokens list",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{},
			want:        []*big.Int{},
		},
		{
			name: "Repeated token in input",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr, ArbAddr},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
			want:        []*big.Int{ArbPrice, ArbPrice},
		},
		{
			name: "Zero price should succeed",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
			},
			inputTokens: []cciptypes.UnknownEncodedAddress{ArbAddr},
			mockPrices:  map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: big.NewInt(0)},
			want:        []*big.Int{big.NewInt(0)},
		},
		{
			name: "Multiple error accounts",
			tokenInfo: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				ArbAddr: ArbInfo,
				EthAddr: EthInfo,
				BtcAddr: BtcInfo,
			},
			inputTokens:   []cciptypes.UnknownEncodedAddress{ArbAddr, EthAddr, BtcAddr},
			mockPrices:    map[cciptypes.UnknownEncodedAddress]*big.Int{ArbAddr: ArbPrice},
			errorAccounts: []cciptypes.UnknownEncodedAddress{EthAddr, BtcAddr},
			want:          nil,
			wantErr:       true,
		},
	}

	for _, tc := range testCases {
		contractReader := createMockReader(t, tc.mockPrices, tc.errorAccounts, tc.tokenInfo)

		feedChain := cciptypes.ChainSelector(1)
		tokenPricesReader := NewPriceReader(
			logger.Test(t),
			map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
				feedChain: contractReader,
			},
			tc.tokenInfo,
			nil,
			feedChain,
		)
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
	mockPrices map[cciptypes.UnknownEncodedAddress]*big.Int,
	errorAccounts []cciptypes.UnknownEncodedAddress,
	tokenInfo map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo,
) *readermock.MockContractReaderFacade {
	reader := readermock.NewMockContractReaderFacade(t)

	// If there are no prices to mock and no error accounts, we don't need to set up any expectations
	if len(mockPrices) == 0 && len(errorAccounts) == 0 {
		return reader
	}

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
		priceResult.SetResult(&LatestRoundData{Answer: big.NewInt(price.Int64())}, nil)
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

	// Set up the mock expectation for BatchGetLatestValues only if we have results to return
	if len(expectedResults) > 0 {
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
	}

	return reader
}
