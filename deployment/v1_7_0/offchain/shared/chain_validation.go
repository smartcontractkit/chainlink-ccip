package shared

import (
	"fmt"
	"strings"

	chainsel "github.com/smartcontractkit/chain-selectors"
)

const unknownChainName = "unknown"

type ChainValidationResult struct {
	NOPAlias      string
	MissingChains []uint64
}

func ValidateNOPChainSupport(
	nopAlias string,
	requiredChains []uint64,
	supportedChains []uint64,
) *ChainValidationResult {
	if len(requiredChains) == 0 {
		return nil
	}

	supportedSet := make(map[uint64]bool, len(supportedChains))
	for _, s := range supportedChains {
		supportedSet[s] = true
	}

	var missing []uint64
	for _, required := range requiredChains {
		if !supportedSet[required] {
			missing = append(missing, required)
		}
	}

	if len(missing) == 0 {
		return nil
	}

	return &ChainValidationResult{
		NOPAlias:      nopAlias,
		MissingChains: missing,
	}
}

func FormatChainValidationError(results []ChainValidationResult) error {
	if len(results) == 0 {
		return nil
	}

	var sb strings.Builder
	sb.WriteString("chain support validation failed:\n")

	for _, result := range results {
		sb.WriteString(fmt.Sprintf("  NOP %q missing chain configs for:\n", result.NOPAlias))
		for _, chainSelector := range result.MissingChains {
			chainName := formatChainName(chainSelector)
			sb.WriteString(fmt.Sprintf("    - %d (%s)\n", chainSelector, chainName))
		}
	}

	sb.WriteString("  action: ensure configs are created for the missing chains on the node")

	return fmt.Errorf("%s", sb.String())
}

func formatChainName(chainSelector uint64) string {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return unknownChainName
	}
	chainID, err := chainsel.GetChainIDFromSelector(chainSelector)
	if err != nil {
		return unknownChainName
	}
	details, err := chainsel.GetChainDetailsByChainIDAndFamily(chainID, family)
	if err != nil {
		return unknownChainName
	}
	return details.ChainName
}
