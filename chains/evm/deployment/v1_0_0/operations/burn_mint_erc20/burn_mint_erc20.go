package burn_mint_erc20

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

var ContractType cldf_deployment.ContractType = "BurnMintERC20Token"

type ConstructorArgs struct {
	Name      string
	Symbol    string
	Decimals  uint8
	MaxSupply *big.Int
	PreMint   *big.Int
}

type RoleAssignment struct {
	Role [32]byte
	To   common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "burn_mint_erc20:deploy",
	Version:          utils.Version_1_0_0,
	Description:      "Deploys the BurnMintERC20 Token contract",
	ContractMetadata: burn_mint_erc20.BurnMintERC20MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *utils.Version_1_0_0).String(): {
			EVM: common.FromHex(burn_mint_erc20.BurnMintERC20Bin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})

var SetCCIPAdmin = contract.NewWrite(contract.WriteParams[string, *burn_mint_erc20.BurnMintERC20]{
	Name:         "burn_mint_erc20:set-ccip-admin",
	Version:      utils.Version_1_0_0,
	Description:  "Set CCIP Admin on BurnMintERC20 token contract",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20.BurnMintERC20ABI,
	NewContract:  burn_mint_erc20.NewBurnMintERC20,
	IsAllowedCaller: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.CallOpts, caller common.Address, input string) (bool, error) {
		// SetCCIPAdmin can only be called by the current CCIP admin
		currentAdmin, err := token.GetCCIPAdmin(opts)
		if err != nil {
			return false, err
		}
		return currentAdmin == caller, nil
	},
	Validate: func(string) error { return nil },
	CallContract: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.TransactOpts, input string) (*types.Transaction, error) {
		return token.SetCCIPAdmin(opts, common.HexToAddress(input))
	},
})

var GrantAdminRole = contract.NewWrite(contract.WriteParams[RoleAssignment, *burn_mint_erc20.BurnMintERC20]{
	Name:         "burn_mint_erc20:grant-admin-role",
	Version:      utils.Version_1_0_0,
	Description:  "Grant admin role on BurnMintERC20 token contract",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20.BurnMintERC20ABI,
	NewContract:  burn_mint_erc20.NewBurnMintERC20,
	IsAllowedCaller: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.CallOpts, caller common.Address, input RoleAssignment) (bool, error) {
		// Check if caller has the admin role for the role being granted
		roleAdmin, err := token.GetRoleAdmin(opts, input.Role)
		if err != nil {
			return false, err
		}
		return token.HasRole(opts, roleAdmin, caller)
	},
	Validate: func(RoleAssignment) error { return nil },
	CallContract: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.TransactOpts, input RoleAssignment) (*types.Transaction, error) {
		return token.GrantRole(opts, input.Role, input.To)
	},
})

var RevokeAdminRole = contract.NewWrite(contract.WriteParams[RoleAssignment, *burn_mint_erc20.BurnMintERC20]{
	Name:         "burn_mint_erc20:revoke-admin-role",
	Version:      utils.Version_1_0_0,
	Description:  "Revoke admin role on BurnMintERC20 token contract",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20.BurnMintERC20ABI,
	NewContract:  burn_mint_erc20.NewBurnMintERC20,
	IsAllowedCaller: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.CallOpts, caller common.Address, input RoleAssignment) (bool, error) {
		// Check if caller has the admin role for the role being revoked
		roleAdmin, err := token.GetRoleAdmin(opts, input.Role)
		if err != nil {
			return false, err
		}
		return token.HasRole(opts, roleAdmin, caller)
	},
	Validate: func(RoleAssignment) error { return nil },
	CallContract: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.TransactOpts, input RoleAssignment) (*types.Transaction, error) {
		return token.RevokeRole(opts, input.Role, input.To)
	},
})

var RenounceAdminRole = contract.NewWrite(contract.WriteParams[RoleAssignment, *burn_mint_erc20.BurnMintERC20]{
	Name:         "burn_mint_erc20:renounce-admin-role",
	Version:      utils.Version_1_0_0,
	Description:  "Renounce admin role on BurnMintERC20 token contract",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20.BurnMintERC20ABI,
	NewContract:  burn_mint_erc20.NewBurnMintERC20,
	IsAllowedCaller: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.CallOpts, caller common.Address, input RoleAssignment) (bool, error) {
		// For renounce, the caller must be the one renouncing their own role
		// The caller can only renounce for themselves
		return caller == input.To, nil
	},
	Validate: func(RoleAssignment) error { return nil },
	CallContract: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.TransactOpts, input RoleAssignment) (*types.Transaction, error) {
		return token.RenounceRole(opts, input.Role, input.To)
	},
})

var GrantMintAndBurnRoles = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc20.BurnMintERC20]{
	Name:         "burn_mint_erc20:grant-mint-and-burn-roles",
	Version:      utils.Version_1_0_0,
	Description:  "Grant mint and burn role on BurnMintERC20 token contract",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc20.BurnMintERC20ABI,
	NewContract:  burn_mint_erc20.NewBurnMintERC20,
	IsAllowedCaller: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		// Check if caller has the admin role for the role being revoked
		roleAdmin, err := token.DEFAULTADMINROLE(opts)
		if err != nil {
			return false, err
		}
		return token.HasRole(opts, roleAdmin, caller)
	},
	Validate: func(address common.Address) error { return nil },
	CallContract: func(token *burn_mint_erc20.BurnMintERC20, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return token.GrantMintAndBurnRoles(opts, input)
	},
})
