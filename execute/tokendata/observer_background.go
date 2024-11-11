package tokendata

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type backgroundObserver struct {
	lggr              logger.Logger
	observer          TokenDataObserver
	numWorkers        int
	cachedTokenData   *inMemTokenDataCache
	msgQueue          *msgQueue
	wg                sync.WaitGroup
	done              chan struct{}
	observeTimeout    time.Duration
	reprocessInterval time.Duration // how long to wait before reprocessing a message
}

// NewBackgroundObserver initializes an observer that retrieves and caches token data in the background.
// It uses the provided observer make the actual Observe calls, storing results in memory for efficient access later.
// Goroutines are spawned to process messages concurrently, numWorkers defines how many.
// cacheExpirationInterval defines for how long in memory token data are considered active.
// cacheCleanupInterval defines how often to check and cleanup inactive data.
// observeTimeout defines how long to wait for the underlying observer to return results.
func NewBackgroundObserver(
	lggr logger.Logger,
	observer TokenDataObserver,
	numWorkers int,
	cacheExpirationInterval, cacheCleanupInterval time.Duration,
	observeTimeout time.Duration,
) TokenDataObserver {
	o := &backgroundObserver{
		lggr:       lggr,
		observer:   observer,
		numWorkers: numWorkers,
		cachedTokenData: newInMemObservationsCache(
			logger.Named(lggr, "inMemObservationsCache"),
			cacheExpirationInterval,
			cacheCleanupInterval,
		),
		msgQueue:          newMsgQueue(logger.Named(lggr, "msgQueue")),
		wg:                sync.WaitGroup{},
		done:              make(chan struct{}),
		observeTimeout:    observeTimeout,
		reprocessInterval: 5 * time.Second,
	}

	o.startWorkers()
	return o
}

// Observe fetches token data for the given messages that are already present in memory.
// If token data are not in memory, it enqueues the messages for background processing.
// Meaning that at least two calls to Observe are needed to get token data for a message.
// If data are already in the queue for processing, it is a nop.
func (o *backgroundObserver) Observe(
	_ context.Context,
	observations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	tokenDataResults := make(exectypes.TokenDataObservations)

	for chainSel, seqNumToMsg := range observations {
		for seqNum, msg := range seqNumToMsg {
			tokenData, exists := o.cachedTokenData.Get(msg.Header.MessageID)
			if exists && !tokenData.IsReady() {
				return nil, fmt.Errorf("internal error, cache contains not ready token data")
			}

			lggr := logger.With(o.lggr, "msgID", msg.Header.MessageID.String(),
				"sourceChain", chainSel.String(), "seqNum", seqNum.String())

			if exists {
				// token data exist so include them in the results
				if _, ok := tokenDataResults[chainSel]; !ok { // initialize this chain if not exists
					tokenDataResults[chainSel] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
				}
				lggr.Infow("token data found in cache")
				tokenDataResults[chainSel][seqNum] = tokenData
			} else {
				// token data not in cache for this message, enqueue the message
				if ok := o.msgQueue.addMsg(msg, 0); ok {
					lggr.Infow("message added to the queue")
				} else {
					lggr.Infow("message already exists in the queue")
				}
			}
		}
	}

	return tokenDataResults, nil
}

// IsTokenSupported simply forwards the call to the underlying observer.
func (o *backgroundObserver) IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool {
	return o.observer.IsTokenSupported(sourceChain, msgToken)
}

func (o *backgroundObserver) startWorkers() {
	o.lggr.Info("waiting for existing (if any) workers to stop")
	o.wg.Wait()
	o.lggr.Info("all workers stopped, new workers are starting...")

	for i := 0; i < o.numWorkers; i++ {
		o.wg.Add(1)
		workerID := i
		go o.worker(workerID)
	}
}

