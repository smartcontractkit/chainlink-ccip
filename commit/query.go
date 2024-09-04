package commit

import (
	"context"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

func (p *Plugin) Query(_ context.Context, outCtx ocr3types.OutcomeContext) (types.Query, error) {
	// Start a timer for MaxQueryDuration-ε
	// >>> t_query = new_timer(MaxQueryDuration-ε)

	// >>> If we are in the building_reports state or t_observations_initial_request has expired:
	//        for rmn_node in p.rmn_nodes:
	//           lanes = lanes that rmn_node can read the source chain based on RmnHomeConfig
	//           ranges = previous_round_ranges_for_lanes(lanes)
	//           req = new_observation_request(lanes, ranges)
	//
	// Make sure a sufficient number of requests are sent for each lane.
	//
	// Start a timer
	// >>> t_observations_initial_request = new_timer(delta)
	//
	// >>> responses, ok = receive a sufficient number of observation responses
	//     if ok:
	//          dest_lane_updates = process_responses(responses)
	//          req = new_report_signature_request(dest_lane_updates)
	//          send(req) to a sufficient number of destination chain signers
	//          t_reports_initial_request = new_timer(delta)
	//
	// if t_reports_initial_request has expired:
	//   req = new_report_signature_request(dest_lane_updates) // previously prepared but not sent
	//
	// Upon receiving sufficient number of report signatures use the signatures in the Query:
	//
	// Upon t_query expiration:
	//    Indicate to the oracles that they should retry.
	//
	// ----------------------------------------------------
	//

	/*
		// low level rmn comm
		type RMNComm interface {
			ObservationRequest(lanes, intervals protoGeneratedTypes) <-chan ObservationResponse
			ReportSignatureRequest(signed_observations protoGeneratedTypes) <-chan ReportSignatureResponse
		}

		// plugin dep
		type RMN interface {
			tbd
		}
	*/

	return types.Query{}, nil
}
