package cctp_through_ccv_token_pool

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// AuthorizedCallerArgs matches applyAuthorizedCallerUpdates input.
type AuthorizedCallerArgs = gobindings.AuthorizedCallersAuthorizedCallerArgs

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[gobindings.AuthorizedCallersAuthorizedCallerArgs, *gobindings.CCTPThroughCCVTokenPool]{
	Name:            "cctp-through-ccv-token-pool:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.CCTPThroughCCVTokenPoolMetaData.ABI,
	NewContract:     gobindings.NewCCTPThroughCCVTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.CCTPThroughCCVTokenPool, gobindings.AuthorizedCallersAuthorizedCallerArgs],
	Validate:        func(gobindings.AuthorizedCallersAuthorizedCallerArgs) error { return nil },
	CallContract: func(c *gobindings.CCTPThroughCCVTokenPool, opts *bind.TransactOpts, args gobindings.AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
		return c.ApplyAuthorizedCallerUpdates(opts, args)
	},
})
