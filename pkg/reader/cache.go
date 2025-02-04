package reader

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type cache struct {
	sync.RWMutex
	data          NogoResponse
	lastRefresh   time.Time
	refreshPeriod time.Duration
}

type CachedChainReader struct {
	*ccipChainReader
	cache *cache
}

func NewCachedChainReader(
	reader *ccipChainReader,
	refreshPeriod time.Duration,
) *CachedChainReader {
	reader.lggr.Infow("Creating new cached chain reader",
		"refreshPeriod", refreshPeriod)
	return &CachedChainReader{
		ccipChainReader: reader,
		cache: &cache{
			refreshPeriod: refreshPeriod,
		},
	}
}

func (r *CachedChainReader) getCachedResponse(ctx context.Context) (NogoResponse, error) {
	r.cache.RLock()
	timeSinceLastRefresh := time.Since(r.cache.lastRefresh)
	if timeSinceLastRefresh < r.cache.refreshPeriod {
		defer r.cache.RUnlock()
		r.lggr.Infow("Cache hit",
			"timeSinceLastRefresh", timeSinceLastRefresh,
			"refreshPeriod", r.cache.refreshPeriod)
		return r.cache.data, nil
	}
	r.cache.RUnlock()

	r.lggr.Infow("Cache miss, refreshing",
		"timeSinceLastRefresh", timeSinceLastRefresh,
		"refreshPeriod", r.cache.refreshPeriod)
	return r.refreshCache(ctx)
}

func (r *CachedChainReader) refreshCache(ctx context.Context) (NogoResponse, error) {
	r.cache.Lock()
	defer r.cache.Unlock()

	timeSinceLastRefresh := time.Since(r.cache.lastRefresh)
	if timeSinceLastRefresh < r.cache.refreshPeriod {
		r.lggr.Infow("Cache was refreshed by another goroutine",
			"timeSinceLastRefresh", timeSinceLastRefresh)
		return r.cache.data, nil
	}

	r.lggr.Infow("Starting cache refresh",
		"lastRefresh", r.cache.lastRefresh)

	startTime := time.Now()
	newData, err := r.ccipChainReader.refresh(ctx)
	refreshDuration := time.Since(startTime)

	if err != nil {
		if !r.cache.lastRefresh.IsZero() {
			r.lggr.Warnw("Failed to refresh cache, using old data",
				"error", err,
				"lastRefresh", r.cache.lastRefresh,
				"refreshDuration", refreshDuration)
			return r.cache.data, nil
		}
		r.lggr.Errorw("Failed to refresh cache, no old data available",
			"error", err,
			"refreshDuration", refreshDuration)
		return NogoResponse{}, fmt.Errorf("failed to refresh cache: %w", err)
	}

	r.cache.data = newData
	r.cache.lastRefresh = time.Now()

	r.lggr.Infow("Cache refresh completed",
		"refreshDuration", refreshDuration,
		"newLastRefresh", r.cache.lastRefresh)

	return newData, nil
}

func (r *CachedChainReader) refresh(ctx context.Context) (NogoResponse, error) {
	return r.getCachedResponse(ctx)
}

func (r *CachedChainReader) ForceRefresh(ctx context.Context) error {
	r.lggr.Infow("Force refreshing cache")
	_, err := r.refreshCache(ctx)
	if err != nil {
		r.lggr.Errorw("Force refresh failed",
			"error", err)
	} else {
		r.lggr.Infow("Force refresh completed successfully")
	}
	return err
}

func (r *CachedChainReader) GetCacheStats() (time.Time, time.Duration) {
	r.cache.RLock()
	defer r.cache.RUnlock()
	r.lggr.Infow("Getting cache stats",
		"lastRefresh", r.cache.lastRefresh,
		"refreshPeriod", r.cache.refreshPeriod,
		"timeSinceLastRefresh", time.Since(r.cache.lastRefresh))
	return r.cache.lastRefresh, r.cache.refreshPeriod
}
