package reader

import (
	"fmt"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
)

const (
	// HomeChainPollingInterval is the interval at which the home chain is polled for updates.
	// It should be used by RMNHome and CCIPHome to poll the home chain for updates.
	// Ethereum was selected for the home chain for CCIP, therefore polling more frequent
	// than block time doesn't bring any value.
	// We selected 15 seconds for simplicity, but this could be extended even further as
	// we accept some delay when fetching the configuration updates.
	// It's advised to use wrap that interval with some jitter to avoid congestion.
	HomeChainPollingInterval = 15 * time.Second
)

func validateReaderExistence[T contractreader.ContractReaderFacade](
	readers map[cciptypes.ChainSelector]T,
	chains ...cciptypes.ChainSelector,
) error {
	for _, ch := range chains {
		_, exists := readers[ch]
		if !exists {
			return fmt.Errorf("chain %d: %w", ch, ErrContractReaderNotFound)
		}
	}
	return nil
}

func getChainAccessor(
	accessors map[cciptypes.ChainSelector]cciptypes.ChainAccessor,
	chainSelector cciptypes.ChainSelector,
) (cciptypes.ChainAccessor, error) {
	if accessor, exists := accessors[chainSelector]; exists {
		return accessor, nil
	}
	return nil, fmt.Errorf("chain %d: %w", chainSelector, ErrChainAccessorNotFound)
}
