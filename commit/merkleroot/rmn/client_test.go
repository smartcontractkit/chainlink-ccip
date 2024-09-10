package rmn

import (
	"sync"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

var (
	chainS1       = chainsel.TEST_90000002.Selector
	chainS1OnRamp = []byte{0, 0xf, 0xf, 0xa, 0x0}

	chainS2       = chainsel.TEST_90000003.Selector
	chainS2OnRamp = []byte{0, 0xf, 0xf, 0xa, 0x1}

	chainD1        = chainsel.TEST_90000004.Selector
	chainD1OffRamp = []byte{0, 0xf, 0xf, 0xa, 0x2}
)

type testSetup struct {
	name           string
	t              *testing.T
	rmnClient      *PBClient
	updateRequests []*rmnpb.FixedDestLaneUpdateRequest
}

func TestPBClientComputeReportSignatures(t *testing.T) {
	resChan := make(chan RawRmnResponse, 100)
	mockClient := newMockRawRmnClient(resChan)
	lggr := logger.Test(t)
	ctx := tests.Context(t)

	newTestSetup := func(t *testing.T) testSetup {
		client := &PBClient{
			lggr:         lggr,
			rawRmnClient: mockClient,
			rmnNodes: []RMNNodeInfo{
				{ID: 1, SupportedSourceChains: mapset.NewSet(chainS1, chainS2), IsSigner: true},
				{ID: 2, SupportedSourceChains: mapset.NewSet(chainS1, chainS2), IsSigner: true},
				{ID: 3, SupportedSourceChains: mapset.NewSet(chainS1, chainS2), IsSigner: true},
				{ID: 4, SupportedSourceChains: mapset.NewSet(chainS1, chainS2), IsSigner: true},
			},
			rmnRemoteAddress:                        []byte{1, 2, 3},
			rmnHomeConfigDigest:                     []byte{0xc, 0x0, 0xf},
			minObservers:                            2,
			minSigners:                              2,
			observationsInitialRequestTimerDuration: 30 * time.Hour,
			reportsInitialRequestTimerDuration:      30 * time.Hour,
		}

		updateRequests := []*rmnpb.FixedDestLaneUpdateRequest{
			{
				LaneSource:     &rmnpb.LaneSource{SourceChainSelector: chainS1, OnrampAddress: chainS1OnRamp},
				ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 10, MaxMsgNr: 20},
			},
			{
				LaneSource:     &rmnpb.LaneSource{SourceChainSelector: chainS2, OnrampAddress: chainS2OnRamp},
				ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 100, MaxMsgNr: 110},
			},
		}

		return testSetup{
			name:           t.Name(),
			t:              t,
			rmnClient:      client,
			updateRequests: updateRequests,
		}
	}

	destChain := &rmnpb.LaneDest{
		DestChainSelector: chainD1,
		OfframpAddress:    chainD1OffRamp,
	}

	t.Run("empty lane update request", func(t *testing.T) {
		ts := newTestSetup(t)
		_, err := ts.rmnClient.ComputeReportSignatures(
			ctx,
			destChain,
			[]*rmnpb.FixedDestLaneUpdateRequest{},
		)
		assert.Error(t, err, ErrNothingToDo)
	})

	t.Run("happy path no retries", func(t *testing.T) {
		ts := newTestSetup(t)
		go func() {
			requestIDs, requestedChains := ts.waitForObservationRequestsToBeSent(mockClient, ts.rmnClient.minObservers)

			ts.nodesRespondToTheObservationRequests(
				mockClient, requestIDs, requestedChains, ts.rmnClient.rmnHomeConfigDigest, destChain)

			requestIDs = ts.waitForInitialReportSignatureRequestsToBeSent(
				t, mockClient, ts.rmnClient.minSigners, ts.rmnClient.minObservers)

			ts.nodesRespondToTheSignatureRequests(mockClient, requestIDs)
		}()

		sigs, err := ts.rmnClient.ComputeReportSignatures(
			ctx,
			destChain,
			ts.updateRequests,
		)
		assert.NoError(t, err)
		assert.Len(t, sigs.LaneUpdates, len(ts.updateRequests))
		assert.Len(t, sigs.Signatures, ts.rmnClient.minSigners)
	})

	t.Run("happy path with retries", func(t *testing.T) {
		ts := newTestSetup(t)

		ts.rmnClient.observationsInitialRequestTimerDuration = time.Nanosecond
		ts.rmnClient.reportsInitialRequestTimerDuration = time.Nanosecond

		go func() {
			requestIDs, requestedChains := ts.waitForObservationRequestsToBeSent(mockClient, ts.rmnClient.minObservers)

			// requests should be sent to at least two nodes
			assert.Greater(t, len(requestIDs), ts.rmnClient.minObservers)
			assert.Len(t, requestedChains, len(ts.rmnClient.rmnNodes))

			ts.nodesRespondToTheObservationRequests(
				mockClient, requestIDs, requestedChains, ts.rmnClient.rmnHomeConfigDigest, destChain)
			time.Sleep(time.Millisecond)

			requestIDs = ts.waitForInitialReportSignatureRequestsToBeSent(
				t, mockClient, ts.rmnClient.minSigners, ts.rmnClient.minObservers)
			// requests should be sent to all nodes, since we hit the timer timeout
			assert.Len(t, requestIDs, len(ts.rmnClient.rmnNodes))

			ts.nodesRespondToTheSignatureRequests(mockClient, requestIDs)
		}()

		sigs, err := ts.rmnClient.ComputeReportSignatures(
			ctx,
			destChain,
			ts.updateRequests,
		)
		assert.NoError(t, err)
		assert.Len(t, sigs.LaneUpdates, len(ts.updateRequests))
		assert.Len(t, sigs.Signatures, ts.rmnClient.minSigners)
	})
}

