package chainfee

import (
	"context"
	"math/big"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Observation will make several calls to fetch:
// - chain fee components
// - native token prices
// - existing chain fee price updates
// - fChain
// The timestamp of the observation is also recorded for consensus purposes.
// Read the Outcome doc for more details about how all this information are used to generate updated chain fees.
func (p *processor) Observation(
	ctx context.Context,
	_ Outcome,
	_ Query,
) (Observation, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	feeComponents := p.obs.getChainsFeeComponents(ctx, lggr)
	nativeTokenPrices := p.obs.getNativeTokenPrices(ctx, lggr)
	chainFeeUpdates := p.obs.getChainFeePriceUpdates(ctx, lggr)
	fChain := p.observeFChain(lggr)
	now := time.Now().UTC()

	lggr.Infow("observed fee components",
		"feeComponents", feeComponents,
		"nativeTokenPrices", nativeTokenPrices,
		"chainFeeUpdates", chainFeeUpdates,
		"fChain", fChain,
		"timestampNow", now,
	)

	uniqueChains := mapset.NewSet[cciptypes.ChainSelector](maps.Keys(feeComponents)...)
	uniqueChains = uniqueChains.Intersect(mapset.NewSet(maps.Keys(nativeTokenPrices)...))

	if len(uniqueChains.ToSlice()) == 0 {
		lggr.Info("observations don't have any unique chains")
		return Observation{}, nil
	}

	obs := Observation{
		FChain:            fChain,
		FeeComponents:     filterMapByUniqueChains(feeComponents, uniqueChains),
		NativeTokenPrices: filterMapByUniqueChains(nativeTokenPrices, uniqueChains),
		ChainFeeUpdates:   chainFeeUpdates,
		TimestampNow:      now,
	}
	return obs, nil
}

// filterMapBySet filters a map based on the keys present in the set.
func filterMapByUniqueChains[T comparable](
	m map[cciptypes.ChainSelector]T,
	s mapset.Set[cciptypes.ChainSelector],
) map[cciptypes.ChainSelector]T {
	filtered := make(map[cciptypes.ChainSelector]T)
	for k, v := range m {
		if s.Contains(k) {
			filtered[k] = v
		}
	}
	return filtered
}

func (p *processor) observeFChain(lggr logger.Logger) map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

func feeUpdatesFromTimestampedBig(
	updates map[cciptypes.ChainSelector]cciptypes.TimestampedBig,
) map[cciptypes.ChainSelector]Update {
	chainFeeUpdates := make(map[cciptypes.ChainSelector]Update, len(updates))
	for chain, u := range updates {
		chainFeeUpdates[chain] = Update{
			ChainFee:  fromPackedFee(u.Value.Int),
			Timestamp: u.Timestamp,
		}
	}
	return chainFeeUpdates
}

// fromPackedFee creates a new Update
// @param packedFee: Is the fee components packed into a single big.Int
// packedFee = (dataAvFeeUSD << 112) | executionFeeUSD
func fromPackedFee(packedFee *big.Int) ComponentsUSDPrices {
	ones112 := big.NewInt(0)
	for i := 0; i < 112; i++ {
		ones112 = ones112.SetBit(ones112, i, 1)
	}

	execFee := new(big.Int).And(packedFee, ones112)
	daFee := new(big.Int).Rsh(packedFee, 112)
	return ComponentsUSDPrices{
		ExecutionFeePriceUSD: execFee,
		DataAvFeePriceUSD:    daFee,
	}
}
