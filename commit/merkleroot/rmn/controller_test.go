package rmn

import (
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
)

var (
	chainS1       = cciptypes.ChainSelector(chainsel.TEST_90000002.Selector)
	chainS1OnRamp = []byte{0, 0xf, 0xf, 0xa, 0x0}

	chainS2       = cciptypes.ChainSelector(chainsel.TEST_90000003.Selector)
	chainS2OnRamp = []byte{0, 0xf, 0xf, 0xa, 0x1}

	chainD1        = cciptypes.ChainSelector(chainsel.TEST_90000004.Selector)
	chainD1OffRamp = []byte{0, 0xf, 0xf, 0xa, 0x2}
)

type testSetup struct {
	lggr           logger.Logger
	ctx            context.Context
	name           string
	t              *testing.T
	rmnController  *controller
	peerClient     *mockPeerClient
	updateRequests []*rmnpb.FixedDestLaneUpdateRequest
	rmnHomeMock    *readerpkg_mock.MockRMNHome
	remoteRMNCfg   rmntypes.RemoteConfig
	minObservers   int
	rmnNodes       []rmntypes.HomeNodeInfo
}

func Test_selectRoots(t *testing.T) {

	root1 := cciptypes.Bytes32{0, 0, 0, 1}
	root2 := cciptypes.Bytes32{0, 0, 0, 2}

	testCases := []struct {
		name         string
		observations []rmnSignedObservationWithMeta
		minObservers map[cciptypes.ChainSelector]int
		expErr       bool
		expRoots     map[cciptypes.ChainSelector]cciptypes.Bytes32
	}{
		{
			name: "happy path",
			observations: []rmnSignedObservationWithMeta{
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root1[:],
								},
							},
						},
					},
				},
			},
			minObservers: map[cciptypes.ChainSelector]int{chainS1: 1},
			expRoots: map[cciptypes.ChainSelector]cciptypes.Bytes32{
				chainS1: root1,
			},
		},
		{
			name: "zero valid roots",
			observations: []rmnSignedObservationWithMeta{
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root1[:],
								},
							},
						},
					},
				},
			},
			minObservers: map[cciptypes.ChainSelector]int{chainS1: 2}, // <-----
			expErr:       true,
		},
		{
			name: "observers not defined",
			observations: []rmnSignedObservationWithMeta{
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root1[:],
								},
							},
						},
					},
				},
			},
			minObservers: map[cciptypes.ChainSelector]int{}, // <-------
			expErr:       true,
		},
		{
			name: "more than one roots but one of them less than f",
			observations: []rmnSignedObservationWithMeta{
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root1[:],
								},
							},
						},
					},
				},
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root1[:],
								},
							},
						},
					},
				},
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root2[:],
								},
							},
						},
					},
				},
			},
			minObservers: map[cciptypes.ChainSelector]int{chainS1: 2},
			expRoots: map[cciptypes.ChainSelector]cciptypes.Bytes32{
				chainS1: root1,
			},
		},
		{
			name: "more than one valid roots",
			observations: []rmnSignedObservationWithMeta{
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root1[:],
								},
							},
						},
					},
				},
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root1[:],
								},
							},
						},
					},
				},
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource: &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1)},
									Root:       root2[:],
								},
							},
						},
					},
				},
			},
			minObservers: map[cciptypes.ChainSelector]int{chainS1: 1},
			expErr:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			roots, err := selectRoots(tc.observations, tc.minObservers)
			if tc.expErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expRoots, roots)
		})
	}
}

