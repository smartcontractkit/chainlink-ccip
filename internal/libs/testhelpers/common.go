package testhelpers

import (
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
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
		ContractAddress: rand.RandomBytes(20),
		ConfigDigest:    rand.RandomBytes32(),
		Signers: []rmntypes.RemoteSignerInfo{
			{
				OnchainPublicKey:      rand.RandomBytes(20),
				NodeIndex:             rand.RandomUint64(),
				SignObservationPrefix: rand.RandomPrefix(),
			},
		},
		MinSigners:       rand.RandomUint64(),
		ConfigVersion:    rand.RandomUint32(),
		RmnReportVersion: rand.RandomReportVersion(),
	}
}
