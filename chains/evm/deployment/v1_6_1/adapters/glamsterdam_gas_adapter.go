package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/glamsterdam"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
)

// resolveUSDCTokenPoolRef finds the chain's non-canonical USDC token pool address ref, if any.
// v1.6.1 has no USDC-specific ContractType; by this version's convention (see
// adapters/non_canonical_usdc_chain.go), the USDC pool is deployed as a
// BurnMintWithLockReleaseFlagTokenPool, and there is at most one per chain, so the first match by
// type is returned regardless of version or qualifier.
func resolveUSDCTokenPoolRef(addrs []datastore.AddressRef, sel uint64) datastore.AddressRef {
	for _, ref := range addrs {
		if ref.ChainSelector == sel && ref.Type == datastore.ContractType(burn_mint_with_lock_release_flag_token_pool.ContractType) {
			return ref
		}
	}
	return datastore.AddressRef{}
}

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
		v1_6_1_adapters.FeeQuoterDestGasOverhead.Name:             result.Output.DestGasOverhead,
		v1_6_1_adapters.FeeQuoterDefaultTokenDestGasOverhead.Name: result.Output.DefaultTokenDestGasOverhead,
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

// DiscoverCandidateTokens returns the chain's USDC token, resolved via its deployed USDC token
// pool, if any. For v1.6.1, only USDC has a special per-token gas config override (table row 5),
// and there is no USDC-specific ContractType to look up directly — per this version's
// non-canonical-USDC convention, the pool is a BurnMintWithLockReleaseFlagTokenPool, and its
// getToken() is the source of truth for which token is "USDC" on this chain. Chains with no such
// pool deployed have nothing to update for this row and return no candidates.
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
	usdcPoolRef := resolveUSDCTokenPoolRef(addrs, srcChainSelector)
	if datastore_utils.IsAddressRefEmpty(usdcPoolRef) {
		// No non-canonical USDC pool deployed on this chain — nothing to update for the
		// USDC-specific row of the v1.6 mapping table.
		return nil, nil
	}

	usdcTokenReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, contract.FunctionInput[struct{}]{
		ChainSelector: srcChainSelector,
		Address:       common.HexToAddress(usdcPoolRef.Address),
		Args:          struct{}{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read underlying token for USDC pool %s on chain %d: %w", usdcPoolRef.Address, srcChainSelector, err)
	}

	return [][]byte{usdcTokenReport.Output.Bytes()}, nil
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

	// Read the current config first so the write preserves every other field (MinFeeUSDCents,
	// MaxFeeUSDCents, DestBytesOverhead, IsEnabled, ...) — building a fresh zero-value struct here
	// would silently write IsEnabled: false, which FeeQuoter rejects for an "enabled" update path
	// and would otherwise disable the override entirely.
	cur, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetTokenTransferFeeConfig, chain, contract.FunctionInput[fee_quoter.GetTokenTransferFeeConfigArgs]{
		ChainSelector: srcChainSelector,
		Address:       feeQuoterAddr,
		Args: fee_quoter.GetTokenTransferFeeConfigArgs{
			DestChainSelector: targetChainSelector,
			Token:             tokenAddr,
		},
	})
	if err != nil {
		return mcms_types.BatchOperation{}, fmt.Errorf("failed to read token transfer fee config: %w", err)
	}
	newConfig := cur.Output
	newConfig.DestGasOverhead = value

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
							Token:                  tokenAddr,
							TokenTransferFeeConfig: newConfig,
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