// worker is a goroutine that processes messages from the queue.
// It will stop using the o.done channel.
func (o *backgroundObserver) worker(id int) {
	lggr := logger.With(o.lggr, "workerID", id)
	lggr.Info("worker started")

	defer o.wg.Done()
	for {
		select {
		case <-o.done:
			lggr.Info("worker stopped after receiving done signal")
			return
		case <-o.msgQueue.newMsgSignalChan:
			lggr.Debug("new job signal received")

			msg, ok := o.msgQueue.pop()
			if !ok {
				lggr.Debug("nothing to work on, waiting for new job signal")
				continue
			}

			lggr := logger.With(lggr,
				"msgID", msg.Header.MessageID.String(),
				"sourceChain", msg.Header.SourceChainSelector.String(),
				"seqNum", msg.Header.SequenceNumber.String(),
				"numTokens", len(msg.TokenAmounts),
			)
			lggr.Infow("processing message")

			// observe only this single message and use a timeout for the observation
			observationTimeoutCtx, cancel := context.WithTimeout(context.Background(), o.observeTimeout)
			tokenData, err := o.observer.Observe(observationTimeoutCtx,
				map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					msg.Header.SourceChainSelector: {msg.Header.SequenceNumber: msg},
				})
			cancel()

			if err != nil {
				lggr.Errorw("message observation failed, message pushed again to the queue", "err", err)
				o.msgQueue.addMsg(msg, o.reprocessInterval)
				continue
			}

			if _, chainExists := tokenData[msg.Header.SourceChainSelector]; !chainExists {
				lggr.Errorw("underlying observer did not return token data for the chain")
				o.msgQueue.addMsg(msg, o.reprocessInterval)
				continue
			}

			if _, seqExists := tokenData[msg.Header.SourceChainSelector][msg.Header.SequenceNumber]; !seqExists {
				lggr.Errorw("underlying observer did not return token data for the sequence number")
				o.msgQueue.addMsg(msg, o.reprocessInterval)
				continue
			}

			if !tokenData[msg.Header.SourceChainSelector][msg.Header.SequenceNumber].IsReady() {
				lggr.Infow("token data not ready by the underlying observer, message pushed again to the queue")
				o.msgQueue.addMsg(msg, o.reprocessInterval)
				continue
			}

			lggr.Infow("message observation successful, token data cached")
			o.cachedTokenData.Set(
				msg.Header.MessageID,
				tokenData[msg.Header.SourceChainSelector][msg.Header.SequenceNumber],
			)
		}
	}

}

// msgQueue is a simple in-memory queue that can be used for async message processing.
type msgQueue struct {
	lggr             logger.Logger
	msgs             []msgWithInfo
	mu               *sync.RWMutex
	newMsgSignalChan chan struct{}
}

type msgWithInfo struct {
	msg cciptypes.Message
	// availableAt is the time when the message can be processed
	availableAt time.Time
	// enqueuedAt is the time when the message was added to the queue
	enqueuedAt time.Time
}

func newMsgQueue(lggr logger.Logger) *msgQueue {
	return &msgQueue{
		lggr:             lggr,
		msgs:             make([]msgWithInfo, 0),
		mu:               &sync.RWMutex{},
		newMsgSignalChan: make(chan struct{}),
	}
}

func (q *msgQueue) addMsg(msg cciptypes.Message, availableAfter time.Duration) bool {
	lggr := logger.With(
		q.lggr, "msgID", msg.Header.MessageID.String(),
		"sourceChain", msg.Header.SourceChainSelector.String(),
		"seqNum", msg.Header.SequenceNumber.String(),
		"availableAfter", availableAfter,
	)
	lggr.Debug("waiting for the lock before adding msg")

	q.mu.RLock()
	containsMsg := q.containsMsg(msg)
	q.mu.RUnlock()
	if containsMsg {
		lggr.Debug("message already exists in the queue")
		return false
	}

	q.mu.Lock()
	q.msgs = append(q.msgs, msgWithInfo{
		msg:         msg,
		availableAt: time.Now().Add(availableAfter).UTC(),
		enqueuedAt:  time.Now().UTC(),
	})
	lggr.Debugw("message added to the queue, new msg signal sent", "numMsgs", len(q.msgs))
	q.mu.Unlock()

	lggr.Debugw("sending to new msg signal channel")
	q.newMsgSignalChan <- struct{}{}
	return true
}

