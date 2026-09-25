package adapters

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/cctp_through_ccv_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/cctp_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/lombard_verifier"
	glamsterdamutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/glamsterdam"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	v2_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

// GlamsterdamGasAdapter implements v2_0_0_adapters.GasUpdateAdapter for EVM chains.
type GlamsterdamGasAdapter struct{}

// tokenPoolKind describes one TokenPool implementation this version's mapping table drives, and
// which FieldSpec (Lombard row 9, or USDC row 10) applies to it. USDC has two possible pool
// implementations — SiloedUSDCTokenPool and CCTPThroughCCVTokenPool (for CCTP-capable EVM
// remotes) — both governed by the same USDC row.
type tokenPoolKind struct {
	ContractType cldf_deployment.ContractType
	Version      *semver.Version
	FieldSpec    glamsterdamutils.FieldSpec[uint32]
}

var glamsterdamTokenPoolKinds = []tokenPoolKind{
	{lombard_token_pool.ContractType, lombard_token_pool.Version, v2_0_0_adapters.LombardTokenPoolDestGasOverhead},
	{siloed_usdc_token_pool.ContractType, siloed_usdc_token_pool.Version, v2_0_0_adapters.USDCTokenPoolDestGasOverhead},
	{cctp_through_ccv_token_pool.ContractType, cctp_through_ccv_token_pool.Version, v2_0_0_adapters.USDCTokenPoolDestGasOverhead},
}

// HasLaneToTarget checks if a lane exists by checking FeeQuoter's DestChainConfig.IsEnabled,
// matching the discovery semantics used elsewhere for this version (see
// v2_0_0/sequences/glamsterdam/discovery.go). OnRamp's Router field is not a reliable signal here:
// it can be set on a chain whose FeeQuoter has no dest chain config for the target at all (e.g.
// this chain's OnRamp is configured for other destinations), which would otherwise cause every
// gas-config write for the rest of the mapping table to be attempted against an undiscovered
// lane.
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

	return result.Output.IsEnabled, nil
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
// findTokenPoolRef returns the datastore ref (and its tokenPoolKind) for the given pool address
// on this chain, if it matches one of glamsterdamTokenPoolKinds. "token" candidates from
// DiscoverCandidateTokens are always pool addresses (see below), so this resolves back to the
// kind whenever the interface hands us one.
func findTokenPoolRef(addrs []datastore.AddressRef, poolAddr common.Address) (datastore.AddressRef, tokenPoolKind, bool) {
	for _, ref := range addrs {
		if common.HexToAddress(ref.Address) != poolAddr {
			continue
		}
		for _, kind := range glamsterdamTokenPoolKinds {
			if ref.Type == datastore.ContractType(kind.ContractType) && ref.Version.Equal(kind.Version) {
				return ref, kind, true
			}
		}
	}
	return datastore.AddressRef{}, tokenPoolKind{}, false
}

// DiscoverCandidateTokens returns every Lombard/USDC token pool address deployed on this chain
// (across every qualifier, since Lombard in particular can have more than one qualifier-scoped
// pool per chain). Each pool wraps exactly one token, so the pool's own address is what
// ReadTokenGasField/WriteTokenGasField/TokenFieldSpec key off of.
func (a *GlamsterdamGasAdapter) DiscoverCandidateTokens(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector uint64,
) ([][]byte, error) {
	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	seen := make(map[common.Address]bool)
	var tokenPools [][]byte

	for _, ref := range addrs {
		for _, kind := range glamsterdamTokenPoolKinds {
			if ref.Type != datastore.ContractType(kind.ContractType) || !ref.Version.Equal(kind.Version) {
				continue
			}
			addr := common.HexToAddress(ref.Address)
			if !seen[addr] {
				seen[addr] = true
				tokenPools = append(tokenPools, addr.Bytes())
			}
		}
	}

	return tokenPools, nil
}

// ReadTokenGasField reads the DestGasOverhead from a TokenPool's TokenTransferFeeConfig. token is
// the pool's own address (see DiscoverCandidateTokens); v2.0.0's TokenPool.getTokenTransferFeeConfig
// takes no token argument since each pool wraps exactly one token.
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

	poolAddr := common.BytesToAddress(token)
	result, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
		ChainSelector: srcChainSelector,
		Address:       poolAddr,
		Args:          token_pool.GetTokenTransferFeeConfigArgs{DestChainSelector: targetChainSelector},
	})
	if err != nil {
		return 0, false, fmt.Errorf("failed to read token transfer fee config from pool %s: %w", poolAddr, err)
	}

	return result.Output.DestGasOverhead, result.Output.IsEnabled, nil
}

// TokenFieldSpec identifies whether token (a pool address) is a Lombard or USDC pool via its
// datastore ContractType, and returns the matching FieldSpec — the two rows have different
// Prague/Glamsterdam baselines and must not be conflated.
func (a *GlamsterdamGasAdapter) TokenFieldSpec(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector uint64,
	token []byte,
) (glamsterdamutils.FieldSpec[uint32], error) {
	addrs := ds.Addresses().Filter(datastore.AddressRefByChainSelector(srcChainSelector))
	poolAddr := common.BytesToAddress(token)
	_, kind, found := findTokenPoolRef(addrs, poolAddr)
	if !found {
		return glamsterdamutils.FieldSpec[uint32]{}, fmt.Errorf(
			"could not resolve token pool kind for address %s on chain %d", poolAddr, srcChainSelector,
		)
	}
	return kind.FieldSpec, nil
}

// WriteTokenGasField writes the DestGasOverhead to the TokenPool at token (a pool address),
// preserving every other field of its current TokenTransferFeeConfig.
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

	poolAddr := common.BytesToAddress(token)

	current, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
		ChainSelector: srcChainSelector,
		Address:       poolAddr,
		Args:          token_pool.GetTokenTransferFeeConfigArgs{DestChainSelector: targetChainSelector},
	})
	if err != nil {
		return mcms_types.BatchOperation{}, fmt.Errorf("failed to read token transfer fee config from pool %s: %w", poolAddr, err)
	}

	updated := current.Output
	updated.DestGasOverhead = value

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
		return mcms_types.BatchOperation{}, fmt.Errorf("failed to apply token transfer fee config update to pool %s: %w", poolAddr, err)
	}

	batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{writeOut.Output})
	if err != nil {
		return mcms_types.BatchOperation{}, fmt.Errorf("failed to build batch operation: %w", err)
	}

	return batchOp, nil
}
