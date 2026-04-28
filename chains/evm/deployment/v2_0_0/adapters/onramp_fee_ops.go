package adapters

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"

	v1_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
)

var _ v1_0_0_adapters.OnRampFeeContractOps = onRampFeeOpsV200{}

// onRampFeeOpsV200 plugs into v1_0_0_adapters.EVMFeeContractResolver. The v2.0
// OnRamp's DynamicConfig differs from v1.6 (3 fields vs 5) but exposes the
// same FeeQuoter field; this struct uses the v2.0 generated bindings.
type onRampFeeOpsV200 struct{}

func (onRampFeeOpsV200) GetFeeContractAddress(ctx context.Context, chain evm.Chain, onRampAddr common.Address) (common.Address, error) {
	c, err := onrampops.NewOnRampContract(onRampAddr, chain.Client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to bind v2.0 OnRamp at %s: %w", onRampAddr.Hex(), err)
	}
	cfg, err := c.GetDynamicConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call v2.0 OnRamp.getDynamicConfig at %s: %w", onRampAddr.Hex(), err)
	}
	return cfg.FeeQuoter, nil
}
