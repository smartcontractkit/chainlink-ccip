package ccip_evm

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/rs/zerolog"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/onramp"
)

const (
	DefaultLokiURL  = "http://localhost:3030/loki/api/v1/push"
	DefaultTempoURL = "http://localhost:4318/v1/traces"
)

/*
Loki labels.
*/
const (
	LokiCCIPMessageSentLabel       = "on-chain-sent"
	LokiExecutionStateChangedLabel = "on-chain-exec"
)

/*
Prometheus metrics.
*/
var (
	msgSentTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "msg_sent_total",
		Help: "Total number of CCIP messages sent",
	}, []string{"from", "to"})
	msgExecTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "msg_exec_total",
		Help: "Total number of CCIP messages executed",
	}, []string{"from", "to"})
	srcDstLatency = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "msg_src_dst_duration_seconds",
		Help:    "Total duration of processing message from src to dst chain",
		Buckets: []float64{1, 2, 5, 10, 15, 20, 30, 45, 60, 90, 120, 180, 240, 300, 400, 500},
	}, []string{"from", "to"})
)

// LaneStreamConfig contains contracts to collect events from and selectors for queries.
type LaneStreamConfig struct {
	FromSelector      uint64
	ToSelector        uint64
	AggregatorAddress string
	AggregatorSince   int64
}

// LaneStreams represents all the events for sent/exec events.
type LaneStreams struct {
	SentEvents []*onramp.OnRampCCIPMessageSent
	ExecEvents []*offramp.OffRampExecutionStateChanged
}

type SentEventPlusMeta struct {
	*onramp.OnRampCCIPMessageSent
	MessageIDHex string
}

type ExecEventPlusMeta struct {
	*offramp.OffRampExecutionStateChanged
	MessageIDHex string
}

func ToAnySlice[T any](slice []T) []any {
	result := make([]any, len(slice))
	for i, v := range slice {
		result[i] = v
	}
	return result
}

// ProcessLaneEvents collects, pushes and observes sent and executed messages for lane.
func ProcessLaneEvents(ctx context.Context, c *CCIP16EVM, lp *LokiPusher, cfg *LaneStreamConfig) error {
	lggr := zerolog.Ctx(ctx)
	lggr.Info().Uint64("FromSelector", cfg.FromSelector).Uint64("ToSelector", cfg.ToSelector).Msg("Processing events")
	streams, err := FetchLaneEvents(ctx, c, cfg)
	if err != nil {
		return err
	}
	fromSelectorStr := fmt.Sprintf("%d", cfg.FromSelector)
	toSelectorStr := fmt.Sprintf("%d", cfg.ToSelector)
	// push Loki streams
	if err := lp.Push(ToAnySlice(streams.SentEvents), map[string]string{
		"job":  LokiCCIPMessageSentLabel,
		"from": fromSelectorStr,
		"to":   toSelectorStr,
	}); err != nil {
		return err
	}
	if err := lp.Push(ToAnySlice(streams.ExecEvents), map[string]string{
		"job":  LokiExecutionStateChangedLabel,
		"from": fromSelectorStr,
		"to":   toSelectorStr,
	}); err != nil {
		return err
	}

	// observe as Prometheus metrics
	logTimeByMsgID := make(map[[32]byte]uint64)
	for _, l := range streams.SentEvents {
		_ = l
		// msgSentTotal.WithLabelValues(fromSelectorStr, toSelectorStr).Inc()
	}
	for _, l := range streams.ExecEvents {
		_ = l
		// blkTimeStarted, ok := logTimeByMsgID[l.MessageId]
		// if !ok {
		// 	continue
		// }
		// elapsed := l.Raw.BlockTimestamp - blkTimeStarted
		// lggr.Debug().
		// 	Any("MsgID", hexutil.Encode(l.MessageId[:])).
		// 	Uint64("Seconds", elapsed).
		// 	Msg("Elapsed time")
		// srcDstLatency.WithLabelValues(fromSelectorStr, toSelectorStr).Observe(float64(elapsed))
		// msgExecTotal.WithLabelValues(fromSelectorStr, toSelectorStr).Inc()
	}
	_ = logTimeByMsgID
	return nil
}

// FetchLaneEvents fetch sent and exec events for lane.
func FetchLaneEvents(ctx context.Context, c *CCIP16EVM, cfg *LaneStreamConfig) (*LaneStreams, error) {
	return &LaneStreams{}, nil
}

// func addSentMetadata(msgs []*onramp.OnRampCCIPMessageSent) []*SentEventPlusMeta {
// 	events := make([]*SentEventPlusMeta, 0)
// 	for _, msg := range msgs {
// 		events = append(events, &SentEventPlusMeta{
// 			OnRampCCIPMessageSent: msg,
// 			MessageIDHex:          hexutil.Encode(msg.MessageId[:]),
// 		})
// 	}
// 	return events
// }

// func addExecMetadata(msgs []*offramp.OffRampExecutionStateChanged) []*ExecEventPlusMeta {
// 	events := make([]*ExecEventPlusMeta, 0)
// 	for _, msg := range msgs {
// 		events = append(events, &ExecEventPlusMeta{
// 			OffRampExecutionStateChanged: msg,
// 			MessageIDHex:                 hexutil.Encode(msg.MessageId[:]),
// 		})
// 	}
// 	return events
// }

// LokiPusher handles pushing logs to Loki
// it does not use Promtail client specifically to avoid dep hell between Prometheus/Loki go deps.
type LokiPusher struct {
	lokiURL string
	client  *resty.Client
}

// LogEntry represents a single log entry for Loki.
type LogEntry struct {
	Timestamp time.Time         `json:"timestamp"`
	Message   any               `json:"message"`
	Labels    map[string]string `json:"labels,omitempty"`
}

// LokiStream represents a stream of log entries with labels.
type LokiStream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"` // [timestamp, log line]
}

// LokiPayload represents the payload structure for Loki API.
type LokiPayload struct {
	Streams []LokiStream `json:"streams"`
}

// NewLokiPusher creates a new LokiPusher instance.
func NewLokiPusher() *LokiPusher {
	lokiURL := os.Getenv("LOKI_URL")
	if lokiURL == "" {
		lokiURL = DefaultLokiURL
	}
	return &LokiPusher{
		lokiURL: lokiURL,
		client:  resty.New().SetTimeout(10 * time.Second),
	}
}

// Push pushes all the messages to a Loki stream.
func (lp *LokiPusher) Push(msgs []any, labels map[string]string) error {
	if len(msgs) == 0 {
		return nil
	}
	values := make([][]string, 0, len(msgs))

	for i := 0; i < len(msgs); i++ {
		combinedMessage := map[string]any{
			"log": msgs[i],
			"ts":  time.Now().Format(time.RFC3339Nano),
		}
		jsonBytes, err := json.Marshal(combinedMessage)
		if err != nil {
			return fmt.Errorf("failed to marshal combined message: %w", err)
		}
		values = append(values, []string{
			fmt.Sprintf("%d", time.Now().UnixNano()),
			string(jsonBytes),
		})
	}

	stream := LokiStream{
		Stream: labels,
		Values: values,
	}
	resp, err := lp.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(LokiPayload{
			Streams: []LokiStream{stream},
		}).
		Post(lp.lokiURL)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	if resp.StatusCode() != 200 && resp.StatusCode() != 204 {
		return fmt.Errorf("loki returned status %d: %s", resp.StatusCode(), resp.String())
	}
	return nil
}
