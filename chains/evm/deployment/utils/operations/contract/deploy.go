package contract

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"reflect"

	"github.com/Masterminds/semver/v3"
	"github.com/aws/smithy-go/ptr"
	"github.com/ethereum/go-ethereum/accounts/abi"
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

var (
	// For testing, these functions can be overridden to mock deployments.
	// In production, they point to the real deployment functions.
	// See deployment_test.go for usage in tests.
	deployZkContract  = deployZkContractImpl
	deployEVMContract = bind.DeployContract
)

// DeployInput is the input structure for the Deploy operation.
type DeployInput[ARGS any] struct {
	// ChainSelector is the selector for the chain on which the contract will be deployed.
	// Required to differentiate between operation runs with the same data targeting different chains.
	ChainSelector uint64 `json:"chainSelector"`
	// TypeAndVersion is the desired type and version of the contract to deploy.
	// The deployment operation must define bytecode for this type and version.
	TypeAndVersion deployment.TypeAndVersion `json:"typeAndVersion"`
	// Qualifier is an optional string to differentiate between multiple deployments of
	//	the same contract type and version on the same chain.
	// if provided, it is stored in the AddressRef returned by the operation.
	Qualifier *string `json:"qualifier,omitempty"`
	// Args are the parameters passed to the contract constructor.
	Args ARGS `json:"args"`
}

// Bytecode specifies the exact bytecode to deploy for each supported VM.
// Specifying multiple possible bytecodes allows callers to inject different bytecodes into the same deployment operation.
// So long as the constructor arguments and ABI remain the same, the bytecode can be swapped out.
type Bytecode struct {
	EVM      []byte
	ZkSyncVM []byte
}

func (b Bytecode) Validate(isZkSyncVM bool) error {
	if isZkSyncVM && len(b.ZkSyncVM) == 0 {
		return errors.New("zkSyncVM bytecode must be defined")
	}
	if !isZkSyncVM && len(b.EVM) == 0 {
		return errors.New("evm bytecode must be defined")
	}
	return nil
}

// DeployParams encapsulates all parameters required to create a deploy operation for an EVM contract.
type DeployParams[ARGS any] struct {
	// Name is the name of the operation.
	Name string
	// Version is the version of the operation.
	Version *semver.Version
	// Description is a brief description of the operation.
	Description string
	// ContractMetadata is the metadata from which the ABI is parsed.
	ContractMetadata *bind.MetaData
	// BytecodeByTypeAndVersion is a map of bytecodes for different types and versions of the contract.
	// The key is the string representation of the type and version.
	BytecodeByTypeAndVersion map[string]Bytecode
	// Validate is an optional function to validate the constructor arguments before deployment.
	Validate func(input ARGS) error
}

