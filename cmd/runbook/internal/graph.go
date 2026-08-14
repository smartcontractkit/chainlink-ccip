package internal

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// GraphStepResult is one step's outcome during a decision-graph walk.
type GraphStepResult struct {
	ID       string          `yaml:"id"`
	Label    string          `yaml:"label"`
	Status   string          `yaml:"status"`           // CONTINUE | STOP | REPORT | AGENT
	Target   string          `yaml:"target,omitempty"` // CONTINUE:step / REPORT:owner
	Reason   string          `yaml:"reason,omitempty"`
	Judgment bool            `yaml:"judgment,omitempty"`
	Auto     bool            `yaml:"automatable"`
	Value    string          `yaml:"value,omitempty"`
	Evidence []QueryEvidence `yaml:"-"`
}

// GraphReport is the trace a graph runbook produces.
type GraphReport struct {
	Name      string            `yaml:"name"`
	Timestamp string            `yaml:"timestamp"`
	Steps     []GraphStepResult `yaml:"steps"`
	Final     string            `yaml:"final"`
}

type actionKind int

const (
	actContinue actionKind = iota
	actStop
	actReport
	actAgent
)

func parseAction(a string) (actionKind, string) {
	a = strings.TrimSpace(a)
	switch {
	case strings.HasPrefix(a, "CONTINUE:"):
		return actContinue, strings.TrimSpace(strings.TrimPrefix(a, "CONTINUE:"))
	case strings.HasPrefix(a, "REPORT:"):
		return actReport, strings.TrimSpace(strings.TrimPrefix(a, "REPORT:"))
	case strings.HasPrefix(a, "AGENT:"):
		return actAgent, strings.TrimSpace(strings.TrimPrefix(a, "AGENT:"))
	case strings.TrimSpace(a) == "STOP":
		return actStop, ""
	default:
		return actStop, ""
	}
}

func (r *GraphReport) step(s GraphStepResult) {
	r.Steps = append(r.Steps, s)
}

// RunGraph walks a decision-graph runbook from root to a terminal outcome,
// executing each automatable step's queries and following its outcomes. Steps
// that require external judgment (Judgment=true) or point at systems this
// runbook has no query for (NotAutomatable=true) are surfaced as AGENT
// handoffs with their raw results — never guessed.
func RunGraph(ctx context.Context, rb *Runbook, ex Querier, in *Inputs, maxSteps int) (*GraphReport, error) {
	rep := &GraphReport{Name: rb.Name, Timestamp: nowISO()}
	byID := map[string]*Step{}
	for i := range rb.Steps {
		byID[rb.Steps[i].ID] = &rb.Steps[i]
	}
	cur := rb.Root
	visited := map[string]bool{}
	for step := 0; step < maxSteps; step++ {
		s := byID[cur]
		if s == nil {
			return nil, fmt.Errorf("step %q not found", cur)
		}
		if visited[cur] {
			return nil, fmt.Errorf("cycle detected at step %q", cur)
		}
		visited[cur] = true

		st := GraphStepResult{ID: s.ID, Label: s.Label, Auto: !s.NotAutomatable, Judgment: s.Judgment}

		// Leaf pointing at an external system we have no query for.
		if s.NotAutomatable {
			st.Status, st.Reason = "AGENT", s.Note
			rep.step(st)
			rep.Final = terminalString(st)
			return rep, nil
		}

		// Run queries, preserving raw results as evidence.
		var combined []Series
		queryFailed := false
		for _, q := range s.Queries {
			query := in.substitute(q)
			series, err := ex.Query(ctx, query)
			ev := QueryEvidence{Query: query, Raw: series}
			if err != nil {
				ev.Error = err.Error()
				queryFailed = true
			}
			st.Evidence = append(st.Evidence, ev)
			combined = append(combined, series...)
		}
		st.Value = seriesString(combined)

		// Empty result -> we genuinely don't know; never treat as 0.
		if len(combined) == 0 {
			st.Status = "AGENT"
			if queryFailed {
				st.Reason = "query failed (see evidence): cannot distinguish 'healthy idle' from 'pipeline broken'. Hand raw results to agent/human."
			} else {
				st.Reason = "all queries returned no series; cannot distinguish 'healthy idle' from 'pipeline broken'. Hand raw results to agent/human."
			}
			rep.step(st)
			rep.Final = terminalString(st)
			return rep, nil
		}

		// Deliberately-fuzzy condition (e.g. "climbing vs flat").
		if s.Judgment {
			st.Status, st.Reason = "AGENT", "condition requires trend/judgment, not a boolean: "+s.Condition
			rep.step(st)
			rep.Final = terminalString(st)
			return rep, nil
		}

		// `by (reason)` classification (Scenario 2 style) when configured.
		if len(s.ReasonOutcomes) > 0 {
			st = classifyByReason(st, s, combined)
			rep.step(st)
			if st.Status != "CONTINUE" {
				rep.Final = terminalString(st)
				return rep, nil
			}
			cur = st.Target
			continue
		}

		// Ordered multi-branch decisions (Scenario 3b oracle-count grading).
		if len(s.Decisions) > 0 {
			var outcome *Outcome
			for _, d := range s.Decisions {
				db, err := MatchResult(in.substitute(d.Condition), in.condEnv(combined))
				if err != nil || db != True {
					continue
				}
				outcome = &Outcome{Action: d.Action, Reason: d.Reason}
				break
			}
			if outcome == nil {
				outcome = s.None
			}
			if outcome == nil {
				st.Status, st.Reason = "AGENT", "no decision matched and no none_outcome"
				rep.step(st)
				rep.Final = terminalString(st)
				return rep, nil
			}
			kind, target := parseAction(outcome.Action)
			st.Status, st.Target, st.Reason = kindStr(kind), target, outcome.Reason
			if kind == actContinue {
				rep.step(st)
				cur = target
				continue
			}
			rep.step(st)
			rep.Final = terminalString(st)
			return rep, nil
		}

		b, err := MatchResult(in.substitute(s.Condition), in.condEnv(combined))
		if err != nil || b == Unknown {
			st.Status = "AGENT"
			if err != nil {
				st.Reason = "could not evaluate condition: " + err.Error()
			} else {
				st.Reason = "condition unresolved (unknown): " + s.Condition
			}
			rep.step(st)
			rep.Final = terminalString(st)
			return rep, nil
		}
		key := "false"
		if b == True {
			key = "true"
		}
		o, ok := s.Outcomes[key]
		if !ok {
			st.Status, st.Reason = "AGENT", "no outcome defined for condition="+key
			rep.step(st)
			rep.Final = terminalString(st)
			return rep, nil
		}
		kind, target := parseAction(o.Action)
		st.Status, st.Target, st.Reason = kindStr(kind), target, o.Reason
		if kind == actContinue {
			if _, ok := byID[target]; !ok {
				st.Status = "AGENT"
				st.Reason = "CONTINUE points at unknown step " + target
				rep.step(st)
				rep.Final = terminalString(st)
				return rep, nil
			}
			rep.step(st)
			cur = target
			continue
		}
		rep.step(st)
		rep.Final = terminalString(st)
		return rep, nil
	}
	return nil, fmt.Errorf("graph exceeded maxSteps (%d)", maxSteps)
}

