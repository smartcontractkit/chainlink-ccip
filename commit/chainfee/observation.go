package chainfee

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"

	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
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

	supportedChains, err := p.chainSupport.SupportedChains(p.oracleID)
	if err != nil {
		return Observation{}, err
	}

	supportedChains.Remove(p.destChain)
	if supportedChains.Cardinality() == 0 {
		lggr.Info("no supported chains other than dest chain to observe")
		return Observation{}, nil
	}

	supportedChainsSlice := supportedChains.ToSlice()
	sort.Slice(supportedChainsSlice, func(i, j int) bool { return supportedChainsSlice[i] < supportedChainsSlice[j] })

	var (
		feeComponents     = map[cciptypes.ChainSelector]types.ChainFeeComponents{}
		nativeTokenPrices = map[cciptypes.ChainSelector]cciptypes.BigInt{}
		chainFeeUpdates   = map[cciptypes.ChainSelector]Update{}
	)

	eg := new(errgroup.Group)

	// Get the fee components for all available chains that we can read from
	eg.Go(func() error {
		feeComponents = p.ccipReader.GetChainsFeeComponents(ctx, supportedChainsSlice)
		return nil
	})

	// Get the native token prices for all available chains that we can read from
	eg.Go(func() error {
		nativeTokenPrices = p.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, supportedChainsSlice)
		return nil
	})

	// Get the latest chain fee price updates for the source chains and
	// Convert them to a map of chain fee updates
	eg.Go(func() error {
		chainFeeUpdates = feeUpdatesFromTimestampedBig(
			p.ccipReader.GetChainFeePriceUpdate(ctx, supportedChainsSlice),
		)
		return nil
	})

	if err := eg.Wait(); err != nil {
		return Observation{}, fmt.Errorf("unexpected error: %w", err)
	}

	fChain := p.observeFChain(lggr)
	now := time.Now().UTC()

	lggr.Infow("observed fee components",
		"supportedChains", supportedChainsSlice,
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
	updates map[cciptypes.ChainSelector]plugintypes.TimestampedBig,
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
