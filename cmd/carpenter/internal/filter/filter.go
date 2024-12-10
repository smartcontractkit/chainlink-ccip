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

func MerkleRootObservation(data parse.Data, object map[string]interface{}) *parse.Data {
	var result *parse.Data
	if data.PluginProcessor == "MerkleRoot" {
		result = &data
		data.Details = data.Message
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
