package testhelpers

import (
	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

func CreateRMNRemoteCfg() ccipocr3.RemoteConfig {
	return ccipocr3.RemoteConfig{
		ContractAddress: rand.RandomBytes(20),
		ConfigDigest:    rand.RandomBytes32(),
		Signers: []ccipocr3.RemoteSignerInfo{
			{OnchainPublicKey: append(rand.RandomBytes(20), byte(1)), NodeIndex: 1},
			{OnchainPublicKey: append(rand.RandomBytes(20), byte(2)), NodeIndex: 2},
		},
		FSign:            1,
		ConfigVersion:    rand.RandomUint32(),
		RmnReportVersion: rand.RandomReportVersion(),
	}
}
