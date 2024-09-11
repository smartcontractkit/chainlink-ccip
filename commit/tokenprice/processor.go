package tokenprice

import (
	"context"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/shared"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type processor struct {
	oracleID         commontypes.OracleID
	lggr             logger.Logger
	cfg              pluginconfig.CommitPluginConfig
	chainSupport     plugincommon.ChainSupport
	tokenPriceReader reader.PriceReader
	homeChain        reader.HomeChain
	bigF             int
}

// nolint: revive
func NewProcessor(
	oracleID commontypes.OracleID,
	lggr logger.Logger,
	cfg pluginconfig.CommitPluginConfig,
	chainSupport plugincommon.ChainSupport,
	tokenPriceReader reader.PriceReader,
	homeChain reader.HomeChain,
	bigF int,
) *processor {
	return &processor{
		oracleID:         oracleID,
		lggr:             lggr,
		cfg:              cfg,
		chainSupport:     chainSupport,
		tokenPriceReader: tokenPriceReader,
		homeChain:        homeChain,
		bigF:             bigF,
	}
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

func (p *processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {

	fChain := p.ObserveFChain()
	if len(fChain) == 0 {
		return Observation{}, nil
	}

	return Observation{
		FeedTokenPrices:       p.ObserveFeedTokenPrices(ctx),
		FeeQuoterTokenUpdates: p.ObserveFeeQuoterTokenUpdates(ctx),
		FChain:                fChain,
		Timestamp:             time.Now().UTC(),
	}, nil
}

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao shared.AttributedObservation[Observation],
) error {
	//TODO: Validate token prices
	return nil
}

func (p *processor) Outcome(
	_ Outcome,
	_ Query,
	aos []shared.AttributedObservation[Observation],
) (Outcome, error) {
	// If set to zero, no prices will be reported (i.e keystone feeds would be active).
	if p.cfg.OffchainConfig.TokenPriceBatchWriteFrequency.Duration() == 0 {
		p.lggr.Debugw("TokenPriceBatchWriteFrequency is set to zero, no prices will be reported")
		return Outcome{}, nil
	}

	consensusObservation, err := p.getConsensusObservation(aos)
	if err != nil {
		return Outcome{}, err
	}

	tokenPriceOutcome := p.selectTokensForUpdate(consensusObservation)
	return Outcome{
		TokenPrices: tokenPriceOutcome,
	}, nil
}

func validateObservedTokenPrices(tokenPrices []cciptypes.TokenPrice) error {
	tokensWithPrice := mapset.NewSet[types.Account]()
	for _, t := range tokenPrices {
		if tokensWithPrice.Contains(t.TokenID) {
			return fmt.Errorf("duplicate token price for token: %s", t.TokenID)
		}
		tokensWithPrice.Add(t.TokenID)

		if t.Price.IsEmpty() {
			return fmt.Errorf("token price must not be empty")
		}
	}
	return nil
}

var _ shared.PluginProcessor[Query, Observation, Outcome] = &processor{}
