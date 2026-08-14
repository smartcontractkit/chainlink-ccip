package internal

import (
	"context"
	"sort"
	"strings"
	"testing"
)

type seriesOpt func(*Series)

func withLabel(k, v string) seriesOpt {
	return func(s *Series) { s.Labels[k] = v }
}

func mk(val float64, opts ...seriesOpt) Series {
	s := Series{Labels: map[string]string{}}
	for _, o := range opts {
		o(&s)
	}
	s.Value = val
	return s
}

func env(series ...Series) CondEnv { return CondEnv{Series: series} }

func TestMatchResult(t *testing.T) {
	cases := []struct {
		name  string
		cond  string
		env   CondEnv
		want  Bool
		parse bool // expect a parse error
	}{
		{name: "result eq 0", cond: "result == 0", env: env(mk(0)), want: True},
		{name: "result gt 0", cond: "result > 0", env: env(mk(5)), want: True},
		{name: "result gt 0 false", cond: "result > 0", env: env(mk(0)), want: False},
		// empty-result rule: any == false, all == true (vacuous)
		{name: "any empty", cond: "any series == 1", env: env(), want: False},
		{name: "all empty", cond: "all series == 0", env: env(), want: True},
		{name: "any multi", cond: "any series > 0", env: env(mk(0), mk(2)), want: True},
		{name: "all multi", cond: "all series == 0", env: env(mk(0), mk(0)), want: True},
		{name: "all multi false", cond: "all series == 0", env: env(mk(0), mk(1)), want: False},
		// label selectors
		{name: "label eq", cond: `any series with status="live" == 1`, env: env(mk(1, withLabel("status", "live")), mk(0, withLabel("status", "rmn"))), want: True},
		{name: "label eq absent", cond: `any series with status="live" == 1`, env: env(mk(0, withLabel("status", "rmn"))), want: False},
		{name: "label regexp", cond: `any series with status=~"skipped_.*" == 1`, env: env(mk(1, withLabel("status", "skipped_price")), mk(0, withLabel("status", "live"))), want: True},
		{name: "label ne", cond: `any series with reason!="stale" > 0`, env: env(mk(3, withLabel("reason", "config_digest_mismatch")), mk(1, withLabel("reason", "stale"))), want: True},
		{name: "label ne absent", cond: `any series with reason!="stale" > 0`, env: env(mk(1, withLabel("reason", "stale"))), want: False},
		{name: "any + no combo", cond: `any series > 0 and no series with reason!="stale" > 0`, env: env(mk(1, withLabel("reason", "stale"))), want: True},
		{name: "any + no combo false", cond: `any series > 0 and no series with reason!="stale" > 0`, env: env(mk(1, withLabel("reason", "stale")), mk(1, withLabel("reason", "boom"))), want: False},
		// numeric params + known guard + chained comparison
		{name: "known true", cond: "known fRoleDON and result < 2*fRoleDON+1", env: CondEnv{Series: []Series{mk(2)}, Params: map[string]float64{"fRoleDON": 1}}, want: True},
		{name: "known false shortcircuit", cond: "known fRoleDON and result < 2*fRoleDON+1", env: env(mk(2)), want: False},
		{name: "not known", cond: "not known fRoleDON", env: env(mk(2)), want: True},
		{name: "chained bound within", cond: "2*fRoleDON+1 <= result < 3*fRoleDON+1", env: CondEnv{Series: []Series{mk(3)}, Params: map[string]float64{"fRoleDON": 1}}, want: True},
		{name: "chained below", cond: "2*fRoleDON+1 <= result < 3*fRoleDON+1", env: CondEnv{Series: []Series{mk(2)}, Params: map[string]float64{"fRoleDON": 1}}, want: False},
		{name: "chained above", cond: "2*fRoleDON+1 <= result < 3*fRoleDON+1", env: CondEnv{Series: []Series{mk(5)}, Params: map[string]float64{"fRoleDON": 1}}, want: False},
		{name: "result empty is 0", cond: "result > 0", env: env(), want: False},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := MatchResult(tc.cond, tc.env)
			if err != nil {
				t.Fatalf("MatchResult(%q): %v", tc.cond, err)
			}
			if got != tc.want {
				t.Errorf("MatchResult(%q) = %v, want %v", tc.cond, got, tc.want)
			}
		})
	}
}

func TestMatchResultBadSyntax(t *testing.T) {
	for _, c := range []string{"result >", "any series", "known", "foo bar", "result >0 x"} {
		if _, err := MatchResult(c, env(mk(1))); err == nil {
			t.Errorf("expected parse error for %q", c)
		}
	}
}

