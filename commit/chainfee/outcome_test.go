package chainfee

import (
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal"

	"github.com/stretchr/testify/mock"

	mapset "github.com/deckarep/golang-set/v2"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	mock_home_chain "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

var ts = time.Now().UTC()

var feeComponentsMap = map[cciptypes.ChainSelector]types.ChainFeeComponents{
	internal.EvmChainSelector:  {ExecutionFee: big.NewInt(100), DataAvailabilityFee: big.NewInt(200)},
	internal.EvmChainSelector2: {ExecutionFee: big.NewInt(150), DataAvailabilityFee: big.NewInt(250)},
}

var chainFeePriceBatchWriteFrequency = *commonconfig.MustNewDuration(time.Minute)

var nativeTokenPricesMap = map[cciptypes.ChainSelector]cciptypes.BigInt{
	internal.EvmChainSelector:  cciptypes.NewBigInt(big.NewInt(1e18)),
	internal.EvmChainSelector2: cciptypes.NewBigInt(big.NewInt(2e18)),
}

var fChains = map[cciptypes.ChainSelector]int{
	internal.EvmChainSelector:  1,
	internal.EvmChainSelector2: 2,
}

var obsNeedUpdate = Observation{
	FeeComponents:     feeComponentsMap,
	NativeTokenPrices: nativeTokenPricesMap,
	FChain:            fChains,
	ChainFeeUpdates: map[cciptypes.ChainSelector]Update{
		internal.EvmChainSelector: {
			Timestamp: ts,
			ChainFee: ComponentsUSDPrices{
				ExecutionFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector,
					feeComponentsMap[internal.EvmChainSelector].ExecutionFee,
					nativeTokenPricesMap[internal.EvmChainSelector].Int,
				),
				DataAvFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector,
					feeComponentsMap[internal.EvmChainSelector].DataAvailabilityFee,
					nativeTokenPricesMap[internal.EvmChainSelector].Int,
				),
			},
		},
		internal.EvmChainSelector2: {
			// Need update because timestamp is older than batch write frequency
			Timestamp: ts.Add(-chainFeePriceBatchWriteFrequency.Duration() * 2),
			ChainFee: ComponentsUSDPrices{
				ExecutionFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector2,
					feeComponentsMap[internal.EvmChainSelector2].ExecutionFee,
					nativeTokenPricesMap[internal.EvmChainSelector2].Int,
				),
				DataAvFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector2,
					feeComponentsMap[internal.EvmChainSelector2].DataAvailabilityFee,
					nativeTokenPricesMap[internal.EvmChainSelector2].Int,
				),
			},
		},
	},
	TimestampNow: ts,
}

var obsNoUpdate = Observation{
	FeeComponents:     feeComponentsMap,
	NativeTokenPrices: nativeTokenPricesMap,
	FChain:            fChains,
	ChainFeeUpdates: map[cciptypes.ChainSelector]Update{
		internal.EvmChainSelector: {
			ChainFee: ComponentsUSDPrices{
				ExecutionFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector,
					feeComponentsMap[internal.EvmChainSelector].ExecutionFee,
					nativeTokenPricesMap[internal.EvmChainSelector].Int,
				),
				DataAvFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector,
					feeComponentsMap[internal.EvmChainSelector].DataAvailabilityFee,
					nativeTokenPricesMap[internal.EvmChainSelector].Int,
				),
			},
			Timestamp: ts,
		},
		internal.EvmChainSelector2: {
			ChainFee: ComponentsUSDPrices{
				ExecutionFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector2,
					feeComponentsMap[internal.EvmChainSelector2].ExecutionFee,
					nativeTokenPricesMap[internal.EvmChainSelector2].Int,
				),
				DataAvFeePriceUSD: internal.MustCalculateUsdPerUnitGas(
					internal.EvmChainSelector2,
					feeComponentsMap[internal.EvmChainSelector2].DataAvailabilityFee,
					nativeTokenPricesMap[internal.EvmChainSelector2].Int,
				),
			},
			Timestamp: ts,
		},
	},
	TimestampNow: ts,
}

