package testhelpers

import (
	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
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

func CreateRMNRemoteCfg() rmntypes.RemoteConfig {
	return rmntypes.RemoteConfig{
		ContractAddress: rand.RandomBytes(20),
		ConfigDigest:    rand.RandomBytes32(),
		Signers: []rmntypes.RemoteSignerInfo{
			{
				OnchainPublicKey: rand.RandomBytes(20),
				NodeIndex:        rand.RandomUint64(),
			},
		},
		MinSigners:       rand.RandomUint64(),
		ConfigVersion:    rand.RandomUint32(),
		RmnReportVersion: rand.RandomReportVersion(),
	}
}

// ReportVersion is the version of the report that the RMNRemote contract supports.
// todo: fetch it from rmnRemoteCfg when the contract supports it and delete this constant.
// currently rmnRemoteCfg.RmnReportVersion is the hash of the report version string.
const ReportVersion string = "RMN_V1_6_ANY2EVM_REPORT"
