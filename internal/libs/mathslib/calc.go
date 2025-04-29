package mathslib

import (
	"fmt"
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
		// Eth has 18 decimals, usdPerFeeCoin represents 1e18 USD/1e18 wei.
		// To get 1e18 USD/gas, we multiply sourceGasPrice by usdPerFeeCoin and divide by 1e18.
		tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
		return tmp.Div(tmp, big.NewInt(1e18)), nil

	case chainsel.FamilySolana:
		// In SVM, sourceGasPrice is denoted in microlamports/cu, or 1e-15 SOL/cu.
		// We need to multiply by 1e3 to be in units of 1e-18 SOL/cu, equivalent to wei/gas.

		// Sol has 9 decimals, usdPerFeeCoin represents 1e18 USD per 1e9 whole sol, or 1e18 USD/1e27wei.
		// We need to divide by 1e9 to be in units of 1e18 USD/1e18 wei, equivalent to EVM model.

		// To keep precision, we want to divide at the very end.
		// The above cancels out, the net result is we need to further divide by 1e6,
		// or 1e18 * 1e6 = 1e24 in total.
		tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
		power24 := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(24), nil)
		return tmp.Div(tmp, power24), nil

	default:
		return nil, fmt.Errorf("unsupported family")
	}

}
