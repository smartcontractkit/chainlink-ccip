package rmn

import (
	"bytes"
	"context"
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

// TODO: testing - unit tests, integration tests, etc...
// TODO: use the populated Query / sig verification / etc..
// TODO: rmn config watcher

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
	// LaneUpdates are the lane updates for each source chain that we got the signatures for.
	// NOTE: A signature[i] corresponds to the whole LaneUpdates slice and NOT LaneUpdates[i].
	LaneUpdates []*rmnpb.FixedDestLaneUpdate
}

// PBClient is the base RMN Client implementation.
type PBClient struct {
	lggr                                    logger.Logger
	rawRmnClient                            RawRmnClient
	rmnNodes                                []RMNNodeInfo
	rmnRemoteAddress                        []byte
	rmnHomeConfigDigest                     []byte
	minObservers                            int
	minSigners                              int
	observationsInitialRequestTimerDuration time.Duration
	reportsInitialRequestTimerDuration      time.Duration
}

// RMNNodeInfo contains the information about an RMN node.
type RMNNodeInfo struct {
	ID                    uint32 // ID is the index of this node in the RMN config
	SupportedSourceChains mapset.Set[uint64]
	IsSigner              bool
}

// NewPBClient creates a new RMN Client to be used by the plugin.
func NewPBClient(
	lggr logger.Logger,
	rawRmnClient RawRmnClient,
	rmnNodes []RMNNodeInfo,
	rmnRemoteAddress []byte,
	rmnHomeConfigDigest []byte,
	minObservers int,
	minSigners int,
	observationsInitialRequestTimerDuration time.Duration,
	reportsInitialRequestTimerDuration time.Duration,
) *PBClient {
	return &PBClient{
		lggr:                                    lggr,
		rawRmnClient:                            rawRmnClient,
		rmnNodes:                                rmnNodes,
		rmnRemoteAddress:                        rmnRemoteAddress,
		rmnHomeConfigDigest:                     rmnHomeConfigDigest,
		minObservers:                            minObservers,
		minSigners:                              minSigners,
		observationsInitialRequestTimerDuration: observationsInitialRequestTimerDuration,
		reportsInitialRequestTimerDuration:      reportsInitialRequestTimerDuration,
	}
}

func (c *PBClient) ComputeReportSignatures(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	requestedUpdates []*rmnpb.FixedDestLaneUpdateRequest,
) (*ReportSignatures, error) {
	// Group the lane update requests by their source chain.
	lursPerChain := make(map[uint64]lurWithMeta)
	for _, updateReq := range requestedUpdates {
		if _, exists := lursPerChain[updateReq.LaneSource.SourceChainSelector]; exists {
			return nil, errors.New("this Client implementation assumes each lane update is for a different chain")
		}

		lursPerChain[updateReq.LaneSource.SourceChainSelector] = lurWithMeta{
			lur:      updateReq,
			rmnNodes: mapset.NewSet[uint32](),
		}
		for _, node := range c.rmnNodes {
			if node.SupportedSourceChains.Contains(updateReq.LaneSource.SourceChainSelector) {
				lursPerChain[updateReq.LaneSource.SourceChainSelector].rmnNodes.Add(node.ID)
			}
		}
	}

	// Filter out the lane update requests for chains without enough RMN nodes supporting them.
	for chain, l := range lursPerChain {
		if l.rmnNodes.Cardinality() < c.minObservers {
			delete(lursPerChain, chain)
		}
	}

	if len(lursPerChain) == 0 {
		return nil, errors.New("no source chains with enough RMN nodes, nothing to do")
	}

	rmnSignedObservations, err := c.getRmnSignedObservations(ctx, destChain, lursPerChain)
	if err != nil {
		return nil, fmt.Errorf("get rmn signed observations: %w", err)
	}

	rmnReportSignatures, err := c.getRmnReportSignatures(ctx, destChain, rmnSignedObservations)
	if err != nil {
		return nil, fmt.Errorf("get rmn report signatures: %w", err)
	}

	return rmnReportSignatures, nil
}

