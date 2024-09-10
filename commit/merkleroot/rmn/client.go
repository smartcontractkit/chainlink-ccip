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
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

// TODO: testing - more unit tests and some code cleanup.
// TODO: use the populated Query / sig verification / etc..
// TODO: rmn config watcher

var (
	// ErrTimeout is returned when the signature computation times out.
	ErrTimeout = errors.New("signature computation timeout")

	// ErrNothingToDo is returned when there are no source chains with enough RMN nodes.
	ErrNothingToDo = errors.New("nothing to observe from the existing RMN nodes, make " +
		"sure RMN is enabled, nodes configured correctly, minObservers value is correct")
)

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
	// Signatures are the ECDSA signatures for the lane updates for each node.
	Signatures []*rmnpb.EcdsaSignature
	// LaneUpdates are the lane updates for each source chain that we got the signatures for.
	// NOTE: A signature[i] corresponds to the whole LaneUpdates slice and NOT LaneUpdates[i].
	LaneUpdates []*rmnpb.FixedDestLaneUpdate
}

// client is the base RMN Client implementation.
type client struct {
	lggr                                    logger.Logger
	rawRmnClient                            RawRmnClient
	rmnNodes                                []RMNNodeInfo
	rmnRemoteAddress                        cciptypes.Bytes
	rmnHomeConfigDigest                     cciptypes.Bytes
	minObservers                            int
	minSigners                              int
	observationsInitialRequestTimerDuration time.Duration
	reportsInitialRequestTimerDuration      time.Duration
}

// RMNNodeInfo contains the information about an RMN node.
type RMNNodeInfo struct {
	// ID is the index of this node in the RMN config
	ID                    uint32
	SupportedSourceChains mapset.Set[cciptypes.ChainSelector]
	IsSigner              bool
}

