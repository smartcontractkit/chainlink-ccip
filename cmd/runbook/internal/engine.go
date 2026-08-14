package internal

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

// Querier runs one instant PromQL query and returns its series. Executor
// implements it; tests substitute a fake.
type Querier interface {
	Query(ctx context.Context, q string) ([]Series, error)
}

// Inputs holds resolved runbook inputs and numeric parameters used during
// substitution and condition evaluation.
type Inputs struct {
	values map[string]string
	params map[string]float64
}

// NewInputs fills the inputs struct from the runbook's declared Input list and
// the values supplied for them (flag-provided). Required-but-missing inputs
// error so a run can't silently proceed with a $-hole in a query.
func NewInputs(rb *Runbook, provided map[string]string) (*Inputs, error) {
	in := &Inputs{values: map[string]string{}, params: map[string]float64{}}
	for _, d := range rb.Inputs {
		v, ok := provided[d.Name]
		if !ok || v == "" {
			v = d.Default
		}
		if !ok && d.Default == "" && d.Required {
			return nil, fmt.Errorf("required input %s is missing (%s)", d.Name, d.Description)
		}
		if d.Type == "integer" && v != "" {
			f, err := parseFloat(v)
			if err != nil {
				return nil, fmt.Errorf("input %s must be an integer, got %q", d.Name, v)
			}
			in.params[d.Name] = f // also make it usable as a bare param in conditions
		}
		in.values[d.Name] = v
	}
	return in, nil
}

// Value returns the string value for a named input, or "" if unset.
func (in *Inputs) Value(name string) string { return in.values[name] }

// Param returns a numeric parameter (e.g. fRoleDON), plus whether it is set.
func (in *Inputs) Param(name string) (float64, bool) {
	v, ok := in.params[name]
	return v, ok
}

// substitute replaces $name placeholders in a query/condition with input values.
func (in *Inputs) substitute(s string) string {
	var b strings.Builder
	for {
		i := strings.IndexByte(s, '$')
		if i < 0 || i+1 >= len(s) {
			b.WriteString(s)
			break
		}
		b.WriteString(s[:i])
		j := i + 1
		for j < len(s) && isIdentChar(rune(s[j])) {
			j++
		}
		name := s[i+1 : j]
		if v, ok := in.values[name]; ok {
			b.WriteString(v)
		} else {
			b.WriteString(s[i:j])
		}
		s = s[j:]
	}
	return b.String()
}

func (in *Inputs) condEnv(series []Series) CondEnv {
	return CondEnv{Series: series, Params: in.params}
}

// --- evidence ---

// QueryEvidence is one executed query and its raw result, emitted so an AI
// layer or human can verify a verdict against ground truth.
type QueryEvidence struct {
	Query string
	Error string
	Raw   []Series
}

// --- checklist runner ---

// CheckResult holds one check's computed status and canonical value.
type CheckResult struct {
	ID            string          `yaml:"id"`
	Status        string          `yaml:"status"`
	Value         string          `yaml:"value"`
	AlwaysEmitted bool            `yaml:"always_emitted"`
	Info          bool            `yaml:"info,omitempty"`
	Group         string          `yaml:"group"`
	Owner         string          `yaml:"owner"`
	Note          string          `yaml:"note"`
	Evidence      []QueryEvidence `yaml:"-"`
}

type ChecklistReport struct {
	DestChain    string        `yaml:"destChain"`
	SourceChains string        `yaml:"sourceChains"`
	Timestamp    string        `yaml:"timestamp"`
	Overall      string        `yaml:"overall"`
	Checks       []CheckResult `yaml:"checks"`
	Concerns     []Concern     `yaml:"concerns"`
}

type Concern struct {
	Checks   []string `yaml:"checks"`
	Severity string   `yaml:"severity"`
	Summary  string   `yaml:"summary"`
	Owner    string   `yaml:"owner"`
}

const (
	statusOK      = "OK"
	statusWARN    = "WARN"
	statusCRIT    = "CRIT"
	statusINFO    = "INFO"
	statusUNKNOWN = "UNKNOWN"
)

var severityRank = func() map[string]int {
	return map[string]int{statusCRIT: 4, statusWARN: 3, statusUNKNOWN: 2, statusINFO: 1, statusOK: 0}
}()

func RunChecklist(ctx context.Context, rb *Runbook, ex Querier, in *Inputs) (*ChecklistReport, error) {
	rep := &ChecklistReport{
		DestChain:    in.Value("destChain"),
		SourceChains: in.Value("sourceChains"),
		Overall:      statusOK,
	}
	if rep.SourceChains == "" {
		rep.SourceChains = ".*"
	}
	rep.Timestamp = nowISO()

	for _, c := range rb.Checks {
		cr := runCheck(ctx, ex, in, c)
		rep.Checks = append(rep.Checks, cr)
		if severityRank[cr.Status] > severityRank[rep.Overall] {
			rep.Overall = cr.Status
		}
	}
	rep.Concerns = aggregateConcerns(rep.Checks)
	return rep, nil
}

