package erc20

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
)

var ContractType cldf_deployment.ContractType = "ERC20Token"

type ConstructorArgs struct {
	Name   string
	Symbol string
}

type ApproveArgs struct {
	Spender common.Address
	Value   *big.Int
}

type TransferArgs struct {
	Receiver common.Address
	Amount   *big.Int
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "erc20:deploy",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the ERC20 Token contract",
	ContractMetadata: erc20.ERC20MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(erc20.ERC20Bin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})

func NewReadBalanceOf(c *erc20.ERC20) *cld_ops.Operation[contract.FunctionInput[common.Address], *big.Int, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[common.Address, *big.Int, *erc20.ERC20]{
		Name:         "erc20:balance-of",
		Version:      utils.Version_1_0_0,
		Description:  "Gets the ERC20 token balance of a specified address",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(token *erc20.ERC20, opts *bind.CallOpts, account common.Address) (*big.Int, error) {
			return token.BalanceOf(opts, account)
		},
	})
}

func NewWriteTransfer(c *erc20.ERC20) *cld_ops.Operation[contract.FunctionInput[TransferArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[TransferArgs, *erc20.ERC20]{
		Name:            "erc20:transfer",
		Version:         utils.Version_1_0_0,
		Description:     "Transfer ERC20 tokens to a specified address",
		ContractType:    ContractType,
		ContractABI:     erc20.ERC20ABI,
		Contract:        c,
		IsAllowedCaller: contract.AllCallersAllowed[*erc20.ERC20, TransferArgs],
		Validate: func(args TransferArgs) error {
			if args.Amount == nil || args.Amount.Cmp(big.NewInt(0)) <= 0 {
				return errors.New("amount must be greater than 0")
			}
			return nil
		},
		CallContract: func(token *erc20.ERC20, opts *bind.TransactOpts, args TransferArgs) (*types.Transaction, error) {
			return token.Transfer(opts, args.Receiver, args.Amount)
		},
	})
}

func NewReadGetSymbol(c *erc20.ERC20) *cld_ops.Operation[contract.FunctionInput[struct{}], string, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, string, *erc20.ERC20]{
		Name:         "erc20:get-symbol",
		Version:      utils.Version_1_0_0,
		Description:  "Gets the ERC20 token symbol",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(token *erc20.ERC20, opts *bind.CallOpts, _ struct{}) (string, error) {
			return token.Symbol(opts)
		},
	})
}

func NewReadGetDecimals(c *erc20.ERC20) *cld_ops.Operation[contract.FunctionInput[struct{}], uint8, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, uint8, *erc20.ERC20]{
		Name:         "erc20:get-decimals",
		Version:      utils.Version_1_0_0,
		Description:  "Gets the ERC20 token decimals",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(token *erc20.ERC20, opts *bind.CallOpts, _ struct{}) (uint8, error) {
			return token.Decimals(opts)
		},
	})
}

func NewWriteApprove(c *erc20.ERC20) *cld_ops.Operation[contract.FunctionInput[ApproveArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[ApproveArgs, *erc20.ERC20]{
		Name:            "erc20:approve",
		Version:         utils.Version_1_0_0,
		Description:     "Approves a spender for ERC20 transfers",
		ContractType:    ContractType,
		ContractABI:     erc20.ERC20ABI,
		Contract:        c,
		IsAllowedCaller: contract.AllCallersAllowed[*erc20.ERC20, ApproveArgs],
		Validate: func(args ApproveArgs) error {
			if args.Spender == (common.Address{}) {
				return errors.New("spender address must be set")
			}
			if args.Value == nil || args.Value.Sign() <= 0 {
				return errors.New("amount must be greater than zero")
			}
			return nil
		},
		CallContract: func(token *erc20.ERC20, opts *bind.TransactOpts, args ApproveArgs) (*types.Transaction, error) {
			return token.Approve(opts, args.Spender, args.Value)
		},
	})
}

// NewWriteApproveProposalOnly is identical to NewWriteApprove but forces the operation into a
// proposal rather than executing directly. Use when the approve must be called by a timelock
// as part of an atomic MCMS batch (e.g., withdraw → approve → deposit flows).
func NewWriteApproveProposalOnly(c *erc20.ERC20) *cld_ops.Operation[contract.FunctionInput[ApproveArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[ApproveArgs, *erc20.ERC20]{
		Name:            "erc20:approve-proposal-only",
		Version:         utils.Version_1_0_0,
		Description:     "Approves a spender for ERC20 transfers (proposal-only, never executed directly)",
		ContractType:    ContractType,
		ContractABI:     erc20.ERC20ABI,
		Contract:        c,
		IsAllowedCaller: contract.NoCallersAllowed[*erc20.ERC20, ApproveArgs],
		Validate: func(args ApproveArgs) error {
			if args.Spender == (common.Address{}) {
				return errors.New("spender address must be set")
			}
			if args.Value == nil || args.Value.Sign() <= 0 {
				return errors.New("amount must be greater than zero")
			}
			return nil
		},
		CallContract: func(token *erc20.ERC20, opts *bind.TransactOpts, args ApproveArgs) (*types.Transaction, error) {
			return token.Approve(opts, args.Spender, args.Value)
		},
	})
}
