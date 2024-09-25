package chainfee

import (
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var ts = time.Now().UTC()

var feeComponentsMap = map[cciptypes.ChainSelector]types.ChainFeeComponents{
	1: {ExecutionFee: big.NewInt(100), DataAvailabilityFee: big.NewInt(200)},
	2: {ExecutionFee: big.NewInt(150), DataAvailabilityFee: big.NewInt(250)},
}

var chainFeePriceBatchWriteFrequency = *commonconfig.MustNewDuration(time.Minute)
var nativeTokenPricesMap = map[cciptypes.ChainSelector]cciptypes.BigInt{
	1: cciptypes.NewBigInt(big.NewInt(10)),
	2: cciptypes.NewBigInt(big.NewInt(20)),
}

var fChains = map[cciptypes.ChainSelector]int{
	1: 1,
	2: 2,
}

var obsNeedUpdate = Observation{
	FeeComponents:    feeComponentsMap,
	NativeTokenPrice: nativeTokenPricesMap,
	FChain:           fChains,
	ChainFeeLatestUpdates: map[cciptypes.ChainSelector]time.Time{
		1: ts,
		2: ts.Add(-chainFeePriceBatchWriteFrequency.Duration() * 2), // Needs updating
	},
	Timestamp: ts,
}

var obsNoUpdate = Observation{
	FeeComponents:    feeComponentsMap,
	NativeTokenPrice: nativeTokenPricesMap,
	FChain:           fChains,
	ChainFeeLatestUpdates: map[cciptypes.ChainSelector]time.Time{
		1: ts,
		2: ts,
	},
	Timestamp: ts,
}

func SameObs(n int, obs Observation) []plugincommon.AttributedObservation[Observation] {
	aos := make([]plugincommon.AttributedObservation[Observation], n)
	for i := 0; i < n; i++ {
		aos[i] = plugincommon.AttributedObservation[Observation]{OracleID: commontypes.OracleID(i), Observation: obs}
	}
	return aos
}
func TestGetConsensusObservation(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr:      lggr,
		destChain: 1,
		fRoleDON:  1,
	}

	// 3 oracles, same observations, will pass destChain 2f+1 for chain selector 1
	aos := SameObs(3, obsNeedUpdate)

	consensusObs, err := p.getConsensusObservation(aos)
	require.NoError(t, err)
	assert.Equal(t, fChains[1], consensusObs.FChain[1])
	assert.Equal(t, fChains[2], consensusObs.FChain[2])

	// Only chain selector 1 will have consensus
	// That's why we assert having only 1 fee component, and 1 native token price.
	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.Timestamp)
	assert.Len(t, consensusObs.FeeComponents, 1)
	assert.Equal(t, feeComponentsMap[1], consensusObs.FeeComponents[1])
	assert.Len(t, consensusObs.NativeTokenPrices, 1)
	assert.Equal(t, nativeTokenPricesMap[1], consensusObs.NativeTokenPrices[1])

	// 5 oracles, same observations, will pass destChain 2f+1 for both chain selectors
	aos = SameObs(5, obsNeedUpdate)

	consensusObs, err = p.getConsensusObservation(aos)
	require.NoError(t, err)
	assert.Equal(t, fChains[1], consensusObs.FChain[1])
	assert.Equal(t, fChains[2], consensusObs.FChain[2])

	// Both chain selectors 1 and 2 will have consensus
	assert.NotNil(t, consensusObs)
	assert.Equal(t, ts, consensusObs.Timestamp)
	assert.Len(t, consensusObs.FeeComponents, 2)
	assert.Equal(t, feeComponentsMap, consensusObs.FeeComponents)
	assert.Len(t, consensusObs.NativeTokenPrices, 2)
}

func TestOutcome(t *testing.T) {
	lggr := logger.Test(t)
	p := &processor{
		lggr:                             lggr,
		destChain:                        1,
		fRoleDON:                         1,
		ChainFeePriceBatchWriteFrequency: chainFeePriceBatchWriteFrequency,
	}

	outcome, err := p.Outcome(Outcome{}, Query{}, SameObs(5, obsNeedUpdate))

	gas1 := new(big.Int)
	// (200 * 10) << 112 | (100 * 10)
	gas1, ok := gas1.SetString("10384593717069655257060992658440193000", 10) // base 10
	require.True(t, ok)

	gas2 := new(big.Int)
	// (250 * 20) << 112 | (150 * 20)
	gas2, ok = gas2.SetString("25961484292674138142652481646100483000", 10) // base 10
	require.True(t, ok)
	// Gas prices sorted by chain selector
	expectedOutcome := Outcome{
		GasPrices: []cciptypes.GasPriceChain{
			{
				ChainSel: 1,
				GasPrice: cciptypes.NewBigInt(gas1),
			},
			{
				ChainSel: 2,
				GasPrice: cciptypes.NewBigInt(gas2),
			},
		},
	}

	require.NoError(t, err)
	assert.Equal(t, expectedOutcome, outcome)
}

func TestProcessor_Outcome(t *testing.T) {
	tests := []struct {
		name                   string
		chainFeeWriteFrequency commonconfig.Duration
		chainFeeLatestUpdates  map[cciptypes.ChainSelector]time.Time
		aos                    []plugincommon.AttributedObservation[Observation]
		expectedError          bool
		expectedOutcome        func() Outcome
	}{
		{
			name:          "Outcome gas prices when earliest update is before batch write frequency duration",
			aos:           SameObs(5, obsNeedUpdate),
			expectedError: false,
			expectedOutcome: func() Outcome {
				gas1 := new(big.Int)
				// {ExecutionFee: big.NewInt(100), DataAvailabilityFee: big.NewInt(200)},
				// (200 * 10) << 112 | (100 * 10)
				gas1, ok := gas1.SetString("10384593717069655257060992658440193000", 10) // base 10
				require.True(t, ok)

				gas2 := new(big.Int)
				// {ExecutionFee: big.NewInt(150), DataAvailabilityFee: big.NewInt(250)}
				// (250 * 20) << 112 | (150 * 20)
				gas2, ok = gas2.SetString("25961484292674138142652481646100483000", 10) // base 10
				require.True(t, ok)
				// Gas prices sorted by chain selector
				expectedOutcome := Outcome{
					GasPrices: []cciptypes.GasPriceChain{
						{
							ChainSel: 1,
							GasPrice: cciptypes.NewBigInt(gas1),
						},
						{
							ChainSel: 2,
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
			aos:           []plugincommon.AttributedObservation[Observation]{},
			expectedError: true, // No f chains to calculate consensus
			expectedOutcome: func() Outcome {
				return Outcome{}
			},
		},
		{
			name:          "Empty Outcome when no need to update",
			aos:           SameObs(5, obsNoUpdate),
			expectedError: false,
			expectedOutcome: func() Outcome {
				return Outcome{}
			},
			chainFeeWriteFrequency: chainFeePriceBatchWriteFrequency,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &processor{
				destChain:                        1,
				fRoleDON:                         1,
				ChainFeePriceBatchWriteFrequency: tt.chainFeeWriteFrequency,
			}

			outcome, err := p.Outcome(Outcome{}, Query{}, tt.aos)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedOutcome(), outcome)
			}
		})
	}
}
