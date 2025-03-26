// Package summary collects data across multiple logs and prints a summary for each OCR round.
package summary

import (
	"fmt"
	"regexp"
	"sort"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/parse"
	"github.com/smartcontractkit/chainlink-ccip/cmd/carpenter/internal/render"
)

func init() {
	render.Register("summary", basicRendererFactory, "Analyze logs and print a summary of each OCR round.")

	reportRegex = regexp.MustCompile("^built (\\d+) reports$")
}

var reportRegex *regexp.Regexp

func basicRendererFactory(options render.Options) render.Renderer {
	sr := summaryRenderer{
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

// summaryRenderer holds metadata collected across multiple log lines.
type summaryRenderer struct {
	commits map[int]map[int]commitSummary
	execs   map[int]map[int]execSummary
}

func (sr summaryRenderer) execCollector(data *parse.Data) {

}

func (sr summaryRenderer) Render(data *parse.Data) {
	switch data.Plugin {
	case "Commit":
		sr.commitCollector(data)
	case "Execute":
		sr.execCollector(data)
	default:
		// TODO: print transmit info?
	}
}

func (sr summaryRenderer) Close() error {
	dons := maps.Keys(sr.commits)
	sort.Ints(dons)
	for _, donID := range dons {
		fmt.Println("Commit Summary for DON", donID)
		keys := maps.Keys(sr.commits[donID])
		sort.Ints(keys)
		for _, key := range keys {
			fmt.Println(sr.commits[donID][key].String())
		}
		fmt.Println()
	}

	dons = maps.Keys(sr.execs)
	sort.Ints(dons)
	for _, donID := range dons {
		fmt.Println("Commit logs for DON", donID)
		keys := maps.Keys(sr.execs[donID])
		sort.Ints(keys)
		for _, key := range keys {
			fmt.Println(sr.execs[donID][key].String())
		}
	}
	return nil
}
