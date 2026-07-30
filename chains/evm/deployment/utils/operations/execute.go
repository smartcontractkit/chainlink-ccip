package operations

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// BindAs returns a contract binder that satisfies interface type I from concrete binder newContract.
func BindAs[I, C any](newContract func(common.Address, bind.ContractBackend) (C, error)) func(common.Address, bind.ContractBackend) (I, error) {
	return func(addr common.Address, backend bind.ContractBackend) (I, error) {
		c, err := newContract(addr, backend)
		if err != nil {
			var zero I
			return zero, err
		}
		return any(c).(I), nil
	}
}

// BindContract binds a contract at addr on chain.
func BindContract[C any](
	chain cldf_evm.Chain,
	addr common.Address,
	newContract func(common.Address, bind.ContractBackend) (C, error),
) (C, error) {
	return newContract(addr, chain.Client)
}

func withContractIdempotency[ARGS any](
	chain cldf_evm.Chain,
	addr common.Address,
	opts []cldf_ops.ExecuteOption[contract.FunctionInput[ARGS], cldf_evm.Chain],
) []cldf_ops.ExecuteOption[contract.FunctionInput[ARGS], cldf_evm.Chain] {
	key := ContractIdempotencyKey(chain.Selector, addr)
	return append([]cldf_ops.ExecuteOption[contract.FunctionInput[ARGS], cldf_evm.Chain]{
		cldf_ops.WithIdempotencyKey[contract.FunctionInput[ARGS], cldf_evm.Chain](key),
	}, opts...)
}

// ExecuteRead runs a v2 read factory operation with contract-scoped idempotency.
func ExecuteRead[C, ARGS, OUT any](
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	addr common.Address,
	newContract func(common.Address, bind.ContractBackend) (C, error),
	newOp func(C) *cldf_ops.Operation[contract.FunctionInput[ARGS], OUT, cldf_evm.Chain],
	args ARGS,
	opts ...cldf_ops.ExecuteOption[contract.FunctionInput[ARGS], cldf_evm.Chain],
) (cldf_ops.Report[contract.FunctionInput[ARGS], OUT], error) {
	c, err := BindContract(chain, addr, newContract)
	if err != nil {
		return cldf_ops.Report[contract.FunctionInput[ARGS], OUT]{}, fmt.Errorf("bind contract at %s: %w", addr.Hex(), err)
	}
	return cldf_ops.ExecuteOperation(
		b,
		newOp(c),
		chain,
		contract.FunctionInput[ARGS]{Args: args},
		withContractIdempotency(chain, addr, opts)...,
	)
}

// ExecuteWrite runs a v2 write factory operation with contract-scoped idempotency.
func ExecuteWrite[C, ARGS any](
	b cldf_ops.Bundle,
	chain cldf_evm.Chain,
	addr common.Address,
	newContract func(common.Address, bind.ContractBackend) (C, error),
	newOp func(C) *cldf_ops.Operation[contract.FunctionInput[ARGS], contract.WriteOutput, cldf_evm.Chain],
	args ARGS,
	opts ...cldf_ops.ExecuteOption[contract.FunctionInput[ARGS], cldf_evm.Chain],
) (cldf_ops.Report[contract.FunctionInput[ARGS], contract.WriteOutput], error) {
	c, err := BindContract(chain, addr, newContract)
	if err != nil {
		return cldf_ops.Report[contract.FunctionInput[ARGS], contract.WriteOutput]{}, fmt.Errorf("bind contract at %s: %w", addr.Hex(), err)
	}
	return cldf_ops.ExecuteOperation(
		b,
		newOp(c),
		chain,
		contract.FunctionInput[ARGS]{Args: args},
		withContractIdempotency(chain, addr, opts)...,
	)
}

// ExecuteDeploy runs a deploy operation with chain-scoped idempotency.
func ExecuteDeploy[ARGS any](
	b cldf_ops.Bundle,
	op *cldf_ops.Operation[contract.DeployInput[ARGS], datastore.AddressRef, cldf_evm.Chain],
	chain cldf_evm.Chain,
	input contract.DeployInput[ARGS],
	opts ...cldf_ops.ExecuteOption[contract.DeployInput[ARGS], cldf_evm.Chain],
) (cldf_ops.Report[contract.DeployInput[ARGS], datastore.AddressRef], error) {
	key := ChainIdempotencyKey(chain.Selector)
	allOpts := append([]cldf_ops.ExecuteOption[contract.DeployInput[ARGS], cldf_evm.Chain]{
		cldf_ops.WithIdempotencyKey[contract.DeployInput[ARGS], cldf_evm.Chain](key),
	}, opts...)
	return cldf_ops.ExecuteOperation(b, op, chain, input, allOpts...)
}

// MaybeDeployContract deploys when no matching address ref exists, using chain-scoped idempotency.
func MaybeDeployContract[ARGS any](
	b cldf_ops.Bundle,
	op *cldf_ops.Operation[contract.DeployInput[ARGS], datastore.AddressRef, cldf_evm.Chain],
	chain cldf_evm.Chain,
	input contract.DeployInput[ARGS],
	existingAddresses []datastore.AddressRef,
	opts ...cldf_ops.ExecuteOption[contract.DeployInput[ARGS], cldf_evm.Chain],
) (datastore.AddressRef, error) {
	for _, ref := range existingAddresses {
		if ref.Type == datastore.ContractType(input.TypeAndVersion.Type) &&
			ref.Version != nil && ref.Version.String() == input.TypeAndVersion.Version.String() {
			if input.Qualifier != nil {
				if ref.Qualifier == *input.Qualifier {
					return ref, nil
				}
			} else {
				return ref, nil
			}
		}
	}
	report, err := ExecuteDeploy(b, op, chain, input, opts...)
	if err != nil {
		return datastore.AddressRef{}, err
	}
	return report.Output, nil
}
