package chainfee_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	sel "github.com/smartcontractkit/chain-selectors"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	commontypes2 "github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	mockchainsupport "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	mockhomechain "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	mockccipreader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

// Package chainfee_test exercises the public chainfee outcome path. The Hedera
// weibar->tinybar normalization lives in mathslib and is invoked from
// outcome.go before the commit DON emits destChainGasPrices.
//
// The 4-cell matrix this test covers:
//
//	                         X -> Hedera                   Hedera -> X
//	  Without fix      10^10x inflated bug          Correct baseline
//	  With fix         Correct                      Correct, no regression
//
// Each cell is asserted on outcome.GasPrices[chain]: the packed
// (daFeeUSD<<112)|execFeeUSD value written to the destination FQ as
// destChainGasPrices[observedChain].

var (
	hederaMainnetSel = cciptypes.ChainSelector(sel.HEDERA_MAINNET.Selector)
	sepoliaSel       = cciptypes.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
)

// Hedera fixtures reproduce the exact on-chain values from
// hedera-fees-investigation-v2.md. They encode the Section 4 mismatch: Hedera
// gas arrives in weibar, while FeeQuoter.getTokenPrice(WHBAR) is tinybar-space.
//
//	Hedera: eth_gasPrice = 1.1e12 weibar (live Hedera RPC, doc §3),
//	        HBAR @ ~$0.0767 → usdPerFeeCoin = 7.671684e26
//	        (live FeeQuoter.getTokenPrice(WHBAR))
//
// Sepolia fixtures are simple synthetic non-Hedera EVM values. They are not
// the live table values; they only prove the Hedera branch is a no-op for
// ordinary EVM chains.
//
//	Sepolia: eth_gasPrice = 2e10 wei (20 gwei),
//	         ETH @ $3,000 → usdPerFeeCoin = 3e21
var (
	hederaGasWeibar   = big.NewInt(1_100_000_000_000)             // 1.1e12 weibar (live Hedera gasPrice)
	hbarUsdPerFeeCoin = mustBigStr("767168400000000000000000000") // 7.671684e26 (live WHBAR getTokenPrice)
	sepoliaGasWei     = big.NewInt(20_000_000_000)                // 2e10 (20 gwei)
	ethUsdPerFeeCoin  = mustBigStr("3000000000000000000000")      // 3e21 (ETH @ $3,000)

	// Expected execFee outputs from mathslib (1e18 USD per gas unit).
	wantHederaExecFeeWithFix = mustBigStr("84388524000")           // 8.4388524e10, the corrected Hedera gas price.
	wantHederaExecFeePreFix  = mustBigStr("843885240000000000000") // 8.4388524e20, the inflated value from the note.
	wantSepoliaExecFee       = mustBigStr("60000000000000")        // 6e13, synthetic non-Hedera baseline.
)

func mustBigStr(s string) *big.Int {
	v, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("bad bigint: " + s)
	}
	return v
}

// preFixCalculateUsdPerUnitGas is the old EVM formula. It constructs the
// "without fix" sentinel values used in this file.
func preFixCalculateUsdPerUnitGas(gasPrice, usdPerFeeCoin *big.Int) *big.Int {
	tmp := new(big.Int).Mul(gasPrice, usdPerFeeCoin)
	return new(big.Int).Div(tmp, big.NewInt(1e18))
}

// unpackExecFee extracts the lower 112 bits of a packed gas price; FeeQuoter
// uses that field for execution-fee pricing.
func unpackExecFee(packed *big.Int) *big.Int {
	mask := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 112), big.NewInt(1))
	return new(big.Int).And(packed, mask)
}