// fakeQuerier returns canned series for queries containing a given substring.
// The longest matching substring wins, so a specific matcher can't be shadowed
// by a more general one in nondeterministic map order.
type fakeQuerier struct {
	matchers []fakeMatcher
}

type fakeMatcher struct {
	sub    string
	series []Series
	err    error
}

func (f *fakeQuerier) match(sub string, series ...Series) *fakeQuerier {
	f.matchers = append(f.matchers, fakeMatcher{sub: sub, series: series})
	return f
}

func (f *fakeQuerier) fail(sub string, err error) *fakeQuerier {
	f.matchers = append(f.matchers, fakeMatcher{sub: sub, err: err})
	return f
}

func (f *fakeQuerier) Query(_ context.Context, q string) ([]Series, error) {
	sort.SliceStable(f.matchers, func(i, j int) bool {
		return len(f.matchers[i].sub) > len(f.matchers[j].sub)
	})
	for _, m := range f.matchers {
		if strings.Contains(q, m.sub) {
			if m.err != nil {
				return nil, m.err
			}
			return m.series, nil
		}
	}
	return nil, nil
}

func loadForTest(t *testing.T, name string) *Runbook {
	t.Helper()
	rb, err := LoadRunbook(name)
	if err != nil {
		t.Fatalf("LoadRunbook(%s): %v", name, err)
	}
	return rb
}

func TestRunChecklist_Healthy(t *testing.T) {
	rb := loadForTest(t, "commit-plugin-health")
	// Cover every always_emitted:true check with a plausible healthy value and
	// leave every always_emitted:false counter unmatched (empty -> OK, the
	// event never fired). Under this the overall should be OK/INFO (some
	// checks are pure-context INFO), never WARN/CRIT/UNKNOWN.
	f := (&fakeQuerier{}).
		mark("phase=\"observation\"", 3).
		mark(`phase="outcome"`, 3).
		mark("config_digest_mismatch{", 0).
		mark("processor_latency_bucket", 0).
		mark("pending_messages{", 0, withLabel("source_network_name", "base")).
		mark("offramp_lane_status{", 1, withLabel("status", "live")).
		mark("rmn_curse_active{", 0, withLabel("curse_type", "global")).
		mark("source_chain_cursed{", 0, withLabel("source_network_name", "base")).
		mark("csa_public_key", 4)
	in, _ := NewInputs(rb, map[string]string{"destChain": "eth"})
	rep, err := RunChecklist(context.Background(), rb, f, in)
	if err != nil {
		t.Fatalf("RunChecklist: %v", err)
	}
	for _, c := range rep.Checks {
		switch c.Status {
		case statusCRIT, statusWARN, statusUNKNOWN:
			t.Errorf("check %s: expected healthy, got %s (note: %s)", c.ID, c.Status, c.Note)
		}
	}
	if rep.Overall == statusCRIT || rep.Overall == statusWARN || rep.Overall == statusUNKNOWN {
		t.Errorf("overall should be OK/INFO, got %s", rep.Overall)
	}
}

// mark is a test helper shorthand for f.match with an optional label error.
func (f *fakeQuerier) mark(sub string, v float64, opts ...seriesOpt) *fakeQuerier {
	return f.match(sub, mk(v, opts...))
}

func TestRunChecklist_EmptyHeartbeatIsUnknown(t *testing.T) {
	rb := loadForTest(t, "commit-plugin-health")
	// No metric matches at all -> every always_emitted:true check is empty.
	f := &fakeQuerier{}
	in, _ := NewInputs(rb, map[string]string{"destChain": "eth"})
	rep, err := RunChecklist(context.Background(), rb, f, in)
	if err != nil {
		t.Fatalf("RunChecklist: %v", err)
	}
	for _, c := range rep.Checks {
		// Info checks carry no verdict (pending_messages, latency) -> skip.
		if c.Info {
			continue
		}
		if c.AlwaysEmitted && c.Status != statusUNKNOWN {
			t.Errorf("check %s: always_emitted with empty result should be UNKNOWN, got %s", c.ID, c.Status)
		}
		// non-Info, non-always-emitted counters empty -> OK (event never fired)
		if !c.AlwaysEmitted && c.Status != statusOK {
			t.Errorf("check %s: non-always-emitted counter with empty result should be OK, got %s", c.ID, c.Status)
		}
	}
}

