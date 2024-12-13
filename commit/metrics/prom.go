package metrics

import (
	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type PromReporter struct {
	lggr    logger.Logger
	chainID string
}

func NewPromReporter(lggr logger.Logger, selector cciptypes.ChainSelector) (*PromReporter, error) {
	chainID, err := sel.GetChainIDFromSelector(uint64(selector))
	if err != nil {
		return nil, err
	}

	return &PromReporter{
		lggr:    lggr,
		chainID: chainID,
	}, nil
}

func (p *PromReporter) TrackObservation(obs committypes.Observation) {
}

func (p *PromReporter) TrackOutcome(outcome committypes.Outcome) {
}

func (p *PromReporter) TrackChainFeeObservation(obs chainfee.Observation) {
}

func (p *PromReporter) TrackChainFeeOutcome(outcome chainfee.Outcome) {
}

func (p *PromReporter) TrackMerkleObservation(obs merkleroot.Observation, state string) {
}

func (p *PromReporter) TrackMerkleOutcome(outcome merkleroot.Outcome, state string) {
}

func (p *PromReporter) TrackTokenPricesObservation(obs tokenprice.Observation) {
}

func (p *PromReporter) TrackTokenPricesOutcome(outcome tokenprice.Outcome) {
}
