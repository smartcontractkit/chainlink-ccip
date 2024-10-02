package discovery

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/commontypes"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// Ensure that PluginProcessor is implemented.
var _ plugincommon.PluginProcessor[dt.Query, dt.Observation, dt.Outcome] = &ContractDiscoveryProcessor{}

// ContractDiscoveryProcessor is a plugin processor for discovering contracts.
type ContractDiscoveryProcessor struct {
	lggr            logger.Logger
	reader          *reader.CCIPReader
	homechain       reader.HomeChain
	dest            cciptypes.ChainSelector
	fRoleDON        int
	oracleIDToP2PID map[commontypes.OracleID]ragep2ptypes.PeerID
}

func NewContractDiscoveryProcessor(
	lggr logger.Logger,
	reader *reader.CCIPReader,
	homechain reader.HomeChain,
	dest cciptypes.ChainSelector,
	fRoleDON int,
	oracleIDToP2PID map[commontypes.OracleID]ragep2ptypes.PeerID,
) *ContractDiscoveryProcessor {
	return &ContractDiscoveryProcessor{
		lggr:            lggr,
		reader:          reader,
		homechain:       homechain,
		dest:            dest,
		fRoleDON:        fRoleDON,
		oracleIDToP2PID: oracleIDToP2PID,
	}
}

// Query is not needed for this processor.
func (cdp *ContractDiscoveryProcessor) Query(_ context.Context, _ dt.Outcome) (dt.Query, error) {
	return nil, nil
}

// Observation reads contract addresses from the destination chain.
// In the future this should be extended to omit observations unless one of the nodes requests addresses.
func (cdp *ContractDiscoveryProcessor) Observation(
	ctx context.Context, _ dt.Outcome, _ dt.Query,
) (dt.Observation, error) {
	// all oracles should be able to read from the home chain, so we
	// can fetch f chain reliably.
	fChain, err := cdp.homechain.GetFChain()
	if err != nil {
		return dt.Observation{}, fmt.Errorf("unable to get fchain: %w", err)
	}

	// TODO: discover the full list of source chain selectors and pass it into DiscoverContracts.
	contracts, err := (*cdp.reader).DiscoverContracts(ctx, nil)
	if err != nil {
		return dt.Observation{}, fmt.Errorf("unable to discover contracts: %w", err)
	}

	return dt.Observation{
		FChain:    fChain,
		Addresses: contracts,
	}, nil
}

func (cdp *ContractDiscoveryProcessor) ValidateObservation(
	_ dt.Outcome, _ dt.Query, ao plugincommon.AttributedObservation[dt.Observation],
) error {
	oraclePeerID, ok := cdp.oracleIDToP2PID[ao.OracleID]
	if !ok {
		// should never happen in practice.
		return fmt.Errorf("no peer ID found for Oracle %d", ao.OracleID)
	}

	supportedChains, err := cdp.homechain.GetSupportedChainsForPeer(oraclePeerID)
	if err != nil {
		return fmt.Errorf("unable to get supported chains for Oracle %d: %w", ao.OracleID, err)
	}

	// check that the oracle is allowed to observe the dest chain.
	if !supportedChains.Contains(cdp.dest) {
		return fmt.Errorf("oracle %d is not allowed to observe chain %s", ao.OracleID, cdp.dest)
	}

	// NOTE: once oracles can also discover things on source chains, we must
	// check that they can read whatever source chain is used to determine the
	// address, e.g source fee quoter / router.

	return nil
}

// aggObs is used to store multiple observations for each value being observed.
type aggObs struct {
	fChain            map[cciptypes.ChainSelector][]int
	onrampAddrs       map[cciptypes.ChainSelector][][]byte
	feeQuoterAddrs    map[cciptypes.ChainSelector][][]byte
	nonceManagerAddrs map[cciptypes.ChainSelector][][]byte
	rmnRemoteAddrs    map[cciptypes.ChainSelector][][]byte
	routerAddrs       map[cciptypes.ChainSelector][][]byte
}

