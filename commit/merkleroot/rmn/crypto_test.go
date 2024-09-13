package rmn

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

var (
	empty32ByteArr [32]byte
	empty20ByteArr [20]byte
)

func Test_verifyObservationSignature(t *testing.T) {
	offchainPublicKey := "0x137d557b609823f3f8f841265d5797e0fe97892b496764af326fd62a82d881aa"

	expSig := "edc308a23ce4aabf7518ad38778447e407b89b13011f7abc10e7641a51728e9ec6dd49a014ca238d8ba5d91fe9aff115d" +
		"00ec3f2e3b7a4391c7ca1cfda4bf40c"

	// Parse offchain pub key
	offchainPublicKeyBytes, err := hex.DecodeString(offchainPublicKey[2:])
	require.NoError(t, err)
	offchainPK := ed25519.PublicKey(offchainPublicKeyBytes)

	expSigBytes, err := hex.DecodeString(expSig)
	require.NoError(t, err)

	signedObs := &rmnpb.SignedObservation{
		Observation: &rmnpb.Observation{
			RmnHomeContractConfigDigest: empty32ByteArr[:],
			LaneDest:                    &rmnpb.LaneDest{DestChainSelector: 1, OfframpAddress: empty20ByteArr[:]},
			FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
				{
					LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 2, OnrampAddress: empty20ByteArr[:]},
					ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 0, MaxMsgNr: 0},
					Root:           empty32ByteArr[:],
				},
			},
			Timestamp: 0,
		},
		Signature: expSigBytes,
	}

	// Init rmn node
	rmnNode := RMNNodeInfo{
		ID:                        123,
		SignObservationsPublicKey: &offchainPK,
		SignObservationPrefix:     "chainlink ccip 1.6 rmn observation",
	}

	err = verifyObservationSignature(rmnNode, signedObs)
	assert.NoError(t, err)

	// After we update one byte in the signature, the signature verification should fail
	signedObs.Signature[0] = signedObs.Signature[0] + 1
	err = verifyObservationSignature(rmnNode, signedObs)
	assert.Error(t, err)
}
