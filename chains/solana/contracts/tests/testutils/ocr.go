package testutils

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
)

type OcrPlugin uint8

func (p OcrPlugin) String() string {
	switch p {
	case OcrCommitPlugin:
		return "CommitPlugin"
	case OcrExecutePlugin:
		return "ExecutePlugin"
	}
	return "INVALID"
}

const (
	OcrCommitPlugin OcrPlugin = iota
	OcrExecutePlugin
)

func GenerateSignersAndTransmitters(t *testing.T, maxOracles int) ([]eth.Signer, []solana.PrivateKey, func() solana.PrivateKey) {
	signers, err := generateUnique(maxOracles, func() (eth.Signer, error) {
		ks, err := eth.GenerateEthPrivateKeys(1)
		if err != nil {
			return eth.Signer{}, err
		}
		return eth.GetSignerFromPk(ks[0])
	})
	require.NoError(t, err)
	transmitters, err := generateUnique(maxOracles, solana.NewRandomPrivateKey)
	require.NoError(t, err)
	getTransmitter := func() solana.PrivateKey {
		index, err := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(len(transmitters))))
		require.NoError(t, err)
		return transmitters[index.Int64()]
	}
	return signers, transmitters, getTransmitter
}

func generateUnique[T fmt.Stringer](count int, generator func() (T, error)) ([]T, error) {
	out := []T{}
	seen := map[string]bool{}

	for len(out) < count {
		v, err := generator()
		if err != nil {
			return nil, err
		}

		if _, exists := seen[v.String()]; !exists {
			out = append(out, v)
			seen[v.String()] = true
		}
	}
	return out, nil
}
