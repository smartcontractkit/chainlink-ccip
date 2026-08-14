package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

// Series is a single returned time series and its instant value.
type Series struct {
	Labels map[string]string
	Value  float64
}

type instantResponse struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  valueTuple        `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

// valueTuple is the [unixSeconds, "<value>"] pair Prometheus returns.
type valueTuple []any

func (v valueTuple) float() (float64, error) {
	if len(v) < 2 {
		return 0, fmt.Errorf("malformed value tuple %v", []any(v))
	}
	s, ok := v[1].(string)
	if !ok {
		// VictoriaMetrics may return a number; tolerate either.
		if f, ok := v[1].(float64); ok {
			return f, nil
		}
		return 0, fmt.Errorf("value not a string or number: %v", v[1])
	}
	return strconv.ParseFloat(s, 64)
}

// Executor runs instant PromQL queries against one datasource, applying an
// optional counter-suffix transformation so the doc-form queries (which carry
// `_total`/`_bucket`, i.e. an OTel->prometheusremotewrite pipeline) still match
// a datasource scraped directly from promauto base metric names.
type Executor struct {
	Endpoint string
	Client   *http.Client
	bearer   string
	user     string
	pass     string
	// Strip=true means this datasource exposes base metric names (no _total),
	// so doc-form queries must have their suffixes removed. Nil = probe on
	// first call.
	Strip *bool
}

func NewExecutor(endpoint string, timeout time.Duration) *Executor {
	return &Executor{Endpoint: endpoint, Client: &http.Client{Timeout: timeout}}
}

func (e *Executor) SetBearer(b string)       { e.bearer = b }
func (e *Executor) SetBasicAuth(u, p string) { e.user, e.pass = u, p }
func (e *Executor) SetStrip(b bool)          { e.Strip = &b }

// FatalSeriesErr distinguishes HTTP/transport problems from a legitimately
// empty result vector, so the engine can map them to UNKNOWN rather than 0.
func isTransportError(err error) bool { return err != nil }

// metricNameRe matches a metric-name token immediately followed by `{` or `[`,
// capturing any counter/histogram suffix. Function names (sum, rate...) are
// followed by `(` and label selectors start with `{`, so neither is captured.
var metricNameRe = regexp.MustCompile(`([a-zA-Z_:][a-zA-Z0-9_:]*?)` +
	`(?:_total|_bucket|_count(?:\.[0-9]+)?)?(\s*[{\[])`)

// Query runs one instant query at the current time.
func (e *Executor) Query(ctx context.Context, q string) ([]Series, error) {
	if e.Strip == nil {
		// Auto: probe once; if the datasource is unreachable, default to the
		// as-written query and let the real query surface connectivity.
		if err := e.probeStrip(ctx); err != nil {
			off := false
			e.Strip = &off
		}
	}
	if *e.Strip {
		q = stripSuffixes(q)
	}
	params := url.Values{}
	params.Set("query", q)
	params.Set("time", strconv.FormatInt(time.Now().Unix(), 10))
	return e.do(ctx, params)
}

// stripSuffixes removes _total/_bucket/_count suffixes from metric-name tokens
// so doc-form queries (OTel pipeline, suffix present) match a datasource that
// scrapes base names directly.
func stripSuffixes(q string) string {
	return metricNameRe.ReplaceAllStringFunc(q, func(m string) string {
		sub := metricNameRe.FindStringSubmatch(m)
		name, brace := sub[1], sub[2]
		return name + brace
	})
}

func (e *Executor) do(ctx context.Context, params url.Values) ([]Series, error) {
	u := e.Endpoint + "/api/v1/query?" + params.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	if e.bearer != "" {
		req.Header.Set("Authorization", "Bearer "+e.bearer)
	}
	if e.user != "" {
		req.SetBasicAuth(e.user, e.pass)
	}
	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("datasource: status %s: %s", resp.Status, string(body))
	}
	var ir instantResponse
	if err := json.NewDecoder(resp.Body).Decode(&ir); err != nil {
		return nil, err
	}
	if ir.Status != "success" {
		return nil, fmt.Errorf("query error: %s", ir.Data.ResultType)
	}
	out := make([]Series, 0, len(ir.Data.Result))
	for _, r := range ir.Data.Result {
		f, err := r.Value.float()
		if err != nil {
			return nil, fmt.Errorf("series %v: %w", r.Metric, err)
		}
		out = append(out, Series{Labels: r.Metric, Value: f})
	}
	return out, nil
}

// probeStrip detects whether this datasource scrapes base metric names (no
// _total suffix) vs. an OTel->prometheusremotewrite pipeline (suffix present).
// Best effort: an unreachable/empty datasource leaves Strip=false (as-written
// doc queries) and warns; the caller surfaces it as UNKNOWN, not 0.
func (e *Executor) probeStrip(ctx context.Context) error {
	const probeName = "ccip_commit_plugin_heartbeat"
	params := url.Values{}
	params.Set("query", `{__name__=~"`+probeName+`(_total)?"}`)
	params.Set("time", strconv.FormatInt(time.Now().Unix(), 10))
	series, err := e.do(ctx, params)
	if err != nil {
		off := false
		e.Strip = &off
		return fmt.Errorf("could not reach datasource to probe metric suffix (assuming as-written _total). If queries then come back empty, set --suffix none): %w", err)
	}
	strip := false
	for _, s := range series {
		if name, ok := s.Labels["__name__"]; ok {
			// Any base (unsuffixed) name present => direct scrape => strip.
			if len(name) < 6 || name[len(name)-6:] != "_total" {
				strip = true
				break
			}
		}
	}
	e.Strip = &strip
	return nil
}
