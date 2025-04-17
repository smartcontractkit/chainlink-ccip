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
	"slices"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/proto"

	typconv "github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv"

	chainsel "github.com/smartcontractkit/chain-selectors"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"

	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	// ErrTimeout is returned when the signature computation times out.
	ErrTimeout = errors.New("signature computation timeout")

	// ErrNothingToDo is returned when there are no source chains with enough RMN nodes.
	ErrNothingToDo = errors.New("nothing to observe from the existing RMN nodes, make " +
		"sure RMN is enabled, nodes configured correctly and F value is correct")

	ErrAllChainsNotReady = errors.New("none of the requested chains is ready to be observed by the RMN nodes")

	// ErrInsufficientObservationResponses is returned when we don't get enough observation responses to cover
	// all the observation requests.
	ErrInsufficientObservationResponses = errors.New("insufficient observation responses")

	// ErrInsufficientSignatureResponses is returned when we don't get enough report signatures.
	ErrInsufficientSignatureResponses = errors.New("insufficient signature responses")

	// ErrNotFound is returned when the RMN node is not found in the RMN home or the Remote contracts.
	ErrNotFound = errors.New("rmn node not found")
)

// Controller contains the high-level functionality required by the plugin to interact with the RMN nodes.
type Controller interface {
	// InitConnection initializes the connection to the generic peer group endpoint and must be called before
	// further Controller interaction. If called twice it overwrites the previous connection.
	InitConnection(
		ctx context.Context,
		commitConfigDigest cciptypes.Bytes32,
		rmnHomeConfigDigest cciptypes.Bytes32,
		oraclePeerIDs []ragep2ptypes.PeerID,
		rmnNodes []rmntypes.HomeNodeInfo,
	) error

	// Close closes the connection to the generic peer group endpoint and all the underlying streams.
	Close() error

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
		rmnRemoteCfg cciptypes.RemoteConfig,
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
	rmnHomeReader readerpkg.RMNHome

	// ed25519Verifier is used to verify the RMN offchain observation signatures.
	ed25519Verifier ED25519Verifier

	// observationsInitialRequestTimerDuration is the duration of the initial observation request timer.
	// After this timer expires we send additional observation requests to the rest of the RMN nodes.
	observationsInitialRequestTimerDuration time.Duration

	// reportsInitialRequestTimerDuration is the duration of the initial report signature request timer.
	// After this timer expires we send additional report signature requests to the rest of the RMN nodes.
	reportsInitialRequestTimerDuration time.Duration

	metricsReporter MetricsReporter
}

// NewController creates a new RMN Controller instance.
func NewController(
	lggr logger.Logger,
	rmnCrypto cciptypes.RMNCrypto,
	signObservationPrefix string,
	peerClient PeerClient,
	rmnHomeReader readerpkg.RMNHome,
	observationsInitialRequestTimerDuration time.Duration,
	reportsInitialRequestTimerDuration time.Duration,
	metricsReporter MetricsReporter,
) Controller {

	lggr.Infow("creating new RMN controller",
		"observationsInitialRequestTimerDuration", observationsInitialRequestTimerDuration,
		"reportsInitialRequestTimerDuration", reportsInitialRequestTimerDuration)

	return &controller{
		lggr:                                    lggr,
		rmnCrypto:                               rmnCrypto,
		signObservationPrefix:                   signObservationPrefix,
		peerClient:                              peerClient,
		rmnHomeReader:                           rmnHomeReader,
		ed25519Verifier:                         NewED25519Verifier(),
		observationsInitialRequestTimerDuration: observationsInitialRequestTimerDuration,
		reportsInitialRequestTimerDuration:      reportsInitialRequestTimerDuration,
		metricsReporter:                         metricsReporter,
	}
}

