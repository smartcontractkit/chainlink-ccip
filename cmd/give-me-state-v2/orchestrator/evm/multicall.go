package evm

import (
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"give-me-state-v2/orchestrator"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Multicall3 configuration
const (
	// MaxBatchSize is the maximum number of calls to bundle into one multicall.
	MaxBatchSize = 16
	// BatchTimeout is how long the batcher waits for more calls before sending a partial batch.
	BatchTimeout = 10 * time.Millisecond
	// batchChanBuf is the buffer size for the pending calls channel.
	batchChanBuf = 256
)

// Multicall3 ABI (only tryAggregate and getChainId)
const multicall3ABIJson = `[{"inputs":[{"internalType":"bool","name":"requireSuccess","type":"bool"},{"components":[{"internalType":"address","name":"target","type":"address"},{"internalType":"bytes","name":"callData","type":"bytes"}],"internalType":"struct Multicall3.Call[]","name":"calls","type":"tuple[]"}],"name":"tryAggregate","outputs":[{"components":[{"internalType":"bool","name":"success","type":"bool"},{"internalType":"bytes","name":"returnData","type":"bytes"}],"internalType":"struct Multicall3.Result[]","name":"returnData","type":"tuple[]"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"getChainId","outputs":[{"internalType":"uint256","name":"chainid","type":"uint256"}],"stateMutability":"view","type":"function"}]`

var (
	multicall3ABI  abi.ABI
	multicall3Addr []byte // 20 bytes
)

func init() {
	var err error
	multicall3ABI, err = abi.JSON(strings.NewReader(multicall3ABIJson))
	if err != nil {
		panic("failed to parse Multicall3 ABI: " + err.Error())
	}
	multicall3Addr, _ = hex.DecodeString("cA11bde05977b3631167028862bE2a173976CA11")
}

// pendingCall represents a single call waiting to be batched.
type pendingCall struct {
	call     orchestrator.Call
	resultCh chan orchestrator.CallResult
}

// chainBatcher collects calls for a single chain and sends them via Multicall3.
type chainBatcher struct {
	pendingCh chan *pendingCall
	orcID     string
	chainID   uint64
	evm       *EVMOrchestrator
}

// run is the main batcher loop. It collects calls and dispatches batches
// when the batch is full (MaxBatchSize) or the timer expires (BatchTimeout).
func (b *chainBatcher) run() {
	var batch []*pendingCall
	timer := time.NewTimer(BatchTimeout)
	timer.Stop()

	for {
		select {
		case pc := <-b.pendingCh:
			batch = append(batch, pc)
			if len(batch) >= MaxBatchSize {
				b.dispatchBatch(batch)
				batch = nil
				timer.Stop()
			} else if len(batch) == 1 {
				// First call in a new batch -- start the countdown
				timer.Reset(BatchTimeout)
			}
		case <-timer.C:
			if len(batch) > 0 {
				b.dispatchBatch(batch)
				batch = nil
			}
		}
	}
}

// dispatchBatch sends a batch in a new goroutine so the batcher can continue collecting.
func (b *chainBatcher) dispatchBatch(batch []*pendingCall) {
	go b.sendBatch(batch)
}

// submit adds a call to the batcher and blocks until the result arrives.
func (b *chainBatcher) submit(call orchestrator.Call) orchestrator.CallResult {
	pc := &pendingCall{
		call:     call,
		resultCh: make(chan orchestrator.CallResult, 1),
	}
	b.pendingCh <- pc
	return <-pc.resultCh
}

// sendBatch encodes calls into a Multicall3 tryAggregate, executes it,
// and distributes the per-call results back to their callers.
func (b *chainBatcher) sendBatch(batch []*pendingCall) {
	// Single call -- skip multicall overhead entirely.
	if len(batch) == 1 {
		batch[0].resultCh <- b.evm.doRequest(b.orcID, batch[0].call)
		return
	}

	// Build the Call[] argument for tryAggregate.
	type MC3Call struct {
		Target   common.Address
		CallData []byte
	}
	calls := make([]MC3Call, len(batch))
	for i, pc := range batch {
		copy(calls[i].Target[:], pc.call.Target)
		calls[i].CallData = pc.call.Data
	}

	// Pack tryAggregate(false, calls)
	data, err := multicall3ABI.Pack("tryAggregate", false, calls)
	if err != nil {
		b.returnErrorToAll(batch, fmt.Errorf("multicall pack: %w", err))
		return
	}

	// Execute the multicall as a single eth_call to the Multicall3 contract.
	mcCall := orchestrator.Call{
		ChainID: b.chainID,
		Target:  multicall3Addr,
		Data:    data,
	}

	result := b.evm.doRequest(b.orcID, mcCall)
	if result.Error != nil {
		// Multicall itself failed -- fall back to individual calls so each
		// caller gets its own result rather than a blanket error.
		b.fallbackIndividual(batch)
		return
	}

	// Unpack the Result[] array.
	unpacked, err := multicall3ABI.Unpack("tryAggregate", result.Data)
	if err != nil {
		b.returnErrorToAll(batch, fmt.Errorf("multicall unpack: %w", err))
		return
	}

	if len(unpacked) == 0 {
		b.returnErrorToAll(batch, fmt.Errorf("multicall: empty response"))
		return
	}

	// Type-assert the returned Result[] = (bool success, bytes returnData)[]
	mcResults, ok := unpacked[0].([]struct {
		Success    bool   `json:"success"`
		ReturnData []byte `json:"returnData"`
	})
	if !ok {
		b.returnErrorToAll(batch, fmt.Errorf("multicall: unexpected result type %T", unpacked[0]))
		return
	}
	if len(mcResults) != len(batch) {
		b.returnErrorToAll(batch, fmt.Errorf("multicall: result count %d != call count %d", len(mcResults), len(batch)))
		return
	}

	// Distribute per-call results.
	for i, pc := range batch {
		if mcResults[i].Success {
			pc.resultCh <- orchestrator.CallResult{Data: mcResults[i].ReturnData}
		} else {
			pc.resultCh <- orchestrator.CallResult{Error: fmt.Errorf("execution reverted")}
		}
	}
}

// fallbackIndividual sends each call in the batch individually (used when the
// multicall request itself fails, e.g. response too large or RPC error).
func (b *chainBatcher) fallbackIndividual(batch []*pendingCall) {
	var wg sync.WaitGroup
	for _, pc := range batch {
		wg.Add(1)
		go func(p *pendingCall) {
			defer wg.Done()
			p.resultCh <- b.evm.doRequest(b.orcID, p.call)
		}(pc)
	}
	wg.Wait()
}

// returnErrorToAll sends the same error to every caller in the batch.
func (b *chainBatcher) returnErrorToAll(batch []*pendingCall, err error) {
	for _, pc := range batch {
		pc.resultCh <- orchestrator.CallResult{Error: err}
	}
}

// probeMulticall checks if Multicall3 is deployed on a chain by calling getChainId().
// Returns true if the call succeeds and returns a non-zero value.
func (e *EVMOrchestrator) probeMulticall(chainID uint64, orcID string) bool {
	data, err := multicall3ABI.Pack("getChainId")
	if err != nil {
		return false
	}

	call := orchestrator.Call{
		ChainID: chainID,
		Target:  multicall3Addr,
		Data:    data,
	}

	result := e.doRequest(orcID, call)
	if result.Error != nil {
		return false
	}

	// getChainId returns uint256 -- must be at least 32 bytes and non-zero.
	if len(result.Data) < 32 {
		return false
	}
	for _, b := range result.Data {
		if b != 0 {
			return true
		}
	}
	return false
}

// initMulticallBatchers probes Multicall3 on every registered chain concurrently
// and starts batcher goroutines for chains that support it.
func (e *EVMOrchestrator) initMulticallBatchers() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for chainID, orcID := range e.orcIDs {
		wg.Add(1)
		go func(cid uint64, oid string) {
			defer wg.Done()
			if e.probeMulticall(cid, oid) {
				batcher := &chainBatcher{
					pendingCh: make(chan *pendingCall, batchChanBuf),
					orcID:     oid,
					chainID:   cid,
					evm:       e,
				}
				mu.Lock()
				e.batchers[cid] = batcher
				mu.Unlock()
				go batcher.run()
				fmt.Fprintf(os.Stderr, "[multicall] chain %d: enabled (batching up to %d calls)\n", cid, MaxBatchSize)
			} else {
				fmt.Fprintf(os.Stderr, "[multicall] chain %d: not available, using individual calls\n", cid)
			}
		}(chainID, orcID)
	}

	wg.Wait()
}
