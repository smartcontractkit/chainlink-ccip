package observer

import (
	"fmt"
	rand2 "math/rand"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_backgroundObserver(t *testing.T) {
	ctx := tests.Context(t)
	lggr := mocks.NullLogger

	numWorkers := 10
	cacheExpirationInterval := 10 * time.Minute
	cacheCleanupInterval := 15 * time.Minute
	observeTimeout := time.Second

	numMsgsPerChain := map[cciptypes.ChainSelector]int{
		1000: 1 + rand2.Intn(10),
		2000: 1 + rand2.Intn(20),
		3000: 1 + rand2.Intn(30),
		4000: 1 + rand2.Intn(40),
	}
	// generate the msg observations
	msgObservations, errorMsgs := generateMsgObservations(numMsgsPerChain)

	errorTokenData := map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]int{}
	for chain, seqNums := range errorMsgs {
		errorTokenData[chain] = map[cciptypes.SeqNum][]int{}
		for _, seqNum := range seqNums {
			errorTokenData[chain][seqNum] = []int{1} // mark only the second token data as error
		}
	}

	baseObserver := &NoopTokenDataObserver{tokenSupported: true, errorTokenData: errorTokenData}
	observer := NewBackgroundObserver(
		lggr,
		baseObserver,
		numWorkers,
		cacheExpirationInterval,
		cacheCleanupInterval,
		observeTimeout,
	)

	// call initial Observe and assert that all token data are empty since jobs were just scheduled
	tokenDataObservations, err := observer.Observe(ctx, msgObservations)
	require.NoError(t, err)
	assertTokenDataEmptyInitially(t, msgObservations, tokenDataObservations)

	// call Observe again until all data are present - the NoOp base observer simply sets ready to true
	require.Eventually(t, func() bool {
		tokenDataObservations, err = observer.Observe(ctx, msgObservations)
		require.NoError(t, err)

		// send another request to make sure it's idempotent
		tokenDataObservations, err = observer.Observe(ctx, msgObservations)
		require.NoError(t, err)

		// make sure all token data observations are present
		if len(msgObservations) != len(tokenDataObservations) {
			return false
		}
		for chain, seqNums := range tokenDataObservations {
			if len(msgObservations[chain]) != len(seqNums) {
				return false
			}

			errorMsgsOfChain := errorMsgs[chain]

			for seqNum, tokenData := range seqNums {
				if slices.Contains(errorMsgsOfChain, seqNum) {
					if len(msgObservations[chain][seqNum].TokenAmounts) != len(tokenData.TokenData) {
						t.Errorf("token data should have an error but token data seem ready for the msg")
					}
				} else {
					if len(msgObservations[chain][seqNum].TokenAmounts) != len(tokenData.TokenData) {
						return false
					}
					for _, td := range tokenData.TokenData {
						if !td.Ready {
							return false
						}
					}
				}
			}
		}
		return true
	}, tests.WaitTimeout(t), 50*time.Millisecond)

	testCacheExpirationAndShutdown(t, observer.(*backgroundObserver), numMsgsPerChain)
}

func assertTokenDataEmptyInitially(t *testing.T,
	msgObservations exectypes.MessageObservations, tokenDataObservations exectypes.TokenDataObservations) {
	require.Equal(t, len(msgObservations), len(tokenDataObservations))
	for chain, seqNums := range tokenDataObservations {
		require.Equal(t, len(msgObservations[chain]), len(seqNums))
		for _, tokenData := range seqNums {
			for _, td := range tokenData.TokenData {
				require.Equal(t, exectypes.TokenData{Supported: true}, td)
			}
		}
	}
}

func testCacheExpirationAndShutdown(
	t *testing.T, rawObserver *backgroundObserver, numMsgsPerChain map[cciptypes.ChainSelector]int) {
	// keep only len(chains) messages in the cache
	msgsToKeep := len(numMsgsPerChain)
	i := 0
	rawObserver.cachedTokenData.mu.Lock()
	for msgID := range rawObserver.cachedTokenData.inMemTokenData {
		if i < msgsToKeep {
			i++
			continue
		}
		rawObserver.cachedTokenData.expiresAt[msgID] = time.Now()
	}
	rawObserver.cachedTokenData.mu.Unlock()
	// run another expiration loop to remove expired messages
	rawObserver.cachedTokenData.runExpirationLoop(time.Millisecond)

	require.Eventually(t, func() bool {
		rawObserver.cachedTokenData.mu.RLock()
		totalMsgs := len(rawObserver.cachedTokenData.inMemTokenData)
		rawObserver.cachedTokenData.mu.RUnlock()
		return msgsToKeep == totalMsgs
	}, tests.WaitTimeout(t), 50*time.Millisecond)

	// graceful shutdown
	assert.NoError(t, rawObserver.Close())
}

func generateMsgObservations(
	numMsgsPerChain map[cciptypes.ChainSelector]int,
) (exectypes.MessageObservations, map[cciptypes.ChainSelector][]cciptypes.SeqNum) {
	msgObservations := exectypes.MessageObservations{}
	errorMsgs := make(map[cciptypes.ChainSelector][]cciptypes.SeqNum)
	for chain, numMsgs := range numMsgsPerChain {
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
					{
						SourcePoolAddress: rand.RandomBytes(32),
						DestTokenAddress:  rand.RandomBytes(32),
						ExtraData:         rand.RandomBytes(32),
						Amount:            cciptypes.NewBigIntFromInt64(123),
					},
					{
						SourcePoolAddress: rand.RandomBytes(32),
						DestTokenAddress:  rand.RandomBytes(32),
						ExtraData:         rand.RandomBytes(32),
						Amount:            cciptypes.NewBigIntFromInt64(123),
					},
				},
			}

			// 10% chance of this msg token data to observer to return an error
			isTokenDataError := rand2.Intn(10) == 0
			if isTokenDataError {
				if _, exists := errorMsgs[chain]; !exists {
					errorMsgs[chain] = []cciptypes.SeqNum{}
				}
				errorMsgs[chain] = append(errorMsgs[chain], seqNum)
			}
		}
	}

	return msgObservations, errorMsgs
}