func (c *PBClient) getRmnSignedObservations(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	lursPerChain map[uint64]lurWithMeta,
) ([]rmnSignedObservationWithMeta, error) {
	requestedNodes := make(map[uint64]mapset.Set[uint32])                   // sourceChain -> requested rmnNodeIDs
	requestsPerNode := make(map[uint32][]*rmnpb.FixedDestLaneUpdateRequest) // grouped requests for each node

	// For each lane update request send an observation request to at most 'minObservers' number of rmn nodes.
	for sourceChain, lur := range lursPerChain {
		requestedNodes[sourceChain] = mapset.NewSet[uint32]()

		// At this point we assume at least minObservers RMN nodes are available for each source chain.
		for nodeID := range lur.rmnNodes.Iter() {
			if requestedNodes[sourceChain].Cardinality() >= c.minObservers {
				break
			}

			requestedNodes[sourceChain].Add(nodeID)
			if _, exists := requestsPerNode[nodeID]; !exists {
				requestsPerNode[nodeID] = make([]*rmnpb.FixedDestLaneUpdateRequest, 0)
			}
			requestsPerNode[nodeID] = append(requestsPerNode[nodeID], lur.lur)
		}
	}

	requestIDs := c.sendObservationRequests(destChain, requestsPerNode)
	signedObservations, err := c.listenForRmnObservationResponses(ctx, destChain, requestIDs, lursPerChain, requestedNodes)
	if err != nil {
		return nil, fmt.Errorf("listen for rmn observation responses: %w", err)
	}

	// Sanity check that we got enough signed observations, should never happen, an error must have been received
	// in the previous step if we didn't get enough observations for each source chain.
	if err := c.ensureEnoughSignedObservations(signedObservations); err != nil {
		return nil, fmt.Errorf("unexpected error, ensuring enough signed observations: %w", err)
	}

	return signedObservations, nil
}

func (c *PBClient) sendObservationRequests(
	destChain *rmnpb.LaneDest,
	requestsPerNode map[uint32][]*rmnpb.FixedDestLaneUpdateRequest,
) (requestIDs mapset.Set[uint64]) {
	requestIDs = mapset.NewSet[uint64]()

	for nodeID, lurs := range requestsPerNode {
		req := &rmnpb.Request{
			RequestId: newRequestID(),
			Request: &rmnpb.Request_ObservationRequest{
				ObservationRequest: &rmnpb.ObservationRequest{
					LaneDest:                    destChain,
					FixedDestLaneUpdateRequests: lurs,
				},
			},
		}

		if err := c.marshalAndSend(req, nodeID); err != nil {
			c.lggr.Errorw("failed to send observation request", "node_id", nodeID, "err", err)
			continue
		}
		requestIDs.Add(req.RequestId)
	}

	return requestIDs
}

