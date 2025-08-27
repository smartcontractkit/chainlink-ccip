package deployment

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
)

// Input is the input structure for the Deploy operation.
type Input[ARGS any] struct {
	// ChainSelector is the selector for the chain on which the contract will be deployed.
	// Required to differentiate between operation runs with the same data targeting different chains.
	ChainSelector uint64 `json:"chainSelector"`
	// Args are the parameters passed to the contract constructor.
	Args ARGS `json:"args"`
}

// AddressRef is the output structure for the Deploy operation.
// This struct serves as a serializable representation of datastore.AddressRef
type AddressRef struct {
	// Address is the address of the contract on the chain.
	Address string `json:"address"`
	// ChainSelector is the chain-selector of the chain where the contract is deployed.
	ChainSelector uint64 `json:"chainSelector"`
	// ContractType is a simple string type for identifying contract types.
	Type deployment.ContractType `json:"type"`
	// Version is the version of the contract.
	Version string `json:"version"`
}

// VMDeployers defines the various deployer functions available for EVM-based chains.
// Currently, it defines an EVM deployer and a ZksyncVM deployer, but can be extended.
type VMDeployers[ARGS any] struct {
	DeployEVM      func(opts *bind.TransactOpts, backend bind.ContractBackend, args ARGS) (common.Address, *types.Transaction, error)
	DeployZksyncVM func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ARGS) (common.Address, error)
}

// New creates a new operation that deploys an EVM contract.
// Any interfacing with gethwrappers should happen in the deploy function.
func New[ARGS any](
	name string,
	version *semver.Version,
	description string,
	contractType deployment.ContractType,
	validate func(input ARGS) error,
	deployers VMDeployers[ARGS],
) *operations.Operation[Input[ARGS], AddressRef, evm.Chain] {
	return operations.NewOperation(
		name,
		version,
		description,
		func(b operations.Bundle, chain evm.Chain, input Input[ARGS]) (AddressRef, error) {
			if validate != nil {
				if err := validate(input.Args); err != nil {
					return AddressRef{}, fmt.Errorf("invalid constructor args for %s: %w", name, err)
				}
			}
			if input.ChainSelector != chain.Selector {
				return AddressRef{}, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			var (
				addr common.Address
				tx   *types.Transaction
				err  error
			)
			if chain.IsZkSyncVM {
				if deployers.DeployZksyncVM == nil {
					return AddressRef{}, fmt.Errorf("no ZkSyncVM deployer defined for %s %s", contractType, version)
				}
				addr, err = deployers.DeployZksyncVM(
					nil,
					chain.ClientZkSyncVM,
					chain.DeployerKeyZkSyncVM,
					chain.Client,
					input.Args,
				)
			} else {
				addr, tx, err = deployers.DeployEVM(
					chain.DeployerKey,
					chain.Client,
					input.Args,
				)
			}
			if err != nil {
				return AddressRef{}, fmt.Errorf("failed to deploy %s %s to %s: %w", contractType, version, chain, err)
			}
			// Non-ZkSyncVM chains require manual confirmation of the deployment transaction.
			if !chain.IsZkSyncVM {
				_, err := chain.Confirm(tx)
				if err != nil {
					return AddressRef{}, fmt.Errorf("failed to confirm deployment of %s %s to %s: %w", contractType, version, chain, err)
				}
			}
			b.Logger.Debugw(fmt.Sprintf("Deployed %s %s to %s", contractType, version, chain), "args", input.Args)

			return AddressRef{
				Address:       addr.Hex(),
				ChainSelector: input.ChainSelector,
				Type:          contractType,
				Version:       version.String(),
			}, err
		},
	)
}
