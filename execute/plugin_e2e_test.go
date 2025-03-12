package execute

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var ocrTypeCodec = ocrtypecodec.DefaultExecCodec

func TestPlugin(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	messages := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(100, srcSelector, dstSelector, true),
		makeMsgWithMetadata(101, srcSelector, dstSelector, true),
		makeMsgWithMetadata(102, srcSelector, dstSelector, false),
		makeMsgWithMetadata(103, srcSelector, dstSelector, false),
		makeMsgWithMetadata(104, srcSelector, dstSelector, false),
		makeMsgWithMetadata(105, srcSelector, dstSelector, false),
	}

	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcSelector}, dstSelector)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour), 1, srcSelector)
	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	// One pending commit report only.
	// Two of the messages are executed which should be indicated in the Outcome.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.CommitReports, 1)
	require.ElementsMatch(t, outcome.CommitReports[0].ExecutedMessages, []cciptypes.SeqNum{100, 101})

	// Round 2 - Get Messages
	// Messages now attached to the pending commit.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.CommitReports, 1)

	// Round 3 - Filter
	// An execute report with the following messages executed: 102, 103, 104, 105.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 1)
	sequenceNumbers := extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 105})
}

// TestExceedSizeObservation tests the case where the observation size exceeds the maximum size.
// Setup multiple commit reports that the total size of the observation exceeds the maximum size.
// Make sure that the observation reports are truncated to fit the maximum size.
func TestExceedSizeObservation(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	// 1 msg * 1 byte    -> 879  | 2 msg * 1 byte -> 1311 | 3 msg * 1 byte -> 1743
	// 3 msg * 2 bytes   -> 882  | 3 msg * 2 byte -> 1319 | 3 msg * 2 byte -> 1755
	// 10 msg * 1 byte   -> 897
	// 100 msg * 1 byte  -> 1077
	// 1000 msg * 1 byte -> 2877
	msgDataSize := 1000
	maxMsgsPerReport := 398
	nReports := 2
	maxMessages := maxMsgsPerReport * nReports // Currently 398 message per report is the max with msgDataSize = 1000
	startSeqNr := cciptypes.SeqNum(100)

	messages := make([]inmem.MessagesWithMetadata, 0, maxMessages)
	for i := 0; i < maxMessages; i++ {
		messages = append(messages,
			makeMsgWithMetadata(
				startSeqNr+cciptypes.SeqNum(i),
				srcSelector,
				dstSelector,
				false,
				withData(make([]byte, msgDataSize)),
			),
		)
	}

	intTest := SetupSimpleTest(t, mocks.NullLogger, []cciptypes.ChainSelector{srcSelector}, dstSelector)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour), nReports, srcSelector)
	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	// Two pending commit reports.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.CommitReports, nReports)

	// Round 2 - Get Messages
	// Only 1 pending report from previous round.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.CommitReports, 2)
	require.Len(t, outcome.CommitReports[0].Messages, maxMsgsPerReport)

	// Round 3 - Filter
	// An execute report with the messages executed until the max per report
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 2)
	sequenceNumbers := extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.Len(t, sequenceNumbers, maxMsgsPerReport)
}

