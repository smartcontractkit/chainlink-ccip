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
	lggr            logger.Logger
	observer        TokenDataObserver
	cachedTokenData *inMemTokenDataCache
	numWorkers      int
	wg              sync.WaitGroup
	msgQueue        *msgQueue
	done            chan struct{}
}

// NewBackgroundObserver creates a new observer that observes token data in the background.
// It uses the provided observer to fetch token data in the background while caching the results in memory.
func NewBackgroundObserver(
	lggr logger.Logger,
	observer TokenDataObserver,
	numWorkers int,
	cacheExpirationInterval, cacheCleanupInterval time.Duration,
) TokenDataObserver {
	o := &backgroundObserver{
		lggr:            lggr,
		observer:        observer,
		cachedTokenData: newInMemObservationsCache(cacheExpirationInterval, cacheCleanupInterval),
		numWorkers:      numWorkers,
		wg:              sync.WaitGroup{},
		msgQueue:        newMsgQueue(),
		done:            make(chan struct{}),
	}

	o.startWorkers()
	return o
}

func (o *backgroundObserver) Observe(ctx context.Context, observations exectypes.MessageObservations) (exectypes.TokenDataObservations, error) {
	tokenDataResults := make(exectypes.TokenDataObservations)

	for chainSel, seqNumToMsg := range observations {
		for seqNum := range seqNumToMsg {
			tokenData, exists := o.cachedTokenData.Get(chainSel, seqNum)

			if exists && !tokenData.IsReady() {
				return nil, fmt.Errorf("internal error, cache contains not ready token data")
			}

			// include cached token data to result set if it exists
			if exists {
				if _, ok := tokenDataResults[chainSel]; !ok {
					tokenDataResults[chainSel] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
				}
				tokenDataResults[chainSel][seqNum] = tokenData
			}

			// add new msgs to the queue for processing
			if !exists {
				for _, msg := range seqNumToMsg {
					if ok := o.msgQueue.addJob(msg); ok {
						o.lggr.Infow("Added msg to queue", "msgID", msg.Header.MessageID.String())
					} else {
						o.lggr.Infow("Job already exists in queue", "msgID", msg.Header.MessageID.String())
					}
				}
			}
		}
	}

	return tokenDataResults, nil
}

func (o *backgroundObserver) IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool {
	return o.observer.IsTokenSupported(sourceChain, msgToken)
}

func (o *backgroundObserver) startWorkers() {
	for i := 0; i < o.numWorkers; i++ {
		o.wg.Add(1)
		workerID := i
		go o.worker(workerID)
	}
}

func (o *backgroundObserver) worker(id int) {
	defer o.wg.Done()
	for {
		select {
		case <-o.done:
			return
		case <-o.msgQueue.newJobSignal():
			msg, ok := o.msgQueue.pop()
			if !ok {
				continue
			}

			observationTimeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // todo: cfg
			tokenData, err := o.observer.Observe(observationTimeoutCtx, map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
				msg.Header.SourceChainSelector: {msg.Header.SequenceNumber: msg},
			})
			cancel()

			if err != nil {
				o.lggr.Errorw("observing message", "msgID", msg.Header.MessageID.String(), "err", err)
				o.msgQueue.addJob(msg) // requeue
				continue
			}

			if !tokenData[msg.Header.SourceChainSelector][msg.Header.SequenceNumber].IsReady() {
				o.lggr.Debugw("token data not ready", "msgID", msg.Header.MessageID.String())
				o.msgQueue.addJob(msg) // requeue
				continue
			}

			o.cachedTokenData.Set(
				msg.Header.SourceChainSelector,
				msg.Header.SequenceNumber,
				tokenData[msg.Header.SourceChainSelector][msg.Header.SequenceNumber],
			)
		}
	}

}

type msgQueue struct {
	msgs             []cciptypes.Message
	mu               *sync.RWMutex
	newMsgSignalChan chan struct{}
}

func newMsgQueue() *msgQueue {
	return &msgQueue{
		msgs:             make([]cciptypes.Message, 0),
		mu:               &sync.RWMutex{},
		newMsgSignalChan: make(chan struct{}),
	}
}

func (q *msgQueue) addJob(msg cciptypes.Message) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.hasMsg(msg) {
		return false
	}

	q.msgs = append(q.msgs, msg)
	q.newMsgSignalChan <- struct{}{}
	return true
}

func (q *msgQueue) pop() (cciptypes.Message, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	// todo: optimize retrying delays

	if len(q.msgs) == 0 {
		return cciptypes.Message{}, false
	}

	job := q.msgs[0]
	q.msgs = q.msgs[1:]
	return job, true
}

func (q *msgQueue) hasMsg(msg cciptypes.Message) bool {
	for _, qMsg := range q.msgs {
		if qMsg.Header.MessageID == msg.Header.MessageID {
			return true
		}
	}

	return false
}

func (q *msgQueue) newJobSignal() <-chan struct{} {
	return q.newMsgSignalChan
}

type inMemTokenDataCache struct {
	expirationInterval time.Duration
	inMemTokenData     map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData
	expiresAt          map[cciptypes.ChainSelector]map[cciptypes.SeqNum]time.Time
	mu                 *sync.RWMutex
}

func newInMemObservationsCache(expirationInterval, cleanupInterval time.Duration) *inMemTokenDataCache {
	c := &inMemTokenDataCache{
		expirationInterval: expirationInterval,
		inMemTokenData:     make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData),
		expiresAt:          make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]time.Time),
		mu:                 &sync.RWMutex{},
	}
	c.runExpirationLoop(cleanupInterval)
	return c
}

func (c *inMemTokenDataCache) Get(
	chainSel cciptypes.ChainSelector,
	seqNum cciptypes.SeqNum,
) (exectypes.MessageTokenData, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	chainData, ok := c.inMemTokenData[chainSel]
	if !ok {
		return exectypes.MessageTokenData{}, false
	}

	msgData, ok := chainData[seqNum]
	if !ok {
		return exectypes.MessageTokenData{}, false
	}

	return msgData, true
}

func (c *inMemTokenDataCache) Set(
	chainSel cciptypes.ChainSelector,
	seqNum cciptypes.SeqNum,
	tokenData exectypes.MessageTokenData,
) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.inMemTokenData[chainSel]; !ok {
		c.inMemTokenData[chainSel] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
	}

	c.inMemTokenData[chainSel][seqNum] = tokenData
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

					for chainSel, seqNums := range c.expiresAt {
						for seqNum := range seqNums {
							if c.hasExpired(chainSel, seqNum) {
								delete(c.expiresAt[chainSel], seqNum)
								delete(c.inMemTokenData[chainSel], seqNum)
							}
						}
					}
				}()
			}
		}
	}()
}

func (c *inMemTokenDataCache) hasExpired(chainSel cciptypes.ChainSelector, seqNum cciptypes.SeqNum) bool {
	if chain, chainExists := c.expiresAt[chainSel]; chainExists {
		if expiresAt, seqExists := chain[seqNum]; seqExists {
			return time.Now().After(expiresAt)
		}
	}
	return false
}
