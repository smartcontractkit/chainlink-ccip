package reader

import (
	"context"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_CachingInstances(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)

	chain1 := readermock.NewMockContractReaderFacade(t)
	chain2 := readermock.NewMockContractReaderFacade(t)

	for _, chain := range []*readermock.MockContractReaderFacade{chain1, chain2} {
		chain.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil).Maybe()
		chain.EXPECT().GetLatestValue(
			mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything,
		).Return(nil).Maybe()
	}

	t.Run("reusing instance for the same chain and address", func(t *testing.T) {
		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()
		poller1 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)
		poller2 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)
		poller3 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)

		require.True(t, poller1.bgPoller == poller2.bgPoller)
		require.True(t, poller2.bgPoller == poller3.bgPoller)

		require.NoError(t, poller1.Close())
		require.NoError(t, poller2.Close())
		require.NoError(t, poller3.Close())

		requirePollerStopped(t, poller1)
	})

	t.Run("creating new instance for different addresses on a single chain", func(t *testing.T) {
		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address1 := rand.RandomAddressBytes()
		address2 := rand.RandomAddressBytes()

		poller1 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address1, chain1)
		poller2 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address2, chain1)

		require.False(t, poller1.bgPoller == poller2.bgPoller)
		require.NoError(t, poller1.Close())
		require.NoError(t, poller2.Close())
		requirePollerStopped(t, poller1)
	})

	t.Run("creating new instance for different chains but same addresses", func(t *testing.T) {
		chainSelector1 := cciptypes.ChainSelector(rand.RandomInt64())
		chainSelector2 := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()

		poller1 := newRMNHomeCasted(ctx, t, lggr, chainSelector1, address, chain1)
		poller2 := newRMNHomeCasted(ctx, t, lggr, chainSelector2, address, chain2)

		require.False(t, poller1.bgPoller == poller2.bgPoller)
		require.NoError(t, poller1.Close())
		require.NoError(t, poller2.Close())
		requirePollerStopped(t, poller1)
		requirePollerStopped(t, poller2)
	})

	t.Run("poller is not closed when there is remaining parent reference", func(t *testing.T) {
		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()
		poller1 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)
		poller2 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)
		poller3 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)

		require.NoError(t, poller1.Close())
		requirePollerStarted(t, poller1)

		require.NoError(t, poller2.Close())
		require.NoError(t, poller2.Close())
		requirePollerStarted(t, poller1)

		require.NoError(t, poller3.Close())
		requirePollerStopped(t, poller1)

		require.NoError(t, poller3.Close())
		requirePollerStopped(t, poller1)
	})

	t.Run("parallel creation of instances doesn't cause any failures", func(t *testing.T) {
		instancesMu.Lock()
		instances = make(map[string]*rmnHomePoller)
		instancesMu.Unlock()

		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()

		pollers := make([]*rmnHome, 1000)

		eg := new(errgroup.Group)
		for i := 0; i < 1000; i++ {
			j := i
			eg.Go(func() error {
				pollers[j] = newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)
				return nil
			})
		}
		require.NoError(t, eg.Wait())
		require.Len(t, instances, 1)

		// 999 closed, but still one reference remains therefore bgPoller is running
		eg = new(errgroup.Group)
		for i := 0; i < 999; i++ {
			j := i
			eg.Go(func() error {
				requirePollerStarted(t, pollers[j])
				require.NoError(t, pollers[j].Close())
				requirePollerStarted(t, pollers[j])
				return nil
			})
		}
		require.NoError(t, eg.Wait())

		// All closed, bgPoller should be stopped
		require.NoError(t, pollers[999].Close())
		requirePollerStopped(t, pollers[0])
	})

	t.Run("create new instance when old one is already stopped", func(t *testing.T) {
		chainSelector := cciptypes.ChainSelector(rand.RandomInt64())
		address := rand.RandomAddressBytes()

		poller1 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)
		require.NoError(t, poller1.Close())
		require.Error(t, poller1.Start(ctx))
		requirePollerStopped(t, poller1)

		poller2 := newRMNHomeCasted(ctx, t, lggr, chainSelector, address, chain1)
		require.NoError(t, poller2.Ready())
		requirePollerStarted(t, poller2)

		require.NoError(t, poller2.Close())
		requirePollerStopped(t, poller2)
	})
}

func requirePollerStarted(t *testing.T, home *rmnHome) {
	require.NoError(t, home.bgPoller.Ready())
	require.Equal(t, "Started", home.bgPoller.sync.State())
}

func requirePollerStopped(t *testing.T, home *rmnHome) {
	require.Error(t, home.bgPoller.Ready())
	require.Equal(t, "Stopped", home.bgPoller.sync.State())
}

func newRMNHomeCasted(
	ctx context.Context,
	t *testing.T,
	lggr logger.Logger,
	selector cciptypes.ChainSelector,
	address cciptypes.Bytes,
	reader *readermock.MockContractReaderFacade,
) *rmnHome {
	rmn, err := NewRMNHomeChainReader(ctx, lggr, 1*time.Millisecond, selector, address, reader)
	require.NoError(t, err)

	err = rmn.Start(ctx)
	require.NoError(t, err)

	casted, ok := rmn.(*rmnHome)
	require.True(t, ok)
	return casted
}
