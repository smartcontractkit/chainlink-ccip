package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
	CanExecute(source ccipocr3.ChainSelector, merkleRoot ccipocr3.Bytes32) bool
	MarkAsExecuted(source ccipocr3.ChainSelector, merkleRoot ccipocr3.Bytes32)
	Snooze(source ccipocr3.ChainSelector, merkleRoot ccipocr3.Bytes32)

	// GetTimestampToQueryFrom returns the timestamp that should be used when querying for commit reports.
	// If the latest finalized fully executed commit root's timestamp is before the message visibility window,
	// the message visibility window is used as the timestamp.
	// Otherwise, the latest finalized fully executed commit root's timestamp is used.
	GetTimestampToQueryFrom() time.Time

	// UpdateLatestFinalizedTimestamp updates the cached latest finalized timestamp.
	// This is used to optimize commit report queries by tracking the newest finalized report.
	UpdateLatestFinalizedTimestamp(timestamp time.Time)
}

func NewCommitRootsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
) CommitsRootsCache {
	return internalNewCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
		CleanupInterval,
		EvictionGracePeriod,
	)
}

func internalNewCommitRootsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
	cleanupInterval time.Duration,
	evictionGracePeriod time.Duration,
) *commitRootsCache {
	snoozedRoots := cache.New(rootSnoozeTime, cleanupInterval)
	executedRoots := cache.New(messageVisibilityInterval+evictionGracePeriod, cleanupInterval)

	return &commitRootsCache{
		lggr:                             lggr,
		rootSnoozeTime:                   rootSnoozeTime,
		executedRoots:                    executedRoots,
		snoozedRoots:                     snoozedRoots,
		messageVisibilityInterval:        messageVisibilityInterval,
		latestFinalizedFullyExecutedRoot: time.Time{}, // Zero value (not set)
		cacheMu:                          sync.Mutex{},
	}
}

type commitRootsCache struct {
	lggr                      logger.Logger
	messageVisibilityInterval time.Duration
	rootSnoozeTime            time.Duration

	cacheMu sync.Mutex

	// snoozedRoots used only for temporary snoozing roots. It's a cache with TTL (usually around 5 minutes,
	// but this configuration is set up on chain using rootSnoozeTime)
	snoozedRoots *cache.Cache
	// executedRoots is a cache with TTL (usually around 8 hours, but this configuration is set up on chain using
	// messageVisibilityInterval). We keep executed roots there to make sure we don't accidentally try to reprocess
	// already executed CommitReport
	executedRoots *cache.Cache
	// latestFinalizedFullyExecutedRoot tracks the timestamp of the latest finalized commit root
	// to optimize database queries by starting from this timestamp
	latestFinalizedFullyExecutedRoot time.Time
}

func getKey(source ccipocr3.ChainSelector, merkleRoot ccipocr3.Bytes32) string {
	return source.String() + "_" + merkleRoot.String()
}

// MarkAsExecuted marks the root as executed. It means that all the messages from the root were executed and the
// ExecutionStateChange event was finalized.
func (r *commitRootsCache) MarkAsExecuted(sel ccipocr3.ChainSelector, merkleRoot ccipocr3.Bytes32) {
	prettyMerkleRoot := getKey(sel, merkleRoot)
	r.lggr.Infow("Marking root as executed and removing entirely from cache", "merkleRoot", prettyMerkleRoot)

	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	r.executedRoots.SetDefault(prettyMerkleRoot, struct{}{})
}

// Snooze temporarily snoozes the root. It means that the root is not eligible for execution for a certain period of
// time.
func (r *commitRootsCache) Snooze(sel ccipocr3.ChainSelector, merkleRoot ccipocr3.Bytes32) {
	prettyMerkleRoot := getKey(sel, merkleRoot)
	r.lggr.Infow("Snoozing root temporarily", "merkleRoot", prettyMerkleRoot, "rootSnoozeTime", r.rootSnoozeTime)
	r.snoozedRoots.SetDefault(prettyMerkleRoot, struct{}{})
}

// CanExecute returns true if the root is not snoozed and not executed.
func (r *commitRootsCache) CanExecute(sel ccipocr3.ChainSelector, merkleRoot ccipocr3.Bytes32) bool {
	prettyMerkleRoot := getKey(sel, merkleRoot)
	return !r.isSnoozed(prettyMerkleRoot) && !r.isExecuted(prettyMerkleRoot)
}

// IsSnoozed returns true if the root is snoozed.
func (r *commitRootsCache) isSnoozed(key string) bool {
	_, snoozed := r.snoozedRoots.Get(key)
	return snoozed
}

// isExecuted returns true if the root is executed.
func (r *commitRootsCache) isExecuted(key string) bool {
	_, executed := r.executedRoots.Get(key)
	return executed
}

func (r *commitRootsCache) GetTimestampToQueryFrom() time.Time {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()

	messageVisibilityWindow := time.Now().Add(-r.messageVisibilityInterval)
	if r.latestFinalizedFullyExecutedRoot.Before(messageVisibilityWindow) {
		r.latestFinalizedFullyExecutedRoot = messageVisibilityWindow
	}
	commitRootsFilterTimestamp := r.latestFinalizedFullyExecutedRoot

	r.lggr.Debugw("Getting timestamp to query from",
		"latestFinalizedTimestamp", r.latestFinalizedFullyExecutedRoot,
		"messageVisibilityWindow", messageVisibilityWindow,
		"commitRootsFilterTimestamp", commitRootsFilterTimestamp)

	// Otherwise use the visibility window based on latest finalized timestamp
	return commitRootsFilterTimestamp
}

func (r *commitRootsCache) UpdateLatestFinalizedTimestamp(timestamp time.Time) {
	if timestamp.IsZero() {
		return
	}

	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()

	// Update if not set or if the new timestamp is newer
	if r.latestFinalizedFullyExecutedRoot.IsZero() || timestamp.After(r.latestFinalizedFullyExecutedRoot) {
		r.lggr.Debugw("Updating latest finalized timestamp",
			"oldTimestamp", r.latestFinalizedFullyExecutedRoot,
			"newTimestamp", timestamp)
		r.latestFinalizedFullyExecutedRoot = timestamp
	}
}