func (c *PBClient) listenForRmnObservationResponses(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	requestIDs mapset.Set[uint64],
	lursPerChain map[uint64]lurWithMeta,
	requestedNodes map[uint64]mapset.Set[uint32],
) ([]rmnSignedObservationWithMeta, error) {
	respChan := c.rawRmnClient.Recv()

	finishedRequestIDs := mapset.NewSet[uint64]()
	rmnObservationResponses := make([]rmnSignedObservationWithMeta, 0)
	chainObservationsCount := make(map[uint64]int)

	initialObservationRequestTimer := time.NewTimer(c.observationsInitialRequestTimerDuration)
	defer initialObservationRequestTimer.Stop()
	for {
		select {
		case resp := <-respChan:
			parsedResp, err := c.parseResponse(&resp, requestIDs, finishedRequestIDs)
			if err != nil {
				c.lggr.Warnw("failed to parse RMN response", "err", err)
				continue
			}
			finishedRequestIDs.Add(parsedResp.RequestId)

			signedObs := parsedResp.GetSignedObservation()
			if signedObs == nil {
				c.lggr.Errorf("RMN node returned an unexpected type of response: %+v", parsedResp.Response)
			} else {
				err := c.validateSignedObservationResponse(resp.RMNNodeID, lursPerChain, signedObs, destChain)
				if err != nil {
					c.lggr.Errorw("failed to validate signed observation response", "err", err)
					continue
				}
				rmnObservationResponses = append(rmnObservationResponses, rmnSignedObservationWithMeta{
					RMNNodeID:         resp.RMNNodeID,
					SignedObservation: signedObs,
				})
				for _, lu := range signedObs.Observation.FixedDestLaneUpdates {
					chainObservationsCount[lu.LaneSource.SourceChainSelector]++
				}
			}

			// We got all the responses we were waiting for.
			if finishedRequestIDs.Equal(requestIDs) {
				return rmnObservationResponses, nil
			}

			// We got sufficient number of observation responses for every source chain.
			allChainsHaveEnoughResponses := true
			for _, count := range chainObservationsCount {
				if count < c.minObservers {
					allChainsHaveEnoughResponses = false
					break
				}
			}
			if allChainsHaveEnoughResponses {
				return rmnObservationResponses, nil
			}
		case <-initialObservationRequestTimer.C:
			// Timer expired, send the observation requests to the rest of the RMN nodes.
			requestsPerNode := make(map[uint32][]*rmnpb.FixedDestLaneUpdateRequest)
			for sourceChain, lur := range lursPerChain {
				for nodeID := range lur.rmnNodes.Iter() {
					if requestedNodes[sourceChain].Contains(nodeID) {
						continue
					}
					requestedNodes[sourceChain].Add(nodeID)
					if _, ok := requestsPerNode[nodeID]; !ok {
						requestsPerNode[nodeID] = make([]*rmnpb.FixedDestLaneUpdateRequest, 0)
					}
					requestsPerNode[nodeID] = append(requestsPerNode[nodeID], lur.lur)
				}
			}
			newRequestIDs := c.sendObservationRequests(destChain, requestsPerNode)
			requestIDs = requestIDs.Union(newRequestIDs)
		case <-ctx.Done():
			return nil, ErrTimeout
		}
	}
}

func (c *PBClient) validateSignedObservationResponse(
	rmnNodeID uint32,
	lurs map[uint64]lurWithMeta,
	signedObs *rmnpb.SignedObservation,
	destChain *rmnpb.LaneDest,
) error {
	if signedObs.Observation.LaneDest != destChain {
		return fmt.Errorf("unexpected lane dest %v", signedObs.Observation.LaneDest)
	}

	if !bytes.Equal(signedObs.Observation.RmnHomeContractConfigDigest, c.rmnHomeConfigDigest) {
		return fmt.Errorf("unexpected rmn home contract config digest %v",
			signedObs.Observation.RmnHomeContractConfigDigest)
	}

	for _, signedObsLu := range signedObs.Observation.FixedDestLaneUpdates {
		lur, exists := lurs[signedObsLu.LaneSource.SourceChainSelector]
		if !exists {
			return fmt.Errorf("unexpected source chain selector %d", signedObsLu.LaneSource.SourceChainSelector)
		}

		if !lur.rmnNodes.Contains(rmnNodeID) {
			return fmt.Errorf("rmn node %d not expected to read chain %d",
				rmnNodeID, signedObsLu.LaneSource.SourceChainSelector)
		}

		if lur.lur.ClosedInterval != signedObsLu.ClosedInterval {
			return fmt.Errorf("unexpected closed interval %v", signedObsLu.ClosedInterval)
		}
		if lur.lur.LaneSource != signedObsLu.LaneSource {
			return fmt.Errorf("unexpected lane source %v", signedObsLu.LaneSource)
		}
		if signedObsLu.Root == nil {
			return errors.New("root is nil")
		}
	}
	return nil
}

