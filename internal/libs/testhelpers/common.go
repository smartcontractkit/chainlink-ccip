package testhelpers

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

func SetupConfigInfo(chainSelector ccipocr3.ChainSelector,
	readers []libocrtypes.PeerID,
	fChain uint8,
	cfg []byte) reader.ChainConfigInfo {
	return reader.ChainConfigInfo{
		ChainSelector: chainSelector,
		ChainConfig: reader.HomeChainConfigMapper{
			Readers: readers,
			FChain:  fChain,
			Config:  cfg,
		},
	}
}

func CreateOracleIDToP2pID(ids ...int) map[commontypes.OracleID]libocrtypes.PeerID {
	res := make(map[commontypes.OracleID]libocrtypes.PeerID)
	for _, id := range ids {
		res[commontypes.OracleID(id)] = libocrtypes.PeerID{byte(id)}
	}
	return res
}

func CreateRMNRemoteCfg() rmntypes.RemoteConfig {
	return rmntypes.RemoteConfig{
		ContractAddress: randomBytes(20),
		ConfigDigest:    randomBytes32(),
		Signers: []rmntypes.RemoteSignerInfo{
			{
				OnchainPublicKey:      randomBytes(20),
				NodeIndex:             randomUint64(),
				SignObservationPrefix: randomPrefix(),
			},
		},
		MinSigners:       randomUint64(),
		ConfigVersion:    randomUint32(),
		RmnReportVersion: randomReportVersion(),
	}
}

func randomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

func randomBytes32() cciptypes.Bytes32 {
	var b cciptypes.Bytes32
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	return b
}

func randomUint32() uint32 {
	n, err := rand.Int(rand.Reader, big.NewInt(256))
	if err != nil {
		panic(err)
	}
	return uint32(n.Int64())
}

func randomUint64() uint64 {
	n, err := rand.Int(rand.Reader, new(big.Int).SetUint64(^uint64(0)))
	if err != nil {
		panic(err)
	}
	return n.Uint64()
}

func randomPrefix() string {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func randomReportVersion() cciptypes.Bytes32 {
	versions := []cciptypes.Bytes32{
		stringToBytes32("RMN_V1_6_ANY2EVM_REPORT"),
		stringToBytes32("RMN_V2_0_ANY2EVM_REPORT"),
		stringToBytes32("RMN_V1_6_EVM2EVM_REPORT"),
	}
	return versions[randomUint32()%uint32(len(versions))]
}

func stringToBytes32(s string) cciptypes.Bytes32 {
	var result cciptypes.Bytes32
	copy(result[:], s)
	return result
}
