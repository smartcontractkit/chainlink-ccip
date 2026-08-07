package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	v2_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/lombard_verifier"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// GlamsterdamGasAdapter implements v2_0_0_adapters.GasUpdateAdapter for EVM chains.
type GlamsterdamGasAdapter struct{}

// HasLaneToTarget checks if a lane exists by checking OnRamp's Router field.
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
	onRampRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, onramp.ContractType, onramp.Version, "")
	if datastore_utils.IsAddressRefEmpty(onRampRef) {
		return false, fmt.Errorf("could not resolve OnRamp address on chain %d", srcChainSelector)
	}

	onRampAddr := common.HexToAddress(onRampRef.Address)
	result, err := cldf_ops.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: srcChainSelector,
		Address:       onRampAddr,
		Args:          targetChainSelector,
	})
	if err != nil {
		return false, fmt.Errorf("failed to read OnRamp dest chain config: %w", err)
	}

	return result.Output.Router != (common.Address{}), nil
}

// ReadDestGasFields reads from all relevant contracts (OnRamp, FeeQuoter, verifiers).
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
	fields := make(map[string]uint32)

	// OnRamp
	onRampRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, onramp.ContractType, onramp.Version, "")
	if !datastore_utils.IsAddressRefEmpty(onRampRef) {
		onRampAddr := common.HexToAddress(onRampRef.Address)
		onRampCur, err := cldf_ops.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       onRampAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read OnRamp: %w", err)
		}
		if onRampCur.Output.Router != (common.Address{}) {
			fields[v2_0_0_adapters.OnRampBaseExecutionGasCost.Name] = onRampCur.Output.BaseExecutionGasCost
		}
	}

	// FeeQuoter
	fqRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, fee_quoter.ContractType, fee_quoter.Version, "")
	if !datastore_utils.IsAddressRefEmpty(fqRef) {
		fqAddr := common.HexToAddress(fqRef.Address)
		fqCur, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       fqAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read FeeQuoter: %w", err)
		}
		fields[v2_0_0_adapters.FeeQuoterDefaultTokenDestGasOverhead.Name] = fqCur.Output.DefaultTokenDestGasOverhead
		fields[v2_0_0_adapters.FeeQuoterMaxPerMsgGasLimit.Name] = fqCur.Output.MaxPerMsgGasLimit
		fields[v2_0_0_adapters.FeeQuoterDestGasPerPayloadByteBase.Name] = uint32(fqCur.Output.DestGasPerPayloadByteBase)
		fields[v2_0_0_adapters.FeeQuoterDefaultTxGasLimit.Name] = fqCur.Output.DefaultTxGasLimit
	}

	// CommitteeVerifier
	cvRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, committee_verifier.ContractType, committee_verifier.Version, "")
	if !datastore_utils.IsAddressRefEmpty(cvRef) {
		cvAddr := common.HexToAddress(cvRef.Address)
		cvCur, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       cvAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read CommitteeVerifier: %w", err)
		}
		if cvCur.Output.RemoteChainConfig.Router != (common.Address{}) {
			fields[v2_0_0_adapters.CommitteeVerifierGasForVerification.Name] = cvCur.Output.RemoteChainConfig.GasForVerification
		}
	}

	// LombardVerifier
	lvRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, lombard_verifier.ContractType, lombard_verifier.Version, "")
	if !datastore_utils.IsAddressRefEmpty(lvRef) {
		lvAddr := common.HexToAddress(lvRef.Address)
		lvCur, err := cldf_ops.ExecuteOperation(b, lombard_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       lvAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read LombardVerifier: %w", err)
		}
		if lvCur.Output.RemoteChainConfig.Router != (common.Address{}) {
			fields[v2_0_0_adapters.LombardVerifierGasForVerification.Name] = lvCur.Output.RemoteChainConfig.GasForVerification
		}
	}

	// CCTPVerifier
	ctpRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, cctp_verifier.ContractType, cctp_verifier.Version, "")
	if !datastore_utils.IsAddressRefEmpty(ctpRef) {
		ctpAddr := common.HexToAddress(ctpRef.Address)
		ctpCur, err := cldf_ops.ExecuteOperation(b, cctp_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       ctpAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read CCTPVerifier: %w", err)
		}
		if ctpCur.Output.RemoteChainConfig.Router != (common.Address{}) {
			fields[v2_0_0_adapters.USDCVerifierGasForVerification.Name] = ctpCur.Output.RemoteChainConfig.GasForVerification
		}
	}

	return fields, nil
}

