// controller.go contains functions and types related to the RMN controller.
// Controller functionality should be high-level functionality that the plugin can directly use.

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
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
)

var (
	// ErrTimeout is returned when the signature computation times out.
	ErrTimeout = errors.New("signature computation timeout")

	// ErrNothingToDo is returned when there are no source chains with enough RMN nodes.
	ErrNothingToDo = errors.New("nothing to observe from the existing RMN nodes, make " +
		"sure RMN is enabled, nodes configured correctly and minObservers value is correct")

	// ErrInsufficientObservationResponses is returned when we don't get enough observation responses to cover
	// all the observation requests.
	ErrInsufficientObservationResponses = errors.New("insufficient observation responses")

	// ErrInsufficientSignatureResponses is returned when we don't get enough report signatures.
	ErrInsufficientSignatureResponses = errors.New("insufficient signature responses")
)

// Controller contains the high-level functionality required by the plugin to interact with the RMN nodes.
type Controller interface {
	// ComputeReportSignatures computes and returns the signatures for the provided lane updates.
	// The returned ReportSignatures might contain a subset of the requested lane updates if some of them were not
	// able to get signed by the RMN nodes.
	//
	// This method abstracts away the RMN specific requests (ObservationRequest, ReportSignaturesRequest) and all the
	// necessary steps to compute the signatures, like retrying and caching which are handled by the implementation.
	ComputeReportSignatures(
		ctx context.Context,
		destChain *rmnpb.LaneDest,
		requestedUpdates []*rmnpb.FixedDestLaneUpdateRequest,
		rmnRemoteCfg rmntypes.RemoteConfig,
	) (*ReportSignatures, error)
}

type ReportSignatures struct {
	// Signatures are the ECDSA signatures for the lane updates for each node.
	Signatures []*rmnpb.EcdsaSignature
	// LaneUpdates are the lane updates for each source chain that we got the signatures for.
	// NOTE: A signature[i] corresponds to the whole LaneUpdates slice and NOT LaneUpdates[i].
	LaneUpdates []*rmnpb.FixedDestLaneUpdate
}

// controller is the base RMN Controller implementation.
type controller struct {
	lggr logger.Logger

	// signObservationPrefix is the prefix used to sign the observation.
	signObservationPrefix string

	// rmnCrypto contains the chain-specific crypto functions required by the RMN controller to verify signatures.
	rmnCrypto cciptypes.RMNCrypto

	// peerClient is the client used to communicate with the RMN nodes.
	peerClient PeerClient

	// rmnHomeReader is used to read the RMN home contract configuration.
	rmnHomeReader reader.RMNHome

	// ed25519Verifier is used to verify the RMN offchain observation signatures.
	ed25519Verifier ED25519Verifier

	// observationsInitialRequestTimerDuration is the duration of the initial observation request timer.
	// After this timer expires we send additional observation requests to the rest of the RMN nodes.
	observationsInitialRequestTimerDuration time.Duration

	// reportsInitialRequestTimerDuration is the duration of the initial report signature request timer.
	// After this timer expires we send additional report signature requests to the rest of the RMN nodes.
	reportsInitialRequestTimerDuration time.Duration
}

// NewController creates a new RMN Controller instance.
func NewController(
	lggr logger.Logger,
	rmnCrypto cciptypes.RMNCrypto,
	signObservationPrefix string,
	peerClient PeerClient,
	rmnHomeReader reader.RMNHome,
	observationsInitialRequestTimerDuration time.Duration,
	reportsInitialRequestTimerDuration time.Duration,
) Controller {
	return &controller{
		lggr:                                    lggr,
		rmnCrypto:                               rmnCrypto,
		signObservationPrefix:                   signObservationPrefix,
		peerClient:                              peerClient,
		rmnHomeReader:                           rmnHomeReader,
		ed25519Verifier:                         NewED25519Verifier(),
		observationsInitialRequestTimerDuration: observationsInitialRequestTimerDuration,
		reportsInitialRequestTimerDuration:      reportsInitialRequestTimerDuration,
	}
}