func TestPlugin_FinalizedUnfinalizedCaching(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	// Create messages for finalized execution report
	finalizedMessages := []inmem.MessagesWithMetadata{
		{
			Message:     makeMsgWithMetadata(100, srcSelector, dstSelector, true).Message, // Executed
			Executed:    true,
			Destination: dstSelector,
		},
		{
			Message:     makeMsgWithMetadata(101, srcSelector, dstSelector, true).Message, // Executed
			Executed:    true,
			Destination: dstSelector,
		},
	}

	// Create messages for non-executed report
	unexecutedMessages := []inmem.MessagesWithMetadata{
		{
			Message:     makeMsgWithMetadata(200, srcSelector, dstSelector, false).Message, // Not executed
			Executed:    false,
			Destination: dstSelector,
		},
		{
			Message:     makeMsgWithMetadata(201, srcSelector, dstSelector, false).Message, // Not executed
			Executed:    false,
			Destination: dstSelector,
		},
	}

	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcSelector}, dstSelector)

	// Set up first report - should be executed and finalized
	intTest.WithMessages(finalizedMessages, 1000, time.Now().Add(-4*time.Hour), 1, srcSelector)

	// Set up second report - should remain available
	intTest.WithMessages(unexecutedMessages, 1001, time.Now().Add(-3*time.Hour), 1, srcSelector)

	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// First round - process both reports
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)

	// First report should be marked as executed and removed from commits
	// Second report should be available since it's not executed
	require.Len(t, outcome.CommitReports, 1)

	// Check that we see the second report
	report := outcome.CommitReports[0]
	require.Equal(t, cciptypes.NewSeqNumRange(200, 201), report.SequenceNumberRange)
}

func TestPlugin_CommitReportTimestampOrdering(t *testing.T) {
	ctx := tests.Context(t)

	srcChainA := cciptypes.ChainSelector(1)
	srcChainB := cciptypes.ChainSelector(3)
	dstSelector := cciptypes.ChainSelector(2)

	// Create messages for multiple commit reports with different timestamps
	baseTime := time.Now().Add(-4 * time.Hour)
	msgSets := []struct {
		messages  []inmem.MessagesWithMetadata
		timestamp time.Time
		blockNum  uint64
		srcChain  cciptypes.ChainSelector
	}{
		{
			// Latest messages
			messages: []inmem.MessagesWithMetadata{
				makeMsgWithMetadata(104, srcChainA, dstSelector, false),
				makeMsgWithMetadata(105, srcChainA, dstSelector, false),
			},
			timestamp: baseTime.Add(2 * time.Hour),
			blockNum:  1002,
			srcChain:  srcChainA,
		},
		{
			// Earliest messages, From a chain that's chronologically after the other chain
			messages: []inmem.MessagesWithMetadata{
				makeMsgWithMetadata(100, srcChainB, dstSelector, false),
				makeMsgWithMetadata(101, srcChainB, dstSelector, false),
			},
			timestamp: baseTime,
			blockNum:  1000,
			srcChain:  srcChainB,
		},
		{
			// Middle messages
			messages: []inmem.MessagesWithMetadata{
				makeMsgWithMetadata(102, srcChainA, dstSelector, false),
				makeMsgWithMetadata(103, srcChainA, dstSelector, false),
			},
			timestamp: baseTime.Add(time.Hour),
			blockNum:  1001,
			srcChain:  srcChainA,
		},
	}

	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcChainA, srcChainB}, dstSelector)

	// Add messages in non-chronological order
	for _, set := range msgSets {
		intTest.WithMessages(set.messages, set.blockNum, set.timestamp, 1, set.srcChain)
	}

	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// GetCommitReports round - should return chronologically ordered reports
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.GetCommitReports, outcome.State)
	require.Len(t, outcome.CommitReports, 3)

	// Verify timestamps are in ascending order
	for i := 0; i < len(outcome.CommitReports)-1; i++ {
		require.True(t, outcome.CommitReports[i].Timestamp.Before(
			outcome.CommitReports[i+1].Timestamp),
			"commit reports should be ordered by timestamp")
	}

	// Verify the specific order matches our expected chronological order
	require.Equal(t, cciptypes.SeqNum(100),
		outcome.CommitReports[0].SequenceNumberRange.Start(), "oldest report should be first")
	require.Equal(t, cciptypes.SeqNum(102),
		outcome.CommitReports[1].SequenceNumberRange.Start(), "middle report should be second")
	require.Equal(t, cciptypes.SeqNum(104),
		outcome.CommitReports[2].SequenceNumberRange.Start(), "newest report should be last")
}
