package wlog

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestIt(t *testing.T) {
	baseLogger := logger.Test(t)
	w := NewWrappedLogger(baseLogger)

	w.WithField("foo", "bar")
	w.WithField("baz", 123)
	w.WithField("str", map[string]any{"a": 1, "b": big.NewInt(123)})

	t.Logf("------------------")

	baseLogger.Debug("some message", "some other message")
	w.Debug("some message", "some other message")

	t.Logf("------------------")

	baseLogger.Debugf("hello %d %s", 123, "world")
	w.Debugf("hello %d %s", 123, "world")

	t.Logf("------------------")

	baseLogger.Debugw("hello friend", "foo", 123, "bar", "baz")
	w.Debugw("hello friend", "foo", 123, "bar", "baz")
}
