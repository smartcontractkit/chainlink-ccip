package testhelpers

import (
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
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
