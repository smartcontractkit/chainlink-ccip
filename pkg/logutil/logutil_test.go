package logutil

import (
	"context"
	"math/rand/v2"
	"testing"

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
