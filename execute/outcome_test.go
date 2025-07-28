package execute

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

func TestOutcome_PseudoDeletedConsensusConflict(t *testing.T) {

	lggr := logger.Test(t)

	destChain := cciptypes.ChainSelector(999)
	srcChain1 := cciptypes.ChainSelector(111)
	srcChain2 := cciptypes.ChainSelector(222)
	srcChain1SeqNr := cciptypes.SeqNum(4001)
	srcChain2SeqNr := cciptypes.SeqNum(5001)
	msg1 := cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:           cciptypes.Bytes32{4, 0, 0, 1},
			SequenceNumber:      srcChain1SeqNr,
			SourceChainSelector: srcChain1,
			DestChainSelector:   destChain,
		},
		Sender: []byte{1, 1, 1, 1, 1},
	}
	msg2 := cciptypes.Message{
		Header: cciptypes.RampMessageHeader{
			MessageID:           cciptypes.Bytes32{5, 0, 0, 1},
			SequenceNumber:      srcChain2SeqNr,
			SourceChainSelector: srcChain2,
			DestChainSelector:   destChain,
		},
		Sender: []byte{1, 1, 1, 1, 1},
	}
	tkData1 := exectypes.NewMessageTokenData(exectypes.NewSuccessTokenData([]byte{1, 1, 1, 1, 1, 1}))
	tkData2 := exectypes.NewMessageTokenData(exectypes.NewSuccessTokenData([]byte{1, 1, 1, 1, 1, 2}))

	oracleObservations := make(map[commontypes.OracleID]exectypes.Observation)

	// oracles 1 and 2 make a full observation with populated msgs
	// oracles 3 and 4 observe pseudo deleted msgs
	for _, oracleID := range []uint8{1, 2, 3, 4} {
		obs := exectypes.Observation{
			CommitReports: nil,
			Messages: exectypes.MessageObservations{
				srcChain1: map[cciptypes.SeqNum]cciptypes.Message{
					srcChain1SeqNr: msg1,
				},
				srcChain2: map[cciptypes.SeqNum]cciptypes.Message{
					srcChain2SeqNr: msg2,
				},
			},
			Hashes: exectypes.MessageHashes{
				srcChain1: map[cciptypes.SeqNum]cciptypes.Bytes32{srcChain1SeqNr: {1, 1, 1, 1, 1, 1}},
				srcChain2: map[cciptypes.SeqNum]cciptypes.Bytes32{srcChain2SeqNr: {1, 1, 1, 1, 1, 2}},
			},
			TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
				srcChain1: {
					srcChain1SeqNr: tkData1,
				},
				srcChain2: {
					srcChain2SeqNr: tkData2,
				},
			},
			FChain: map[cciptypes.ChainSelector]int{
				srcChain1: 1,
				srcChain2: 1,
				destChain: 1,
			},
		}

		// oracles 3 and 4 observe chain 2 msg as pseudo deleted
		if oracleID == 3 || oracleID == 4 {
			// Create a pseudo deleted message for srcChain2
			msg2PseudoDeleted := createEmptyMessageWithIDAndSeqNum(obs.Messages[srcChain2][srcChain2SeqNr])
			obs.Messages[srcChain2][srcChain2SeqNr] = msg2PseudoDeleted
			obs.TokenData[srcChain2][srcChain2SeqNr] = exectypes.NewMessageTokenData()
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
	require.Equal(t, msg1, consensusObs.Messages[srcChain1][srcChain1SeqNr])
	require.Equal(t, msg2, consensusObs.Messages[srcChain2][srcChain2SeqNr])
	require.Equal(t, tkData1, consensusObs.TokenData[srcChain1][srcChain1SeqNr])
	require.Equal(t, tkData2, consensusObs.TokenData[srcChain2][srcChain2SeqNr])
}
