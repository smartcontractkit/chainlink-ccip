package mathslib

import (
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
