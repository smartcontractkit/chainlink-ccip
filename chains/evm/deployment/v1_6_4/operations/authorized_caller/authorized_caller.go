package authorized_caller

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_4/usdc_token_pool"
)

var Version *semver.Version = semver.MustParse("1.6.4")

type AuthorizedCallerUpdateArgs = usdc_token_pool.AuthorizedCallersAuthorizedCallerArgs

// Note: The AuthorizedCallers Contract does not have its own gobinding as it should never be used on its own, and
// is only present as a parent to other contracts included in v1_6_4. It is being included here in its own
// file to allow for the use of the ApplyAuthorizedCallerUpdates operation on a variety of different contracts
// that implement the AuthorizedCallers interface (e.g. USDCTokenPool, SiloedUSDCTokenPool, etc.).
var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerUpdateArgs, *usdc_token_pool.USDCTokenPool]{
	Name:            "authorized-caller:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Applies authorized caller updates to a contract implementing the AuthorizedCallers interface",
	ContractType:    "AuthorizedCallers",
	ContractABI:     usdc_token_pool.USDCTokenPoolABI,
	NewContract:     usdc_token_pool.NewUSDCTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*usdc_token_pool.USDCTokenPool, AuthorizedCallerUpdateArgs],
	Validate:        func(AuthorizedCallerUpdateArgs) error { return nil },
	CallContract: func(usdcTokenPool *usdc_token_pool.USDCTokenPool, opts *bind.TransactOpts, args AuthorizedCallerUpdateArgs) (*types.Transaction, error) {
		return usdcTokenPool.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
