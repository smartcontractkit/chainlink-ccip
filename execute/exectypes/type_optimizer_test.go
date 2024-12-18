package exectypes

import (
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_truncateLastCommit(t *testing.T) {
	tests := []struct {
		name        string
		chain       cciptypes.ChainSelector
		observation Observation
		expected    Observation
	}{
		{
			name:  "no commits to truncate",
			chain: 1,
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					1: {},
				},
			},
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					1: {},
				},
			},
		},
		{
			name:  "truncate last commit",
			chain: 1,
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData{
					1: {
						1:  NewMessageTokenData(),
						11: NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x02}},
			},
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData{
					1: {
						1: NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := NewObservationOptimizer(10000)
			truncated, _ := op.truncateLastCommit(tt.observation, tt.chain)
			require.Equal(t, tt.expected, truncated)
		})
	}
}

func Test_truncateChain(t *testing.T) {
	tests := []struct {
		name        string
		chain       cciptypes.ChainSelector
		observation Observation
		expected    Observation
	}{
		{
			name:  "truncate chain data",
			chain: 1,
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData{
					1: {
						1: NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x01}},
			},
			expected: Observation{
				CommitReports:  map[cciptypes.ChainSelector][]CommitData{},
				Messages:       map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{},
				TokenData:      map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData{},
				CostlyMessages: []cciptypes.Bytes32{},
			},
		},
		{
			name:  "truncate non existent chain",
			chain: 2, // non existent chain
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData{
					1: {
						1: NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x01}},
			},
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData{
					1: {
						1: NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x01}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := NewObservationOptimizer(10000) // size not important here
			truncated := op.truncateChain(tt.observation, tt.chain)
			require.Equal(t, tt.expected, truncated)
		})
	}
}

func Test_truncateObservation(t *testing.T) {
	tests := []struct {
		name        string
		observation Observation
		maxSize     int
		expected    Observation
		deletedMsgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum]struct{}
		wantErr     bool
	}{
		{
			name: "no truncation needed",
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
			},
			maxSize: 1000,
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
			},
		},
		{
			name: "truncate last commit",
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
			},
		},
		{
			name: "truncate one message from first chain",
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
			observation: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{
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
			expected: Observation{
				CommitReports: map[cciptypes.ChainSelector][]CommitData{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.observation.TokenData = makeNoopTokenDataObservations(tt.observation.Messages)
			tt.expected.TokenData = tt.observation.TokenData
			op := NewObservationOptimizer(tt.maxSize)
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
	opts ...msgOption) MessageObservations {

	obs := make(MessageObservations)
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

func makeNoopTokenDataObservations(msgs MessageObservations) TokenDataObservations {
	tokenData := make(TokenDataObservations)
	for src, seqNumToMsg := range msgs {
		tokenData[src] = make(map[cciptypes.SeqNum]MessageTokenData)
		for seq := range seqNumToMsg {
			tokenData[src][seq] = NewMessageTokenData()
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
