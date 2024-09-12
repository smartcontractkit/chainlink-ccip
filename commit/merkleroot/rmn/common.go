package rmn

import (
	"fmt"
	"math/big"

	mapset "github.com/deckarep/golang-set/v2"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

func (c *client) marshalAndSend(req *rmnpb.Request, nodeID NodeID) error {
	reqBytes, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("proto marshal RMN request: %w", err)
	}

	if err := c.rawRmnClient.Send(nodeID, reqBytes); err != nil {
		return fmt.Errorf("send rmn request: %w", err)
	}

	return nil
}

// parseResponse parses the response from the RMN and returns the response.
// Validates that the response is expected and not a duplicate.
func (c *client) parseResponse(
	resp *RawRmnResponse, requestIDs, gotResponses mapset.Set[uint64]) (*rmnpb.Response, error) {

	c.lggr.Infof("requests we are waiting for: %s", requestIDs.String())

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

type ReportData struct {
	DestChainEvmID              *big.Int          `abi:"destChainId"`
	DestChainSelector           uint64            `abi:"destChainSelector"`
	RmnRemoteContractAddress    string            `abi:"rmnRemoteContractAddress"`
	OfframpAddress              string            `abi:"offrampAddress"`
	RmnHomeContractConfigDigest cciptypes.Bytes32 `abi:"rmnHomeContractConfigDigest"`
	LaneUpdates                 []LaneUpdate      `abi:"destLaneUpdates"`
}

type LaneUpdate struct {
	SourceChainSelector uint64            `abi:"sourceChainSelector"`
	OnRampAddress       []byte            `abi:"onRampAddress"`
	MinSeqNr            uint64            `abi:"minSeqNr"`
	MaxSeqNr            uint64            `abi:"maxSeqNr"`
	MerkleRoot          cciptypes.Bytes32 `abi:"merkleRoot"`
}

// NewLaneUpdatesFromPBType creates a LaneUpdate from protobuf LaneUpdate type.
func NewLaneUpdatesFromPBType(pb []*rmnpb.FixedDestLaneUpdate) []LaneUpdate {
	laneUpdates := make([]LaneUpdate, len(pb))
	for i, pbLaneUpdate := range pb {
		root32 := [32]byte{}
		copy(root32[:], pbLaneUpdate.Root)

		laneUpdates[i] = LaneUpdate{
			SourceChainSelector: pbLaneUpdate.LaneSource.SourceChainSelector,
			MinSeqNr:            pbLaneUpdate.ClosedInterval.MinMsgNr,
			MaxSeqNr:            pbLaneUpdate.ClosedInterval.MaxMsgNr,
			MerkleRoot:          root32,
			OnRampAddress:       pbLaneUpdate.LaneSource.OnrampAddress,
		}
	}
	return laneUpdates
}
