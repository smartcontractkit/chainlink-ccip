package siloed_usdc_token_pool

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "SiloedUSDCTokenPool"

var Version = semver.MustParse("1.7.0")

type LockBoxConfig = siloed_usdc_token_pool.SiloedLockReleaseTokenPoolLockBoxConfig

type AuthorizedCallerArgs = siloed_usdc_token_pool.AuthorizedCallersAuthorizedCallerArgs

var ConfigureLockBoxes = contract.NewWrite(contract.WriteParams[[]LockBoxConfig, *siloed_usdc_token_pool.SiloedUSDCTokenPool]{
	Name:            "siloed-usdc-token-pool:configure-lock-boxes",
	Version:         Version,
	Description:     "Configures lock box mappings on the SiloedUSDCTokenPool",
	ContractType:    ContractType,
	ContractABI:     siloed_usdc_token_pool.SiloedUSDCTokenPoolABI,
	NewContract:     siloed_usdc_token_pool.NewSiloedUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*siloed_usdc_token_pool.SiloedUSDCTokenPool, []LockBoxConfig],
	Validate: func(configs []LockBoxConfig) error {
		for i, cfg := range configs {
			if cfg.LockBox == (common.Address{}) {
				return fmt.Errorf("lock box config %d has zero address", i)
			}
		}
		return nil
	},
	CallContract: func(pool *siloed_usdc_token_pool.SiloedUSDCTokenPool, opts *bind.TransactOpts, args []LockBoxConfig) (*types.Transaction, error) {
		return pool.ConfigureLockBoxes(opts, args)
	},
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *siloed_usdc_token_pool.SiloedUSDCTokenPool]{
	Name:            "siloed-usdc-token-pool:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Applies authorized caller updates on the SiloedUSDCTokenPool",
	ContractType:    ContractType,
	ContractABI:     siloed_usdc_token_pool.SiloedUSDCTokenPoolABI,
	NewContract:     siloed_usdc_token_pool.NewSiloedUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*siloed_usdc_token_pool.SiloedUSDCTokenPool, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(pool *siloed_usdc_token_pool.SiloedUSDCTokenPool, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return pool.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
