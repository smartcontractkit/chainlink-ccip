package token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var (
	ContractType cldf_deployment.ContractType = "TokenPool"
	Version      *semver.Version              = semver.MustParse("1.5.1")
)

var GetToken = contract.NewRead(contract.ReadParams[struct{}, common.Address, *token_pool.TokenPool]{
	Name:         "token-pool:get-token",
	Version:      Version,
	Description:  "Gets the token address managed by the TokenPool 1.5.1 contract",
	ContractType: ContractType,
	NewContract:  token_pool.NewTokenPool,
	CallContract: func(tp *token_pool.TokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return tp.GetToken(opts)
	},
})
