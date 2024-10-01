package merkleroot

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
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
	dstChainInfo := &rmnpb.LaneDest{
		DestChainSelector: uint64(w.cfg.DestChain),
		OfframpAddress:    offRampAddress,
	}

	reqUpdates := make([]*rmnpb.FixedDestLaneUpdateRequest, 0, len(prevOutcome.RangesSelectedForReport))
	for _, sourceChainRange := range prevOutcome.RangesSelectedForReport {
		onRampAddress, err := w.ccipReader.GetContractAddress(consts.ContractNameOnRamp, sourceChainRange.ChainSel)
		if err != nil {
			return Query{}, fmt.Errorf("get onRamp address for chain %v: %w", sourceChainRange.ChainSel, err)
		}

		reqUpdates = append(reqUpdates, &rmnpb.FixedDestLaneUpdateRequest{
			LaneSource: &rmnpb.LaneSource{
				SourceChainSelector: uint64(sourceChainRange.ChainSel),
				OnrampAddress:       onRampAddress,
			},
			ClosedInterval: &rmnpb.ClosedInterval{
				MinMsgNr: uint64(sourceChainRange.SeqNumRange.Start()),
				MaxMsgNr: uint64(sourceChainRange.SeqNumRange.End()),
			},
		})
	}

	if prevOutcome.RMNRemoteCfg.IsEmpty() {
		return Query{}, fmt.Errorf("RMN report config is empty")
	}

	ctxQuery, cancel := context.WithTimeout(ctx, w.cfg.RMNSignaturesTimeout)
	defer cancel()

	sigs, err := w.rmnClient.ComputeReportSignatures(ctxQuery, dstChainInfo, reqUpdates)
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