func (c *controller) ComputeReportSignatures(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	updateRequests []*rmnpb.FixedDestLaneUpdateRequest,
	rmnRemoteCfg rmntypes.RemoteConfig,
) (*ReportSignatures, error) {
	// Group the lane update requests by their source chain and mark the RMN nodes that can sign each update
	// based on whether it supports the source chain or not.
	updatesPerChain := make(map[uint64]updateRequestWithMeta)
	for _, updateReq := range updateRequests {
		if _, exists := updatesPerChain[updateReq.LaneSource.SourceChainSelector]; exists {
			return nil, errors.New("controller implementation assumes each lane update is for a different chain")
		}

		updatesPerChain[updateReq.LaneSource.SourceChainSelector] = updateRequestWithMeta{
			Data:     updateReq,
			RmnNodes: mapset.NewSet[rmntypes.NodeID](),
		}
		rmnNodes, err := c.rmnHomeReader.GetRMNNodesInfo(rmnRemoteCfg.ConfigDigest)
		if err != nil {
			return nil, fmt.Errorf("get rmn nodes info: %w", err)
		}

		for _, node := range rmnNodes {
			// If RMN node supports the chain add it to the list of RMN nodes that can sign the update.
			if node.SupportedSourceChains.Contains(cciptypes.ChainSelector(updateReq.LaneSource.SourceChainSelector)) {
				updatesPerChain[updateReq.LaneSource.SourceChainSelector].RmnNodes.Add(node.ID)
			}
		}
	}

	minObserversMap, err := c.rmnHomeReader.GetMinObservers(rmnRemoteCfg.ConfigDigest)
	if err != nil {
		return nil, fmt.Errorf("get min observers: %w", err)
	}
	// Filter out the lane update requests for chains without enough RMN nodes supporting them.
	for chain, l := range updatesPerChain {
		if _, exists := minObserversMap[cciptypes.ChainSelector(chain)]; !exists {
			return nil, fmt.Errorf("no min observers for chain %d", chain)
		}
		if l.RmnNodes.Cardinality() < minObserversMap[cciptypes.ChainSelector(chain)] {
			c.lggr.Warnw("chain skipped, not enough RMN nodes to support it",
				"chain", chain,
				"minObservers", minObserversMap[cciptypes.ChainSelector(chain)],
				"nodes", l.RmnNodes.ToSlice(),
			)
			delete(updatesPerChain, chain)
		}
	}

	if len(updatesPerChain) == 0 {
		return nil, ErrNothingToDo
	}

	tStart := time.Now()
	rmnSignedObservations, err := c.getRmnSignedObservations(
		ctx,
		destChain,
		updatesPerChain,
		rmnRemoteCfg.ConfigDigest,
		minObserversMap)
	// minObserversMap[cciptypes.ChainSelector(destChain.DestChainSelector)])
	if err != nil {
		return nil, fmt.Errorf("get rmn signed observations: %w", err)
	}
	c.lggr.Infow("received RMN signed observations",
		"requestedUpdates", updatesPerChain,
		"signedObservations", rmnSignedObservations,
		"duration", time.Since(tStart),
	)

	tStart = time.Now()
	rmnReportSignatures, err := c.getRmnReportSignatures(
		ctx,
		destChain,
		rmnSignedObservations,
		updatesPerChain,
		rmnRemoteCfg)
	if err != nil {
		return nil, fmt.Errorf("get rmn report signatures: %w", err)
	}
	c.lggr.Infow("received RMN report signatures",
		"signedObservations", rmnSignedObservations,
		"reportSignatures", rmnReportSignatures,
		"duration", time.Since(tStart),
	)

	return rmnReportSignatures, nil
}

// getRmnSignedObservations guarantees to return at least #minObservers signed observations for each source chain.
func (c *controller) getRmnSignedObservations(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	updateRequestsPerChain map[uint64]updateRequestWithMeta,
	configDigest cciptypes.Bytes32,
	minObserversMap map[cciptypes.ChainSelector]int,
) ([]rmnSignedObservationWithMeta, error) {
	requestedNodes := make(map[uint64]mapset.Set[rmntypes.NodeID])                   // sourceChain -> requested rmnNodeIDs
	requestsPerNode := make(map[rmntypes.NodeID][]*rmnpb.FixedDestLaneUpdateRequest) // grouped requests for each node

	// For each lane update request send an observation request to at most 'minObservers' number of rmn nodes.
	// At this point we can safely assume that we have at least #minObservers supporting each source chain.
	for sourceChain, updateRequest := range updateRequestsPerChain {
		requestedNodes[sourceChain] = mapset.NewSet[rmntypes.NodeID]()
		minObservers, exist := minObserversMap[cciptypes.ChainSelector(sourceChain)]
		if !exist {
			return nil, fmt.Errorf("no min observers for chain %d", sourceChain)
		}

		for nodeID := range updateRequest.RmnNodes.Iter() {
			if requestedNodes[sourceChain].Cardinality() >= minObservers {
				break // We have enough initial observers for this source chain.
			}

			requestedNodes[sourceChain].Add(nodeID)
			if _, exists := requestsPerNode[nodeID]; !exists {
				requestsPerNode[nodeID] = make([]*rmnpb.FixedDestLaneUpdateRequest, 0)
			}
			requestsPerNode[nodeID] = append(requestsPerNode[nodeID], updateRequest.Data)
		}
	}

	requestIDs := c.sendObservationRequests(destChain, requestsPerNode)

	signedObservations, err := c.listenForRmnObservationResponses(
		ctx, destChain, requestIDs, updateRequestsPerChain, requestedNodes, configDigest, minObserversMap)
	if err != nil {
		return nil, fmt.Errorf("listen for rmn observation responses: %w", err)
	}

	// Sanity check that we got enough signed observations for every source chain.
	// In practice this should never happen, an error must have been received earlier.
	if !gotSufficientObservationResponses(updateRequestsPerChain, signedObservations, minObserversMap) {
		return nil, fmt.Errorf("not enough signed observations after sanity check")
	}

	return signedObservations, nil
}

