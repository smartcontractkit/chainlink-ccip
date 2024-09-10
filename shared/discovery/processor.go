package discovery

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/shared"
)

// Ensure that shared.PluginProcessor is implemented.
var _ shared.PluginProcessor[Query, Observation, Outcome] = &ContractDiscoveryProcessor{}

// Outcome isn't needed for this processor.
type Outcome []byte

// Query isn't needed for this processor.
type Query []byte

// Observation of contract addresses.
type Observation struct {
	OnRamp map[ccipocr3.ChainSelector][]byte

	// TODO: some sort of request flag.
	// Request bool
}

// ContractDiscoveryProcessor is a plugin processor for discovering contracts.
type ContractDiscoveryProcessor struct {
	lggr      logger.Logger
	reader    reader.CCIPReader
	homechain reader.HomeChain
	dest      ccipocr3.ChainSelector
}

// Query is not needed for this processor.
func (cdp *ContractDiscoveryProcessor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return nil, nil
}

// Observation reads contract addresses from the destination chain.
// In the future this should be extended to omit observations unless one of the nodes requests addresses.
func (cdp *ContractDiscoveryProcessor) Observation(
	ctx context.Context, prevOutcome Outcome, query Query,
) (Observation, error) {
	contracts, err := cdp.reader.DiscoverContracts(ctx, cdp.dest)
	if err != nil {
		if errors.Is(err, reader.ErrContractReaderNotFound) {
			// not a dest reader, no observations (disable entire processor?).
			return Observation{}, nil
		}
		return Observation{}, fmt.Errorf("unable to discover contracts: %w", err)
	}

	return Observation{
		OnRamp: contracts[consts.ContractNameOnRamp],
	}, nil
}

func (cdp *ContractDiscoveryProcessor) ValidateObservation(
	prevOutcome Outcome, query Query, ao shared.AttributedObservation[Observation],
) error {
	return nil
}

// Outcome
func (cdp *ContractDiscoveryProcessor) Outcome(
	prevOutcome Outcome, query Query, aos []shared.AttributedObservation[Observation],
) (Outcome, error) {
	// come to consensus on the onramp addresses and update the chainreader.

	fChain, err := cdp.homechain.GetFChain()
	if err != nil {
		return Outcome{}, fmt.Errorf("unable to get fchain: %w", err)
	}

	onrampAddrs := make(map[ccipocr3.ChainSelector][][]byte)
	for _, ao := range aos {
		for chain, addr := range ao.Observation.OnRamp {
			onrampAddrs[chain] = append(onrampAddrs[chain], addr)
		}
	}

	contracts := make(map[string]map[ccipocr3.ChainSelector][]byte)
	contracts[consts.ContractNameOnRamp] = shared.GetConsensusMap(cdp.lggr, "onramp", onrampAddrs, fChain)
	if err := cdp.reader.Sync(context.Background(), contracts); err != nil {
		return Outcome{}, fmt.Errorf("unable to sync contracts: %w", err)
	}

	return Outcome{}, nil
}
