package plugincommon_test

import (
	"errors"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/require"

	plugincommon2 "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
)

func TestGetTransmissionSchedule(t *testing.T) {
	testCases := []struct {
		name                        string
		allTheOracles               []commontypes.OracleID
		oraclesSupportingDest       []commontypes.OracleID
		transmissionDelayMultiplier time.Duration
		expectedTransmitters        []commontypes.OracleID
		expectedTransmissionDelays  []time.Duration
		expectedError               bool
		chainSupportReturnsError    bool
	}{
		{
			name:                        "no oracles supporting dest leads to error",
			allTheOracles:               []commontypes.OracleID{1, 2, 3},
			oraclesSupportingDest:       []commontypes.OracleID{},
			transmissionDelayMultiplier: 5 * time.Second,
			expectedTransmitters:        []commontypes.OracleID{},
			expectedTransmissionDelays:  []time.Duration{},
			expectedError:               true,
		},
		{
			name:                        "some transmitters supporting dest",
			allTheOracles:               []commontypes.OracleID{1, 2, 3, 4, 5},
			oraclesSupportingDest:       []commontypes.OracleID{1, 3, 4},
			transmissionDelayMultiplier: 5 * time.Second,
			expectedTransmitters:        []commontypes.OracleID{1, 3, 4},
			expectedTransmissionDelays:  []time.Duration{0 * time.Second, 5 * time.Second, 10 * time.Second},
			expectedError:               false,
		},
		{
			name:                        "chainsupport returns error",
			allTheOracles:               []commontypes.OracleID{1, 2, 3},
			oraclesSupportingDest:       []commontypes.OracleID{1, 3},
			transmissionDelayMultiplier: 5 * time.Second,
			expectedError:               true,
			chainSupportReturnsError:    true,
		},
		{
			name:                        "determinism check",
			allTheOracles:               []commontypes.OracleID{3, 1, 2}, // <------ not ordered
			oraclesSupportingDest:       []commontypes.OracleID{1, 3},
			transmissionDelayMultiplier: 5 * time.Second,
			expectedTransmitters:        []commontypes.OracleID{1, 3},
			expectedTransmissionDelays:  []time.Duration{0 * time.Second, 5 * time.Second},
			expectedError:               false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cs := plugincommon.NewMockChainSupport(t)
			destSupportingOraclesSet := mapset.NewSet(tc.oraclesSupportingDest...)
			for _, oracleID := range tc.allTheOracles {
				var err error
				if tc.chainSupportReturnsError {
					err = errors.New("some error")
				}
				cs.On("SupportsDestChain", oracleID).
					Return(destSupportingOraclesSet.Contains(oracleID), err).Maybe()
			}

			transmissionSchedule, err := plugincommon2.GetTransmissionSchedule(
				cs,
				tc.allTheOracles,
				tc.transmissionDelayMultiplier,
			)
			if tc.expectedError {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedTransmitters, transmissionSchedule.Transmitters)
			require.Equal(t, tc.expectedTransmissionDelays, transmissionSchedule.TransmissionDelays)
		})
	}
}
