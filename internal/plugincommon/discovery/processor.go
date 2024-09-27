package discovery

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

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

	contracts, err := (*cdp.reader).DiscoverContracts(ctx, cdp.dest)
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
		FChain:           fChain,
		OnRamp:           contracts[consts.ContractNameOnRamp],
		DestNonceManager: contracts[consts.ContractNameNonceManager][cdp.dest],
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

// Outcome comes to consensus on the contract addresses and updates the chainreader. It doesn't actually
// return an Outcome.
func (cdp *ContractDiscoveryProcessor) Outcome(
	_ dt.Outcome, _ dt.Query, aos []plugincommon.AttributedObservation[dt.Observation],
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

	// collect onramp and nonce manager addresses
	onrampAddrs := make(map[cciptypes.ChainSelector][][]byte)
	var nonceManagerAddrs [][]byte
	for _, ao := range aos {
		for chain, addr := range ao.Observation.OnRamp {
			// we don't want invalid observations to "poison" the consensus.
			if len(addr) == 0 {
				cdp.lggr.Warnf("skipping empty onramp address in observation from Oracle %d", ao.OracleID)
				continue
			}
			onrampAddrs[chain] = append(onrampAddrs[chain], addr)
		}

		if len(ao.Observation.DestNonceManager) == 0 {
			cdp.lggr.Warnf("skipping empty nonce manager address in observation from Oracle %d", ao.OracleID)
			continue
		}
		nonceManagerAddrs = append(
			nonceManagerAddrs,
			ao.Observation.DestNonceManager,
		)
	}

	contracts := make(reader.ContractAddresses)
	// onramps and dest nonce managers are determined by reading the _dest_ chain, therefore
	// we MUST use the fChain value of the dest chain to determine
	// the consensus onramp.
	fDestChainThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](
		consensus.TwoFPlus1(fChain[cdp.dest]),
	)
	onrampConsensus := consensus.GetConsensusMap(
		cdp.lggr,
		"onramp",
		onrampAddrs,
		fDestChainThresh,
	)
	cdp.lggr.Infow("Determined consensus onramps",
		"onrampConsensus", onrampConsensus,
		"onrampAddrs", onrampAddrs,
		"fDestChainThresh", fDestChainThresh,
	)
	if len(onrampConsensus) == 0 {
		cdp.lggr.Warnw("No consensus on onramps, onrampConsensus map is empty")
	}
	contracts[consts.ContractNameOnRamp] = onrampConsensus

	nonceManagerConsensus := consensus.GetConsensusMap(
		cdp.lggr,
		"nonceManager",
		map[cciptypes.ChainSelector][][]byte{cdp.dest: nonceManagerAddrs},
		fDestChainThresh,
	)
	cdp.lggr.Infow("Determined consensus nonce manager",
		"nonceManagerConsensus", nonceManagerConsensus,
		"nonceManagerAddrs", nonceManagerAddrs,
		"fDestChainThresh", fDestChainThresh,
	)
	contracts[consts.ContractNameNonceManager] = nonceManagerConsensus

	// call Sync to bind contracts.
	if err := (*cdp.reader).Sync(context.Background(), contracts); err != nil {
		return dt.Outcome{}, fmt.Errorf("unable to sync contracts: %w", err)
	}

	return dt.Outcome{}, nil
}
