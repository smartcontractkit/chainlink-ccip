package reader

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_USDCMessageReader_New(t *testing.T) {

}

func Test_USDCMessageReader_MessageHashes(t *testing.T) {

}

func Test_MessageSentEvent_unpackID(t *testing.T) {
	nonEmptyEvent := eventID{}
	for i := 0; i < 32; i++ {
		nonEmptyEvent[i] = byte(i)
	}

	fullPayload := make([]byte, 0, 64)
	fullPayload = append(fullPayload, nonEmptyEvent[:]...)
	fullPayload = append(fullPayload, nonEmptyEvent[:]...)

	tests := []struct {
		name    string
		data    []byte
		want    eventID
		wantErr bool
	}{
		{
			name:    "event too short",
			data:    make([]byte, 31),
			wantErr: true,
		},
		{
			name: "event with proper length but empty",
			data: make([]byte, 32),
			want: eventID{},
		},
		{
			name: "event with proper length and data",
			data: fullPayload,
			want: nonEmptyEvent,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event := MessageSentEvent{Arg0: tt.data}
			got, err := event.unpackID()

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.Error(t, err)
			}
		})

	}
}

func Test_SourceTokenDataPayload_ToBytes(t *testing.T) {
	tests := []struct {
		nonce        uint64
		sourceDomain uint32
	}{
		{
			nonce:        0,
			sourceDomain: 0,
		},
		{
			nonce:        2137,
			sourceDomain: 4,
		},
		{
			nonce:        2,
			sourceDomain: 2137,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("nonce=%d,sourceDomain=%d", test.nonce, test.sourceDomain), func(t *testing.T) {
			payload1 := NewSourceTokenDataPayload(test.nonce, test.sourceDomain)
			bytes := payload1.ToBytes()

			payload2, err := NewSourceTokenDataPayloadFromBytes(bytes)
			require.NoError(t, err)
			require.Equal(t, *payload1, *payload2)
		})
	}
}

func Test_SourceTokenDataPayload_FromBytes(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		want    *SourceTokenDataPayload
		wantErr bool
	}{
		{
			name:    "too short data",
			data:    []byte{0x01, 0x02, 0x03},
			wantErr: true,
		},
		{
			name: "empty data but with proper length",
			data: make([]byte, 12),
			want: NewSourceTokenDataPayload(0, 0),
		},
		{
			name: "data with nonce and source domain",
			data: []byte{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 2},
			want: NewSourceTokenDataPayload(1, 2),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSourceTokenDataPayloadFromBytes(tt.data)

			if !tt.wantErr {
				require.NoError(t, err)
				require.Equal(t, tt.want, got)
			} else {
				require.Error(t, err)
			}
		})
	}
}
