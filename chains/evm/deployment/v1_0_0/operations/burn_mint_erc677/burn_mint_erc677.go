package burn_mint_erc677

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc677"
)

var ContractType cldf_deployment.ContractType = cciputils.BurnMintToken

// AuthorityKind records whether an address may call grantMintAndBurnRoles on BurnMintERC677.
// The chainlink-evm token is Ownable (not AccessControl): only the owner may grant pool roles.
type AuthorityKind string

const (
	AuthorityOwner        AuthorityKind = "owner"
	AuthorityUnauthorized AuthorityKind = "unauthorized"
)

type ConstructorArgs struct {
	Name      string
	Symbol    string
	Decimals  uint8
	MaxSupply *big.Int
}

type GrantMintAndBurnRolesAuthority struct {
	Kind  AuthorityKind
	Owner common.Address
}

func (a GrantMintAndBurnRolesAuthority) CanGrantMintAndBurnRoles() bool {
	return a.Kind == AuthorityOwner
}

// ResolveGrantMintAndBurnRolesAuthority returns whether caller is the Ownable owner of the token.
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

	token, err := burn_mint_erc677.NewBurnMintERC677(tokenAddress, backend)
	if err != nil {
		return GrantMintAndBurnRolesAuthority{}, err
	}
	owner, err := token.Owner(&bind.CallOpts{Context: ctx})
	if err != nil {
		return GrantMintAndBurnRolesAuthority{}, fmt.Errorf("failed to read token owner: %w", err)
	}
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

// PrepareGrantMintAndBurnRoles plans grantMintAndBurnRoles for the pool on a BurnMintERC677 token.
// The on-chain function is owner-gated. IsAllowedCaller on GrantMintAndBurnRoles uses AllCallersAllowed
// because MCMS simulations often use the deployer key while the token owner is the timelock.
// When proposalExecutor is set and differs from the deployer, it must be the token owner.
func PrepareGrantMintAndBurnRoles(
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	input contract.FunctionInput[common.Address],
	proposalExecutor common.Address,
) ([]contract.WriteOutput, error) {
	if proposalExecutor != (common.Address{}) && proposalExecutor != chain.DeployerKey.From {
		auth, err := ResolveGrantMintAndBurnRolesAuthority(b.GetContext(), chain.Client, input.Address, proposalExecutor)
		if err != nil {
			return nil, fmt.Errorf("failed to validate proposal executor %s: %w", proposalExecutor, err)
		}
		if auth.Kind != AuthorityOwner {
			return nil, fmt.Errorf(
				"proposal executor %s is not the token owner (owner=%s) for token %s; cannot grant mint/burn roles",
				proposalExecutor, auth.Owner, input.Address,
			)
		}
	}

	grantReport, err := cldf_ops.ExecuteOperation(b, GrantMintAndBurnRoles, chain, input)
	if err != nil {
		return nil, err
	}

	return []contract.WriteOutput{grantReport.Output}, nil
}

var GrantMintAndBurnRoles = contract.NewWrite(contract.WriteParams[common.Address, *burn_mint_erc677.BurnMintERC677]{
	Name:         "burn_mint_erc677:grant-mint-and-burn-roles",
	Version:      cciputils.Version_1_0_0,
	Description:  "Grant mint and burn roles on BurnMintERC677 (owner-only on-chain)",
	ContractType: ContractType,
	ContractABI:  burn_mint_erc677.BurnMintERC677ABI,
	NewContract:  burn_mint_erc677.NewBurnMintERC677,
	// On-chain only the owner may call grantMintAndBurnRoles. Do not use OnlyOwner here:
	// MCMS/timelock flows simulate with the deployer key while ownership is the timelock
	IsAllowedCaller: contract.AllCallersAllowed[*burn_mint_erc677.BurnMintERC677, common.Address],
	Validate: func(address common.Address) error {
		if address == (common.Address{}) {
			return errors.New("burn and minter address cannot be zero")
		}
		return nil
	},
	CallContract: func(token *burn_mint_erc677.BurnMintERC677, opts *bind.TransactOpts, input common.Address) (*types.Transaction, error) {
		return token.GrantMintAndBurnRoles(opts, input)
	},
})

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "burn_mint_erc677:deploy",
	Version:          cciputils.Version_1_0_0,
	Description:      "Deploys the BurnMintERC677 token contract",
	ContractMetadata: burn_mint_erc677.BurnMintERC677MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *cciputils.Version_1_0_0).String(): {
			EVM: common.FromHex(burn_mint_erc677.BurnMintERC677Bin),
		},
	},
	Validate: func(args ConstructorArgs) error {
		if args.Name == "" {
			return errors.New("name is required")
		}
		if args.Symbol == "" {
			return errors.New("symbol is required")
		}
		if args.MaxSupply == nil {
			return errors.New("maxSupply is required")
		}
		return nil
	},
})
