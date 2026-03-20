package erc20

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/latest/erc20"
)

// ApproveProposalOnly is identical to Approve but forces the operation into a proposal
// rather than executing directly. Use when the approve must be called by a timelock
// as part of an atomic MCMS batch (e.g., withdraw → approve → deposit flows).
var ApproveProposalOnly = contract.NewWrite(contract.WriteParams[ApproveArgs, *erc20_bindings.ERC20]{
	Name:            "erc20:approve-proposal-only",
	Version:         Version,
	Description:     "Approves a spender for ERC20 transfers (proposal-only, never executed directly)",
	ContractType:    ContractType,
	ContractABI:     erc20_bindings.ERC20ABI,
	NewContract:     erc20_bindings.NewERC20,
	IsAllowedCaller: contract.NoCallersAllowed[*erc20_bindings.ERC20, ApproveArgs],
	Validate:        validateApproveArgs,
	CallContract: func(token *erc20_bindings.ERC20, opts *bind.TransactOpts, args ApproveArgs) (*types.Transaction, error) {
		return token.Approve(opts, args.Spender, args.Amount)
	},
})

func validateApproveArgs(args ApproveArgs) error {
	if args.Spender == (common.Address{}) {
		return fmt.Errorf("spender address must be set")
	}
	if args.Amount == nil || args.Amount.Sign() <= 0 {
		return fmt.Errorf("amount must be greater than zero")
	}
	return nil
}