// TestIntegration_HederaWeibarFix_FourCellMatrix runs the chainfee processor
// with Hedera and Sepolia as observed chains, then checks the gas prices that
// would be emitted to FeeQuoter.
//
// Test wiring:
//   - destChain = Sepolia (any destination works for this outcome path)
//   - 5 oracles producing identical observations (2f+1 consensus, f=2)
//   - feeComponents observed for both Hedera and Sepolia
//   - chainFeeUpdates timestamped 2x the heartbeat ago to force emission
func TestIntegration_HederaWeibarFix_FourCellMatrix(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)
	destChain := sepoliaSel
	heartbeat := commonconfig.MustNewDuration(time.Minute)
	staleTs := time.Now().UTC().Add(-2 * heartbeat.Duration())
	nowTs := time.Now().UTC()
	const numOracles = 5

	// ---- mocks ----
	homeChain := mockhomechain.NewMockHomeChain(t)
	homeChain.EXPECT().GetChainConfig(mock.Anything).Return(reader.ChainConfig{
		FChain:         1,
		SupportedNodes: mapset.NewSet(libocrtypes.PeerID{1}, libocrtypes.PeerID{2}),
		Config: chainconfig.ChainConfig{
			GasPriceDeviationPPB:    cciptypes.NewBigInt(big.NewInt(1)),
			DAGasPriceDeviationPPB:  cciptypes.NewBigInt(big.NewInt(1)),
			OptimisticConfirmations: 1,
		},
	}, nil).Maybe()

	ccipReader := mockccipreader.NewMockCCIPReader(t)
	chainSupport := mockchainsupport.NewMockChainSupport(t)

	// ---- processor (public API) ----
	proc := chainfee.NewProcessor(
		lggr,
		commontypes.OracleID(0),
		destChain,
		homeChain,
		ccipReader,
		pluginconfig.CommitOffchainConfig{
			RemoteGasPriceBatchWriteFrequency: *heartbeat,
			ChainFeeAsyncObserverDisabled:     true, // we inject observations directly
		},
		chainSupport,
		1, // fRoleDON
		plugincommon.NoopReporter{},
	)
	t.Cleanup(func() { _ = proc.Close() })

	// ---- observation: Hedera + Sepolia as observed source chains ----
	observation := chainfee.Observation{
		FeeComponents: map[cciptypes.ChainSelector]commontypes2.ChainFeeComponents{
			hederaMainnetSel: {
				ExecutionFee:        hederaGasWeibar, // weibar from Hedera RPC
				DataAvailabilityFee: big.NewInt(0),
			},
			sepoliaSel: {
				ExecutionFee:        sepoliaGasWei,
				DataAvailabilityFee: big.NewInt(0),
			},
		},
		NativeTokenPrices: map[cciptypes.ChainSelector]cciptypes.BigInt{
			hederaMainnetSel: cciptypes.NewBigInt(hbarUsdPerFeeCoin),
			sepoliaSel:       cciptypes.NewBigInt(ethUsdPerFeeCoin),
		},
		ChainFeeUpdates: map[cciptypes.ChainSelector]chainfee.Update{
			// Stale (>heartbeat ago) → outcome must emit fresh price.
			hederaMainnetSel: {Timestamp: staleTs, ChainFee: chainfee.ComponentsUSDPrices{
				ExecutionFeePriceUSD: big.NewInt(1), DataAvFeePriceUSD: big.NewInt(0),
			}},
			sepoliaSel: {Timestamp: staleTs, ChainFee: chainfee.ComponentsUSDPrices{
				ExecutionFeePriceUSD: big.NewInt(1), DataAvFeePriceUSD: big.NewInt(0),
			}},
		},
		FChain: map[cciptypes.ChainSelector]int{
			destChain:        1,
			hederaMainnetSel: 1,
			sepoliaSel:       1,
		},
		TimestampNow: nowTs,
	}

	aos := make([]plugincommon.AttributedObservation[chainfee.Observation], numOracles)
	for i := 0; i < numOracles; i++ {
		aos[i] = plugincommon.AttributedObservation[chainfee.Observation]{
			OracleID:    commontypes.OracleID(i),
			Observation: observation,
		}
	}

	// ---- execute ----
	outcome, err := proc.Outcome(ctx, chainfee.Outcome{}, chainfee.Query{}, aos)
	require.NoError(t, err)
	require.Len(t, outcome.GasPrices, 2, "expected gas price entries for Hedera and Sepolia")

	gotByChain := make(map[cciptypes.ChainSelector]*big.Int)
	for _, gp := range outcome.GasPrices {
		gotByChain[gp.ChainSel] = gp.GasPrice.Int
	}

	// Cell 1/2: Hedera as the observed chain.
	// Without the fix it would emit 8.4388524e20; with the fix it emits
	// 8.4388524e10, matching the corrected gas price in the note.
	t.Run("Cell_1_and_2_X_to_Hedera", func(t *testing.T) {
		packedHedera, ok := gotByChain[hederaMainnetSel]
		require.True(t, ok, "Hedera entry missing from outcome.GasPrices")
		gotExecFee := unpackExecFee(packedHedera)

		require.Equal(t, wantHederaExecFeeWithFix.String(), gotExecFee.String(),
			"Cell 2: Hedera execFee with fix expected %s, got %s",
			wantHederaExecFeeWithFix.String(), gotExecFee.String())

		// Cell 1 sentinel: the old formula produces the inflated value.
		preFix := preFixCalculateUsdPerUnitGas(hederaGasWeibar, hbarUsdPerFeeCoin)
		require.Equal(t, wantHederaExecFeePreFix.String(), preFix.String(),
			"Cell 1 sentinel: pre-fix expected %s, got %s",
			wantHederaExecFeePreFix.String(), preFix.String())
		require.NotEqual(t, preFix.String(), gotExecFee.String(),
			"Cell 1→2: outcome must NOT match the buggy pre-fix value")

		// Bug magnitude: fixed outcome * 10^10 == pre-fix outcome.
		ratio := new(big.Int).Quo(preFix, gotExecFee)
		require.Equal(t, "10000000000", ratio.String(),
			"bug magnitude must be exactly 10^10× (preFix/withFix)")
	})

	// Cell 3/4: Sepolia as the observed chain. The Hedera branch must be a no-op.
	t.Run("Cell_3_and_4_Hedera_to_X_no_regression", func(t *testing.T) {
		packedSepolia, ok := gotByChain[sepoliaSel]
		require.True(t, ok, "Sepolia entry missing from outcome.GasPrices")
		gotExecFee := unpackExecFee(packedSepolia)

		require.Equal(t, wantSepoliaExecFee.String(), gotExecFee.String(),
			"Cell 4: Sepolia execFee expected %s, got %s",
			wantSepoliaExecFee.String(), gotExecFee.String())

		// Cell 3 sentinel: ordinary EVM math is identical before and after the fix.
		preFix := preFixCalculateUsdPerUnitGas(sepoliaGasWei, ethUsdPerFeeCoin)
		require.Equal(t, preFix.String(), gotExecFee.String(),
			"Cell 3↔4: non-Hedera chains must produce identical results before and after fix")
	})
}

