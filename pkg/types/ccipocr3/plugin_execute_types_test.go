package ccipocr3

import (
	"fmt"
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
		data := append([]byte{1}, []byte(`{"Bogus":1}`)...)
		_, err := DecodeExecuteReportInfo(data)
		require.ErrorContains(t, err, "unknown field")
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
			name: "zero object",
			reportInfo: []MerkleRootChain{
				{},
			},
			//nolint:lll
			want:    append([]byte{1}, []byte(`[{"chain":0,"onRampAddress":"0x","seqNumsRange":[0,0],"merkleRoot":"0x0000000000000000000000000000000000000000000000000000000000000000"}]`)...),
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.reportInfo.Encode()
			fmt.Println(string(got))
			tt.wantErr(t, err, "Encode()")
			require.Equalf(t, tt.want, got, "Encode()")

			eri2, err := DecodeExecuteReportInfo(got)
			tt.wantErr(t, err, "Decode()")
			assert.Equalf(t, tt.reportInfo, eri2, "Decode()")
		})
	}
}
