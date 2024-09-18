package internal

import (
	"crypto/rand"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"
)

func MessageWithTokens(t *testing.T, tokenPoolAddr ...string) cciptypes.Message {
	onRampTokens := make([]cciptypes.RampTokenAmount, len(tokenPoolAddr))
	for i, addr := range tokenPoolAddr {
		b, err := cciptypes.NewBytesFromString(addr)
		require.NoError(t, err)
		onRampTokens[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: b,
			Amount:            cciptypes.NewBigIntFromInt64(int64(i + 1)),
		}
	}
	return cciptypes.Message{
		TokenAmounts: onRampTokens,
	}
}

func RandBytes() cciptypes.Bytes {
	var array [32]byte
	_, err := rand.Read(array[:])
	if err != nil {
		panic(err)
	}
	return array[:]
}
