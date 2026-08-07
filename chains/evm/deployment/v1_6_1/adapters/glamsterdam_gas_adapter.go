package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/glamsterdam"
	tar_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// getAllConfiguredTokensInput is the input to getAllConfiguredTokens.
type getAllConfiguredTokensInput struct {
	StartIndex uint64
	MaxCount   uint64
}

// tokenAdminRegistryGetAllConfiguredTokens reads every token TokenAdminRegistry knows about on a chain.
var tokenAdminRegistryGetAllConfiguredTokens = contract.NewRead(contract.ReadParams[getAllConfiguredTokensInput, []common.Address, *tar_bindings.TokenAdminRegistry]{
	Name:         "glamsterdam:token-admin-registry:get-all-configured-tokens",
	Version:      token_admin_registry.Version,
	Description:  "Calls getAllConfiguredTokens on TokenAdminRegistry",
	ContractType: token_admin_registry.ContractType,
	NewContract:  tar_bindings.NewTokenAdminRegistry,
	CallContract: func(c *tar_bindings.TokenAdminRegistry, opts *bind.CallOpts, args getAllConfiguredTokensInput) ([]common.Address, error) {
		return c.GetAllConfiguredTokens(opts, args.StartIndex, args.MaxCount)
	},
})

// GlamsterdamGasAdapter implements v1_6_1_adapters.GasUpdateAdapter for EVM chains.
type GlamsterdamGasAdapter struct{}

// HasLaneToTarget checks if a lane exists from srcChainSelector to targetChainSelector
// by reading the FeeQuoter's dest chain config.
func (a *GlamsterdamGasAdapter) HasLaneToTarget(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
) (bool, error) {
	chain, ok := chains.EVMChains()[srcChainSelector]
	if !ok {
		return false, fmt.Errorf("EVM chain %d not found", srcChainSelector)
	}

	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	fqRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, fee_quoter.ContractType, fee_quoter.Version, "")
	if datastore_utils.IsAddressRefEmpty(fqRef) {
		return false, fmt.Errorf("could not resolve FeeQuoter address on chain %d", srcChainSelector)
	}

	feeQuoterAddr := common.HexToAddress(fqRef.Address)
	result, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: srcChainSelector,
		Address:       feeQuoterAddr,
		Args:          targetChainSelector,
	})
	if err != nil {
		return false, fmt.Errorf("failed to read FeeQuoter dest chain config: %w", err)
	}

	// A lane is enabled if IsEnabled is true
	return result.Output.IsEnabled, nil
}

// ReadDestGasFields reads the current gas field values from FeeQuoter.
func (a *GlamsterdamGasAdapter) ReadDestGasFields(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
) (map[string]uint32, error) {
	chain, ok := chains.EVMChains()[srcChainSelector]
	if !ok {
		return nil, fmt.Errorf("EVM chain %d not found", srcChainSelector)
	}

	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	fqRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, fee_quoter.ContractType, fee_quoter.Version, "")
	if datastore_utils.IsAddressRefEmpty(fqRef) {
		return nil, fmt.Errorf("could not resolve FeeQuoter address on chain %d", srcChainSelector)
	}

	feeQuoterAddr := common.HexToAddress(fqRef.Address)
	result, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: srcChainSelector,
		Address:       feeQuoterAddr,
		Args:          targetChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read FeeQuoter dest chain config: %w", err)
	}

	return map[string]uint32{
		v1_6_1_adapters.FeeQuoterDestGasOverhead.Name:              result.Output.DestGasOverhead,
		v1_6_1_adapters.FeeQuoterDefaultTokenDestGasOverhead.Name:  result.Output.DefaultTokenDestGasOverhead,
	}, nil
}

// WriteDestGasFields writes the resolved gas field values to FeeQuoter via MCMS.
func (a *GlamsterdamGasAdapter) WriteDestGasFields(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
	resolved map[string]uint32,
) ([]mcms_types.BatchOperation, error) {
	chain, ok := chains.EVMChains()[srcChainSelector]
	if !ok {
		return nil, fmt.Errorf("EVM chain %d not found", srcChainSelector)
	}

	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	fqRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, fee_quoter.ContractType, fee_quoter.Version, "")
	if datastore_utils.IsAddressRefEmpty(fqRef) {
		return nil, fmt.Errorf("could not resolve FeeQuoter address on chain %d", srcChainSelector)
	}

	feeQuoterAddr := common.HexToAddress(fqRef.Address)

	// Read current config to build the updated one
	current, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: srcChainSelector,
		Address:       feeQuoterAddr,
		Args:          targetChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read FeeQuoter dest chain config: %w", err)
	}

	// Build updated config
	updated := current.Output
	if val, ok := resolved[v1_6_1_adapters.FeeQuoterDestGasOverhead.Name]; ok {
		updated.DestGasOverhead = val
	}
	if val, ok := resolved[v1_6_1_adapters.FeeQuoterDefaultTokenDestGasOverhead.Name]; ok {
		updated.DefaultTokenDestGasOverhead = val
	}

	// Execute the write operation (exported from sequences)
	writeOut, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyFeeQuoterDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
		ChainSelector: srcChainSelector,
		Address:       feeQuoterAddr,
		Args: []fee_quoter.DestChainConfigArgs{
			{DestChainSelector: targetChainSelector, DestChainConfig: updated},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to apply FeeQuoter update: %w", err)
	}

	batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{writeOut.Output})
	if err != nil {
		return nil, fmt.Errorf("failed to build batch operation: %w", err)
	}

	return []mcms_types.BatchOperation{batchOp}, nil
}

