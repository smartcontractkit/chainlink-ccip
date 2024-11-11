package tokendata

import (
	"fmt"
	rand2 "math/rand"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_backgroundObserver(t *testing.T) {
	ctx := tests.Context(t)
	lggr := mocks.NullLogger

	baseObserver := &NoopTokenDataObserver{}
	numWorkers := 10
	cacheExpirationInterval := 10 * time.Minute
	cacheCleanupInterval := 15 * time.Minute
	observeTimeout := time.Second

	msgsPerChain := map[cciptypes.ChainSelector]int{
		1000: 1 + rand2.Intn(10),
		2000: 1 + rand2.Intn(20),
		3000: 1 + rand2.Intn(30),
		4000: 1 + rand2.Intn(40),
	}

	observer := NewBackgroundObserver(
		lggr,
		baseObserver,
		numWorkers,
		cacheExpirationInterval,
		cacheCleanupInterval,
		observeTimeout,
	)

	msgIDs := mapset.NewSet[cciptypes.Bytes32]()

	msgObservations := exectypes.MessageObservations{}
	for chain, numMsgs := range msgsPerChain {
		msgObservations[chain] = make(map[cciptypes.SeqNum]cciptypes.Message, numMsgs)
		for i := 0; i < numMsgs; i++ {
			seqNum := cciptypes.SeqNum(i)
			msgIDStr := fmt.Sprintf("%d-%d", chain, i)
			msgID := cciptypes.Bytes32{}
			copy(msgID[:], msgIDStr)

			msgObservations[chain][seqNum] = cciptypes.Message{
				Header: cciptypes.RampMessageHeader{
					SequenceNumber:      cciptypes.SeqNum(i),
					SourceChainSelector: chain,
					MessageID:           msgID,
				},
				TokenAmounts: []cciptypes.RampTokenAmount{
					{},
				},
			}

			if msgIDs.Contains(msgID) {
				t.Fatalf("duplicate msgID: %s", msgID)
			}
			msgIDs.Add(msgID)
		}
	}

	tokenDataObservations, err := observer.Observe(ctx, msgObservations)
	require.NoError(t, err)
	require.Equal(t, len(msgObservations), len(tokenDataObservations))
	for chain, seqNums := range tokenDataObservations {
		require.Equal(t, len(msgObservations[chain]), len(seqNums))
		for _, tokenData := range seqNums {
			for _, td := range tokenData.TokenData {
				require.Empty(t, td)
			}
		}
	}

	require.Eventually(t, func() bool {
		tokenDataObservations, err = observer.Observe(ctx, msgObservations)
		require.NoError(t, err)
		// send another request to make sure it's idempotent
		tokenDataObservations, err = observer.Observe(ctx, msgObservations)
		require.NoError(t, err)

		t.Logf("len(tokenDataObservations): %d", len(tokenDataObservations))
		t.Logf("len(msgObservations): %d", len(msgObservations))
		if len(msgObservations) != len(tokenDataObservations) {
			return false
		}

		for chain, msgs := range msgObservations {
			t.Logf("chain: %d, len(msgs): %d, len(tokenDataObservations[chain]): %d",
				chain, len(msgs), len(tokenDataObservations[chain]))
			if len(msgs) != len(tokenDataObservations[chain]) {
				return false
			}

			for chain, seqNums := range tokenDataObservations {
				require.Equal(t, len(msgObservations[chain]), len(seqNums))
				for _, tokenData := range seqNums {
					for _, td := range tokenData.TokenData {
						if td.Data == nil {
							return false
						}
					}
				}
			}
		}

		return true
	}, tests.WaitTimeout(t), 50*time.Millisecond)

	// test expiration
	rawObserver := observer.(*backgroundObserver)
	// keep only len(chains) messages in the cache
	msgsToKeep := len(msgsPerChain)
	i := 0
	for msgID := range rawObserver.cachedTokenData.inMemTokenData {
		if i < msgsToKeep {
			i++
			continue
		}
		rawObserver.cachedTokenData.expiresAt[msgID] = time.Now()
	}

	// run another expiration loop to remove expired messages
	rawObserver.cachedTokenData.runExpirationLoop(10 * time.Millisecond)

	require.Eventually(t, func() bool {
		rawObserver.cachedTokenData.mu.RLock()
		totalMsgs := len(rawObserver.cachedTokenData.inMemTokenData)
		rawObserver.cachedTokenData.mu.RUnlock()
		t.Logf("totalMsgs: %d", totalMsgs)
		t.Logf("msgsToKeep: %d", msgsToKeep)
		return msgsToKeep == totalMsgs
	}, tests.WaitTimeout(t), 50*time.Millisecond)

	// graceful shutdown
	rawObserver.Close()
}