func (ts *testSetup) waitForObservationRequestsToBeSent(
	rmnClient *mockRawRmnClient,
	minObservers int,
) (map[uint32]uint64, map[uint32]mapset.Set[uint64]) {
	requestIDs := make(map[uint32]uint64)
	requestedChains := make(map[uint32]mapset.Set[uint64])

	for {
		time.Sleep(time.Millisecond)
		recvReqs := rmnClient.getReceivedRequests()
		requestsPerChain := map[uint64]int{}
		for _, reqs := range recvReqs {
			for _, req := range reqs {
				for _, src := range req.GetObservationRequest().GetFixedDestLaneUpdateRequests() {
					requestsPerChain[src.LaneSource.SourceChainSelector]++
				}
			}
		}
		if requestsPerChain[chainS1] >= minObservers && requestsPerChain[chainS2] >= minObservers {
			for nodeID, reqs := range recvReqs {
				requestIDs[nodeID] = reqs[0].RequestId
				requestedChains[nodeID] = mapset.NewSet[uint64]()
				for _, src := range reqs[0].GetObservationRequest().GetFixedDestLaneUpdateRequests() {
					requestedChains[nodeID].Add(src.LaneSource.SourceChainSelector)
				}

			}

			rmnClient.resetReceivedRequests()
			break // all nodes have received the observation requests
		}
		continue
	}
	ts.t.Logf("nodes received the observation requests: %v", requestIDs)
	return requestIDs, requestedChains
}

func (ts *testSetup) nodesRespondToTheObservationRequests(
	rmnClient *mockRawRmnClient,
	requestIDs map[uint32]uint64,
	requestedChains map[uint32]mapset.Set[uint64],
	rmnHomeConfigDigest []byte,
	destChain *rmnpb.LaneDest,
) {
	allLaneUpdates := []*rmnpb.FixedDestLaneUpdate{
		{
			LaneSource:     &rmnpb.LaneSource{SourceChainSelector: chainS1, OnrampAddress: chainS1OnRamp},
			ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 10, MaxMsgNr: 20},
			Root:           []byte{1, 10, 20},
		},
		{
			LaneSource:     &rmnpb.LaneSource{SourceChainSelector: chainS2, OnrampAddress: chainS2OnRamp},
			ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 100, MaxMsgNr: 110},
			Root:           []byte{2, 100, 110},
		},
	}

	for nodeID, requestID := range requestIDs {
		laneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0)
		for _, laneUpdate := range allLaneUpdates {
			if requestedChains[nodeID].Contains(laneUpdate.LaneSource.SourceChainSelector) {
				laneUpdates = append(laneUpdates, laneUpdate)
			}
		}

		resp := &rmnpb.Response{
			RequestId: requestID,
			Response: &rmnpb.Response_SignedObservation{
				SignedObservation: &rmnpb.SignedObservation{
					Observation: &rmnpb.Observation{
						RmnHomeContractConfigDigest: rmnHomeConfigDigest,
						LaneDest:                    destChain,
						FixedDestLaneUpdates:        laneUpdates,
						Timestamp:                   uint64(time.Now().UnixMilli()),
					},
					Signature: []byte{byte(nodeID), 1, 2, 3},
				},
			},
		}

		b, err := proto.Marshal(resp)
		require.NoError(ts.t, err)

		rmnClient.resChan <- RawRmnResponse{
			RMNNodeID: nodeID,
			Body:      b,
		}
	}
}

