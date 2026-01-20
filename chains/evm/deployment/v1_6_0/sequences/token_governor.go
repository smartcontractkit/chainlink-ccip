package sequences

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	tg_bindings "github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/token_governor"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/token_governor"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
)

// Role represents a valid role string for TokenGovernor
type Role string

// Valid role constants for TokenGovernor
const (
	RoleMinter             Role = "minter"
	RoleBridgeMinterBurner Role = "bridge_minter_or_burner"
	RoleBurner             Role = "burner"
	RoleFreezer            Role = "freezer"
	RoleUnfreezer          Role = "unfreezer"
	RolePauser             Role = "pauser"
	RoleUnpauser           Role = "unpauser"
	RoleRecovery           Role = "recovery"
	RoleCheckerAdmin       Role = "checker_admin"
	RoleDefaultAdmin       Role = "default_admin"
)

// IsValid checks if the role is a valid TokenGovernor role
func (r Role) IsValid() bool {
	switch r {
	case RoleMinter, RoleBridgeMinterBurner, RoleBurner, RoleFreezer, RoleUnfreezer,
		RolePauser, RoleUnpauser, RoleRecovery, RoleCheckerAdmin, RoleDefaultAdmin:
		return true
	default:
		return false
	}
}

// DeployTokenGovernorInput is the input for deploying a TokenGovernor contract
type DeployTokenGovernorInput struct {
	Token               string      `yaml:"token" json:"token"`
	InitialDelay        *big.Int    `yaml:"initial-delay" json:"initialDelay"`
	InitialDefaultAdmin string      `yaml:"initial-default-admin" json:"initialDefaultAdmin"`
	MCMS                *mcms.Input `yaml:"mcms,omitempty" json:"mcms,omitempty"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ChainSelector     uint64
	ExistingDataStore datastore.DataStore
}

// TokenGovernorGrantRole represents a role assignment to an account
type TokenGovernorGrantRole struct {
	Role    string         `yaml:"role" json:"role"` // e.g., "minter", "burner", "pauser"
	Account common.Address `yaml:"account" json:"account"`
}

// TokenGovernorRoleInput is the input for role management sequences
type TokenGovernorRoleInput struct {
	Tokens map[uint64]map[string]TokenGovernorGrantRole `yaml:"tokens" json:"tokens"`
	MCMS   *mcms.Input                                  `yaml:"mcms,omitempty" json:"mcms,omitempty"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ExistingDataStore datastore.DataStore
}

// TokenGovernorOwnershipInput is the input for ownership transfer sequences and default admin transfer sequences
// Tokens is a map of chain selector -> token symbol -> owner info/admin info
type TokenGovernorOwnershipInput struct {
	Tokens map[uint64]map[string]common.Address `yaml:"tokens" json:"tokens"`
	MCMS   *mcms.Input                          `yaml:"mcms,omitempty" json:"mcms,omitempty"`
	// below are not specified by the user, filled in by the deployment system to pass to chain operations
	ExistingDataStore datastore.DataStore
}

// GetRoleFromTokenGovernor returns the role bytes32 from the token governor contract.
// Valid role strings: minter, bridge_minter_or_burner, burner, freezer, unfreezer, pauser, unpauser, recovery, checker_admin, default_admin
func GetRoleFromTokenGovernor(ctx context.Context, tokenGovernor *tg_bindings.TokenGovernor, role string) ([32]byte, error) {
	if tokenGovernor == nil {
		return [32]byte{}, errors.New("token governor is nil")
	}

	switch strings.ToLower(role) {
	case "minter":
		r, err := tokenGovernor.MINTERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch minter role: %w", err)
		}
		return r, nil
	case "bridge_minter_or_burner":
		r, err := tokenGovernor.BRIDGEMINTERORBURNERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch bridge minter or burner role: %w", err)
		}
		return r, nil
	case "burner":
		r, err := tokenGovernor.BURNERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch burner role: %w", err)
		}
		return r, nil
	case "freezer":
		r, err := tokenGovernor.FREEZERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch freezer role: %w", err)
		}
		return r, nil
	case "unfreezer":
		r, err := tokenGovernor.UNFREEZERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch unfreezer role: %w", err)
		}
		return r, nil
	case "pauser":
		r, err := tokenGovernor.PAUSERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch pauser role: %w", err)
		}
		return r, nil
	case "unpauser":
		r, err := tokenGovernor.UNPAUSERROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch unpauser role: %w", err)
		}
		return r, nil
	case "recovery":
		r, err := tokenGovernor.RECOVERYROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch recovery role: %w", err)
		}
		return r, nil
	case "checker_admin":
		r, err := tokenGovernor.CHECKERADMINROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch checker admin role: %w", err)
		}
		return r, nil
	case "default_admin":
		r, err := tokenGovernor.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
		if err != nil {
			return [32]byte{}, fmt.Errorf("failed to fetch default admin role: %w", err)
		}
		return r, nil
	default:
		return [32]byte{}, fmt.Errorf("unknown role: %s. Valid roles: minter, bridge_minter_or_burner, burner, freezer, unfreezer, pauser, unpauser, recovery, checker_admin, default_admin", role)
	}
}

