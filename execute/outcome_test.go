package execute

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestOutcome_NonDeterministicObservationsThatAreHittingSizeLimits(t *testing.T) {
	ctx := t.Context()

	var buf bytes.Buffer
	lggr := logger.NewWithSync(&buf)

	destChain := cciptypes.ChainSelector(999)
	srcChain1 := cciptypes.ChainSelector(111)
	srcChain2 := cciptypes.ChainSelector(222)

	oracleObservations := make(map[commontypes.OracleID]exectypes.Observation)

	// oracles 1 and 2 make a full observation with populated msgs
	// oracles 3 and 4 observe non-populated msgs
	for _, oracleID := range []uint8{1, 2, 3, 4} {
		obs := exectypes.Observation{
			CommitReports: nil,
			Messages: exectypes.MessageObservations{
				srcChain1: map[cciptypes.SeqNum]cciptypes.Message{
					4001: {
						Header: cciptypes.RampMessageHeader{
							MessageID:      cciptypes.Bytes32{4, 0, 0, 1},
							SequenceNumber: 4001,
						},
						Sender: []byte{1, 1, 1, 1, 1},
					},
				},
				srcChain2: map[cciptypes.SeqNum]cciptypes.Message{
					5001: {
						Header: cciptypes.RampMessageHeader{
							MessageID:      cciptypes.Bytes32{5, 0, 0, 1},
							SequenceNumber: 5001,
						},
						Sender: []byte{1, 1, 1, 1, 1},
					},
				},
			},
			Hashes: exectypes.MessageHashes{
				srcChain1: map[cciptypes.SeqNum]cciptypes.Bytes32{4001: {1, 1, 1, 1, 1, 1}},
				srcChain2: map[cciptypes.SeqNum]cciptypes.Bytes32{5001: {1, 1, 1, 1, 1, 2}},
			},
			FChain: map[cciptypes.ChainSelector]int{
				srcChain1: 1,
				srcChain2: 1,
				destChain: 1,
			},
		}

		// oracles 3 and 4 do not observe populated msg on chain 2 (sufficient to just populated the sender field)
		if oracleID == 3 || oracleID == 4 {
			msg2 := obs.Messages[srcChain2][5001]
			msg2.Sender = []byte{}
			obs.Messages[srcChain2][5001] = msg2
		}
		oracleObservations[commontypes.OracleID(oracleID)] = obs
	}

	p := &Plugin{
		ocrTypeCodec: ocrTypeCodec,
		lggr:         lggr,
		destChain:    destChain,
	}

	prevOutcome := exectypes.Outcome{
		State: exectypes.GetCommitReports,
	}
	prevOutcomeB, err := ocrTypeCodec.EncodeOutcome(prevOutcome)
	require.NoError(t, err)

	aos := make([]types.AttributedObservation, 0)
	for oracleID, observation := range oracleObservations {
		b, err := ocrTypeCodec.EncodeObservation(observation)
		require.NoError(t, err)
		aos = append(aos, types.AttributedObservation{
			Observation: b,
			Observer:    oracleID,
		})
	}

	_, err = p.Outcome(ctx, ocr3types.OutcomeContext{
		SeqNr:           0,
		PreviousOutcome: prevOutcomeB,
	}, types.Query{}, aos)
	require.NoError(t, err)

	errorForReproduce1 := "unexpected number of message hashes: expected 1, got 2"
	errorForReproduce2 := "more than one message reached consensus for a sequence number, skipping it"
	if strings.Contains(buf.String(), errorForReproduce1) && strings.Contains(buf.String(), errorForReproduce2) {
		t.Log("Reproduced!")
	}
	fmt.Println(buf.String())
}