func (c *controller) ComputeReportSignatures(
	ctx context.Context,
	destChain *rmnpb.LaneDest,
	updateRequests []*rmnpb.FixedDestLaneUpdateRequest,
	rmnRemoteCfg cciptypes.RemoteConfig,
) (*ReportSignatures, error) {
	lggr := logutil.WithContextValues(ctx, c.lggr)

	rmnNodes, err := c.rmnHomeReader.GetRMNNodesInfo(rmnRemoteCfg.ConfigDigest)
	if err != nil {
		return nil, fmt.Errorf("get rmn nodes info: %w", err)
	}

	rmnNodeInfo := make(map[rmntypes.NodeID]rmntypes.HomeNodeInfo, len(rmnNodes))
	for _, node := range rmnNodes {
		rmnNodeInfo[node.ID] = node
	}

	lggr.Infow("got RMN nodes info", "nodes", rmnNodeInfo)
	lggr.Infow("requested updates", "updates", updateRequests)

	updatesPerChain, err := populateUpdatesPerChain(updateRequests, rmnNodes)
	if err != nil {
		return nil, fmt.Errorf("process update requests: %w", err)
	}

	homeFMap, err := c.rmnHomeReader.GetFObserve(rmnRemoteCfg.ConfigDigest)
	if err != nil {
		return nil, fmt.Errorf("get home F: %w", err)
	}
	// Filter out the lane update requests for chains without enough RMN nodes supporting them.
	for chain, l := range updatesPerChain {
		homeChainF, exists := homeFMap[cciptypes.ChainSelector(chain)]
		if !exists {
			lggr.Errorw("no home F for chain, chain skipped", "chain", chain)
			delete(updatesPerChain, chain)
		}

		if consensus.LtFPlusOne(homeChainF, l.RmnNodes.Cardinality()) {
			lggr.Warnw("chain skipped, not enough RMN nodes to support it",
				"chain", chain,
				"homeF", homeFMap[cciptypes.ChainSelector(chain)],
				"nodes", l.RmnNodes.ToSlice(),
			)
			delete(updatesPerChain, chain)
		}
	}

	if len(updatesPerChain) == 0 {
		return nil, ErrNothingToDo
	}

	tStart := time.Now()
	rmnSignedObservations, laneUpdatesToMakeProgressWith, err := c.getRmnSignedObservations(
		ctx,
		lggr,
		destChain,
		updatesPerChain,
		rmnRemoteCfg.ConfigDigest,
		homeFMap,
		rmnNodeInfo)
	if err != nil {
		return nil, fmt.Errorf("get rmn signed observations: %w", err)
	}
	lggr.Infow("received RMN signed observations",
		"requestedUpdates", updatesPerChain,
		"signedObservations", rmnSignedObservations,
		"duration", time.Since(tStart),
	)

	tStart = time.Now()
	rmnReportSignatures, err := c.getRmnReportSignatures(
		ctx,
		lggr,
		destChain,
		rmnSignedObservations,
		laneUpdatesToMakeProgressWith,
		rmnRemoteCfg,
		rmnNodeInfo)
	if err != nil {
		return nil, fmt.Errorf("get rmn report signatures: %w", err)
	}
	lggr.Infow("received RMN report signatures",
		"signedObservations", rmnSignedObservations,
		"reportSignatures", rmnReportSignatures,
		"duration", time.Since(tStart),
	)

	return rmnReportSignatures, nil
}

// populateUpdatesPerChain processes a list of update requests, groups the lane updates by their source chain
// and populates the items with metadata for each update request, including a set of RMN nodes supporting that request.
func populateUpdatesPerChain(
	updateRequests []*rmnpb.FixedDestLaneUpdateRequest,
	rmnNodes []rmntypes.HomeNodeInfo,
) (map[uint64]updateRequestWithMeta, error) {
	updatesPerChain := make(map[uint64]updateRequestWithMeta)

	for _, updateReq := range updateRequests {
		if _, exists := updatesPerChain[updateReq.LaneSource.SourceChainSelector]; exists {
			return nil, errors.New("controller implementation assumes each lane update is for a different chain")
		}

		updatesPerChain[updateReq.LaneSource.SourceChainSelector] = updateRequestWithMeta{
			Data:     updateReq,
			RmnNodes: mapset.NewSet[rmntypes.NodeID](),
		}

		for _, node := range rmnNodes {
			// If RMN node supports the chain add it to the list of RMN nodes that can sign the update.
			if node.SupportedSourceChains.Contains(cciptypes.ChainSelector(updateReq.LaneSource.SourceChainSelector)) {
				updatesPerChain[updateReq.LaneSource.SourceChainSelector].RmnNodes.Add(node.ID)
			}
		}
	}

	return updatesPerChain, nil
}

func (c *controller) InitConnection(
	ctx context.Context,
	commitConfigDigest cciptypes.Bytes32,
	rmnHomeConfigDigest cciptypes.Bytes32,
	oraclePeerIDs []ragep2ptypes.PeerID,
	rmnNodes []rmntypes.HomeNodeInfo,
) error {
	return c.peerClient.InitConnection(ctx, commitConfigDigest, rmnHomeConfigDigest, oraclePeerIDs, rmnNodes)
}

func (c *controller) Close() error {
	return c.peerClient.Close()
}