// WriteDestGasFields writes the resolved gas field values to all relevant contracts via MCMS.
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
	var writes []contract.WriteOutput

	// OnRamp: BaseExecutionGasCost
	onRampRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, onramp.ContractType, onramp.Version, "")
	if !datastore_utils.IsAddressRefEmpty(onRampRef) {
		onRampAddr := common.HexToAddress(onRampRef.Address)
		onRampCur, err := cldf_ops.ExecuteOperation(b, onramp.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       onRampAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read OnRamp dest chain config: %w", err)
		}

		updated := onRampCur.Output
		if val, ok := resolved[v2_0_0_adapters.OnRampBaseExecutionGasCost.Name]; ok {
			updated.BaseExecutionGasCost = val
		}

		onRampWrite, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyOnRampDestChainConfigUpdates, chain, contract.FunctionInput[[]onramp.DestChainConfigArgs]{
			ChainSelector: srcChainSelector,
			Address:       onRampAddr,
			Args: []onramp.DestChainConfigArgs{
				{
					DestChainSelector:         targetChainSelector,
					Router:                    updated.Router,
					AddressBytesLength:        updated.AddressBytesLength,
					TokenReceiverAllowed:      updated.TokenReceiverAllowed,
					MessageNetworkFeeUSDCents: updated.MessageNetworkFeeUSDCents,
					TokenNetworkFeeUSDCents:   updated.TokenNetworkFeeUSDCents,
					BaseExecutionGasCost:      updated.BaseExecutionGasCost,
					DefaultExecutor:           updated.DefaultExecutor,
					DefaultCCVs:               updated.DefaultCCVs,
					LaneMandatedCCVs:          updated.LaneMandatedCCVs,
					OffRamp:                   updated.OffRamp,
				},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply OnRamp update: %w", err)
		}
		writes = append(writes, onRampWrite.Output)
	}

	// FeeQuoter: DefaultTokenDestGasOverhead, MaxPerMsgGasLimit, DestGasPerPayloadByteBase, DefaultTxGasLimit
	fqRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, fee_quoter.ContractType, fee_quoter.Version, "")
	if !datastore_utils.IsAddressRefEmpty(fqRef) {
		fqAddr := common.HexToAddress(fqRef.Address)
		fqCur, err := cldf_ops.ExecuteOperation(b, fee_quoter.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       fqAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read FeeQuoter dest chain config: %w", err)
		}

		updated := fqCur.Output
		if val, ok := resolved[v2_0_0_adapters.FeeQuoterDefaultTokenDestGasOverhead.Name]; ok {
			updated.DefaultTokenDestGasOverhead = val
		}
		if val, ok := resolved[v2_0_0_adapters.FeeQuoterMaxPerMsgGasLimit.Name]; ok {
			updated.MaxPerMsgGasLimit = val
		}
		if val, ok := resolved[v2_0_0_adapters.FeeQuoterDestGasPerPayloadByteBase.Name]; ok {
			updated.DestGasPerPayloadByteBase = uint8(val)
		}
		if val, ok := resolved[v2_0_0_adapters.FeeQuoterDefaultTxGasLimit.Name]; ok {
			updated.DefaultTxGasLimit = val
		}

		fqWrite, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyFeeQuoterDestChainConfigUpdates, chain, contract.FunctionInput[[]fee_quoter.DestChainConfigArgs]{
			ChainSelector: srcChainSelector,
			Address:       fqAddr,
			Args: []fee_quoter.DestChainConfigArgs{
				{DestChainSelector: targetChainSelector, DestChainConfig: updated},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to apply FeeQuoter update: %w", err)
		}
		writes = append(writes, fqWrite.Output)
	}

	// CommitteeVerifier: GasForVerification
	cvRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, committee_verifier.ContractType, committee_verifier.Version, "")
	if !datastore_utils.IsAddressRefEmpty(cvRef) {
		cvAddr := common.HexToAddress(cvRef.Address)
		cvCur, err := cldf_ops.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       cvAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read CommitteeVerifier remote chain config: %w", err)
		}

		if cvCur.Output.RemoteChainConfig.Router != (common.Address{}) {
			updated := cvCur.Output.RemoteChainConfig
			if val, ok := resolved[v2_0_0_adapters.CommitteeVerifierGasForVerification.Name]; ok {
				updated.GasForVerification = val
			}

			cvWrite, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyCommitteeVerifierRemoteChainConfigUpdates, chain, contract.FunctionInput[[]committee_verifier.RemoteChainConfigArgs]{
				ChainSelector: srcChainSelector,
				Address:       cvAddr,
				Args:          []committee_verifier.RemoteChainConfigArgs{updated},
			})
			if err != nil {
				return nil, fmt.Errorf("failed to apply CommitteeVerifier update: %w", err)
			}
			writes = append(writes, cvWrite.Output)
		}
	}

	// LombardVerifier: GasForVerification
	lvRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, lombard_verifier.ContractType, lombard_verifier.Version, "")
	if !datastore_utils.IsAddressRefEmpty(lvRef) {
		lvAddr := common.HexToAddress(lvRef.Address)
		lvCur, err := cldf_ops.ExecuteOperation(b, lombard_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       lvAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read LombardVerifier remote chain config: %w", err)
		}

		if lvCur.Output.RemoteChainConfig.Router != (common.Address{}) {
			updated := lvCur.Output.RemoteChainConfig
			if val, ok := resolved[v2_0_0_adapters.LombardVerifierGasForVerification.Name]; ok {
				updated.GasForVerification = val
			}

			lvWrite, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyLombardVerifierRemoteChainConfigUpdates, chain, contract.FunctionInput[[]lombard_verifier.RemoteChainConfigArgs]{
				ChainSelector: srcChainSelector,
				Address:       lvAddr,
				Args:          []lombard_verifier.RemoteChainConfigArgs{updated},
			})
			if err != nil {
				return nil, fmt.Errorf("failed to apply LombardVerifier update: %w", err)
			}
			writes = append(writes, lvWrite.Output)
		}
	}

	// CCTPVerifier: GasForVerification
	ctpRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, cctp_verifier.ContractType, cctp_verifier.Version, "")
	if !datastore_utils.IsAddressRefEmpty(ctpRef) {
		ctpAddr := common.HexToAddress(ctpRef.Address)
		ctpCur, err := cldf_ops.ExecuteOperation(b, cctp_verifier.GetRemoteChainConfig, chain, contract.FunctionInput[uint64]{
			ChainSelector: srcChainSelector,
			Address:       ctpAddr,
			Args:          targetChainSelector,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to read CCTPVerifier remote chain config: %w", err)
		}

		if ctpCur.Output.RemoteChainConfig.Router != (common.Address{}) {
			updated := ctpCur.Output.RemoteChainConfig
			if val, ok := resolved[v2_0_0_adapters.USDCVerifierGasForVerification.Name]; ok {
				updated.GasForVerification = val
			}

			ctpWrite, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyCCTPVerifierRemoteChainConfigUpdates, chain, contract.FunctionInput[[]cctp_verifier.RemoteChainConfigArgs]{
				ChainSelector: srcChainSelector,
				Address:       ctpAddr,
				Args:          []cctp_verifier.RemoteChainConfigArgs{updated},
			})
			if err != nil {
				return nil, fmt.Errorf("failed to apply CCTPVerifier update: %w", err)
			}
			writes = append(writes, ctpWrite.Output)
		}
	}

	// Convert all writes to a single batch operation
	if len(writes) == 0 {
		return []mcms_types.BatchOperation{}, nil
	}

	batchOp, err := contract.NewBatchOperationFromWrites(writes)
	if err != nil {
		return nil, fmt.Errorf("failed to build batch operation: %w", err)
	}

	return []mcms_types.BatchOperation{batchOp}, nil
}

