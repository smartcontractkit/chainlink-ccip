package chainfee

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
)

type processor struct {
	destChain                        cciptypes.ChainSelector
	lggr                             logger.Logger
	homeChain                        reader.HomeChain
	ccipReader                       readerpkg.CCIPReader
	ChainFeePriceBatchWriteFrequency commonconfig.Duration
	fRoleDON                         int
}

// nolint: revive
func NewProcessor(
	lggr logger.Logger,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	chainFeePriceBatchWriteFrequency commonconfig.Duration,
	fRoleDON int,
) *processor {
	return &processor{
		lggr:                             lggr,
		destChain:                        destChain,
		homeChain:                        homeChain,
		ccipReader:                       ccipReader,
		ChainFeePriceBatchWriteFrequency: chainFeePriceBatchWriteFrequency,
		fRoleDON:                         fRoleDON,
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
	feeComponents := p.ccipReader.GetAvailableChainsFeeComponents(ctx)
	nativeTokenPrices := p.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, maps.Keys(feeComponents))
	chainFeePriceUpdates := p.ccipReader.GetChainFeePriceUpdate(ctx, maps.Keys(feeComponents))
	latestTimestamps := make(map[cciptypes.ChainSelector]time.Time, len(chainFeePriceUpdates))
	for chain, update := range chainFeePriceUpdates {
		latestTimestamps[chain] = update.Timestamp
	}

	return Observation{
		FChain:                p.ObserveFChain(),
		FeeComponents:         feeComponents,
		NativeTokenPrices:     nativeTokenPrices,
		ChainFeeLatestUpdates: latestTimestamps,
		Timestamp:             time.Now().UTC(),
	}, nil
}

func (p *processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {

	consensusObs, err := p.getConsensusObservation(aos)
	if err != nil {
		return Outcome{}, err
	}
	// No need to update yet
	if len(consensusObs.FeeComponents) == 0 {
		return Outcome{}, nil
	}

	// Stop early if earliest updated timestamp is still fresh
	earliestUpdateTime := consensus.EarliestTimestamp(maps.Values(consensusObs.ChainFeeLatestUpdates))
	nextUpdateTime := earliestUpdateTime.Add(p.ChainFeePriceBatchWriteFrequency.Duration())
	if nextUpdateTime.After(consensusObs.Timestamp) {
		return Outcome{}, nil
	}

	gasPrices := make([]cciptypes.GasPriceChain, 0, len(consensusObs.FeeComponents))
	for chain, feeComp := range consensusObs.FeeComponents {
		// GasPrice is a Bitwise operation here like:
		// (dataAvFeeNative * nativeTokenPriceUSD) << 112 | (executionFeeNative * nativeTokenPrice)
		// e.g. (dataAvFeeWei * WeiUSD) << 112 | (executionFeeWei * WeiUSD)
		// nolint:lll
		// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
		nativeTokenPriceUSD := consensusObs.NativeTokenPrices[chain].Int

		dataAvailabilityPriceUSD := new(big.Int).Mul(nativeTokenPriceUSD, feeComp.DataAvailabilityFee)
		execPriceUSD := new(big.Int).Mul(nativeTokenPriceUSD, feeComp.ExecutionFee)

		// Price is the combination ofr both dataAvailabilityPriceUSD and execPriceUSD
		price := dataAvailabilityPriceUSD.Lsh(dataAvailabilityPriceUSD, 112)
		combinedPricesUSD := new(big.Int).Or(price, execPriceUSD)

		gasPrice := cciptypes.GasPriceChain{
			ChainSel: chain,
			GasPrice: cciptypes.NewBigInt(combinedPricesUSD),
		}
		gasPrices = append(gasPrices, gasPrice)
	}

	// sort gasPrices based on chainSel
	sort.Slice(gasPrices, func(i, j int) bool {
		return gasPrices[i].ChainSel < gasPrices[j].ChainSel
	})

	return Outcome{
		GasPrices: gasPrices,
	}, nil
}

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	// TODO: Validate token prices
	return nil
}

func validateObservedGasPrices(gasPrices []cciptypes.GasPriceChain) error {
	// Duplicate gas prices must not appear for the same chain and must not be empty.
	gasPriceChains := mapset.NewSet[cciptypes.ChainSelector]()
	for _, g := range gasPrices {
		if gasPriceChains.Contains(g.ChainSel) {
			return fmt.Errorf("duplicate gas price for chain %d", g.ChainSel)
		}
		gasPriceChains.Add(g.ChainSel)
		if g.GasPrice.IsEmpty() {
			return fmt.Errorf("gas price must not be empty")
		}
	}

	return nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}