// sendObservationRequests sends observation requests to the RMN nodes.
// If a specific request fails, it is logged and not included in the returned requestIDs mapping.
func (c *controller) sendObservationRequests(
	destChain *rmnpb.LaneDest,
	requestsPerNode map[rmntypes.NodeID][]*rmnpb.FixedDestLaneUpdateRequest,
) (requestIDs mapset.Set[uint64]) {
	requestIDs = mapset.NewSet[uint64]()

	for nodeID, requests := range requestsPerNode {
		sort.Slice(requests, func(i, j int) bool {
			return requests[i].LaneSource.SourceChainSelector < requests[j].LaneSource.SourceChainSelector
		})

		req := &rmnpb.Request{
			RequestId: newRequestID(),
			Request: &rmnpb.Request_ObservationRequest{
				ObservationRequest: &rmnpb.ObservationRequest{
					LaneDest:                    destChain,
					FixedDestLaneUpdateRequests: requests,
				},
			},
		}

		lggr := logger.With(c.lggr, "node", nodeID, "requestID", req.RequestId)
		lggr.Infow("sending observation request", "laneUpdateRequests", requests)
		if err := c.marshalAndSend(req, nodeID); err != nil {
			lggr.Errorw("failed to send observation request", "err", err)
			continue
		}

		requestIDs.Add(req.RequestId)
	}

	return requestIDs
}

// listenForRmnObservationResponses listens for the RMN observation responses.
// It waits for the responses until all the requests are finished or until the context is done.
// It is responsible for sending additional observation requests if the initial requests timeout.
// nolint:gocyclo // todo
func (c *controller) listenForRmnObservationResponses(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	requestIDs mapset.Set[uint64],
	lursPerChain map[uint64]updateRequestWithMeta,
	requestedNodes map[uint64]mapset.Set[rmntypes.NodeID],
	configDigest cciptypes.Bytes32,
	minObserversMap map[cciptypes.ChainSelector]int,
) ([]rmnSignedObservationWithMeta, error) {
	c.lggr.Infow("listening for RMN observation responses", "requestIDs", requestIDs.String())

	finishedRequestIDs := mapset.NewSet[uint64]()
	rmnObservationResponses := make([]rmnSignedObservationWithMeta, 0)

	initialObservationRequestTimer := time.NewTimer(c.observationsInitialRequestTimerDuration)
	timerExpired := false

	defer initialObservationRequestTimer.Stop()
	for {
		select {
		case resp := <-c.peerClient.Recv():
			parsedResp, err := c.parseResponse(&resp, requestIDs, finishedRequestIDs)
			if err != nil {
				c.lggr.Debugw("skipping an unexpected RMN response", "err", err)
				continue
			}

			finishedRequestIDs.Add(parsedResp.RequestId)

			err = c.validateSignedObservationResponse(
				parsedResp,
				resp.RMNNodeID,
				lursPerChain,
				destChain,
				configDigest,
			)
			if err != nil {
				c.lggr.Warnw("skipping an invalid RMN observation response", "err", err)
				initialObservationRequestTimer.Reset(0) // immediately schedule the additional requests
			} else {
				rmnObservationResponses = append(rmnObservationResponses, rmnSignedObservationWithMeta{
					SignedObservation: parsedResp.GetSignedObservation(),
					RMNNodeID:         resp.RMNNodeID,
				})
			}

			allChainsHaveEnoughResponses := gotSufficientObservationResponses(
				lursPerChain, rmnObservationResponses, minObserversMap)
			if allChainsHaveEnoughResponses {
				c.lggr.Infof("all chains have enough observation responses with matching roots")
				return rmnObservationResponses, nil
			}

			// We got all the responses we were waiting for, but they are not sufficient for all chains.
			if timerExpired && finishedRequestIDs.Equal(requestIDs) {
				c.lggr.Warnw("observation requests were finished, but results are not sufficient")
				return rmnObservationResponses, ErrInsufficientObservationResponses
			}
		case <-initialObservationRequestTimer.C:
			if timerExpired {
				continue
			}
			timerExpired = true

			c.lggr.Warn("sending additional RMN observation requests")
			requestsPerNode := make(map[rmntypes.NodeID][]*rmnpb.FixedDestLaneUpdateRequest)
			for sourceChain, updateReq := range lursPerChain {
				for _, nodeID := range randomShuffle(updateReq.RmnNodes.ToSlice()) {
					if requestedNodes[sourceChain].Contains(nodeID) {
						continue
					}
					requestedNodes[sourceChain].Add(nodeID)
					requestsPerNode[nodeID] = append(requestsPerNode[nodeID], updateReq.Data)
				}
			}
			newRequestIDs := c.sendObservationRequests(destChain, requestsPerNode)
			requestIDs = requestIDs.Union(newRequestIDs)
		case <-ctx.Done():
			return nil, ErrTimeout
		}
	}
}