// ReadImmutableSanityFields reads OffRamp immutable fields.
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
		return map[string]uint32{}, nil
	}

	offRampAddr := common.HexToAddress(offRampRef.Address)
	result, err := cldf_ops.ExecuteOperation(b, offramp.GetStaticConfig, chain, contract.FunctionInput[struct{}]{
		ChainSelector: srcChainSelector,
		Address:       offRampAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read OffRamp: %w", err)
	}

	return map[string]uint32{
		"OffRamp.GasForCallExactCheck":      uint32(result.Output.GasForCallExactCheck),
		"OffRamp.MaxGasBufferToUpdateState": result.Output.MaxGasBufferToUpdateState,
	}, nil
}

// DiscoverCandidateTokens returns token pool addresses from the datastore.
// For v2.0.0, this returns the addresses of token pools (Lombard/USDC) that are registered.
func (a *GlamsterdamGasAdapter) DiscoverCandidateTokens(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector uint64,
) ([][]byte, error) {
	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	var tokenPools [][]byte

	// Look for token pool addresses in the datastore
	for _, ref := range addrs {
		// For v2.0.0 Glamsterdam, we discover any TokenPool type addresses
		// In practice, these are Lombard and USDC pools
		if ref.Type == datastore.ContractType(token_pool.ContractType) {
			tokenAddr := common.HexToAddress(ref.Address)
			tokenPools = append(tokenPools, tokenAddr.Bytes())
		}
	}

	return tokenPools, nil
}

