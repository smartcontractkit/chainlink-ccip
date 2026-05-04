package siloed_lock_release_token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var GetAllLockBoxConfigs = contract.NewRead(contract.ReadParams[struct{}, []gobindings.SiloedLockReleaseTokenPoolLockBoxConfig, *gobindings.SiloedLockReleaseTokenPool]{
	Name:         "siloed-lock-release-token-pool:get-all-lock-box-configs",
	Version:      Version,
	Description:  "Calls getAllLockBoxConfigs on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewSiloedLockReleaseTokenPool,
	CallContract: func(c *gobindings.SiloedLockReleaseTokenPool, opts *bind.CallOpts, args struct{}) ([]gobindings.SiloedLockReleaseTokenPoolLockBoxConfig, error) {
		return c.GetAllLockBoxConfigs(opts)
	},
})
