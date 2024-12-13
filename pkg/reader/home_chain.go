package reader

import (
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	reader_internal "github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

type HomeChain = reader_internal.HomeChain

type ChainConfig = reader_internal.ChainConfig

type ChainConfigInfo = reader_internal.ChainConfigInfo

type OCR3ConfigWithMeta = reader_internal.OCR3ConfigWithMeta

type ActiveAndCandidate = reader_internal.ActiveAndCandidate

type OCR3Config = reader_internal.OCR3Config

type OCR3Node = reader_internal.OCR3Node

func NewObservedHomeChainReader(
	homeChainReader types.ContractReader,
	lggr logger.Logger,
	pollingInterval time.Duration,
	ccipConfigBoundContract types.BoundContract,
	chainID string,
) HomeChain {
	return reader_internal.NewHomeChainConfigPoller(
		contractreader.NewObserverReader(
			homeChainReader,
			lggr,
			chainID,
		),
		lggr,
		pollingInterval,
		ccipConfigBoundContract,
	)
}
