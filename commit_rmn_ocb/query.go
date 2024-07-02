package commitrmnocb

import (
	"context"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// Query TODO: doc
func (p *Plugin) Query(_ context.Context, outCtx ocr3types.OutcomeContext) (types.Query, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)

	switch nextState {
	case SelectingRangesForReport:
		return p.BuildRmnSeqNumsQuery()

	case BuildingReport:
		return p.BuildSignedRootsQuery(previousOutcome)

	default:
		return types.Query{}, nil
	}
}

// BuildRmnSeqNumsQuery TODO: doc
func (p *Plugin) BuildRmnSeqNumsQuery() (types.Query, error) {
	rmnMaxSourceSeqNums, err := p.rmn.RequestMaxSeqNums(p.cfg.AllSourceChains)
	if err != nil {
		return types.Query{}, err
	}

	encodedQuery, err := NewCommitQuery(rmnMaxSourceSeqNums, nil).Encode()
	if err != nil {
		return types.Query{}, err
	}

	return encodedQuery, nil
}

// BuildSignedRootsQuery TODO: doc
func (p *Plugin) BuildSignedRootsQuery(previousOutcome CommitPluginOutcome) (types.Query, error) {
	signedRoots, err := p.rmn.RequestSignedIntervals(previousOutcome.RangesSelectedForReport)
	if err != nil {
		return types.Query{}, err
	}

	encodedQuery, err := NewCommitQuery(nil, signedRoots).Encode()
	if err != nil {
		return types.Query{}, err
	}

	return encodedQuery, nil
}
