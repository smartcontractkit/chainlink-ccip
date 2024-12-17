package ccipocr3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeExecuteReportInfo(t *testing.T) {
	// unknown version
	{
		data := append([]byte{2}, []byte("{}")...)
		_, err := DecodeExecuteReportInfo(data)
		require.ErrorContains(t, err, "unknown execute report info version (2)")
	}

	// unknown field
	{
		data := append([]byte{1}, []byte(`[{"Bogus":1}]`)...)
		_, err := DecodeExecuteReportInfo(data)
		require.ErrorContains(t, err, "unknown field")
	}

	// Not a slice
	{
		data := append([]byte{1}, []byte(`{"Bogus":1}`)...)
		_, err := DecodeExecuteReportInfo(data)
		require.ErrorContains(t, err, "object") // not super helpful...
	}
}

func TestExecuteReportInfo_EncodeDecode(t *testing.T) {
	tests := []struct {
		name       string
		reportInfo ExecuteReportInfo
		want       []byte
		wantErr    require.ErrorAssertionFunc
	}{
		{
			name: "object",
			reportInfo: []MerkleRootChain{
				{
					ChainSel:      10,
					OnRampAddress: mustNewUnknownAddress(t, "0x04D4cC5972ad487F71b85654d48b27D32b13a22F"),
					SeqNumsRange:  NewSeqNumRange(100, 200),
					MerkleRoot:    Bytes32{},
				},
			},
			//nolint:lll
			want:    append([]byte{1}, []byte(`[{"chain":10,"onRampAddress":"0x04d4cc5972ad487f71b85654d48b27d32b13a22f","seqNumsRange":[100,200],"merkleRoot":"0x0000000000000000000000000000000000000000000000000000000000000000"}]`)...),
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.reportInfo.Encode()
			tt.wantErr(t, err, "Encode()")
			require.Equalf(t, tt.want, got, "Encode()")

			eri2, err := DecodeExecuteReportInfo(got)
			tt.wantErr(t, err, "Decode()")
			assert.Equalf(t, tt.reportInfo, eri2, "Decode()")
		})
	}
}
