package logpoller

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/sqlutil"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/mailbox"
	"github.com/smartcontractkit/chainlink-evm/pkg/client"
	"github.com/smartcontractkit/chainlink-evm/pkg/config"
	"github.com/smartcontractkit/chainlink-evm/pkg/heads"
	"github.com/smartcontractkit/chainlink-evm/pkg/logpoller"
)

func NewLogPoller(ctx context.Context, cfg *config.ChainScoped, l logger.Logger, ds sqlutil.DataSource, client client.Client, mailMon *mailbox.Monitor) logpoller.LogPoller {
	headTrackerORM := heads.NewORM(*cfg.EVM().ChainID(), ds, 100)
	headSaver := heads.NewSaver(l, headTrackerORM, cfg.EVM(), cfg.EVM().HeadTracker())
	headBroadcaster := heads.NewBroadcaster(l)
	headTracker := heads.NewTracker(l, client, cfg.EVM(), cfg.EVM().HeadTracker(), headBroadcaster, headSaver, mailMon)
	headTracker.Start(ctx)
	head, err := headTracker.LatestSafeBlock(ctx)
	if err != nil {
		l.Panicf("Error getting latest safe block: %v", err)
	}
	l.Info("Latest safe block", "block", head)

	lpOpts := logpoller.Opts{
		PollPeriod:               cfg.EVM().LogPollInterval(),
		UseFinalityTag:           cfg.EVM().FinalityTagEnabled(),
		FinalityDepth:            int64(cfg.EVM().FinalityDepth()),
		BackfillBatchSize:        int64(cfg.EVM().LogBackfillBatchSize()),
		RPCBatchSize:             int64(cfg.EVM().RPCDefaultBatchSize()),
		KeepFinalizedBlocksDepth: int64(cfg.EVM().LogKeepBlocksDepth()),
		LogPrunePageSize:         int64(cfg.EVM().LogPrunePageSize()),
		BackupPollerBlockDelay:   int64(cfg.EVM().BackupLogPollerBlockDelay()),
		ClientErrors:             cfg.EVM().NodePool().Errors(),
	}

	lpORM, _ := logpoller.NewObservedORM(cfg.EVM().ChainID(), ds, l)
	return logpoller.NewLogPoller(lpORM, client, l, headTracker, lpOpts)
}
