package lock_release_token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

type LockBoxConfig = siloed_lock_release_token_pool.SiloedLockReleaseTokenPoolLockBoxConfig

var GetLockBoxSiloed = contract.NewRead(contract.ReadParams[uint64, common.Address, *siloed_lock_release_token_pool.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-lock-box",
	Version:      Version,
	Description:  "Gets the lock box address for a specific remote chain from a v2.0 SiloedLockReleaseTokenPool",
	ContractType: SiloedContractType,
	NewContract:  siloed_lock_release_token_pool.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *siloed_lock_release_token_pool.SiloedLockReleaseTokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (common.Address, error) {
		return c.GetLockBox(opts, remoteChainSelector)
	},
})

var GetAllLockBoxConfigs = contract.NewRead(contract.ReadParams[any, []LockBoxConfig, *siloed_lock_release_token_pool.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-all-lock-box-configs",
	Version:      Version,
	Description:  "Gets all configured lock box mappings from a v2.0 SiloedLockReleaseTokenPool",
	ContractType: SiloedContractType,
	NewContract:  siloed_lock_release_token_pool.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *siloed_lock_release_token_pool.SiloedLockReleaseTokenPool, opts *bind.CallOpts, _ any) ([]LockBoxConfig, error) {
		return c.GetAllLockBoxConfigs(opts)
	},
})