var defaultChainConfig = reader.ChainConfig{
	FChain: 1,
	// not necessary for test, using some peerIDs
	SupportedNodes: mapset.NewSet(libocrtypes.PeerID{1}, libocrtypes.PeerID{2}),
	Config: chainconfig.ChainConfig{
		GasPriceDeviationPPB:    cciptypes.NewBigInt(big.NewInt(1)),
		DAGasPriceDeviationPPB:  cciptypes.NewBigInt(big.NewInt(1)),
		OptimisticConfirmations: 1,
	},
}

// sameObs returns n observations with the same observation but from different oracle ids
func sameObs(n int, obs Observation) []plugincommon.AttributedObservation[Observation] {
	aos := make([]plugincommon.AttributedObservation[Observation], n)
	for i := 0; i < n; i++ {
		aos[i] = plugincommon.AttributedObservation[Observation]{OracleID: commontypes.OracleID(i), Observation: obs}
	}
	return aos
}

type FeeInfo struct {
	// ExecDeviationPPB is the deviation threshold in parts per billion that determines whether or not
	// the exec portion of the gas price has deviated and needs to be reported on chain.
	ExecDeviationPPB cciptypes.BigInt `json:"execDeviationPPB"`

	// DataAvailabilityDeviationPPB is the deviation threshold in parts per billion that determines whether or not
	// the data availability portion of the gas price has deviated and needs to be reported on chain.
	DataAvailabilityDeviationPPB cciptypes.BigInt `json:"dataAvailabilityDeviationPPB"`

	// ChainFeeDeviationDisabled is a flag to disable deviation-based reporting. If true, we will only report
	// prices based on the heartbeat.
	ChainFeeDeviationDisabled bool `json:"chainFeeDeviationDisabled"`
}

func TestGetConsensusObservation(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr:      lggr,
		destChain: internal.EvmChainSelector,
		fRoleDON:  1,
	}

	// 3 oracles, same observations, will pass destChain 2f+1 for chain selector 1
	aos := sameObs(3, obsNeedUpdate)

	consensusObs, err := p.getConsensusObservation(lggr, aos)
	require.NoError(t, err)
	assert.Equal(t, fChains[1], consensusObs.FChain[1])
	assert.Equal(t, fChains[2], consensusObs.FChain[2])

	// Only chain selector 1 will have consensus
	// That's why we assert having only 1 fee component, and 1 native token price.
	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.TimestampNow)
	assert.Len(t, consensusObs.FeeComponents, 1)
	assert.Equal(t, feeComponentsMap[1], consensusObs.FeeComponents[1])
	assert.Len(t, consensusObs.NativeTokenPrices, 1)
	assert.Equal(t, nativeTokenPricesMap[1], consensusObs.NativeTokenPrices[1])

	// 5 oracles, same observations, will pass destChain 2f+1 for both chain selectors
	aos = sameObs(5, obsNeedUpdate)

	consensusObs, err = p.getConsensusObservation(lggr, aos)
	require.NoError(t, err)
	assert.Equal(t, fChains[1], consensusObs.FChain[1])
	assert.Equal(t, fChains[2], consensusObs.FChain[2])

	// Both chain selectors 1 and 2 will have consensus
	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.TimestampNow)
	assert.Len(t, consensusObs.FeeComponents, 2)
	assert.Equal(t, feeComponentsMap, consensusObs.FeeComponents)
	assert.Len(t, consensusObs.NativeTokenPrices, 2)
}

