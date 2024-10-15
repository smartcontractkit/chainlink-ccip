package reader

import (
	"time"

	reader_internal "github.com/smartcontractkit/chainlink-ccip/internal/reader"

	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
)

type RMNHome = reader_internal.RMNHome

func NewRMNHomePoller(
	contractReader contractreader.ContractReaderFacade,
	rmnHomeBoundContract types.BoundContract,
	lggr logger.Logger,
	pollingInterval time.Duration,
) RMNHome {
	return reader_internal.NewRMNHomePoller(
		contractReader,
		rmnHomeBoundContract,
		lggr,
		pollingInterval,
	)
}
