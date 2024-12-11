package logutil

import (
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
		"donID":        donID,
		"oracleID":     oracleID,
		"plugin":       "TestPlugin",
		"configDigest": digest.String(),
	}

	// Initial wrapping of base logger.
	wrapped := WithPluginConstants(lggr, "TestPlugin", donID, oracleID, digest)
	wrapped.Info("Where do thumbs hang out at work?")
	require.Equal(t, hook.Len(), 1)
	require.Len(t, hook.All()[0].Context, len(expected))
	require.Equal(t, expected, hook.All()[0].ContextMap())

	// Second wrapping.
	wrapped2 := WithContext(wrapped, "TestProcessor")
	expected["context"] = "TestProcessor"
	wrapped2.Info("The space bar.")
	require.Equal(t, hook.Len(), 2)
	require.Len(t, hook.All()[1].Context, len(expected))
	require.Equal(t, expected, hook.All()[1].ContextMap())
}

func TestNamed(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)

	// Name the base logger.
	namedLggr := logger.Named(lggr, "ElToroLoco")
	namedLggr.Info("Monster Jam")
	require.Equal(t, 1, hook.Len())
	require.Equal(t, "ElToroLoco", hook.All()[0].LoggerName)

	// Name the named logger.
	namedLggr2 := logger.Named(namedLggr, "ObiWan")
	namedLggr2.Info("Star Wars")

	require.Equal(t, 2, hook.Len())
	require.Equal(t, "ElToroLoco.ObiWan", hook.All()[1].LoggerName)
}

func TestLogCopy(t *testing.T) {
	lggr, hook := logger.TestObserved(t, zapcore.DebugLevel)

	var donID uint32 = 2
	var oracleID commontypes.OracleID = 3
	var digest types.ConfigDigest
	digest[0] = 1
	digest[31] = 31

	expected := map[string]interface{}{
		"donID":        donID,
		"oracleID":     oracleID,
		"plugin":       "TestPlugin",
		"configDigest": digest.String(),
	}

	// Initial wrapping of base logger.
	wrapped := WithPluginConstants(lggr, "TestPlugin", donID, oracleID, digest)
	wrapped.Info("Where do thumbs hang out at work?")
	require.Equal(t, hook.Len(), 1)
	require.Equal(t, expected, hook.All()[0].ContextMap())

	// Second wrapping.
	logCopy := wrapped
	wrapped2 := WithContext(logCopy, "TestProcessor")
	expected["context"] = "TestProcessor"
	wrapped2.Info("The space bar.")
	require.Equal(t, hook.Len(), 2)
	require.Len(t, hook.All()[1].Context, len(expected))
	require.Equal(t, expected, hook.All()[1].ContextMap())
}
