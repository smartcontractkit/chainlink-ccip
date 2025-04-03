package rmn

import (
	"context"
	"crypto/ed25519"
	"crypto/sha256"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	chainsel "github.com/smartcontractkit/chain-selectors"
	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"

	rmnpb "github.com/smartcontractkit/chainlink-protos/rmn/v1.6/go/serialization"

	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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
	remoteRMNCfg   cciptypes.RemoteConfig
	homeF          int
	rmnNodes       []rmntypes.HomeNodeInfo
}

func Test_selectRoots(t *testing.T) {

	root1 := cciptypes.Bytes32{0, 0, 0, 1}
	root2 := cciptypes.Bytes32{0, 0, 0, 2}

	testCases := []struct {
		name         string
		observations []rmnSignedObservationWithMeta
		homeF        map[cciptypes.ChainSelector]int
		expErr       bool
		expRoots     map[cciptypes.ChainSelector]cciptypes.Bytes32
	}{
		{
			name: "happy path F+1 observations",
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
			},
			homeF: map[cciptypes.ChainSelector]int{chainS1: 1},
			expRoots: map[cciptypes.ChainSelector]cciptypes.Bytes32{
				chainS1: root1,
			},
		},
		{
			name: "F observations instead of minimum F+1",
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
			homeF:  map[cciptypes.ChainSelector]int{chainS1: 1},
			expErr: true,
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
			homeF:  map[cciptypes.ChainSelector]int{chainS1: 2}, // <-----
			expErr: true,
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
			homeF:  map[cciptypes.ChainSelector]int{}, // <-------
			expErr: true,
		},
		{
			name: "more than one roots but one of them less than F+1",
			//nolint:dupl // to be fixed
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
			homeF: map[cciptypes.ChainSelector]int{chainS1: 2},
			expRoots: map[cciptypes.ChainSelector]cciptypes.Bytes32{
				chainS1: root1,
			},
		},
		{
			name: "more than one valid roots",
			//nolint:dupl // to be fixed
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
			homeF:  map[cciptypes.ChainSelector]int{chainS1: 1},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			roots, err := selectRoots(tc.observations, tc.homeF)
			if tc.expErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.expRoots, roots)
		})
	}
}