// gotSufficientObservationResponses checks if we got enough observation responses for each source chain.
// Enough meaning that we got at least #minObservers observing the same merkle root for a target chain.
func gotSufficientObservationResponses(
	updateRequests map[uint64]updateRequestWithMeta,
	rmnObservationResponses []rmnSignedObservationWithMeta,
	minObserversMap map[cciptypes.ChainSelector]int,
) bool {
	merkleRootsCount := make(map[uint64]map[cciptypes.Bytes32]int)
	for _, signedObs := range rmnObservationResponses {
		for _, lu := range signedObs.SignedObservation.Observation.FixedDestLaneUpdates {
			if _, exists := merkleRootsCount[lu.LaneSource.SourceChainSelector]; !exists {
				merkleRootsCount[lu.LaneSource.SourceChainSelector] = make(map[cciptypes.Bytes32]int)
			}
			merkleRootsCount[lu.LaneSource.SourceChainSelector][cciptypes.Bytes32(lu.Root)]++
		}
	}

	for sourceChain := range updateRequests {
		// make sure we got at least #minObservers observing the same merkle root for a target chain.
		countsPerRoot, ok := merkleRootsCount[sourceChain]
		if !ok || len(countsPerRoot) == 0 {
			return false
		}
		minObservers, exists := minObserversMap[cciptypes.ChainSelector(sourceChain)]
		if !exists {
			return false
		}

		values := maps.Values(countsPerRoot)
		sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
		if values[len(values)-1] < minObservers {
			return false
		}
	}

	return true
}

