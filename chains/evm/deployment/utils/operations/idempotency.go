package operations

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// ContractIdempotencyKey scopes read/write idempotency to a chain selector and contract address.
func ContractIdempotencyKey(chainSelector uint64, addr common.Address) string {
	return fmt.Sprintf("%d:%s", chainSelector, addr.Hex())
}

// ChainIdempotencyKey scopes deploy idempotency to a chain selector.
func ChainIdempotencyKey(chainSelector uint64) string {
	return fmt.Sprintf("%d", chainSelector)
}