func Test_populateUpdatesPerChain(t *testing.T) {
	testCases := []struct {
		name           string
		updateRequests []*rmnpb.FixedDestLaneUpdateRequest
		rmnNodes       []rmntypes.HomeNodeInfo
		rmnNodeInfo    map[rmntypes.NodeID]rmntypes.HomeNodeInfo
		expectedResult map[uint64]updateRequestWithMeta
		expectedError  error
	}{
		{
			name: "single update request, single supported node",
			updateRequests: []*rmnpb.FixedDestLaneUpdateRequest{
				{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
			},
			rmnNodes: []rmntypes.HomeNodeInfo{
				{
					ID:                    1,
					SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(1)),
				},
			},
			expectedResult: map[uint64]updateRequestWithMeta{
				1: {
					Data:     &rmnpb.FixedDestLaneUpdateRequest{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
					RmnNodes: mapset.NewSet[rmntypes.NodeID](rmntypes.NodeID(1)),
				},
			},
			expectedError: nil,
		},
		{
			name: "duplicate sourceChainSelector error",
			updateRequests: []*rmnpb.FixedDestLaneUpdateRequest{
				{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
				{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
			},
			rmnNodes: []rmntypes.HomeNodeInfo{
				{
					ID:                    1,
					SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(1)),
				},
			},
			rmnNodeInfo:   map[rmntypes.NodeID]rmntypes.HomeNodeInfo{},
			expectedError: errors.New("controller implementation assumes each lane update is for a different chain"),
		},
		{
			name: "Single Update Request, No Supported Nodes",
			updateRequests: []*rmnpb.FixedDestLaneUpdateRequest{
				{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
			},
			rmnNodes: []rmntypes.HomeNodeInfo{
				{
					ID:                    2,
					SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(2)),
				},
			},
			expectedResult: map[uint64]updateRequestWithMeta{
				1: {
					Data:     &rmnpb.FixedDestLaneUpdateRequest{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
					RmnNodes: mapset.NewSet[rmntypes.NodeID](),
				},
			},
			expectedError: nil,
		},
		{
			name: "multiple update requests, multiple nodes",
			updateRequests: []*rmnpb.FixedDestLaneUpdateRequest{
				{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
				{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 2}},
				{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 3}},
			},
			rmnNodes: []rmntypes.HomeNodeInfo{
				{ID: 1, SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(1))},
				{ID: 2, SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(1))},
				{ID: 3, SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(2))},
				{ID: 4, SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(1))},
			},
			expectedResult: map[uint64]updateRequestWithMeta{
				1: {
					Data:     &rmnpb.FixedDestLaneUpdateRequest{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 1}},
					RmnNodes: mapset.NewSet[rmntypes.NodeID](rmntypes.NodeID(1), rmntypes.NodeID(2), rmntypes.NodeID(4)),
				},
				2: {
					Data:     &rmnpb.FixedDestLaneUpdateRequest{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 2}},
					RmnNodes: mapset.NewSet[rmntypes.NodeID](rmntypes.NodeID(3)),
				},
				3: {
					Data:     &rmnpb.FixedDestLaneUpdateRequest{LaneSource: &rmnpb.LaneSource{SourceChainSelector: 3}},
					RmnNodes: mapset.NewSet[rmntypes.NodeID](),
				},
			},
			expectedError: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			// Run the function
			result, err := populateUpdatesPerChain(tt.updateRequests, tt.rmnNodes)

			// Check for expected error
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}

			// Check for expected result
			assert.Equal(t, tt.expectedResult, result)

			// Check if rmnNodeInfo was populated correctly
			for id, info := range tt.rmnNodeInfo {
				assert.Equal(t, info, tt.rmnNodeInfo[id])
			}
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

		const numNodes = 8
		rmnNodes := make([]rmntypes.HomeNodeInfo, numNodes)
		for i := 0; i < numNodes; i++ {
			rmnNodes[i] = rmntypes.HomeNodeInfo{
				ID:                    rmntypes.NodeID(i + 1),
				PeerID:                [32]byte{1, 2, 3},
				SupportedSourceChains: mapset.NewSet(chainS1, chainS2),
				OffchainPublicKey:     getDeterministicPubKey(t),
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
			metricsReporter:                         NoopMetrics{},
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

		rmnRemoteCfg := cciptypes.RemoteConfig{
			ContractAddress: []byte{1, 2, 3},
			ConfigDigest:    cciptypes.Bytes32{0x1, 0x2, 0x3},
			FSign:           2,
			Signers: []cciptypes.RemoteSignerInfo{
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
			homeF:          2,
			rmnNodes:       rmnNodes,
		}
	}

	destChain := &rmnpb.LaneDest{
		DestChainSelector: uint64(chainD1),
		OfframpAddress:    chainD1OffRamp,
	}

	t.Run("empty lane update request", func(t *testing.T) {
		ts := newTestSetup(t)

		ts.rmnHomeMock.On("GetFObserve", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(
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
		ts.rmnHomeMock.On("GetFObserve", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(
			map[cciptypes.ChainSelector]int{chainS1: 2, chainS2: 2, chainD1: 2}, nil)
		go func() {
			requestIDs, requestedChains := ts.waitForObservationRequestsToBeSent(
				ts.peerClient, ts.homeF)

			ts.nodesRespondToTheObservationRequests(
				ts.peerClient, requestIDs, requestedChains, ts.remoteRMNCfg.ConfigDigest, destChain)

			requestIDs = ts.waitForReportSignatureRequestsToBeSent(
				t, ts.peerClient,
				int(ts.remoteRMNCfg.FSign)+1,
				ts.homeF,
			)

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
		assert.Len(t, sigs.Signatures, int(ts.remoteRMNCfg.FSign+1))
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
		ts.rmnHomeMock.On("GetFObserve", cciptypes.Bytes32{0x1, 0x2, 0x3}).Return(
			map[cciptypes.ChainSelector]int{chainS1: 2, chainS2: 2}, nil)

		go func() {
			requestIDs, requestedChains := ts.waitForObservationRequestsToBeSent(
				ts.peerClient, ts.homeF)

			// requests should be sent to at least two nodes
			assert.GreaterOrEqual(t, len(requestIDs), ts.homeF)
			assert.GreaterOrEqual(t, len(requestedChains), ts.homeF)

			ts.nodesRespondToTheObservationRequests(
				ts.peerClient, requestIDs, requestedChains, ts.remoteRMNCfg.ConfigDigest, destChain)
			time.Sleep(time.Millisecond)

			requestIDs = ts.waitForReportSignatureRequestsToBeSent(
				t,
				ts.peerClient,
				len(ts.remoteRMNCfg.Signers), // wait until all signers have responded
				ts.homeF,
			)
			time.Sleep(time.Millisecond)

			t.Logf("requestIDs: %v", requestIDs)

			// requests should be sent to more than F+1 nodes, since we hit the timer timeout
			assert.Greater(t, len(requestIDs), int(ts.remoteRMNCfg.FSign)+1)

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
		assert.Len(t, sigs.Signatures, int(ts.remoteRMNCfg.FSign+1))
	})
}

func Test_controller_validateSignedObservationResponse(t *testing.T) {
	configDigest123 := [32]byte{1, 2, 3}

	testCases := []struct {
		name                      string
		signedObservationResponse *rmnpb.SignedObservation
		rmnNodeID                 rmntypes.NodeID
		lurs                      map[uint64]updateRequestWithMeta
		destChain                 *rmnpb.LaneDest
		homeConfigDigest          [32]byte
		rmnNodesInfo              []rmntypes.HomeNodeInfo
		expErrContains            string
	}{
		{
			name: "single valid observation",
			signedObservationResponse: &rmnpb.SignedObservation{
				Observation: &rmnpb.Observation{
					RmnHomeContractConfigDigest: configDigest123[:],
					LaneDest:                    &rmnpb.LaneDest{DestChainSelector: 1},
					FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
						{
							LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 2},
							Root:           []byte{1, 2, 33},
							ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 1, MaxMsgNr: 2},
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
				Signature: []byte{10, 20, 30},
			},
			rmnNodeID: 20,
			lurs: map[uint64]updateRequestWithMeta{
				2: {
					Data: &rmnpb.FixedDestLaneUpdateRequest{
						LaneSource: &rmnpb.LaneSource{SourceChainSelector: 2},
						ClosedInterval: &rmnpb.ClosedInterval{
							MinMsgNr: 1,
							MaxMsgNr: 2,
						},
					},
					RmnNodes: mapset.NewSet(rmntypes.NodeID(20)),
				},
			},
			destChain:        &rmnpb.LaneDest{DestChainSelector: 1},
			homeConfigDigest: configDigest123,
			rmnNodesInfo: []rmntypes.HomeNodeInfo{
				{
					ID:                    20,
					SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(2)),
					OffchainPublicKey:     getDeterministicPubKey(t),
				},
			},
		},
		{
			name: "duplicate valid source lane updates should be rejected",
			signedObservationResponse: &rmnpb.SignedObservation{
				Observation: &rmnpb.Observation{
					RmnHomeContractConfigDigest: configDigest123[:],
					LaneDest:                    &rmnpb.LaneDest{DestChainSelector: 1},
					FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
						{
							LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 2},
							Root:           []byte{1, 2, 33},
							ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 1, MaxMsgNr: 2},
						},
						{
							LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 2},
							Root:           []byte{1, 2, 33},
							ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 1, MaxMsgNr: 2},
						},
					},
					Timestamp: uint64(time.Now().Unix()),
				},
			},
			rmnNodeID: 20,
			lurs: map[uint64]updateRequestWithMeta{
				2: {
					Data: &rmnpb.FixedDestLaneUpdateRequest{
						LaneSource: &rmnpb.LaneSource{SourceChainSelector: 2},
						ClosedInterval: &rmnpb.ClosedInterval{
							MinMsgNr: 1,
							MaxMsgNr: 2,
						},
					},
					RmnNodes: mapset.NewSet(rmntypes.NodeID(20)),
				},
			},
			destChain:        &rmnpb.LaneDest{DestChainSelector: 1},
			homeConfigDigest: configDigest123,
			rmnNodesInfo: []rmntypes.HomeNodeInfo{
				{
					ID:                    20,
					SupportedSourceChains: mapset.NewSet[cciptypes.ChainSelector](cciptypes.ChainSelector(2)),
				},
			},
			expErrContains: "duplicate source chain 2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)

			rmnHomeReaderMock := readerpkg_mock.NewMockRMNHome(t)
			rmnHomeReaderMock.EXPECT().GetRMNNodesInfo(cciptypes.Bytes32(tc.homeConfigDigest)).
				Return(tc.rmnNodesInfo, nil)

			cl := &controller{
				lggr:            lggr,
				ed25519Verifier: signatureVerifierAlwaysTrue{},
				rmnCrypto:       signatureVerifierAlwaysTrue{},
				rmnHomeReader:   rmnHomeReaderMock,
			}

			err := cl.validateSignedObservationResponse(
				&rmnpb.Response{
					RequestId: 0,
					Response:  &rmnpb.Response_SignedObservation{SignedObservation: tc.signedObservationResponse},
				},
				tc.rmnNodeID,
				tc.lurs,
				tc.destChain,
				tc.homeConfigDigest,
			)
			if tc.expErrContains != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrContains)
				return
			}
			require.NoError(t, err)
		})
	}
}