// Newclient creates a new RMN Client to be used by the plugin.
func Newclient(
	lggr logger.Logger,
	rawRmnClient RawRmnClient,
	rmnNodes []RMNNodeInfo,
	rmnRemoteAddress cciptypes.Bytes,
	rmnHomeConfigDigest cciptypes.Bytes,
	minObservers int,
	minSigners int,
	observationsInitialRequestTimerDuration time.Duration,
	reportsInitialRequestTimerDuration time.Duration,
) Client {
	return &client{
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

// ComputeReportSignatures sends a request to each rmn node to handle requests and build signatures.
func (c *client) ComputeReportSignatures(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	updateRequests []*rmnpb.FixedDestLaneUpdateRequest,
) (*ReportSignatures, error) {
	// Group the lane update requests by their source chain.
	updatesPerChain := make(map[uint64]updateRequestWithMeta)
	for _, updateReq := range updateRequests {
		if _, exists := updatesPerChain[updateReq.LaneSource.SourceChainSelector]; exists {
			return nil, errors.New("this Client implementation assumes each lane update is for a different chain")
		}

		updatesPerChain[updateReq.LaneSource.SourceChainSelector] = updateRequestWithMeta{
			data:     updateReq,
			rmnNodes: mapset.NewSet[uint32](),
		}
		for _, node := range c.rmnNodes {
			if node.SupportedSourceChains.Contains(cciptypes.ChainSelector(updateReq.LaneSource.SourceChainSelector)) {
				updatesPerChain[updateReq.LaneSource.SourceChainSelector].rmnNodes.Add(node.ID)
			}
		}
	}

	// Filter out the lane update requests for chains without enough RMN nodes supporting them.
	for chain, l := range updatesPerChain {
		if l.rmnNodes.Cardinality() < c.minObservers {
			c.lggr.Warnw("not enough RMN nodes for chain", "chain", chain, "minObservers", c.minObservers)
			delete(updatesPerChain, chain)
		}
	}

	if len(updatesPerChain) == 0 {
		return nil, ErrNothingToDo
	}

	rmnSignedObservations, err := c.getRmnSignedObservations(ctx, destChain, updatesPerChain)
	if err != nil {
		return nil, fmt.Errorf("get rmn signed observations: %w", err)
	}

	rmnReportSignatures, err := c.getRmnReportSignatures(ctx, destChain, rmnSignedObservations)
	if err != nil {
		return nil, fmt.Errorf("get rmn report signatures: %w", err)
	}

	return rmnReportSignatures, nil
}

func (c *client) getRmnSignedObservations(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	updateRequestsPerChain map[uint64]updateRequestWithMeta,
) ([]rmnSignedObservationWithMeta, error) {
	requestedNodes := make(map[uint64]mapset.Set[uint32])                   // sourceChain -> requested rmnNodeIDs
	requestsPerNode := make(map[uint32][]*rmnpb.FixedDestLaneUpdateRequest) // grouped requests for each node

	c.lggr.Infof("update requests per chain: %v", updateRequestsPerChain)

	// For each lane update request send an observation request to at most 'minObservers' number of rmn nodes.
	for sourceChain, updateRequest := range updateRequestsPerChain {
		requestedNodes[sourceChain] = mapset.NewSet[uint32]()

		// At this point we assume at least minObservers RMN nodes are available for each source chain.
		for nodeID := range updateRequest.rmnNodes.Iter() {
			if requestedNodes[sourceChain].Cardinality() >= c.minObservers {
				break
			}

			requestedNodes[sourceChain].Add(nodeID)
			if _, exists := requestsPerNode[nodeID]; !exists {
				requestsPerNode[nodeID] = make([]*rmnpb.FixedDestLaneUpdateRequest, 0)
			}
			requestsPerNode[nodeID] = append(requestsPerNode[nodeID], updateRequest.data)
		}
	}

	requestIDs := c.sendObservationRequests(destChain, requestsPerNode)
	signedObservations, err := c.listenForRmnObservationResponses(
		ctx, destChain, requestIDs, updateRequestsPerChain, requestedNodes)
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

func (c *client) sendObservationRequests(
	destChain *rmnpb.LaneDest,
	requestsPerNode map[uint32][]*rmnpb.FixedDestLaneUpdateRequest,
) (requestIDs mapset.Set[uint64]) {
	requestIDs = mapset.NewSet[uint64]()

	for nodeID, requests := range requestsPerNode {
		req := &rmnpb.Request{
			RequestId: newRequestID(),
			Request: &rmnpb.Request_ObservationRequest{
				ObservationRequest: &rmnpb.ObservationRequest{
					LaneDest:                    destChain,
					FixedDestLaneUpdateRequests: requests,
				},
			},
		}

		c.lggr.Infow("sending observation request", "node", nodeID, "requestID", req.RequestId)
		if err := c.marshalAndSend(req, nodeID); err != nil {
			c.lggr.Errorw("failed to send observation request", "node_id", nodeID, "err", err)
			continue
		}
		requestIDs.Add(req.RequestId)
	}
	return requestIDs
}

func (c *client) listenForRmnObservationResponses(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	requestIDs mapset.Set[uint64],
	lursPerChain map[uint64]updateRequestWithMeta,
	requestedNodes map[uint64]mapset.Set[uint32],
) ([]rmnSignedObservationWithMeta, error) {
	c.lggr.Infow("Waiting for observation requests", "requestIDs", requestIDs.String())

	respChan := c.rawRmnClient.Recv()

	finishedRequestIDs := mapset.NewSet[uint64]()
	rmnObservationResponses := make([]rmnSignedObservationWithMeta, 0)
	merkleRootsCount := make(map[uint64]map[string]int) // sourceChain -> merkleRoot -> count

	initialObservationRequestTimer := time.NewTimer(c.observationsInitialRequestTimerDuration)
	defer initialObservationRequestTimer.Stop()
	for {
		select {
		case resp := <-respChan:
			parsedResp, err := c.parseResponse(&resp, requestIDs, finishedRequestIDs)
			if err != nil {
				c.lggr.Warnw("failed to parse RMN observation response", "err", err)
				continue
			}
			finishedRequestIDs.Add(parsedResp.RequestId)

			signedObs := parsedResp.GetSignedObservation()
			if signedObs == nil {
				c.lggr.Infof("RMN node returned an unexpected type of response: %+v", parsedResp.Response)
			} else {
				err := c.validateSignedObservationResponse(resp.RMNNodeID, lursPerChain, signedObs, destChain)
				if err != nil {
					c.lggr.Warnw("failed to validate signed observation response", "err", err)
					continue
				}

				c.lggr.Infow("received signed observation",
					"node", resp.RMNNodeID, "requestID", parsedResp.RequestId)

				rmnObservationResponses = append(rmnObservationResponses, rmnSignedObservationWithMeta{
					RMNNodeID:         resp.RMNNodeID,
					SignedObservation: signedObs,
				})
				for _, lu := range signedObs.Observation.FixedDestLaneUpdates {
					if _, exists := merkleRootsCount[lu.LaneSource.SourceChainSelector]; !exists {
						merkleRootsCount[lu.LaneSource.SourceChainSelector] = make(map[string]int)
					}
					merkleRootsCount[lu.LaneSource.SourceChainSelector][cciptypes.Bytes(lu.Root).String()]++
				}
			}

			// We got all the responses we were waiting for.
			if finishedRequestIDs.Equal(requestIDs) {
				c.lggr.Infof("all observation requests were finished")
				return rmnObservationResponses, nil
			}

			// We got sufficient number of observation responses with matching roots for every source chain.
			allChainsHaveEnoughResponses := true
			for chainRequest := range lursPerChain {
				var merkleRoot cciptypes.Bytes
				for merkleRootStr, count := range merkleRootsCount[chainRequest] {
					if count >= c.minObservers {
						merkleRoot, err = cciptypes.NewBytesFromString(merkleRootStr)
						if err != nil {
							return nil, fmt.Errorf("failed to parse merkle root: %w", err)
						}
						break
					}
				}
				if merkleRoot == nil {
					allChainsHaveEnoughResponses = false
					break
				}
			}

			if allChainsHaveEnoughResponses {
				c.lggr.Infof("all chains have enough observation responses with matching roots")
				return rmnObservationResponses, nil
			}
		case <-initialObservationRequestTimer.C:
			c.lggr.Warn("initial observation request timer expired, sending additional requests")
			// Timer expired, send the observation requests to the rest of the RMN nodes.
			requestsPerNode := make(map[uint32][]*rmnpb.FixedDestLaneUpdateRequest)
			for sourceChain, updateReq := range lursPerChain {
				for _, nodeID := range randomShuffle(updateReq.rmnNodes.ToSlice()) {
					if requestedNodes[sourceChain].Contains(nodeID) {
						continue
					}
					requestedNodes[sourceChain].Add(nodeID)
					if _, ok := requestsPerNode[nodeID]; !ok {
						requestsPerNode[nodeID] = make([]*rmnpb.FixedDestLaneUpdateRequest, 0)
					}
					requestsPerNode[nodeID] = append(requestsPerNode[nodeID], updateReq.data)
				}
			}
			newRequestIDs := c.sendObservationRequests(destChain, requestsPerNode)
			requestIDs = requestIDs.Union(newRequestIDs)
		case <-ctx.Done():
			return nil, ErrTimeout
		}
	}
}

func (c *client) validateSignedObservationResponse(
	rmnNodeID uint32,
	lurs map[uint64]updateRequestWithMeta,
	signedObs *rmnpb.SignedObservation,
	destChain *rmnpb.LaneDest,
) error {
	if signedObs.Observation.LaneDest.DestChainSelector != destChain.DestChainSelector {
		return fmt.Errorf("unexpected lane dest chain selector %v", signedObs.Observation.LaneDest)
	}
	if !bytes.Equal(signedObs.Observation.LaneDest.OfframpAddress, destChain.OfframpAddress) {
		return fmt.Errorf("unexpected lane dest offramp %v", signedObs.Observation.LaneDest)
	}

	if !bytes.Equal(signedObs.Observation.RmnHomeContractConfigDigest, c.rmnHomeConfigDigest) {
		return fmt.Errorf("unexpected rmn home contract config digest %x",
			signedObs.Observation.RmnHomeContractConfigDigest)
	}

	for _, signedObsLu := range signedObs.Observation.FixedDestLaneUpdates {
		updateReq, exists := lurs[signedObsLu.LaneSource.SourceChainSelector]
		if !exists {
			return fmt.Errorf("unexpected source chain selector %d", signedObsLu.LaneSource.SourceChainSelector)
		}

		if !updateReq.rmnNodes.Contains(rmnNodeID) {
			return fmt.Errorf("rmn node %d not expected to read chain %d",
				rmnNodeID, signedObsLu.LaneSource.SourceChainSelector)
		}

		if updateReq.data.ClosedInterval.MinMsgNr != signedObsLu.ClosedInterval.MinMsgNr {
			return fmt.Errorf("unexpected closed interval %v", signedObsLu.ClosedInterval)
		}
		if updateReq.data.ClosedInterval.MaxMsgNr != signedObsLu.ClosedInterval.MaxMsgNr {
			return fmt.Errorf("unexpected closed interval %v", signedObsLu.ClosedInterval)
		}

		if updateReq.data.LaneSource.SourceChainSelector != signedObsLu.LaneSource.SourceChainSelector {
			return fmt.Errorf("unexpected lane source %v", signedObsLu.LaneSource)
		}
		if !bytes.Equal(updateReq.data.LaneSource.OnrampAddress, signedObsLu.LaneSource.OnrampAddress) {
			return fmt.Errorf("unexpected lane source %v", signedObsLu.LaneSource)
		}

		if signedObsLu.Root == nil {
			return errors.New("root is nil")
		}
	}
	return nil
}

func (c *client) getRmnReportSignatures(
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

	var sigObservations []*rmnpb.AttributedSignedObservation
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

		c.lggr.Infow("sending report signature request", "node", node.ID, "requestID", req.RequestId)

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
	merkleRootCounts := make(map[uint64]map[string]int)
	for _, signedObs := range sigObservations {
		for _, lu := range signedObs.SignedObservation.Observation.FixedDestLaneUpdates {
			if _, exists := merkleRootCounts[lu.LaneSource.SourceChainSelector]; !exists {
				merkleRootCounts[lu.LaneSource.SourceChainSelector] = make(map[string]int)
			}
			merkleRootCounts[lu.LaneSource.SourceChainSelector][cciptypes.Bytes(lu.Root).String()]++
		}
	}

	rootsPerSourceChain := make(map[uint64]cciptypes.Bytes)
	for chain, counts := range merkleRootCounts {
		var root cciptypes.Bytes
		for merkleRootStr, count := range counts {
			if count >= c.minSigners {
				root, err = cciptypes.NewBytesFromString(merkleRootStr)
				if err != nil {
					return nil, fmt.Errorf("failed to parse merkle root: %w", err)
				}
				break
			}
		}
		if root == nil {
			return nil, fmt.Errorf("not enough roots for chain %d", chain)
		}
		rootsPerSourceChain[chain] = root
	}

	fixedDestLaneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0)
	addedUpdates := mapset.NewSet[string]()
	for _, signedObs := range sigObservations {
		for _, lu := range signedObs.SignedObservation.Observation.FixedDestLaneUpdates {
			rootStr := cciptypes.Bytes(lu.Root).String()
			if addedUpdates.Contains(rootStr) {
				continue
			}
			addedUpdates.Add(rootStr)
			fixedDestLaneUpdates = append(fixedDestLaneUpdates, lu)
		}
	}

	return &ReportSignatures{
		Signatures:  ecdsaSignatures,
		LaneUpdates: fixedDestLaneUpdates,
	}, nil
}

func (c *client) listenForRmnReportSignatures(
	ctx context.Context,
	requestIDs mapset.Set[uint64],
	reportSigReq *rmnpb.ReportSignatureRequest,
	signersRequested mapset.Set[uint32],
) ([]*rmnpb.EcdsaSignature, error) {
	tReportsInitialRequest := time.NewTimer(c.reportsInitialRequestTimerDuration)
	reportSigs := make([]*rmnpb.ReportSignature, 0)
	finishedRequests := mapset.NewSet[uint64]()
	respChan := c.rawRmnClient.Recv()
	requestIDs = requestIDs.Clone()

	for {
		select {
		case resp := <-respChan:
			responseTyp, err := c.parseResponse(&resp, requestIDs, finishedRequests)
			if err != nil {
				c.lggr.Errorw("failed to parse RMN signature response", "err", err)
				continue
			}

			reportSig := responseTyp.GetReportSignature()
			if reportSig == nil {
				c.lggr.Infof("RMN node returned an unexpected type of response: %+v", responseTyp.Response)
			} else {
				c.lggr.Infow("received report signature", "node", resp.RMNNodeID, "requestID", responseTyp.RequestId)
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
			c.lggr.Warnw("initial report signatures request timer expired, sending additional requests")
			// Send to the rest of the signers
			for _, node := range randomShuffle(c.rmnNodes) {
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

func (c *client) ensureEnoughSignedObservations(rmnSignedObservations []rmnSignedObservationWithMeta) error {
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

type updateRequestWithMeta struct {
	data     *rmnpb.FixedDestLaneUpdateRequest
	rmnNodes mapset.Set[uint32]
}

type rmnSignedObservationWithMeta struct {
	SignedObservation *rmnpb.SignedObservation
	RMNNodeID         uint32
}

func randomShuffle[T any](s []T) []T {
	ret := make([]T, len(s))
	for i, randIndex := range rand.Perm(len(s)) {
		ret[i] = s[randIndex]
	}
	return ret
}

func newRequestID() uint64 {
	b := make([]byte, 8)
	_, err := crand.Read(b)
	if err != nil {
		panic(err)
	}
	randomUint64 := binary.LittleEndian.Uint64(b)
	return randomUint64
}
