package executor

import (
	"container/heap"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

// timedMessageItem represents a message with its delivery time
type timedMessageItem struct {
	message   modsectypes.Message
	deliverAt time.Time
	index     int // for heap.Interface
}

// timedMessageHeap implements heap.Interface for managing timed messages
type timedMessageHeap []*timedMessageItem

func (h timedMessageHeap) Len() int { return len(h) }

func (h timedMessageHeap) Less(i, j int) bool {
	return h[i].deliverAt.Before(h[j].deliverAt)
}

func (h timedMessageHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].index = i
	h[j].index = j
}

func (h *timedMessageHeap) Push(x interface{}) {
	n := len(*h)
	item := x.(*timedMessageItem)
	item.index = n
	*h = append(*h, item)
}

func (h *timedMessageHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*h = old[0 : n-1]
	return item
}

// timedMessageChannel implements TimedMessageChannel
type timedMessageChannel struct {
	mu         sync.Mutex
	heap       *timedMessageHeap
	outputChan chan modsectypes.Message
	stopChan   chan struct{}
	workerDone chan struct{}
	closed     bool
}

// NewTimedMessageChannel creates a new TimedMessageChannel with the specified buffer size
func NewTimedMessageChannel(bufferSize int) modsectypes.TimedMessageChannel {
	tmc := &timedMessageChannel{
		heap:       &timedMessageHeap{},
		outputChan: make(chan modsectypes.Message, bufferSize),
		stopChan:   make(chan struct{}),
		workerDone: make(chan struct{}),
	}

	heap.Init(tmc.heap)
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

	heap.Push(tmc.heap, item)
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

// worker processes the timed messages and sends them to the output channel when ready
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

	now := time.Now()

	// Process all messages that are ready to be delivered
	for tmc.heap.Len() > 0 {
		item := (*tmc.heap)[0]
		if item.deliverAt.After(now) {
			break // No more messages ready
		}

		// Remove from heap
		heap.Pop(tmc.heap)

		// Send to output channel (non-blocking)
		select {
		case tmc.outputChan <- item.message:
		default:
			// Channel is full, could log this or handle overflow
		}
	}
}