func runCheck(ctx context.Context, ex Querier, in *Inputs, c Check) CheckResult {
	res := CheckResult{
		ID:            c.ID,
		Group:         c.Group,
		Owner:         c.Owner,
		AlwaysEmitted: c.AlwaysEmitted,
		Info:          c.Info,
	}
	if c.Info {
		res.Status = statusINFO
	}
	query := in.substitute(c.Query)
	series, err := ex.Query(ctx, query)
	ev := QueryEvidence{Query: query, Raw: series}
	if err != nil {
		ev.Error = err.Error()
		res.Evidence = append(res.Evidence, ev)
		res.Status = statusUNKNOWN
		res.Note = "query failed: " + err.Error()
		return res
	}
	res.Evidence = append(res.Evidence, ev)
	res.Value = seriesString(series)

	if c.Info {
		return res
	}

	// Empty-result rule: an always_emitted check that returns nothing means the
	// pipeline is suspect -> UNKNOWN. Otherwise treat empty as vacuous
	// (evaluate conditions normally: all/any over no series).
	if len(series) == 0 && c.AlwaysEmitted {
		res.Status = statusUNKNOWN
		res.Note = "always_emitted metric returned no series: telemetry may be stalled, not healthy"
		return res
	}

	env := in.condEnv(series)
	hitUnknown := false
	for _, rule := range c.Rules {
		b, err := MatchResult(rule.Condition, env)
		if err != nil {
			res.Status = statusUNKNOWN
			res.Note = "bad severity condition: " + err.Error()
			return res
		}
		switch b {
		case True:
			res.Status = rule.Status
			return res
		case Unknown:
			hitUnknown = true
		}
	}
	if hitUnknown {
		res.Status = statusUNKNOWN
		res.Note = "severity conditions unresolved"
		return res
	}
	res.Status = statusOK
	return res
}

// aggregateConcerns folds non-OK checks into the report's concerns, applying
// the checklist's documented combination rules: fchain_read_errors +
// consensus_observation_failed + consensus_dropped are one incident,
// live_oracle_count takes precedence over fchain_read_errors, everything else
// becomes its own concern worst-first, and UNKNOWN is never silently dropped
// to OK.
func aggregateConcerns(checks []CheckResult) []Concern {
	find := func(id string) *CheckResult {
		for i := range checks {
			if checks[i].ID == id {
				return &checks[i]
			}
		}
		return nil
	}
	isBad := func(id string) bool {
		c := find(id)
		return c != nil && c.Status != statusOK && c.Status != statusINFO
	}

	consensus := isBad("consensus_observation_failed")
	dropped := isBad("consensus_dropped")
	h1 := isBad("fchain_read_errors")
	live := find("live_oracle_count")
	liveBad := live != nil && (live.Status == statusCRIT || live.Status == statusWARN)

	var concerns []Concern

	// One folded finding for the consensus family, when any member is bad.
	if consensus || dropped || h1 {
		ids := []string{}
		if consensus {
			ids = append(ids, "consensus_observation_failed")
		}
		if dropped {
			ids = append(ids, "consensus_dropped")
		}
		if h1 {
			ids = append(ids, "fchain_read_errors")
		}
		sev := statusWARN
		if consensus {
			sev = statusCRIT
		}
		owner, summary := "ccip-commit-oncall", "DON-wide FChain consensus failure"
		if liveBad {
			summary += fmt.Sprintf("; live_oracle_count is %s and takes precedence (likely explains this on its own)", live.Status)
		} else if consensus && !h1 {
			summary += "; check consensus_dropped for split vs insufficient_agreement (H2 vs H1) / threshold_not_defined"
		} else if h1 && !consensus {
			owner, summary = "home-chain-infra-oncall", "fchain_read_errors spiking (H1, home-chain read outage)"
		}
		concerns = append(concerns, Concern{Checks: ids, Severity: sev, Summary: summary, Owner: owner})
	}

	folded := make(map[string]bool)
	for _, cn := range concerns {
		for _, cid := range cn.Checks {
			folded[cid] = true
		}
	}

	// Remaining non-OK and UNKNOWN checks become their own concerns.
	var rest []CheckResult
	for _, c := range checks {
		if c.Status == statusOK || c.Status == statusINFO || folded[c.ID] {
			continue
		}
		rest = append(rest, c)
	}
	sort.SliceStable(rest, func(i, j int) bool {
		return severityRank[rest[i].Status] > severityRank[rest[j].Status]
	})
	for _, c := range rest {
		concerns = append(concerns, Concern{
			Checks:   []string{c.ID},
			Severity: c.Status,
			Summary:  c.Note,
			Owner:    c.Owner,
		})
	}
	return concerns
}

// seriesString renders a result set in the contract's canonical multi-series
// form: all grouping labels joined + ":" + value, series separated by "; ".
func seriesString(series []Series) string {
	if len(series) == 0 {
		return "no data"
	}
	parts := make([]string, 0, len(series))
	for _, s := range series {
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
		b.WriteString(fmt.Sprintf("%g", s.Value))
		parts = append(parts, b.String())
	}
	return strings.Join(parts, "; ")
}

func parseFloat(s string) (float64, error) {
	var f float64
	_, err := fmt.Sscanf(s, "%g", &f)
	return f, err
}