// nolint:gocyclo // todo
func (c *controller) validateSignedObservationResponse(
	parsedResp *rmnpb.Response,
	rmnNodeID rmntypes.NodeID,
	lurs map[uint64]updateRequestWithMeta,
	destChain *rmnpb.LaneDest,
	rmnHomeConfigDigest cciptypes.Bytes32,
) error {
	signedObs := parsedResp.GetSignedObservation()
	if signedObs == nil {
		return fmt.Errorf("got an unexpected type of response %T", parsedResp.Response)
	}

	rmnNode, exists := c.getHomeNodeByID(rmnHomeConfigDigest, rmnNodeID)
	if !exists {
		return fmt.Errorf("rmn node %d not found", rmnNodeID)
	}

	if signedObs.Observation.LaneDest.DestChainSelector != destChain.DestChainSelector {
		return fmt.Errorf("unexpected lane dest chain selector %v", signedObs.Observation.LaneDest)
	}
	if !bytes.Equal(signedObs.Observation.LaneDest.OfframpAddress, destChain.OfframpAddress) {
		return fmt.Errorf("unexpected lane dest offramp %v", signedObs.Observation.LaneDest)
	}

	if !bytes.Equal(signedObs.Observation.RmnHomeContractConfigDigest, rmnHomeConfigDigest[:]) {
		return fmt.Errorf("unexpected rmn home contract config digest %x",
			signedObs.Observation.RmnHomeContractConfigDigest)
	}

	for _, signedObsLu := range signedObs.Observation.FixedDestLaneUpdates {
		updateReq, exists := lurs[signedObsLu.LaneSource.SourceChainSelector]
		if !exists {
			return fmt.Errorf("unexpected source chain selector %d", signedObsLu.LaneSource.SourceChainSelector)
		}

		if !updateReq.RmnNodes.Contains(rmnNodeID) {
			return fmt.Errorf("rmn node %d not expected to read chain %d",
				rmnNodeID, signedObsLu.LaneSource.SourceChainSelector)
		}

		if updateReq.Data.ClosedInterval.MinMsgNr != signedObsLu.ClosedInterval.MinMsgNr {
			return fmt.Errorf("unexpected closed interval %v", signedObsLu.ClosedInterval)
		}
		if updateReq.Data.ClosedInterval.MaxMsgNr != signedObsLu.ClosedInterval.MaxMsgNr {
			return fmt.Errorf("unexpected closed interval %v", signedObsLu.ClosedInterval)
		}

		if updateReq.Data.LaneSource.SourceChainSelector != signedObsLu.LaneSource.SourceChainSelector {
			return fmt.Errorf("unexpected lane source %v", signedObsLu.LaneSource)
		}
		if !bytes.Equal(updateReq.Data.LaneSource.OnrampAddress, signedObsLu.LaneSource.OnrampAddress) {
			return fmt.Errorf("unexpected lane source %v", signedObsLu.LaneSource)
		}

		if signedObsLu.Root == nil {
			return errors.New("root is nil")
		}
	}

	if err := verifyObservationSignature(rmnNode, c.signObservationPrefix, signedObs, c.ed25519Verifier); err != nil {
		return fmt.Errorf("failed to verify observation signature: %w", err)
	}
	return nil
}

func (c *controller) getRmnReportSignatures(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	rmnSignedObservations []rmnSignedObservationWithMeta,
	updatesPerChain map[uint64]updateRequestWithMeta,
	rmnRemoteCfg rmntypes.RemoteConfig,
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
	// node3: getObservations(1)          ->                         responds_ok (requested after node1 timeout)

	// At this point it is also possible that the signed observations contain
	// different roots for the same source chain and interval.

	rootsPerChain := getMostVotedRootsFromObservations(rmnSignedObservations)

	fixedDestLaneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0)
	for sourceChain, updateReq := range updatesPerChain {
		mostVotedRoot, ok := rootsPerChain[cciptypes.ChainSelector(sourceChain)]
		if !ok {
			return nil, fmt.Errorf("no most voted root for source chain %d", sourceChain)
		}
		fixedDestLaneUpdates = append(fixedDestLaneUpdates, &rmnpb.FixedDestLaneUpdate{
			LaneSource:     updateReq.Data.LaneSource,
			ClosedInterval: updateReq.Data.ClosedInterval,
			Root:           mostVotedRoot[:],
		})
	}
	sort.Slice(fixedDestLaneUpdates, func(i, j int) bool {
		return fixedDestLaneUpdates[i].LaneSource.SourceChainSelector < fixedDestLaneUpdates[j].LaneSource.SourceChainSelector
	})

	destChainInfo, exists := chainsel.ChainBySelector(destChain.DestChainSelector)
	if !exists {
		return nil, fmt.Errorf("unknown dest chain selector %d", destChain.DestChainSelector)
	}

	laneUpdates, err := NewLaneUpdatesFromPB(fixedDestLaneUpdates)
	if err != nil {
		return nil, fmt.Errorf("failed to convert lane updates from protobuf: %w", err)
	}

	rmnReport := cciptypes.RMNReport{
		ReportVersion:               rmnRemoteCfg.RmnReportVersion.String(),
		DestChainID:                 cciptypes.NewBigIntFromInt64(int64(destChainInfo.EvmChainID)),
		DestChainSelector:           cciptypes.ChainSelector(destChain.DestChainSelector),
		RmnRemoteContractAddress:    rmnRemoteCfg.ContractAddress,
		OfframpAddress:              destChain.OfframpAddress,
		RmnHomeContractConfigDigest: rmnRemoteCfg.ConfigDigest,
		LaneUpdates:                 laneUpdates,
	}

	chainInfo, exists := chainsel.ChainBySelector(destChain.DestChainSelector)
	if !exists {
		return nil, fmt.Errorf("unknown chain selector %d", destChain.DestChainSelector)
	}

	reportSigReq := &rmnpb.ReportSignatureRequest{
		Context: &rmnpb.ReportContext{
			EvmDestChainId:              chainInfo.EvmChainID,
			RmnRemoteContractAddress:    rmnRemoteCfg.ContractAddress,
			RmnHomeContractConfigDigest: rmnRemoteCfg.ConfigDigest[:],
			LaneDest:                    destChain,
		},
		AttributedSignedObservations: transformAndSortObservations(rmnSignedObservations),
	}
	minSigners := rmnRemoteCfg.MinSigners
	signers := rmnRemoteCfg.Signers
	requestIDs, signersRequested, err := c.sendReportSignatureRequest(
		reportSigReq,
		signers,
		int(minSigners))
	if err != nil {
		return nil, fmt.Errorf("send report signature request: %w", err)
	}

	ecdsaSignatures, err := c.listenForRmnReportSignatures(
		ctx,
		requestIDs,
		rmnReport,
		reportSigReq,
		signersRequested,
		signers,
		int(minSigners))
	if err != nil {
		return nil, fmt.Errorf("listen for rmn report signatures: %w", err)
	}

	return &ReportSignatures{
		Signatures:  ecdsaSignatures,
		LaneUpdates: fixedDestLaneUpdates,
	}, nil
}

