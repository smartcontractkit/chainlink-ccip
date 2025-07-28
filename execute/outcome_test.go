package execute

import (
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"testing"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func TestOutcome_PseudoDeletedConsensusConflict(t *testing.T) {

	lggr := logger.Test(t)

	destChain := cciptypes.ChainSelector(999)
	srcChain1 := cciptypes.ChainSelector(111)
	srcChain2 := cciptypes.ChainSelector(222)
	srcChain1SeqNr := cciptypes.SeqNum(4001)
	srcChain2SeqNr := cciptypes.SeqNum(5001)
	fullObs := exectypes.Observation{
		CommitReports: nil,
		Messages: exectypes.MessageObservations{
			srcChain1: map[cciptypes.SeqNum]cciptypes.Message{
				srcChain1SeqNr: {
					Header: cciptypes.RampMessageHeader{
						MessageID:      cciptypes.Bytes32{4, 0, 0, 1},
						SequenceNumber: srcChain1SeqNr,
					},
					Sender: []byte{1, 1, 1, 1, 1},
				},
			},
			srcChain2: map[cciptypes.SeqNum]cciptypes.Message{
				srcChain2SeqNr: {
					Header: cciptypes.RampMessageHeader{
						MessageID:      cciptypes.Bytes32{5, 0, 0, 1},
						SequenceNumber: srcChain2SeqNr,
					},
					Sender: []byte{1, 1, 1, 1, 1},
				},
			},
		},
		Hashes: exectypes.MessageHashes{
			srcChain1: map[cciptypes.SeqNum]cciptypes.Bytes32{srcChain1SeqNr: {1, 1, 1, 1, 1, 1}},
			srcChain2: map[cciptypes.SeqNum]cciptypes.Bytes32{srcChain2SeqNr: {1, 1, 1, 1, 1, 2}},
		},
		TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
			srcChain1: {
				srcChain1SeqNr: exectypes.NewMessageTokenData(),
			},
			srcChain2: {
				srcChain2SeqNr: exectypes.NewMessageTokenData(),
			},
		},
		FChain: map[cciptypes.ChainSelector]int{
			srcChain1: 1,
			srcChain2: 1,
			destChain: 1,
		},
	}

	oracleObservations := make(map[commontypes.OracleID]exectypes.Observation)

	// oracles 1 and 2 make a full observation with populated msgs
	// oracles 3 and 4 observe pseudo deleted msgs
	for _, oracleID := range []uint8{1, 2, 3, 4} {
		obs := fullObs

		// oracles 3 and 4 observe chain 2 msg as pseudo deleted
		if oracleID == 3 || oracleID == 4 {
			// Create a pseudo deleted message for srcChain2
			msg2 := createEmptyMessageWithIDAndSeqNum(obs.Messages[srcChain2][srcChain2SeqNr])
			obs.Messages[srcChain2][srcChain2SeqNr] = msg2
		}
		oracleObservations[commontypes.OracleID(oracleID)] = obs
	}

	aos := make([]plugincommon.AttributedObservation[exectypes.Observation], 0)
	for oracleID, observation := range oracleObservations {
		aos = append(aos, plugincommon.AttributedObservation[exectypes.Observation]{
			Observation: observation,
			OracleID:    oracleID,
		})
	}

	consensusObs, err := computeConsensusObservation(lggr, aos, 999, 1)
	require.NoError(t, err)
	require.Len(t, consensusObs.Messages, 2)
	require.Equal(t, fullObs.Messages[srcChain1][srcChain1SeqNr], consensusObs.Messages[srcChain1][srcChain1SeqNr])
	require.Equal(t, fullObs.Messages[srcChain2][srcChain2SeqNr], consensusObs.Messages[srcChain2][srcChain2SeqNr])
}
