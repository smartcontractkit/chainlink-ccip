package hooks_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/config/network"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/domain"
	cldverification "github.com/smartcontractkit/chainlink-deployments-framework/engine/cld/verification"

	"github.com/smartcontractkit/chainlink-ccip/deployment/hooks"
)

// concurrentStepsLimit must match hooks.concurrentVerificationsLimit so the test documents the
// expected ceiling for in-flight step goroutines per network.
const concurrentStepsLimit = 5

type noopVerifiable struct{}

func (noopVerifiable) String() string { return "noop" }

func (noopVerifiable) IsVerified(context.Context) (bool, error) { return true, nil }

func (noopVerifiable) Verify(context.Context) error { return nil }

// iterateParallelTestVerifier drives IterateVerifiers with configurable networks and always
// returns a cheap noop Verifiable so the step callback is the locus of concurrency.
type iterateParallelTestVerifier struct {
	networks []network.Network
}

func (v *iterateParallelTestVerifier) FilterNetworks(_ string, _ domain.Domain, _ logger.Logger) (*network.Config, error) {
	return network.NewConfig(v.networks), nil
}

func (v *iterateParallelTestVerifier) NeedsVerification(_ datastore.AddressRef) bool {
	return true
}

func (v *iterateParallelTestVerifier) ForEachNetwork(
	_ context.Context,
	_ network.Network,
	_ uint64,
	_ logger.Logger,
	_ string,
) (hooks.VerifierBuilderForNetwork, bool) {
	return func(_ context.Context, _ datastore.AddressRef) (cldverification.Verifiable, error) {
		return noopVerifiable{}, nil
	}, false
}

// TestIterateVerifiers_VerificationStepsRespectConcurrencyLimit spawns more address refs than the
// per-network verification limit; peak concurrent step executions should reach the limit (not 1),
// proving errgroup-limited parallelism within a single network.
func TestIterateVerifiers_VerificationStepsRespectConcurrencyLimit(t *testing.T) {
	t.Parallel()

	chain, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)

	var (
		mu            sync.Mutex
		inFlight      int
		maxConcurrent int
	)

	verifier := &iterateParallelTestVerifier{
		networks: []network.Network{{
			Type:          network.NetworkTypeMainnet,
			ChainSelector: chain.Selector,
			RPCs:          []network.RPC{{HTTPURL: "http://localhost"}},
		}},
	}

	// AddressRef store keys are (chain, type, version, qualifier) — not address. Use a unique
	// qualifier per row so six contracts on the same chain can coexist.
	addrs := []string{
		"0x1000000000000000000000000000000000000001",
		"0x2000000000000000000000000000000000000001",
		"0x3000000000000000000000000000000000000001",
		"0x4000000000000000000000000000000000000001",
		"0x5000000000000000000000000000000000000001",
		"0x6000000000000000000000000000000000000001",
	}

	ds := datastore.NewMemoryDataStore()
	for i, addr := range addrs {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain.Selector,
			Type:            "C",
			Version:         semver.MustParse("1.0.0"),
			Qualifier:       fmt.Sprintf("p%d", i),
			Address:         addr,
		}))
	}

	cfg, err := verifier.FilterNetworks("", domain.Domain{}, logger.Test(t))
	require.NoError(t, err)

	err = hooks.IterateVerifiers(t.Context(), ds.Seal(), cfg, logger.Test(t), "test", verifier,
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			mu.Lock()
			inFlight++
			if inFlight > maxConcurrent {
				maxConcurrent = inFlight
			}
			cur := inFlight
			mu.Unlock()

			if cur > concurrentStepsLimit {
				t.Errorf("in-flight steps %d exceeds limit %d", cur, concurrentStepsLimit)
			}

			time.Sleep(5 * time.Millisecond)

			mu.Lock()
			inFlight--
			mu.Unlock()
			return nil
		},
	)
	require.NoError(t, err)
	require.Equal(t, concurrentStepsLimit, maxConcurrent,
		"with %d refs and limit %d, peak concurrency should hit the limit", concurrentStepsLimit+1, concurrentStepsLimit)
}

// TestIterateVerifiers_MultipleNetworksRunInParallel uses three networks with one ref each and a
// non-trivial step delay. If network iterations were strictly sequential, wall time would be ~3×
// the delay; with parallel network goroutines, elapsed time should stay near one delay.
func TestIterateVerifiers_MultipleNetworksRunInParallel(t *testing.T) {
	t.Parallel()

	eth, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET.Selector)
	require.True(t, ok)
	poly, ok := chainsel.ChainBySelector(chainsel.POLYGON_MAINNET.Selector)
	require.True(t, ok)
	arb, ok := chainsel.ChainBySelector(chainsel.ETHEREUM_MAINNET_ARBITRUM_1.Selector)
	require.True(t, ok)

	verifier := &iterateParallelTestVerifier{
		networks: []network.Network{
			{Type: network.NetworkTypeMainnet, ChainSelector: eth.Selector, RPCs: []network.RPC{{HTTPURL: "http://localhost"}}},
			{Type: network.NetworkTypeMainnet, ChainSelector: poly.Selector, RPCs: []network.RPC{{HTTPURL: "http://localhost"}}},
			{Type: network.NetworkTypeMainnet, ChainSelector: arb.Selector, RPCs: []network.RPC{{HTTPURL: "http://localhost"}}},
		},
	}

	ds := datastore.NewMemoryDataStore()
	for _, ch := range []chainsel.Chain{eth, poly, arb} {
		require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: ch.Selector,
			Type:            "C",
			Version:         semver.MustParse("1.0.0"),
			Address:         "0x0000000000000000000000000000000000000001",
		}))
	}

	cfg, err := verifier.FilterNetworks("", domain.Domain{}, logger.Test(t))
	require.NoError(t, err)

	const stepDelay = 80 * time.Millisecond
	start := time.Now()
	err = hooks.IterateVerifiers(t.Context(), ds.Seal(), cfg, logger.Test(t), "test", verifier,
		func(_ context.Context, _ cldverification.Verifiable, _ datastore.AddressRef, _ uint64) error {
			time.Sleep(stepDelay)
			return nil
		},
	)
	require.NoError(t, err)
	elapsed := time.Since(start)

	// Sequential network work would be ~3 * stepDelay; parallel network goroutines overlap.
	require.Less(t, elapsed, 2*stepDelay,
		"elapsed %v should be well under 3×%v if networks run in parallel", elapsed, stepDelay)
}
