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

func Test_verifyObservationSignature2(t *testing.T) {
	offchainPublicKey := "b1f20030f19ace75552a59125500d58d2d43305af537cb680da126e9a01d4100"
	expSig := "e6e7eaab2f66f8cabbf98ed301827992d096c795d036322146671b5acb18c54f8e7945f0847a0d5da5218a2cb4dd0c471f" +
		"67fb498bfdc4a8fc2afba157e5b709"

	// Parse offchain pub key
	offchainPublicKeyBytes, err := hex.DecodeString(offchainPublicKey)
	require.NoError(t, err)
	offchainPK := ed25519.PublicKey(offchainPublicKeyBytes)

	expSigBytes, err := hex.DecodeString(expSig)
	require.NoError(t, err)

	cfgDigest := "95efcbf7b2a42dc297d2a074e88046d21ac03fbfa1e7c87e4b873bcba76a1614"
	cfgDigestBytes, err := hex.DecodeString(cfgDigest)
	require.NoError(t, err)

	offRampAddress := "ab6e3971c5cb5490ec0762d6c4d09e9ea74ef9fe"
	offRampAddressBytes, err := hex.DecodeString(offRampAddress)
	require.NoError(t, err)

	onRampAddress := "986c36c599792220aa96e6179ead9e6176cc425c"
	onRampAddressBytes, err := hex.DecodeString(onRampAddress)
	require.NoError(t, err)

	rootStr := "36750744c3539dcf2d19f6385c0dd4c91d96111b7d170ac4e488484f81b9d64c"
	rootBytes, err := hex.DecodeString(rootStr)
	require.NoError(t, err)

	signedObs := &rmnpb.SignedObservation{
		Observation: &rmnpb.Observation{
			RmnHomeContractConfigDigest: cfgDigestBytes,
			LaneDest: &rmnpb.LaneDest{
				DestChainSelector: 11718334693212889894,
				OfframpAddress:    offRampAddressBytes,
			},
			FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
				{
					LaneSource: &rmnpb.LaneSource{
						SourceChainSelector: 3447725707992745283,
						OnrampAddress:       onRampAddressBytes,
					},
					ClosedInterval: &rmnpb.ClosedInterval{
						MinMsgNr: 10380311269462515579,
						MaxMsgNr: 16302188641602378505,
					},
					Root: rootBytes,
				},
			},
			Timestamp: 17436336888368921341,
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
