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

}

func Test_SourceTokenDataPayload_ToBytes(t *testing.T) {
	tt := []struct {
		nonce        uint64
		sourceDomain uint32
	}{
		{
			nonce:        0,
			sourceDomain: 0,
		},
		{
			nonce:        1978987,
			sourceDomain: 1,
		},
		{
			nonce:        0,
			sourceDomain: 2123,
		},
	}

	for _, test := range tt {
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
