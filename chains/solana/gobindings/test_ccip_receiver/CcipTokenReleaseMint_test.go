// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package test_ccip_receiver

import (
	"bytes"
	ag_gofuzz "github.com/gagliardetto/gofuzz"
	ag_require "github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestEncodeDecode_CcipTokenReleaseMint(t *testing.T) {
	fu := ag_gofuzz.New().NilChance(0)
	for i := 0; i < 1; i++ {
		t.Run("CcipTokenReleaseMint"+strconv.Itoa(i), func(t *testing.T) {
			{
				params := new(CcipTokenReleaseMint)
				fu.Fuzz(params)
				params.AccountMetaSlice = nil
				buf := new(bytes.Buffer)
				err := encodeT(*params, buf)
				ag_require.NoError(t, err)
				got := new(CcipTokenReleaseMint)
				err = decodeT(got, buf.Bytes())
				got.AccountMetaSlice = nil
				ag_require.NoError(t, err)
				ag_require.Equal(t, params, got)
			}
		})
	}
}
