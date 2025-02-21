package mathslib

import (
	"math"
	"math/big"
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

// CalculateUsdPerUnitGas returns: (sourceGasPrice * usdPerFeeCoin) / 1e18
func CalculateUsdPerUnitGas(sourceGasPrice *big.Int, usdPerFeeCoin *big.Int) *big.Int {
	// (wei / gas) * (usd / eth) * (1 eth / 1e18 wei)  = usd/gas
	tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
	return tmp.Div(tmp, big.NewInt(1e18))
}

// DeviatesOnCurve calculates a deviation threshold on the fly using xNew. For now it's only used for gas price
// deviation calculation. It's important to make sure the order of xNew and xOld is correct when passed into this
// function to get an accurate deviation threshold.
func DeviatesOnCurve(xNew, xOld, noDeviationLowerBound *big.Int) bool {
	// If xNew < noDeviationLowerBound, Deviates should never be true
	if xNew.Cmp(noDeviationLowerBound) < 0 {
		return false
	}

	xNewFloat := new(big.Float).SetInt(xNew)
	xNewFloat64, _ := xNewFloat.Float64()

	// We use xNew to generate the threshold so that when going from cheap --> expensive, xNew generates a smaller
	// deviation threshold so we are more likely to update the gas price on chain. When going from expensive --> cheap,
	// xNew generates a larger deviation threshold since it's not as urgent to update the gas price on chain.
	curveThresholdPPB := calculateCurveThresholdPPB(xNewFloat64)
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
func calculateCurveThresholdPPB(x float64) int64 {
	const constantFactor = 10e11
	const exponent = 0.665
	xPower := math.Pow(x, exponent)
	threshold := constantFactor / xPower

	// Convert curve output percentage to PPB
	thresholdPPB := int64(threshold * 1e7)
	return thresholdPPB
}
