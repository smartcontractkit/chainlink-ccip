package commit

import (
	"context"
	"errors"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// ObservationNext contains new Observation logic with breaking changes that should replace the existing
// Observation in the next release.
//
// NOTE: If you are building a feature make sure to include your changes here.
func (p *Plugin) ObservationNext(
	ctx context.Context, outCtx ocr3types.OutcomeContext, q types.Query,
) (types.Observation, error) {
	return types.Observation{}, errors.New("not implemented")
}

// OutcomeNext contains new Outcome logic with breaking changes that should replace the existing
// Outcome in the next release.
//
// NOTE: If you are building a feature make sure to include your changes here.
func (p *Plugin) OutcomeNext(
	ctx context.Context, outCtx ocr3types.OutcomeContext, q types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	return ocr3types.Outcome{}, errors.New("not implemented")
}