func TestClient_ComputeReportSignatures(t *testing.T) {
	newTestSetup := func(t *testing.T) testSetup {
		lggr := logger.Test(t)
		ctx := tests.Context(t)
		resChan := make(chan PeerResponse, 200)
		peerClient := newMockPeerClient(resChan)
		rmnHomeReaderMock := readerpkg_mock.NewMockRMNHome(t)

		const numNodes = 4
		rmnNodes := make([]rmntypes.HomeNodeInfo, numNodes)
		for i := 0; i < numNodes; i++ {
			// deterministically create a public key by seeding with a 32char string.
			publicKey, _, err := ed25519.GenerateKey(
				strings.NewReader(strconv.Itoa(i) + strings.Repeat("x", 31)))
			require.NoError(t, err)
			rmnNodes[i] = rmntypes.HomeNodeInfo{
				ID:                    rmntypes.NodeID(i + 1),
				PeerID:                ragep2ptypes.PeerID([32]byte{1, 2, 3}),
				SupportedSourceChains: mapset.NewSet(chainS1, chainS2),
				OffchainPublicKey:     &publicKey,
			}
		}

		cl := &controller{
			lggr:                                    lggr,
			peerClient:                              peerClient,
			rmnHomeReader:                           rmnHomeReaderMock,
			observationsInitialRequestTimerDuration: time.Minute,
			reportsInitialRequestTimerDuration:      time.Minute,
			ed25519Verifier:                         signatureVerifierAlwaysTrue{},
			rmnCrypto:                               signatureVerifierAlwaysTrue{},
		}

		updateRequests := []*rmnpb.FixedDestLaneUpdateRequest{
			{
				LaneSource:     &rmnpb.LaneSource{SourceChainSelector: uint64(chainS1), OnrampAddress: chainS1OnRamp},
				ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 10, MaxMsgNr: 20},
			},
			{
				LaneSource:     &rmnpb.LaneSource{SourceChainSelector: uint64(chainS2), OnrampAddress: chainS2OnRamp},
				ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 100, MaxMsgNr: 110},
			},
		}

		rmnRemoteCfg := rmntypes.RemoteConfig{
			ContractAddress: []byte{1, 2, 3},
			ConfigDigest:    cciptypes.Bytes32{0x1, 0x2, 0x3},
			MinSigners:      2,
			Signers: []rmntypes.RemoteSignerInfo{
				{
					OnchainPublicKey: []byte{1, 2, 3},
					NodeIndex:        1,
				},
				{
					OnchainPublicKey: []byte{4, 5, 6},
					NodeIndex:        2,
				},
				{
					OnchainPublicKey: []byte{7, 8, 9},
					NodeIndex:        3,
				},
				{
					OnchainPublicKey: []byte{10, 11, 12},
					NodeIndex:        4,
				},
			},
			ConfigVersion:    1,
			RmnReportVersion: cciptypes.Bytes32{0x1, 0x2, 0x3},
		}

		return testSetup{
			name:           t.Name(),
			t:              t,
			ctx:            ctx,
			lggr:           lggr,
			rmnController:  cl,
			peerClient:     peerClient,
			updateRequests: updateRequests,
			rmnHomeMock:    rmnHomeReaderMock,
			remoteRMNCfg:   rmnRemoteCfg,
			minObservers:   2,
			rmnNodes:       rmnNodes,
		}
	}

	destChain := &rmnpb.LaneDest{
		DestChainSelector: uint64(chainD1),
		OfframpAddress:    chainD1OffRamp,
	}

	t.Run("empty lane update request", func(t *testing.T) {
		ts := newTestSetup(t)

		ts.rmnHomeMock.On("GetMinObservers", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(
			map[cciptypes.ChainSelector]int{chainS1: 2, chainS2: 2}, nil)

		ts.rmnHomeMock.On("GetRMNNodesInfo", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(ts.rmnNodes, nil)

		_, err := ts.rmnController.ComputeReportSignatures(
			ts.ctx,
			destChain,
			[]*rmnpb.FixedDestLaneUpdateRequest{},
			ts.remoteRMNCfg,
		)
		assert.Error(t, err, ErrNothingToDo)
	})

	t.Run("happy path no retries", func(t *testing.T) {
		ts := newTestSetup(t)

		ts.rmnHomeMock.On("GetRMNNodesInfo", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(ts.rmnNodes, nil)
		ts.rmnHomeMock.On("GetMinObservers", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(
			map[cciptypes.ChainSelector]int{chainS1: 2, chainS2: 2, chainD1: 2}, nil)
		go func() {
			requestIDs, requestedChains := ts.waitForObservationRequestsToBeSent(
				ts.peerClient, ts.minObservers)

			ts.nodesRespondToTheObservationRequests(
				ts.peerClient, requestIDs, requestedChains, ts.remoteRMNCfg.ConfigDigest, destChain)

			requestIDs = ts.waitForReportSignatureRequestsToBeSent(
				t, ts.peerClient, int(ts.remoteRMNCfg.MinSigners),
				ts.minObservers)

			ts.nodesRespondToTheSignatureRequests(ts.peerClient, requestIDs)
		}()

		sigs, err := ts.rmnController.ComputeReportSignatures(
			ts.ctx,
			destChain,
			ts.updateRequests,
			ts.remoteRMNCfg,
		)
		assert.NoError(t, err)
		assert.Len(t, sigs.LaneUpdates, len(ts.updateRequests))
		assert.Len(t, sigs.Signatures, int(ts.remoteRMNCfg.MinSigners))
		// Make sure signature are in ascending signer address order
		for i := 1; i < len(sigs.Signatures); i++ {
			assert.True(t, sigs.Signatures[i].R[0] > sigs.Signatures[i-1].R[0])
			assert.True(t, sigs.Signatures[i].S[0] > sigs.Signatures[i-1].S[0])
		}
	})

	t.Run("happy path with retries", func(t *testing.T) {
		ts := newTestSetup(t)

		ts.rmnController.observationsInitialRequestTimerDuration = time.Nanosecond
		ts.rmnController.reportsInitialRequestTimerDuration = time.Nanosecond

		ts.rmnHomeMock.On("GetRMNNodesInfo", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(ts.rmnNodes, nil)
		ts.rmnHomeMock.On("GetMinObservers", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(
			map[cciptypes.ChainSelector]int{chainS1: 2, chainS2: 2}, nil)

		go func() {
			requestIDs, requestedChains := ts.waitForObservationRequestsToBeSent(
				ts.peerClient, ts.minObservers)

			// requests should be sent to at least two nodes
			assert.GreaterOrEqual(t, len(requestIDs), ts.minObservers)
			assert.GreaterOrEqual(t, len(requestedChains), ts.minObservers)

			ts.nodesRespondToTheObservationRequests(
				ts.peerClient, requestIDs, requestedChains, ts.remoteRMNCfg.ConfigDigest, destChain)
			time.Sleep(time.Millisecond)

			requestIDs = ts.waitForReportSignatureRequestsToBeSent(
				t, ts.peerClient, len(ts.remoteRMNCfg.Signers), ts.minObservers)
			time.Sleep(time.Millisecond)

			t.Logf("requestIDs: %v", requestIDs)

			// requests should be sent to all nodes, since we hit the timer timeout
			assert.Equal(t, len(requestIDs), len(ts.remoteRMNCfg.Signers))

			ts.nodesRespondToTheSignatureRequests(ts.peerClient, requestIDs)
		}()

		sigs, err := ts.rmnController.ComputeReportSignatures(
			ts.ctx,
			destChain,
			ts.updateRequests,
			ts.remoteRMNCfg,
		)
		assert.NoError(t, err)
		assert.Len(t, sigs.LaneUpdates, len(ts.updateRequests))
		assert.Len(t, sigs.Signatures, int(ts.remoteRMNCfg.MinSigners))
	})
}

func (ts *testSetup) waitForObservationRequestsToBeSent(
	rmnClient *mockPeerClient,
	minObservers int,
) (map[rmntypes.NodeID]uint64, map[rmntypes.NodeID]mapset.Set[uint64]) {
	requestIDs := make(map[rmntypes.NodeID]uint64)
	requestedChains := make(map[rmntypes.NodeID]mapset.Set[uint64])

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
		if requestsPerChain[uint64(chainS1)] >= minObservers && requestsPerChain[uint64(chainS2)] >= minObservers {
			for nodeID, reqs := range recvReqs {
				requestIDs[nodeID] = reqs[0].RequestId
				requestedChains[nodeID] = mapset.NewSet[uint64]()

				for _, req := range reqs {
					reqUpdates := req.GetObservationRequest().GetFixedDestLaneUpdateRequests()

					assert.True(ts.t, sort.SliceIsSorted(reqUpdates, func(i, j int) bool {
						return reqUpdates[i].LaneSource.SourceChainSelector < reqUpdates[j].LaneSource.SourceChainSelector
					}), "SourceChainSelector should be in ASC order")

					for _, src := range reqUpdates {
						requestedChains[nodeID].Add(src.LaneSource.SourceChainSelector)
					}
				}
			}

			rmnClient.resetReceivedRequests()
			break // all nodes have received the observation requests
		}
		continue
	}
	ts.t.Logf("test/mock nodes received the observation requests: %v", requestIDs)
	return requestIDs, requestedChains
}

func (ts *testSetup) nodesRespondToTheObservationRequests(
	rmnClient *mockPeerClient,
	requestIDs map[rmntypes.NodeID]uint64,
	requestedChains map[rmntypes.NodeID]mapset.Set[uint64],
	rmnHomeConfigDigest [32]byte,
	destChain *rmnpb.LaneDest,
) {
	allLaneUpdates := ts.updateRequests

	for nodeID, requestID := range requestIDs {
		laneUpdates := make([]*rmnpb.FixedDestLaneUpdate, 0)
		for _, laneUpdate := range allLaneUpdates {
			if requestedChains[nodeID].Contains(laneUpdate.LaneSource.SourceChainSelector) {
				root := sha256.Sum256([]byte(fmt.Sprintf("%d[%d,%d]",
					laneUpdate.LaneSource.SourceChainSelector,
					laneUpdate.ClosedInterval.MinMsgNr,
					laneUpdate.ClosedInterval.MaxMsgNr,
				)))
				laneUpdates = append(laneUpdates, &rmnpb.FixedDestLaneUpdate{
					LaneSource:     laneUpdate.LaneSource,
					ClosedInterval: laneUpdate.ClosedInterval,
					Root:           root[:],
				})
			}
		}

		resp := &rmnpb.Response{
			RequestId: requestID,
			Response: &rmnpb.Response_SignedObservation{
				SignedObservation: &rmnpb.SignedObservation{
					Observation: &rmnpb.Observation{
						RmnHomeContractConfigDigest: rmnHomeConfigDigest[:],
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

		ts.t.Logf("test/mock node %d responds to %d", nodeID, resp.RequestId)
		rmnClient.resChan <- PeerResponse{
			RMNNodeID: nodeID,
			Body:      b,
		}
	}
}

func (ts *testSetup) waitForReportSignatureRequestsToBeSent(
	t *testing.T,
	rmnClient *mockPeerClient,
	expectedResponses int,
	minObservers int,
) map[rmntypes.NodeID]uint64 {
	requestIDs := make(map[rmntypes.NodeID]uint64)
	// plugin now has received the observation responses and should send
	// the report requests to the nodes, wait for them to be received by the nodes
	// should a total of minSigners requests each one containing the observation requests
	for {
		time.Sleep(time.Millisecond)

		recvReqs := rmnClient.getReceivedRequests()
		if len(recvReqs) < expectedResponses {
			continue
		}

		cntValid := 0
		// check that the requests the node received are correct
		for _, reqs := range recvReqs {
			for _, req := range reqs {
				if req.GetReportSignatureRequest() == nil {
					continue
				}
				assert.True(t, len(req.GetReportSignatureRequest().AttributedSignedObservations) >= minObservers)

				aos := req.GetReportSignatureRequest().AttributedSignedObservations

				assert.True(t, sort.SliceIsSorted(aos, func(i, j int) bool {
					return aos[i].SignerNodeIndex < aos[j].SignerNodeIndex
				}), "SignerNodeIndex should be in ASC order")

				for _, ao := range aos {
					lus := ao.SignedObservation.Observation.FixedDestLaneUpdates
					assert.True(t, sort.SliceIsSorted(lus, func(i, j int) bool {
						return lus[i].LaneSource.SourceChainSelector < lus[j].LaneSource.SourceChainSelector
					}), "LaneSource.SourceChainSelector should be in ASC order")
				}
				cntValid++
			}
		}

		if cntValid < expectedResponses {
			continue
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
	rmnClient *mockPeerClient,
	requestIDs map[rmntypes.NodeID]uint64,
) {
	// now the plugin is waiting for rmn node responses for all this requests
	for nodeID, reqID := range requestIDs {
		r := [32]byte{byte(nodeID), 1, 2, 3}
		s := [32]byte{byte(nodeID), 4, 5, 6}

		resp := &rmnpb.Response{
			RequestId: reqID,
			Response: &rmnpb.Response_ReportSignature{
				ReportSignature: &rmnpb.ReportSignature{
					Signature: &rmnpb.EcdsaSignature{R: r[:], S: s[:]},
				},
			},
		}
		b, err := proto.Marshal(resp)
		require.NoError(ts.t, err)

		rmnClient.resChan <- PeerResponse{
			RMNNodeID: nodeID,
			Body:      b,
		}
	}
}

type mockPeerClient struct {
	resChan          chan PeerResponse
	receivedRequests map[rmntypes.NodeID][]*rmnpb.Request
	mu               *sync.RWMutex
}

func newMockPeerClient(resChan chan PeerResponse) *mockPeerClient {
	return &mockPeerClient{
		mu:      &sync.RWMutex{},
		resChan: resChan,
	}
}

func (m *mockPeerClient) getReceivedRequests() map[rmntypes.NodeID][]*rmnpb.Request {
	cp := make(map[rmntypes.NodeID][]*rmnpb.Request)
	m.mu.RLock()
	defer m.mu.RUnlock()
	for k, v := range m.receivedRequests {
		cp[k] = v
	}
	return cp
}

func (m *mockPeerClient) resetReceivedRequests() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.receivedRequests = make(map[rmntypes.NodeID][]*rmnpb.Request)
}

func (m *mockPeerClient) InitConnection(_ context.Context, _ cciptypes.Bytes32, _ cciptypes.Bytes32, _ []string) error {
	return nil
}

func (m *mockPeerClient) Close() error {
	return nil
}

func (m *mockPeerClient) Send(rmnNode rmntypes.HomeNodeInfo, request []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.receivedRequests == nil {
		m.receivedRequests = map[rmntypes.NodeID][]*rmnpb.Request{}
	}

	if _, ok := m.receivedRequests[rmnNode.ID]; !ok {
		m.receivedRequests[rmnNode.ID] = []*rmnpb.Request{}
	}

	req := &rmnpb.Request{}
	err := proto.Unmarshal(request, req)
	if err != nil {
		return err
	}

	m.receivedRequests[rmnNode.ID] = append(m.receivedRequests[rmnNode.ID], req)
	return nil
}

func (m *mockPeerClient) Recv() <-chan PeerResponse {
	return m.resChan
}

// signatureVerifierAlwaysTrue is a signature verifier that always returns true.
type signatureVerifierAlwaysTrue struct{}

func (a signatureVerifierAlwaysTrue) Verify(_ ed25519.PublicKey, _, _ []byte) bool {
	return true
}

func (a signatureVerifierAlwaysTrue) VerifyReportSignatures(
	_ context.Context, _ []cciptypes.RMNECDSASignature, _ cciptypes.RMNReport, _ []cciptypes.Bytes) error {
	return nil
}
