package plugincommon

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

type BackgroundReaderSyncer struct {
	lggr          logger.Logger
	reader        reader.CCIPReader
	syncTimeout   time.Duration
	syncFrequency time.Duration

	bgSyncCtx    context.Context
	bgSyncCf     context.CancelFunc
	bgSyncWG     *sync.WaitGroup
	bgSyncTicker *time.Ticker
}

var syncTimeoutRecommendedRange = [2]time.Duration{500 * time.Millisecond, 15 * time.Second}
var syncFrequencyRecommendedRange = [2]time.Duration{time.Second, time.Hour}

func NewBackgroundReaderSyncer(
	lggr logger.Logger,
	reader reader.CCIPReader,
	syncTimeout time.Duration,
	syncFrequency time.Duration,
) *BackgroundReaderSyncer {

	if syncTimeout < syncTimeoutRecommendedRange[0] || syncTimeout > syncTimeoutRecommendedRange[1] {
		lggr.Warnw("syncTimeout outside recommended range", "syncTimeout", syncTimeout)
	}

	if syncFrequency < syncFrequencyRecommendedRange[0] || syncFrequency > syncFrequencyRecommendedRange[1] {
		lggr.Warnw("syncFrequency outside recommended range", "syncFrequency", syncFrequency)
	}

	return &BackgroundReaderSyncer{
		lggr:          lggr,
		reader:        reader,
		syncTimeout:   syncTimeout,
		syncFrequency: syncFrequency,
	}
}

func (b *BackgroundReaderSyncer) Start(ctx context.Context) error {
	if b.bgSyncCtx != nil {
		return fmt.Errorf("background syncer already started")
	}

	b.bgSyncCtx, b.bgSyncCf = context.WithCancel(ctx)
	b.bgSyncWG = &sync.WaitGroup{}
	b.bgSyncWG.Add(1)
	b.bgSyncTicker = time.NewTicker(b.syncFrequency)

	backgroundReaderSync(
		b.bgSyncCtx,
		b.bgSyncWG,
		b.lggr,
		b.reader,
		b.syncTimeout,
		b.bgSyncTicker.C,
	)

	return nil
}

func (b *BackgroundReaderSyncer) Close() error {
	if b.bgSyncCtx == nil {
		return fmt.Errorf("background syncer not started")
	}

	if b.bgSyncCf != nil {
		b.bgSyncCf()
		b.bgSyncWG.Wait()
	}

	b.bgSyncTicker.Stop()

	return nil
}

// backgroundReaderSync runs a background process that periodically syncs the provided CCIPReader.
func backgroundReaderSync(
	ctx context.Context,
	wg *sync.WaitGroup,
	lggr logger.Logger,
	reader reader.CCIPReader,
	syncTimeout time.Duration,
	ticker <-chan time.Time,
) {
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				lggr.Debug("backgroundReaderSync context done")
				return
			case <-ticker:
				if err := syncReader(ctx, reader, syncTimeout); err != nil {
					lggr.Errorw("runBackgroundReaderSync failed", "err", err)
				}
				lggr.Infow("runBackgroundReaderSync success")
			}
		}
	}()
}

func syncReader(
	ctx context.Context,
	reader reader.CCIPReader,
	syncTimeout time.Duration,
) error {
	timeoutCtx, cf := context.WithTimeout(ctx, syncTimeout)
	defer cf()

	err := reader.Sync(timeoutCtx, nil)
	if err != nil {
		return fmt.Errorf("syncReader: %w", err)
	}

	return nil
}
