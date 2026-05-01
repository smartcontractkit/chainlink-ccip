package tests

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// RunCurseTests runs the curse test suite using the given environment and selectors.
//
// For example, to run this in devenv on EVM, you can do the following:
//
// 1. Run "ccip up env.toml,env-anvil.toml,env-cl-rebuild.toml"
//
// 2. Run "ccip obs up" as Loki is needed when asserting logs in the test.
//
// 3. Run "go test -v -run "TestE2ECurse" ./tests/e2e/"
func RunCurseTests(t *testing.T, e *deployment.Environment, selectors []uint64) {
	selectorsToImpl := buildImplsMap(t, e, selectors)

	// get two distinct selectors
	fromImpl, toImpl := selectorsToImpl[selectors[0]], selectorsToImpl[selectors[1]]
	require.NotEqual(t, fromImpl, toImpl)

	destFamily, err := chainsel.GetSelectorFamily(toImpl.ChainSelector())
	require.NoError(t, err)

	curseAdapter, ok := fastcurse.GetCurseRegistry().GetCurseAdapter(destFamily, semver.MustParse("1.6.0"))
	if !ok {
		t.Skipf("no curse adapter registered for chain family: %s and version: %s", destFamily, semver.MustParse("1.6.0"))
	}
	require.NotNil(t, curseAdapter, "registered curse adapter is nil")

	require.NoError(t, curseAdapter.Initialize(*e, toImpl.ChainSelector()))

	// Ping the loki URL to ensure its available, else skip the test.
	err = PingLoki(t.Context(), DefaultLokiURL)
	if err != nil {
		t.Skipf("failed to ping Loki WebSocket URL, is loki running?: %v", err)
	}

	// Need to wrap the ctx with the zerolog logger to get the
	// logs from the adapter in the test output.
	ctx := log.
		Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		Level(zerolog.InfoLevel).
		WithContext(t.Context())

	// curse the source selector on the destination chain if its not already cursed.
	subject := selectorToSubject(fromImpl.ChainSelector())
	alreadyCursed, err := curseAdapter.IsSubjectCursedOnChain(*e, toImpl.ChainSelector(), subject)
	require.NoError(t, err)
	if !alreadyCursed {
		t.Logf("cursing source selector %d on destination chain %d", fromImpl.ChainSelector(), toImpl.ChainSelector())
		curseCS := fastcurse.CurseChangeset(fastcurse.GetCurseRegistry(), changesets.GetRegistry())
		output, err := curseCS.Apply(*e, fastcurse.RMNCurseConfig{
			CurseActions: []fastcurse.CurseActionInput{
				{
					ChainSelector:        toImpl.ChainSelector(),
					SubjectChainSelector: fromImpl.ChainSelector(),
					Version:              semver.MustParse("1.6.0"),
					IsGlobalCurse:        false,
				},
			},
		})
		require.NoError(t, err)
		require.Greater(t, len(output.Reports), 0)
	} else {
		t.Logf("source selector %d is already cursed on destination chain %d, skipping curse", fromImpl.ChainSelector(), toImpl.ChainSelector())
	}

	// Confirm that the subject is cursed on the destination chain.
	isCursed, err := curseAdapter.IsSubjectCursedOnChain(*e, toImpl.ChainSelector(), subject)
	require.NoError(t, err)
	require.True(t, isCursed, "subject should be cursed on destination chain")

	// Send a message from the source chain to the destination chain
	// The plugin should ignore the message because the dest is cursing the source selector.
	// block := toImpl.CurrentBlock(t)
	msg, err := fromImpl.BuildMessage(testadapters.MessageComponents{
		DestChainSelector: toImpl.ChainSelector(),
		Receiver:          toImpl.CCIPReceiver(),
		Data:              []byte("should not commit"),
		FeeToken:          "",
		ExtraArgs:         nil,
		TokenAmounts:      nil,
	})
	require.NoError(t, err)

	seqNr, messageID, err := fromImpl.SendMessage(ctx, toImpl.ChainSelector(), msg)
	require.NoError(t, err)
	t.Logf("sent message with seqNr: %d and messageID: %s", seqNr, messageID)

	// Assertion: If the plugin is reading the curse state correctly, we should see a log from
	// the commit plugin that the source chain is being ignored due to being cursed on the destination.
	// In this particular case, since we only have 2 chains and one is cursed, we should see the log
	// "nothing to observe from the offRamp, no active source chains exist".
	var (
		expectedMsg   = "nothing to observe from the offRamp, no active source chains exist"
		containerName = "don-node3" // TODO: pull dynamically"?
		logQuery      = fmt.Sprintf(
			`{container="%s"} | json | caller="merkleroot/observation.go:520" | msg="%s" | component="MerkleRoot"`,
			containerName,
			expectedMsg,
		)
	)

	t.Logf("waiting for log line with query: %s to appear in Loki", logQuery)
	logLine, err := WaitForLokiLogs(
		ctx,
		DefaultLokiURL,
		logQuery,
	)
	require.NoError(t, err)

	t.Logf("found log line asserting that the source chain is being ignored due to being cursed on the destination: %s", logLine)

	// TODO: could uncurse and wait for the message to commit/execute?
}

func selectorToSubject(selector uint64) [16]byte {
	var b [16]byte
	binary.BigEndian.PutUint64(b[8:], selector)
	return b
}
