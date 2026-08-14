package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"
)

// promServer is a minimal Prometheus-compatible endpoint used to exercise the
// HTTP path (suffix probe, JSON decode, strip/no-strip).
type promServer struct {
	mu            sync.Mutex
	name          string // __name__ returned for every series
	received      []string
	suffixExposed bool
}

func (s *promServer) handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q, _ := url.QueryUnescape(r.URL.Query().Get("query"))
		s.mu.Lock()
		s.received = append(s.received, q)
		s.mu.Unlock()

		// Probe query: {__name__=~"...(_total)?"} -> report the exposed spelling.
		if strings.Contains(q, `__name__=~`) {
			name := s.name
			if !s.suffixExposed {
				name = strings.TrimSuffix(s.name, "_total")
			}
			fmt.Fprintf(w, `{"status":"success","data":{"resultType":"vector","result":[{"metric":{"__name__":%q},"value":[1700000000,"1"]}]}}`, name)
			return
		}
		fmt.Fprintf(w, `{"status":"success","data":{"resultType":"vector","result":[{"metric":{"__name__":%q,"phase":"observation"},"value":[1700000000,"5"]}]}}`, s.name)
	})
}

func (s *promServer) latestQuery() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	if n := len(s.received); n > 0 {
		return s.received[n-1]
	}
	return ""
}

func TestExecutorProbeAndStrip(t *testing.T) {
	tests := []struct {
		name          string
		exposedName   string // what the datasource actually exposes
		wantStripRes  bool   // expected resolved strip value after auto probe
		wantSentER    string // substring we expect the real query to contain
		notWantSentER string // substring we expect the real query to NOT contain
	}{
		{
			name:          "otel pipeline exposes _total",
			exposedName:   "ccip_commit_plugin_heartbeat_total",
			wantStripRes:  false,
			wantSentER:    "ccip_commit_plugin_heartbeat_total",
			notWantSentER: "",
		},
		{
			name:          "direct scrape exposes base name",
			exposedName:   "ccip_commit_plugin_heartbeat",
			wantStripRes:  true,
			wantSentER:    "ccip_commit_plugin_heartbeat{",
			notWantSentER: "_total",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := &promServer{name: tc.exposedName, suffixExposed: strings.HasSuffix(tc.exposedName, "_total")}
			srv := httptest.NewServer(svc.handler())
			defer srv.Close()

			ex := NewExecutor(srv.URL, 5*time.Second) // Strip nil -> auto probe
			series, err := ex.Query(context.Background(), "sum(rate(ccip_commit_plugin_heartbeat_total{chainID=\"x\"}[1m]))")
			if err != nil {
				t.Fatalf("Query: %v", err)
			}
			if len(series) != 1 || series[0].Value != 5 {
				t.Fatalf("unexpected series: %+v", series)
			}
			if ex.Strip == nil || *ex.Strip != tc.wantStripRes {
				t.Errorf("resolved strip = %v, want %v", ex.Strip, tc.wantStripRes)
			}
			last := svc.latestQuery()
			if tc.notWantSentER != "" && strings.Contains(last, tc.notWantSentER) {
				t.Errorf("query %q should not contain %q", last, tc.notWantSentER)
			}
			if !strings.Contains(last, tc.wantSentER) {
				t.Errorf("query %q should contain %q", last, tc.wantSentER)
			}
		})
	}
}

func TestExecutorForcedSuffixNoProbe(t *testing.T) {
	svc := &promServer{name: "ccip_commit_plugin_heartbeat", suffixExposed: true}
	srv := httptest.NewServer(svc.handler())
	defer srv.Close()
	ex := NewExecutor(srv.URL, 5*time.Second)
	ex.SetStrip(true) // --suffix strip

	q := "ccip_commit_plugin_heartbeat_total{x=\"1\"}"
	if _, err := ex.Query(context.Background(), q); err != nil {
		t.Fatalf("Query: %v", err)
	}
	if got := svc.latestQuery(); strings.Contains(got, "_total") {
		t.Errorf("query should strip _total under --suffix strip, got %q", got)
	}
}

func TestExecutorTransportError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer srv.Close()
	ex := NewExecutor(srv.URL, 5*time.Second)
	if _, err := ex.Query(context.Background(), "x{}"); err == nil {
		t.Errorf("expected error for 500, got nil")
	}
}