// aggregateObservations combines observations for multiple objects into aggObs, which is a convenient
// format for consensus.GetConsensusMap.
func aggregateObservations(
	lggr logger.Logger,
	dest cciptypes.ChainSelector,
	aos []plugincommon.AttributedObservation[dt.Observation],
) aggObs {
	obs := aggObs{
		fChain:            make(map[cciptypes.ChainSelector][]int),
		onrampAddrs:       make(map[cciptypes.ChainSelector][][]byte),
		feeQuoterAddrs:    make(map[cciptypes.ChainSelector][][]byte),
		nonceManagerAddrs: make(map[cciptypes.ChainSelector][][]byte),
		rmnRemoteAddrs:    make(map[cciptypes.ChainSelector][][]byte),
		routerAddrs:       make(map[cciptypes.ChainSelector][][]byte),
	}
	for _, ao := range aos {
		for chainSel, f := range ao.Observation.FChain {
			obs.fChain[chainSel] = append(obs.fChain[chainSel], f)
		}

		for chain, addr := range ao.Observation.Addresses[consts.ContractNameOnRamp] {
			// we don't want invalid observations to "poison" the consensus.
			if isZero(addr) {
				lggr.Warnf("skipping empty onramp address in observation from Oracle %d", ao.OracleID)
				continue
			}
			obs.onrampAddrs[chain] = append(obs.onrampAddrs[chain], addr)
		}

		if isZero(ao.Observation.Addresses[consts.ContractNameNonceManager][dest]) {
			lggr.Warnf("skipping empty nonce manager address in observation from Oracle %d", ao.OracleID)
		} else {
			obs.nonceManagerAddrs[dest] = append(
				obs.nonceManagerAddrs[dest],
				ao.Observation.Addresses[consts.ContractNameNonceManager][dest],
			)
		}

		if isZero(ao.Observation.Addresses[consts.ContractNameRMNRemote][dest]) {
			lggr.Warnf("skipping empty RMNRemote address in observation from Oracle %d", ao.OracleID)
		} else {
			obs.rmnRemoteAddrs[dest] = append(
				obs.rmnRemoteAddrs[dest],
				ao.Observation.Addresses[consts.ContractNameRMNRemote][dest],
			)
		}

		for chain, addr := range ao.Observation.Addresses[consts.ContractNameRouter] {
			if isZero(addr) {
				lggr.Warnf("skipping empty Router address in observation from Oracle %d", ao.OracleID)
				continue
			}
			obs.routerAddrs[chain] = append(
				obs.routerAddrs[chain],
				ao.Observation.Addresses[consts.ContractNameRouter][chain],
			)
		}

		for chain, addr := range ao.Observation.Addresses[consts.ContractNameFeeQuoter] {
			// we don't want invalid observations to "poison" the consensus.
			if isZero(addr) {
				lggr.Warnf("skipping empty fee quoter address in observation from Oracle %d", ao.OracleID)
				continue
			}
			obs.feeQuoterAddrs[chain] = append(obs.feeQuoterAddrs[chain], addr)
		}
	}
	return obs
}

// getOnRampDestRouterConsensus figures out a single dest router address from the multiple observed addresses from
// multiple offramp source chain configurations. This happens in two phases:
//  1. get consensus on the router address for each offRamp source chain config according the threshold (2*fChain+1).
//  2. with consensus across the source chains, select whichever address was observed the most times.
//     * if there is a tie, return an error.
//
// For a typical configuration, we expect all source chains to report the same destination router address.
// See https://github.com/smartcontractkit/chainlink/blob/05c8b45da51340aa9575ac909486decc9c164792/contracts/src/v0.8/ccip/offRamp/OffRamp.sol#L105
//
//nolint:lll // it's a url
func getOnRampDestRouterConsensus(
	lggr logger.Logger,
	routerAddrs map[cciptypes.ChainSelector][][]byte,
	thresh consensus.MultiThreshold[cciptypes.ChainSelector],
) ([]byte, error) {
	routerConsensus := consensus.GetConsensusMap(
		lggr,
		"router",
		routerAddrs,
		thresh,
	)
	// All the routers in routerAddrs are "valid", but we can only choose one.
	routerCounts := make(map[string]int)
	for _, addr := range routerConsensus {
		routerCounts[string(addr)]++
	}
	var tied bool
	var maxCount int
	var address []byte
	for addr, count := range routerCounts {
		if count > maxCount {
			maxCount = count
			tied = false
			address = []byte(addr)
		} else if count == maxCount {
			tied = true
		}
	}
	if tied {
		return nil, fmt.Errorf("no consensus on router, there is a tie multiple routers were seen %d times: %v, %v",
			maxCount, routerCounts, routerConsensus)
	}
	if maxCount == 0 {
		return nil, fmt.Errorf("no consensus on router, routerAddrs: %v", routerAddrs)
	}
	return address, nil
}

// isZero returns true if data is nil or all zeros, otherwise returns false.
func isZero(data []byte) bool {
	for _, v := range data {
		if v != 0 {
			return false
		}
	}
	return true
}

