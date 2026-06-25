package mathslib

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	sel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// calculateUsdPerUnitGasPreFix is the old EVM formula. It gives these tests a
// precise "before" value without touching production code.
func calculateUsdPerUnitGasPreFix(sourceGasPrice, usdPerFeeCoin *big.Int) *big.Int {
	tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
	return new(big.Int).Div(tmp, big.NewInt(1e18))
}

var (
	hederaMainnetSelector = ccipocr3.ChainSelector(sel.HEDERA_MAINNET.Selector)
	hederaTestnetSelector = ccipocr3.ChainSelector(sel.HEDERA_TESTNET.Selector)
	sepoliaSelector       = ccipocr3.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
)

// TestHederaWeibarFix_FourCellMatrix covers the bidirectional x
// {with, without fix} matrix from hedera-fees-investigation-v2.md:
//
//	                         X -> Hedera                   Hedera -> X
//	  Without fix      10^10x inflated bug          Correct baseline
//	  With fix         Correct                      Correct, no regression
//
// The Hedera-side fixtures reproduce the exact unit mismatch in the note:
// Hedera RPC reports gas in weibar, while FeeQuoter.getTokenPrice(WHBAR) is in
// tinybar-space. Section 4 explains why 1.5 kept those spaces separate.
//
//	Hedera fixture (X -> Hedera):
//	  eth_gasPrice = 1.1e12 weibar/gas (live Hedera gas price from the note)
//	  HBAR @ ~$0.0767 → usdPerFeeCoin = 7.671684e26
//	    ("WHBAR=7.672e26" from FeeQuoter.getTokenPrice(WHBAR))
//	    Derivation: price × 1e18 / 10^decimals
//	                = 0.07671684 × 1e18 × 1e18 / 1e8
//	                = 7.671684e26  (uses CCIPHome.tokenInfo[HBAR].Decimals = 8)
//
//	Sepolia fixture (Hedera -> X no-regression guard):
//	  eth_gasPrice = 2e10 wei/gas (20 gwei)
//	  ETH @ $3,000 → usdPerFeeCoin = 3000 * 1e18 * 1e18 / 1e18 = 3e21
//	  These Sepolia values are intentionally simple synthetic values; they only
//	  prove non-Hedera EVM chains are unchanged by the Hedera branch.
//
// Expected outputs (each in "1e18 USD per gas unit"):
//
//	Cell 1 (X -> Hedera, no fix): 1.1e12 * 7.671684e26 / 1e18 = 8.4388524e20
//	Cell 2 (X -> Hedera, fix):    1.1e12 * 7.671684e26 / 1e28 = 8.4388524e10
//	Cell 3 (Hedera -> X, no fix): 2e10 * 3e21 / 1e18           = 6e13
//	Cell 4 (Hedera -> X, fix):    identical to Cell 3
//
// The live note uses Hedera testnet. This matrix uses the mainnet selector to
// exercise the same Hedera branch; TestHederaWeibarFix_TestnetCovered asserts
// that the testnet selector takes that branch too.
func TestHederaWeibarFix_FourCellMatrix(t *testing.T) {
	hederaGasWeibar := big.NewInt(1_100_000_000_000)         // 1.1e12 weibar (live Hedera gasPrice)
	hbarUsdPerFeeCoin := MustBigIntSetString("7671684", 20)  // 7.671684e26 (live WHBAR getTokenPrice)
	sepoliaGasWei := big.NewInt(20_000_000_000)              // 2e10 (20 gwei)
	ethUsdPerFeeCoin := MustBigIntSetString("3", 21)         // 3e21 (ETH @ $3,000)

	cases := []struct {
		cell      string
		direction string
		applyFix  bool
		source    ccipocr3.ChainSelector
		gasPrice  *big.Int
		usdPerFee *big.Int
		want      *big.Int
	}{
		{
			cell:      "Cell 1: X→Hedera, WITHOUT fix",
			direction: "Sepolia commit DON observes Hedera as source; Hedera RPC returns weibar",
			applyFix:  false,
			source:    hederaMainnetSelector,
			gasPrice:  hederaGasWeibar,
			usdPerFee: hbarUsdPerFeeCoin,
			want:      MustBigIntSetString("84388524", 13), // 8.4388524e20 — BUGGY (live Sepolia FQ stored value, doc §3.5)
		},
		{
			cell:      "Cell 2: X→Hedera, WITH fix",
			direction: "Same scenario as Cell 1; mathslib applies weibar→tinybar normalization",
			applyFix:  true,
			source:    hederaMainnetSelector,
			gasPrice:  hederaGasWeibar,
			usdPerFee: hbarUsdPerFeeCoin,
			want:      MustBigIntSetString("84388524", 3), // 8.4388524e10 — CORRECT (≈ doc §7 target 8.4e10)
		},
		{
			cell:      "Cell 3: Hedera→X (Sepolia source), WITHOUT fix",
			direction: "Hedera commit DON observes Sepolia as source; baseline correct",
			applyFix:  false,
			source:    sepoliaSelector,
			gasPrice:  sepoliaGasWei,
			usdPerFee: ethUsdPerFeeCoin,
			want:      MustBigIntSetString("6", 13), // 6e13 — correct baseline (~$6e-5/gas)
		},
		{
			cell:      "Cell 4: Hedera→X (Sepolia source), WITH fix — NO REGRESSION",
			direction: "Same as Cell 3 but fix applied; must be byte-identical (no-op for non-Hedera)",
			applyFix:  true,
			source:    sepoliaSelector,
			gasPrice:  sepoliaGasWei,
			usdPerFee: ethUsdPerFeeCoin,
			want:      MustBigIntSetString("6", 13), // 6e13 — identical to Cell 3
		},
	}

	for _, tc := range cases {
		t.Run(tc.cell, func(t *testing.T) {
			var got *big.Int
			if tc.applyFix {
				res, err := CalculateUsdPerUnitGas(tc.source, tc.gasPrice, tc.usdPerFee)
				require.NoError(t, err)
				got = res
			} else {
				got = calculateUsdPerUnitGasPreFix(tc.gasPrice, tc.usdPerFee)
			}
			require.Equal(t, tc.want.String(), got.String(),
				"%s\ndirection: %s\nexpected: %s\n     got: %s",
				tc.cell, tc.direction, tc.want.String(), got.String())
		})
	}
}

