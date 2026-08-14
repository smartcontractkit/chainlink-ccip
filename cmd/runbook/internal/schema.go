package internal

import (
	"embed"
	"fmt"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

//go:embed runbooks/*.yaml
var runbookFS embed.FS

// Runbook is the machine-readable form of one docs/runbooks file. It is meant
// to be a faithful, executable transcription of the fenced YAML block in the
// markdown doc — the two must not silently drift (same step/check IDs, same
// owners, same queries). Type selects the control structure: "checklist"
// (proactive/health — run all checks, aggregate) or "graph" (reactive/triage —
// walk a decision tree from root to a terminal outcome).
type Runbook struct {
	Name        string  `yaml:"name"`
	Description string  `yaml:"description"`
	Type        string  `yaml:"type"` // "checklist" | "graph"
	Owner       string  `yaml:"owner"`
	Inputs      []Input `yaml:"inputs"`
	Checks      []Check `yaml:"checks"`
	Root        string  `yaml:"root"`  // graph only
	Steps       []Step  `yaml:"steps"` // graph only
}

type Input struct {
	Name        string `yaml:"name"`
	Type        string `yaml:"type"`
	Description string `yaml:"description"`
	Required    bool   `yaml:"required"`
	Default     string `yaml:"default"`
}

// Validate returns an error if the runbook is structurally inconsistent.
func (r *Runbook) Validate() error {
	if r.Name == "" || r.Type == "" {
		return fmt.Errorf("runbook missing name or type")
	}
	seen := map[string]bool{}
	switch r.Type {
	case "checklist":
		for _, c := range r.Checks {
			if err := c.Validate(); err != nil {
				return fmt.Errorf("check %q: %w", c.ID, err)
			}
			if seen[c.ID] {
				return fmt.Errorf("duplicate check id %q", c.ID)
			}
			seen[c.ID] = true
		}
	case "graph":
		for _, s := range r.Steps {
			if err := s.Validate(); err != nil {
				return fmt.Errorf("step %q: %w", s.ID, err)
			}
			if seen[s.ID] {
				return fmt.Errorf("duplicate step id %q", s.ID)
			}
			seen[s.ID] = true
		}
		if r.Root == "" || !seen[r.Root] {
			return fmt.Errorf("graph root %q not present in steps", r.Root)
		}
	default:
		return fmt.Errorf("unknown runbook type %q", r.Type)
	}
	return nil
}

// Check is a single independent health check (checklist runbook).
type Check struct {
	ID            string         `yaml:"id"`
	Group         string         `yaml:"group"`
	AlwaysEmitted bool           `yaml:"always_emitted"`
	Query         string         `yaml:"query"`
	Owner         string         `yaml:"owner"`
	Info          bool           `yaml:"info"`     // pure-context, no verdict
	Rules         []SeverityRule `yaml:"severity"` // evaluated in order; first hit wins
	Note          string         `yaml:"note"`
}

func (c *Check) Validate() error {
	if c.ID == "" || c.Query == "" {
		return fmt.Errorf("missing id or query")
	}
	if !c.Info && len(c.Rules) == 0 {
		return fmt.Errorf("non-info check has no severity rules")
	}
	for _, r := range c.Rules {
		if r.Status == "" || r.Condition == "" {
			return fmt.Errorf("severity rule missing status or condition")
		}
	}
	return nil
}

// SeverityRule maps a boolean condition to a status when the check is not Info.
type SeverityRule struct {
	Status    string `yaml:"status"` // OK | WARN | CRIT | INFO
	Condition string `yaml:"condition"`
}

// Step is a node in a decision-graph runbook. A step that points at a system
// this runbook has no query for (chain explorer, TXM dashboard) is marked
// NotAutomatable: the engine emits an AGENT/HUMAN handoff instead of guessing.
type Step struct {
	ID      string   `yaml:"id"`
	Label   string   `yaml:"label"`
	Queries []string `yaml:"queries"`
	// Condition is a boolean expression over the combined query results.
	Condition string `yaml:"condition"`
	// Judgment marks the condition as deliberately fuzzy (e.g. "is this series
	// climbing vs flat") — the engine stops and hands the raw results to an
	// agent/human rather than force a boolean.
	Judgment       bool `yaml:"judgment,omitempty"`
	NotAutomatable bool `yaml:"not_automatable,omitempty"`
	// Outcomes keyed by "true" / "false" (a bare boolean step).
	Outcomes map[string]Outcome `yaml:"outcomes"`
	// ReasonOutcomes classify a `by (reason)` query's positive series by reason
	// label; "default" is the fallback for unmapped reason strings.
	ReasonOutcomes []ReasonOutcome `yaml:"reason_outcomes,omitempty"`
	// None is the outcome applied when a reasoning/decision step has no
	// positive match (e.g. scenario2 with all rejection counts zero).
	None *Outcome `yaml:"none_outcome,omitempty"`
	// Decisions is an ordered list of (condition -> outcome) evaluated in
	// order; first true wins, falling back to None. Used for multi-branch
	// numeric thresholds (e.g. oracle-count grading).
	Decisions []Decision `yaml:"decisions,omitempty"`
	Note      string     `yaml:"note"`
}

// Decision is one branch of a multi-way step.
type Decision struct {
	Condition string `yaml:"condition"`
	Action    string `yaml:"action"`
	Reason    string `yaml:"reason"`
}

func (s *Step) Validate() error {
	if s.ID == "" {
		return fmt.Errorf("missing id")
	}
	if s.NotAutomatable {
		return nil // handoff leaf; no queries/condition required
	}
	fuzzy := s.Judgment
	hasOutcomes := len(s.Outcomes) > 0
	hasReasons := len(s.ReasonOutcomes) > 0
	hasDecisions := len(s.Decisions) > 0
	matched := 0
	for _, b := range []bool{hasOutcomes, hasReasons, hasDecisions} {
		if b {
			matched++
		}
	}
	if !fuzzy && matched != 1 {
		return fmt.Errorf("step must have exactly one of outcomes, reason_outcomes, decisions")
	}
	if fuzzy && s.Condition == "" {
		return fmt.Errorf("judgment step missing condition")
	}
	// Boolean steps need both outcomes; reason steps need reason_outcomes;
	// multi-branch steps need decisions.
	if !fuzzy && matched == 1 {
		if hasOutcomes {
			if s.Condition == "" {
				return fmt.Errorf("boolean step missing condition")
			}
			if _, ok := s.Outcomes["true"]; !ok {
				return fmt.Errorf("step has no outcome for condition=true")
			}
			if _, ok := s.Outcomes["false"]; !ok {
				return fmt.Errorf("step has no outcome for condition=false")
			}
		}
		if hasDecisions {
			for _, d := range s.Decisions {
				if d.Condition == "" || d.Action == "" {
					return fmt.Errorf("decision missing condition or action")
				}
			}
		}
	}
	return nil
}

// Outcome is a terminal or continuation verdict a step can reach.
type Outcome struct {
	Action string `yaml:"action"` // CONTINUE:<id> | STOP | REPORT:<owner> | AGENT:<text>
	Reason string `yaml:"reason"`
}

// ReasonOutcome classifies a positive series of a by(reason) query.
type ReasonOutcome struct {
	Label  string `yaml:"reason"` // exact reason label; "default" = fallback
	Action string `yaml:"action"`
	Text   string `yaml:"reason_text"`
}

// LoadRunbook reads and validates a runbook by its canonical name.
func LoadRunbook(name string) (*Runbook, error) {
	names, err := EmbeddedRunbooks()
	if err != nil {
		return nil, err
	}
	for _, n := range names {
		if strings.TrimSuffix(n, ".yaml") == name || n == name {
			data, err := runbookFS.ReadFile("runbooks/" + n)
			if err != nil {
				return nil, err
			}
			var rb Runbook
			if err := yaml.Unmarshal(data, &rb); err != nil {
				return nil, fmt.Errorf("parsing %s: %w", n, err)
			}
			if err := rb.Validate(); err != nil {
				return nil, fmt.Errorf("validating %s: %w", n, err)
			}
			return &rb, nil
		}
	}
	return nil, fmt.Errorf("unknown runbook %q (available: %s)", name, strings.Join(names, ", "))
}

// EmbeddedRunbooks returns the canonical names of the bundled runbooks.
func EmbeddedRunbooks() ([]string, error) {
	entries, err := runbookFS.ReadDir("runbooks")
	if err != nil {
		return nil, err
	}
	var names []string
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".yaml") {
			names = append(names, e.Name())
		}
	}
	sort.Strings(names)
	return names, nil
}