func TestProcessor_Outcome(t *testing.T) {
	oneMinuteAgo := time.Now().Add(-time.Minute).UTC()
	numOracles := 5 // Use a consistent number for generating aos

	cases := []struct {
		name                   string
		chainFeeWriteFrequency commonconfig.Duration
		feeInfo                map[cciptypes.ChainSelector]FeeInfo
		aos                    []plugincommon.AttributedObservation[Observation]
		expectedError          bool
		expectedOutcome        func() Outcome
	}{
		{
			name:          "Outcome gas prices when earliest update is before batch write frequency duration",
			aos:           sameObs(numOracles, obsNeedUpdate),
			expectedError: false,
			expectedOutcome: func() Outcome {
				gas2 := new(big.Int)
				// {ExecutionFee: big.NewInt(150), DataAvailabilityFee: big.NewInt(250)}
				// (250 * 2e18/e18) << 112 | (150 * 2e18/e18) -- check `CalculateUsdPerUnitGas`
				gas2, ok := gas2.SetString("2596148429267413814265248164610048300", 10) // base 10
				require.True(t, ok)
				// Only chain selector 2 will be updated because last update is stale
				expectedOutcome := Outcome{
					GasPrices: []cciptypes.GasPriceChain{
						{
							ChainSel: internal.EvmChainSelector2,
							GasPrice: cciptypes.NewBigInt(gas2),
						},
					},
				}
				return expectedOutcome
			},
			chainFeeWriteFrequency: chainFeePriceBatchWriteFrequency,
		},
		{
			name:          "no consensus",
			aos:           []plugincommon.AttributedObservation[Observation]{}, // Empty aos slice
			expectedError: true,                                                // No f chains to calculate consensus
			expectedOutcome: func() Outcome {
				return Outcome{}
			},
		},
		{
			name:          "Empty Outcome when no need to update",
			aos:           sameObs(numOracles, obsNoUpdate),
			expectedError: false,
			expectedOutcome: func() Outcome {
				return Outcome{}
			},
			chainFeeWriteFrequency: chainFeePriceBatchWriteFrequency,
		},
		{
			name:                   "happy path with a price deviation",
			chainFeeWriteFrequency: *commonconfig.MustNewDuration(time.Hour),
			feeInfo: map[cciptypes.ChainSelector]FeeInfo{
				internal.EvmChainSelector: {
					ExecDeviationPPB:             cciptypes.NewBigInt(big.NewInt(1)),
					DataAvailabilityDeviationPPB: cciptypes.NewBigInt(big.NewInt(1)),
				},
				internal.EvmChainSelector2: {
					ExecDeviationPPB:             cciptypes.NewBigInt(big.NewInt(1)),
					DataAvailabilityDeviationPPB: cciptypes.NewBigInt(big.NewInt(1)),
				},
			},
			aos: sameObs(numOracles, Observation{
				FeeComponents: map[cciptypes.ChainSelector]types.ChainFeeComponents{
					internal.EvmChainSelector:  {ExecutionFee: big.NewInt(2), DataAvailabilityFee: big.NewInt(1)},
					internal.EvmChainSelector2: {ExecutionFee: big.NewInt(2), DataAvailabilityFee: big.NewInt(1)},
				},
				NativeTokenPrices: map[cciptypes.ChainSelector]cciptypes.BigInt{
					// <----------- token price increased deviation reached
					internal.EvmChainSelector: cciptypes.NewBigInt(big.NewInt(2e18)),
					// <----------- token price same deviation not reached
					internal.EvmChainSelector2: cciptypes.NewBigInt(big.NewInt(1e18)),
				},
				ChainFeeUpdates: map[cciptypes.ChainSelector]Update{
					internal.EvmChainSelector: {
						Timestamp: oneMinuteAgo,
						ChainFee: ComponentsUSDPrices{
							ExecutionFeePriceUSD: big.NewInt(2), DataAvFeePriceUSD: big.NewInt(1),
						},
					},
					internal.EvmChainSelector2: {
						Timestamp: oneMinuteAgo,
						ChainFee: ComponentsUSDPrices{
							ExecutionFeePriceUSD: big.NewInt(2), DataAvFeePriceUSD: big.NewInt(1),
						},
					},
				},
				FChain:       map[cciptypes.ChainSelector]int{internal.EvmChainSelector: 1},
				TimestampNow: time.Now().UTC(),
			}),
			expectedError: false,
			expectedOutcome: func() Outcome {
				var b big.Int
				exp, ok := b.SetString("10384593717069655257060992658440196", 10)
				require.True(t, ok)
				return Outcome{
					GasPrices: []cciptypes.GasPriceChain{
						{GasPrice: cciptypes.NewBigInt(exp), ChainSel: internal.EvmChainSelector}, // only chainSel=1
					},
				}
			},
		},
		{
			name:                   "deviation check disabled",
			chainFeeWriteFrequency: *commonconfig.MustNewDuration(time.Hour),
			feeInfo: map[cciptypes.ChainSelector]FeeInfo{
				internal.EvmChainSelector: {
					ExecDeviationPPB:             cciptypes.NewBigInt(big.NewInt(1)),
					DataAvailabilityDeviationPPB: cciptypes.NewBigInt(big.NewInt(1)),
					ChainFeeDeviationDisabled:    true,
				},
			},
			aos: sameObs(numOracles, Observation{
				FeeComponents: map[cciptypes.ChainSelector]types.ChainFeeComponents{
					internal.EvmChainSelector: {ExecutionFee: big.NewInt(2), DataAvailabilityFee: big.NewInt(1)},
				},
				NativeTokenPrices: map[cciptypes.ChainSelector]cciptypes.BigInt{
					// <----------- token price increased deviation reached
					internal.EvmChainSelector: cciptypes.NewBigInt(big.NewInt(2e18)),
				},
				ChainFeeUpdates: map[cciptypes.ChainSelector]Update{
					internal.EvmChainSelector: {
						Timestamp: oneMinuteAgo,
						ChainFee: ComponentsUSDPrices{
							ExecutionFeePriceUSD: big.NewInt(2), DataAvFeePriceUSD: big.NewInt(1),
						},
					},
				},
				FChain:       map[cciptypes.ChainSelector]int{internal.EvmChainSelector: 1},
				TimestampNow: time.Now().UTC(),
			}),
			expectedError: false,
			expectedOutcome: func() Outcome {
				return Outcome{
					GasPrices: nil,
				}
			},
		},
		{
			name: "Empty Observations with only FChain and TimestampNow",
			aos: func() []plugincommon.AttributedObservation[Observation] {
				aos := make([]plugincommon.AttributedObservation[Observation], numOracles)
				for i := 0; i < numOracles; i++ {
					aos[i] = plugincommon.AttributedObservation[Observation]{
						OracleID: commontypes.OracleID(i),
						Observation: Observation{
							FChain:       fChains,
							TimestampNow: ts,
							// FeeComponents, NativeTokenPrices, ChainFeeUpdates are zero/empty
						},
					}
				}
				return aos
			}(),
			expectedError: false, // No error expected, just an empty outcome
			expectedOutcome: func() Outcome {
				return Outcome{}
			},
			chainFeeWriteFrequency: chainFeePriceBatchWriteFrequency, // Needs a frequency
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tests.Context(t)
			homeChainMock := mock_home_chain.NewMockHomeChain(t)
			if tt.feeInfo == nil {
				homeChainMock.EXPECT().GetChainConfig(mock.Anything).
					Return(defaultChainConfig, nil).Maybe()
			}
			for chain, info := range tt.feeInfo {
				homeChainMock.EXPECT().GetChainConfig(chain).Return(reader.ChainConfig{
					FChain: 1,
					// not necessary for test, using some peerIDs
					SupportedNodes: mapset.NewSet(libocrtypes.PeerID{1}, libocrtypes.PeerID{2}),
					Config: chainconfig.ChainConfig{
						GasPriceDeviationPPB:      info.ExecDeviationPPB,
						DAGasPriceDeviationPPB:    info.DataAvailabilityDeviationPPB,
						OptimisticConfirmations:   1,
						ChainFeeDeviationDisabled: info.ChainFeeDeviationDisabled,
					},
				}, nil).Maybe()
			}
			p := &processor{
				lggr:      logger.Test(t),
				destChain: internal.EvmChainSelector,
				fRoleDON:  1,
				cfg: pluginconfig.CommitOffchainConfig{
					RemoteGasPriceBatchWriteFrequency: tt.chainFeeWriteFrequency,
				},
				metricsReporter: plugincommon.NoopReporter{},
				homeChain:       homeChainMock,
			}

			outcome, err := p.Outcome(ctx, Outcome{}, Query{}, tt.aos)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedOutcome(), outcome)
			}
		})
	}
}
