package discovery

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

// Ensure that PluginProcessor is implemented.
var _ plugincommon.PluginProcessor[dt.Query, dt.Observation, dt.Outcome] = &ContractDiscoveryProcessor{}

// ContractDiscoveryProcessor is a plugin processor for discovering contracts.
type ContractDiscoveryProcessor struct {
	lggr      logger.Logger
	reader    *reader.CCIPReader
	homechain reader.HomeChain
	dest      cciptypes.ChainSelector
	fRoleDON  int
}

func NewContractDiscoveryProcessor(
	lggr logger.Logger,
	reader *reader.CCIPReader,
	homechain reader.HomeChain,
	dest cciptypes.ChainSelector,
	fRoleDON int,
) *ContractDiscoveryProcessor {
	return &ContractDiscoveryProcessor{
		lggr:      lggr,
		reader:    reader,
		homechain: homechain,
		dest:      dest,
		fRoleDON:  fRoleDON,
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
	contracts, err := (*cdp.reader).DiscoverContracts(ctx, cdp.dest)
	if err != nil {
		if errors.Is(err, reader.ErrContractReaderNotFound) {
			// Not a dest reader, no observations will be made.
			// Processor is not disabled because the outcome phase will bind observed contracts.
			return dt.Observation{}, nil
		}
		return dt.Observation{}, fmt.Errorf("unable to discover contracts: %w", err)
	}

	fChain, err := cdp.homechain.GetFChain()
	if err != nil {
		return dt.Observation{}, fmt.Errorf("unable to get fchain: %w", err)
	}

	return dt.Observation{
		FChain: fChain,
		OnRamp: contracts[consts.ContractNameOnRamp],
	}, nil
}

func (cdp *ContractDiscoveryProcessor) ValidateObservation(
	_ dt.Outcome, _ dt.Query, _ plugincommon.AttributedObservation[dt.Observation],
) error {
	// TODO: make sure all observations come from dest readers.
	return nil
}

// Outcome comes to consensus on the contract addresses and updates the chainreader. It doesn't actually
// return an Outcome.
func (cdp *ContractDiscoveryProcessor) Outcome(
	_ context.Context, _ dt.Outcome, _ dt.Query, aos []plugincommon.AttributedObservation[dt.Observation],
) (dt.Outcome, error) {
	// come to consensus on the onramp addresses and update the chainreader.

	// fChain consensus - uses the role DON F value because all nodes can observe the home chain.
	fChainObs := make(map[cciptypes.ChainSelector][]int)
	for _, ao := range aos {
		for chainSel, f := range ao.Observation.FChain {
			fChainObs[chainSel] = append(fChainObs[chainSel], f)
		}
	}
	fMin := make(map[cciptypes.ChainSelector]int)
	for chain := range fChainObs {
		fMin[chain] = cdp.fRoleDON
	}

	fChain := plugincommon.GetConsensusMap(cdp.lggr, "fChain", fChainObs, fMin)

	// onramp address consensus
	onrampAddrs := make(map[cciptypes.ChainSelector][][]byte)
	for _, ao := range aos {
		for chain, addr := range ao.Observation.OnRamp {
			onrampAddrs[chain] = append(onrampAddrs[chain], addr)
		}
	}

	// call Sync to bind contracts.
	contracts := make(map[string]map[cciptypes.ChainSelector][]byte)
	contracts[consts.ContractNameOnRamp] = plugincommon.GetConsensusMap(cdp.lggr, "onramp", onrampAddrs, fChain)
	if err := (*cdp.reader).Sync(context.Background(), contracts); err != nil {
		return dt.Outcome{}, fmt.Errorf("unable to sync contracts: %w", err)
	}

	return dt.Outcome{}, nil
}
