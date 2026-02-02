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

type ConstructorArgs struct {
	Token              common.Address
	LocalTokenDecimals uint8
	AdvancedPoolHooks  common.Address
	RMNProxy           common.Address
	Router             common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "siloed-usdc-token-pool:deploy",
	Version:          Version,
	Description:      "Deploys the SiloedUSDCTokenPool contract",
	ContractMetadata: siloed_usdc_token_pool.SiloedUSDCTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(siloed_usdc_token_pool.SiloedUSDCTokenPoolBin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})

var ConfigureLockBoxes = contract.NewWrite(contract.WriteParams[[]LockBoxConfig, *siloed_usdc_token_pool.SiloedUSDCTokenPool]{
	Name:            "siloed-usdc-token-pool:configure-lock-boxes",
	Version:         Version,
	Description:     "Configures lock box mappings on the SiloedUSDCTokenPool",
	ContractType:    ContractType,
	ContractABI:     siloed_usdc_token_pool.SiloedUSDCTokenPoolABI,
	NewContract:     siloed_usdc_token_pool.NewSiloedUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*siloed_usdc_token_pool.SiloedUSDCTokenPool, []LockBoxConfig],
	Validate: func(configs []LockBoxConfig) error {
		for _, cfg := range configs {
			if cfg.LockBox == (common.Address{}) {
				return fmt.Errorf("lock box config for chain %d has zero address", cfg.RemoteChainSelector)
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

var GetAllLockBoxConfigs = contract.NewRead(contract.ReadParams[any, []LockBoxConfig, *siloed_usdc_token_pool.SiloedUSDCTokenPool]{
	Name:         "siloed-usdc-token-pool:get-all-lock-box-configs",
	Version:      Version,
	Description:  "Gets all lock box configurations on the SiloedUSDCTokenPool",
	ContractType: ContractType,
	NewContract:  siloed_usdc_token_pool.NewSiloedUSDCTokenPool,
	CallContract: func(pool *siloed_usdc_token_pool.SiloedUSDCTokenPool, opts *bind.CallOpts, args any) ([]LockBoxConfig, error) {
		return pool.GetAllLockBoxConfigs(opts)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[any, []common.Address, *siloed_usdc_token_pool.SiloedUSDCTokenPool]{
	Name:         "siloed-usdc-token-pool:get-all-authorized-callers",
	Version:      Version,
	Description:  "Gets all authorized callers on the SiloedUSDCTokenPool",
	ContractType: ContractType,
	NewContract:  siloed_usdc_token_pool.NewSiloedUSDCTokenPool,
	CallContract: func(pool *siloed_usdc_token_pool.SiloedUSDCTokenPool, opts *bind.CallOpts, args any) ([]common.Address, error) {
		return pool.GetAllAuthorizedCallers(opts)
	},
})
