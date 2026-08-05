package rmn

import (
	"crypto/ed25519"
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

var (
	empty32ByteArr [32]byte
	empty20ByteArr [20]byte
)

func Test_verifyObservationSignature(t *testing.T) {
	// NOTE: The following test case data are shared w/ RMN team.
	// RMN team has generated the public keys and the signatures.

	testCases := []struct {
		name              string
		offchainPublicKey string
		expSig            string
		signedObs         *rmnpb.SignedObservation
		cfgDigest         string
		offRampAddress    string
		onRampAddress     string
		rootStr           string
	}{
		{
			name:              "empty observation",
			offchainPublicKey: "137d557b609823f3f8f841265d5797e0fe97892b496764af326fd62a82d881aa",
			expSig: "edc308a23ce4aabf7518ad38778447e407b89b13011f7abc10e7641a51728e9ec6dd49a014ca238d8ba5d" +
				"91fe9aff115d00ec3f2e3b7a4391c7ca1cfda4bf40c",
			signedObs: &rmnpb.SignedObservation{
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
			},
		},
		{
			name:              "full observation",
			offchainPublicKey: "b1f20030f19ace75552a59125500d58d2d43305af537cb680da126e9a01d4100",
			expSig: "e6e7eaab2f66f8cabbf98ed301827992d096c795d036322146671b5acb18c54f8e7945f0847a0d5da5218a" +
				"2cb4dd0c471f67fb498bfdc4a8fc2afba157e5b709",
			signedObs: &rmnpb.SignedObservation{
				Observation: &rmnpb.Observation{
					LaneDest: &rmnpb.LaneDest{
						DestChainSelector: 11718334693212889894,
					},
					FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
						{
							LaneSource: &rmnpb.LaneSource{
								SourceChainSelector: 3447725707992745283,
							},
							ClosedInterval: &rmnpb.ClosedInterval{
								MinMsgNr: 10380311269462515579,
								MaxMsgNr: 16302188641602378505,
							},
						},
					},
					Timestamp: 17436336888368921341,
				},
			},
			cfgDigest:      "95efcbf7b2a42dc297d2a074e88046d21ac03fbfa1e7c87e4b873bcba76a1614",
			offRampAddress: "ab6e3971c5cb5490ec0762d6c4d09e9ea74ef9fe",
			onRampAddress:  "986c36c599792220aa96e6179ead9e6176cc425c",
			rootStr:        "36750744c3539dcf2d19f6385c0dd4c91d96111b7d170ac4e488484f81b9d64c",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			offchainPublicKeyBytes, err := hex.DecodeString(tc.offchainPublicKey)
			require.NoError(t, err)
			offchainPK := ed25519.PublicKey(offchainPublicKeyBytes)

			if tc.expSig != "" {
				tc.signedObs.Signature, _ = hex.DecodeString(tc.expSig)
			}
			if tc.cfgDigest != "" {
				tc.signedObs.Observation.RmnHomeContractConfigDigest, _ = hex.DecodeString(tc.cfgDigest)
			}
			if tc.offRampAddress != "" {
				tc.signedObs.Observation.LaneDest.OfframpAddress, _ = hex.DecodeString(tc.offRampAddress)
			}
			if tc.onRampAddress != "" {
				tc.signedObs.Observation.FixedDestLaneUpdates[0].LaneSource.OnrampAddress, _ = hex.DecodeString(tc.onRampAddress)
			}
			if tc.rootStr != "" {
				tc.signedObs.Observation.FixedDestLaneUpdates[0].Root, _ = hex.DecodeString(tc.rootStr)
			}

			rmnNode := rmntypes.HomeNodeInfo{
				OffchainPublicKey: &offchainPK,
			}
			signObservationPrefix1 := "chainlink ccip 1.6 rmn observation"

			err = verifyObservationSignature(rmnNode, signObservationPrefix1, tc.signedObs, NewED25519Verifier())
			assert.NoError(t, err)

			rmnNode2 := rmnNode
			signObservationPrefix2 := "chainlink ccip 1.6 rmn observation2----------"
			err = verifyObservationSignature(rmnNode2, signObservationPrefix2, tc.signedObs, NewED25519Verifier())
			assert.Error(t, err)

			// After we update one byte in the signature, the signature verification should fail
			tc.signedObs.Signature[0] = tc.signedObs.Signature[0] + 1
			err = verifyObservationSignature(rmnNode, signObservationPrefix1, tc.signedObs, NewED25519Verifier())
			assert.Error(t, err)
		})
	}
}
