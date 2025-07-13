package tokenprice

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
)

func TestProcessor_Outcome_cacheInvalidation(t *testing.T) {
	lggr := logger.Test(t)
	ctx := t.Context()

	aobs := &asyncObserver{
		lggr:            lggr,
		cancelFunc:      nil,
		mu:              sync.RWMutex{},
		triggerSyncChan: make(chan time.Time),
	}

	p := &processor{
		lggr: lggr,
		obs:  aobs,
	}

	aobs.mu.Lock()
	aobs.tokenPriceMap = cciptypes.TokenPriceMap{"A": cciptypes.NewBigIntFromInt64(123)}
	aobs.mu.Unlock()

	// cache is not invalidated with a normal context
	_, _ = p.Outcome(ctx, Outcome{}, Query{}, nil)
	aobs.mu.Lock()
	updates := aobs.tokenPriceMap
	aobs.mu.Unlock()
	require.Len(t, updates, 1)

	// cache is not invalidated with invalidation context set to false
	ctx = context.WithValue(ctx, consts.InvalidateCacheKey, false)
	_, _ = p.Outcome(ctx, Outcome{}, Query{}, nil)
	aobs.mu.Lock()
	updates = aobs.tokenPriceMap
	aobs.mu.Unlock()
	require.Len(t, updates, 1)

	// cache is invalidated with invalidation context set to true and a sync op is triggered
	wg := sync.WaitGroup{}
	go func() {
		<-aobs.triggerSyncChan
	}()
	ctx = context.WithValue(ctx, consts.InvalidateCacheKey, true)
	_, _ = p.Outcome(ctx, Outcome{}, Query{}, nil)
	aobs.mu.Lock()
	updates = aobs.tokenPriceMap
	aobs.mu.Unlock()
	require.Len(t, updates, 0)
	wg.Wait() // wait until receiving the sync operation signal
}
