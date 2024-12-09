package filter

import (
	"fmt"
	"os"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func init() {
	parse.RegisterDataFilter("Merkle Root Observation", MerkleRootObservation)
}

var _ parse.DataFilter = MerkleRootObservation

func MerkleRootObservation(object map[string]interface{}) *parse.Data {
	var result *parse.Data
	switch {
	case nameContains("CommitPlugin.MerkleRootProcessor", object):
		result = &parse.Data{
			Level:           getString("level", object),
			Caller:          getString("caller", object),
			Plugin:          "Commit",
			PluginProcessor: "MerkleRoot",
			SequenceNumber:  0,
			Details:         getString("msg", object),
		}
	}

	if result == nil {
		return nil
	}

	time, err := getTimestamp("ts", object)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse timestamp: %s", err)
	} else {
		result.Timestamp = time
	}

	return result
}
