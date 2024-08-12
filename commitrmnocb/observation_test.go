package commitrmnocb

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func Test_computeMerkleRoot(t *testing.T) {
	testCases := []struct {
		name           string
		messageHeaders []cciptypes.RampMessageHeader
		messageHasher  cciptypes.MessageHasher
		expMerkleRoot  string
		expErr         bool
	}{
		{
			name: "Happy path",
			messageHeaders: []cciptypes.RampMessageHeader{
				{
					MessageID:      mustNewMessageID("0x1a"),
					SequenceNumber: 112,
				},
				{
					MessageID:      mustNewMessageID("0x23"),
					SequenceNumber: 113,
				},
				{
					MessageID:      mustNewMessageID("0x87"),
					SequenceNumber: 114,
				}},
			messageHasher: NewMessageHasher(),
			expMerkleRoot: "94c7e711e6f2acf41dca598ced55b6925e55aaed83520dc5ea6cbc054344564b",
			expErr:        false,
		},
		{
			name: "Sequence number gap",
			messageHeaders: []cciptypes.RampMessageHeader{
				{
					MessageID:      mustNewMessageID("0x10"),
					SequenceNumber: 34,
				},
				{
					MessageID:      mustNewMessageID("0x12"),
					SequenceNumber: 36,
				}},
			messageHasher: NewMessageHasher(),
			expMerkleRoot: "",
			expErr:        true,
		},
		{
			name:           "Empty messages",
			messageHeaders: []cciptypes.RampMessageHeader{},
			messageHasher:  NewMessageHasher(),
			expMerkleRoot:  "",
			expErr:         true,
		},
		{
			name: "Bad hasher",
			messageHeaders: []cciptypes.RampMessageHeader{
				{
					MessageID:      mustNewMessageID("0x1a"),
					SequenceNumber: 112,
				},
				{
					MessageID:      mustNewMessageID("0x23"),
					SequenceNumber: 113,
				},
				{
					MessageID:      mustNewMessageID("0x87"),
					SequenceNumber: 114,
				}},
			messageHasher: NewBadMessageHasher(),
			expMerkleRoot: "94c7e711e6f2acf41dca598ced55b6925e55aaed83520dc5ea6cbc054344564b",
			expErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := Plugin{
				lggr:      logger.Test(t),
				msgHasher: tc.messageHasher,
			}

			msgs := make([]cciptypes.Message, 0)
			for _, h := range tc.messageHeaders {
				msgs = append(msgs, cciptypes.Message{Header: h})
			}

			rootBytes, err := p.computeMerkleRoot(context.Background(), msgs)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			rootString := hex.EncodeToString(rootBytes[:])
			assert.Equal(t, tc.expMerkleRoot, rootString)
		})
	}
}

func mustNewMessageID(msgIDHex string) cciptypes.Bytes32 {
	msgID, err := cciptypes.NewBytes32FromString(msgIDHex)
	if err != nil {
		panic(err)
	}
	return msgID
}

type MessageHasher struct{}

func NewMessageHasher() *MessageHasher {
	return &MessageHasher{}
}

func (m *MessageHasher) Hash(ctx context.Context, msg cciptypes.Message) (cciptypes.Bytes32, error) {
	// simply return the msg id as bytes32
	var b32 [32]byte
	copy(b32[:], msg.Header.MessageID[:])
	return b32, nil
}

type BadMessageHasher struct{}

func NewBadMessageHasher() *BadMessageHasher {
	return &BadMessageHasher{}
}

// Always returns an error
func (m *BadMessageHasher) Hash(ctx context.Context, msg cciptypes.Message) (cciptypes.Bytes32, error) {
	return cciptypes.Bytes32{}, fmt.Errorf("failed to hash")
}
