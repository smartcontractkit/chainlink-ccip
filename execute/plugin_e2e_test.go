package execute

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks/inmem"
)

func TestPlugin(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	srcSelector := cciptypes.ChainSelector(1)
	dstSelector := cciptypes.ChainSelector(2)

	messages := []inmem.MessagesWithMetadata{
		makeMsg(100, srcSelector, dstSelector, true),
		makeMsg(101, srcSelector, dstSelector, true),
		makeMsg(102, srcSelector, dstSelector, false),
		makeMsg(103, srcSelector, dstSelector, false),
		makeMsg(104, srcSelector, dstSelector, false),
		makeMsg(105, srcSelector, dstSelector, false),
	}

	runner, server := SetupSimpleTest(ctx, t, lggr, srcSelector, dstSelector, messages)
	defer server.Close()

	// Contract Discovery round.
	res, err := runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err := exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Equal(t, exectypes.Initialized, outcome.State)

	// Round 1 - Get Commit Reports
	// One pending commit report only.
	// Two of the messages are executed which should be indicated in the Outcome.
	res, err = runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err = exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)
	require.ElementsMatch(t, outcome.PendingCommitReports[0].ExecutedMessages, []cciptypes.SeqNum{100, 101})

	// Round 2 - Get Messages
	// Messages now attached to the pending commit.
	res, err = runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err = exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	require.Len(t, outcome.Report.ChainReports, 0)
	require.Len(t, outcome.PendingCommitReports, 1)

	// Round 3 - Filter
	// An execute report with the following messages executed: 102, 103, 104, 105.
	res, err = runner.RunRound(ctx)
	require.NoError(t, err)
	outcome, err = exectypes.DecodeOutcome(res.Outcome)
	require.NoError(t, err)
	sequenceNumbers := slicelib.Map(outcome.Report.ChainReports[0].Messages, func(m cciptypes.Message) cciptypes.SeqNum {
		return m.Header.SequenceNumber
	})
	require.ElementsMatch(t, sequenceNumbers, []cciptypes.SeqNum{102, 103, 104, 105})

}
