package logutil

import (
	"context"
	"math/rand/v2"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

func TestLogWrapper(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)

	var donID uint32 = 2
	var oracleID commontypes.OracleID = 3
	var digest types.ConfigDigest
	digest[0] = 1
	digest[31] = 31

	expected := map[string]interface{}{
		donIDLoggerKey:        donID,
		oracleIDLoggerKey:     oracleID,
		pluginLoggerKey:       "TestPlugin",
		configDigestLoggerKey: digest.String(),
	}

	// Initial wrapping of base logger.
	wrapped := WithPluginConstants(lggr, "TestPlugin", donID, oracleID, digest)
	wrapped.Info("Where do thumbs hang out at work?")
	require.Equal(t, hook.Len(), 1)
	require.Len(t, hook.All()[0].Context, len(expected))
	require.Equal(t, expected, hook.All()[0].ContextMap())

	// Second wrapping.
	wrapped2 := WithComponent(wrapped, "TestProcessor")
	expected[componentLoggerKey] = "TestProcessor"
	wrapped2.Info("The space bar.")
	require.Equal(t, hook.Len(), 2)
	require.Len(t, hook.All()[1].Context, len(expected))
	require.Equal(t, expected, hook.All()[1].ContextMap())
}

func TestLogCopy(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)

	var donID uint32 = 2
	var oracleID commontypes.OracleID = 3
	var digest types.ConfigDigest
	digest[0] = 1
	digest[31] = 31

	expected := map[string]interface{}{
		donIDLoggerKey:        donID,
		oracleIDLoggerKey:     oracleID,
		pluginLoggerKey:       "TestPlugin",
		configDigestLoggerKey: digest.String(),
	}

	// Initial wrapping of base logger.
	wrapped := WithPluginConstants(lggr, "TestPlugin", donID, oracleID, digest)
	wrapped.Info("Where do thumbs hang out at work?")
	require.Equal(t, hook.Len(), 1)
	require.Equal(t, expected, hook.All()[0].ContextMap())

	// Second wrapping.
	logCopy := wrapped
	wrapped2 := WithComponent(logCopy, "TestProcessor")
	expected[componentLoggerKey] = "TestProcessor"
	wrapped2.Info("The space bar.")
	require.Equal(t, hook.Len(), 2)
	require.Len(t, hook.All()[1].Context, len(expected))
	require.Equal(t, expected, hook.All()[1].ContextMap())
}

func Test_GetSeqNr(t *testing.T) {
	ctx := context.TODO()
	seqNr := rand.Uint64()
	phase := PhaseObservation

	// Add sequence number to context
	ctxWithSeqNr := ctxWithOCRVals(ctx, seqNr, phase)

	// Retrieve sequence number from context
	retrievedSeqNr := GetSeqNr(ctxWithSeqNr)
	require.Equal(t, seqNr, retrievedSeqNr, "The sequence number should match the one set in the context")

	retreivedPhase := ctxWithSeqNr.Value(ocrPhaseKey).(string)
	require.Equal(t, phase, retreivedPhase, "The phase should match the one set in the context")

	// Retrieve sequence number from context without sequence number
	retrievedSeqNr = GetSeqNr(ctx)
	require.Equal(t, uint64(0), retrievedSeqNr)
}

func Test_WithContextValues(t *testing.T) {
	ctx := context.TODO()
	seqNr := rand.Uint64()
	phase := PhaseObservation

	// Add sequence number to context
	ctxWithSeqNr := ctxWithOCRVals(ctx, seqNr, phase)

	lggr, observed := logger.TestObserved(t, zapcore.InfoLevel)
	lggr = WithContextValues(ctxWithSeqNr, lggr)

	lggr.Info("Test message")

	seqNrLogs := observed.FilterField(
		zapcore.Field{
			Key:     ocrSeqNrLoggerKey,
			Type:    zapcore.Uint64Type,
			Integer: int64(seqNr),
		},
	)
	require.Len(t, seqNrLogs.All(), 1)
}

func Test_WithOCRSeqNr(t *testing.T) {
	ctx := context.TODO()
	seqNr := rand.Uint64()
	phase := PhaseObservation

	lggr, observed := logger.TestObserved(t, zapcore.InfoLevel)

	newCtx, newLggr := WithOCRInfo(ctx, lggr, seqNr, phase)

	// Check if the context has the correct sequence number
	retrievedSeqNr := GetSeqNr(newCtx)
	require.Equal(t, seqNr, retrievedSeqNr, "The sequence number should match the one set in the context")

	// Check if the logger has the correct sequence number in its fields
	newLggr.Info("Test message")
	seqNrLogs := observed.FilterField(
		zapcore.Field{
			Key:     ocrSeqNrLoggerKey,
			Type:    zapcore.Uint64Type,
			Integer: int64(seqNr),
		},
	)
	require.Len(t, seqNrLogs.All(), 1)

	phaseLogs := observed.FilterField(
		zapcore.Field{
			Key:    ocrPhaseLoggerKey,
			Type:   zapcore.StringType,
			String: phase,
		},
	)
	require.Len(t, phaseLogs.All(), 1)
}

func TestPlugin_logWhenExceedFrequency(t *testing.T) {
	testCases := []struct {
		name          string
		frequency     time.Duration
		calls         []time.Duration // delays between calls
		expectedCalls int             // expected number of times logFunc is called
		description   string
	}{
		{
			name:          "first call always logs",
			frequency:     time.Minute,
			calls:         []time.Duration{0},
			expectedCalls: 1,
			description:   "First call should always execute logFunc",
		},
		{
			name:          "second call too soon",
			frequency:     time.Minute,
			calls:         []time.Duration{0, time.Second},
			expectedCalls: 1,
			description:   "Second call within frequency should not execute logFunc",
		},
		{
			name:          "second call after frequency",
			frequency:     100 * time.Millisecond,
			calls:         []time.Duration{0, 150 * time.Millisecond},
			expectedCalls: 2,
			description:   "Second call after frequency should execute logFunc",
		},
		{
			name:      "multiple calls mixed timing",
			frequency: 100 * time.Millisecond,
			calls: []time.Duration{
				0,                     // first call, should log
				50 * time.Millisecond, // too soon
				60 * time.Millisecond, // total 110ms, should log and reset timer
				90 * time.Millisecond, // too soon again
				10 * time.Millisecond, // total 100ms, should log and reset timer
			},
			expectedCalls: 3,
			description:   "Only calls that exceed frequency should execute logFunc",
		},
		{
			name:          "zero frequency",
			frequency:     0,
			calls:         []time.Duration{0, 0, 0},
			expectedCalls: 3,
			description:   "Zero frequency should allow all calls",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var lastLog atomic.Pointer[time.Time]
			callCount := 0

			logFunc := func() {
				callCount++
			}

			// Execute calls with specified delays
			for i, delay := range tt.calls {
				if i > 0 {
					time.Sleep(delay)
				}
				LogWhenExceedFrequency(&lastLog, tt.frequency, logFunc)
			}

			require.Equal(t, tt.expectedCalls, callCount, tt.description)
		})
	}
}
