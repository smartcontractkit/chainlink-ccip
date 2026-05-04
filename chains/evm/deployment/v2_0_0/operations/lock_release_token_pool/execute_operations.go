package lock_release_token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var GetLockBox = contract.NewRead(contract.ReadParams[struct{}, common.Address, *gobindings.LockReleaseTokenPool]{
	Name:         "lock-release-token-pool:get-lock-box",
	Version:      Version,
	Description:  "Calls getLockBox on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewLockReleaseTokenPool,
	CallContract: func(c *gobindings.LockReleaseTokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetLockBox(opts)
	},
})