// TestIntegration_HederaWeibarFix_TestnetSelector protects the live Hedera
// testnet lane from a mainnet-only selector check.
func TestIntegration_HederaWeibarFix_TestnetSelector(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)
	destChain := sepoliaSel
	hederaTestnetSel := cciptypes.ChainSelector(sel.HEDERA_TESTNET.Selector)
	heartbeat := commonconfig.MustNewDuration(time.Minute)
	staleTs := time.Now().UTC().Add(-2 * heartbeat.Duration())
	const numOracles = 5

	homeChain := mockhomechain.NewMockHomeChain(t)
	homeChain.EXPECT().GetChainConfig(mock.Anything).Return(reader.ChainConfig{
		FChain:         1,
		SupportedNodes: mapset.NewSet(libocrtypes.PeerID{1}),
		Config: chainconfig.ChainConfig{
			GasPriceDeviationPPB:    cciptypes.NewBigInt(big.NewInt(1)),
			DAGasPriceDeviationPPB:  cciptypes.NewBigInt(big.NewInt(1)),
			OptimisticConfirmations: 1,
		},
	}, nil).Maybe()

	proc := chainfee.NewProcessor(
		lggr,
		commontypes.OracleID(0),
		destChain,
		homeChain,
		mockccipreader.NewMockCCIPReader(t),
		pluginconfig.CommitOffchainConfig{
			RemoteGasPriceBatchWriteFrequency: *heartbeat,
			ChainFeeAsyncObserverDisabled:     true,
		},
		mockchainsupport.NewMockChainSupport(t),
		1,
		plugincommon.NoopReporter{},
	)
	t.Cleanup(func() { _ = proc.Close() })

	obs := chainfee.Observation{
		FeeComponents: map[cciptypes.ChainSelector]commontypes2.ChainFeeComponents{
			hederaTestnetSel: {
				ExecutionFee:        hederaGasWeibar,
				DataAvailabilityFee: big.NewInt(0),
			},
		},
		NativeTokenPrices: map[cciptypes.ChainSelector]cciptypes.BigInt{
			hederaTestnetSel: cciptypes.NewBigInt(hbarUsdPerFeeCoin),
		},
		ChainFeeUpdates: map[cciptypes.ChainSelector]chainfee.Update{
			hederaTestnetSel: {Timestamp: staleTs, ChainFee: chainfee.ComponentsUSDPrices{
				ExecutionFeePriceUSD: big.NewInt(1), DataAvFeePriceUSD: big.NewInt(0),
			}},
		},
		FChain: map[cciptypes.ChainSelector]int{
			destChain:        1,
			hederaTestnetSel: 1,
		},
		TimestampNow: time.Now().UTC(),
	}

	aos := make([]plugincommon.AttributedObservation[chainfee.Observation], numOracles)
	for i := 0; i < numOracles; i++ {
		aos[i] = plugincommon.AttributedObservation[chainfee.Observation]{
			OracleID:    commontypes.OracleID(i),
			Observation: obs,
		}
	}

	outcome, err := proc.Outcome(ctx, chainfee.Outcome{}, chainfee.Query{}, aos)
	require.NoError(t, err)
	require.Len(t, outcome.GasPrices, 1)

	got := unpackExecFee(outcome.GasPrices[0].GasPrice.Int)
	require.Equal(t, wantHederaExecFeeWithFix.String(), got.String(),
		"Hedera testnet selector must apply the same normalization as mainnet")
}
