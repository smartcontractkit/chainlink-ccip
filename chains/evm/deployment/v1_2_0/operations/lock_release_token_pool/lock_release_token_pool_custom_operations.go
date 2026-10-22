package lock_release_token_pool

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

// ReleaseOrMintArgs carries the arguments of the flat v1.2.0
// releaseOrMint(bytes,address,uint256,uint64,bytes) that the lock-release pool actually reads. The
// original sender, remote chain selector and extra data are ignored by v1.2.0 and sent empty.
type ReleaseOrMintArgs struct {
	Receiver common.Address
	Amount   *big.Int
}

// ReleaseOrMintProposalOnly forces the offRamp-gated release into an MCMS proposal. The on-chain
// gate is the offRamp set, which none of operations-gen's access aliases describe, so it is never
// executed directly with the deployer key.
var ReleaseOrMintProposalOnly = contract.NewWrite(contract.WriteParams[ReleaseOrMintArgs, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:            "lock_release_token_pool:release-or-mint-proposal-only",
	Version:         Version,
	Description:     "Calls releaseOrMint on the LockReleaseTokenPool 1.2.0 contract (proposal-only, never executed directly)",
	ContractType:    ContractType,
	ContractABI:     lock_release_token_pool.LockReleaseTokenPoolABI,
	NewContract:     lock_release_token_pool.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.NoCallersAllowed[*lock_release_token_pool.LockReleaseTokenPool, ReleaseOrMintArgs],
	Validate: func(args ReleaseOrMintArgs) error {
		if args.Receiver == (common.Address{}) {
			return errors.New("receiver must not be the zero address")
		}
		if args.Amount == nil || args.Amount.Sign() <= 0 {
			return errors.New("amount must be positive")
		}
		return nil
	},
	CallContract: func(c *lock_release_token_pool.LockReleaseTokenPool, opts *bind.TransactOpts, args ReleaseOrMintArgs) (*types.Transaction, error) {
		return c.ReleaseOrMint(opts, nil, args.Receiver, args.Amount, 0, nil)
	},
})
