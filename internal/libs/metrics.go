package libs

import (
	sel "github.com/smartcontractkit/chain-selectors"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// GetChainInfoFromSelector returns the chain ID and family for a given chain selector
// If the selector is invalid, it returns "unknown" for both values and doesn't raise any error.
// This case should never happen in practice, but it is handled gracefully to avoid panics.
// It's very convenient to use this function for Prometheus metrics
func GetChainInfoFromSelector(selector cciptypes.ChainSelector) (chainFamily string, chainID string, ok bool) {
	chainID, err1 := sel.GetChainIDFromSelector(uint64(selector))
	chainFamily, err2 := sel.GetSelectorFamily(uint64(selector))

	if err1 != nil || err2 != nil {
		return "unknown", "unknown", false
	}

	return chainFamily, chainID, true
}

func GetNameFromIDAndFamily(chainID string, family string) (string, error) {
	details, err := sel.GetChainDetailsByChainIDAndFamily(chainID, family)
	if err != nil {
		return "", err
	}
	return details.ChainName, nil
}
