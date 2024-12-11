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
	type fields struct {
		Roots []Bytes32
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr require.ErrorAssertionFunc
	}{
		{
			name: "nil roots",
			fields: fields{
				Roots: nil,
			},
			want:    append([]byte{1}, []byte(`{"Roots":null}`)...),
			wantErr: require.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eri := ExecuteReportInfo{
				Roots: tt.fields.Roots,
			}
			got, err := eri.Encode()
			tt.wantErr(t, err, fmt.Sprintf("Encode()"))
			require.Equalf(t, tt.want, got, "Encode()")

			eri2, err := DecodeExecuteReportInfo(got)
			tt.wantErr(t, err, fmt.Sprintf("Decode()"))
			assert.Equalf(t, eri, eri2, "Decode()")
		})
	}
}
