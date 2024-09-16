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

type NodeID uint32

// client is the base RMN Client implementation.
type client struct {
	lggr            logger.Logger
	rmnCrypto       cciptypes.RMNCrypto
	rawRmnClient    RawRmnClient
	rmnCfg          Config
	ed25519Verifier ED25519Verifier

	observationsInitialRequestTimerDuration time.Duration
	reportsInitialRequestTimerDuration      time.Duration
}

// NewClient creates a new RMN Client to be used by the plugin.
func NewClient(
	lggr logger.Logger,
	rmnCrypto cciptypes.RMNCrypto,
	rawRmnClient RawRmnClient,
	rmnConfig Config,
	observationsInitialRequestTimerDuration time.Duration,
	reportsInitialRequestTimerDuration time.Duration,
) Client {
	return &client{
		lggr:            lggr,
		rmnCrypto:       rmnCrypto,
		rawRmnClient:    rawRmnClient,
		rmnCfg:          rmnConfig,
		ed25519Verifier: NewED25519Verifier(),

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
			rmnNodes: mapset.NewSet[NodeID](),
		}
		for _, node := range c.rmnCfg.Home.RmnNodes {
			if node.SupportedSourceChains.Contains(cciptypes.ChainSelector(updateReq.LaneSource.SourceChainSelector)) {
				updatesPerChain[updateReq.LaneSource.SourceChainSelector].rmnNodes.Add(node.ID)
			}
		}
	}

	// Filter out the lane update requests for chains without enough RMN nodes supporting them.
	for chain, l := range updatesPerChain {
		if l.rmnNodes.Cardinality() < c.rmnCfg.Remote.MinObservers {
			c.lggr.Warnw("not enough RMN nodes for chain",
				"chain", chain, "minObservers", c.rmnCfg.Remote.MinObservers)
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
	requestedNodes := make(map[uint64]mapset.Set[NodeID])                   // sourceChain -> requested rmnNodeIDs
	requestsPerNode := make(map[NodeID][]*rmnpb.FixedDestLaneUpdateRequest) // grouped requests for each node

	c.lggr.Infof("update requests per chain: %v", updateRequestsPerChain)

	// For each lane update request send an observation request to at most 'minObservers' number of rmn nodes.
	for sourceChain, updateRequest := range updateRequestsPerChain {
		requestedNodes[sourceChain] = mapset.NewSet[NodeID]()

		// At this point we assume at least minObservers RMN nodes are available for each source chain.
		for nodeID := range updateRequest.rmnNodes.Iter() {
			if requestedNodes[sourceChain].Cardinality() >= c.rmnCfg.Remote.MinObservers {
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
	requestsPerNode map[NodeID][]*rmnpb.FixedDestLaneUpdateRequest,
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
	requestedNodes map[uint64]mapset.Set[NodeID],
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
				err := c.validateSignedObservationResponse(
					resp.RMNNodeID, lursPerChain, signedObs, destChain, c.rmnCfg.Home.ConfigDigest)
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
					if count >= c.rmnCfg.Remote.MinObservers {
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
			requestsPerNode := make(map[NodeID][]*rmnpb.FixedDestLaneUpdateRequest)
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
	rmnNodeID NodeID,
	lurs map[uint64]updateRequestWithMeta,
	signedObs *rmnpb.SignedObservation,
	destChain *rmnpb.LaneDest,
	rmnHomeConfigDigest cciptypes.Bytes32,
) error {
	rmnNode, exists := c.getRmnNodeByID(rmnNodeID)
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

	if err := verifyObservationSignature(rmnNode, signedObs, c.ed25519Verifier); err != nil {
		return fmt.Errorf("failed to verify observation signature: %w", err)
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
			SignerNodeIndex:   uint32(resp.RMNNodeID),
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
			RmnRemoteContractAddress:    c.rmnCfg.Remote.ContractAddress,
			RmnHomeContractConfigDigest: c.rmnCfg.Home.ConfigDigest[:],
			LaneDest:                    destChain,
		},
		AttributedSignedObservations: sigObservations,
	}

	requestIDs := mapset.NewSet[uint64]()
	signersRequested := mapset.NewSet[NodeID]()

	// Send the report signature request to at least minSigners
	for _, node := range c.rmnCfg.Home.RmnNodes {
		if requestIDs.Cardinality() >= c.rmnCfg.Remote.MinSigners {
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
		var err error
		for merkleRootStr, count := range counts {
			if count >= c.rmnCfg.Remote.MinSigners {
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

	ecdsaSignatures, err := c.listenForRmnReportSignatures(
		ctx, requestIDs, fixedDestLaneUpdates, reportSigReq, signersRequested, destChain)
	if err != nil {
		return nil, fmt.Errorf("listen for rmn report signatures: %w", err)
	}

	return &ReportSignatures{
		Signatures:  ecdsaSignatures,
		LaneUpdates: fixedDestLaneUpdates,
	}, nil
}

// reportSigWithNodeID is a helper struct to store the report signature and the node ID that provided it.
type reportSigWithNodeID struct {
	reportSig     *rmnpb.ReportSignature
	signerAddress cciptypes.Bytes
}

func (c *client) listenForRmnReportSignatures(
	ctx context.Context,
	requestIDs mapset.Set[uint64],
	fixedDestLaneUpdates []*rmnpb.FixedDestLaneUpdate,
	reportSigReq *rmnpb.ReportSignatureRequest,
	signersRequested mapset.Set[NodeID],
	destChain *rmnpb.LaneDest,
) ([]*rmnpb.EcdsaSignature, error) {
	tReportsInitialRequest := time.NewTimer(c.reportsInitialRequestTimerDuration)
	reportSigs := make([]reportSigWithNodeID, 0)
	finishedRequests := mapset.NewSet[uint64]()
	respChan := c.rawRmnClient.Recv()
	requestIDs = requestIDs.Clone()
	c.lggr.Infof("Waiting for report signatures, requestIDs: %s", requestIDs.String())

	for {
		select {
		case resp := <-respChan:
			responseTyp, err := c.parseResponse(&resp, requestIDs, finishedRequests)
			if err != nil {
				c.lggr.Infow("failed to parse RMN signature response", "err", err)
				continue
			}

			reportSig := responseTyp.GetReportSignature()
			if reportSig == nil {
				c.lggr.Infof("RMN node returned an unexpected type of response: %+v", responseTyp.Response)
				continue
			}

			rmnNode, exists := c.getRmnNodeByID(resp.RMNNodeID)
			if !exists {
				c.lggr.Warnw("rmn node that appears in report signature does not exist", "nodeID", resp.RMNNodeID)
				continue
			}

			ch, exists := chainsel.ChainBySelector(destChain.DestChainSelector)
			if !exists {
				c.lggr.Warnw("unknown chain selector", "chainSelector", destChain.DestChainSelector)
				continue
			}

			laneUpdates, err := NewLaneUpdatesFromPB(fixedDestLaneUpdates)
			if err != nil {
				return nil, fmt.Errorf("failed to convert lane updates from protobuf: %w", err)
			}

			rmnReport := cciptypes.RMNReport{
				ReportVersion:               c.rmnCfg.Home.RmnReportVersion,
				DestChainID:                 cciptypes.NewBigIntFromInt64(int64(ch.EvmChainID)),
				DestChainSelector:           cciptypes.ChainSelector(destChain.DestChainSelector),
				RmnRemoteContractAddress:    c.rmnCfg.Remote.ContractAddress,
				OfframpAddress:              destChain.OfframpAddress,
				RmnHomeContractConfigDigest: c.rmnCfg.Home.ConfigDigest,
				LaneUpdates:                 laneUpdates,
			}

			sig, err := NewECDSASigFromPB(reportSig.Signature)
			if err != nil {
				return nil, fmt.Errorf("failed to convert signature from protobuf: %w", err)
			}

			err = c.rmnCrypto.VerifyReportSignatures(
				ctx,
				[]cciptypes.RMNECDSASignature{*sig},
				rmnReport,
				[]cciptypes.Bytes{rmnNode.SignReportsAddress},
			)
			if err != nil {
				c.lggr.Warnw("failed to verify report signature", "err", err)
				continue
			}

			c.lggr.Infow("received report signature", "node", resp.RMNNodeID, "requestID", responseTyp.RequestId)
			reportSigs = append(reportSigs, reportSigWithNodeID{
				reportSig:     reportSig,
				signerAddress: rmnNode.SignReportsAddress,
			})

			if len(reportSigs) >= c.rmnCfg.Remote.MinSigners {
				// Sort report sigs by signer address
				// Similar to RMNRemote.verify (signatures must be sorted in ascending order by signer address).
				sort.Slice(reportSigs, func(i, j int) bool {
					return reportSigs[i].signerAddress.String() < reportSigs[j].signerAddress.String()
				})

				ecdsaSigs := make([]*rmnpb.EcdsaSignature, 0)
				for _, rs := range reportSigs {
					ecdsaSigs = append(ecdsaSigs, rs.reportSig.Signature)
				}
				return ecdsaSigs, nil
			}
		case <-tReportsInitialRequest.C:
			c.lggr.Warnw("initial report signatures request timer expired, sending additional requests")
			// Send to the rest of the signers
			for _, node := range randomShuffle(c.rmnCfg.Home.RmnNodes) {
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
		if count < c.rmnCfg.Remote.MinObservers {
			return fmt.Errorf("not enough observations for chain=%d count=%d minObservers=%d",
				chain, count, c.rmnCfg.Remote.MinObservers)
		}
	}

	return nil
}

func (c *client) getRmnNodeByID(nodeID NodeID) (RMNNodeInfo, bool) {
	for _, node := range c.rmnCfg.Home.RmnNodes {
		if node.ID == nodeID {
			return node, true
		}
	}
	return RMNNodeInfo{}, false
}

type RawRmnClient interface {
	Send(rmnNodeID NodeID, request []byte) error
	// Recv returns a channel which can be used to listen on for
	// responses by all RMN nodes. This is expected to be monitored
	// by the plugin in order to get RMN responses.
	Recv() <-chan RawRmnResponse
}

type RawRmnResponse struct {
	RMNNodeID NodeID
	Body      []byte // pb
}

type updateRequestWithMeta struct {
	data     *rmnpb.FixedDestLaneUpdateRequest
	rmnNodes mapset.Set[NodeID]
}

type rmnSignedObservationWithMeta struct {
	SignedObservation *rmnpb.SignedObservation
	RMNNodeID         NodeID
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
