package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// TimeProvider interface to make time testable
type TimeProvider interface {
	Now() time.Time
}

// RealTimeProvider provides actual current time
type RealTimeProvider struct{}

func (r *RealTimeProvider) Now() time.Time {
	return time.Now().UTC()
}

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
	// It will return the message visibility window or the earliest unexecuted root timestamp,
	// whichever is later, to optimize the query window.
	GetTimestampToQueryFrom() time.Time

	// UpdateEarliestUnexecutedRoot updates the earliest unexecuted root timestamp
	// based on the remaining unexecuted and executed but not finalized reports.
	UpdateEarliestUnexecutedRoot(remainingReports map[ccipocr3.ChainSelector][]exectypes.CommitData)

	UpdateLatestEmptyRootTimestamp(timestamp time.Time)
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
		&RealTimeProvider{},
	)
}

// For testing - create with custom time provider
func NewCommitRootsCacheWithTimeProvider(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
	timeProvider TimeProvider,
) CommitsRootsCache {
	return internalNewCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
		timeProvider,
	)
}

func internalNewCommitRootsCache(
	lggr logger.Logger,
	messageVisibilityInterval time.Duration,
	rootSnoozeTime time.Duration,
	timeProvider TimeProvider,
) *commitRootsCache {
	lggr.Debugw(
		"Creating CommitRootsCache",
		"messageVisibilityInterval", messageVisibilityInterval,
		"rootSnoozeTime", rootSnoozeTime,
		"cleanupInterval", CleanupInterval,
		"evictionGracePeriod", EvictionGracePeriod,
	)

	snoozedRoots := cache.New(rootSnoozeTime, CleanupInterval)
	executedRoots := cache.New(messageVisibilityInterval+EvictionGracePeriod, CleanupInterval)

	return &commitRootsCache{
		lggr:                      lggr,
		rootSnoozeTime:            rootSnoozeTime,
		executedRoots:             executedRoots,
		snoozedRoots:              snoozedRoots,
		messageVisibilityInterval: messageVisibilityInterval,
		cacheMu:                   sync.RWMutex{},
		timeProvider:              timeProvider,
	}
}

type commitRootsCache struct {
	lggr                      logger.Logger
	messageVisibilityInterval time.Duration
	rootSnoozeTime            time.Duration
	timeProvider              TimeProvider

	cacheMu sync.RWMutex

	// snoozedRoots used only for temporary snoozing roots. It's a cache with TTL (usually around 5 minutes,
	// but this configuration is set up on chain using rootSnoozeTime)
	snoozedRoots *cache.Cache
	// executedRoots is a cache with TTL (usually around 8 hours, but this configuration is set up on chain using
	// messageVisibilityInterval). We keep executed roots there to make sure we don't accidentally try to reprocess
	// already executed CommitReport
	executedRoots *cache.Cache
	// earliestUnexecutedRoot tracks the timestamp of the earliest unexecuted root
	// to optimize database queries by potentially starting from a later timestamp
	earliestUnexecutedRoot time.Time

	// latestEmptyRoot tracks the timestamp of the latest Finalized empty root,
	// This is to ensure that the plugin is moving forward if all reports within round limit are empty

	latestEmptyRoot time.Time
}

func (r *commitRootsCache) UpdateLatestEmptyRootTimestamp(timestamp time.Time) {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()
	if timestamp.After(r.latestEmptyRoot) {
		r.latestEmptyRoot = timestamp
	}
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

// GetTimestampToQueryFrom returns the timestamp to use when querying for commit reports.
// It optimizes the query window by using the later of:
// 1. The message visibility window
// 2. Min(earliestUnexecutedRoot, latestEmptyRoot) if it's after the message visibility window
// TODO: add diagram in github instead of using external link
// For illustration check https://app.excalidraw.com/l/AdjkJ3DaenS/84EpHxkgbND
func (r *commitRootsCache) GetTimestampToQueryFrom() time.Time {
	r.cacheMu.RLock()
	defer r.cacheMu.RUnlock()

	// Calculate current visibility window based on stored interval
	messageVisibilityWindow := r.timeProvider.Now().Add(-r.messageVisibilityInterval).UTC()

	// Start with message visibility window as the default (lower bound)
	commitRootsFilterTimestamp := messageVisibilityWindow

	// Determine the earliest timestamp between unexecuted root and latest empty root
	var minTimestamp time.Time

	// Check if earliestUnexecutedRoot is set
	if !r.earliestUnexecutedRoot.IsZero() {
		minTimestamp = r.earliestUnexecutedRoot
		r.lggr.Debugw("MinTimeStamp set to earliest unexecuted root")
	}

	// Check if latestEmptyRoot is set
	if !r.latestEmptyRoot.IsZero() {
		// If minTimestamp is not set or latestEmptyRoot is earlier, use latestEmptyRoot
		if minTimestamp.IsZero() || r.latestEmptyRoot.Before(minTimestamp) {
			minTimestamp = r.latestEmptyRoot
		}
		r.lggr.Debugw("MinTimeStamp set to latest empty finalized root")
	}

	// If we know the earliest unexecuted root or latestEmptymTimestamp and it's AFTER the visibility window,
	// we can optimize by starting our query from that timestamp instead
	if minTimestamp.After(messageVisibilityWindow) {
		commitRootsFilterTimestamp = minTimestamp
		r.lggr.Debugw("Using minTimestamp to optimize query",
			"minTimestamp", minTimestamp,
			"messageVisibilityWindow", messageVisibilityWindow)
	}

	r.lggr.Debugw("Getting timestamp to query from",
		"earliestUnexecutedRoot", r.earliestUnexecutedRoot,
		"latestEmptyRoot", r.latestEmptyRoot,
		"messageVisibilityWindow", messageVisibilityWindow,
		"commitRootsFilterTimestamp", commitRootsFilterTimestamp,
		"optimized", commitRootsFilterTimestamp.After(messageVisibilityWindow))

	return commitRootsFilterTimestamp
}

// UpdateEarliestUnexecutedRoot updates the earliest unexecuted root timestamp
// based on the remaining unexecuted reports.
func (r *commitRootsCache) UpdateEarliestUnexecutedRoot(
	remainingReports map[ccipocr3.ChainSelector][]exectypes.CommitData) {
	r.cacheMu.Lock()
	defer r.cacheMu.Unlock()

	// If no unexecuted reports remain, we keep the current value
	// as it represents the last known earliest unexecuted root
	if len(remainingReports) == 0 {
		return
	}

	// Find the earliest timestamp among all remaining reports
	var earliestTimestamp time.Time
	for _, reports := range remainingReports {
		for _, report := range reports {
			if earliestTimestamp.IsZero() || report.Timestamp.Before(earliestTimestamp) {
				earliestTimestamp = report.Timestamp
			}
		}
	}

	// Check if we found a valid timestamp
	if earliestTimestamp.IsZero() {
		return
	}

	// Only update if the new timestamp is later than the current one
	if earliestTimestamp.After(r.earliestUnexecutedRoot) {
		r.lggr.Debugw("Updating earliest unexecuted root",
			"oldTimestamp", r.earliestUnexecutedRoot,
			"newTimestamp", earliestTimestamp)
		r.earliestUnexecutedRoot = earliestTimestamp
	}
}
