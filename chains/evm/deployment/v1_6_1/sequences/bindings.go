package sequences

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
)

// NewTokenPool binds a TokenPool contract and returns it as TokenPoolInterface for use with generated v2 operation factories.
func NewTokenPool(addr common.Address, backend bind.ContractBackend) (gobindings.TokenPoolInterface, error) {
	return gobindings.NewTokenPool(addr, backend)
}
