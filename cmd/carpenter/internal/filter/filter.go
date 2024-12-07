package filter

import (
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func init() {
	parse.RegisterDataFilter("Merkle Root Observation", MerkleRootObservation)
}

var _ parse.DataFilter = MerkleRootObservation

func MerkleRootObservation(object map[string]interface{}) *parse.Data {
	if nameContains("CommitPlugin.MerkleRootProcessor", object) {
		return &parse.Data{
			Timestamp:       getString("ts", object),
			Level:           getString("level", object),
			Caller:          getString("caller", object),
			Plugin:          "Commit",
			PluginProcessor: "MerkleRoot",
			SequenceNumber:  0,
			Details:         getString("msg", object),
		}
	}
	return nil
}
