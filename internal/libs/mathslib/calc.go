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

// CalculateUsdPerUnitGas returns: (sourceGasPrice * usdPerFeeCoin) / 1e18
func CalculateUsdPerUnitGas(sourceChainSelector ccipocr3.ChainSelector, sourceGasPrice *big.Int, usdPerFeeCoin *big.Int) (*big.Int, error) {

	family, err := chainsel.GetSelectorFamily(uint64(sourceChainSelector))
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", sourceChainSelector, err)
	}

	switch family {
	case chainsel.FamilyEVM:
		// (wei / gas) * (usd / eth) * (1 eth / 1e18 wei)  = usd/gas
		tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
		return tmp.Div(tmp, big.NewInt(1e18)), nil

	case chainsel.FamilySolana:
		// (micro lamport / compute units) * (usd / 1 sol) * (1 sol / 1e15 micro lamport)  = usd/cu
		tmp := new(big.Int).Mul(sourceGasPrice, usdPerFeeCoin)
		return tmp.Div(tmp, big.NewInt(1e15)), nil

	default:
		return nil, fmt.Errorf("unsupported family for extra args type %s", family)
	}

}
