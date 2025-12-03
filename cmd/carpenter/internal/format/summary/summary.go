// Package summary collects data across multiple logs and prints a summary for each OCR round.
package summary

import (
	"fmt"
	"maps"
	"slices"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/format"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
)

func init() {
	format.Register("summary", basicFormatterFactory, "Analyze logs and print a summary of each OCR round.")
}

func basicFormatterFactory(options format.Options) format.Formatter {
	sr := summaryFormatter{
		commits: make(map[int]map[int]commitSummary),
		execs:   make(map[int]map[int]execSummary),
	}
	return sr
}

type execSummary struct {
	logs      []*parse.Data
	seqNumber int
}

func (es execSummary) String() string {
	return ""
}

// summaryFormatter holds metadata collected across multiple log lines.
type summaryFormatter struct {
	commits map[int]map[int]commitSummary
	execs   map[int]map[int]execSummary
}

func (sr summaryFormatter) execCollector(data *parse.Data) {

}

func (sr summaryFormatter) Format(data *parse.Data) {
	switch data.Plugin {
	case "Commit":
		sr.commitCollector(data)
	case "Execute":
		sr.execCollector(data)
	default:
		// TODO: print transmit info?
	}
}

func (sr summaryFormatter) Close() error {
	dons := slices.Sorted(maps.Keys(sr.commits))
	for _, donID := range dons {
		fmt.Println("Commit Summary for DON", donID)
		keys := slices.Sorted(maps.Keys(sr.commits[donID]))
		for _, key := range keys {
			fmt.Println(sr.commits[donID][key].String())
		}
		fmt.Println()
	}

	dons = slices.Sorted(maps.Keys(sr.execs))
	for _, donID := range dons {
		fmt.Println("Commit logs for DON", donID)
		keys := slices.Sorted(maps.Keys(sr.execs[donID]))
		for _, key := range keys {
			fmt.Println(sr.execs[donID][key].String())
		}
	}
	return nil
}
