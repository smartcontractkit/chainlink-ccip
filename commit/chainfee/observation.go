package chainfee

import (
	"context"
	"math/big"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/asynclib"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
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

	if invalidateCache, ok := ctx.Value(consts.InvalidateCacheKey).(bool); ok && invalidateCache {
		p.obs.invalidateCaches(ctx, lggr)
	}

	var (
		feeComponents     map[cciptypes.ChainSelector]types.ChainFeeComponents
		nativeTokenPrices map[cciptypes.ChainSelector]cciptypes.BigInt
		chainFeeUpdates   map[cciptypes.ChainSelector]Update
		fChain            map[cciptypes.ChainSelector]int
	)

	operations := asynclib.AsyncNoErrOperationsMap{
		"getChainsFeeComponents": func(ctx context.Context, l logger.Logger) {
			feeComponents = p.obs.getChainsFeeComponents(ctx, l)
		},
		"getNativeTokenPrices": func(ctx context.Context, l logger.Logger) {
			nativeTokenPrices = p.obs.getNativeTokenPrices(ctx, l)
		},
		"getChainFeePriceUpdates": func(ctx context.Context, l logger.Logger) {
			chainFeeUpdates = p.obs.getChainFeePriceUpdates(ctx, l)
		},
		"observeFChain": func(_ context.Context, l logger.Logger) {
			fChain = p.observeFChain(l)
		},
	}

	asynclib.WaitForAllNoErrOperations(ctx, p.cfg.ChainFeeAsyncObserverSyncTimeout, operations, lggr)
	now := time.Now().UTC()

	chainsWithNativeTokenPrices := mapset.NewSet(maps.Keys(feeComponents)...).
		Intersect(
			mapset.NewSet(maps.Keys(nativeTokenPrices)...),
		)
	chainsWithoutNativeTokenPrices := mapset.NewSet(maps.Keys(feeComponents)...).
		Difference(
			mapset.NewSet(maps.Keys(nativeTokenPrices)...),
		)

	lggr.Infow("observed fee components",
		"feeComponents", feeComponents,
		"nativeTokenPrices", nativeTokenPrices,
		"chainFeeUpdates", chainFeeUpdates,
		"fChain", fChain,
		"timestampNow", now,
		"chainsWithNativeTokenPrices", chainsWithNativeTokenPrices.ToSlice(),
		"chainsWithoutNativeTokenPrices", chainsWithoutNativeTokenPrices.ToSlice(),
	)

	if len(chainsWithNativeTokenPrices.ToSlice()) == 0 {
		lggr.Infow("don't have any chains with native token prices",
			"chainsWithoutNativeTokenPrices", chainsWithoutNativeTokenPrices.ToSlice())
		return Observation{
			FChain:          fChain,
			TimestampNow:    now,
			ChainFeeUpdates: chainFeeUpdates,
		}, nil
	}

	obs := Observation{
		FChain:            fChain,
		FeeComponents:     selectMapKeysInSet(feeComponents, chainsWithNativeTokenPrices),
		NativeTokenPrices: selectMapKeysInSet(nativeTokenPrices, chainsWithNativeTokenPrices),
		ChainFeeUpdates:   chainFeeUpdates,
		TimestampNow:      now,
	}
	return obs, nil
}

// selectMapKeysInSet returns a new map containing only the key-value pairs
// from the input map whose keys are present in the specified set.
// This effectively performs an intersection between the map keys and the set.
func selectMapKeysInSet[T comparable](
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
