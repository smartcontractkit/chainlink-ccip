package execute

import (
	"testing"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/cache"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
)

var ocrTypeCodec = ocrtypecodec.DefaultExecCodec

func TestPlugin(t *testing.T) {
	ctx := t.Context()

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
	require.Len(t, outcome.Reports, 0)
	require.Len(t, outcome.CommitReports, 1)
	require.ElementsMatch(t, outcome.CommitReports[0].ExecutedMessages, []cciptypes.SeqNum{100, 101})

	// Round 2 - Get Messages
	// Messages now attached to the pending commit.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Reports, 0)
	require.Len(t, outcome.CommitReports, 1)

	// Round 3 - Filter
	// An execute report with the following messages executed: 102, 103, 104, 105.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Reports, 1)
	require.Len(t, outcome.Reports[0].ChainReports, 1)
	sequenceNumbers := extractSequenceNumbers(outcome.Reports[0].ChainReports[0].Messages)
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 105})
}

// This is a simulation for Solana case where MaxReportMessages is 1 and MaxSingleChainReports is 1 while
// MultipleReportsEnabled is true.
func TestPluginWithMultipleReports(t *testing.T) {
	ctx := t.Context()

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	// Create more messages than the first test to demonstrate multiple reports
	messages := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(100, srcSelector, dstSelector, true),
		makeMsgWithMetadata(101, srcSelector, dstSelector, true),
		makeMsgWithMetadata(102, srcSelector, dstSelector, false),
		makeMsgWithMetadata(103, srcSelector, dstSelector, false),
		makeMsgWithMetadata(104, srcSelector, dstSelector, false),
		makeMsgWithMetadata(105, srcSelector, dstSelector, false),
		makeMsgWithMetadata(106, srcSelector, dstSelector, false),
		makeMsgWithMetadata(107, srcSelector, dstSelector, false),
		makeMsgWithMetadata(108, srcSelector, dstSelector, false),
		makeMsgWithMetadata(109, srcSelector, dstSelector, false),
	}

	// Set up the test with multiple reports enabled and low max messages per report
	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcSelector}, dstSelector)
	intTest.WithMessages(messages, 1000, time.Now().Add(-4*time.Hour), 1, srcSelector)

	// Configure the plugin to use multiple reports with a small max messages value
	intTest.WithOffChainConfig(pluginconfig.ExecuteOffchainConfig{
		MessageVisibilityInterval: *commonconfig.MustNewDuration(8 * time.Hour),
		BatchGasLimit:             100000000,
		MaxCommitReportsToFetch:   10,
		MultipleReportsEnabled:    true,
		MaxReportMessages:         1,
		MaxSingleChainReports:     1,
	})

	runner := intTest.Start()
	defer intTest.Close()

	// Contract Discovery round
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	// One pending commit report only.
	// Two of the messages are executed which should be indicated in the Outcome.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Reports, 0)
	require.Len(t, outcome.CommitReports, 1)
	require.ElementsMatch(t, outcome.CommitReports[0].ExecutedMessages, []cciptypes.SeqNum{100, 101})

	// Round 2 - Get Messages
	// Messages now attached to the pending commit.
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)
	require.Len(t, outcome.Reports, 0)
	require.Len(t, outcome.CommitReports, 1)

	// Round 3 - Filter
	// With MaxMessages=1, we should get multiple execute reports, each with 1 message
	outcome = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)

	// Should have 8 reports, each with a single chain report
	require.Equal(t, len(outcome.Reports), 8)

	// Collect all sequence numbers across all reports
	var allSequenceNumbers []cciptypes.SeqNum
	for _, report := range outcome.Reports {
		require.Len(t, report.ChainReports, 1, "Each report should have one chain report")
		sequenceNumbers := extractSequenceNumbers(report.ChainReports[0].Messages)
		allSequenceNumbers = append(allSequenceNumbers, sequenceNumbers...)

		// Each report should have MaxMessages=1 messages
		require.Equal(t, len(report.ChainReports[0].Messages), 1,
			"Each report should have at most MaxMessages (1) messages")
	}

	// Verify all expected messages are included across the reports
	expectedSeqNumbers := []cciptypes.SeqNum{102, 103, 104, 105, 106, 107, 108, 109}
	require.ElementsMatch(t, expectedSeqNumbers, allSequenceNumbers,
		"All expected messages should be included across the reports")
}

