package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ManualRegistrationSequenceInput struct {
	Address       common.Address
	ChainSelector uint64
	AdminAddress  common.Address
	TokenAddress  common.Address
}

var ManualRegistrationSequence = operations.NewSequence(
	"token-admin-registry:manual-registration",
	semver.MustParse("1.5.0"),
	"Transfer or propose admin role for a token in the TokenAdminRegistry contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input ManualRegistrationSequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}

		tar, err := token_admin_registry.NewTokenAdminRegistry(input.Address, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to bind to token admin registry at address %q on chain %d: %w", input.Address, input.ChainSelector, err)
		}

		cfg, err := tar.GetTokenConfig(&bind.CallOpts{Context: b.GetContext()}, input.TokenAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config for token %q from token admin registry at address %q: %w", input.TokenAddress, input.Address, err)
		}

		writes := make([]contract.WriteOutput, 0)
		if cfg.Administrator == (common.Address{}) {
			report, err := operations.ExecuteOperation(b,
				tarops.ProposeAdministrator,
				chain,
				contract.FunctionInput[tarops.ProposeAdministratorArgs]{
					Address:       input.Address,
					ChainSelector: input.ChainSelector,
					Args: tarops.ProposeAdministratorArgs{
						Administrator: input.AdminAddress,
						TokenAddress:  input.TokenAddress,
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute ProposeAdministrator operation on %q: %w", chain, err)
			}
			writes = append(writes, report.Output)
		} else {
			report, err := operations.ExecuteOperation(b,
				tarops.TransferAdminRole,
				chain,
				contract.FunctionInput[tarops.TransferAdminRoleArgs]{
					Address:       input.Address,
					ChainSelector: input.ChainSelector,
					Args: tarops.TransferAdminRoleArgs{
						Administrator: input.AdminAddress,
						TokenAddress:  input.TokenAddress,
					},
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to execute TransferAdminRole operation on %q: %w", chain, err)
			}
			writes = append(writes, report.Output)
		}

		batch, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}
		return sequences.OnChainOutput{
			BatchOps: []mcms_types.BatchOperation{batch},
		}, nil
	})
