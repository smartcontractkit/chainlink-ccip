package tests

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
)

func RunQAInteractiveTests(t *testing.T, e *deployment.Environment,
	tonSelector, evmSelector uint64) {

	getAdapter := func(selector uint64) ccip.CCIP16ProductConfiguration {
		family, err := chainsel.GetSelectorFamily(selector)
		require.NoError(t, err)
		chainID, err := chainsel.GetChainIDFromSelector(selector)
		require.NoError(t, err)
		adapter, err := ccip.NewCCIPImplFromNetwork(family, chainID)
		require.NoError(t, err)
		adapter.SetCLDF(e)
		return adapter
	}

	tonAdapter := getAdapter(tonSelector)
	evmAdapter := getAdapter(evmSelector)

	type chainPair struct {
		src  ccip.CCIP16ProductConfiguration
		dest ccip.CCIP16ProductConfiguration
	}
	type testCase struct {
		name     string
		lanes    []chainPair
		testFunc func(t *testing.T, lane chainPair)
	}
	testCases := []testCase{
		{
			"message to eoa",
			[]chainPair{
				{evmAdapter, tonAdapter},
			},
			func(t *testing.T, lane chainPair) {
				receiver := lane.dest.EOAReceiver(t)
				extraArgs, err := lane.dest.GetExtraArgs(receiver, lane.src.Family())
				require.NoError(t, err)

				msg, err := lane.src.BuildMessage(testadapters.MessageComponents{
					Receiver:          receiver,
					Data:              []byte("hello eoa"),
					FeeToken:          "",
					ExtraArgs:         extraArgs,
					TokenAmounts:      nil,
					DestChainSelector: lane.dest.ChainSelector(),
				})
				require.NoError(t, err)

				seqNr, msgID, err := lane.src.SendMessage(t.Context(), lane.dest.ChainSelector(), msg)
				t.Logf("sendMsgRequireNoError got msgID: %s", msgID)
				require.NoError(t, err, "failed to send message")
				seqNumRange := ccipocr3.NewSeqNumRange(ccipocr3.SeqNum(seqNr), ccipocr3.SeqNum(seqNr))
				lane.dest.ValidateCommit(t, lane.src.ChainSelector(), nil, seqNumRange)
				waitForUserAction(t, fmt.Sprintf("Check that message was received but is stuck in \"In Progress\". Then press ENTER\n%s", msgInfo(msgID)))
			},
		},
		{
			"not enough gas; manual re-exec",
			[]chainPair{
				{tonAdapter, evmAdapter},
				{evmAdapter, tonAdapter},
			},
			func(t *testing.T, lane chainPair) {
				receiver := lane.dest.CCIPReceiver()

				extraArgs, err := lane.dest.GetExtraArgs(receiver, lane.src.Family(), testadapters.NewGasLimitExtraArg(lane.dest.LowGasLimit()))
				require.NoError(t, err)

				msg, err := lane.src.BuildMessage(testadapters.MessageComponents{
					Receiver:          receiver,
					Data:              []byte("hello world"),
					ExtraArgs:         extraArgs,
					DestChainSelector: lane.dest.ChainSelector(),
				})
				require.NoError(t, err)

				_, msgID := sendMsgRequireErrorOnDestChain(t, lane.src, lane.dest, msg)

				waitForUserAction(t, fmt.Sprintf("Check that message was received but is on \"Manual Exec\". Execute manually with higher gas limit TODO and press ENTER\n%s", msgInfo(msgID)))
			},
		},
		{
			"receiver fails; manual re-exec",
			[]chainPair{
				{evmAdapter, tonAdapter},
			},
			func(t *testing.T, lane chainPair) {
				isRejectAll := false
				setReceiverRejectAll := func(rejectAll bool) {
					err := lane.dest.SetReceiverRejectAll(t.Context(), t, rejectAll)
					require.NoError(t, err)
					isRejectAll = rejectAll
					t.Cleanup(func() {
						if !isRejectAll {
							return
						}
						err := lane.dest.SetReceiverRejectAll(t.Context(), t, false)
						require.NoError(t, err, "failed to reset receiver reject all flag, manual cleanup may be required")
					})
				}
				receiver := lane.dest.CCIPReceiver()

				setReceiverRejectAll(true)

				extraArgs, err := lane.dest.GetExtraArgs(receiver, lane.src.Family())
				require.NoError(t, err)

				msg, err := lane.src.BuildMessage(testadapters.MessageComponents{
					Receiver:          receiver,
					Data:              []byte("hello world"),
					ExtraArgs:         extraArgs,
					DestChainSelector: lane.dest.ChainSelector(),
				})
				require.NoError(t, err)

				seqNr, msgID := sendMsgRequireErrorOnDestChain(t, lane.src, lane.dest, msg)

				setReceiverRejectAll(false)

				// Assuming the error is due to not enough gas, we can attempt a manual re-execution with a higher gas limit

				waitForExec := make(chan struct{})
				go func() {
					defer close(waitForExec)
					startBlock := lane.dest.CurrentBlock(t) + 1
					lane.dest.ValidateExecSucceeds(t, lane.src.ChainSelector(), &startBlock, []uint64{seqNr})
				}()

				waitForUserAction(t, fmt.Sprintf("Check that message was received but is on \"Manual Exec\". Execute manually and press ENTER\n%s", msgInfo(msgID)))

				<-waitForExec
			},
		},
	}

	for _, tc := range testCases {
		for _, lane := range tc.lanes {
			laneTag := fmt.Sprintf("%s->%s", lane.src.Family(), lane.dest.Family())
			t.Run(fmt.Sprintf("%s:%s", laneTag, tc.name), func(t *testing.T) {
				tc.testFunc(t, lane)
			})
		}
	}
}

const (
	RESET     = "\033[0m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
)

func waitForUserAction(t *testing.T, prompt string) {
	t.Helper()
	fmt.Printf("\a\n%s%s>>> ACTION REQUIRED: %s%s\nPress ENTER when done... ", BOLD, YELLOW, prompt, RESET)
	_, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func msgInfo(msgID string) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprint(RESET, "Message ID: ", BOLD, msgID, "\n"))
	builder.WriteString(fmt.Sprint(RESET, "Explorer: ", BLUE, UNDERLINE, ccipExplorer(msgID), "\n"))
	builder.WriteString(RESET)
	return builder.String()
}

func ccipExplorer(msgID string) string {
	return urlFmt("https://ccip.chain.link/#/side-drawer/msg/%s", msgID)
}

func urlFmt(format string, a ...any) string {
	url := fmt.Sprintf(format, a...)
	return fmt.Sprintf("%s%s%s", BLUE, url, RESET)
}

func sendMsgRequireErrorOnDestChain(t *testing.T, fromImpl, toImpl ccip.CCIP16ProductConfiguration, msg any) (uint64, string) {
	t.Helper()
	startBlock := toImpl.CurrentBlock(t) + 1
	seqNr, messageID, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
	require.NoError(t, err, "sendMsgRequireErrorOnDestChain failed to send message")
	t.Logf("sendMsgRequireErrorOnDestChain got messageID: %s", messageID)

	require.NoError(t, err)
	seqNrUint := ccipocr3.SeqNum(seqNr)
	seqNumRange := ccipocr3.NewSeqNumRange(seqNrUint, seqNrUint)
	toImpl.ValidateCommit(t, fromImpl.ChainSelector(), &startBlock, seqNumRange)
	toImpl.ValidateExecFails(t, fromImpl.ChainSelector(), &startBlock, []uint64{seqNr})

	return seqNr, messageID
}