// getRmnSignedObservations guarantees to return at least F+1 signed observations for each source chain.
func (c *controller) getRmnSignedObservations(
	ctx context.Context,
	lggr logger.Logger,
	destChain *rmnpb.LaneDest,
	updateRequestsPerChain map[uint64]updateRequestWithMeta,
	configDigest cciptypes.Bytes32,
	homeFMap map[cciptypes.ChainSelector]int,
	rmnNodeInfo map[rmntypes.NodeID]rmntypes.HomeNodeInfo,
) ([]rmnSignedObservationWithMeta, map[uint64]updateRequestWithMeta, error) {
	requestedNodes := make(map[uint64]mapset.Set[rmntypes.NodeID])                   // sourceChain -> requested rmnNodeIDs
	requestsPerNode := make(map[rmntypes.NodeID][]*rmnpb.FixedDestLaneUpdateRequest) // grouped requests for each node

	// Send to every RMN node all the lane update requests it supports until all chains have a sufficient amount.
	// of initial observers. Upon timer expiration, additional requests are sent to the rest of the RMN nodes.

	chainsWithEnoughRequests := mapset.NewSet[uint64]()
	for nodeID := range rmnNodeInfo {
		if chainsWithEnoughRequests.Cardinality() == len(updateRequestsPerChain) {
			break // We have enough initial observers for all source chains.
		}

		for sourceChain, updateRequest := range updateRequestsPerChain {
			// if this node cannot support the source chain, skip it
			if !updateRequest.RmnNodes.Contains(nodeID) {
				continue
			}

			// add the node as a requested observer for this source chain
			if _, ok := requestedNodes[sourceChain]; !ok {
				requestedNodes[sourceChain] = mapset.NewSet[rmntypes.NodeID]()
			}
			requestedNodes[sourceChain].Add(nodeID)
			requestsPerNode[nodeID] = append(requestsPerNode[nodeID], updateRequest.Data)

			// if we already have enough requests for this source chain, mark it
			homeChainF, exist := homeFMap[cciptypes.ChainSelector(sourceChain)]
			if !exist {
				lggr.Errorw("no home F for chain", "chain", sourceChain)
				continue
			}
			if consensus.GteFPlusOne(homeChainF, requestedNodes[sourceChain].Cardinality()) {
				chainsWithEnoughRequests.Add(sourceChain)
			}
		}
	}

	requestIDs := c.sendObservationRequests(lggr, destChain, requestsPerNode, rmnNodeInfo)

	signedObservations, err := c.listenForRmnObservationResponses(
		ctx, lggr, destChain, requestIDs, updateRequestsPerChain, requestedNodes, configDigest, homeFMap, rmnNodeInfo)
	if err != nil && !errors.Is(err, ErrInsufficientObservationResponses) {
		return nil, nil, fmt.Errorf("listen for rmn observation responses: %w", err)
	}

	laneUpdatesToMakeProgressWith := make(map[uint64]updateRequestWithMeta)

	chainsToMakeProgressWith := chainsWithSufficientObservationResponses(
		lggr, updateRequestsPerChain, signedObservations, homeFMap)
	if chainsToMakeProgressWith.Cardinality() == 0 {
		return nil, nil, ErrAllChainsNotReady
	}

	for chain := range chainsToMakeProgressWith.Iter() {
		laneUpdatesToMakeProgressWith[uint64(chain)] = updateRequestsPerChain[uint64(chain)]
	}

	// Sanity check that we got enough signed observations for every source chain we want to make progress with.
	chainsWithSufficientResps := chainsWithSufficientObservationResponses(
		lggr, laneUpdatesToMakeProgressWith, signedObservations, homeFMap)
	if chainsWithSufficientResps.Cardinality() != len(laneUpdatesToMakeProgressWith) {
		lggr.Errorw("not enough signed observations after sanity check",
			"chainsWithSufficientResps", chainsWithSufficientResps.String(),
			"laneUpdatesToMakeProgressWith", laneUpdatesToMakeProgressWith,
			"chainsToMakeProgressWith", chainsToMakeProgressWith.String())
		return nil, nil, fmt.Errorf("not enough signed observations after sanity check")
	}

	return signedObservations, laneUpdatesToMakeProgressWith, nil
}

// sendObservationRequests sends observation requests to the RMN nodes.
// If a specific request fails, it is logged and not included in the returned requestIDs mapping.
// Returns a set of inflight requests (a mapping from request ID to InFlightRequest)
func (c *controller) sendObservationRequests(
	lggr logger.Logger,
	destChain *rmnpb.LaneDest,
	requestsPerNode map[rmntypes.NodeID][]*rmnpb.FixedDestLaneUpdateRequest,
	rmnNodeInfo map[rmntypes.NodeID]rmntypes.HomeNodeInfo,
) (inFlightRequests map[uint64]InFlightRmnRequest) {
	inFlightRequests = make(map[uint64]InFlightRmnRequest)

	for nodeID, requests := range requestsPerNode {
		sort.Slice(requests, func(i, j int) bool {
			return requests[i].LaneSource.SourceChainSelector < requests[j].LaneSource.SourceChainSelector
		})

		fixedDestLaneUpdateRequests := make([]*rmnpb.FixedDestLaneUpdateRequest, 0)

		// convert OnRamp address from 32bytes (abi encoded) to 20bytes (evm address) before sending the request
		// todo: make this part chain-agnostic
		for _, request := range requests {
			if destChain.DestChainSelector == request.LaneSource.SourceChainSelector {
				// explicit check that the destination chain should never be in any of the requested updates.
				lggr.Errorw("skipping observation request, request source chain is equal to the destination",
					"request", request)
				continue
			}

			fixedDestLaneUpdateRequests = append(fixedDestLaneUpdateRequests, &rmnpb.FixedDestLaneUpdateRequest{
				LaneSource: &rmnpb.LaneSource{
					SourceChainSelector: request.LaneSource.SourceChainSelector,
					// TODO check if we can remove the call for keepNRightBytes
					// https://github.com/smartcontractkit/chainlink-ccip/pull/647/files#r1966165319
					OnrampAddress: typconv.KeepNRightBytes(request.LaneSource.OnrampAddress, 20),
				},
				ClosedInterval: request.ClosedInterval,
			})
		}

		req := &rmnpb.Request{
			RequestId: newRequestID(lggr),
			Request: &rmnpb.Request_ObservationRequest{
				ObservationRequest: &rmnpb.ObservationRequest{
					LaneDest:                    destChain,
					FixedDestLaneUpdateRequests: fixedDestLaneUpdateRequests,
				},
			},
		}

		rmnNode, exists := rmnNodeInfo[nodeID]
		if !exists {
			lggr.Errorw("rmn node info not found skipped", "node", nodeID)
			continue
		}

		lggr := logger.With(lggr, "node", nodeID, "requestID", req.RequestId)
		lggr.Infow("sending observation request", "laneUpdateRequests", requests)
		if err := c.marshalAndSend(req, rmnNode); err != nil {
			c.metricsReporter.TrackRmnRequest(RmnMethodObservation, 0, uint64(nodeID), "failed_to_send_request")
			lggr.Errorw("failed to send observation request", "err", err)
			continue
		}

		inFlightRequests[req.RequestId] = NewInFlightRmnRequest(uint64(nodeID))
	}

	return inFlightRequests
}