var DeployTokenGovernor = cldf_ops.NewSequence(
	"deploy-token-governor",
	token_governor.Version,
	"Deploy token governor contract",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input DeployTokenGovernorInput) (sequences.OnChainOutput, error) {
		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract.WriteOutput, 0)
		chain := chains.EVMChains()[input.ChainSelector]
		var err error
		var tokenGovernorRef datastore.AddressRef
		tokenAddr, err := erc20.NewERC20(common.HexToAddress(input.Token), chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate ERC20 token at %s: %w", input.Token, err)
		}
		symbol, err := tokenAddr.Symbol(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get symbol for ERC20 token at %s: %w", input.Token, err)
		}
		qualifier := symbol
		admin := chain.DeployerKey.From
		if input.InitialDefaultAdmin != "" {
			admin = common.HexToAddress(input.InitialDefaultAdmin)
		}
		tokenGovernorRef, err = contract.MaybeDeployContract(b, token_governor.Deploy, chain, contract.DeployInput[token_governor.ConstructorArgs]{
			TypeAndVersion: token_governor.TypeAndVersion,
			ChainSelector:  chain.Selector,
			Args: token_governor.ConstructorArgs{
				Token:               common.HexToAddress(input.Token),
				InitialDelay:        input.InitialDelay,
				InitialDefaultAdmin: admin,
			},
			Qualifier: &qualifier,
		}, nil)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy TokenGovernor: %w", err)
		}
		addresses = append(addresses, tokenGovernorRef)

		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{
			Addresses: addresses,
			BatchOps:  []mcms_types.BatchOperation{batchOp},
		}, nil
	},
)

