package main

import (
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/cmd/runbook/internal"
)

func writer() io.Writer { return os.Stdout }

// extractEvidence pulls the raw query results (kept off the YAML report proper
// via yaml:"-") into a side map so an agent or human can verify a verdict
// against ground truth. Format per entry: query -> "label=val,...:value; ...".
func extractEvidence(v any) map[string]any {
	out := map[string]any{}
	switch r := v.(type) {
	case *internal.ChecklistReport:
		for _, c := range r.Checks {
			if len(c.Evidence) == 0 {
				continue
			}
			out["check:"+c.ID] = evidenceEntries(c.Evidence)
		}
	case *internal.GraphReport:
		for _, s := range r.Steps {
			if len(s.Evidence) == 0 {
				continue
			}
			out["step:"+s.ID] = evidenceEntries(s.Evidence)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}

func evidenceEntries(ev []internal.QueryEvidence) any {
	entries := make([]map[string]any, 0, len(ev))
	for _, e := range ev {
		entry := map[string]any{"query": e.Query}
		if e.Error != "" {
			entry["error"] = e.Error
		}
		var series []string
		for _, s := range e.Raw {
			keys := make([]string, 0, len(s.Labels))
			for k := range s.Labels {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			var b strings.Builder
			for _, k := range keys {
				b.WriteString(k)
				b.WriteString("=")
				b.WriteString(s.Labels[k])
				b.WriteString(",")
			}
			b.WriteString(trimFloat(s.Value))
			series = append(series, b.String())
		}
		if len(series) > 0 {
			entry["series"] = series
		}
		entries = append(entries, entry)
	}
	return entries
}

func trimFloat(f float64) string {
	return strconv.FormatFloat(f, 'g', -1, 64)
}