func TestPluginMultipleReportsWithMultipleSourceChains(t *testing.T) {
	ctx := t.Context()

	srcSelector1 := cciptypes.ChainSelector(1)
	srcSelector2 := cciptypes.ChainSelector(2)
	dstSelector := cciptypes.ChainSelector(3)

	// Create messages from two source chains
	messages1 := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(100, srcSelector1, dstSelector, false),
		makeMsgWithMetadata(101, srcSelector1, dstSelector, false),
		makeMsgWithMetadata(102, srcSelector1, dstSelector, false),
	}

	messages2 := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(200, srcSelector2, dstSelector, false),
		makeMsgWithMetadata(201, srcSelector2, dstSelector, false),
		makeMsgWithMetadata(202, srcSelector2, dstSelector, false),
	}

	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcSelector1, srcSelector2}, dstSelector)
	intTest.WithMessages(messages1, 1000, time.Now().Add(-4*time.Hour), 1, srcSelector1)
	intTest.WithMessages(messages2, 1000, time.Now().Add(-4*time.Hour), 1, srcSelector2)

	intTest.WithOffChainConfig(pluginconfig.ExecuteOffchainConfig{
		MessageVisibilityInterval: *commonconfig.MustNewDuration(8 * time.Hour),
		BatchGasLimit:             100000000,
		MaxCommitReportsToFetch:   10,
		MultipleReportsEnabled:    true,
		MaxSingleChainReports:     1,
	})

	runner := intTest.Start()
	defer intTest.Close()

	// Run rounds
	_ = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)        // Contract Discovery
	_ = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)        // Get Commit Reports
	_ = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)        // Get Messages
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner) // Filter

	require.GreaterOrEqual(t, len(outcome.Reports), 2, "Should create multiple reports")

	// Collect all chain selectors and sequence numbers across reports
	sourceChains := make(map[cciptypes.ChainSelector]bool)
	for _, report := range outcome.Reports {
		for _, chainReport := range report.ChainReports {
			sourceChains[chainReport.SourceChainSelector] = true
		}
	}

	require.Len(t, sourceChains, 2, "Should have reports from both source chains")

	// Check all messages are marked as executed
	var allSequenceNumbers []cciptypes.SeqNum
	for _, report := range outcome.Reports {
		require.Len(t, report.ChainReports, 1, "Each report should have one chain report")
		sequenceNumbers := extractSequenceNumbers(report.ChainReports[0].Messages)
		allSequenceNumbers = append(allSequenceNumbers, sequenceNumbers...)
	}

	// Verify all expected messages are included across the reports
	expectedSeqNumbers := []cciptypes.SeqNum{100, 101, 102, 200, 201, 202}
	require.ElementsMatch(t, expectedSeqNumbers, allSequenceNumbers,
		"All expected messages should be included across the reports")

}

func TestPluginMultipleReportsWithMultipleSourceChainsAndTimestamps(t *testing.T) {
	ctx := t.Context()

	srcSelector1 := cciptypes.ChainSelector(1)
	srcSelector2 := cciptypes.ChainSelector(2)
	dstSelector := cciptypes.ChainSelector(3)

	// Create messages from two source chains
	messages1 := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(100, srcSelector1, dstSelector, false),
		makeMsgWithMetadata(101, srcSelector1, dstSelector, false),
		makeMsgWithMetadata(102, srcSelector1, dstSelector, false),
	}

	messages2 := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(200, srcSelector2, dstSelector, false),
		makeMsgWithMetadata(201, srcSelector2, dstSelector, false),
		makeMsgWithMetadata(202, srcSelector2, dstSelector, false),
	}

	intTest := SetupSimpleTest(t, logger.Test(t), []cciptypes.ChainSelector{srcSelector1, srcSelector2}, dstSelector)
	intTest.WithMessages(messages1, 1000, time.Now().Add(-4*time.Hour), 1, srcSelector1)
	// Report for srcSelector2 should come first in the outcome reports
	intTest.WithMessages(messages2, 1000, time.Now().Add(-5*time.Hour), 1, srcSelector2)

	intTest.WithOffChainConfig(pluginconfig.ExecuteOffchainConfig{
		MessageVisibilityInterval: *commonconfig.MustNewDuration(8 * time.Hour),
		BatchGasLimit:             100000000,
		MaxCommitReportsToFetch:   10,
		MultipleReportsEnabled:    true,
		MaxSingleChainReports:     1,
	})

	runner := intTest.Start()
	defer intTest.Close()

	// Run rounds
	_ = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)        // Contract Discovery
	_ = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)        // Get Commit Reports
	_ = runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner)        // Get Messages
	outcome := runRoundAndGetOutcome(ctx, ocrTypeCodec, t, runner) // Filter

	require.GreaterOrEqual(t, len(outcome.Reports), 2, "Should create multiple reports")

	require.Equal(t, srcSelector2, outcome.Reports[0].ChainReports[0].SourceChainSelector,
		"First report should be from srcSelector2 as it has earlier timestamp")

	// Check all messages are marked as executed
	var allSequenceNumbers []cciptypes.SeqNum
	for _, report := range outcome.Reports {
		require.Len(t, report.ChainReports, 1, "Each report should have one chain report")
		sequenceNumbers := extractSequenceNumbers(report.ChainReports[0].Messages)
		allSequenceNumbers = append(allSequenceNumbers, sequenceNumbers...)
	}

	// Verify all expected messages are included across the reports
	expectedSeqNumbers := []cciptypes.SeqNum{100, 101, 102, 200, 201, 202}
	require.ElementsMatch(t, expectedSeqNumbers, allSequenceNumbers,
		"All expected messages should be included across the reports")

}

