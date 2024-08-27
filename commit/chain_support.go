package commit

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
)

// ChainSupport contains functions that enable an oracle to determine which chains are accessible by itself and
// other oracles
type ChainSupport interface {
	// SupportedChains returns the set of chains that the given Oracle is configured to access
	SupportedChains(oracleID commontypes.OracleID) (mapset.Set[cciptypes.ChainSelector], error)

	// SupportsDestChain returns true if the given oracle supports the dest chain, returns false otherwise
	SupportsDestChain(oracle commontypes.OracleID) (bool, error)

	// KnownSourceChainsSlice returns a list of all known source chains
	KnownSourceChainsSlice() ([]cciptypes.ChainSelector, error)
}

type CCIPChainSupport struct {
	lggr            logger.Logger
	homeChain       reader.HomeChain
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID
	nodeID          commontypes.OracleID
	destChain       cciptypes.ChainSelector
}

func (c CCIPChainSupport) KnownSourceChainsSlice() ([]cciptypes.ChainSelector, error) {
	knownSourceChains, err := c.homeChain.GetKnownCCIPChains()
	if err != nil {
		c.lggr.Errorw("error getting known chains", "err", err)
		return nil, fmt.Errorf("error getting known chains: %w", err)
	}
	knownSourceChainsSlice := knownSourceChains.ToSlice()
	return slicelib.Filter(knownSourceChainsSlice, func(ch cciptypes.ChainSelector) bool { return ch != c.destChain }), nil
}

// SupportedChains returns the set of chains that the given Oracle is configured to access
func (c CCIPChainSupport) SupportedChains(oracleID commontypes.OracleID) (mapset.Set[cciptypes.ChainSelector], error) {
	p2pID, exists := c.oracleIDToP2pID[oracleID]
	if !exists {
		return nil, fmt.Errorf("oracle ID %d not found in oracleIDToP2pID", c.nodeID)
	}
	supportedChains, err := c.homeChain.GetSupportedChainsForPeer(p2pID)
	if err != nil {
		c.lggr.Warnw("error getting supported chains", err)
		return mapset.NewSet[cciptypes.ChainSelector](), fmt.Errorf("error getting supported chains: %w", err)
	}

	return supportedChains, nil
}

// SupportsDestChain returns true if the given oracle supports the dest chain, returns false otherwise
func (c CCIPChainSupport) SupportsDestChain(oracle commontypes.OracleID) (bool, error) {
	destChainConfig, err := c.homeChain.GetChainConfig(c.destChain)
	if err != nil {
		return false, fmt.Errorf("get chain config: %w", err)
	}
	return destChainConfig.SupportedNodes.Contains(c.oracleIDToP2pID[oracle]), nil
}

// Interface compliance check
var _ ChainSupport = (*CCIPChainSupport)(nil)