func (c *PBClient) getRmnReportSignatures(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	rmnSignedObservations []rmnSignedObservationWithMeta,
) (*ReportSignatures, error) {
	// At this point we might have multiple signedObservations for different nodes but never for the same source chain
	// from the same node.
	//
	// e.g.
	// The following nodes support the following chains and minObservers=2:
	// node1: [1]	node2:[1,2,3]	node3:[1,2,3]
	//
	// node1: getObservations(1)          ->             never_responds
	// node2: getObservations(1,2,3)      -> responds_ok
	// node3: getObservations(2,3)        ->     responds_ok
	// node3: getObservations(1)          ->                         responds_ok (after timeout we made a new request)

	sigObservations := make([]*rmnpb.AttributedSignedObservation, 0)
	for _, resp := range rmnSignedObservations {
		// order observations by source chain
		sort.Slice(resp.SignedObservation.Observation.FixedDestLaneUpdates, func(i, j int) bool {
			return resp.SignedObservation.Observation.FixedDestLaneUpdates[i].LaneSource.SourceChainSelector <
				resp.SignedObservation.Observation.FixedDestLaneUpdates[j].LaneSource.SourceChainSelector
		})
		sigObservations = append(sigObservations, &rmnpb.AttributedSignedObservation{
			SignedObservation: resp.SignedObservation,
			SignerNodeIndex:   resp.RMNNodeID,
		})
	}

	// order by node index and then by source chain selector of the first observation request (should exist)
	sort.Slice(sigObservations, func(i, j int) bool {
		if sigObservations[i].SignerNodeIndex == sigObservations[j].SignerNodeIndex {
			return sigObservations[i].SignedObservation.Observation.FixedDestLaneUpdates[0].LaneSource.SourceChainSelector <
				sigObservations[j].SignedObservation.Observation.FixedDestLaneUpdates[0].LaneSource.SourceChainSelector
		}
		return sigObservations[i].SignerNodeIndex < sigObservations[j].SignerNodeIndex
	})

	chainInfo, exists := chainsel.ChainBySelector(destChain.DestChainSelector)
	if !exists {
		return nil, fmt.Errorf("unknown chain selector %d", destChain.DestChainSelector)
	}

	reportSigReq := &rmnpb.ReportSignatureRequest{
		Context: &rmnpb.ReportContext{
			EvmDestChainId:              chainInfo.EvmChainID,
			RmnRemoteContractAddress:    c.rmnRemoteAddress,
			RmnHomeContractConfigDigest: c.rmnHomeConfigDigest,
			LaneDest:                    destChain,
		},
		AttributedSignedObservations: sigObservations,
	}

	requestIDs := mapset.NewSet[uint64]()
	signersRequested := mapset.NewSet[uint32]()

	// Send the report signature request to at least minSigners
	for _, node := range c.rmnNodes {
		if requestIDs.Cardinality() >= c.minSigners {
			break
		}

		if !node.IsSigner {
			continue
		}

		req := &rmnpb.Request{
			RequestId: newRequestID(),
			Request: &rmnpb.Request_ReportSignatureRequest{
				ReportSignatureRequest: reportSigReq,
			},
		}

		err := c.marshalAndSend(req, node.ID)
		if err != nil {
			return nil, fmt.Errorf("send rmn report signature request: %w", err)
		}
		requestIDs.Add(req.RequestId)
		signersRequested.Add(node.ID)
	}

	ecdsaSignatures, err := c.listenForRmnReportSignatures(ctx, requestIDs, reportSigReq, signersRequested)
	if err != nil {
		return nil, fmt.Errorf("listen for rmn report signatures: %w", err)
	}

	// At this point we expect that attributed observations are correct, but we might have different roots
	// coming from different RMN nodes. In that case we return an error since something is breached.
	rootsPerSourceChain := make(map[uint64][]byte)
	fixedDestLaneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0)
	for _, signedObs := range sigObservations {
		for _, lu := range signedObs.SignedObservation.Observation.FixedDestLaneUpdates {
			existingRoot, exists := rootsPerSourceChain[lu.LaneSource.SourceChainSelector]
			if !exists {
				rootsPerSourceChain[lu.LaneSource.SourceChainSelector] = lu.Root
				fixedDestLaneUpdates = append(fixedDestLaneUpdates, lu)
				continue
			}
			if !bytes.Equal(existingRoot, lu.Root) {
				return nil, fmt.Errorf("found different roots (%x, %x) for the same source chain %d",
					existingRoot, lu.Root, lu.LaneSource.SourceChainSelector)
			}
		}
	}

	return &ReportSignatures{
		Signatures:  ecdsaSignatures,
		LaneUpdates: fixedDestLaneUpdates,
	}, nil
}

