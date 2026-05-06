package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog"
)

// DefaultLokiURL is the default Loki WebSocket URL for log streaming in tests.
// TODO: ideally we'd get this from a config or an env var.
const DefaultLokiURL = "ws://localhost:3030"

// lokiPingQuery is a minimal valid LogQL selector for /tail. Loki rejects an empty `{}`
// selector (HTTP 400 before the WebSocket upgrade, seen as "bad handshake").
// Devenv pushes CCIP streams with a `job` label (see chainimpl monitoring).
const lokiPingQuery = `{job=~".+"}`

type lokiReadResult struct {
	data []byte
	err  error
}

// PingLoki opens a short-lived WebSocket to Loki's tail API to verify the server
// accepts connections at lokiWSURL (same shape as DefaultLokiURL, e.g. ws://host:3030).
func PingLoki(ctx context.Context, lokiWSURL string) error {
	if lokiWSURL == "" {
		return fmt.Errorf("loki WebSocket URL is empty")
	}
	wsURL := fmt.Sprintf("%s/loki/api/v1/tail?query=%s", lokiWSURL, url.QueryEscape(lokiPingQuery))
	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		return lokiHandshakeErr(resp, err)
	}
	return conn.Close()
}

func lokiHandshakeErr(resp *http.Response, err error) error {
	if resp == nil {
		return fmt.Errorf("failed to connect to Loki WebSocket: %w", err)
	}
	defer resp.Body.Close()
	body, readErr := io.ReadAll(io.LimitReader(resp.Body, 512))
	if readErr != nil {
		return fmt.Errorf("failed to connect to Loki WebSocket (status %d): %w", resp.StatusCode, err)
	}
	detail := strings.TrimSpace(string(body))
	if detail != "" {
		return fmt.Errorf("failed to connect to Loki WebSocket (status %d): %s: %w", resp.StatusCode, detail, err)
	}
	return fmt.Errorf("failed to connect to Loki WebSocket (status %d): %w", resp.StatusCode, err)
}

// WaitForLokiLogs waits for logs with the given query to appear in Loki
// and returns the latest log line.
//
// logQuery is the same log query that is used in the Grafana UI, e.g.
// `{container="don-node1"} | json | logger="SomeLoggerName" | msg="SomeMessage"`.
func WaitForLokiLogs(ctx context.Context, lokiWSURL, logQuery string) (string, error) {
	l := zerolog.Ctx(ctx)

	// URL-encode the query and build the tail URL
	encodedQuery := url.QueryEscape(logQuery)
	wsURL := fmt.Sprintf("%s/loki/api/v1/tail?query=%s", lokiWSURL, encodedQuery)

	l.Info().Str("log_query", logQuery).Str("loki_ws_url", wsURL).Msg("Waiting for Loki logs")

	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		return "", lokiHandshakeErr(resp, err)
	}

	var closeOnce sync.Once
	closeConn := func() {
		closeOnce.Do(func() { _ = conn.Close() })
	}
	defer closeConn()

	for {
		readCh := make(chan lokiReadResult, 1)
		go func() {
			_, data, readErr := conn.ReadMessage()
			readCh <- lokiReadResult{data: data, err: readErr}
		}()

		var message []byte
		select {
		case <-ctx.Done():
			closeConn()
			return "", ctx.Err()
		case r := <-readCh:
			if r.err != nil {
				return "", fmt.Errorf("WebSocket read error: %w", r.err)
			}
			message = r.data
		}

		var response lokiStreamResponse
		if err := json.Unmarshal(message, &response); err != nil {
			continue // Skip malformed messages
		}

		if len(response.Streams) == 0 {
			continue // Skip empty responses
		}

		stream := response.Streams[0]
		if len(stream.Values) == 0 {
			continue // Skip empty values
		}

		latestValue := stream.Values[len(stream.Values)-1]
		if len(latestValue) != 2 {
			continue // Skip invalid values
		}

		// 2nd element is the log line.
		latestLog := latestValue[1]
		l.Info().Str("latest_log", latestLog).Msg("Found latest log, returning it")

		return latestLog, nil
	}
}

// lokiStreamResponse represents the JSON structure from Loki's tail API.
type lokiStreamResponse struct {
	Streams []struct {
		// Stream is a map of labels associated with the logs.
		Stream map[string]string `json:"stream"`

		// Values is an array of log entries, where each entry is a two-element array:
		// [<timestamp_nanoseconds>, <log_line_string>]
		Values [][]string `json:"values"`
	} `json:"streams"`
}