// NewDeploy creates a new operation that deploys an EVM contract.
// Any interfacing with gethwrappers should happen in the deploy function.
func NewDeploy[ARGS any](params DeployParams[ARGS]) *operations.Operation[DeployInput[ARGS], datastore.AddressRef, evm.Chain] {
	return operations.NewOperation(
		params.Name,
		params.Version,
		params.Description,
		func(b operations.Bundle, chain evm.Chain, input DeployInput[ARGS]) (datastore.AddressRef, error) {
			// BEGIN Validation
			if params.Validate != nil {
				if err := params.Validate(input.Args); err != nil {
					return datastore.AddressRef{}, fmt.Errorf("invalid constructor args for %s: %w", params.Name, err)
				}
			}
			if input.ChainSelector != chain.Selector {
				return datastore.AddressRef{}, fmt.Errorf("mismatch between inputted chain selector and selector defined within dependencies: %d != %d", input.ChainSelector, chain.Selector)
			}
			if params.ContractMetadata == nil {
				return datastore.AddressRef{}, fmt.Errorf("contract metadata must be defined for %s", params.Name)
			}
			bytecode, ok := params.BytecodeByTypeAndVersion[input.TypeAndVersion.String()]
			if !ok {
				return datastore.AddressRef{}, fmt.Errorf("no bytecode defined for %s", input.TypeAndVersion)
			}
			// END Validation

			parsedABI, err := params.ContractMetadata.GetAbi()
			if err != nil {
				return datastore.AddressRef{}, fmt.Errorf("failed to parse ABI for %s: %w", input.TypeAndVersion, err)
			}
			if parsedABI == nil {
				return datastore.AddressRef{}, fmt.Errorf("abi is nil for %s", input.TypeAndVersion)
			}
			args, err := arrayify(input.Args)
			if err != nil {
				return datastore.AddressRef{}, fmt.Errorf("failed to arrayify constructor args for %s: %w", input.TypeAndVersion, err)
			}

			var (
				addr      common.Address
				tx        *types.Transaction
				deployErr error
			)
			if chain.IsZkSyncVM {
				addr, deployErr = deployZkContract(
					nil,
					bytecode.ZkSyncVM,
					chain.ClientZkSyncVM,
					chain.DeployerKeyZkSyncVM,
					parsedABI,
					args...,
				)
			} else {
				addr, tx, _, deployErr = deployEVMContract(
					chain.DeployerKey,
					*parsedABI,
					bytecode.EVM,
					chain.Client,
					args...,
				)
			}
			if !chain.IsZkSyncVM {
				// Non-ZkSyncVM chains require manual confirmation of the deployment transaction.
				// We attempt to decode any errors with the provided ABI.
				_, confirmErr := deployment.ConfirmIfNoErrorWithABI(chain, tx, params.ContractMetadata.ABI, deployErr)
				if confirmErr != nil {
					return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s to %s with args %+v: %w", input.TypeAndVersion, chain, input.Args, confirmErr)
				}
				b.Logger.Debugw(fmt.Sprintf("Confirmed %s tx on %s", params.Name, chain), "hash", tx.Hash().Hex())
			} else if deployErr != nil {
				// For ZkSyncVM chains, if an error is returned from initial deployment, we return it here.
				return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s to %s with args %+v: %w", input.TypeAndVersion, chain, input.Args, deployErr)
			}
			b.Logger.Debugw(fmt.Sprintf("Deployed %s to %s", input.TypeAndVersion, chain), "args", input.Args)
			return datastore.AddressRef{
				Address:       addr.Hex(),
				ChainSelector: input.ChainSelector,
				Type:          datastore.ContractType(input.TypeAndVersion.Type),
				Version:       &input.TypeAndVersion.Version,
				Qualifier:     ptr.ToString(input.Qualifier),
			}, nil
		},
	)
}

// deployZkContractImpl deploys a contract on a ZkSync VM chain using the provided parameters.
func deployZkContractImpl(
	deployOpts *accounts.TransactOpts,
	bytecode []byte,
	client *clients.Client,
	wallet *accounts.Wallet,
	parsedABI *abi.ABI,
	args ...interface{},
) (common.Address, error) {
	var calldata []byte
	var err error
	if len(args) > 0 {
		calldata, err = parsedABI.Pack("", args...)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to pack constructor args: %w", err)
		}
	}

	salt := make([]byte, 32)
	n, err := rand.Read(salt)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to read random bytes: %w", err)
	}
	if n != len(salt) {
		return common.Address{}, fmt.Errorf("failed to read random bytes: expected %d, got %d", len(salt), n)
	}

	txHash, err := wallet.Deploy(deployOpts, accounts.Create2Transaction{
		Bytecode: bytecode,
		Calldata: calldata,
		Salt:     salt,
	})
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy zk contract: %w", err)
	}

	receipt, err := client.WaitMined(context.Background(), txHash)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to confirm zk contract deployment: %w", err)
	}

	return receipt.ContractAddress, nil
}

// arrayify converts a struct or pointer to struct into a slice of its field values.
func arrayify[ARGS any](args ARGS) ([]interface{}, error) {
	v := reflect.ValueOf(args)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected struct or pointer to struct, got %s", v.Kind())
	}

	result := make([]interface{}, v.NumField())
	for i := 0; i < v.NumField(); i++ {
		result[i] = v.Field(i).Interface()
	}
	return result, nil
}

func MaybeDeployContract[ARGS any](
	b operations.Bundle,
	op *operations.Operation[DeployInput[ARGS], datastore.AddressRef, evm.Chain],
	chain evm.Chain,
	input DeployInput[ARGS],
	existingAddresses []datastore.AddressRef) (datastore.AddressRef, error) {
	for _, ref := range existingAddresses {
		if ref.Type == datastore.ContractType(input.TypeAndVersion.Type) &&
			ref.Version.String() == input.TypeAndVersion.Version.String() {
			if input.Qualifier != nil {
				if ref.Qualifier == *input.Qualifier {
					return ref, nil
				}
			} else {
				return ref, nil
			}
		}
	}
	report, err := operations.ExecuteOperation(b, op, chain, input)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to deploy %s %s: %w", input.TypeAndVersion.Type, op.Def().Version, err)
	}
	return report.Output, nil
}
