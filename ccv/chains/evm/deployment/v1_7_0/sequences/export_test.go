package sequences

import (
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// MaybeAddRouterOnRampsAddsConfigArgForTest exposes maybeAddRouterOnRampsAddsConfigArg for testing.
func MaybeAddRouterOnRampsAddsConfigArgForTest(b cldf_ops.Bundle, chain evm.Chain, input lanes.UpdateLanesInput) ([]router.OnRamp, error) {
	return maybeAddRouterOnRampsAddsConfigArg(b, chain, input)
}
