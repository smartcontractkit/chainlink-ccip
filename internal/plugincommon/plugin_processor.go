package plugincommon

import (
	"context"

	"github.com/smartcontractkit/libocr/commontypes"
)

type AttributedObservation[ObservationType any] struct {
	OracleID    commontypes.OracleID
	Observation ObservationType
}

// PluginProcessor is to encapsulate logic for multiple processors under a OCR plugin.
// This makes it easier to manage and test when there are multiple logical components  of a single plugin.
// Some of them will implement state machines (e.g. merkleroot),
// others might implement simpler logic. (e.g. tokenprices, chainfee)
// The OCR plugin becomes a coordinator/collector of these processors.
// Example Pseudo code:
//
//			type OCRPlugin struct {
//		 	   merkleRootsProcessor
//			   tokenPriceProcessor
//		 	   chainFeeProcessor
//			}
//
//	     // Observation excludes error handling for brevity.
//		 	func (p *OCRPlugin) Observation() ocrtype.Observation {
//		 	 	 return ocrtype.Observation {
//		           merkleRoots: p.merkleRootsProcessor.Observation(...)
//		           tokenPrices: p.tokenPriceProcessor.Observation(...)
//		           chainFees: p.chainFeeProcessor.Observation(...)
//				 }
//	      	}
//
// Notice all interface functions are using prevOutcome instead of outCtx.
// We're interested in the prevOutcome, and it makes it easier to have all decoding on the top level (OCR plugin),
// otherwise there might be cyclic dependencies or just complicating the code more.
type PluginProcessor[QueryType any, ObservationType any, OutcomeType any] interface {
	Query(ctx context.Context, prevOutcome OutcomeType) (QueryType, error)
	Observation(ctx context.Context, prevOutcome OutcomeType, query QueryType) (ObservationType, error)
	ValidateObservation(prevOutcome OutcomeType, query QueryType, ao AttributedObservation[ObservationType]) error
	Outcome(prevOutcome OutcomeType, query QueryType, aos []AttributedObservation[ObservationType]) (OutcomeType, error)
}
