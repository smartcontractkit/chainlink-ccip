package adapters

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"

	v1_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
)

var _ v1_0_0_adapters.OnRampFeeContractOps = onRampFeeOpsV150{}

// onRampFeeOpsV150 plugs into v1_0_0_adapters.EVMFeeContractResolver. For v1.5
// the EVM2EVMOnRamp itself holds token-transfer fee config, so the on-ramp
// address is the fee-contract address.
type onRampFeeOpsV150 struct{}

func (onRampFeeOpsV150) GetFeeContractAddress(_ context.Context, _ evm.Chain, onRampAddr common.Address) (common.Address, error) {
	return onRampAddr, nil
}
