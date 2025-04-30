package chainfee

import (
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_validateFeeComponentsAndChainFeeUpdates(t *testing.T) {
	fourHoursAgo := time.Now().Add(-4 * time.Hour).UTC().Truncate(time.Hour)
	tests := []struct {
		name                         string
		feeComponents                map[ccipocr3.ChainSelector]types.ChainFeeComponents
		chainFeeUpdates              map[ccipocr3.ChainSelector]Update
		expectedFeeComponentError    string
		expectedChainFeeUpdatesError string
	}{
		{
			name: "valid fee components",
			feeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(10),
					DataAvailabilityFee: big.NewInt(20),
				},
			},
			expectedFeeComponentError: "",
		},
		{
			name: "nil execution fee",
			feeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        nil,
					DataAvailabilityFee: big.NewInt(20),
				},
			},
			expectedFeeComponentError: "nil or negative execution fee: <nil>",
		},
		{
			name: "negative execution fee",
			feeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(-1),
					DataAvailabilityFee: big.NewInt(20),
				},
			},
			expectedFeeComponentError: "nil or negative execution fee: -1",
		},
		{
			name: "nil data availability fee",
			feeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(10),
					DataAvailabilityFee: nil,
				},
			},
			expectedFeeComponentError: "nil or negative data availability fee: <nil>",
		},
		{
			name: "negative data availability fee",
			feeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
				1: {
					ExecutionFee:        big.NewInt(10),
					DataAvailabilityFee: big.NewInt(-1),
				},
			},
			expectedFeeComponentError: "nil or negative data availability fee: -1",
		},
		{
			name: "valid chain fee updates",
			chainFeeUpdates: map[ccipocr3.ChainSelector]Update{
				1: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(10),
						DataAvFeePriceUSD:    big.NewInt(20),
					},
					Timestamp: fourHoursAgo,
				},
			},
			expectedChainFeeUpdatesError: "",
		},
		{
			name: "nil execution fee price - chain fee updates",
			chainFeeUpdates: map[ccipocr3.ChainSelector]Update{
				1: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: nil,
						DataAvFeePriceUSD:    big.NewInt(20),
					},
					Timestamp: fourHoursAgo,
				},
			},
			expectedChainFeeUpdatesError: "nil or negative execution fee: <nil>",
		},
		{
			name: "negative execution fee price - chain fee updates",
			chainFeeUpdates: map[ccipocr3.ChainSelector]Update{
				1: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(-1),
						DataAvFeePriceUSD:    big.NewInt(20),
					},
					Timestamp: fourHoursAgo,
				},
			},
			expectedChainFeeUpdatesError: "nil or negative execution fee: -1",
		},
		{
			name: "nil data availability fee price - chain fee updates",
			chainFeeUpdates: map[ccipocr3.ChainSelector]Update{
				1: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(10),
						DataAvFeePriceUSD:    nil,
					},
					Timestamp: fourHoursAgo,
				},
			},
			expectedChainFeeUpdatesError: "nil or negative data availability fee: <nil>",
		},
		{
			name: "negative data availability fee price - chain fee updates",
			chainFeeUpdates: map[ccipocr3.ChainSelector]Update{
				1: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(10),
						DataAvFeePriceUSD:    big.NewInt(-1),
					},
					Timestamp: fourHoursAgo,
				},
			},
			expectedChainFeeUpdatesError: "nil or negative data availability fee: -1",
		},
		{
			name: "zero timestamp - chain fee updates",
			chainFeeUpdates: map[ccipocr3.ChainSelector]Update{
				1: {
					ChainFee: ComponentsUSDPrices{
						ExecutionFeePriceUSD: big.NewInt(10),
						DataAvFeePriceUSD:    big.NewInt(20),
					},
				},
			},
			expectedChainFeeUpdatesError: "timestamp cannot be zero",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ao := plugincommon.AttributedObservation[Observation]{
				Observation: Observation{
					FeeComponents:   tt.feeComponents,
					ChainFeeUpdates: tt.chainFeeUpdates,
				},
			}
			err := validateFeeComponents(ao)
			if tt.expectedFeeComponentError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedFeeComponentError)
			}
			err = validateChainFeeUpdates(ao)
			if tt.expectedChainFeeUpdatesError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedChainFeeUpdatesError)
			}
		})
	}
}

func Test_validateObservedChains(t *testing.T) {
	fourHoursAgo := time.Now().Add(-4 * time.Hour).UTC().Truncate(time.Hour)
	tests := []struct {
		name                    string
		ao                      plugincommon.AttributedObservation[Observation]
		observerSupportedChains mapset.Set[ccipocr3.ChainSelector]
		expectedError           string
	}{
		{
			name: "valid observed chains",
			ao: plugincommon.AttributedObservation[Observation]{
				Observation: Observation{
					FeeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
						1: {
							ExecutionFee:        big.NewInt(10),
							DataAvailabilityFee: big.NewInt(20),
						},
					},
					NativeTokenPrices: map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						1: {
							Int: big.NewInt(100),
						},
					},
					ChainFeeUpdates: map[ccipocr3.ChainSelector]Update{
						1: {
							ChainFee: ComponentsUSDPrices{
								ExecutionFeePriceUSD: big.NewInt(10),
								DataAvFeePriceUSD:    big.NewInt(20),
							},
							Timestamp: fourHoursAgo,
						},
					},
				},
			},
			observerSupportedChains: mapset.NewSet[ccipocr3.ChainSelector](1),
			expectedError:           "",
		},
		{
			name: "unsupported chain",
			ao: plugincommon.AttributedObservation[Observation]{
				Observation: Observation{
					FeeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
						1: {
							ExecutionFee:        big.NewInt(10),
							DataAvailabilityFee: big.NewInt(20),
						},
					},
					NativeTokenPrices: map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						1: {
							Int: big.NewInt(100),
						},
					},
					ChainFeeUpdates: map[ccipocr3.ChainSelector]Update{
						1: {
							ChainFee: ComponentsUSDPrices{
								ExecutionFeePriceUSD: big.NewInt(10),
								DataAvFeePriceUSD:    big.NewInt(20),
							},
							Timestamp: fourHoursAgo,
						},
					},
				},
			},
			observerSupportedChains: mapset.NewSet[ccipocr3.ChainSelector](2),
			expectedError:           "chain 1 is not supported by observer",
		},
		{
			name: "different observed chains in fee components and native token prices",
			ao: plugincommon.AttributedObservation[Observation]{
				Observation: Observation{
					FeeComponents: map[ccipocr3.ChainSelector]types.ChainFeeComponents{
						1: {
							ExecutionFee:        big.NewInt(10),
							DataAvailabilityFee: big.NewInt(20),
						},
					},
					NativeTokenPrices: map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						2: {
							Int: big.NewInt(100),
						},
					},
					ChainFeeUpdates: map[ccipocr3.ChainSelector]Update{
						1: {
							ChainFee: ComponentsUSDPrices{
								ExecutionFeePriceUSD: big.NewInt(10),
								DataAvFeePriceUSD:    big.NewInt(20),
							},
							Timestamp: fourHoursAgo,
						},
					},
				},
			},
			observerSupportedChains: mapset.NewSet[ccipocr3.ChainSelector](1, 2),
			expectedError:           "fee components and native token prices have different observed chains",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateObservedChains(tt.ao, tt.observerSupportedChains)
			if tt.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedError)
			}
		})
	}
}
