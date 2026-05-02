package burn_mint_erc677

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "BurnMintToken"

const burnMintERC677ABI = `[{"inputs":[],"name":"BURN_MINT_ADMIN_ROLE","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"DEFAULT_ADMIN_ROLE","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"}],"name":"getRoleAdmin","outputs":[{"internalType":"bytes32","name":"","type":"bytes32"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"grantRole","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes32","name":"role","type":"bytes32"},{"internalType":"address","name":"account","type":"address"}],"name":"hasRole","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"burnAndMinter","type":"address"}],"name":"grantMintAndBurnRoles","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

type AuthorityKind string

const (
	AuthorityBurnMintAdmin AuthorityKind = "burn-mint-admin"
	AuthorityDefaultAdmin  AuthorityKind = "default-admin"
	AuthorityOwner         AuthorityKind = "owner"
	AuthorityUnauthorized  AuthorityKind = "unauthorized"
)

type GrantMintAndBurnRolesAuthority struct {
	Kind              AuthorityKind
	BurnMintAdminRole [32]byte
	AdminRole         [32]byte
	Owner             common.Address
}

func (a GrantMintAndBurnRolesAuthority) CanGrantMintAndBurnRoles() bool {
	return a.Kind == AuthorityBurnMintAdmin || a.Kind == AuthorityOwner
}

type RoleAssignment struct {
	Role [32]byte
	To   common.Address
}

type burnMintERC677 struct {
	address  common.Address
	contract *bind.BoundContract
}

func newBurnMintERC677(address common.Address, backend bind.ContractBackend) (*burnMintERC677, error) {
	parsed, err := abi.JSON(strings.NewReader(burnMintERC677ABI))
	if err != nil {
		return nil, err
	}

	return &burnMintERC677{
		address:  address,
		contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

func (token *burnMintERC677) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := token.contract.Call(opts, &out, "owner")
	if err != nil {
		return common.Address{}, err
	}

	owner := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return owner, nil
}

func (token *burnMintERC677) BurnMintAdminRole(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := token.contract.Call(opts, &out, "BURN_MINT_ADMIN_ROLE")
	if err != nil {
		return [32]byte{}, err
	}

	role := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return role, nil
}

func (token *burnMintERC677) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := token.contract.Call(opts, &out, "getRoleAdmin", role)
	if err != nil {
		return [32]byte{}, err
	}

	adminRole := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return adminRole, nil
}

func (token *burnMintERC677) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := token.contract.Call(opts, &out, "hasRole", role, account)
	if err != nil {
		return false, err
	}

	hasRole := *abi.ConvertType(out[0], new(bool)).(*bool)
	return hasRole, nil
}

func (token *burnMintERC677) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return token.contract.Transact(opts, "grantRole", role, account)
}

func (token *burnMintERC677) GrantMintAndBurnRoles(opts *bind.TransactOpts, burnAndMinter common.Address) (*types.Transaction, error) {
	return token.contract.Transact(opts, "grantMintAndBurnRoles", burnAndMinter)
}

func (token *burnMintERC677) resolveGrantMintAndBurnRolesAuthority(
	opts *bind.CallOpts,
	caller common.Address,
) (GrantMintAndBurnRolesAuthority, error) {
	burnMintAdminRole, accessControlErr := token.BurnMintAdminRole(opts)
	if accessControlErr == nil {
		hasBurnMintAdminRole, err := token.HasRole(opts, burnMintAdminRole, caller)
		if err != nil {
			return GrantMintAndBurnRolesAuthority{}, fmt.Errorf("failed to check burn/mint admin role for %s: %w", caller, err)
		}
		if hasBurnMintAdminRole {
			return GrantMintAndBurnRolesAuthority{
				Kind:              AuthorityBurnMintAdmin,
				BurnMintAdminRole: burnMintAdminRole,
			}, nil
		}

		adminRole, err := token.GetRoleAdmin(opts, burnMintAdminRole)
		if err != nil {
			return GrantMintAndBurnRolesAuthority{}, fmt.Errorf("failed to get burn/mint admin role admin: %w", err)
		}
		hasRoleAdmin, err := token.HasRole(opts, adminRole, caller)
		if err != nil {
			return GrantMintAndBurnRolesAuthority{}, fmt.Errorf("failed to check burn/mint role admin for %s: %w", caller, err)
		}
		if hasRoleAdmin {
			return GrantMintAndBurnRolesAuthority{
				Kind:              AuthorityDefaultAdmin,
				BurnMintAdminRole: burnMintAdminRole,
				AdminRole:         adminRole,
			}, nil
		}

		return GrantMintAndBurnRolesAuthority{
			Kind:              AuthorityUnauthorized,
			BurnMintAdminRole: burnMintAdminRole,
			AdminRole:         adminRole,
		}, nil
	}

	owner, ownerErr := token.Owner(opts)
	if ownerErr == nil {
		if owner == caller {
			return GrantMintAndBurnRolesAuthority{
				Kind:  AuthorityOwner,
				Owner: owner,
			}, nil
		}
		return GrantMintAndBurnRolesAuthority{
			Kind:  AuthorityUnauthorized,
			Owner: owner,
		}, nil
	}

	return GrantMintAndBurnRolesAuthority{}, fmt.Errorf(
		"token does not expose a supported burn/mint role authority interface: BURN_MINT_ADMIN_ROLE failed: %w; owner failed: %w",
		accessControlErr,
		ownerErr,
	)
}

func ResolveGrantMintAndBurnRolesAuthority(
	ctx context.Context,
	backend bind.ContractBackend,
	tokenAddress common.Address,
	caller common.Address,
) (GrantMintAndBurnRolesAuthority, error) {
	if tokenAddress == (common.Address{}) {
		return GrantMintAndBurnRolesAuthority{}, errors.New("token address cannot be zero")
	}
	if caller == (common.Address{}) {
		return GrantMintAndBurnRolesAuthority{}, errors.New("caller address cannot be zero")
	}

	token, err := newBurnMintERC677(tokenAddress, backend)
	if err != nil {
		return GrantMintAndBurnRolesAuthority{}, err
	}
	return token.resolveGrantMintAndBurnRolesAuthority(&bind.CallOpts{Context: ctx}, caller)
}

func PrepareGrantMintAndBurnRoles(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	input contract.FunctionInput[common.Address],
	proposalExecutor common.Address,
) ([]contract.WriteOutput, error) {
	writes := []contract.WriteOutput{}
	deployer := chain.DeployerKey.From
	if proposalExecutor == (common.Address{}) || proposalExecutor == deployer {
		deployerAuthority, err := ResolveGrantMintAndBurnRolesAuthority(b.GetContext(), chain.Client, input.Address, deployer)
		if err == nil && deployerAuthority.Kind == AuthorityDefaultAdmin {
			grantAdminReport, execErr := cldf_ops.ExecuteOperation(b, GrantRole, chain, contract.FunctionInput[RoleAssignment]{
				ChainSelector: input.ChainSelector,
				Address:       input.Address,
				Args: RoleAssignment{
					Role: deployerAuthority.BurnMintAdminRole,
					To:   deployer,
				},
			})
			if execErr != nil {
				return nil, fmt.Errorf("failed to grant burn/mint admin role to deployer %s: %w", deployer, execErr)
			}
			writes = append(writes, grantAdminReport.Output)
		}
	}

	grantReport, err := cldf_ops.ExecuteOperation(b, GrantMintAndBurnRoles, chain, input)
	if err != nil {
		return nil, err
	}
	writes = append(writes, grantReport.Output)
	if grantReport.Output.Executed() || proposalExecutor == (common.Address{}) {
		return writes, nil
	}

	proposalAuthority, err := ResolveGrantMintAndBurnRolesAuthority(b.GetContext(), chain.Client, input.Address, proposalExecutor)
	if err != nil {
		return nil, fmt.Errorf("failed to validate proposal executor %s can grant burn/mint roles: %w", proposalExecutor, err)
	}
	switch proposalAuthority.Kind {
	case AuthorityBurnMintAdmin, AuthorityOwner:
		return writes, nil
	case AuthorityDefaultAdmin:
		grantAdminReport, execErr := cldf_ops.ExecuteOperation(b, GrantRole, chain, contract.FunctionInput[RoleAssignment]{
			ChainSelector: input.ChainSelector,
			Address:       input.Address,
			Args: RoleAssignment{
				Role: proposalAuthority.BurnMintAdminRole,
				To:   proposalExecutor,
			},
		})
		if execErr != nil {
			return nil, fmt.Errorf("failed to grant burn/mint admin role to proposal executor %s: %w", proposalExecutor, execErr)
		}
		return append([]contract.WriteOutput{grantAdminReport.Output}, writes...), nil
	default:
		return nil, fmt.Errorf("proposal executor %s cannot grant burn/mint roles for token %s", proposalExecutor, input.Address)
	}
}

var GrantRole = contract.NewWrite(contract.WriteParams[RoleAssignment, *burnMintERC677]{
	Name:         "burn_mint_erc677:grant-role",
	Version:      cciputils.Version_1_0_0,
	Description:  "Grant role on AccessControl-compatible burn/mint token contract",
	ContractType: ContractType,
	ContractABI:  burnMintERC677ABI,
	NewContract:  newBurnMintERC677,
	IsAllowedCaller: func(token *burnMintERC677, opts *bind.CallOpts, caller common.Address, input RoleAssignment) (bool, error) {
		roleAdmin, err := token.GetRoleAdmin(opts, input.Role)
		if err != nil {
			return false, err
		}
		return token.HasRole(opts, roleAdmin, caller)
	},
	Validate: func(input RoleAssignment) error {
		if input.To == (common.Address{}) {
			return errors.New("role assignee address cannot be zero")
		}
		return nil
	},
	CallContract: func(token *burnMintERC677, opts *bind.TransactOpts, input RoleAssignment) (*types.Transaction, error) {
		return token.GrantRole(opts, input.Role, input.To)
	},
})

var GrantMintAndBurnRoles = contract.NewWrite(contract.WriteParams[common.Address, *burnMintERC677]{
	Name:         "burn_mint_erc677:grant-mint-and-burn-roles",
	Version:      cciputils.Version_1_0_0,
	Description:  "Grant mint and burn role on BurnMintERC677 token contract",
	ContractType: ContractType,
	ContractABI:  burnMintERC677ABI,
	NewContract:  newBurnMintERC677,
	IsAllowedCaller: func(token *burnMintERC677, opts *bind.CallOpts, caller common.Address, input common.Address) (bool, error) {
		authority, err := token.resolveGrantMintAndBurnRolesAuthority(opts, caller)
		if err != nil {
			return false, err
		}
		return authority.CanGrantMintAndBurnRoles(), nil
	},
	Validate: func(address common.Address) error {
		if address == (common.Address{}) {
			return errors.New("burn and minter address cannot be zero")
		}
		return nil
	},
	CallContract: func(token *burnMintERC677, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return token.GrantMintAndBurnRoles(opts, input)
	},
})
