package plugintypes

import (
	"encoding/json"
	"fmt"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ///////////////////////
// Execute Observation //
// ///////////////////////

type ExecutePluginCommitDataWithMessages struct {
	ExecutePluginCommitData
	Messages []cciptypes.CCIPMsg `json:"messages"`
}

// ExecutePluginCommitData is the data that is committed to the chain.
type ExecutePluginCommitData struct {
	// SourceChain of the chain that contains the commit report.
	SourceChain cciptypes.ChainSelector `json:"chainSelector"`
	// Timestamp of the block that contains the commit.
	Timestamp time.Time `json:"timestamp"`
	// BlockNum of the block that contains the commit.
	BlockNum uint64 `json:"blockNum"`
	// MerkleRoot of the messages that are in this commit report.
	MerkleRoot cciptypes.Bytes32 `json:"merkleRoot"`
	// SequenceNumberRange of the messages that are in this commit report.
	SequenceNumberRange cciptypes.SeqNumRange `json:"sequenceNumberRange"`
	// ExecutedMessages are the messages in this report that have already been executed.
	ExecutedMessages []cciptypes.SeqNum `json:"executed"`
}

type ExecutePluginCommitObservations map[cciptypes.ChainSelector][]ExecutePluginCommitDataWithMessages
type ExecutePluginMessageObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.CCIPMsg

// ExecutePluginObservation is the observation of the ExecutePlugin.
// TODO: revisit observation types. The maps used here are more space efficient and easier to work
// with but require more transformations compared to the on-chain representations.
type ExecutePluginObservation struct {
	// CommitReports are determined during the first phase of execute.
	// It contains the commit reports we would like to execute in the following round.
	CommitReports ExecutePluginCommitObservations `json:"commitReports"`
	// Messages are determined during the second phase of execute.
	// Ideally, it contains all the messages identified by the previous outcome's
	// NextCommits. With the previous outcome, and these messsages, we can build the
	// execute report.
	Messages ExecutePluginMessageObservations `json:"messages"`
	// TODO: some of the nodes configuration may need to be included here.
}

func NewExecutePluginObservation(
	commitReports ExecutePluginCommitObservations, messages ExecutePluginMessageObservations) ExecutePluginObservation {
	return ExecutePluginObservation{
		CommitReports: commitReports,
		Messages:      messages,
	}
}

func (obs ExecutePluginObservation) Encode() ([]byte, error) {
	return json.Marshal(obs)
}

func DecodeExecutePluginObservation(b []byte) (ExecutePluginObservation, error) {
	obs := ExecutePluginObservation{}
	err := json.Unmarshal(b, &obs)
	return obs, err
}

// ///////////////////
// Execute Outcome //
// ///////////////////

// ExecutePluginOutcome is the outcome of the ExecutePlugin.
type ExecutePluginOutcome struct {
	// PendingCommitReports are the oldest reports with pending commits. The slice is
	// sorted from oldest to newest.
	PendingCommitReports []ExecutePluginCommitDataWithMessages `json:"commitReports"`

	// Report is built from the oldest pending commit reports.
	Report cciptypes.ExecutePluginReport `json:"report"`
}

func NewExecutePluginOutcome(
	pendingCommits []ExecutePluginCommitDataWithMessages,
	report cciptypes.ExecutePluginReport,
) ExecutePluginOutcome {
	return ExecutePluginOutcome{
		PendingCommitReports: pendingCommits,
		Report:               report,
	}
}

func (o ExecutePluginOutcome) Encode() ([]byte, error) {
	return json.Marshal(o)
}

func DecodeExecutePluginOutcome(b []byte) (ExecutePluginOutcome, error) {
	o := ExecutePluginOutcome{}
	err := json.Unmarshal(b, &o)
	return o, err
}

func (o ExecutePluginOutcome) String() string {
	return fmt.Sprintf("NextCommits: %v", o.PendingCommitReports)
}
