package v1_2

import (
	"fmt"

	"call-orchestrator-demo/views"

	"github.com/ethereum/go-ethereum/accounts/abi"

	router "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/router"
)

// Parsed ABIs for v1.2 contracts
var (
	RouterABI abi.ABI
)

func init() {
	// Parse the Router ABI once at startup (uses ccv local bindings)
	parsed, err := router.RouterMetaData.GetAbi()
	if err != nil {
		panic(fmt.Sprintf("failed to parse Router ABI: %v", err))
	}
	RouterABI = *parsed

	// Register v1.2 views
	views.Register("evm", "Router", "1.2.0", ViewRouter)
	views.Register("evm", "PriceRegistry", "1.2.0", ViewPriceRegistry)
}
