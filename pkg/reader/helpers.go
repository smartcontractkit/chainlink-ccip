package reader

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

// bindReaderContract is a generic helper for binding contracts to readers, the addresses input is the same object
// returned by DiscoverContracts.
//
// No error is returned if contractName is not found in the contracts. This allows calling the function before all
// contracts are discovered.
func bindReaderContract[T contractreader.ContractReaderFacade](
	ctx context.Context,
	lggr logger.Logger,
	readers map[cciptypes.ChainSelector]T,
	chainSel cciptypes.ChainSelector,
	contractName string,
	address []byte,
	codec cciptypes.AddressCodec,
) (types.BoundContract, error) {
	if err := validateReaderExistence(readers, chainSel); err != nil {
		return types.BoundContract{}, fmt.Errorf("validate reader existence: %w", err)
	}

	addressStr, err := codec.AddressBytesToString(address, chainSel)
	if err != nil {
		return types.BoundContract{}, fmt.Errorf("unable to convert address bytes to string: %w, address: %v", err, address)
	}

	contract := types.BoundContract{
		Address: addressStr,
		Name:    contractName,
	}

	lggr.Debugw("Binding contract",
		"chainSel", chainSel,
		"contractName", contractName,
		"address", addressStr,
	)
	// Bind the contract address to the reader.
	// If the same address exists -> no-op
	// If the address is changed -> updates the address, overwrites the existing one
	// If the contract not bound -> binds to the new address
	if err := readers[chainSel].Bind(ctx, []types.BoundContract{contract}); err != nil {
		return types.BoundContract{}, fmt.Errorf("unable to bind %s %s for chain %d: %w", contractName, addressStr, chainSel, err)
	}

	return contract, nil
}

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