// listenForRmnObservationResponses listens for the RMN observation responses.
// It waits for the responses until all the requests are finished or until the context is done.
// It is responsible for sending additional observation requests if the initial requests timeout.
//
//nolint:gocyclo // todo
func (c *controller) listenForRmnObservationResponses(
	ctx context.Context,
	lggr logger.Logger,
	destChain *rmnpb.LaneDest,
	inFlightRequests map[uint64]InFlightRmnRequest, // map from requestIDs to request info
	lursPerChain map[uint64]updateRequestWithMeta,
	requestedNodes map[uint64]mapset.Set[rmntypes.NodeID],
	configDigest cciptypes.Bytes32,
	homeFMap map[cciptypes.ChainSelector]int,
	rmnNodeInfo map[rmntypes.NodeID]rmntypes.HomeNodeInfo,
) ([]rmnSignedObservationWithMeta, error) {
	lggr.Infow("listening for RMN observation responses", "requestIDs",
		mapset.NewSetFromMapKeys(inFlightRequests).String())

	finishedRequestIDs := mapset.NewSet[uint64]()
	rmnObservationResponses := make([]rmnSignedObservationWithMeta, 0)

	initialObservationRequestTimer := time.NewTimer(c.observationsInitialRequestTimerDuration)
	timerExpired := false

	defer initialObservationRequestTimer.Stop()
	for {
		select {
		case resp := <-c.peerClient.Recv():
			parsedResp, latency, err := c.parseResponse(&resp, inFlightRequests, finishedRequestIDs)
			if err != nil {
				lggr.Debugw("skipping an unexpected RMN response", "err", err)
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
				c.metricsReporter.TrackRmnRequest(RmnMethodObservation, latency, uint64(resp.RMNNodeID), "invalid_response")
				lggr.Warnw("skipping an invalid RMN observation response", "err", err)
				initialObservationRequestTimer.Reset(0) // immediately schedule the additional requests
			} else {
				c.metricsReporter.TrackRmnRequest(RmnMethodObservation, latency, uint64(resp.RMNNodeID), "")
				rmnObservationResponses = append(rmnObservationResponses, rmnSignedObservationWithMeta{
					SignedObservation: parsedResp.GetSignedObservation(),
					RMNNodeID:         resp.RMNNodeID,
				})
			}

			chainsWithSufficientResponses := chainsWithSufficientObservationResponses(
				lggr,
				lursPerChain,
				rmnObservationResponses,
				homeFMap)

			if chainsWithSufficientResponses.Cardinality() == len(lursPerChain) {
				lggr.Info("all chains have enough observation responses with matching roots")
				return rmnObservationResponses, nil
			}

			if chainsWithSufficientResponses.Cardinality() > len(lursPerChain) {
				lggr.Errorw("internal bug, reported chains with sufficient responses more than expected",
					"chainsWithSufficientResponses", chainsWithSufficientResponses.String(),
					"lursPerChain", lursPerChain)
				return nil, fmt.Errorf("reported chains with sufficient responses are more than expected")
			}

			// We got all the responses we were waiting for, but they are not sufficient for all chains.
			if timerExpired && finishedRequestIDs.Equal(mapset.NewSetFromMapKeys(inFlightRequests)) {
				lggr.Warnw("observation requests were finished, but results are not sufficient")
				return rmnObservationResponses, ErrInsufficientObservationResponses
			}
		case <-initialObservationRequestTimer.C:
			if timerExpired {
				continue
			}
			timerExpired = true

			lggr.Warn("sending additional RMN observation requests")
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
			newInFlightRequests := c.sendObservationRequests(lggr, destChain, requestsPerNode, rmnNodeInfo)
			maps.Copy(inFlightRequests, newInFlightRequests)
		case <-ctx.Done():
			// Report metrics for requests we never received responses for
			for requestID, requestInfo := range inFlightRequests {
				if !finishedRequestIDs.Contains(requestID) {
					c.metricsReporter.TrackRmnRequest(RmnMethodObservation, requestInfo.Latency(),
						requestInfo.nodeID, "timeout")
					lggr.Warnw("Timed out waiting for an observation response from RMN",
						"requestID", requestID, "nodeID", requestInfo.nodeID, "latency", requestInfo.Latency())
				}
			}
			return nil, ErrTimeout
		}
	}
}

