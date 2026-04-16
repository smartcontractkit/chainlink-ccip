package tests

import (
	"encoding/binary"
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/deployment/testadapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"
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

	// curse the source selector on the destination chain
	subject := selectorToSubject(fromImpl.ChainSelector())
	err := curser.Curse(t.Context(), subject)
	require.NoError(t, err)

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

	seqNr, messageID, err := fromImpl.SendMessage(t.Context(), toImpl.ChainSelector(), msg)
	require.NoError(t, err)
	t.Logf("sent message with seqNr: %d and messageID: %s", seqNr, messageID)

	// It won't get committed onchain but a buggy implementation will keep trying and reverting.
	// TODO: whats the best way to assert this? that we see e.g. one revert with an expected reason?
	// toImpl.ValidateCommit(t, fromImpl.ChainSelector(), &block, ccipocr3.NewSeqNumRange(ccipocr3.SeqNum(seqNr), ccipocr3.SeqNum(seqNr)))

	// Alternatively, if the plugin is reading the curse state correctly, we should see a log from
	// the commit plugin that the source chain is being ignored due to being cursed on the destination.
	// In this particular case, since we only have 2 chains and one is cursed, we should see the log
	// "nothing to observe from the offRamp, no active source chains exist".
}

func selectorToSubject(selector uint64) [16]byte {
	var b [16]byte
	binary.BigEndian.PutUint64(b[8:], selector)
	return b
}
