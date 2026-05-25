package erc20

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
)

var ContractType cldf_deployment.ContractType = "ERC20Token"

type ConstructorArgs struct {
	Name   string
	Symbol string
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

var Transfer = contract.NewWrite(contract.WriteParams[TransferArgs, *erc20.ERC20]{
	Name:            "erc20:transfer",
	Version:         utils.Version_1_0_0,
	Description:     "Transfer ERC20 tokens to a specified address",
	ContractType:    ContractType,
	ContractABI:     erc20.ERC20ABI,
	NewContract:     erc20.NewERC20,
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

var GetSymbol = contract.NewRead(contract.ReadParams[struct{}, string, *erc20.ERC20]{
	Name:         "erc20:get-symbol",
	Version:      utils.Version_1_0_0,
	Description:  "Gets the ERC20 token symbol",
	ContractType: ContractType,
	NewContract:  erc20.NewERC20,
	CallContract: func(token *erc20.ERC20, opts *bind.CallOpts, _ struct{}) (string, error) {
		return token.Symbol(opts)
	},
})

var GetDecimals = contract.NewRead(contract.ReadParams[struct{}, uint8, *erc20.ERC20]{
	Name:         "erc20:get-decimals",
	Version:      utils.Version_1_0_0,
	Description:  "Gets the ERC20 token decimals",
	ContractType: ContractType,
	NewContract:  erc20.NewERC20,
	CallContract: func(token *erc20.ERC20, opts *bind.CallOpts, _ struct{}) (uint8, error) {
		return token.Decimals(opts)
	},
})
