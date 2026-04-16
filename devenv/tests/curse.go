package tests

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

func RunCurseTests(t *testing.T, e *deployment.Environment, selectors []uint64) {
	selectorsToImpl := buildImplsMap(t, e, selectors)

	// get two distinct selectors
	fromImpl, toImpl := selectorsToImpl[selectors[0]], selectorsToImpl[selectors[1]]
	require.NotEqual(t, fromImpl, toImpl)

	curser, ok := toImpl.(testadapters.Curser)
	if !ok {
		t.Skipf("destination chain %d does not implement Curser, chain family: %s", toImpl.ChainSelector(), toImpl.Family())
	}

	// Ping the loki URL to ensure its available, else skip the test.
	err := PingLoki(t.Context(), DefaultLokiURL)
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
	alreadyCursed, err := curser.IsCursed(ctx, subject)
	require.NoError(t, err)
	if !alreadyCursed {
		t.Logf("cursing source selector %d on destination chain %d", fromImpl.ChainSelector(), toImpl.ChainSelector())
		err := curser.Curse(ctx, subject)
		require.NoError(t, err)
	} else {
		t.Logf("source selector %d is already cursed on destination chain %d, skipping curse", fromImpl.ChainSelector(), toImpl.ChainSelector())
	}

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

	// It won't get committed onchain but a buggy implementation will keep trying and reverting.
	// TODO: whats the best way to assert this? that we see e.g. one revert with an expected reason?
	// toImpl.ValidateCommit(t, fromImpl.ChainSelector(), &block, ccipocr3.NewSeqNumRange(ccipocr3.SeqNum(seqNr), ccipocr3.SeqNum(seqNr)))

	// Alternatively, if the plugin is reading the curse state correctly, we should see a log from
	// the commit plugin that the source chain is being ignored due to being cursed on the destination.
	// In this particular case, since we only have 2 chains and one is cursed, we should see the log
	// "nothing to observe from the offRamp, no active source chains exist".
	family, err := chainsel.GetSelectorFamily(toImpl.ChainSelector())
	require.NoError(t, err)

	chainID, err := chainsel.GetChainIDFromSelector(toImpl.ChainSelector())
	require.NoError(t, err)

	// Get offramp from the datastore
	offRampRef, err := e.DataStore.Addresses().Get(datastore.NewAddressRefKey(
		toImpl.ChainSelector(),
		datastore.ContractType(offramp.ContractType), // TODO: evm only?
		offramp.Version,
		"",
	))
	require.NoError(t, err)

	var (
		expectedLogger = fmt.Sprintf("CCIPCommitPlugin.%s.%s.%d.%s", family, chainID, toImpl.ChainSelector(), offRampRef.Address)
		expectedMsg    = "nothing to observe from the offRamp, no active source chains exist"
		containerName  = "don-node3" // TODO: pull dynamically"?
		logQuery       = fmt.Sprintf(
			`{container="%s"} | json | msg="%s" | logger="%s"`,
			containerName,
			expectedMsg,
			expectedLogger,
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

	// // Uncurse the source selector on the destination chain to restore uncursed state.
	// TODO: could also wait for the message to commit after being uncursed?
	// err = curser.Uncurse(ctx, subject)
	// require.NoError(t, err)
	// t.Logf("uncursed source selector %d on destination chain %d", fromImpl.ChainSelector(), toImpl.ChainSelector())
}

func selectorToSubject(selector uint64) [16]byte {
	var b [16]byte
	binary.BigEndian.PutUint64(b[8:], selector)
	return b
}