// transformAndSortObservations transforms the given slice to AttributedSignedObservations.
// Results are ordered by the RMN node index and then by the source chain selector of the first observation request.
func transformAndSortObservations(
	rmnSignedObservations []rmnSignedObservationWithMeta) []*rmnpb.AttributedSignedObservation {
	var attrSigObservations []*rmnpb.AttributedSignedObservation

	for _, so := range rmnSignedObservations {
		// order observations by source chain
		sort.Slice(so.SignedObservation.Observation.FixedDestLaneUpdates, func(i, j int) bool {
			return so.SignedObservation.Observation.FixedDestLaneUpdates[i].LaneSource.SourceChainSelector <
				so.SignedObservation.Observation.FixedDestLaneUpdates[j].LaneSource.SourceChainSelector
		})

		attrSigObservations = append(attrSigObservations, &rmnpb.AttributedSignedObservation{
			SignedObservation: so.SignedObservation,
			SignerNodeIndex:   uint32(so.RMNNodeID),
		})
	}

	// order by node index and then by source chain selector of the first observation request (should exist)
	sort.Slice(attrSigObservations, func(i, j int) bool {
		if attrSigObservations[i].SignerNodeIndex == attrSigObservations[j].SignerNodeIndex {
			return attrSigObservations[i].SignedObservation.Observation.FixedDestLaneUpdates[0].LaneSource.SourceChainSelector <
				attrSigObservations[j].SignedObservation.Observation.FixedDestLaneUpdates[0].LaneSource.SourceChainSelector
		}
		return attrSigObservations[i].SignerNodeIndex < attrSigObservations[j].SignerNodeIndex
	})

	return attrSigObservations
}

// getMostVotedRootsFromObservations returns the most voted roots for each source chain.
func getMostVotedRootsFromObservations(
	observations []rmnSignedObservationWithMeta) map[cciptypes.ChainSelector]cciptypes.Bytes32 {
	votesPerRoot := make(map[cciptypes.ChainSelector]map[cciptypes.Bytes32]int)
	for _, so := range observations {
		for _, lu := range so.SignedObservation.Observation.FixedDestLaneUpdates {
			if _, exists := votesPerRoot[cciptypes.ChainSelector(lu.LaneSource.SourceChainSelector)]; !exists {
				votesPerRoot[cciptypes.ChainSelector(lu.LaneSource.SourceChainSelector)] = make(map[cciptypes.Bytes32]int)
			}
			votesPerRoot[cciptypes.ChainSelector(lu.LaneSource.SourceChainSelector)][cciptypes.Bytes32(lu.Root)]++
		}
	}

	mostVotedRoots := make(map[cciptypes.ChainSelector]cciptypes.Bytes32)
	for chain, votes := range votesPerRoot {
		var mostVotedRoot cciptypes.Bytes32
		maxVotes := 0
		for root, vote := range votes {
			if vote > maxVotes {
				mostVotedRoot = root
				maxVotes = vote
			}
		}
		mostVotedRoots[chain] = mostVotedRoot
	}

	return mostVotedRoots
}

