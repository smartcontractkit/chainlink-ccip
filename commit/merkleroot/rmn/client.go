package rmn

import (
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"google.golang.org/protobuf/proto"

	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

// ErrTimeout is returned when the signature computation times out.
var ErrTimeout = errors.New("signature computation timeout")

// Client contains the methods required by the plugin to interact with the RMN nodes.
type Client interface {
	// ComputeReportSignatures computes and returns the signatures for the provided lane updates.
	//
	// This method abstracts away the RMN specific requests (ObservationRequest, ReportSignaturesRequest) and all the
	// necessary steps to compute the signatures, like retrying and caching which are handled by the implementation.
	ComputeReportSignatures(
		ctx context.Context,
		destChain *rmnpb.LaneDest,
		requestedUpdates []*rmnpb.FixedDestLaneUpdateRequest,
	) (*ReportSignatures, error)
}

type ReportSignatures struct {
	// ReportSignatures are the ECDSA signatures for the lane updates for each node.
	Signatures []*rmnpb.EcdsaSignature
}

// PBClient is the base RMN Client implementation.
type PBClient struct {
	rawRmnClient        rawRmnClient
	rmnNodes            []RMNNodeInfo
	rmnRemoteAddress    []byte
	rmnHomeConfigDigest []byte
	minObservers        int
	minSigners          int

	lggr logger.Logger

	observationsInitialRequestTimeout time.Duration
	reportsInitialRequestTimeout      time.Duration
}

type RMNNodeInfo struct {
	ID                    uint32 // ID is the index of this node in the RMN config
	SupportedSourceChains mapset.Set[uint64]
	IsSigner              bool
}

type laneUpdateRequestWithStatus struct {
	request *rmnpb.FixedDestLaneUpdateRequest
	sent    bool
}

func (c *PBClient) ComputeReportSignatures(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	requestedUpdates []*rmnpb.FixedDestLaneUpdateRequest,
) (*ReportSignatures, error) {
	// In this implementation we assume that there is a different source chain for each lane update request.
	chainsSet := mapset.NewSet[uint64]()
	for _, req := range requestedUpdates {
		if chainsSet.Contains(req.LaneSource.SourceChainSelector) {
			return nil, fmt.Errorf("source chain %d is duplicated", req.LaneSource.SourceChainSelector)
		}
		chainsSet.Add(req.LaneSource.SourceChainSelector)
	}

	// We are given a list of lane update requests.
	// e.g.:
	//    sourceChain1: [10, 20]
	//    sourceChain2: [30, 40]

	// We have a list of RMN nodes some of them supporting sourceChain1 and some of them supporting sourceChain2.
	// Group all the lane update requests for each RMN node.
	// e.g.:
	//    rmnNode1: [sourceChain1: [10, 20], sourceChain2: [30, 40]]
	//	  rmnNode2: [sourceChain1: [10, 20]]
	//	  rmnNode3: [sourceChain2: [30, 40]]
	nodeLaneUpdateRequests := make(map[uint32][]laneUpdateRequestWithStatus)
	for _, node := range c.rmnNodes {
		for _, req := range requestedUpdates {
			if node.SupportedSourceChains.Contains(req.LaneSource.SourceChainSelector) {
				continue
			}
			nodeLaneUpdateRequests[node.ID] = append(nodeLaneUpdateRequests[node.ID], laneUpdateRequestWithStatus{
				request: req,
				sent:    false,
			})
		}
	}

	// For each lane update request send an observation request to at most 'minObservers' number of rmn nodes.
	requestCounts := make(map[uint64]int) // source chain -> count
	requestIDs := mapset.NewSet[uint64]()

	for _, node := range copyAndGetInRandomOrder(c.rmnNodes) {
		observationRequest := &rmnpb.ObservationRequest{
			LaneDest:                    destChain,
			FixedDestLaneUpdateRequests: make([]*rmnpb.FixedDestLaneUpdateRequest, 0),
		}

		for i, req := range nodeLaneUpdateRequests[node.ID] {
			if requestCounts[req.request.LaneSource.SourceChainSelector] >= c.minObservers {
				continue
			}
			requestCounts[req.request.LaneSource.SourceChainSelector]++
			observationRequest.FixedDestLaneUpdateRequests = append(
				observationRequest.FixedDestLaneUpdateRequests,
				req.request,
			)
			nodeLaneUpdateRequests[node.ID][i].sent = true
		}

		if len(observationRequest.FixedDestLaneUpdateRequests) == 0 {
			continue
		}

		req := &rmnpb.Request{
			RequestId: newRequestID(),
			Request: &rmnpb.Request_ObservationRequest{
				ObservationRequest: observationRequest,
			},
		}

		reqBytes, err := proto.Marshal(req)
		if err != nil {
			return nil, fmt.Errorf("proto marshal: %w", err)
		}
		if err := c.rawRmnClient.Send(node.ID, reqBytes); err != nil {
			return nil, fmt.Errorf("send rmn request: %w", err)
		}

		if requestIDs.Contains(req.RequestId) {
			return nil, fmt.Errorf("request id %d is duplicated, newRequestID fn bug", req.RequestId)
		}
		requestIDs.Add(req.RequestId)
	}

	// So far we have submitted the observation requests to the RMN nodes.
	// e.g.:
	//    rmnNode1: // nothing sent
	//	  rmnNode2: [sourceChain1: [10, 20]] // requestID1
	//	  rmnNode3: [sourceChain2: [30, 40]] // requestID2

	respChan := c.rawRmnClient.Recv()

	// Wait for the responses
	tInitialRequest := time.NewTimer(c.observationsInitialRequestTimeout)
	defer tInitialRequest.Stop()

	type rmnObservationResponse struct {
		RMNNodeIDX uint32
		SO         *rmnpb.SignedObservation
	}
	rmnObservationResponses := make([]rmnObservationResponse, 0)
	finishedRequests := mapset.NewSet[uint64]()

observationsLoop:
	for {
		select {
		case resp := <-respChan:
			responseTyp := &rmnpb.Response{}
			err := proto.Unmarshal(resp.Body, responseTyp)
			if err != nil {
				return nil, fmt.Errorf("proto unmarshal: %w", err)
			}

			if !requestIDs.Contains(responseTyp.RequestId) {
				continue // not something we are waiting for
			}
			if finishedRequests.Contains(responseTyp.RequestId) {
				c.lggr.Warnw("got an RMN duplicate response", "request_id", responseTyp.RequestId)
				continue // already got this response
			}

			finishedRequests.Add(responseTyp.RequestId)

			signedObs := responseTyp.GetSignedObservation()
			if signedObs == nil {
				c.lggr.Errorf("RMN node returned an unexpected type of response: %+v", responseTyp.Response)
			} else {
				rmnObservationResponses = append(rmnObservationResponses, rmnObservationResponse{
					RMNNodeIDX: resp.RMNNodeIDX,
					SO:         signedObs,
				})
			}

			// We got all the responses we were waiting for.
			if finishedRequests.Equal(requestIDs) {
				break
			}

			// We got sufficient number of responses for every source chain.
			cnts := make(map[uint64]int)
			for _, rlu := range requestedUpdates {
				cnts[rlu.LaneSource.SourceChainSelector] = 0
			}
			for _, o := range rmnObservationResponses {
				for _, obsUpdates := range o.SO.Observation.FixedDestLaneUpdates {
					cnts[obsUpdates.LaneSource.SourceChainSelector]++
				}
			}
			allSufficient := true
			for _, rlu := range requestedUpdates {
				if cnts[rlu.LaneSource.SourceChainSelector] < c.minObservers {
					allSufficient = false
					break
				}
			}
			if allSufficient {
				break observationsLoop
			}
		case <-tInitialRequest.C:
			for _, node := range copyAndGetInRandomOrder(c.rmnNodes) {
				observationRequest := &rmnpb.ObservationRequest{
					LaneDest:                    destChain,
					FixedDestLaneUpdateRequests: make([]*rmnpb.FixedDestLaneUpdateRequest, 0),
				}
				for i, req := range nodeLaneUpdateRequests[node.ID] {
					if nodeLaneUpdateRequests[node.ID][i].sent {
						continue
					}
					observationRequest.FixedDestLaneUpdateRequests = append(
						observationRequest.FixedDestLaneUpdateRequests,
						req.request,
					)
					nodeLaneUpdateRequests[node.ID][i].sent = true
				}
				if len(observationRequest.FixedDestLaneUpdateRequests) == 0 {
					continue
				}
				req := &rmnpb.Request{
					RequestId: newRequestID(),
					Request: &rmnpb.Request_ObservationRequest{
						ObservationRequest: observationRequest,
					},
				}
				if requestIDs.Contains(req.RequestId) {
					return nil, fmt.Errorf("request id %d is duplicated, newRequestID fn bug", req.RequestId)
				}
				requestIDs.Add(req.RequestId)
				reqBytes, err := proto.Marshal(req)
				if err != nil {
					return nil, fmt.Errorf("proto marshal: %w", err)
				}
				if err := c.rawRmnClient.Send(node.ID, reqBytes); err != nil {
					return nil, fmt.Errorf("send rmn request: %w", err)
				}
			}
		case <-ctx.Done():
			return nil, ErrTimeout
		}
	}

	// At this point we got all the observations we were waiting for but maybe not a sufficient number 'minObservers'
	// for each source chain.
	//
	// e.g.:
	//    rmnNode1: [sourceChain1: [10, 20], sourceChain2: [30, 40]] // someSignature1
	//	  rmnNode2: [sourceChain1: [10, 20]]                         // someSignature2

	attributedSignedObs := make([]*rmnpb.AttributedSignedObservation, 0)
	for _, resp := range rmnObservationResponses {
		sort.Slice(resp.SO.Observation.FixedDestLaneUpdates, func(i, j int) bool {
			return resp.SO.Observation.FixedDestLaneUpdates[i].LaneSource.SourceChainSelector <
				resp.SO.Observation.FixedDestLaneUpdates[j].LaneSource.SourceChainSelector
		})
		attributedSignedObs = append(attributedSignedObs, &rmnpb.AttributedSignedObservation{
			SignedObservation: resp.SO,
			SignerNodeIndex:   resp.RMNNodeIDX,
		})
	}
	sort.Slice(attributedSignedObs, func(i, j int) bool {
		return attributedSignedObs[i].SignerNodeIndex < attributedSignedObs[j].SignerNodeIndex
	})

	chainInfo, exists := chainsel.ChainBySelector(destChain.DestChainSelector)
	if !exists {
		return nil, fmt.Errorf("unknown chain selector %d", destChain.DestChainSelector)
	}

	reportSigReq := rmnpb.ReportSignatureRequest{
		Context: &rmnpb.ReportContext{
			EvmDestChainId:              chainInfo.EvmChainID,
			RmnRemoteContractAddress:    c.rmnRemoteAddress,
			RmnHomeContractConfigDigest: c.rmnHomeConfigDigest,
			LaneDest:                    destChain,
		},
		AttributedSignedObservations: attributedSignedObs,
	}

	requestIDs = mapset.NewSet[uint64]()
	signersRequested := mapset.NewSet[uint32]()

	// Send the report signature request to at least minSigners
	for _, node := range copyAndGetInRandomOrder(c.rmnNodes) {
		if !node.IsSigner {
			continue
		}

		req := &rmnpb.Request{
			RequestId: newRequestID(),
			Request: &rmnpb.Request_ReportSignatureRequest{
				ReportSignatureRequest: &reportSigReq,
			},
		}

		reqBytes, err := proto.Marshal(req)
		if err != nil {
			return nil, fmt.Errorf("proto marshal: %w", err)
		}
		if err := c.rawRmnClient.Send(node.ID, reqBytes); err != nil {
			return nil, fmt.Errorf("send rmn report signature request: %w", err)
		}

		if requestIDs.Contains(req.RequestId) {
			return nil, fmt.Errorf("request id %d is duplicated, newRequestID fn bug", req.RequestId)
		}
		requestIDs.Add(req.RequestId)
		signersRequested.Add(node.ID)

		if requestIDs.Cardinality() >= c.minSigners {
			break
		}
	}

	tReportsInitialRequest := time.NewTimer(c.reportsInitialRequestTimeout)
	reportSigs := make([]*rmnpb.ReportSignature, 0)
	finishedRequests = mapset.NewSet[uint64]()

	for {
		select {
		case resp := <-respChan:
			responseTyp := &rmnpb.Response{}
			err := proto.Unmarshal(resp.Body, responseTyp)
			if err != nil {
				return nil, fmt.Errorf("proto unmarshal: %w", err)
			}

			if !requestIDs.Contains(responseTyp.RequestId) {
				continue // not something we are waiting for
			}
			if finishedRequests.Contains(responseTyp.RequestId) {
				c.lggr.Warnw("got an RMN duplicate response", "request_id", responseTyp.RequestId)
				continue // already got this response
			}

			reportSig := responseTyp.GetReportSignature()
			if reportSig == nil {
				c.lggr.Errorf("RMN node returned an unexpected type of response: %+v", responseTyp.Response)
			} else {
				reportSigs = append(reportSigs, reportSig)
			}

			if len(reportSigs) >= c.minSigners {
				ecdsaSigs := make([]*rmnpb.EcdsaSignature, 0)
				for _, rs := range reportSigs {
					ecdsaSigs = append(ecdsaSigs, rs.Signature)
				}

				return &ReportSignatures{
					Signatures: ecdsaSigs,
				}, nil
			}
		case <-tReportsInitialRequest.C:
			// Send to the rest of the signers
			for _, node := range copyAndGetInRandomOrder(c.rmnNodes) {
				if !node.IsSigner || signersRequested.Contains(node.ID) {
					continue
				}

				req := &rmnpb.Request{
					RequestId: newRequestID(),
					Request: &rmnpb.Request_ReportSignatureRequest{
						ReportSignatureRequest: &reportSigReq,
					},
				}
				reqBytes, err := proto.Marshal(req)
				if err != nil {
					return nil, fmt.Errorf("proto marshal: %w", err)
				}
				if err := c.rawRmnClient.Send(node.ID, reqBytes); err != nil {
					return nil, fmt.Errorf("send rmn report signature request: %w", err)
				}

				if requestIDs.Contains(req.RequestId) {
					return nil, fmt.Errorf("request id %d is duplicated, newRequestID fn bug", req.RequestId)
				}
				requestIDs.Add(req.RequestId)
				signersRequested.Add(node.ID)
			}
		case <-ctx.Done():
			return nil, ErrTimeout
		}
	}
}

type rawRmnClient interface {
	Send(rmnNodeID uint32, request []byte) error
	// Recv returns a channel which can be used to listen on for
	// responses by all RMN nodes. This is expected to be monitored
	// by the plugin in order to get RMN responses.
	Recv() <-chan rawRmnResponse
}

type rawRmnResponse struct {
	RMNNodeIDX uint32
	Body       []byte // pb
}

func copyAndGetInRandomOrder[T any](s []T) []T {
	copied := make([]T, len(s))
	copy(copied, s)
	rand.Shuffle(len(copied), func(i, j int) {
		copied[i], copied[j] = copied[j], copied[i]
	})
	return copied
}

func newRequestID() uint64 {
	now := uint64(time.Now().UTC().UnixNano())

	// Generate a random 32-bit number
	var randBytes [4]byte
	_, err := crand.Read(randBytes[:])
	if err != nil {
		panic("failed to generate a random number")
	}

	randPart := uint64(binary.LittleEndian.Uint32(randBytes[:]))
	return (now << 32) | randPart
}
