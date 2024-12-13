package plugincommon

import (
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// ValidateFChain validates that the FChain values are positive.
func ValidateFChain(fChain map[cciptypes.ChainSelector]int) error {
	for chainSelector, f := range fChain {
		if f <= 0 {
			return fmt.Errorf("fChain for chain %d is not positive: %d", chainSelector, f)
		}
	}
	return nil
}