// sendReportSignatureRequest sends the report signature request to #minSigners random RMN nodes.
// If not enough requests were sent, it returns an error.
func (c *controller) sendReportSignatureRequest(
	reportSigReq *rmnpb.ReportSignatureRequest,
	remoteSigners []rmntypes.RemoteSignerInfo,
	minSigners int,
) (
	requestIDs mapset.Set[uint64], signersRequested mapset.Set[rmntypes.NodeID], err error) {
	requestIDs = mapset.NewSet[uint64]()
	signersRequested = mapset.NewSet[rmntypes.NodeID]()

	// Send the report signature request to at least minSigners
	for _, node := range randomShuffle(remoteSigners) {
		if requestIDs.Cardinality() >= minSigners {
			break
		}

		req := &rmnpb.Request{
			RequestId: newRequestID(),
			Request: &rmnpb.Request_ReportSignatureRequest{
				ReportSignatureRequest: reportSigReq,
			},
		}

		c.lggr.Infow("sending report signature request", "node", node.NodeIndex, "requestID", req.RequestId)

		err := c.marshalAndSend(req, rmntypes.NodeID(node.NodeIndex))
		if err != nil {
			c.lggr.Warnw("failed to send report signature request", "node", node.NodeIndex, "err", err)
			continue
		}

		requestIDs.Add(req.RequestId)
		signersRequested.Add(rmntypes.NodeID(node.NodeIndex))
	}

	if requestIDs.Cardinality() < minSigners {
		return requestIDs, signersRequested, fmt.Errorf("not able to send to enough report signers")
	}
	return requestIDs, signersRequested, nil
}

// reportSigWithSignerAddress is a helper struct to store the report signature and
// the address of the private key that signed it.
type reportSigWithSignerAddress struct {
	reportSig     *rmnpb.ReportSignature
	signerAddress cciptypes.Bytes
}

// nolint:gocyclo // todo
func (c *controller) listenForRmnReportSignatures(
	ctx context.Context,
	requestIDs mapset.Set[uint64],
	rmnReport cciptypes.RMNReport,
	reportSigReq *rmnpb.ReportSignatureRequest,
	signersRequested mapset.Set[rmntypes.NodeID],
	signers []rmntypes.RemoteSignerInfo,
	minSigners int,
) ([]*rmnpb.EcdsaSignature, error) {
	tReportsInitialRequest := time.NewTimer(c.reportsInitialRequestTimerDuration)
	timerExpired := false

	reportSigs := make([]reportSigWithSignerAddress, 0)
	finishedRequests := mapset.NewSet[uint64]()
	requestIDs = requestIDs.Clone()
	c.lggr.Infof("waiting for report signatures, requestIDs: %s", requestIDs.String())

	for {
		select {
		case resp := <-c.peerClient.Recv():
			responseTyp, err := c.parseResponse(&resp, requestIDs, finishedRequests)
			if err != nil {
				c.lggr.Infow("failed to parse RMN signature response", "err", err)
				continue
			}
			finishedRequests.Add(responseTyp.RequestId)

			reportSig, err := c.validateReportSigResponse(ctx, responseTyp, resp.RMNNodeID, signers, rmnReport)
			if err != nil {
				c.lggr.Warnw("skipping an invalid RMN report signature response", "err", err)
				tReportsInitialRequest.Reset(0) // schedule additional requests if any
			} else {
				c.lggr.Infow("received valid report signature", "node", resp.RMNNodeID, "requestID", responseTyp.RequestId)
				reportSigs = append(reportSigs, *reportSig)
			}

			if len(reportSigs) >= minSigners {
				c.lggr.Infof("got enough RMN report signatures")
				return sortAndParseReportSigs(reportSigs), nil
			}

			if timerExpired && finishedRequests.Equal(requestIDs) {
				c.lggr.Warn("report signature requests were finished, but results are not sufficient")
				return nil, ErrInsufficientSignatureResponses
			}
		case <-tReportsInitialRequest.C:
			if timerExpired {
				continue
			}
			timerExpired = true

			c.lggr.Warnw("sending additional RMN signature requests")

			for _, node := range randomShuffle(signers) {
				if signersRequested.Contains(rmntypes.NodeID(node.NodeIndex)) {
					continue
				}
				req := &rmnpb.Request{
					RequestId: newRequestID(),
					Request: &rmnpb.Request_ReportSignatureRequest{
						ReportSignatureRequest: reportSigReq,
					},
				}

				c.lggr.Infow("sending report signature request", "node", node.NodeIndex, "requestID", req.RequestId)
				if err := c.marshalAndSend(req, rmntypes.NodeID(node.NodeIndex)); err != nil {
					c.lggr.Errorw("failed to send report signature request", "node", node.NodeIndex, "err", err)
					continue
				}
				requestIDs.Add(req.RequestId)
				signersRequested.Add(rmntypes.NodeID(node.NodeIndex))
			}
		case <-ctx.Done():
			return nil, ErrTimeout
		}
	}
}

