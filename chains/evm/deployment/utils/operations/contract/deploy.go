package contract

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
)

// DeployInput is the input structure for the Deploy operation.
type DeployInput[ARGS any] struct {
	// ChainSelector is the selector for the chain on which the contract will be deployed.
	// Required to differentiate between operation runs with the same data targeting different chains.
	ChainSelector uint64 `json:"chainSelector"`
	// Args are the parameters passed to the contract constructor.
	Args ARGS `json:"args"`
}

// VMDeployers defines the various deployer functions available for EVM-based chains.
// Currently, it defines an EVM deployer and a ZksyncVM deployer, but can be extended.
type VMDeployers[ARGS any] struct {
	DeployEVM      func(opts *bind.TransactOpts, backend bind.ContractBackend, args ARGS) (common.Address, *types.Transaction, error)
	DeployZksyncVM func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ARGS) (common.Address, error)
}

// NewDeploy creates a new operation that deploys an EVM contract.
// Any interfacing with gethwrappers should happen in the deploy function.
func NewDeploy[ARGS any](
	name string,
	version *semver.Version,
	description string,
	contractType deployment.ContractType,
	contractABI string,
	validate func(input ARGS) error,
	deployers VMDeployers[ARGS],
) *operations.Operation[DeployInput[ARGS], datastore.AddressRef, evm.Chain] {
	return operations.NewOperation(
		name,
		version,
		description,
		func(b operations.Bundle, chain evm.Chain, input DeployInput[ARGS]) (datastore.AddressRef, error) {
			if validate != nil {
				if err := validate(input.Args); err != nil {
					return datastore.AddressRef{}, fmt.Errorf("invalid constructor args for %s: %w", name, err)
				}
			}
			if input.ChainSelector != chain.Selector {
				return datastore.AddressRef{}, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			var (
				addr      common.Address
				tx        *types.Transaction
				deployErr error
			)
			if chain.IsZkSyncVM {
				if deployers.DeployZksyncVM == nil {
					return datastore.AddressRef{}, fmt.Errorf("no ZkSyncVM deployer defined for %s %s", contractType, version)
				}
				addr, deployErr = deployers.DeployZksyncVM(
					nil,
					chain.ClientZkSyncVM,
					chain.DeployerKeyZkSyncVM,
					chain.Client,
					input.Args,
				)
			} else {
				addr, tx, deployErr = deployers.DeployEVM(
					chain.DeployerKey,
					chain.Client,
					input.Args,
				)
			}
			if !chain.IsZkSyncVM {
				// Non-ZkSyncVM chains require manual confirmation of the deployment transaction.
				// We attempt to decode any errors with the provided ABI.
				_, confirmErr := deployment.ConfirmIfNoErrorWithABI(chain, tx, contractABI, deployErr)
				if confirmErr != nil {
					return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s %s to %s with args %+v: %w", contractType, version, chain, input.Args, confirmErr)
				}
				b.Logger.Debugw(fmt.Sprintf("Confirmed %s tx on %s", name, chain), "hash", tx.Hash().Hex())
			} else if deployErr != nil {
				// For ZkSyncVM chains, if an error is returned from initial deployment, we return it here.
				return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s %s to %s with args %+v: %w", contractType, version, chain, input.Args, deployErr)
			}
			b.Logger.Debugw(fmt.Sprintf("Deployed %s %s to %s", contractType, version, chain), "args", input.Args)

			return datastore.AddressRef{
				Address:       addr.Hex(),
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(contractType),
				Version:       version,
			}, nil
		},
	)
}
