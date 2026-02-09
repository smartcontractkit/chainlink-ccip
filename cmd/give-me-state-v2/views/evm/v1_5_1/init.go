package v1_5_1

import (
	"fmt"

	"give-me-state-v2/views"

	"github.com/ethereum/go-ethereum/accounts/abi"

	token_pool "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
)

// Parsed ABI for token pools
var TokenPoolABI abi.ABI

func init() {
	parsed, err := token_pool.TokenPoolMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse TokenPool ABI: %v", err))
	}
	TokenPoolABI = *parsed

	// Register v1.5.1 views
	views.Register("evm", "TokenPoolFactory", "1.5.1", ViewTokenPoolFactory)
	views.Register("evm", "BurnMintTokenPool", "1.5.1", ViewBurnMintTokenPool)
	views.Register("evm", "LockReleaseTokenPool", "1.5.1", ViewLockReleaseTokenPool)
	views.Register("evm", "BurnFromMintTokenPool", "1.5.1", ViewBurnFromMintTokenPool)
	views.Register("evm", "BurnWithFromMintTokenPool", "1.5.1", ViewBurnWithFromMintTokenPool)
}