func (c *PBClient) listenForRmnReportSignatures(
	ctx context.Context,
	requestIDs mapset.Set[uint64],
	reportSigReq *rmnpb.ReportSignatureRequest,
	signersRequested mapset.Set[uint32],
) ([]*rmnpb.EcdsaSignature, error) {
	tReportsInitialRequest := time.NewTimer(c.reportsInitialRequestTimerDuration)
	reportSigs := make([]*rmnpb.ReportSignature, 0)
	finishedRequests := mapset.NewSet[uint64]()
	respChan := c.rawRmnClient.Recv()

	for {
		select {
		case resp := <-respChan:
			responseTyp, err := c.parseResponse(&resp, requestIDs, finishedRequests)
			if err != nil {
				c.lggr.Errorw("failed to parse RMN response", "err", err)
				continue
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
				return ecdsaSigs, nil
			}
		case <-tReportsInitialRequest.C:
			// Send to the rest of the signers
			for node := range randomIter(c.rmnNodes) {
				if !node.IsSigner || signersRequested.Contains(node.ID) {
					continue
				}
				req := &rmnpb.Request{
					RequestId: newRequestID(),
					Request: &rmnpb.Request_ReportSignatureRequest{
						ReportSignatureRequest: reportSigReq,
					},
				}
				if err := c.marshalAndSend(req, node.ID); err != nil {
					c.lggr.Errorw("failed to send report signature request", "node_id", node.ID, "err", err)
					continue
				}
				requestIDs.Add(req.RequestId)
				signersRequested.Add(node.ID)
			}
		case <-ctx.Done():
			return nil, ErrTimeout
		}
	}
}

func (c *PBClient) ensureEnoughSignedObservations(rmnSignedObservations []rmnSignedObservationWithMeta) error {
	counts := make(map[uint64]int)
	for _, so := range rmnSignedObservations {
		for _, lu := range so.SignedObservation.Observation.FixedDestLaneUpdates {
			counts[lu.LaneSource.SourceChainSelector]++
		}
	}

	for chain, count := range counts {
		if count < c.minObservers {
			return fmt.Errorf("not enough observations for chain=%d count=%d minObservers=%d",
				chain, count, c.minObservers)
		}
	}

	return nil
}

type RawRmnClient interface {
	Send(rmnNodeID uint32, request []byte) error
	// Recv returns a channel which can be used to listen on for
	// responses by all RMN nodes. This is expected to be monitored
	// by the plugin in order to get RMN responses.
	Recv() <-chan RawRmnResponse
}

type RawRmnResponse struct {
	RMNNodeID uint32
	Body      []byte // pb
}

type lurWithMeta struct {
	lur      *rmnpb.FixedDestLaneUpdateRequest
	rmnNodes mapset.Set[uint32]
}

type rmnSignedObservationWithMeta struct {
	SignedObservation *rmnpb.SignedObservation
	RMNNodeID         uint32
}

func randomIter[T any](s []T) chan T {
	ch := make(chan T)
	go func() {
		defer close(ch)
		perm := rand.Perm(len(s))
		for _, i := range perm {
			ch <- s[i]
		}
	}()
	return ch
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
