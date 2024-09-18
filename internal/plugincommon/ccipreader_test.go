package plugincommon

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
)

func TestBackgroundReaderSyncer(t *testing.T) {
	lggr := logger.Test(t)
	mockReader := reader_mock.NewMockCCIPReader(t)

	t.Run("start/stop checks", func(t *testing.T) {
		readerSyncer := NewBackgroundReaderSyncer(lggr, mockReader, time.Hour, time.Hour)

		ctx, cf := context.WithCancel(context.Background())

		err := readerSyncer.Close()
		assert.Error(t, err, "closing a non-started syncer")

		err = readerSyncer.Start(ctx)
		assert.NoError(t, err, "start success")

		err = readerSyncer.Start(ctx)
		assert.Error(t, err, "cannot be started twice")

		err = readerSyncer.Close()
		assert.NoError(t, err, "closing a started syncer")

		err = readerSyncer.Start(ctx)
		assert.Error(t, err, "restarting")

		cf()
		err = readerSyncer.Close()
		assert.NoError(t, err, "closing a syncer with expired context")
	})

	t.Run("syncing", func(t *testing.T) {
		ctx := context.Background()
		mockReader.On("Sync", mock.Anything, mock.Anything).Return(nil)
		readerSyncer := NewBackgroundReaderSyncer(lggr, mockReader, time.Second, time.Millisecond)
		err := readerSyncer.Start(ctx)
		assert.NoError(t, err, "start success")
		assert.Eventually(t, func() bool {
			return mockReader.AssertExpectations(t)
		}, 3*time.Second, 10*time.Millisecond)
		err = readerSyncer.Close()
		assert.NoError(t, err, "closing a started syncer")
	})
}

func Test_backgroundReaderSync(t *testing.T) {
	ctx, cf := context.WithCancel(context.Background())
	lggr := logger.Test(t)
	reader := reader_mock.NewMockCCIPReader(t)
	syncTimeout := 50 * time.Millisecond
	ticker := make(chan time.Time)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// start background syncing
	backgroundReaderSync(ctx, wg, lggr, reader, syncTimeout, ticker)

	// send a tick to trigger the first sync that errors
	reader.On("Sync", mock.Anything, mock.Anything).Return(fmt.Errorf("some err")).Once()
	ticker <- time.Now()

	// send a tick to trigger the second sync that succeeds without changes
	reader.On("Sync", mock.Anything, mock.Anything).Return(nil).Once()
	ticker <- time.Now()

	// make sync hang to see the context timeout
	reader.On("Sync", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ctx := args.Get(0).(context.Context)
		for { // simulate endless work until context times out
			select {
			case <-ctx.Done():
				t.Log("context cancelled as expected")
				return
			default:
				time.Sleep(time.Millisecond) // sleep to not block the CPU
			}
		}
	}).Return(nil).Once()
	ticker <- time.Now()

	// send a tick to trigger the fourth sync that succeeds with changes
	reader.On("Sync", mock.Anything, mock.Anything).Return(nil).Once()
	ticker <- time.Now()

	cf()      // trigger bg sync to stop
	wg.Wait() // wait for it to stop
	reader.AssertExpectations(t)
}
