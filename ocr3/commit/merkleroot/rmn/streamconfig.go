package rmn

import (
	"math"
	"time"

	"github.com/smartcontractkit/libocr/networking"
	"github.com/smartcontractkit/libocr/ragep2p"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"
)

const (
	// estimatedMaxNumberOfSourceChains is the estimated maximum number of source chains
	// that the current stream configuration supports and can be increased if required.
	// This value does not correlate to the maximum number of source chains that CCIP can support.
	estimatedMaxNumberOfSourceChains = 500

	// initialObservationRequest + observationRequestWithOtherSourcesAfterTimeout + reportSignatureRequest
	maxNumOfMsgsPerRound = 3

	rateScale       = 2
	capacityScale   = 5
	minMsgLimitRate = 1.0

	// bufferSize should be set to 1 as advised by the RMN team.
	outgoingBufferSize = 1
	incomingBufferSize = 1
)

var (
	maxObservationResponseBytes int
	maxReportSigResponseBytes   int
)

func newStreamConfig(
	lggr logger.Logger,
	streamName string,
	estimatedOcrRoundInterval time.Duration,
) networking.NewStreamArgs1 {
	cfg := networking.NewStreamArgs1{
		StreamName:         streamName,
		OutgoingBufferSize: outgoingBufferSize,
		IncomingBufferSize: incomingBufferSize,
		MaxMessageLength:   maxMessageLength(),
		MessagesLimit:      messagesLimit(estimatedOcrRoundInterval),
		BytesLimit:         bytesLimit(estimatedOcrRoundInterval),
	}

	lggr.Infow("new stream config",
		"streamName", streamName,
		"cfg", cfg,
		"maxObservationResponseBytes", maxObservationResponseBytes,
		"maxReportSigResponseBytes", maxReportSigResponseBytes,
	)

	return cfg
}

func maxMessageLength() int {
	return max(
		maxObservationResponseBytes,
		maxReportSigResponseBytes,
	)
}

func messagesLimit(roundInterval time.Duration) ragep2p.TokenBucketParams {
	numMsgsPerRound := float64(maxNumOfMsgsPerRound)
	roundIntervalSecF64 := float64(roundInterval.Milliseconds()) / 1000.0

	rate := (numMsgsPerRound / roundIntervalSecF64) * rateScale
	capacity := numMsgsPerRound * capacityScale

	if rate < minMsgLimitRate {
		rate = minMsgLimitRate
	}

	return ragep2p.TokenBucketParams{
		Rate:     rate,
		Capacity: uint32(capacity),
	}
}

func bytesLimit(roundInterval time.Duration) ragep2p.TokenBucketParams {
	maxSumLenInboundPerRound := maxObservationResponseBytes + maxReportSigResponseBytes
	roundIntervalSecF64 := float64(roundInterval.Milliseconds()) / 1000.0

	rate := (float64(maxSumLenInboundPerRound) / roundIntervalSecF64) * rateScale
	capacity := uint32(maxSumLenInboundPerRound * capacityScale)

	return ragep2p.TokenBucketParams{
		Rate:     rate,
		Capacity: capacity,
	}
}

// compute max observation request size and max report signatures request size
func init() {
	fixedDestLaneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0, estimatedMaxNumberOfSourceChains)
	for i := 0; i < estimatedMaxNumberOfSourceChains; i++ {
		fixedDestLaneUpdates = append(fixedDestLaneUpdates, &rmnpb.FixedDestLaneUpdate{
			LaneSource: &rmnpb.LaneSource{
				SourceChainSelector: math.MaxUint64,
				OnrampAddress:       make([]byte, 32),
			},
			ClosedInterval: &rmnpb.ClosedInterval{
				MinMsgNr: math.MaxUint64,
				MaxMsgNr: math.MaxUint64,
			},
			Root: make([]byte, 32),
		})
	}

	obsResponse := &rmnpb.Response{
		RequestId: math.MaxUint64,
		Response: &rmnpb.Response_SignedObservation{
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
		},
	}

	sigResponse := &rmnpb.Response{
		RequestId: math.MaxUint64,
		Response: &rmnpb.Response_ReportSignature{
			ReportSignature: &rmnpb.ReportSignature{
				Signature: &rmnpb.EcdsaSignature{
					R: make([]byte, 32),
					S: make([]byte, 32),
				},
			},
		},
	}

	obsResponseBytes, err := proto.Marshal(obsResponse)
	if err != nil {
		panic(err)
	}

	sigResponseBytes, err := proto.Marshal(sigResponse)
	if err != nil {
		panic(err)
	}

	maxObservationResponseBytes = len(obsResponseBytes)
	maxReportSigResponseBytes = len(sigResponseBytes)
}