// ReadTokenGasField reads the DestGasOverhead from a TokenPool's TokenTransferFeeConfig.
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
	poolRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, token_pool.ContractType, token_pool.Version, "")
	if datastore_utils.IsAddressRefEmpty(poolRef) {
		return 0, false, fmt.Errorf("could not resolve TokenPool address on chain %d", srcChainSelector)
	}

	poolAddr := common.HexToAddress(poolRef.Address)
	tokenAddr := common.BytesToAddress(token)

	result, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
		ChainSelector: srcChainSelector,
		Address:       poolAddr,
		Args: token_pool.GetTokenTransferFeeConfigArgs{
			Arg0:              tokenAddr,
			DestChainSelector: targetChainSelector,
			Arg2:              [4]byte{},
			Arg3:              []byte{},
		},
	})
	if err != nil {
		return 0, false, fmt.Errorf("failed to read token transfer fee config: %w", err)
	}

	return result.Output.DestGasOverhead, result.Output.IsEnabled, nil
}

// WriteTokenGasField writes the DestGasOverhead to a TokenPool via MCMS.
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
	poolRef := datastore_utils.GetAddressRef(addrs, srcChainSelector, token_pool.ContractType, token_pool.Version, "")
	if datastore_utils.IsAddressRefEmpty(poolRef) {
		return mcms_types.BatchOperation{}, fmt.Errorf("could not resolve TokenPool address on chain %d", srcChainSelector)
	}

	poolAddr := common.HexToAddress(poolRef.Address)
	tokenAddr := common.BytesToAddress(token)

	// Read current config to preserve other fields
	current, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
		ChainSelector: srcChainSelector,
		Address:       poolAddr,
		Args: token_pool.GetTokenTransferFeeConfigArgs{
			Arg0:              tokenAddr,
			DestChainSelector: targetChainSelector,
			Arg2:              [4]byte{},
			Arg3:              []byte{},
		},
	})
	if err != nil {
		return mcms_types.BatchOperation{}, fmt.Errorf("failed to read current token transfer fee config: %w", err)
	}

	// Update only the DestGasOverhead field
	updated := current.Output
	updated.DestGasOverhead = value

	// Execute the write operation
	writeOut, err := cldf_ops.ExecuteOperation(b, glamsterdam.ApplyTokenPoolTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs]{
		ChainSelector: srcChainSelector,
		Address:       poolAddr,
		Args: token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
			TokenTransferFeeConfigArgs: []token_pool.TokenTransferFeeConfigArgs{
				{DestChainSelector: targetChainSelector, TokenTransferFeeConfig: updated},
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
