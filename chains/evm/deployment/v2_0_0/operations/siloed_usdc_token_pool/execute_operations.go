package siloed_usdc_token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// AuthorizedCallerArgs matches applyAuthorizedCallerUpdates input.
type AuthorizedCallerArgs = gobindings.AuthorizedCallersAuthorizedCallerArgs

// LockBoxConfig is a single lock box entry for configureLockBoxes.
type LockBoxConfig = gobindings.SiloedLockReleaseTokenPoolLockBoxConfig

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[struct{}, []common.Address, *gobindings.SiloedUSDCTokenPool]{
	Name:         "siloed-usdc-token-pool:get-all-authorized-callers",
	Version:      Version,
	Description:  "Calls getAllAuthorizedCallers on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedUSDCTokenPool,
	CallContract: func(c *gobindings.SiloedUSDCTokenPool, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
		return c.GetAllAuthorizedCallers(opts)
	},
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[gobindings.AuthorizedCallersAuthorizedCallerArgs, *gobindings.SiloedUSDCTokenPool]{
	Name:            "siloed-usdc-token-pool:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.SiloedUSDCTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewSiloedUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.SiloedUSDCTokenPool, gobindings.AuthorizedCallersAuthorizedCallerArgs],
	Validate:        func(gobindings.AuthorizedCallersAuthorizedCallerArgs) error { return nil },
	CallContract: func(c *gobindings.SiloedUSDCTokenPool, opts *bind.TransactOpts, args gobindings.AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
		return c.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var ConfigureLockBoxes = contract.NewWrite(contract.WriteParams[[]gobindings.SiloedLockReleaseTokenPoolLockBoxConfig, *gobindings.SiloedUSDCTokenPool]{
	Name:            "siloed-usdc-token-pool:configure-lock-boxes",
	Version:         Version,
	Description:     "Calls configureLockBoxes on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.SiloedUSDCTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewSiloedUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.SiloedUSDCTokenPool, []gobindings.SiloedLockReleaseTokenPoolLockBoxConfig],
	Validate:        func([]gobindings.SiloedLockReleaseTokenPoolLockBoxConfig) error { return nil },
	CallContract: func(c *gobindings.SiloedUSDCTokenPool, opts *bind.TransactOpts, args []gobindings.SiloedLockReleaseTokenPoolLockBoxConfig) (*types.Transaction, error) {
		return c.ConfigureLockBoxes(opts, args)
	},
})

var GetAllLockBoxConfigs = contract.NewRead(contract.ReadParams[struct{}, []gobindings.SiloedLockReleaseTokenPoolLockBoxConfig, *gobindings.SiloedUSDCTokenPool]{
	Name:         "siloed-usdc-token-pool:get-all-lock-box-configs",
	Version:      Version,
	Description:  "Calls getAllLockBoxConfigs on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedUSDCTokenPool,
	CallContract: func(c *gobindings.SiloedUSDCTokenPool, opts *bind.CallOpts, args struct{}) ([]gobindings.SiloedLockReleaseTokenPoolLockBoxConfig, error) {
		return c.GetAllLockBoxConfigs(opts)
	},
})