var GrantRole = cldf_ops.NewSequence(
	"GrantRole",
	token_governor.Version,
	"grants the given role to the given account on the given chains.",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input TokenGovernorRoleInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)
		for chainSelector, tokenMap := range input.Tokens {
			chain, ok := chains.EVMChains()[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
			}
			for tokenSymbol, tokenGovernorRole := range tokenMap {
				// get token governor address from datastore
				tgAddr, err := GetTokenGovernor(input.ExistingDataStore, chainSelector, tokenSymbol)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token governor address for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}
				// instantiate token governor
				tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token governor at %s: %w", tgAddr.Hex(), err)
				}
				// get role bytes32 from token governor (validates role string internally)
				role, err := GetRoleFromTokenGovernor(b.GetContext(), tg, tokenGovernorRole.Role)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get role from token governor at %s: %w", tgAddr.Hex(), err)
				}
				// check if account already has role
				hasRole, err := tg.HasRole(&bind.CallOpts{Context: b.GetContext()}, role, tokenGovernorRole.Account)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to check if account %s has role on token governor at %s: %w", tokenGovernorRole.Account.Hex(), tgAddr.Hex(), err)
				}
				if hasRole {
					return sequences.OnChainOutput{}, fmt.Errorf("account %s already has role %s", tokenGovernorRole.Account.Hex(), tokenGovernorRole.Role)
				}
				// execute GrantRole operation
				report, err := cldf_ops.ExecuteOperation(b, token_governor.GrantRole, chain, contract.FunctionInput[token_governor.RoleAssignment]{
					ChainSelector: chain.Selector,
					Address:       tgAddr,
					Args: token_governor.RoleAssignment{
						Role: role,
						To:   tokenGovernorRole.Account,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute GrantRole on chain %d for token %s: %w", chainSelector, tokenSymbol, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

var RevokeRole = cldf_ops.NewSequence(
	"RevokeRole",
	token_governor.Version,
	"revokes the given role to the given account on the given chains.",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input TokenGovernorRoleInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)
		for chainSelector, tokenMap := range input.Tokens {
			chain, ok := chains.EVMChains()[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
			}
			for tokenSymbol, tokenGovernorRole := range tokenMap {
				// get token governor address from datastore
				tgAddr, err := GetTokenGovernor(input.ExistingDataStore, chainSelector, tokenSymbol)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token governor address for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}
				// instantiate token governor
				tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token governor at %s: %w", tgAddr.Hex(), err)
				}
				// get role bytes32 from token governor (validates role string internally)
				role, err := GetRoleFromTokenGovernor(b.GetContext(), tg, tokenGovernorRole.Role)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get role from token governor at %s: %w", tgAddr.Hex(), err)
				}
				// check if account already has role
				hasRole, err := tg.HasRole(&bind.CallOpts{Context: b.GetContext()}, role, tokenGovernorRole.Account)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to check if account %s has role on token governor at %s: %w", tokenGovernorRole.Account.Hex(), tgAddr.Hex(), err)
				}
				if !hasRole {
					return sequences.OnChainOutput{}, fmt.Errorf("account %s doesn't have role %s", tokenGovernorRole.Account.Hex(), tokenGovernorRole.Role)
				}
				// execute RevokeRole operation
				report, err := cldf_ops.ExecuteOperation(b, token_governor.RevokeRole, chain, contract.FunctionInput[token_governor.RoleAssignment]{
					ChainSelector: chain.Selector,
					Address:       tgAddr,
					Args: token_governor.RoleAssignment{
						Role: role,
						To:   tokenGovernorRole.Account,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute RevokeRole on chain %d for token %s: %w", chainSelector, tokenSymbol, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

var RenounceRole = cldf_ops.NewSequence(
	"RenounceRole",
	token_governor.Version,
	"renounce the given role to the given account on the given chains.",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input TokenGovernorRoleInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)
		for chainSelector, tokenMap := range input.Tokens {
			chain, ok := chains.EVMChains()[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
			}
			for tokenSymbol, tokenGovernorRole := range tokenMap {
				// get token governor address from datastore
				tgAddr, err := GetTokenGovernor(input.ExistingDataStore, chainSelector, tokenSymbol)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token governor address for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}
				// instantiate token governor
				tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token governor at %s: %w", tgAddr.Hex(), err)
				}
				// get role bytes32 from token governor (validates role string internally)
				role, err := GetRoleFromTokenGovernor(b.GetContext(), tg, tokenGovernorRole.Role)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get role from token governor at %s: %w", tgAddr.Hex(), err)
				}
				// check if account already has role
				hasRole, err := tg.HasRole(&bind.CallOpts{Context: b.GetContext()}, role, tokenGovernorRole.Account)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to check if account %s has role on token governor at %s: %w", tokenGovernorRole.Account.Hex(), tgAddr.Hex(), err)
				}
				if !hasRole {
					return sequences.OnChainOutput{}, fmt.Errorf("account %s doesn't have role %s", tokenGovernorRole.Account.Hex(), tokenGovernorRole.Role)
				}
				// execute RenounceRole operation
				report, err := cldf_ops.ExecuteOperation(b, token_governor.RenounceRole, chain, contract.FunctionInput[token_governor.RoleAssignment]{
					ChainSelector: chain.Selector,
					Address:       tgAddr,
					Args: token_governor.RoleAssignment{
						Role: role,
						To:   tokenGovernorRole.Account,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute RenounceRole on chain %d for token %s: %w", chainSelector, tokenSymbol, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

func GetTokenGovernor(ds datastore.DataStore, chainSelector uint64, tokenSymbol string) (common.Address, error) {
	// fetch token governor from the data store
	tokenGovernorAddr, err := datastore_utils.FindAndFormatRef(ds, datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(token_governor.ContractType),
		Qualifier:     tokenSymbol,
	}, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("token governor for token with symbol '%s' is not found in datastore, %v", tokenSymbol, err)
	}
	return common.HexToAddress(tokenGovernorAddr.Address), nil
}

var TransferOwnership = cldf_ops.NewSequence(
	"TransferOwnership",
	token_governor.Version,
	"transfers ownership of TokenGovernor to a new owner (requires acceptance by new owner).",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input TokenGovernorOwnershipInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)

		for chainSelector, tokenMap := range input.Tokens {
			chain, ok := chains.EVMChains()[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
			}

			for tokenSymbol, newOwner := range tokenMap {
				// get token governor address from datastore
				tgAddr, err := GetTokenGovernor(input.ExistingDataStore, chainSelector, tokenSymbol)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token governor address for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}

				// instantiate token governor
				tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token governor at %s: %w", tgAddr.Hex(), err)
				}

				// verify current owner
				currentOwner, err := tg.Owner(&bind.CallOpts{Context: b.GetContext()})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get current owner: %w", err)
				}
				if currentOwner == newOwner {
					return sequences.OnChainOutput{}, fmt.Errorf("new owner %s is already the current owner for token %s on chain %d", newOwner.Hex(), tokenSymbol, chainSelector)
				}

				// execute TransferOwnership operation
				report, err := cldf_ops.ExecuteOperation(b, token_governor.TransferOwnership, chain, contract.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       tgAddr,
					Args:          newOwner,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute TransferOwnership for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

var AcceptOwnership = cldf_ops.NewSequence(
	"AcceptOwnership",
	token_governor.Version,
	"accepts ownership of TokenGovernor",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input TokenGovernorOwnershipInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)

		for chainSelector, tokenMap := range input.Tokens {
			chain, ok := chains.EVMChains()[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
			}

			for tokenSymbol, newOwner := range tokenMap {
				// get token governor address from datastore
				tgAddr, err := GetTokenGovernor(input.ExistingDataStore, chainSelector, tokenSymbol)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token governor address for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}

				// execute AcceptOwnership operation
				report, err := cldf_ops.ExecuteOperation(b, token_governor.AcceptOwnership, chain, contract.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       tgAddr,
					Args:          newOwner,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute AcceptOwnership for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

var BeginDefaultAdminTransfer = cldf_ops.NewSequence(
	"BeginDefaultAdminTransfer",
	token_governor.Version,
	"begins the transfer of default admin role to a new admin (requires acceptance by new admin after delay).",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input TokenGovernorOwnershipInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)

		for chainSelector, tokenMap := range input.Tokens {
			chain, ok := chains.EVMChains()[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
			}

			for tokenSymbol, newAdmin := range tokenMap {
				// get token governor address from datastore
				tgAddr, err := GetTokenGovernor(input.ExistingDataStore, chainSelector, tokenSymbol)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token governor address for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}

				// instantiate token governor
				tg, err := tg_bindings.NewTokenGovernor(tgAddr, chain.Client)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token governor at %s: %w", tgAddr.Hex(), err)
				}

				// verify current default admin
				currentAdmin, err := tg.DefaultAdmin(&bind.CallOpts{Context: b.GetContext()})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get current default admin: %w", err)
				}
				if currentAdmin == newAdmin {
					return sequences.OnChainOutput{}, fmt.Errorf("new admin %s is already the current default admin for token %s on chain %d", newAdmin.Hex(), tokenSymbol, chainSelector)
				}

				// execute BeginDefaultAdminTransfer operation
				report, err := cldf_ops.ExecuteOperation(b, token_governor.BeginDefaultAdminTransfer, chain, contract.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       tgAddr,
					Args:          newAdmin,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute BeginDefaultAdminTransfer for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})

var AcceptDefaultAdminTransfer = cldf_ops.NewSequence(
	"AcceptDefaultAdminTransfer",
	token_governor.Version,
	"accepts the pending default admin role transfer (must be called by the pending admin after delay).",
	func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input TokenGovernorOwnershipInput) (sequences.OnChainOutput, error) {
		writes := make([]contract.WriteOutput, 0)

		for chainSelector, tokenMap := range input.Tokens {
			chain, ok := chains.EVMChains()[chainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", chainSelector)
			}

			for tokenSymbol, admin := range tokenMap {
				// get token governor address from datastore
				tgAddr, err := GetTokenGovernor(input.ExistingDataStore, chainSelector, tokenSymbol)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get token governor address for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}

				// execute AcceptDefaultAdminTransfer operation
				report, err := cldf_ops.ExecuteOperation(b, token_governor.AcceptDefaultAdminTransfer, chain, contract.FunctionInput[common.Address]{
					ChainSelector: chain.Selector,
					Address:       tgAddr,
					Args:          admin,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to execute AcceptDefaultAdminTransfer for token %s on chain %d: %w", tokenSymbol, chainSelector, err)
				}
				writes = append(writes, report.Output)
			}
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})
