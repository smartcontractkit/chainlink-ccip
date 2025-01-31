package merkleroot

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"
)

func (p *Processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	if !p.offchainCfg.RMNEnabled {
		return Query{}, nil
	}

	nextState := prevOutcome.nextState()
	if nextState != buildingReport {
		return Query{}, nil
	}

	if prevOutcome.RMNRemoteCfg.IsEmpty() {
		return Query{}, fmt.Errorf("RMN report config is empty")
	}

	if err := p.prepareRMNController(ctx, lggr, prevOutcome); err != nil {
		return Query{}, fmt.Errorf("initialize RMN controller: %w", err)
	}

	offRampAddress, err := p.ccipReader.GetContractAddress(consts.ContractNameOffRamp, p.destChain)
	if err != nil {
		return Query{}, fmt.Errorf("get offRamp contract address: %w", err)
	}
	dstChainInfo := &rmnpb.LaneDest{
		DestChainSelector: uint64(p.destChain),
		OfframpAddress:    offRampAddress,
	}

	reqUpdates := make([]*rmnpb.FixedDestLaneUpdateRequest, 0, len(prevOutcome.RangesSelectedForReport))
	for _, sourceChainRange := range prevOutcome.RangesSelectedForReport {
		onRampAddress, err := p.ccipReader.GetContractAddress(consts.ContractNameOnRamp, sourceChainRange.ChainSel)
		if err != nil {
			lggr.Warnw("failed to get onRamp address", "chain", sourceChainRange.ChainSel, "err", err)
			continue
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

	ctxQuery, cancel := context.WithTimeout(ctx, p.offchainCfg.RMNSignaturesTimeout)
	defer cancel()

	rmnReportStart := time.Now()
	// Get signatures for the requested updates. The signatures might contain a subset of the requested updates.
	// While building the report the plugin should exclude source chain updates without signatures.
	sigs, err := p.rmnController.ComputeReportSignatures(ctxQuery, dstChainInfo, reqUpdates, prevOutcome.RMNRemoteCfg)
	if err != nil {
		p.metricsReporter.TrackRmnReport(float64(time.Since(rmnReportStart).Milliseconds()), false)
		if errors.Is(err, rmn.ErrTimeout) {
			lggr.Errorf("RMN timeout while computing signatures for %d updates for chain %v",
				len(reqUpdates), dstChainInfo)
			return Query{RetryRMNSignatures: true}, nil
		}

		if errors.Is(err, rmn.ErrAllChainsNotReady) {
			lggr.Infow("none of the observation requests were ready to be observed by RMN, got empty responses")
			return Query{}, nil
		}

		return Query{}, fmt.Errorf("compute RMN signatures: %w", err)
	}

	p.metricsReporter.TrackRmnReport(time.Since(rmnReportStart).Seconds(), true)
	return Query{RetryRMNSignatures: false, RMNSignatures: sigs}, nil
}
