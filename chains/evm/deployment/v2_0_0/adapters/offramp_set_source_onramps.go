package adapters

import (
	"bytes"
	"errors"
	"fmt"
	"math"
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

var (
	_ ccvadapters.OffRampSourceOnRampSetter = (*ChainFamilyAdapter)(nil)
	_ ccvadapters.OffRampSourceOnRampReader = (*ChainFamilyAdapter)(nil)
)

// GetOffRampSourceOnRamps implements [ccvadapters.OffRampSourceOnRampReader].
func (a *ChainFamilyAdapter) GetOffRampSourceOnRamps(
	e cldf.Environment,
	localChainSelector uint64,
	sourceChainSelector uint64,
) ([]string, error) {
	chain, ok := e.BlockChains.EVMChains()[localChainSelector]
	if !ok {
		return nil, fmt.Errorf("EVM chain %d not found in environment", localChainSelector)
	}

	offRampBytes, err := a.GetOffRampAddress(e.DataStore, localChainSelector)
	if err != nil {
		return nil, fmt.Errorf("resolve OffRamp on chain %d: %w", localChainSelector, err)
	}
	offRampAddr := common.BytesToAddress(offRampBytes)

	report, err := cldf_ops.ExecuteOperation(e.OperationsBundle, offramp.GetSourceChainConfig, chain, contract.FunctionInput[uint64]{
		ChainSelector: chain.Selector,
		Address:       offRampAddr,
		Args:          sourceChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("get source chain config for %d on OffRamp %s: %w",
			sourceChainSelector, offRampAddr, err)
	}

	onRamps := make([]string, 0, len(report.Output.OnRamps))
	for _, onRamp := range report.Output.OnRamps {
		onRamps = append(onRamps, hexutil.Encode(onRamp))
	}
	return onRamps, nil
}

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

// parseOffRampSourceOnRampAddresses decodes the operator-supplied onramps into the bytes
// stored on the OffRamp. The source chain may be of any family, so the bytes are taken
// exactly as given: the OffRamp matches an incoming message by hashing the onramp bytes it
// carries, and only the source family knows how it encodes them.
func parseOffRampSourceOnRampAddresses(addrs []string) ([][]byte, error) {
	out := make([][]byte, 0, len(addrs))
	seen := make(map[string]struct{}, len(addrs))
	for i, addr := range addrs {
		raw, err := decodeOffRampOnRampHex(addr)
		if err != nil {
			return nil, fmt.Errorf("onRamps[%d]: %w", i, err)
		}
		key := string(raw)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, raw)
	}
	if len(out) == 0 {
		return nil, errors.New("no valid onRamp addresses after parsing")
	}
	return out, nil
}

// decodeOffRampOnRampHex decodes a 0x-prefixed hex string into the raw onramp bytes. The
// only bound is the wire format's: MessageV1 length-prefixes the onramp address with a
// uint8, so it cannot exceed 255 bytes.
func decodeOffRampOnRampHex(addr string) ([]byte, error) {
	raw, err := hexutil.Decode(strings.TrimSpace(addr))
	if err != nil {
		return nil, fmt.Errorf("invalid hex address %q: %w", addr, err)
	}
	if len(raw) == 0 || len(raw) > math.MaxUint8 {
		return nil, fmt.Errorf("address %q must be 1-%d bytes, got %d", addr, math.MaxUint8, len(raw))
	}
	return raw, nil
}
