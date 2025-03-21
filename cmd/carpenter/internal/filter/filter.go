package filter

import (
	"fmt"
	"os"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func init() {
	parse.RegisterDataFilter("CommitFilter", CommitFilter)
}

var _ parse.DataFilter = CommitFilter

func CommitFilter(data parse.Data) *parse.Data {
	var result *parse.Data
	if data.Plugin == "Commit" {
		//if data.PluginProcessor == "MerkleRoot" {
		result = &data
		//data.Details = data.Message
	}

	if result == nil {
		return nil
	}

	time, err := getTimestamp("ts", data.RawLoggerFields)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse timestamp: %s", err)
	} else {
		result.ProdTimestamp = time.String()
	}

	return result
}
