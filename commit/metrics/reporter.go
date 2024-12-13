package metrics

import (
	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
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

var _ Reporter = &Noop{}

type Noop struct{}

func (n *Noop) TrackObservation(obs committypes.Observation, state string) {}

func (n *Noop) TrackOutcome(outcome committypes.Outcome, state string) {}

func (n *Noop) TrackChainFeeObservation(obs chainfee.Observation, state string) {}

func (n *Noop) TrackChainFeeOutcome(outcome chainfee.Outcome, state string) {}

func (n *Noop) TrackMerkleObservation(obs merkleroot.Observation, state string) {}

func (n *Noop) TrackMerkleOutcome(outcome merkleroot.Outcome, state string) {}

func (n *Noop) TrackTokenPricesObservation(obs merkleroot.Observation, state string) {}

func (n *Noop) TrackTokenPricesOutcome(outcome merkleroot.Outcome, state string) {}