func Test_newRequestID(t *testing.T) {
	ids := map[uint64]struct{}{}
	for i := 0; i < 1000; i++ {
		id := newRequestID(logger.Test(t))
		_, ok := ids[id]
		assert.False(t, ok)
		ids[id] = struct{}{}
	}
}

func Test_chainsWithSufficientObservationResponses(t *testing.T) {
	b1 := [32]byte{0, 0, 0, 0, 1}
	b2 := [32]byte{0, 0, 0, 0, 2}
	b3 := [32]byte{0, 0, 0, 0, 3}

	testCases := []struct {
		name                    string
		updateRequests          map[uint64]updateRequestWithMeta
		rmnObservationResponses []rmnSignedObservationWithMeta
		homeFMap                map[cciptypes.ChainSelector]int
		expChains               mapset.Set[cciptypes.ChainSelector]
	}{
		{
			name: "one case with all scenarios",
			updateRequests: map[uint64]updateRequestWithMeta{
				1: { // all good for source chain 1
					Data: &rmnpb.FixedDestLaneUpdateRequest{
						LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 1},
						ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 10, MaxMsgNr: 20},
					},
				},
				2: { // f not found
					Data: &rmnpb.FixedDestLaneUpdateRequest{
						LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 2},
						ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 20, MaxMsgNr: 30},
					},
				},
				3: { // not enough observations
					Data: &rmnpb.FixedDestLaneUpdateRequest{
						LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 3},
						ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 30, MaxMsgNr: 40},
					},
				},
				4: { // zero observations
					Data: &rmnpb.FixedDestLaneUpdateRequest{
						LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 4},
						ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 40, MaxMsgNr: 50},
					},
				},
			},
			rmnObservationResponses: []rmnSignedObservationWithMeta{
				{
					SignedObservation: &rmnpb.SignedObservation{
						Observation: &rmnpb.Observation{
							FixedDestLaneUpdates: []*rmnpb.FixedDestLaneUpdate{
								{
									LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 1, OnrampAddress: chainS1OnRamp},
									Root:           b1[:],
									ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 10, MaxMsgNr: 20},
								},
								{
									LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 2, OnrampAddress: chainS1OnRamp},
									Root:           b2[:],
									ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 20, MaxMsgNr: 30},
								},
								{
									LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 3, OnrampAddress: chainS1OnRamp},
									Root:           b2[:],
									ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 30, MaxMsgNr: 40},
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
									LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 1},
									Root:           b1[:],
									ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 10, MaxMsgNr: 20},
								},
								{
									LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 2, OnrampAddress: chainS1OnRamp},
									Root:           b2[:],
									ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 20, MaxMsgNr: 30},
								},
								{
									LaneSource:     &rmnpb.LaneSource{SourceChainSelector: 3, OnrampAddress: chainS1OnRamp},
									Root:           b3[:],
									ClosedInterval: &rmnpb.ClosedInterval{MinMsgNr: 30, MaxMsgNr: 40},
								},
							},
						},
					},
				},
			},
			homeFMap: map[cciptypes.ChainSelector]int{
				1: 1, // at least f+1 = 2
				3: 2, // at least f+1 = 3
			},
			expChains: mapset.NewSet(cciptypes.ChainSelector(1)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lggr := logger.Test(t)
			chains := chainsWithSufficientObservationResponses(
				lggr, tc.updateRequests, tc.rmnObservationResponses, tc.homeFMap)
			assert.True(t, tc.expChains.Equal(chains))
		})
	}
}

