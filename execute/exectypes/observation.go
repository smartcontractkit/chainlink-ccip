package exectypes

import (
	"encoding/json"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type CommitObservations map[cciptypes.ChainSelector][]CommitData
type MessageObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message

// Observation is the observation of the ExecutePlugin.
// TODO: revisit observation types. The maps used here are more space efficient and easier to work
// with but require more transformations compared to the on-chain representations.
type Observation struct {
	// CommitReports are determined during the first phase of execute.
	// It contains the commit reports we would like to execute in the following round.
	CommitReports CommitObservations `json:"commitReports"`
	// Messages are determined during the second phase of execute.
	// Ideally, it contains all the messages identified by the previous outcome's
	// NextCommits. With the previous outcome, and these messsages, we can build the
	// execute report.
	Messages MessageObservations `json:"messages"`
	// TODO: some of the nodes configuration may need to be included here.
}

func NewObservation(
	commitReports CommitObservations, messages MessageObservations) Observation {
	return Observation{
		CommitReports: commitReports,
		Messages:      messages,
	}
}

func (obs Observation) Encode() ([]byte, error) {
	return json.Marshal(obs)
}

func DecodeObservation(b []byte) (Observation, error) {
	obs := Observation{}
	err := json.Unmarshal(b, &obs)
	return obs, err
}
