package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

const (
	// EvictionGracePeriod defines how long after the messageVisibilityInterval a root is still kept in the cache
	EvictionGracePeriod = 1 * time.Hour
	// CleanupInterval defines how often roots cache is scanned to evict stale roots
	CleanupInterval = 30 * time.Minute
)

// CommitsRootsCache keeps track of commit roots (and messages?) that are eligible for execution.
//
// This cache is used when:
//   - a commit root is fully executed and we want to skip it in the future. This would be called
//     in the GetCommitReports phase after checking if each SeqNr is executed (or if the Txm state is finalized).
//   - remember the oldest pending commit root to limit the database scan to only the unfinalized part of the chain.
type CommitsRootsCache interface {
	// TODO: Some function to check if a root is eligible for execution
	//RootsEligibleForExecution(ctx context.Context) ([]ccip.CommitStoreReport, error)
	MarkAsExecuted(merkleRoot ccipocr3.Bytes32)
	Snooze(merkleRoot ccipocr3.Bytes32)
}

func NewCommitRootsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
) CommitsRootsCache {
	return newCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
		CleanupInterval,
		EvictionGracePeriod,
	)
}

func newCommitRootsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
	cleanupInterval time.Duration,
	evictionGracePeriod time.Duration,
) *commitRootsCache {
	snoozedRoots := cache.New(rootSnoozeTime, cleanupInterval)
	executedRoots := cache.New(messageVisibilityInterval+evictionGracePeriod, cleanupInterval)

	return &commitRootsCache{
		lggr:                        lggr,
		rootSnoozeTime:              rootSnoozeTime,
		executedRoots:               executedRoots,
		snoozedRoots:                snoozedRoots,
		messageVisibilityInterval:   messageVisibilityInterval,
		latestFinalizedCommitRootTs: time.Now().Add(-messageVisibilityInterval),
		cacheMu:                     sync.RWMutex{},
	}
}

type commitRootsCache struct {
	lggr                      logger.Logger
	messageVisibilityInterval time.Duration
	rootSnoozeTime            time.Duration

	cacheMu sync.RWMutex

	// snoozedRoots used only for temporary snoozing roots. It's a cache with TTL (usually around 5 minutes, but this configuration is set up on chain using rootSnoozeTime)
	snoozedRoots *cache.Cache
	// executedRoots is a cache with TTL (usually around 8 hours, but this configuration is set up on chain using messageVisibilityInterval).
	// We keep executed roots there to make sure we don't accidentally try to reprocess already executed CommitReport
	executedRoots *cache.Cache
	// latestFinalizedCommitRootTs is the timestamp of the latest finalized commit root (youngest in terms of timestamp).
	// It's used get only the logs that were considered as unfinalized in a previous run.
	// This way we limit database scans to the minimum and keep polling "unfinalized" part of the ReportAccepted events queue.
	latestFinalizedCommitRootTs time.Time
}

// MarkAsExecuted marks the root as executed. It means that all the messages from the root were executed and the ExecutionStateChange event was finalized.
// Executed roots are removed from the cache.
func (r *commitRootsCache) MarkAsExecuted(merkleRoot ccipocr3.Bytes32) {
	prettyMerkleRoot := merkleRoot.String()
	r.lggr.Infow("Marking root as executed and removing entirely from cache", "merkleRoot", prettyMerkleRoot)

	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	r.executedRoots.SetDefault(prettyMerkleRoot, struct{}{})
}

// Snooze temporarily snoozes the root. It means that the root is not eligible for execution for a certain period of time.
// Snoozed roots are skipped when calling RootsEligibleForExecution
func (r *commitRootsCache) Snooze(merkleRoot ccipocr3.Bytes32) {
	prettyMerkleRoot := merkleRoot.String()
	r.lggr.Infow("Snoozing root temporarily", "merkleRoot", prettyMerkleRoot, "rootSnoozeTime", r.rootSnoozeTime)
	r.snoozedRoots.SetDefault(prettyMerkleRoot, struct{}{})
}

func (r *commitRootsCache) isSnoozed(merkleRoot ccipocr3.Bytes32) bool {
	_, snoozed := r.snoozedRoots.Get(merkleRoot.String())
	return snoozed
}

func (r *commitRootsCache) isExecuted(merkleRoot ccipocr3.Bytes32) bool {
	_, executed := r.executedRoots.Get(merkleRoot.String())
	return executed
}
