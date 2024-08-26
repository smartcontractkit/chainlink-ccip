package commit

import (
	"context"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

func (p *Plugin) Query(_ context.Context, outCtx ocr3types.OutcomeContext) (types.Query, error) {
	return types.Query{}, nil
}
