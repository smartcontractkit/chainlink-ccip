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

	TrackMerkleObservation(obs merkleroot.Observation, state string)
	TrackMerkleOutcome(outcome merkleroot.Outcome, state string)

	TrackChainFeeObservation(obs chainfee.Observation)
	TrackChainFeeOutcome(outcome chainfee.Outcome)

	TrackTokenPricesObservation(obs tokenprice.Observation)
	TrackTokenPricesOutcome(outcome tokenprice.Outcome)
}

type CommitPluginReporter interface {
	TrackObservation(obs committypes.Observation, state string)
	TrackOutcome(outcome committypes.Outcome, state string)
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

type Noop struct{}

func (n *Noop) TrackObservation(committypes.Observation, string) {}

func (n *Noop) TrackOutcome(committypes.Outcome, string) {}

func (n *Noop) TrackChainFeeObservation(chainfee.Observation) {}

func (n *Noop) TrackChainFeeOutcome(chainfee.Outcome) {}

func (n *Noop) TrackMerkleObservation(merkleroot.Observation, string) {}

func (n *Noop) TrackMerkleOutcome(merkleroot.Outcome, string) {}

func (n *Noop) TrackTokenPricesObservation(tokenprice.Observation) {}

func (n *Noop) TrackTokenPricesOutcome(tokenprice.Outcome) {}

var _ Reporter = &Noop{}
var _ Reporter = &PromReporter{}
var _ CommitPluginReporter = &PromReporter{}
var _ chainfee.MetricsReporter = &PromReporter{}
var _ merkleroot.MetricsReporter = &PromReporter{}
var _ tokenprice.MetricsReporter = &PromReporter{}