func fakeUncommittedQuerier(gaveUp float64) *fakeQuerier {
	// Tune the triage path: message observed (onramp ahead of offramp), no
	// per-chain consensus failure, nothing cursed, and give_up = gaveUp. Every
	// prefix step must return non-empty series or the graph treats it as an
	// AGENT handoff (empty -> unknown), so zero-valued matchers are included.
	return (&fakeQuerier{}).
		mark("plugin_heartbeat_total{", 1).
		mark("offramp_next_seq_num", 10).
		mark("onramp_max_seq_num", 20).
		mark("consensus_dropped_total", 0).
		mark("offramp_consensus_insufficient_total", 0).
		mark("consensus_observation_failed", 0).
		mark("rmn_curse_active", 0).
		mark("source_chain_cursed", 0).
		mark("report_transmission_gave_up", gaveUp)
}

func TestRunGraph_BacklogReport(t *testing.T) {
	rb := loadForTest(t, "uncommitted-message")
	// gaveUp == 0 -> step5 false -> REPORT backlog, stops.
	f := fakeUncommittedQuerier(0)
	in, _ := NewInputs(rb, map[string]string{"destChain": "eth", "sourceChain": "base", "seqNum": "20"})
	rep, err := RunGraph(context.Background(), rb, f, in, 30)
	if err != nil {
		t.Fatalf("RunGraph: %v", err)
	}
	if !strings.HasPrefix(rep.Final, "REPORT: ccip-commit-oncall") {
		t.Errorf("expected REPORT backlog terminal, got %q", rep.Final)
	}
	for _, s := range rep.Steps {
		if s.Status == "AGENT" && !s.Judgment {
			t.Errorf("unexpected AGENT without judgment at step %s", s.ID)
		}
	}
}

func TestRunGraph_Scenario2Reason(t *testing.T) {
	rb := loadForTest(t, "uncommitted-message")
	f := fakeUncommittedQuerier(1) // gave_up>0 -> scenario2
	// At scenario2, a positive config_digest_mismatch rejection must map to REPORT:ccip-commit-oncall.
	f.match("report_validation_rejected", mk(1, withLabel("reason", "stale")), mk(2, withLabel("reason", "config_digest_mismatch")))
	in, _ := NewInputs(rb, map[string]string{"destChain": "eth", "sourceChain": "base", "seqNum": "20"})
	rep, err := RunGraph(context.Background(), rb, f, in, 30)
	if err != nil {
		t.Fatalf("RunGraph: %v", err)
	}
	if !strings.HasPrefix(rep.Final, "REPORT: ccip-commit-oncall") {
		t.Errorf("expected REPORT on config_digest_mismatch, got %q", rep.Final)
	}
}

func TestRunGraph_Scenario2UnmappedReason(t *testing.T) {
	rb := loadForTest(t, "uncommitted-message")
	f := fakeUncommittedQuerier(1)
	f.match("report_validation_rejected", mk(1, withLabel("reason", "forced_test_failure")))
	in, _ := NewInputs(rb, map[string]string{"destChain": "eth", "sourceChain": "base", "seqNum": "20"})
	rep, err := RunGraph(context.Background(), rb, f, in, 30)
	if err != nil {
		t.Fatalf("RunGraph: %v", err)
	}
	if !strings.HasPrefix(rep.Final, "AGENT") {
		t.Errorf("unmapped reason should be an AGENT handoff, got %q", rep.Final)
	}
}

func TestRunGraph_Scenario3FuzzyJudgment(t *testing.T) {
	rb := loadForTest(t, "uncommitted-message")
	f := fakeUncommittedQuerier(0)
	// consensus_observation_failed > 0 -> CONTINUE:scenario3 (judgment leaf).
	// Longer substring than the default so it wins the longest-match tiebreak.
	f.match("consensus_observation_failed_total{", mk(2, withLabel("chain_id", "eth")))
	f.match("pending_messages{dest_network_name", mk(5, withLabel("source_network_name", "base"))) // unused; scenario3 is a judgment leaf
	in, _ := NewInputs(rb, map[string]string{"destChain": "eth", "sourceChain": "base", "seqNum": "20"})
	rep, err := RunGraph(context.Background(), rb, f, in, 30)
	if err != nil {
		t.Fatalf("RunGraph: %v", err)
	}
	if !strings.HasPrefix(rep.Final, "AGENT") {
		t.Errorf("fuzzy scenario3 should hand off to AGENT, got %q", rep.Final)
	}
}

func TestNewInputs_RequiresRequired(t *testing.T) {
	rb := loadForTest(t, "uncommitted-message")
	if _, err := NewInputs(rb, map[string]string{"destChain": "eth"}); err == nil {
		t.Errorf("expected missing-required-input error for sourceChain/seqNum")
	}
}
