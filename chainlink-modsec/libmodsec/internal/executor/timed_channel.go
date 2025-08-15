package executor

import (
	"sync"
	"time"

	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

// timedMessageItem represents a message with its delivery time
type timedMessageItem struct {
	message   modsectypes.Message
	deliverAt time.Time
}

// timedMessageChannel implements TimedMessageChannel
type timedMessageChannel struct {
	mu         sync.Mutex
	queue      *priorityqueue.Queue
	outputChan chan modsectypes.Message
	stopChan   chan struct{}
	workerDone chan struct{}
	closed     bool
}

// NewTimedMessageChannel creates a new TimedMessageChannel with the specified buffer size
func NewTimedMessageChannel(bufferSize int) modsectypes.TimedMessageChannel {
	tmc := &timedMessageChannel{
		queue: priorityqueue.NewWith(func(a, b interface{}) int {
			itemA := a.(*timedMessageItem)
			itemB := b.(*timedMessageItem)
			return itemA.deliverAt.Compare(itemB.deliverAt)
		}),
		outputChan: make(chan modsectypes.Message, bufferSize),
		stopChan:   make(chan struct{}),
		workerDone: make(chan struct{}),
	}

	go tmc.worker()

	return tmc
}

// SendMessage adds a message to be delivered after the specified tick duration
func (tmc *timedMessageChannel) SendMessage(msg modsectypes.Message, tick time.Duration) {
	tmc.mu.Lock()
	defer tmc.mu.Unlock()

	if tmc.closed {
		return
	}

	deliverAt := time.Now().Add(tick)
	item := &timedMessageItem{
		message:   msg,
		deliverAt: deliverAt,
	}

	tmc.queue.Enqueue(item)
}

// Messages returns the channel that receives messages after their tick duration has elapsed
func (tmc *timedMessageChannel) Messages() <-chan modsectypes.Message {
	return tmc.outputChan
}

// Close closes the channel and stops processing
func (tmc *timedMessageChannel) Close() {
	tmc.mu.Lock()
	defer tmc.mu.Unlock()

	if tmc.closed {
		return
	}

	tmc.closed = true
	close(tmc.stopChan)
	<-tmc.workerDone
	close(tmc.outputChan)
}

// worker processes the timed messages using event-driven timers
func (tmc *timedMessageChannel) worker() {
	defer close(tmc.workerDone)

	ticker := time.NewTicker(100 * time.Millisecond) // Check every 100ms
	defer ticker.Stop()

	for {
		select {
		case <-tmc.stopChan:
			return
		case <-ticker.C:
			tmc.processReadyMessages()
		}
	}
}

// processReadyMessages checks for messages that are ready to be delivered
func (tmc *timedMessageChannel) processReadyMessages() {
	tmc.mu.Lock()
	defer tmc.mu.Unlock()

	if tmc.closed {
		return
	}

	// Get the next message
	next, found := tmc.queue.Dequeue()
	if !found {
		return
	}

	item := next.(*timedMessageItem)

	// Check if it's actually ready (in case of clock drift or delays)
	if item.deliverAt.After(time.Now()) {
		// Put it back and reschedule
		tmc.queue.Enqueue(item)
		return
	}

	// Send to output channel (non-blocking)
	select {
	case tmc.outputChan <- item.message:
		// Message sent successfully
	default:
		// Channel is full, could log this or handle overflow
	}
}
