package plugincommon

import (
	"context"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// BackgroundReaderSync runs a background process that periodically syncs the provider CCIP reader.
func BackgroundReaderSync(
	ctx context.Context,
	wg *sync.WaitGroup,
	lggr logger.Logger,
	reader reader.CCIP,
	syncTimeout time.Duration,
	ticker <-chan time.Time,
) {
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker:
				if err := syncReader(ctx, lggr, reader, syncTimeout); err != nil {
					lggr.Errorw("runBackgroundReaderSync failed", "err", err)
				}
			}
		}
	}()
}

func syncReader(
	ctx context.Context,
	lggr logger.Logger,
	reader reader.CCIP,
	syncTimeout time.Duration,
) error {
	timeoutCtx, cf := context.WithTimeout(ctx, syncTimeout)
	defer cf()

	updated, err := reader.Sync(timeoutCtx)
	if err != nil {
		return err
	}

	if !updated {
		lggr.Debug("no updates found after trying to sync")
	} else {
		lggr.Info("ccip reader sync success")
	}

	return nil
}