// pop returns the first available message from the queue
func (q *msgQueue) pop() (cciptypes.Message, bool) {
	q.lggr.Debug("waiting for the lock before popping msg")

	q.mu.Lock()
	defer q.mu.Unlock()

	q.lggr.Debug("lock acquired before popping msg")

	if len(q.msgs) == 0 {
		q.lggr.Debug("no messages in the queue")
		return cciptypes.Message{}, false
	}

	for i, msg := range q.msgs {
		if time.Now().UTC().After(msg.availableAt) {
			q.lggr.Debugw("message popped from the queue",
				"msgID", msg.msg.Header.MessageID.String(),
				"sourceChain", msg.msg.Header.SourceChainSelector.String(),
				"seqNum", msg.msg.Header.SequenceNumber.String(),
				"enqueuedSince", time.Since(msg.enqueuedAt),
				"availableSince", time.Since(msg.availableAt),
			)
			q.msgs = append(q.msgs[:i], q.msgs[i+1:]...)
			return msg.msg, true
		}
	}

	return cciptypes.Message{}, false
}

func (q *msgQueue) containsMsg(msg cciptypes.Message) bool {
	for _, qMsg := range q.msgs {
		if qMsg.msg.Header.MessageID == msg.Header.MessageID {
			q.lggr.Debugw("message already exists in the queue", "msg", msg, "qMsg", qMsg)
			return true
		}
	}

	return false
}

type inMemTokenDataCache struct {
	lggr               logger.Logger
	expirationInterval time.Duration
	inMemTokenData     map[cciptypes.Bytes32]exectypes.MessageTokenData
	expiresAt          map[cciptypes.Bytes32]time.Time
	mu                 *sync.RWMutex
}

// newInMemObservationsCache initializes an in-memory cache for token data.
// It uses a background goroutine to periodically check and remove expired data.
// cleanupInterval specifies the frequency for checking and cleaning up inactive data.
// Setting a low value is discouraged, as the cleanup process holds a lock.
func newInMemObservationsCache(
	lggr logger.Logger, expirationInterval, cleanupInterval time.Duration) *inMemTokenDataCache {
	c := &inMemTokenDataCache{
		lggr:               lggr,
		expirationInterval: expirationInterval,
		inMemTokenData:     make(map[cciptypes.Bytes32]exectypes.MessageTokenData),
		expiresAt:          make(map[cciptypes.Bytes32]time.Time),
		mu:                 &sync.RWMutex{},
	}
	c.runExpirationLoop(cleanupInterval)
	return c
}

func (c *inMemTokenDataCache) Get(msgID cciptypes.Bytes32) (exectypes.MessageTokenData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	msgData, ok := c.inMemTokenData[msgID]
	if !ok {
		return exectypes.MessageTokenData{}, false
	}

	return msgData, true
}

func (c *inMemTokenDataCache) Set(msgID cciptypes.Bytes32, tokenData exectypes.MessageTokenData) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.inMemTokenData[msgID] = tokenData
	c.expiresAt[msgID] = time.Now().Add(c.expirationInterval).UTC()
	c.lggr.Debugw("token data cached", "msgID", msgID, "expiresAt", c.expiresAt[msgID])
}

func (c *inMemTokenDataCache) runExpirationLoop(cleanupInterval time.Duration) {
	go func() {
		ticker := time.NewTicker(cleanupInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				func() {
					c.mu.Lock()
					defer c.mu.Unlock()

					for msgID, expiresAt := range c.expiresAt {
						if c.hasExpired(msgID) {
							c.lggr.Debugw("token data expired and removed from cache",
								"msgID", msgID.String(),
								"expiresAt", expiresAt,
								"now", time.Now().UTC(),
							)

							delete(c.inMemTokenData, msgID)
							delete(c.expiresAt, msgID)
						}
					}
				}()
			}
		}
	}()
}

func (c *inMemTokenDataCache) hasExpired(msgID cciptypes.Bytes32) bool {
	expiresAt, ok := c.expiresAt[msgID]
	if !ok {
		// if the data is not in the cache, it is considered expired
		return true
	}

	return time.Now().UTC().After(expiresAt)
}
