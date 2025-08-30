package plugincommon

import (
	"fmt"
	"sort"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
)

// GetTransmissionSchedule returns a TransmissionSchedule for the provided oracles.
// It uses the ChainSupport service to query which oracles support the destination chain.
// It returns an error if no oracles support the destination chain.
// Read more about TransmissionDelay at ocr3types.TransmissionSchedule
// The transmissionDelayMultiplier is used in the following way:
//
//	Assume that we have transmitters: [1, 3, 5]
//	And transmissionDelayMultiplier = 5s
//	Then the transmission delays will be: [0s, 5s, 10s]
func GetTransmissionSchedule(
	chainSupport ChainSupport,
	allTheOracles []commontypes.OracleID,
	transmissionDelayMultiplier time.Duration,
) (*ocr3types.TransmissionSchedule, error) {
	transmitters := make([]commontypes.OracleID, 0, len(allTheOracles))
	for _, oracleID := range allTheOracles {
		supportsDestChain, err := chainSupport.SupportsDestChain(oracleID)
		if err != nil {
			return nil, fmt.Errorf("supports dest chain %d: %w", oracleID, err)
		}
		if supportsDestChain {
			transmitters = append(transmitters, oracleID)
		}
	}

	// transmissionSchedule must be deterministic
	sort.Slice(transmitters, func(i, j int) bool { return transmitters[i] < transmitters[j] })

	transmissionDelays := make([]time.Duration, len(transmitters))

	for i := range transmissionDelays {
		transmissionDelays[i] = (transmissionDelayMultiplier) * time.Duration(i)
	}

	if len(transmitters) == 0 {
		return nil, fmt.Errorf("no transmitters")
	}

	if len(transmitters) != len(transmissionDelays) {
		return nil, fmt.Errorf("critical issue mismatched transmitters and transmission delays")
	}

	return &ocr3types.TransmissionSchedule{
		Transmitters:       transmitters,
		TransmissionDelays: transmissionDelays,
	}, nil
}
