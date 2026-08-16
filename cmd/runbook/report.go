package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/smartcontractkit/chainlink-ccip/cmd/runbook/internal"
)

func writer() io.Writer { return os.Stdout }

// emitConcerns is the default, concise output for a checklist run: a
// worst-finding table, not a full YAML report. Use --verbose for the complete
// structured report.
func emitConcerns(rep *internal.ChecklistReport) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	fmt.Fprintf(w, "health:\t%s\n\tdestChain:\t%s\n\tsourceChains:\t%s\n\tchecks run:\t%d\n",
		rep.Overall, rep.DestChain, rep.SourceChains, len(rep.Checks))
	if len(rep.Concerns) == 0 {
		fmt.Fprintln(w, "concerns:\tnone (all checks OK/INFO; use --verbose for full detail)")
	} else {
		fmt.Fprintln(w, "concerns:")
		fmt.Fprintf(w, "\t%8s\t%-28s\t%-20s\t%s\n", "severity", "checks", "owner", "summary")
		for _, c := range rep.Concerns {
			fmt.Fprintf(w, "\t%8s\t%-28s\t%-20s\t%s\n",
				c.Severity, strings.Join(c.Checks, ","), c.Owner, truncate(c.Summary, 120))
		}
	}
	return w.Flush()
}

// emitTrace is the default, concise output for a decision-graph run: the path
// taken and the terminal finding.
func emitTrace(rep *internal.GraphReport) error {
	w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	fmt.Fprintln(w, "trace:")
	fmt.Fprintf(w, "\t%-12s\t%-10s\t%s\n", "step", "status", "outcome")
	for _, s := range rep.Steps {
		var out string
		if s.Target != "" {
			out += s.Target
		}
		if s.Reason != "" {
			if out != "" {
				out += " — "
			}
			out += truncate(s.Reason, 100)
		}
		if out == "" {
			out = "-"
		}
		fmt.Fprintf(w, "\t%-12s\t%-10s\t%s\n", s.ID, s.Status, out)
	}
	fmt.Fprintln(w, "final: ", rep.Final)
	fmt.Fprintln(w, "(use --verbose for the full YAML report)")
	return w.Flush()
}

// truncate shortens a string for tabular display, appending an ellipsis.
func truncate(s string, n int) string {
	r := []rune(s)
	if len(r) <= n {
		return s
	}
	return string(r[:n]) + "…"
}

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