func getDeterministicPubKey(t *testing.T) *ed25519.PublicKey {
	// deterministically create a public key by seeding with a 32char string.
	publicKey, _, err := ed25519.GenerateKey(
		strings.NewReader(strconv.Itoa(1) + strings.Repeat("x", 31)))
	require.NoError(t, err)
	return &publicKey
}

func (ts *testSetup) waitForObservationRequestsToBeSent(
	rmnClient *mockPeerClient,
	homeF int,
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
		if requestsPerChain[uint64(chainS1)] >= homeF+1 && requestsPerChain[uint64(chainS2)] >= homeF+1 {
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
	homeF int,
) map[rmntypes.NodeID]uint64 {
	requestIDs := make(map[rmntypes.NodeID]uint64)
	// plugin now has received the observation responses and should send
	// the report requests to the nodes, wait for them to be received by the nodes
	// should a total of #remoteF requests each one containing the observation requests
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
				assert.Greater(t, len(req.GetReportSignatureRequest().AttributedSignedObservations), homeF)

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

func (m *mockPeerClient) InitConnection(
	_ context.Context,
	_ cciptypes.Bytes32,
	_ cciptypes.Bytes32,
	_ []ragep2ptypes.PeerID,
	_ []rmntypes.HomeNodeInfo) error {
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
	_ context.Context, _ []cciptypes.RMNECDSASignature, _ cciptypes.RMNReport, _ []cciptypes.UnknownAddress) error {
	return nil
}
