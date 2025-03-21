package optimizers

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_truncateLastCommit(t *testing.T) {
	tests := []struct {
		name        string
		chain       cciptypes.ChainSelector
		observation exectypes.Observation
		expected    exectypes.Observation
	}{
		{
			name:  "no commits to truncate",
			chain: 1,
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {},
				},
			},
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {},
				},
			},
		},
		{
			name:  "truncate last commit",
			chain: 1,
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
						{SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1:  {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
						11: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x02}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1:  exectypes.NewMessageTokenData(),
						11: exectypes.NewMessageTokenData(),
					},
				},
			},
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lggr := logger.Test(t)
			op := NewObservationOptimizer(lggr, 10000, ocrtypecodec.DefaultExecCodec)
			truncated, _ := op.truncateLastCommit(tt.observation, tt.chain)
			require.Equal(t, tt.expected, truncated)
		})
	}
}

func Test_truncateChain(t *testing.T) {
	tests := []struct {
		name        string
		chain       cciptypes.ChainSelector
		observation exectypes.Observation
		expected    exectypes.Observation
	}{
		{
			name:  "truncate chain data",
			chain: 1,
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
						{SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
			},
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{},
				Messages:      map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{},
				TokenData:     map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{},
			},
		},
		{
			name:  "truncate non existent chain",
			chain: 2, // non existent chain
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
			},
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lggr := logger.Test(t)
			op := NewObservationOptimizer(lggr, 10000, ocrtypecodec.DefaultExecCodec) // size not important here
			truncated := op.truncateChain(tt.observation, tt.chain)
			require.Equal(t, tt.expected, truncated)
		})
	}
}

func Test_truncateObservation(t *testing.T) {
	tests := []struct {
		name        string
		observation exectypes.Observation
		maxSize     int
		expected    exectypes.Observation
		deletedMsgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum]struct{}
		wantErr     bool
	}{
		{
			name: "no truncation needed",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
			},
			maxSize: 1000,
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
			},
		},
		{
			name: "truncate last commit",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
				Messages: makeMessageObservation(
					map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
						1: {1, 4},
					},
					withData(make([]byte, 100)),
				),
			},
			maxSize: 2600, // this number is calculated by checking encoded sizes for the observation we expect
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
					},
				},
				// Not going to check Messages in the test
			},
			deletedMsgs: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]struct{}{
				1: {
					3: struct{}{},
					4: struct{}{},
				},
			},
		},
		{
			name: "truncate entire chain",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
					},
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
				Messages: makeMessageObservation(
					map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
						1: {1, 2},
						2: {3, 4},
					},
					withData(make([]byte, 100)),
				),
			},
			maxSize: 1789,
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
			},
		},
		{
			name: "truncate one message from first chain",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
					},
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
				Messages: makeMessageObservation(
					map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
						1: {1, 2},
						2: {3, 4},
					},
					withData(make([]byte, 100)),
				),
			},
			maxSize: 3159, // chain 1, message 1 will be truncated
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
					},
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
			},
			deletedMsgs: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]struct{}{
				1: {
					1: struct{}{},
				},
			},
		},
		{
			name: "truncate all",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20)},
					},
				},
				Messages: makeMessageObservation(
					map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
						1: {1, 10},
						2: {11, 20},
					},
				),
			},
			maxSize: 50, // less than what can fit a single commit report for single chain
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.observation.TokenData = makeNoopTokenDataObservations(tt.observation.Messages)
			tt.expected.TokenData = tt.observation.TokenData
			lggr := logger.Test(t)
			op := NewObservationOptimizer(lggr, tt.maxSize, ocrtypecodec.NewExecCodecJSON()) // todo: use proto
			obs, err := op.TruncateObservation(tt.observation)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.expected.CommitReports, obs.CommitReports)

			// Check only deleted messages were deleted
			for chain, seqNums := range obs.Messages {
				for seqNum := range seqNums {
					if _, ok := tt.deletedMsgs[chain]; ok {
						if _, ok2 := tt.deletedMsgs[chain][seqNum]; ok2 {
							require.True(t, obs.Messages[chain][seqNum].IsEmpty())
							continue
						}
					}
					require.False(t, obs.Messages[chain][seqNum].IsEmpty())
				}
			}
		})
	}
}

func makeMessageObservation(
	srcToSeqNumRange map[cciptypes.ChainSelector]cciptypes.SeqNumRange,
	opts ...msgOption) exectypes.MessageObservations {

	obs := make(exectypes.MessageObservations)
	for src, seqNumRng := range srcToSeqNumRange {
		obs[src] = make(map[cciptypes.SeqNum]cciptypes.Message)
		for i := seqNumRng.Start(); i <= seqNumRng.End(); i++ {
			msg := cciptypes.Message{
				Header: cciptypes.RampMessageHeader{
					SourceChainSelector: src,
					SequenceNumber:      i,
					MessageID:           rand.RandomBytes32(),
				},
				FeeValueJuels: cciptypes.NewBigIntFromInt64(100),
			}
			for _, opt := range opts {
				opt(&msg)
			}
			obs[src][i] = msg
		}
	}
	return obs
}

func makeNoopTokenDataObservations(msgs exectypes.MessageObservations) exectypes.TokenDataObservations {
	tokenData := make(exectypes.TokenDataObservations)
	for src, seqNumToMsg := range msgs {
		tokenData[src] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
		for seq := range seqNumToMsg {
			tokenData[src][seq] = exectypes.NewMessageTokenData()
		}
	}
	return tokenData

}

type msgOption func(*cciptypes.Message)

func withData(data []byte) msgOption {
	return func(m *cciptypes.Message) {
		m.Data = data
	}
}