// classifyByReason maps the most severe positive `by (reason)` series to an
// outcome. Priority is by declaration order in the YAML (the doc lists the most
// actionable reasons first, e.g. config_digest_mismatch before the benign
// stale). An unknown positive reason falls through to the "default" mapping;
// no positive counts falls back to the step's None outcome.
func classifyByReason(st GraphStepResult, s *Step, combined []Series) GraphStepResult {
	// positive reason labels present
	positive := map[string]bool{}
	for _, series := range combined {
		if series.Value > 0 {
			positive[series.Labels["reason"]] = true
		}
	}
	if len(positive) == 0 {
		if s.None != nil {
			kind, target := parseAction(s.None.Action)
			st.Status, st.Target, st.Reason = kindStr(kind), target, s.None.Reason
			return st
		}
		st.Status, st.Reason = "STOP", "no positive counts in window"
		return st
	}
	// Highest-priority declared reason that is present.
	for _, ro := range s.ReasonOutcomes {
		if ro.Label == "default" {
			continue // evaluated last
		}
		if positive[ro.Label] {
			kind, target := parseAction(ro.Action)
			st.Status, st.Target, st.Reason = kindStr(kind), target, ro.Text
			return st
		}
	}
	// A positive reason with no mapping -> default.
	for _, ro := range s.ReasonOutcomes {
		if ro.Label == "default" {
			kind, target := parseAction(ro.Action)
			st.Status, st.Target, st.Reason = kindStr(kind), target, ro.Text
			return st
		}
	}
	st.Status, st.Reason = "AGENT", "unmapped reason"
	return st
}

func terminalString(st GraphStepResult) string {
	s := st.Status
	if st.Target != "" {
		s += ": " + st.Target
	}
	if st.Reason != "" {
		s += " — " + st.Reason
	}
	return s
}

func kindStr(k actionKind) string {
	switch k {
	case actContinue:
		return "CONTINUE"
	case actReport:
		return "REPORT"
	case actAgent:
		return "AGENT"
	default:
		return "STOP"
	}
}

func suffixOwner(target string) string {
	if target == "" {
		return ""
	}
	return ": " + target
}

func nowISO() string { return time.Now().Format(time.RFC3339) }
