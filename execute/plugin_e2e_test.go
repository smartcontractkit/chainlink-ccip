package execute

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

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

// Testing first scenario from the diagram:
// TODO: add diagram in github instead of using external link
// https://app.excalidraw.com/l/AdjkJ3DaenS/84EpHxkgbND
func TestPluginSkipEmptyReports(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	crBlockNumber := uint64(1000)
	currentTimestamp := time.Now().Add(-4 * time.Hour)

	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcSelector}, dstSelector)
	// Add empty reports to the reader, these are to mock price reports without merkle roots.
	// All of them are finalized,
	// Note: As the time of writing this test unfinalized in ContractReader includes finalized and unfinalized.
	for i := 0; i < maxCommitReportsToFetch; i++ {
		currentTimestamp = currentTimestamp.Add(time.Second)
		crBlockNumber++
		commitReportWithMeta := cciptypes.CommitPluginReportWithMeta{
			Report:    cciptypes.CommitPluginReport{},
			BlockNum:  crBlockNumber,
			Timestamp: currentTimestamp,
		}
		intTest.ccipReader.UnfinalizedReports = append(intTest.ccipReader.UnfinalizedReports, commitReportWithMeta)
		intTest.ccipReader.FinalizedReports = append(intTest.ccipReader.FinalizedReports, commitReportWithMeta)
	}

	// Add messages, These will be in an unfinalized report.
	messages := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(100, srcSelector, dstSelector, true),
		makeMsgWithMetadata(101, srcSelector, dstSelector, true),
		makeMsgWithMetadata(102, srcSelector, dstSelector, false),
		makeMsgWithMetadata(103, srcSelector, dstSelector, false),
		makeMsgWithMetadata(104, srcSelector, dstSelector, false),
		makeMsgWithMetadata(105, srcSelector, dstSelector, false),
	}
	intTest.WithMessages(messages, crBlockNumber, currentTimestamp, 1, srcSelector)

	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round.
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit UnfinalizedReports
	// No pending commit reports, all lenientMaxMsgsPerObs reports are with no merkelRoots.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.CommitReports, 0)

	// Should have updated
	// Round 1 - Get Commit UnfinalizedReports
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

func TestPlugin_EncodingSizeLimits(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	// Create messages with large data payloads
	largeMessages := []inmem.MessagesWithMetadata{}
	// TODO: make size co related with maxObservationSize
	nMessages := 10
	for i := 1; i <= nMessages; i++ {
		// Only  (half of the messages - 1) will be included in the observation
		// Notice that there's encoding overhead from other fields in observation, meaning that the message in the
		// middle of the observation will be truncated.
		size := lenientMaxObservationLength / (nMessages / 2)
		largeMessages = append(largeMessages, inmem.MessagesWithMetadata{
			Message:     makeMessageWithData(i, size, srcSelector, dstSelector).Message,
			Executed:    false,
			Destination: dstSelector,
		})
	}

	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcSelector}, dstSelector)

	// Add large messages
	intTest.WithMessages(largeMessages, 1000, time.Now().Add(-4*time.Hour), 1, srcSelector)

	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// GetCommitReports round
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.GetCommitReports, outcome.State)
	require.Len(t, outcome.CommitReports, 1)

	// GetMessages round - encoding size should limit number of messages
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.GetMessages, outcome.State)

	// Count messages with data (not empty placeholder messages)
	fullMsgCount := 0
	for _, report := range outcome.CommitReports {
		for _, msg := range report.Messages {
			if len(msg.Data) > 0 {
				fullMsgCount++
			}
		}
	}

	// Only half of the messages should be included in the observation
	require.GreaterOrEqual(t, len(largeMessages)/2, fullMsgCount,
		"Encoding size limit should prevent all messages from being included with their data")

	// Filter round
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Filter, outcome.State)

	// Verify report was created with the messages that were included
	require.NotEmpty(t, outcome.Report.ChainReports)
	sequenceNumbers := extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{1, 2, 3, 4})

	// Do another full round
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.GetCommitReports, outcome.State)
	require.NotEmpty(t, outcome.CommitReports)

	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.GetMessages, outcome.State)
	require.Len(t, outcome.CommitReports, 1)

	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Filter, outcome.State)

	seqNums := make([]cciptypes.SeqNum, 0)
	for _, report := range outcome.CommitReports {
		for _, msg := range report.Messages {
			if len(msg.Data) > 0 {
				seqNums = append(seqNums, msg.Header.SequenceNumber)
			}
		}
	}
	// next 4 messages
	require.ElementsMatch(t, seqNums, []cciptypes.SeqNum{5, 6, 7, 8})

	// Do another full round
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.GetCommitReports, outcome.State)
	require.NotEmpty(t, outcome.CommitReports)

	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.GetMessages, outcome.State)
	require.Len(t, outcome.CommitReports, 1)

	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Filter, outcome.State)

	seqNums = make([]cciptypes.SeqNum, 0)
	for _, report := range outcome.CommitReports {
		for _, msg := range report.Messages {
			if len(msg.Data) > 0 {
				seqNums = append(seqNums, msg.Header.SequenceNumber)
			}
		}
	}
	// next 4 messages
	require.ElementsMatch(t, seqNums, []cciptypes.SeqNum{9, 10})
}

func makeMessageWithData(seqNum, byteSize int, src, dst cciptypes.ChainSelector) inmem.MessagesWithMetadata {
	// Create a message with large data payload to test encoding size limits
	msg := makeMsgWithMetadata(cciptypes.SeqNum(seqNum), src, dst, false)

	largeData := make([]byte, byteSize)
	for i := range largeData {
		largeData[i] = byte(i % 256)
	}

	msg.Message.Data = largeData
	return msg
}