// TestHederaWeibarFix_RatioBetweenBugAndFix keeps the core invariant explicit:
// the bad Hedera value is exactly 10^10 times the fixed value.
func TestHederaWeibarFix_RatioBetweenBugAndFix(t *testing.T) {
	hederaGasWeibar := big.NewInt(1_100_000_000_000)
	hbarUsdPerFeeCoin := MustBigIntSetString("7671684", 20)

	buggy := calculateUsdPerUnitGasPreFix(hederaGasWeibar, hbarUsdPerFeeCoin)
	correct, err := CalculateUsdPerUnitGas(hederaMainnetSelector, hederaGasWeibar, hbarUsdPerFeeCoin)
	require.NoError(t, err)

	ratio := new(big.Int).Quo(buggy, correct)
	require.Equal(t, "10000000000", ratio.String(),
		"bug must be exactly 10^10× — got buggy=%s correct=%s ratio=%s",
		buggy.String(), correct.String(), ratio.String())
}

// TestHederaWeibarFix_TestnetCovered protects the live testnet lane from a
// mainnet-only selector check.
func TestHederaWeibarFix_TestnetCovered(t *testing.T) {
	gasWeibar := big.NewInt(1_100_000_000_000)
	hbarUsd := MustBigIntSetString("7671684", 20)

	mainnet, err := CalculateUsdPerUnitGas(hederaMainnetSelector, gasWeibar, hbarUsd)
	require.NoError(t, err)
	testnet, err := CalculateUsdPerUnitGas(hederaTestnetSelector, gasWeibar, hbarUsd)
	require.NoError(t, err)

	require.Equal(t, mainnet.String(), testnet.String(),
		"both Hedera selectors must normalize identically (mainnet=%s testnet=%s)",
		mainnet.String(), testnet.String())
}

// TestHederaWeibarFix_NonHederaEVMUnaffected is the compact no-regression check
// for ordinary EVM chains.
func TestHederaWeibarFix_NonHederaEVMUnaffected(t *testing.T) {
	gasWei := big.NewInt(20_000_000_000) // 20 gwei
	ethUsd := MustBigIntSetString("3", 21)

	withFix, err := CalculateUsdPerUnitGas(sepoliaSelector, gasWei, ethUsd)
	require.NoError(t, err)
	preFix := calculateUsdPerUnitGasPreFix(gasWei, ethUsd)

	require.Equal(t, preFix.String(), withFix.String(),
		"non-Hedera EVM must produce identical output before and after fix (preFix=%s withFix=%s)",
		preFix.String(), withFix.String())
}

// TestHederaWeibarFix_ZeroGasPrice guards the Hedera branch's zero edge case.
func TestHederaWeibarFix_ZeroGasPrice(t *testing.T) {
	hbarUsd := MustBigIntSetString("7671684", 20)

	got, err := CalculateUsdPerUnitGas(hederaMainnetSelector, big.NewInt(0), hbarUsd)
	require.NoError(t, err)
	require.Equal(t, "0", got.String(), "zero gas price must produce zero output")
}
