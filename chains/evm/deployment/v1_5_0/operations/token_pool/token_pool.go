package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "TokenPool"
var Version *semver.Version = semver.MustParse("1.5.0")

var GetToken = contract.NewRead(contract.ReadParams[any, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-token",
	Version:      Version,
	Description:  "Gets the local token address for a TokenPool",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tokenPool *token_pool.TokenPool, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenPool.GetToken(opts)
	},
})
