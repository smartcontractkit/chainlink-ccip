package discovery

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/commontypes"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
	reporter plugincommon.MetricsReporter,
) plugincommon.PluginProcessor[dt.Query, dt.Observation, dt.Outcome] {
	p := &ContractDiscoveryProcessor{
		lggr:            lggr,
		reader:          reader,
		homechain:       homechain,
		dest:            dest,
		fRoleDON:        fRoleDON,
		oracleIDToP2PID: oracleIDToP2PID,
	}
	return plugincommon.NewTrackedProcessor(lggr, p, "discovery", reporter)
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
	seqNr := logutil.GetSeqNr(ctx)

	// all oracles should be able to read from the home chain, so we
	// can fetch f chain reliably.
	fChain, err := cdp.homechain.GetFChain()
	if err != nil {
		return dt.Observation{}, fmt.Errorf("unable to get fchain: %w, seqNr: %d", err, seqNr)
	}

	chainConfigs, err := cdp.homechain.GetAllChainConfigs()
	if err != nil {
		return dt.Observation{}, fmt.Errorf("unable to get chain configs: %w, seqNr: %d", err, seqNr)
	}

	contracts, err := (*cdp.reader).DiscoverContracts(ctx, maps.Keys(chainConfigs))
	if err != nil {
		return dt.Observation{}, fmt.Errorf("unable to discover contracts: %w, seqNr: %d", err, seqNr)
	}

	return dt.Observation{
		FChain:    fChain,
		Addresses: contracts,
	}, nil
}

func (cdp *ContractDiscoveryProcessor) ValidateObservation(
	_ dt.Outcome, _ dt.Query, ao plugincommon.AttributedObservation[dt.Observation],
) error {

	if ao.Observation.IsEmpty() {
		return nil
	}

	oraclePeerID, ok := cdp.oracleIDToP2PID[ao.OracleID]
	if !ok {
		// should never happen in practice.
		return fmt.Errorf("no peer ID found for Oracle %d", ao.OracleID)
	}

	supportedChains, err := cdp.homechain.GetSupportedChainsForPeer(oraclePeerID)
	if err != nil {
		return fmt.Errorf("unable to get supported chains for Oracle %d: %w", ao.OracleID, err)
	}

	err = plugincommon.ValidateFChain(ao.Observation.FChain)
	if err != nil {
		return fmt.Errorf("invalid FChain: %w", err)
	}

	for contract, addrs := range ao.Observation.Addresses {
		// some contract addresses come from the destination, others are from the source.
		switch contract {
		// discovered on the chain that the contract is deployed on.
		case consts.ContractNameFeeQuoter,
			consts.ContractNameRouter:
			for chain := range addrs {
				if !supportedChains.Contains(chain) {
					return fmt.Errorf(
						"oracle %d is not allowed to observe chain %s for %s", ao.OracleID, chain, contract)
				}
			}
		// discovered on the destination chain.
		case consts.ContractNameOnRamp,
			consts.ContractNameOffRamp,
			consts.ContractNameNonceManager,
			consts.ContractNameRMNRemote:

			if !supportedChains.Contains(cdp.dest) {
				return fmt.Errorf(
					"oracle %d is not allowed to observe contract (%s) on the destination chain %s",
					ao.OracleID, contract, cdp.dest)
			}
		default:
			return fmt.Errorf("unknown contract name %s", contract)
		}
	}

	return nil
}

