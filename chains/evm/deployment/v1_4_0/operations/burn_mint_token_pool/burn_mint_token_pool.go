// Package burn_mint_token_pool exposes the v1.4.0 BurnMintTokenPool.
//
// A surviving pool of this type sits behind a v1.5.0 BurnMintTokenPoolAndProxy,
// which is the contract registered in the TokenAdminRegistry. Reads and writes
// here target the legacy pool itself, through an ABI that is a superset of the
// shared TokenPool surface in the token_pool package; only the type-specific
// reads live here.
//
// There is deliberately no deploy operation: deploying a new v1.4.0 pool is
// never correct today.
package burn_mint_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_4_0/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "BurnMintTokenPool"
var Version *semver.Version = semver.MustParse("1.4.0")
var TypeAndVersion = cldf_deployment.NewTypeAndVersion(ContractType, *Version)

var TypeAndVersionRead = contract.NewRead(contract.ReadParams[struct{}, string, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "burn-mint-token-pool:type-and-version",
	Version:      Version,
	Description:  "Gets the type and version of a v1.4.0 BurnMintTokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(pool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, _ struct{}) (string, error) {
		return pool.TypeAndVersion(opts)
	},
})

var GetAllowList = contract.NewRead(contract.ReadParams[struct{}, []common.Address, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "burn-mint-token-pool:get-allow-list",
	Version:      Version,
	Description:  "Gets the allow list of a v1.4.0 BurnMintTokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(pool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, _ struct{}) ([]common.Address, error) {
		return pool.GetAllowList(opts)
	},
})

var GetAllowListEnabled = contract.NewRead(contract.ReadParams[struct{}, bool, *burn_mint_token_pool.BurnMintTokenPool]{
	Name:         "burn-mint-token-pool:get-allow-list-enabled",
	Version:      Version,
	Description:  "Reports whether the allow list is enabled on a v1.4.0 BurnMintTokenPool",
	ContractType: ContractType,
	NewContract:  burn_mint_token_pool.NewBurnMintTokenPool,
	CallContract: func(pool *burn_mint_token_pool.BurnMintTokenPool, opts *bind.CallOpts, _ struct{}) (bool, error) {
		return pool.GetAllowListEnabled(opts)
	},
})