// chainsWithSufficientObservationResponses checks for which chains we got enough observation responses.
// Enough meaning that we got at least F+1 observing the same merkle root for a target chain.
func chainsWithSufficientObservationResponses(
	lggr logger.Logger,
	updateRequests map[uint64]updateRequestWithMeta,
	rmnObservationResponses []rmnSignedObservationWithMeta,
	homeFMap map[cciptypes.ChainSelector]int,
) mapset.Set[cciptypes.ChainSelector] {
	merkleRootsCount := make(map[uint64]map[cciptypes.Bytes32]int)
	for _, signedObs := range rmnObservationResponses {
		for _, lu := range signedObs.SignedObservation.Observation.FixedDestLaneUpdates {
			if _, exists := merkleRootsCount[lu.LaneSource.SourceChainSelector]; !exists {
				merkleRootsCount[lu.LaneSource.SourceChainSelector] = make(map[cciptypes.Bytes32]int)
			}
			merkleRootsCount[lu.LaneSource.SourceChainSelector][cciptypes.Bytes32(lu.Root)]++
		}
	}

	resultChains := mapset.NewSet[cciptypes.ChainSelector]()

	for sourceChain := range updateRequests {
		countsPerRoot, ok := merkleRootsCount[sourceChain]
		if !ok || len(countsPerRoot) == 0 {
			lggr.Infow("chain skipped, zero roots observed", "chain", sourceChain)
			continue
		}

		fObserve, exists := homeFMap[cciptypes.ChainSelector(sourceChain)]
		if !exists {
			lggr.Errorw("no F for chain", "chain", sourceChain)
			continue
		}

		mostVotedRootVotes := slices.Max(maps.Values(countsPerRoot))
		if consensus.LtFPlusOne(fObserve, mostVotedRootVotes) {
			lggr.Infow("chain skipped, maximally voted root doesn't have enough votes",
				"chain", sourceChain,
				"fObserve", fObserve,
				"mostVotedRootVotes", mostVotedRootVotes)
			continue
		}

		lggr.Infow("chain has enough observed roots",
			"chain", sourceChain, "mostVotedRootVotes", mostVotedRootVotes, "fObserve", fObserve)
		resultChains.Add(cciptypes.ChainSelector(sourceChain))
	}

	return resultChains
}