// aggObs is used to store multiple observations for each value being observed.
type aggObs struct {
	fChain            map[cciptypes.ChainSelector][]int
	onrampAddrs       map[cciptypes.ChainSelector][]cciptypes.UnknownAddress
	feeQuoterAddrs    map[cciptypes.ChainSelector][]cciptypes.UnknownAddress
	nonceManagerAddrs map[cciptypes.ChainSelector][]cciptypes.UnknownAddress
	rmnRemoteAddrs    map[cciptypes.ChainSelector][]cciptypes.UnknownAddress
	routerAddrs       map[cciptypes.ChainSelector][]cciptypes.UnknownAddress
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
		onrampAddrs:       make(map[cciptypes.ChainSelector][]cciptypes.UnknownAddress),
		feeQuoterAddrs:    make(map[cciptypes.ChainSelector][]cciptypes.UnknownAddress),
		nonceManagerAddrs: make(map[cciptypes.ChainSelector][]cciptypes.UnknownAddress),
		rmnRemoteAddrs:    make(map[cciptypes.ChainSelector][]cciptypes.UnknownAddress),
		routerAddrs:       make(map[cciptypes.ChainSelector][]cciptypes.UnknownAddress),
	}
	for _, ao := range aos {
		for chainSel, f := range ao.Observation.FChain {
			obs.fChain[chainSel] = append(obs.fChain[chainSel], f)
		}

		for chain, addr := range ao.Observation.Addresses[consts.ContractNameOnRamp] {
			// we don't want invalid observations to "poison" the consensus.
			if isZero(addr) {
				lggr.Debugf("skipping empty onramp address in observation from Oracle %d", ao.OracleID)
				continue
			}
			obs.onrampAddrs[chain] = append(obs.onrampAddrs[chain], addr)
		}

		if isZero(ao.Observation.Addresses[consts.ContractNameNonceManager][dest]) {
			lggr.Debugf("skipping empty nonce manager address in observation from Oracle %d", ao.OracleID)
		} else {
			obs.nonceManagerAddrs[dest] = append(
				obs.nonceManagerAddrs[dest],
				ao.Observation.Addresses[consts.ContractNameNonceManager][dest],
			)
		}

		if isZero(ao.Observation.Addresses[consts.ContractNameRMNRemote][dest]) {
			lggr.Debugf("skipping empty RMNRemote address in observation from Oracle %d", ao.OracleID)
		} else {
			obs.rmnRemoteAddrs[dest] = append(
				obs.rmnRemoteAddrs[dest],
				ao.Observation.Addresses[consts.ContractNameRMNRemote][dest],
			)
		}

		for chain, addr := range ao.Observation.Addresses[consts.ContractNameRouter] {
			if isZero(addr) {
				lggr.Debugf("skipping empty Router address in observation from Oracle %d", ao.OracleID)
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
				lggr.Debugf("skipping empty fee quoter address in observation from Oracle %d", ao.OracleID)
				continue
			}
			obs.feeQuoterAddrs[chain] = append(obs.feeQuoterAddrs[chain], addr)
		}
	}
	return obs
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
	lggr := logutil.WithContextValues(ctx, cdp.lggr)
	lggr.Infow("Processing contract discovery outcome")
	contracts := make(reader.ContractAddresses)

	agg := aggregateObservations(lggr, cdp.dest, aos)

	// fChain consensus - uses the role DON F value because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(cdp.fRoleDON))
	fChain := consensus.GetConsensusMap(lggr, "fChain", agg.fChain, donThresh)
	fChainThresh := consensus.MakeMultiThreshold(fChain, consensus.TwoFPlus1)

	// We read onramp addresses from destChain offramp configs
	if _, exists := fChain[cdp.dest]; !exists {
		lggr.Warnf("missing fChain for dest (fChain[%d]), skipping onramp address lookup", cdp.dest)
	} else {
		destThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fChain[cdp.dest]))

		onrampConsensus := consensus.GetConsensusMap(
			lggr,
			"onramp",
			agg.onrampAddrs,
			destThresh,
		)
		lggr.Infow("Determined consensus onramps",
			"onrampConsensus", onrampConsensus,
			"onrampAddrs", agg.onrampAddrs,
			"fChainThresh", fChainThresh,
		)
		if len(onrampConsensus) == 0 {
			lggr.Warnw("No consensus on onramps, onrampConsensus map is empty")
		}
		contracts[consts.ContractNameOnRamp] = onrampConsensus
	}

	nonceManagerConsensus := consensus.GetConsensusMap(
		lggr,
		"nonceManager",
		agg.nonceManagerAddrs,
		fChainThresh,
	)
	lggr.Infow("Determined consensus nonce manager",
		"nonceManagerConsensus", nonceManagerConsensus,
		"nonceManagerAddrs", agg.nonceManagerAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(nonceManagerConsensus) == 0 {
		lggr.Warnw("No consensus on nonce manager, nonceManagerConsensus map is empty")
	}
	contracts[consts.ContractNameNonceManager] = nonceManagerConsensus

	// RMNRemote address consensus
	rmnRemoteConsensus := consensus.GetConsensusMap(
		lggr,
		"rmnRemote",
		agg.rmnRemoteAddrs,
		fChainThresh,
	)
	lggr.Infow("Determined consensus RMNRemote",
		"rmnRemoteConsensus", rmnRemoteConsensus,
		"rmnRemoteAddrs", agg.rmnRemoteAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(rmnRemoteConsensus) == 0 {
		lggr.Warnw("No consensus on RMNRemote, rmnRemoteConsensus map is empty")
	}
	contracts[consts.ContractNameRMNRemote] = rmnRemoteConsensus

	feeQuoterConsensus := consensus.GetConsensusMap(
		lggr,
		"fee quoter",
		agg.feeQuoterAddrs,
		fChainThresh,
	)
	lggr.Infow("Determined consensus fee quoter",
		"feeQuoterConsensus", feeQuoterConsensus,
		"feeQuoterAddrs", agg.feeQuoterAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(feeQuoterConsensus) == 0 {
		lggr.Warnw("No consensus on fee quoters, feeQuoterConsensus map is empty")
	}
	contracts[consts.ContractNameFeeQuoter] = feeQuoterConsensus

	// Router address consensus
	routerConsensus := consensus.GetConsensusMap(
		lggr,
		"router",
		agg.routerAddrs,
		fChainThresh,
	)
	lggr.Infow("Determined consensus router",
		"routerConsensus", routerConsensus,
		"routerAddrs", agg.routerAddrs,
		"fChainThresh", fChainThresh,
	)
	if len(routerConsensus) == 0 {
		lggr.Warnw("No consensus on router, routerConsensus map is empty")
	}
	contracts[consts.ContractNameRouter] = routerConsensus

	// call Sync to bind contracts.
	// NOTE: since Sync may make network calls, it could potentially fail and we don't want to
	// fail the entire outcome because of that. The reason being is that if this node is a leader
	// of an OCR round, it will NOT be able to complete the round due to failing to compute the Outcome.
	// TODO: we should move Sync calls to observation but that requires updates to the Outcome struct for discovery.
	if err := (*cdp.reader).Sync(ctx, contracts); err != nil {
		lggr.Errorw(
			"unable to sync contracts - this is usually due to RPC issues,"+
				" please check your RPC endpoints and their health!",
			"err", err)
	}

	return dt.Outcome{}, nil
}

func (cdp *ContractDiscoveryProcessor) Close() error {
	return nil
}
