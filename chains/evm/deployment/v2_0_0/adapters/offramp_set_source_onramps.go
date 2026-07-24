package adapters

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

var _ ccvadapters.OffRampSourceOnRampSetter = (*ChainFamilyAdapter)(nil)

// SetOffRampSourceOnRamps updates the OffRamp source-chain onramp whitelist on an EVM chain.
func (a *ChainFamilyAdapter) SetOffRampSourceOnRamps(
	e cldf.Environment,
	update ccvadapters.OffRampSetSourceOnRampsEntry,
) (*mcms_types.BatchOperation, bool, error) {
	chain, ok := e.BlockChains.EVMChains()[update.LocalChainSelector]
	if !ok {
		return nil, false, fmt.Errorf("EVM chain %d not found in environment", update.LocalChainSelector)
	}

	offRampBytes, err := a.GetOffRampAddress(e.DataStore, update.LocalChainSelector)
	if err != nil {
		return nil, false, fmt.Errorf("resolve OffRamp on chain %d: %w", update.LocalChainSelector, err)
	}
	offRampAddr := common.BytesToAddress(offRampBytes)

	desiredOnRamps, err := parseOffRampSourceOnRampAddresses(update.OnRamps)
	if err != nil {
		return nil, false, err
	}

	currentReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, offramp.GetSourceChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chain.Selector,
		Address:       offRampAddr,
		Args:          update.SourceChainSelector,
	})
	if err != nil {
		return nil, false, fmt.Errorf("get source chain config for %d on OffRamp %s: %w",
			update.SourceChainSelector, offRampAddr, err)
	}
	current := currentReport.Output

	if sequences.UnorderedSliceEqual(current.OnRamps, desiredOnRamps, bytes.Equal) {
		e.Logger.Infow("OffRamp source onramp whitelist already matches desired state, skipping",
			"localChain", update.LocalChainSelector,
			"sourceChain", update.SourceChainSelector,
			"offRamp", offRampAddr.Hex(),
			"onRampCount", len(desiredOnRamps),
		)
		return nil, true, nil
	}

	desired := offramp.SourceChainConfigArgs{
		Router:              current.Router,
		SourceChainSelector: update.SourceChainSelector,
		IsEnabled:           current.IsEnabled,
		OnRamps:             desiredOnRamps,
		DefaultCCVs:         current.DefaultCCVs,
		LaneMandatedCCVs:    current.LaneMandatedCCVs,
	}

	offRampReport, err := cldf_ops.ExecuteOperation(e.OperationsBundle, offramp.ApplySourceChainConfigUpdates, chain, contract.FunctionInput[[]offramp.SourceChainConfigArgs]{
		ChainSelector: chain.Selector,
		Address:       offRampAddr,
		Args:          []offramp.SourceChainConfigArgs{desired},
	})
	if err != nil {
		return nil, false, fmt.Errorf("apply source chain config on OffRamp %s: %w", offRampAddr, err)
	}

	batchOp, err := contract.NewBatchOperationFromWrites([]contract.WriteOutput{offRampReport.Output})
	if err != nil {
		return nil, false, fmt.Errorf("build batch operation: %w", err)
	}
	return &batchOp, false, nil
}

func parseOffRampSourceOnRampAddresses(addrs []string) ([][]byte, error) {
	out := make([][]byte, 0, len(addrs))
	seen := make(map[string]struct{}, len(addrs))
	for i, addr := range addrs {
		raw, err := decodeOffRampOnRampHex(addr)
		if err != nil {
			return nil, fmt.Errorf("onRamps[%d]: %w", i, err)
		}
		// Same left-pad-to-32 used when writing OffRamp source config in configure_chain_for_lanes.
		encoded := common.LeftPadBytes(raw, 32)
		key := string(encoded)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, encoded)
	}
	if len(out) == 0 {
		return nil, errors.New("no valid onRamp addresses after parsing")
	}
	return out, nil
}

// decodeOffRampOnRampHex follows the parseRemoteAdapter pattern (20-byte IsHexAddress or hexutil.Decode).
func decodeOffRampOnRampHex(addr string) ([]byte, error) {
	trimmed := strings.TrimSpace(addr)
	if common.IsHexAddress(trimmed) {
		return common.HexToAddress(trimmed).Bytes(), nil
	}
	raw, err := hexutil.Decode(trimmed)
	if err != nil {
		return nil, fmt.Errorf("invalid hex address %q: %w", addr, err)
	}
	if len(raw) == 0 || len(raw) > 32 {
		return nil, fmt.Errorf("address %q must be 1-32 bytes, got %d", addr, len(raw))
	}
	if len(raw) != common.AddressLength && len(raw) != 32 {
		return nil, fmt.Errorf("address %q must be 20 or 32 bytes, got %d", addr, len(raw))
	}
	return raw, nil
}
