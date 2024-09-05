package merkleroot

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
)

func (w *Processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	if !w.cfg.RMNEnabled {
		return Query{}, nil
	}

	nextState := prevOutcome.NextState()
	if nextState != BuildingReport {
		return Query{}, nil
	}

	offRampAddress, err := w.ccipReader.GetContractAddress(consts.ContractNameOffRamp, w.cfg.DestChain)
	if err != nil {
		return Query{}, fmt.Errorf("get offRamp contract address: %w", err)
	}
	dstChainInfo := rmn.DestChainInfo{
		Chain:          w.cfg.DestChain,
		OffRampAddress: offRampAddress,
	}

	reqUpdates := make([]rmn.FixedDestLaneUpdateRequest, 0, len(prevOutcome.RangesSelectedForReport))
	for _, sourceChainRange := range prevOutcome.RangesSelectedForReport {
		onRampAddress, err := w.ccipReader.GetContractAddress(consts.ContractNameOnRamp, sourceChainRange.ChainSel)
		if err != nil {
			return Query{}, fmt.Errorf("get onRamp address for chain %v: %w", sourceChainRange.ChainSel, err)
		}

		reqUpdates = append(reqUpdates, rmn.FixedDestLaneUpdateRequest{
			SourceChainInfo: rmn.SourceChainInfo{
				Chain:         sourceChainRange.ChainSel,
				OnRampAddress: onRampAddress,
			},
			SeqNumRange: sourceChainRange.SeqNumRange,
		})
	}

	ctxQuery, cancel := context.WithTimeout(ctx, w.cfg.RMNSignaturesTimeout)
	defer cancel()

	sigs, err := w.rmnClient.ComputeSignatures(ctxQuery, dstChainInfo, reqUpdates)
	if err != nil {
		if errors.Is(err, rmn.ErrTimeout) {
			w.lggr.Errorf("RMN timeout while computing signatures for %d updates for chain %v",
				len(reqUpdates), dstChainInfo)
			return Query{RetryRMNSignatures: true}, nil
		}
		return Query{}, fmt.Errorf("compute RMN signatures: %w", err)
	}

	return Query{RetryRMNSignatures: false, RMNSignatures: sigs}, nil
}
