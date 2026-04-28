package adapters

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"

	v1_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
)

var _ v1_0_0_adapters.OnRampFeeContractOps = onRampFeeOpsV160{}

// onRampFeeOpsV160 plugs into v1_0_0_adapters.EVMFeeContractResolver. For v1.6
// the FeeQuoter holds token-transfer fee config and is reachable via the
// OnRamp's dynamic config.
type onRampFeeOpsV160 struct{}

func (onRampFeeOpsV160) GetFeeContractAddress(ctx context.Context, chain evm.Chain, onRampAddr common.Address) (common.Address, error) {
	c, err := onrampops.NewOnRampContract(onRampAddr, chain.Client)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to bind v1.6 OnRamp at %s: %w", onRampAddr.Hex(), err)
	}
	cfg, err := c.GetDynamicConfig(&bind.CallOpts{Context: ctx})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to call v1.6 OnRamp.getDynamicConfig at %s: %w", onRampAddr.Hex(), err)
	}
	return cfg.FeeQuoter, nil
}
