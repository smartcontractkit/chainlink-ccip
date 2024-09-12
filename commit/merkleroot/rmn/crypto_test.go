package rmn

import (
	"crypto/ed25519"
	"encoding/hex"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func Test_VerifyRmnReportSignatures(t *testing.T) {
	onchainRmnRemoteAddr := "0x7821bcd6944457d17c631157efeb0c621baa76eb"

	// Init rmn node
	rmnNode := RMNNodeInfo{
		ID:                    123,
		SignReportsAddress:    common.HexToAddress(onchainRmnRemoteAddr),
		SignObservationPrefix: "chainlink ccip 1.6 rmn observation",
	}

	rmnHomeContractConfigDigestHex := "0x785936570d1c7422ef30b7da5555ad2f175fa2dd97a2429a2e71d1e07c94e060"
	rmnHomeContractConfigDigest := common.FromHex(rmnHomeContractConfigDigestHex)
	require.Len(t, rmnHomeContractConfigDigest, 32)
	var rmnHomeContractConfigDigest32 [32]byte
	copy(rmnHomeContractConfigDigest32[:], rmnHomeContractConfigDigest)

	rootHex := "0x48e688aefc20a04fdec6b8ff19df358fd532455659dcf529797cda358e9e5205"
	root := common.FromHex(rootHex)
	require.Len(t, root, 32)
	var root32 [32]byte
	copy(root32[:], root)

	onRampAddr := common.HexToAddress("0x6662cb20464f4be557262693bea0409f068397ed")
	abiDefinition := `[{"name": "", "type": "address"}]`
	onRampAddrAbi, err := abiEncode(abiDefinition, onRampAddr)
	if err != nil {
		log.Fatalf("Failed to ABI encode: %v", err)
	}

	destChainEvmID := uint64(4083663998511321420)
	reportData := ReportData{
		DestChainEvmID:              big.NewInt(0).SetUint64(destChainEvmID),
		DestChainSelector:           5266174733271469989,
		RmnRemoteContractAddress:    "0x3d015cec4411357eff4ea5f009a581cc519f75d3",
		OfframpAddress:              "0xc5cdb7711a478058023373b8ae9e7421925140f8",
		RmnHomeContractConfigDigest: rmnHomeContractConfigDigest32,
		LaneUpdates: NewLaneUpdatesFromPBType([]*rmnpb.FixedDestLaneUpdate{
			{
				LaneSource: &rmnpb.LaneSource{
					SourceChainSelector: 8258882951688608272,
					OnrampAddress:       onRampAddrAbi,
				},
				ClosedInterval: &rmnpb.ClosedInterval{
					MinMsgNr: 9018980618932210108,
					MaxMsgNr: 8239368306600774074,
				},
				Root: root32[:],
			},
		}),
	}

	rmnNodeSig := rmnpb.EcdsaSignature{
		R: common.FromHex("0x89546b4652d0377062a398e413344e4da6034ae877c437d0efe0e5246b70a9a1"),
		S: common.FromHex("0x95eef2d24d856ccac3886db8f4aebea60684ed73942392692908fed79a679b4e"),
	}

	err = VerifyRmnReportSignatures(
		reportData,
		[]*rmnpb.EcdsaSignature{&rmnNodeSig},
		[]RMNNodeInfo{rmnNode},
	)
	assert.NoError(t, err)
}
