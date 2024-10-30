package rmn

import (
	"math"
	"time"

	"github.com/smartcontractkit/libocr/networking"
	"github.com/smartcontractkit/libocr/ragep2p"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

const (
	// estimatedMaxNumberOfSourceChains is the estimated maximum number of source chains
	// that the current stream configuration supports and can be increased if required.
	// This value does not correlate to the maximum number of source chains that CCIP can support.
	estimatedMaxNumberOfSourceChains = 500

	// initialObservationRequest + observationRequestWithOtherSourcesAfterTimeout + reportSignatureRequest
	maxNumOfMsgsPerRound = 3

	// values below chosen by research team
	rateScale     = 1.2
	capacityScale = 3

	// bufferSize should be set to 1 as advised by the RMN team.
	outgoingBufferSize = 1
	incomingBufferSize = 1

	estimatedRoundInterval = time.Second
)

var (
	maxObservationRequestBytes int
	maxReportSigRequestBytes   int
)

func newStreamConfig(
	lggr logger.Logger,
	streamName string,
) networking.NewStreamArgs1 {
	cfg := networking.NewStreamArgs1{
		StreamName:         streamName,
		OutgoingBufferSize: outgoingBufferSize,
		IncomingBufferSize: incomingBufferSize,
		MaxMessageLength:   maxMessageLength(),
		MessagesLimit:      messagesLimit(),
		BytesLimit:         bytesLimit(),
	}

	lggr.Infow("new stream config",
		"streamName", streamName,
		"cfg", cfg,
		"maxObservationRequestBytes", maxObservationRequestBytes,
		"maxReportSigRequestBytes", maxReportSigRequestBytes,
	)

	return cfg
}

func maxMessageLength() int {
	return max(
		maxObservationRequestBytes,
		maxReportSigRequestBytes,
	)
}

func messagesLimit() ragep2p.TokenBucketParams {
	return ragep2p.TokenBucketParams{
		Rate:     rateScale * (float64(maxNumOfMsgsPerRound) / estimatedRoundInterval.Seconds()),
		Capacity: maxNumOfMsgsPerRound * capacityScale,
	}
}

func bytesLimit() ragep2p.TokenBucketParams {
	maxSumLenOutboundPerRound := (2 * maxObservationRequestBytes) + maxReportSigRequestBytes

	return ragep2p.TokenBucketParams{
		Rate:     (float64(maxSumLenOutboundPerRound) / estimatedRoundInterval.Seconds()) * rateScale,
		Capacity: uint32(maxSumLenOutboundPerRound) * capacityScale,
	}
}

// compute max observation request size and max report signatures request size
func init() {
	req := &rmnpb.Request{
		Request: &rmnpb.Request_ObservationRequest{
			ObservationRequest: &rmnpb.ObservationRequest{
				LaneDest: &rmnpb.LaneDest{
					DestChainSelector: math.MaxUint64,
					OfframpAddress:    make([]byte, 32),
				},
				FixedDestLaneUpdateRequests: make([]*rmnpb.FixedDestLaneUpdateRequest, 0, estimatedMaxNumberOfSourceChains),
			},
		},
	}
	for i := 0; i < estimatedMaxNumberOfSourceChains; i++ {
		req.GetObservationRequest().FixedDestLaneUpdateRequests = append(
			req.GetObservationRequest().FixedDestLaneUpdateRequests, &rmnpb.FixedDestLaneUpdateRequest{
				LaneSource: &rmnpb.LaneSource{
					SourceChainSelector: math.MaxUint64,
					OnrampAddress:       make([]byte, 32),
				},
				ClosedInterval: &rmnpb.ClosedInterval{
					MinMsgNr: math.MaxUint64,
					MaxMsgNr: math.MaxUint64,
				},
			},
		)
	}
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	maxObservationRequestBytes = len(reqBytes)

	req = &rmnpb.Request{
		Request: &rmnpb.Request_ReportSignatureRequest{
			ReportSignatureRequest: &rmnpb.ReportSignatureRequest{
				Context:                      &rmnpb.ReportContext{},
				AttributedSignedObservations: make([]*rmnpb.AttributedSignedObservation, 0, estimatedMaxNumberOfSourceChains),
			},
		},
	}
	fixedDestLaneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0, estimatedMaxNumberOfSourceChains)
	for i := 0; i < estimatedMaxNumberOfSourceChains; i++ {
		fixedDestLaneUpdates = append(fixedDestLaneUpdates, &rmnpb.FixedDestLaneUpdate{
			LaneSource: &rmnpb.LaneSource{
				SourceChainSelector: math.MaxUint64,
				OnrampAddress:       make([]byte, 32),
			},
			ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: math.MaxUint64, MaxMsgNr: math.MaxUint64},
			Root:           make([]byte, 32),
		})
	}

	for i := 0; i < estimatedMaxNumberOfSourceChains; i++ {
		req.GetReportSignatureRequest().AttributedSignedObservations = append(
			req.GetReportSignatureRequest().AttributedSignedObservations, &rmnpb.AttributedSignedObservation{
				SignedObservation: &rmnpb.SignedObservation{
					Observation: &rmnpb.Observation{
						RmnHomeContractConfigDigest: make([]byte, 32),
						LaneDest: &rmnpb.LaneDest{
							DestChainSelector: math.MaxUint64,
							OfframpAddress:    make([]byte, 32),
						},
						FixedDestLaneUpdates: fixedDestLaneUpdates,
						Timestamp:            math.MaxUint64,
					},
					Signature: make([]byte, 256),
				},
				SignerNodeIndex: math.MaxUint32,
			})
	}

	reqBytes, err = proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	maxReportSigRequestBytes = len(reqBytes)
}
