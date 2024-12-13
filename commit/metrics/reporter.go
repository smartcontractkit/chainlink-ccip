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

type Reporter interface {
	TrackObservation(obs committypes.Observation, state string)
	TrackOutcome(outcome committypes.Outcome, state string)

	TrackChainFeeObservation(obs chainfee.Observation, state string)
	TrackChainFeeOutcome(outcome chainfee.Outcome, state string)

	TrackMerkleObservation(obs merkleroot.Observation, state string)
	TrackMerkleOutcome(outcome merkleroot.Outcome, state string)

	TrackTokenPricesObservation(obs merkleroot.Observation, state string)
	TrackTokenPricesOutcome(outcome merkleroot.Outcome, state string)
}

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

func (p *PromReporter) TrackObservation(obs committypes.Observation, state string) {
}

func (p *PromReporter) TrackOutcome(outcome committypes.Outcome, state string) {
}

func (p *PromReporter) TrackChainFeeObservation(obs chainfee.Observation, state string) {
}

func (p *PromReporter) TrackChainFeeOutcome(outcome chainfee.Outcome, state string) {
}

func (p *PromReporter) TrackMerkleObservation(obs merkleroot.Observation, state string) {
}

func (p *PromReporter) TrackMerkleOutcome(outcome merkleroot.Outcome, state string) {
}

func (p *PromReporter) TrackTokenPricesObservation(obs merkleroot.Observation, state string) {
}

func (p *PromReporter) TrackTokenPricesOutcome(outcome merkleroot.Outcome, state string) {
}

type Noop struct{}

func (n *Noop) TrackObservation(committypes.Observation, string) {}

func (n *Noop) TrackOutcome(committypes.Outcome, string) {}

func (n *Noop) TrackChainFeeObservation(chainfee.Observation, string) {}

func (n *Noop) TrackChainFeeOutcome(chainfee.Outcome, string) {}

func (n *Noop) TrackMerkleObservation(merkleroot.Observation, string) {}

func (n *Noop) TrackMerkleOutcome(merkleroot.Outcome, string) {}

func (n *Noop) TrackTokenPricesObservation(merkleroot.Observation, string) {}

func (n *Noop) TrackTokenPricesOutcome(merkleroot.Outcome, string) {}

var _ Reporter = &Noop{}
var _ Reporter = &PromReporter{}
var _ chainfee.MetricsReporter = &PromReporter{}
var _ merkleroot.MetricsReporter = &PromReporter{}
var _ tokenprice.MetricsReporter = &PromReporter{}