// ReadImmutableSanityFields reads OffRamp's immutable fields for sanity checking.
func (a *GlamsterdamGasAdapter) ReadImmutableSanityFields(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector uint64,
) (map[string]uint32, error) {
	chain, ok := chains.EVMChains()[srcChainSelector]
	if !ok {
		return nil, fmt.Errorf("EVM chain %d not found", srcChainSelector)
	}

	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	offRampRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, offramp.ContractType, offramp.Version, "")
	if datastore_utils.IsAddressRefEmpty(offRampRef) {
		// OffRamp optional for sanity checks
		return map[string]uint32{}, nil
	}

	offRampAddr := common.HexToAddress(offRampRef.Address)
	result, err := cldf_ops.ExecuteOperation(b, offramp.GetStaticConfig, chain, contract.FunctionInput[struct{}]{
		ChainSelector: srcChainSelector,
		Address:       offRampAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read OffRamp static config: %w", err)
	}

	return map[string]uint32{
		"OffRamp.GasForCallExactCheck": uint32(result.Output.GasForCallExactCheck),
	}, nil
}

// DiscoverCandidateTokens returns all tokens known to TokenAdminRegistry on the chain.
func (a *GlamsterdamGasAdapter) DiscoverCandidateTokens(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector uint64,
) ([][]byte, error) {
	chain, ok := chains.EVMChains()[srcChainSelector]
	if !ok {
		return nil, fmt.Errorf("EVM chain %d not found", srcChainSelector)
	}

	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	tarRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, token_admin_registry.ContractType, token_admin_registry.Version, "")
	if datastore_utils.IsAddressRefEmpty(tarRef) {
		return nil, fmt.Errorf("could not resolve TokenAdminRegistry address on chain %d", srcChainSelector)
	}

	// Read all configured tokens
	result, err := cldf_ops.ExecuteOperation(b, tokenAdminRegistryGetAllConfiguredTokens, chain, contract.FunctionInput[getAllConfiguredTokensInput]{
		ChainSelector: srcChainSelector,
		Address:       common.HexToAddress(tarRef.Address),
		Args: getAllConfiguredTokensInput{
			StartIndex: 0,
			MaxCount:   1000,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read configured tokens: %w", err)
	}

	// Convert common.Address to []byte
	var tokens [][]byte
	for _, addr := range result.Output {
		tokens = append(tokens, addr.Bytes())
	}

	return tokens, nil
}

// ReadTokenGasField reads a token's gas field from FeeQuoter.TokenTransferFeeConfig.
func (a *GlamsterdamGasAdapter) ReadTokenGasField(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
	token []byte,
) (uint32, bool, error) {
	chain, ok := chains.EVMChains()[srcChainSelector]
	if !ok {
		return 0, false, fmt.Errorf("EVM chain %d not found", srcChainSelector)
	}

	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	fqRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, fee_quoter.ContractType, fee_quoter.Version, "")
	if datastore_utils.IsAddressRefEmpty(fqRef) {
		return 0, false, fmt.Errorf("could not resolve FeeQuoter address on chain %d", srcChainSelector)
	}

	feeQuoterAddr := common.HexToAddress(fqRef.Address)
	tokenAddr := common.BytesToAddress(token)

	// Read TokenTransferFeeConfig
	result, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetTokenTransferFeeConfig, chain, contract.FunctionInput[fee_quoter.GetTokenTransferFeeConfigArgs]{
		ChainSelector: srcChainSelector,
		Address:       feeQuoterAddr,
		Args: fee_quoter.GetTokenTransferFeeConfigArgs{
			DestChainSelector: targetChainSelector,
			Token:             tokenAddr,
		},
	})
	if err != nil {
		return 0, false, fmt.Errorf("failed to read token transfer fee config: %w", err)
	}

	// Check if configured (IsEnabled flag)
	return result.Output.DestGasOverhead, result.Output.IsEnabled, nil
}

// WriteTokenGasField writes a token's gas field to FeeQuoter via MCMS.
func (a *GlamsterdamGasAdapter) WriteTokenGasField(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
	token []byte,
	value uint32,
) (mcms_types.BatchOperation, error) {
	chain, ok := chains.EVMChains()[srcChainSelector]
	if !ok {
		return mcms_types.BatchOperation{}, fmt.Errorf("EVM chain %d not found", srcChainSelector)
	}

	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	fqRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, fee_quoter.ContractType, fee_quoter.Version, "")
	if datastore_utils.IsAddressRefEmpty(fqRef) {
		return mcms_types.BatchOperation{}, fmt.Errorf("could not resolve FeeQuoter address on chain %d", srcChainSelector)
	}

	feeQuoterAddr := common.HexToAddress(fqRef.Address)
	tokenAddr := common.BytesToAddress(token)

	// Execute the write operation (exported from sequences)
	writeOut, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyFeeQuoterTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs]{
		ChainSelector: srcChainSelector,
		Address:       feeQuoterAddr,
		Args: fee_quoter.ApplyTokenTransferFeeConfigUpdatesArgs{
			TokenTransferFeeConfigArgs: []fee_quoter.TokenTransferFeeConfigArgs{
				{
					DestChainSelector: targetChainSelector,
					TokenTransferFeeConfigs: []fee_quoter.TokenTransferFeeConfigSingleTokenArgs{
						{
							Token: tokenAddr,
							TokenTransferFeeConfig: fee_quoter.TokenTransferFeeConfig{
								DestGasOverhead: value,
							},
						},
					},
				},
			},
		},
	})
	if err != nil {
		return mcms_types.BatchOperation{}, fmt.Errorf("failed to apply token transfer fee config update: %w", err)
	}

	batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{writeOut.Output})
	if err != nil {
		return mcms_types.BatchOperation{}, fmt.Errorf("failed to build batch operation: %w", err)
	}

	return batchOp, nil
}
