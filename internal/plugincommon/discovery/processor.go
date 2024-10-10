package discovery

import (
	"context"
	"errors"
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
		if errors.Is(err, reader.ErrContractReaderNotFound) {
			// Not a dest reader, only fChain observation will be made.
			// Processor is not disabled because the outcome phase will bind observed contracts.
			return dt.Observation{
				FChain: fChain,
			}, nil
		}

		// otherwise a legitimate error occurred when discovering.
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

type aggObs struct {
	onrampAddrs       map[cciptypes.ChainSelector][][]byte
	feeQuoterAddrs    map[cciptypes.ChainSelector][][]byte
	nonceManagerAddrs map[cciptypes.ChainSelector][][]byte
	rmnRemoteAddrs    map[cciptypes.ChainSelector][][]byte
	routerAddrs       map[cciptypes.ChainSelector][][]byte
}

func aggregateObservations(
	lggr logger.Logger,
	dest cciptypes.ChainSelector,
	aos []plugincommon.AttributedObservation[dt.Observation],
) aggObs {
	obs := aggObs{
		onrampAddrs:       make(map[cciptypes.ChainSelector][][]byte),
		feeQuoterAddrs:    make(map[cciptypes.ChainSelector][][]byte),
		nonceManagerAddrs: make(map[cciptypes.ChainSelector][][]byte),
		rmnRemoteAddrs:    make(map[cciptypes.ChainSelector][][]byte),
		routerAddrs:       make(map[cciptypes.ChainSelector][][]byte),
	}
	for _, ao := range aos {
		for chain, addr := range ao.Observation.Addresses[consts.ContractNameOnRamp] {
			// we don't want invalid observations to "poison" the consensus.
			if len(addr) == 0 {
				lggr.Warnf("skipping empty onramp address in observation from Oracle %d", ao.OracleID)
				continue
			}
			obs.onrampAddrs[chain] = append(obs.onrampAddrs[chain], addr)
		}

		if len(ao.Observation.Addresses[consts.ContractNameNonceManager][dest]) == 0 {
			lggr.Warnf("skipping empty nonce manager address in observation from Oracle %d", ao.OracleID)
		} else {
			obs.nonceManagerAddrs[dest] = append(
				obs.nonceManagerAddrs[dest],
				ao.Observation.Addresses[consts.ContractNameNonceManager][dest],
			)
		}

		if len(ao.Observation.Addresses[consts.ContractNameRMNRemote][dest]) == 0 {
			lggr.Warnf("skipping empty RMNRemote address in observation from Oracle %d", ao.OracleID)
		} else {
			obs.rmnRemoteAddrs[dest] = append(
				obs.rmnRemoteAddrs[dest],
				ao.Observation.Addresses[consts.ContractNameRMNRemote][dest],
			)
		}

		if len(ao.Observation.Addresses[consts.ContractNameRouter][dest]) == 0 {
			lggr.Warnf("skipping empty Router address in observation from Oracle %d", ao.OracleID)
		} else {
			obs.routerAddrs[dest] = append(
				obs.routerAddrs[dest],
				ao.Observation.Addresses[consts.ContractNameRouter][dest],
			)
		}

		for chain, addr := range ao.Observation.Addresses[consts.ContractNameFeeQuoter] {
			// we don't want invalid observations to "poison" the consensus.
			if len(addr) == 0 {
				lggr.Warnf("skipping empty fee quoter address in observation from Oracle %d", ao.OracleID)
				continue
			}
			obs.feeQuoterAddrs[chain] = append(obs.feeQuoterAddrs[chain], addr)
		}
	}
	return obs
}

// Outcome comes to consensus on the contract addresses and updates the chainreader. It doesn't actually
// return an Outcome.
func (cdp *ContractDiscoveryProcessor) Outcome(
	ctx context.Context, _ dt.Outcome, _ dt.Query, aos []plugincommon.AttributedObservation[dt.Observation],
) (dt.Outcome, error) {
	cdp.lggr.Infow("Processing contract discovery outcome", "observations", aos)
	// come to consensus on the onramp addresses and update the chainreader.

	// fChain consensus - uses the role DON F value because all nodes can observe the home chain.
	fChainObs := make(map[cciptypes.ChainSelector][]int)
	for _, ao := range aos {
		for chainSel, f := range ao.Observation.FChain {
			fChainObs[chainSel] = append(fChainObs[chainSel], f)
		}
	}
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(cdp.fRoleDON))
	fChain := consensus.GetConsensusMap(cdp.lggr, "fChain", fChainObs, donThresh)
	fChainThresh := consensus.MakeMultiThreshold(fChain, consensus.TwoFPlus1)
	destThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fChain[cdp.dest]))

	agg := aggregateObservations(cdp.lggr, cdp.dest, aos)

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

	// Router address consensus
	routerConsensus := consensus.GetConsensusMap(
		cdp.lggr,
		"router",
		agg.routerAddrs,
		fChainThresh,
	)
	cdp.lggr.Infow("Determined consensus Router",
		"RouterConsensus", routerConsensus,
		"RouterAddrs", agg.routerAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(routerConsensus) == 0 {
		cdp.lggr.Warnw("No consensus on router, routerConsensus map is empty")
	}
	contracts[consts.ContractNameRouter] = routerConsensus

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

	// call Sync to bind contracts.
	if err := (*cdp.reader).Sync(ctx, contracts); err != nil {
		return dt.Outcome{}, fmt.Errorf("unable to sync contracts: %w", err)
	}

	return dt.Outcome{}, nil
}
