package mathslib

import (
	"fmt"
	"math"
	"math/big"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Deviates checks if x1 and x2 deviates based on the provided ppb (parts per billion)
// ppb is calculated based on the smaller value of the two
// e.g, if x1 > x2, deviation_parts_per_billion = ((x1 - x2) / x2) * 1e9
func Deviates(x1, x2 *big.Int, ppb int64) bool {
	// handle cases when either one or both of the numbers are 0
	if x1.BitLen() == 0 || x2.BitLen() == 0 {
		return x1.Cmp(x2) != 0
	}
	// ensure x1 > x2
	if x1.Cmp(x2) < 0 {
		x1, x2 = x2, x1
	}
	diff := big.NewInt(0).Sub(x1, x2) // diff = x1-x2
	diff.Mul(diff, big.NewInt(1e9))   // diff = diff * 1e9
	// dividing by the smaller value gives consistent ppb regardless of input order, and supports >100% deviation.
	diff.Div(diff, x2)
	return diff.Cmp(big.NewInt(ppb)) > 0 // diff > ppb
}

// CalculateUsdPerUnitGas accepts source gas price denoted in chain-specific price units,
// and usdPerFeeCoin denoted in 1e18 USD per 1e18 smallest token unit (the FeeQuoter notation),
// and returns the source gas price denoted in 1e18 USD per gas unit.
func CalculateUsdPerUnitGas(
	sourceChainSelector ccipocr3.ChainSelector,
	sourceGasPrice *big.Int,
	usdPerFeeCoin *big.Int) (*big.Int, error) {

	family, err := chainsel.GetSelectorFamily(uint64(sourceChainSelector))
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", sourceChainSelector, err)
	}

	switch family {
	case chainsel.FamilyEVM:
		// In EVM, sourceGasPrice is denoted in wei/gas.
		// Eth has 18 decimals, usdPerFeeCoin represents 1e18 USD * 1e18 / wei.
		// To get 1e18 USD/gas, we have
		//   sourceGasPrice * usdPerFeeCoin / 1e18
		//     = (wei / gas) * (1e18 USD * 1e18 / wei) / 1e18
		//     = 1e18 USD * 1e18 / gas / 1e18
		//     = 1e18 USD / gas
		tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
		return tmp.Div(tmp, big.NewInt(1e18)), nil

	case chainsel.FamilySolana:
		// In SVM, sourceGasPrice is denoted in microlamports/cu, or 1e-15 SOL/cu.
		// sourceGasPrice * 1e3 become in units of 1e-18 SOL/cu == wei/gas.

		// Sol has 9 decimals, usdPerFeeCoin represents 1e18 USD per 1e9 whole sol, or 1e18 USD/1e27wei.
		// usdPerFeeCoin / 1e9 become in units of 1e18 USD * 1e18 / wei.

		// To get 1e18 USD/gas, we have
		//   (sourceGasPrice * 1e3) * (usdPerFeeCoin / 1e9) / 1e18
		//     = (wei / gas) * (1e18 USD * 1e18 / wei) / 1e18
		//     = 1e18 USD * 1e18 / gas / 1e18
		//     = 1e18 USD / gas

		// To keep precision, we want to divide at the very end,
		//   (sourceGasPrice * 1e3) * (usdPerFeeCoin / 1e9) / 1e18
		//      = (sourceGasPrice * usdPerFeeCoin) / 1e24
		tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
		power24 := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(24), nil)
		return tmp.Div(tmp, power24), nil

	default:
		return nil, fmt.Errorf("unsupported family %s", family)
	}

}

// DeviatesOnCurve calculates a deviation threshold on the fly using xNew. For now it's only used for gas price
// deviation calculation. It's important to make sure the order of xNew and xOld is correct when passed into this
// function to get an accurate deviation threshold. This is used instead of Deviates() to fit gas price updates
// into a curve that is more sensitive to gas price changes at lower gas prices and less sensitive at higher gas prices.
// See this article on power law curves: https://en.wikipedia.org/wiki/Power_law
func DeviatesOnCurve(xNew, xOld, noDeviationLowerBound *big.Int) bool {
	// If xNew < noDeviationLowerBound, Deviates should never be true
	if xNew.Cmp(noDeviationLowerBound) < 0 {
		return false
	}

	// We use xNew to generate the threshold so that when going from cheap --> expensive, xNew generates a smaller
	// deviation threshold so we are more likely to update the gas price on chain. When going from expensive --> cheap,
	// xNew generates a larger deviation threshold since it's not as urgent to update the gas price on chain.
	xNewFloat := new(big.Float).SetInt(xNew)
	curveThresholdPPB := calculateCurveThresholdPPB(xNewFloat)
	return Deviates(xNew, xOld, curveThresholdPPB)
}

// calculateCurveThresholdPPB calculates the deviation threshold percentage with x using the formula:
// y = (10e11) / (x^0.665). This sliding scale curve was created by collecting several thousands of historical
// PriceRegistry gas price update samples from chains with volatile gas prices like Zircuit and Mode and then using that
// historical data to define thresholds of gas deviations that were acceptable given their USD value. The gas prices
// (X coordinates) and these new thresholds (Y coordinates) were used to fit this sliding scale curve that returns large
// thresholds at low gas prices and smaller thresholds at higher gas prices. Constructing the curve in USD terms allows
// us to more easily reason about these thresholds and also better translates to non-evm chains. For example, when the
// per unit gas price is 0.000006 USD, (6e12 USDgwei), the curve will output a threshold of around 3,000%. However, when
// the per unit gas price is 0.005 USD (5e15 USDgwei), the curve will output a threshold of only ~30%.
func calculateCurveThresholdPPB(x *big.Float) int64 {
	const constantFactor = 10e11
	const exponent = 0.665
	xFloat64, _ := x.Float64()
	xPower := math.Pow(xFloat64, exponent)
	threshold := constantFactor / xPower

	// Convert curve output percentage to PPB
	thresholdPPB := int64(threshold * 1e7)
	return thresholdPPB
}
