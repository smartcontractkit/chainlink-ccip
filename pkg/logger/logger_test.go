package logger

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"
)

func TestLogWrapper(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)

	expected := map[string]interface{}{
		"donID":    uint32(2),
		"oracleID": commontypes.OracleID(3),
		"plugin":   "TestPlugin",
	}
	var donID uint32 = 2
	var oracleID commontypes.OracleID = 3

	// Initial wrapping of base logger.
	wrapped := WithPluginConstants(lggr, "TestPlugin", donID, oracleID)
	wrapped.Info("Where do thumbs hang out at work?")
	require.Equal(t, hook.Len(), 1)
	require.Len(t, hook.All()[0].Context, len(expected))
	require.Equal(t, expected, hook.All()[0].ContextMap())

	// Second wrapping.
	wrapped2 := WithProcessor(wrapped, "TestProcessor")
	expected["processor"] = "TestProcessor"
	wrapped2.Info("The space bar.")
	require.Equal(t, hook.Len(), 2)
	require.Len(t, hook.All()[1].Context, len(expected))
	require.Equal(t, expected, hook.All()[1].ContextMap())
}

func TestNamed(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)

	// Name the base logger.
	namedLggr := Named(lggr, "ElToroLoco")
	namedLggr.Info("Monster Jam")
	require.Equal(t, 1, hook.Len())
	require.Equal(t, "ElToroLoco", hook.All()[0].LoggerName)

	// Name the named logger.
	namedLggr2 := Named(namedLggr, "ObiWan")
	namedLggr2.Info("Star Wars")

	require.Equal(t, 2, hook.Len())
	require.Equal(t, "ElToroLoco.ObiWan", hook.All()[1].LoggerName)
}

func TestLogPointers(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)
	namedLggr := Named(lggr, "ElToroLoco")
	namePtr := Named(&namedLggr, "ObiWan")

	namePtr.Info("Hello1")
	require.Equal(t, 1, hook.Len())
	require.Equal(t, "ElToroLoco.ObiWan", hook.All()[0].LoggerName)

	ptrWrap := WithProcessor(namePtr, "Star Wars")
	ptrWrap.Info("Hello2")
	require.Equal(t, 2, hook.Len())
	require.Equal(t, "ElToroLoco.ObiWan", hook.All()[1].LoggerName)
	require.Equal(t, "Star Wars", hook.All()[1].ContextMap()["processor"])
}

func TestLogCopy(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)

	expected := map[string]interface{}{
		"donID":    uint32(2),
		"oracleID": commontypes.OracleID(3),
		"plugin":   "TestPlugin",
	}
	var donID uint32 = 2
	var oracleID commontypes.OracleID = 3

	// Initial wrapping of base logger.
	wrapped := WithPluginConstants(lggr, "TestPlugin", donID, oracleID)
	wrapped.Info("Where do thumbs hang out at work?")
	require.Equal(t, hook.Len(), 1)
	require.Equal(t, expected, hook.All()[0].ContextMap())

	// Second wrapping.
	logCopy := wrapped
	wrapped2 := WithProcessor(logCopy, "TestProcessor")
	expected["processor"] = "TestProcessor"
	wrapped2.Info("The space bar.")
	require.Equal(t, hook.Len(), 2)
	require.Len(t, hook.All()[1].Context, len(expected))
	require.Equal(t, expected, hook.All()[1].ContextMap())
}