// Outcome comes to consensus on the contract addresses and updates the chainreader and update the
// CCIPReader. It doesn't actually return an Outcome.
func (cdp *ContractDiscoveryProcessor) Outcome(
	ctx context.Context, _ dt.Outcome, _ dt.Query, aos []plugincommon.AttributedObservation[dt.Observation],
) (dt.Outcome, error) {
	cdp.lggr.Infow("Processing contract discovery outcome", "observations", aos)

	agg := aggregateObservations(cdp.lggr, cdp.dest, aos)

	// fChain consensus - uses the role DON F value because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(cdp.fRoleDON))
	fChain := consensus.GetConsensusMap(cdp.lggr, "fChain", agg.fChain, donThresh)
	fChainThresh := consensus.MakeMultiThreshold(fChain, consensus.TwoFPlus1)
	destThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fChain[cdp.dest]))

	contracts := make(reader.ContractAddresses)
	onrampConsensus := consensus.GetConsensusMap(
		cdp.lggr,
		"onramp",
		agg.onrampAddrs,
		destThresh,
	)
	cdp.lggr.Infow("Determined consensus onramps",
		"onrampConsensus", onrampConsensus,
		"onrampAddrs", agg.onrampAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(onrampConsensus) == 0 {
		cdp.lggr.Warnw("No consensus on onramps, onrampConsensus map is empty")
	}
	contracts[consts.ContractNameOnRamp] = onrampConsensus

	nonceManagerConsensus := consensus.GetConsensusMap(
		cdp.lggr,
		"nonceManager",
		agg.nonceManagerAddrs,
		fChainThresh,
	)
	cdp.lggr.Infow("Determined consensus nonce manager",
		"nonceManagerConsensus", nonceManagerConsensus,
		"nonceManagerAddrs", agg.nonceManagerAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(nonceManagerConsensus) == 0 {
		cdp.lggr.Warnw("No consensus on nonce manager, nonceManagerConsensus map is empty")
	}
	contracts[consts.ContractNameNonceManager] = nonceManagerConsensus

	// RMNRemote address consensus
	rmnRemoteConsensus := consensus.GetConsensusMap(
		cdp.lggr,
		"rmnRemote",
		agg.rmnRemoteAddrs,
		fChainThresh,
	)
	cdp.lggr.Infow("Determined consensus RMNRemote",
		"rmnRemoteConsensus", rmnRemoteConsensus,
		"rmnRemoteAddrs", agg.rmnRemoteAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(rmnRemoteConsensus) == 0 {
		cdp.lggr.Warnw("No consensus on RMNRemote, rmnRemoteConsensus map is empty")
	}
	contracts[consts.ContractNameRMNRemote] = rmnRemoteConsensus

	feeQuoterConsensus := consensus.GetConsensusMap(
		cdp.lggr,
		"fee quoter",
		agg.feeQuoterAddrs,
		fChainThresh,
	)
	cdp.lggr.Infow("Determined consensus fee quoter",
		"feeQuoterConsensus", feeQuoterConsensus,
		"feeQuoterAddrs", agg.feeQuoterAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(feeQuoterConsensus) == 0 {
		cdp.lggr.Warnw("No consensus on fee quoters, feeQuoterConsensus map is empty")
	}
	contracts[consts.ContractNameFeeQuoter] = feeQuoterConsensus

	// Router address consensus
	contracts[consts.ContractNameRouter] = map[cciptypes.ChainSelector][]byte{}
	routerAddr, err := getOnRampDestRouterConsensus(cdp.lggr, agg.routerAddrs, fChainThresh)
	if err != nil {
		cdp.lggr.Warnw("Unable to reach consensus on router address",
			"RouterAddrs", agg.routerAddrs,
			"fChainThresh", fChainThresh,
			"err", err,
		)
	} else {
		cdp.lggr.Infow("Determined consensus Router",
			"RouterConsensus", routerAddr,
			"RouterAddrs", agg.routerAddrs,
			"fChainThresh", fChainThresh,
		)
		if len(routerAddr) == 0 {
			cdp.lggr.Warnw("No consensus on router, routerConsensus map is empty")
		} else {
			contracts[consts.ContractNameRouter][cdp.dest] = routerAddr
		}
	}

	// call Sync to bind contracts.
	if err := (*cdp.reader).Sync(ctx, contracts); err != nil {
		return dt.Outcome{}, fmt.Errorf("unable to sync contracts: %w", err)
	}

	return dt.Outcome{}, nil
}