func sortAndParseReportSigs(reportSigs []reportSigWithSignerAddress) []*rmnpb.EcdsaSignature {
	// Sort report sigs by signer address
	// Similar to RMNRemote.verify (signatures must be sorted in ascending order by signer address).
	sort.Slice(reportSigs, func(i, j int) bool {
		return reportSigs[i].signerAddress.String() < reportSigs[j].signerAddress.String()
	})

	ecdsaSigs := make([]*rmnpb.EcdsaSignature, 0)
	for _, rs := range reportSigs {
		ecdsaSigs = append(ecdsaSigs, rs.reportSig.Signature)
	}
	return ecdsaSigs
}

func (c *controller) validateReportSigResponse(
	ctx context.Context,
	responseTyp *rmnpb.Response,
	nodeID rmntypes.NodeID,
	signers []rmntypes.RemoteSignerInfo,
	rmnReport cciptypes.RMNReport,
) (*reportSigWithSignerAddress, error) {
	signerNode, exists := c.getSignerNodeByID(signers, nodeID)
	if !exists {
		return nil, fmt.Errorf("rmn node %d not found", nodeID)
	}

	reportSig := responseTyp.GetReportSignature()
	if reportSig == nil {
		return nil, fmt.Errorf("got an unexpected type of response %T", responseTyp.Response)
	}

	sig, err := NewECDSASigFromPB(reportSig.Signature)
	if err != nil {
		return nil, fmt.Errorf("failed to convert signature from protobuf: %w", err)
	}

	err = c.rmnCrypto.VerifyReportSignatures(
		ctx,
		[]cciptypes.RMNECDSASignature{*sig},
		rmnReport,
		[]cciptypes.Bytes{signerNode.OnchainPublicKey},
	)

	if err != nil {
		return nil, fmt.Errorf("failed to verify report signature: %w", err)
	}

	return &reportSigWithSignerAddress{
		reportSig:     reportSig,
		signerAddress: signerNode.OnchainPublicKey,
	}, nil
}

func (c *controller) getHomeNodeByID(
	configDigest cciptypes.Bytes32,
	nodeID rmntypes.NodeID) (rmntypes.HomeNodeInfo, bool) {
	rmnNodes, err := c.rmnHomeReader.GetRMNNodesInfo(configDigest)
	if err != nil {
		return rmntypes.HomeNodeInfo{}, false
	}

	for _, node := range rmnNodes {
		if node.ID == nodeID {
			return node, true
		}
	}
	return rmntypes.HomeNodeInfo{}, false
}

func (c *controller) getSignerNodeByID(
	rmnNodes []rmntypes.RemoteSignerInfo,
	nodeID rmntypes.NodeID) (rmntypes.RemoteSignerInfo, bool) {
	for _, node := range rmnNodes {
		if rmntypes.NodeID(node.NodeIndex) == nodeID {
			return node, true
		}
	}
	return rmntypes.RemoteSignerInfo{}, false
}

type updateRequestWithMeta struct {
	Data     *rmnpb.FixedDestLaneUpdateRequest
	RmnNodes mapset.Set[rmntypes.NodeID]
}

type rmnSignedObservationWithMeta struct {
	SignedObservation *rmnpb.SignedObservation
	RMNNodeID         rmntypes.NodeID
}

func (c *controller) marshalAndSend(req *rmnpb.Request, nodeID rmntypes.NodeID) error {
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("proto marshal RMN request: %w", err)
	}

	if err := c.peerClient.Send(nodeID, reqBytes); err != nil {
		return fmt.Errorf("send rmn request: %w", err)
	}

	return nil
}

// parseResponse parses the response from the RMN and returns the response.
// Validates that the response is expected and not a duplicate.
func (c *controller) parseResponse(
	resp *PeerResponse, requestIDs, gotResponses mapset.Set[uint64]) (*rmnpb.Response, error) {
	responseTyp := &rmnpb.Response{}
	err := proto.Unmarshal(resp.Body, responseTyp)
	if err != nil {
		return nil, fmt.Errorf("proto unmarshal: %w", err)
	}

	if !requestIDs.Contains(responseTyp.RequestId) {
		return nil, fmt.Errorf(
			"got an RMN response that we are not waiting for: %d (%s)", responseTyp.RequestId, requestIDs.String())
	}

	if gotResponses.Contains(responseTyp.RequestId) {
		return nil, fmt.Errorf("got a duplicate RMN response: %d", responseTyp.RequestId)
	}

	return responseTyp, nil
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