//nolint:gocyclo // todo
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

	rmnNode, err := c.getHomeNodeByID(rmnHomeConfigDigest, rmnNodeID)
	if err != nil {
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

	seenSourceChainSelectors := mapset.NewSet[uint64]()

	for _, signedObsLu := range signedObs.Observation.FixedDestLaneUpdates {
		updateReq, exists := lurs[signedObsLu.LaneSource.SourceChainSelector]
		if !exists {
			return fmt.Errorf("unexpected source chain selector %d", signedObsLu.LaneSource.SourceChainSelector)
		}

		if seenSourceChainSelectors.Contains(signedObsLu.LaneSource.SourceChainSelector) {
			return fmt.Errorf("duplicate source chain %d", signedObsLu.LaneSource.SourceChainSelector)
		}
		seenSourceChainSelectors.Add(signedObsLu.LaneSource.SourceChainSelector)

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

		// todo: The original updateReq contains abi encoded onRamp address, the one in the RMN response
		// is 20 bytes evm address. This is chain specific and should be handled in a chain specific way.

		// TODO check if we can remove the call for keepNRightBytes
		// https://github.com/smartcontractkit/chainlink-ccip/pull/647/files#r1966165319
		expOnRampAddress := typconv.KeepNRightBytes(updateReq.Data.LaneSource.OnrampAddress, 20)
		if !bytes.Equal(expOnRampAddress, signedObsLu.LaneSource.OnrampAddress) {
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
	lggr logger.Logger,
	destChain *rmnpb.LaneDest,
	rmnSignedObservations []rmnSignedObservationWithMeta,
	updatesPerChain map[uint64]updateRequestWithMeta,
	rmnRemoteCfg cciptypes.RemoteConfig,
	rmnNodeInfo map[rmntypes.NodeID]rmntypes.HomeNodeInfo,
) (*ReportSignatures, error) {
	// At this point we might have multiple signedObservations for different nodes but never for the same source chain
	// from the same node.
	//
	// e.g.
	// The following nodes support the following chains and F=1:
	// node1: [1]	node2:[1,2,3]	node3:[1,2,3]
	//
	// node1: getObservations(1)          ->             never_responds
	// node2: getObservations(1,2,3)      -> responds_ok
	// node3: getObservations(2,3)        ->     responds_ok
	// node3: getObservations(1)          ->                         responds_ok (requested after node1 timeout)

	// At this point it is also possible that the signed observations contain
	// different roots for the same source chain and interval.

	homeChainF, err := c.rmnHomeReader.GetFObserve(rmnRemoteCfg.ConfigDigest)
	if err != nil {
		return nil, fmt.Errorf("get home reader F: %w", err)
	}

	rootsPerChain, err := selectRoots(rmnSignedObservations, homeChainF)
	if err != nil {
		return nil, fmt.Errorf("get most voted roots from observations: %w", err)
	}

	fixedDestLaneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0)
	for sourceChain, updateReq := range updatesPerChain {
		selectedRoot, ok := rootsPerChain[cciptypes.ChainSelector(sourceChain)]
		if !ok {
			return nil, fmt.Errorf("no most voted root for source chain %d", sourceChain)
		}
		fixedDestLaneUpdates = append(fixedDestLaneUpdates, &rmnpb.FixedDestLaneUpdate{
			LaneSource:     updateReq.Data.LaneSource,
			ClosedInterval: updateReq.Data.ClosedInterval,
			Root:           selectedRoot[:],
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

	rmnReport := cciptypes.NewRMNReport(
		rmnRemoteCfg.RmnReportVersion,
		cciptypes.NewBigIntFromInt64(int64(destChainInfo.EvmChainID)),
		cciptypes.ChainSelector(destChain.DestChainSelector),
		rmnRemoteCfg.ContractAddress,
		destChain.OfframpAddress,
		rmnRemoteCfg.ConfigDigest,
		laneUpdates,
	)

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
	remoteF := int(rmnRemoteCfg.FSign)
	signers := rmnRemoteCfg.Signers
	inFlightRequests, signersRequested, err := c.sendReportSignatureRequest(
		lggr,
		reportSigReq,
		signers,
		remoteF,
		rmnNodeInfo)
	if err != nil {
		return nil, fmt.Errorf("send report signature request: %w", err)
	}

	ecdsaSignatures, err := c.listenForRmnReportSignatures(
		ctx,
		lggr,
		inFlightRequests,
		rmnReport,
		reportSigReq,
		signersRequested,
		signers,
		remoteF,
		rmnNodeInfo)
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

// selectRoots selects the roots from the signed observations.
// If there are more than one valid roots based on the provided F it returns an error.
func selectRoots(
	observations []rmnSignedObservationWithMeta,
	homeFMap map[cciptypes.ChainSelector]int,
) (map[cciptypes.ChainSelector]cciptypes.Bytes32, error) {
	votesPerRoot := make(map[cciptypes.ChainSelector]map[cciptypes.Bytes32]int)
	for _, so := range observations {
		for _, lu := range so.SignedObservation.Observation.FixedDestLaneUpdates {
			if _, exists := votesPerRoot[cciptypes.ChainSelector(lu.LaneSource.SourceChainSelector)]; !exists {
				votesPerRoot[cciptypes.ChainSelector(lu.LaneSource.SourceChainSelector)] = make(map[cciptypes.Bytes32]int)
			}
			votesPerRoot[cciptypes.ChainSelector(lu.LaneSource.SourceChainSelector)][cciptypes.Bytes32(lu.Root)]++
		}
	}

	selectedRoots := make(map[cciptypes.ChainSelector]cciptypes.Bytes32)
	for chain, votes := range votesPerRoot {
		homeF, exists := homeFMap[chain]
		if !exists {
			return nil, fmt.Errorf("no home F for chain %d", chain)
		}

		var selectedRoot cciptypes.Bytes32

		for root, vote := range votes {
			if consensus.LtFPlusOne(homeF, vote) {
				continue
			}

			if !selectedRoot.IsEmpty() {
				return nil, fmt.Errorf("more than one valid root for chain %d", chain)
			}

			selectedRoot = root
		}

		if selectedRoot.IsEmpty() {
			return nil, fmt.Errorf("no valid root for chain %d", chain)
		}
		selectedRoots[chain] = selectedRoot
	}

	return selectedRoots, nil
}

// sendReportSignatureRequest sends the report signature request to #remoteF+1 random RMN nodes.
// If not enough requests were sent, it returns an error.
func (c *controller) sendReportSignatureRequest(
	lggr logger.Logger,
	reportSigReq *rmnpb.ReportSignatureRequest,
	remoteSigners []cciptypes.RemoteSignerInfo,
	remoteF int,
	rmnNodeInfo map[rmntypes.NodeID]rmntypes.HomeNodeInfo,
) (
	inFlightRequests map[uint64]InFlightRmnRequest, signersRequested mapset.Set[rmntypes.NodeID], err error) {
	inFlightRequests = make(map[uint64]InFlightRmnRequest)
	signersRequested = mapset.NewSet[rmntypes.NodeID]()

	// Send the report signature request to at least #remoteF+1
	for _, node := range randomShuffle(remoteSigners) {
		if consensus.GteFPlusOne(remoteF, len(inFlightRequests)) {
			break
		}

		req := &rmnpb.Request{
			RequestId: newRequestID(lggr),
			Request: &rmnpb.Request_ReportSignatureRequest{
				ReportSignatureRequest: reportSigReq,
			},
		}

		lggr.Infow("sending report signature request", "node", node.NodeIndex, "requestID", req.RequestId)

		rmnNode, exists := rmnNodeInfo[rmntypes.NodeID(node.NodeIndex)]
		if !exists {
			lggr.Errorw("rmn node info not found skipped", "node", node.NodeIndex)
			continue
		}

		err := c.marshalAndSend(req, rmnNode)
		if err != nil {
			lggr.Warnw("failed to send report signature request", "node", node.NodeIndex, "err", err)
			c.metricsReporter.TrackRmnRequest(RmnMethodReportSignature, 0, node.NodeIndex, "failed_to_send_request")
			continue
		}

		inFlightRequests[req.RequestId] = NewInFlightRmnRequest(node.NodeIndex)
		signersRequested.Add(rmntypes.NodeID(node.NodeIndex))
	}

	if consensus.LtFPlusOne(remoteF, len(inFlightRequests)) {
		return inFlightRequests, signersRequested, fmt.Errorf("not able to send to enough report signers")
	}
	return inFlightRequests, signersRequested, nil
}

// reportSigWithSignerAddress is a helper struct to store the report signature and
// the address of the private key that signed it.
type reportSigWithSignerAddress struct {
	reportSig     *rmnpb.ReportSignature
	signerAddress cciptypes.UnknownAddress
}

//nolint:gocyclo // todo
func (c *controller) listenForRmnReportSignatures(
	ctx context.Context,
	lggr logger.Logger,
	inFlightRequests map[uint64]InFlightRmnRequest,
	rmnReport cciptypes.RMNReport,
	reportSigReq *rmnpb.ReportSignatureRequest,
	signersRequested mapset.Set[rmntypes.NodeID],
	signers []cciptypes.RemoteSignerInfo,
	remoteF int,
	rmnNodeInfo map[rmntypes.NodeID]rmntypes.HomeNodeInfo,
) ([]*rmnpb.EcdsaSignature, error) {
	tReportsInitialRequest := time.NewTimer(c.reportsInitialRequestTimerDuration)
	timerExpired := false

	reportSigs := make([]reportSigWithSignerAddress, 0)
	finishedRequests := mapset.NewSet[uint64]()
	inFlightRequests = maps.Clone(inFlightRequests)
	requestIDs := mapset.NewSetFromMapKeys(inFlightRequests)
	lggr.Infof("waiting for report signatures, requestIDs: %s", requestIDs.String())

	for {
		select {
		case resp := <-c.peerClient.Recv():
			requestIDs = mapset.NewSetFromMapKeys(inFlightRequests)
			responseTyp, latency, err := c.parseResponse(&resp, inFlightRequests, finishedRequests)
			if err != nil {
				lggr.Infow("failed to parse RMN signature response", "err", err)
				continue
			}
			finishedRequests.Add(responseTyp.RequestId)

			reportSig, err := c.validateReportSigResponse(ctx, responseTyp, resp.RMNNodeID, signers, rmnReport)

			if err != nil {
				c.metricsReporter.TrackRmnRequest(RmnMethodReportSignature, latency, uint64(resp.RMNNodeID), "invalid_response")
				lggr.Warnw("skipping an invalid RMN report signature response", "err", err)
				tReportsInitialRequest.Reset(0) // schedule additional requests if any
			} else {
				c.metricsReporter.TrackRmnRequest(RmnMethodReportSignature, latency, uint64(resp.RMNNodeID), "")
				lggr.Infow("received valid report signature", "node", resp.RMNNodeID, "requestID", responseTyp.RequestId)
				reportSigs = append(reportSigs, *reportSig)
			}

			if consensus.GteFPlusOne(remoteF, len(reportSigs)) {
				lggr.Infof("got enough RMN report signatures")
				return sortAndParseReportSigs(reportSigs), nil
			}

			if timerExpired && finishedRequests.Equal(requestIDs) {
				lggr.Warn("report signature requests were finished, but results are not sufficient")
				return nil, ErrInsufficientSignatureResponses
			}
		case <-tReportsInitialRequest.C:
			if timerExpired {
				continue
			}
			timerExpired = true

			lggr.Warnw("sending additional RMN signature requests")

			for _, node := range randomShuffle(signers) {
				nodeIndex := node.NodeIndex
				if signersRequested.Contains(rmntypes.NodeID(nodeIndex)) {
					continue
				}
				req := &rmnpb.Request{
					RequestId: newRequestID(lggr),
					Request: &rmnpb.Request_ReportSignatureRequest{
						ReportSignatureRequest: reportSigReq,
					},
				}

				rmnNode, exists := rmnNodeInfo[rmntypes.NodeID(nodeIndex)]
				if !exists {
					lggr.Errorw("rmn node info not found skipped", "node", nodeIndex)
					continue
				}

				lggr.Infow("sending report signature request", "node", nodeIndex, "requestID", req.RequestId)
				if err := c.marshalAndSend(req, rmnNode); err != nil {
					lggr.Errorw("failed to send report signature request", "node", nodeIndex, "err", err)
					continue
				}
				inFlightRequests[req.RequestId] = NewInFlightRmnRequest(nodeIndex)
				signersRequested.Add(rmntypes.NodeID(nodeIndex))
			}
		case <-ctx.Done():
			// Report metrics for requests we never received responses for
			for requestID, requestInfo := range inFlightRequests {
				if !finishedRequests.Contains(requestID) {
					c.metricsReporter.TrackRmnRequest(RmnMethodReportSignature, requestInfo.Latency(),
						requestInfo.nodeID, "timeout")
					lggr.Warnw("Timed out waiting for a report signature response from RMN",
						"requestID", requestID, "nodeID", requestInfo.nodeID, "latency", requestInfo.Latency())
				}
			}
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
	signers []cciptypes.RemoteSignerInfo,
	rmnReport cciptypes.RMNReport,
) (*reportSigWithSignerAddress, error) {
	signerNode, err := c.getSignerNodeByID(signers, nodeID)
	if err != nil {
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
		[]cciptypes.UnknownAddress{signerNode.OnchainPublicKey},
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
	nodeID rmntypes.NodeID) (rmntypes.HomeNodeInfo, error) {
	rmnNodes, err := c.rmnHomeReader.GetRMNNodesInfo(configDigest)
	if err != nil {
		return rmntypes.HomeNodeInfo{}, err
	}
	for _, node := range rmnNodes {
		if node.ID == nodeID {
			return node, nil
		}
	}

	return rmntypes.HomeNodeInfo{}, ErrNotFound
}

func (c *controller) getSignerNodeByID(
	rmnNodes []cciptypes.RemoteSignerInfo,
	nodeID rmntypes.NodeID) (cciptypes.RemoteSignerInfo, error) {

	// Search for the node with the specified ID
	for _, node := range rmnNodes {
		if rmntypes.NodeID(node.NodeIndex) == nodeID {
			// Return the found node and nil error
			return node, nil
		}
	}

	// If the node was not found, return a "not found" error
	return cciptypes.RemoteSignerInfo{}, ErrNotFound
}

type updateRequestWithMeta struct {
	Data     *rmnpb.FixedDestLaneUpdateRequest
	RmnNodes mapset.Set[rmntypes.NodeID]
}

type rmnSignedObservationWithMeta struct {
	SignedObservation *rmnpb.SignedObservation
	RMNNodeID         rmntypes.NodeID
}

func (c *controller) marshalAndSend(req *rmnpb.Request, rmnNode rmntypes.HomeNodeInfo) error {
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("proto marshal RMN request: %w", err)
	}

	// TODO: include ctx in the func call to pass thru OCR seqNr.
	if err := c.peerClient.Send(rmnNode, reqBytes); err != nil {
		return fmt.Errorf("send rmn request: %w", err)
	}

	return nil
}

// parseResponse parses the response from the RMN and returns the response and the response latency
// Validates that the response is expected and not a duplicate.
func (c *controller) parseResponse(
	resp *PeerResponse, inFlightRequests map[uint64]InFlightRmnRequest, gotResponses mapset.Set[uint64],
) (*rmnpb.Response, float64, error) {
	responseTyp := &rmnpb.Response{}
	err := proto.Unmarshal(resp.Body, responseTyp)
	if err != nil {
		return nil, 0, fmt.Errorf("proto unmarshal: %w", err)
	}

	requestInfo, exists := inFlightRequests[responseTyp.RequestId]
	if !exists {
		requestIDs := mapset.NewSetFromMapKeys(inFlightRequests).String()
		return nil, 0, fmt.Errorf(
			"got an RMN response that we are not waiting for: %d (%s)", responseTyp.RequestId, requestIDs)
	}

	if requestInfo.nodeID != uint64(resp.RMNNodeID) {
		return nil, 0, fmt.Errorf(
			"got an RMN response from node %d, but the request was made to node %d, requestID: %d",
			resp.RMNNodeID, requestInfo.nodeID, responseTyp.RequestId)
	}

	if gotResponses.Contains(responseTyp.RequestId) {
		return nil, 0, fmt.Errorf("got a duplicate RMN response: %d", responseTyp.RequestId)
	}

	return responseTyp, requestInfo.Latency(), nil
}

func randomShuffle[T any](s []T) []T {
	ret := make([]T, len(s))
	for i, randIndex := range rand.Perm(len(s)) {
		ret[i] = s[randIndex]
	}
	return ret
}

// newRequestID generates a new unique request ID.
func newRequestID(lggr logger.Logger) uint64 {
	b := make([]byte, 8)
	_, err := crand.Read(b)
	if err != nil {
		// fallback to time-based id in the very rare scenario that the random number generator fails
		lggr.Warnw("failed to generate random request id, falling back to golang.org/x/exp/rand",
			"err", err,
		)
		//nolint:gosec // this is unlikely to occur and is not a security concern.
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		return rnd.Uint64()
	}
	randomUint64 := binary.LittleEndian.Uint64(b)
	return randomUint64
}

// InFlightRmnRequest tracks an in-flight RMN request, used for metrics
// RMN requests are sent async, so in order to compute request latencies we need to store a mapping from request ID to
// InFlightRmnRequest, and then reference when we get a response (or if we don't receive a response, we are able to
// mark the InFlightRmnRequests that have timed out, and mark the offending node IDs).
type InFlightRmnRequest struct {
	sent   time.Time // The time the request was sent
	nodeID uint64    // The node that the request was sent to
}

func NewInFlightRmnRequest(nodeID uint64) InFlightRmnRequest {
	return InFlightRmnRequest{
		sent:   time.Now(),
		nodeID: nodeID,
	}
}

func (i *InFlightRmnRequest) Latency() float64 {
	return float64(time.Since(i.sent).Milliseconds())
}
