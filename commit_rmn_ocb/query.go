package commitrmnocb

import (
	"context"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// Query Depending on the current state, queries RMN for sequence numbers or signed roots
func (p *Plugin) Query(_ context.Context, outCtx ocr3types.OutcomeContext) (types.Query, error) {
	previousOutcome, nextState := p.decodeOutcome(outCtx.PreviousOutcome)

	switch nextState {
	case SelectingRangesForReport:
		return p.BuildRmnSeqNumsQuery()

	case BuildingReport:
		return p.BuildMerkleRootsQuery(previousOutcome)

	default:
		return types.Query{}, nil
	}
}

// BuildRmnSeqNumsQuery builds a Query that contains OnRamp max seq nums from RMN
func (p *Plugin) BuildRmnSeqNumsQuery() (types.Query, error) {
	rmnMaxSourceSeqNums, err := p.rmn.RequestOnRampMaxSeqNums(p.knownSourceChainsSlice())
	if err != nil {
		return types.Query{}, err
	}

	encodedQuery, err := NewCommitQuery(rmnMaxSourceSeqNums, nil).Encode()
	if err != nil {
		return types.Query{}, err
	}

	return encodedQuery, nil
}

// BuildMerkleRootsQuery builds a Query that contains RMN signed roots
func (p *Plugin) BuildMerkleRootsQuery(previousOutcome CommitPluginOutcome) (types.Query, error) {
	signedRoots, err := p.rmn.RequestMerkleRoots(previousOutcome.RangesSelectedForReport)
	if err != nil {
		return types.Query{}, err
	}

	encodedQuery, err := NewCommitQuery(nil, signedRoots).Encode()
	if err != nil {
		return types.Query{}, err
	}

	return encodedQuery, nil
}
