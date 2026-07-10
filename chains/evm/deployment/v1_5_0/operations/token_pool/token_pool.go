package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
)

var ContractType cldf_deployment.ContractType = "TokenPool"
var Version *semver.Version = semver.MustParse("1.5.0")

func NewReadGetToken(c *gobindings.TokenPool) *cld_ops.Operation[contract.FunctionInput[struct{}], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.TokenPool]{
		Name:         "token-pool:get-token",
		Version:      Version,
		Description:  "Gets the local token address for a TokenPool",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tokenPool *gobindings.TokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
			return tokenPool.GetToken(opts)
		},
	})
}