func (ts *testSetup) waitForInitialReportSignatureRequestsToBeSent(
	t *testing.T,
	rmnClient *mockRawRmnClient,
	minSigners int,
	minObservers int,
) map[uint32]uint64 {
	requestIDs := make(map[uint32]uint64)
	// plugin now has received the observation responses and should send
	// the report requests to the nodes, wait for them to be received by the nodes
	// should a total of minSigners requests each one containing the observation requests
	for {
		time.Sleep(time.Millisecond)

		recvReqs := rmnClient.getReceivedRequests()
		if len(recvReqs) < minSigners {
			continue
		}

		// check that the requests the node received are correct
		for _, reqs := range recvReqs {
			for _, req := range reqs {
				require.NotNil(t, req.GetReportSignatureRequest())
				assert.True(t, len(req.GetReportSignatureRequest().AttributedSignedObservations) >= minObservers)
			}
		}

		for nodeID, reqs := range recvReqs {
			for _, req := range reqs {
				requestIDs[nodeID] = req.RequestId
			}
		}

		rmnClient.resetReceivedRequests()
		return requestIDs
	}
}

func (ts *testSetup) nodesRespondToTheSignatureRequests(
	rmnClient *mockRawRmnClient,
	requestIDs map[uint32]uint64,
) {
	// now the plugin is waiting for rmn node responses for all this requests
	for nodeID, reqID := range requestIDs {
		resp := &rmnpb.Response{
			RequestId: reqID,
			Response: &rmnpb.Response_ReportSignature{
				ReportSignature: &rmnpb.ReportSignature{
					Signature: &rmnpb.EcdsaSignature{
						R: []byte{byte(nodeID), 1, 2, 3},
						S: []byte{byte(nodeID), 4, 5, 6},
					},
				},
			},
		}
		b, err := proto.Marshal(resp)
		require.NoError(ts.t, err)

		rmnClient.resChan <- RawRmnResponse{
			RMNNodeID: nodeID,
			Body:      b,
		}
	}
}

type mockRawRmnClient struct {
	resChan          chan RawRmnResponse
	receivedRequests map[uint32][]*rmnpb.Request
	mu               *sync.RWMutex
}

func newMockRawRmnClient(resChan chan RawRmnResponse) *mockRawRmnClient {
	return &mockRawRmnClient{
		mu:      &sync.RWMutex{},
		resChan: resChan,
	}
}

func (m *mockRawRmnClient) getReceivedRequests() map[uint32][]*rmnpb.Request {
	cp := make(map[uint32][]*rmnpb.Request)
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.receivedRequests {
		cp[k] = v
	}
	return cp
}

func (m *mockRawRmnClient) resetReceivedRequests() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.receivedRequests = make(map[uint32][]*rmnpb.Request)
}

func (m *mockRawRmnClient) Send(rmnNodeID uint32, request []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.receivedRequests == nil {
		m.receivedRequests = map[uint32][]*rmnpb.Request{}
	}

	if _, ok := m.receivedRequests[rmnNodeID]; !ok {
		m.receivedRequests[rmnNodeID] = []*rmnpb.Request{}
	}

	req := &rmnpb.Request{}
	err := proto.Unmarshal(request, req)
	if err != nil {
		return err
	}

	m.receivedRequests[rmnNodeID] = append(m.receivedRequests[rmnNodeID], req)
	return nil
}

func (m *mockRawRmnClient) Recv() <-chan RawRmnResponse {
	return m.resChan
}
