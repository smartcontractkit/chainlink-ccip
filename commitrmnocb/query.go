package commitrmnocb

import (
	"context"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// Query Depending on the current state, queries RMN for sequence numbers or signed roots
func (p *Plugin) Query(_ context.Context, outCtx ocr3types.OutcomeContext) (types.Query, error) {
	_, nextState := p.decodeOutcome(outCtx.PreviousOutcome)

	switch nextState {
	case SelectingRangesForReport:
		return p.BuildRmnSeqNumsQuery()

	case BuildingReport:
		return types.Query{}, nil

	default:
		return types.Query{}, nil
	}
}

// BuildRmnSeqNumsQuery builds a Query that contains OnRamp max seq nums from RMN
func (p *Plugin) BuildRmnSeqNumsQuery() (types.Query, error) {
	rmnMaxSourceSeqNums := make([]plugintypes.SeqNumChain, 0)

	encodedQuery, err := NewCommitQuery(rmnMaxSourceSeqNums, nil).Encode()
	if err != nil {
		return types.Query{}, err
	}

	return encodedQuery, nil
}
