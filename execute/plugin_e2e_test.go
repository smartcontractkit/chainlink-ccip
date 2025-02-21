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
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
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

	intTest := SetupSimpleTest(t, logger.Test(t), srcSelector, dstSelector)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour), 1)
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

func Test_ExcludingCostlyMessages(t *testing.T) {
	ctx := tests.Context(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	messages := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(100, srcSelector, dstSelector, false, withFeeValueJuels(100)),
		makeMsgWithMetadata(101, srcSelector, dstSelector, false, withFeeValueJuels(200)),
		makeMsgWithMetadata(102, srcSelector, dstSelector, false, withFeeValueJuels(300)),
	}

	messageTimestamp := time.Now().Add(-4 * time.Hour)
	tm := timeMachine{now: messageTimestamp}

	intTest := SetupSimpleTest(t, logger.Test(t), srcSelector, dstSelector)
	intTest.WithMessages(messages, 1000, messageTimestamp, 1)
	intTest.WithCustomFeeBoosting(1.0, tm.Now, map[cciptypes.Bytes32]plugintypes.USD18{
		messages[0].Header.MessageID: plugintypes.NewUSD18(40000),
		messages[1].Header.MessageID: plugintypes.NewUSD18(200000),
		messages[2].Header.MessageID: plugintypes.NewUSD18(200000),
	})

	runner := intTest.Start()
	defer intTest.Close()

	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// First outcome is empty - all messages are too expensive to be executed
	// Message1 cost=40000,  fee=10000
	// Message2 cost=200000, fee=20000
	// Message3 cost=200000, fee=30000
	for i := 0; i < 3; i++ {
		outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	}
	require.Len(t, outcome.Report.ChainReports, 0)

	// 4 hours later, we agree to pay higher fee, but only for the first message
	// Message1 cost=40000,  fee=50000 boosted original_fee * (1 + 4*1.0),
	// Message2 cost=200000, fee=20000
	// Message3 cost=200000, fee=30000
	tm.SetNow(time.Now())
	for i := 0; i < 3; i++ {
		outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	}
	sequenceNumbers := extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{100})

	// Second message execution cost drops, it should be included in the outcome
	// the first message is excluded by the inflight message cache.
	// Message1 cost=40000,  fee=50000   boosted original_fee * (1 + 4*1.0),
	// Message2 cost=40000,  fee=100000
	// Message3 cost=200000, fee=150000
	intTest.UpdateExecutionCost(messages[1].Header.MessageID, 40000)
	for i := 0; i < 3; i++ {
		outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	}
	sequenceNumbers = extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{101})

	// 3 hours in the future, we agree to pay higher fee for the third message (7 hours since the message was sent)
	// the first and second message are excluded by the inflight message cache.
	// Message1 cost=40000,  fee=80000  boosted 10000 * (1 + 7*1.0),
	// Message2 cost=40000,  fee=160000
	// Message3 cost=200000, fee=240000
	tm.SetNow(time.Now().Add(3 * time.Hour))
	for i := 0; i < 3; i++ {
		outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	}
	sequenceNumbers = extractSequenceNumbers(outcome.Report.ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102})
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

	intTest := SetupSimpleTest(t, mocks.NullLogger, srcSelector, dstSelector)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour), nReports)
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

	intTest := SetupSimpleTest(t, logger.Test(t), srcSelector, dstSelector)

	// Set up first report - should be executed and finalized
	intTest.WithMessages(finalizedMessages, 1000, time.Now().Add(-4*time.Hour), 1)

	// Set up second report - should remain available
	intTest.WithMessages(unexecutedMessages, 1001, time.Now().Add(-3*time.Hour), 1)

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