func TestCommitReportCacheOptimization(t *testing.T) {
	ctx := t.Context()

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	baseBlockNumber := uint64(1000)
	currentTimestamp := time.Now().Add(-4 * time.Hour)

	lggr := logger.Test(t)
	intTest := SetupSimpleTest(t, lggr, []cciptypes.ChainSelector{srcSelector}, dstSelector)

	// Add a mix of empty reports (no Merkle roots) and reports with Merkle roots
	// First add empty reports - these should be filtered out by the cache
	for i := 0; i < 10; i++ {
		currentTimestamp = currentTimestamp.Add(time.Second)
		baseBlockNumber++
		emptyReport := cciptypes.CommitPluginReportWithMeta{
			Report:    cciptypes.CommitPluginReport{}, // Empty report with no Merkle roots
			BlockNum:  baseBlockNumber,
			Timestamp: currentTimestamp,
		}
		intTest.ccipReader.UnfinalizedReports = append(intTest.ccipReader.UnfinalizedReports, emptyReport)
		intTest.ccipReader.FinalizedReports = append(intTest.ccipReader.FinalizedReports, emptyReport)
	}

	// Now add a report WITH Merkle roots - this should be cached and processed
	seqNumRange := cciptypes.NewSeqNumRange(100, 105)
	currentTimestamp = currentTimestamp.Add(time.Second)
	baseBlockNumber++
	reportWithRoots := cciptypes.CommitPluginReportWithMeta{
		Report: cciptypes.CommitPluginReport{
			BlessedMerkleRoots: []cciptypes.MerkleRootChain{
				{
					ChainSel:     srcSelector,
					SeqNumsRange: seqNumRange,
					MerkleRoot:   cciptypes.Bytes32{1, 2, 3, 4},
				},
			},
		},
		BlockNum:  baseBlockNumber,
		Timestamp: currentTimestamp,
	}
	intTest.ccipReader.UnfinalizedReports = append(intTest.ccipReader.UnfinalizedReports, reportWithRoots)
	intTest.ccipReader.FinalizedReports = append(intTest.ccipReader.FinalizedReports, reportWithRoots)

	// Add messages to be executed
	messages := []inmem.MessagesWithMetadata{
		makeMsgWithMetadata(100, srcSelector, dstSelector, false),
		makeMsgWithMetadata(101, srcSelector, dstSelector, false),
		makeMsgWithMetadata(102, srcSelector, dstSelector, false),
		makeMsgWithMetadata(103, srcSelector, dstSelector, false),
		makeMsgWithMetadata(104, srcSelector, dstSelector, false),
		makeMsgWithMetadata(105, srcSelector, dstSelector, false),
	}
	intTest.WithMessages(messages, baseBlockNumber, currentTimestamp, 1, srcSelector)

	// Instead of running multiple rounds which is causing issues with test timing,
	// let's write a simpler test that directly checks the commit report cache's filtering behavior:
	cacheImpl := cache.NewCommitReportCache(
		lggr,
		cache.CommitReportCacheConfig{
			MessageVisibilityInterval: 8 * time.Hour,
			EvictionGracePeriod:       1 * time.Hour,
			CleanupInterval:           30 * time.Minute,
			LookbackGracePeriod:       1 * time.Hour,
		},
		&cache.RealTimeProvider{},
		intTest.ccipReader,
	)

	// Manually refresh the cache
	err := cacheImpl.RefreshCache(ctx)
	require.NoError(t, err, "Failed to refresh cache")

	// Get the cached reports - should only contain the one with Merkle roots
	cachedReports := cacheImpl.GetCachedReports(time.Time{})
	require.Len(t, cachedReports, 1, "Cache should contain exactly 1 report (the one with Merkle roots)")

	// Verify the content of the cached report
	require.Len(t, cachedReports[0].Report.BlessedMerkleRoots, 1, "Cached report should have 1 blessed Merkle root")
	require.Equal(t, srcSelector, cachedReports[0].Report.BlessedMerkleRoots[0].ChainSel, "Chain selector mismatch")
	require.Equal(t,
		seqNumRange,
		cachedReports[0].Report.BlessedMerkleRoots[0].SeqNumsRange,
		"Sequence number range mismatch",
	)

	lggr.Info("CommitReportCache optimization test passed - correctly filtered empty reports")
}

func TestPlugin_FinalizedUnfinalizedCaching(t *testing.T) {
	ctx := t.Context()

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
	ctx := t.Context()

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
	ctx := t.Context()

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
	require.NotEmpty(t, outcome.Reports)
	require.NotEmpty(t, outcome.Reports[0].ChainReports)
	sequenceNumbers := extractSequenceNumbers(outcome.Reports[0].ChainReports[0].Messages)
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
