package wlog

import (
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type WrappedLogger struct {
	baseLogger logger.Logger
	data       map[string]any
}

func NewWrappedLogger(baseLogger logger.Logger) *WrappedLogger {
	return &WrappedLogger{
		baseLogger: baseLogger,
		data:       make(map[string]any),
	}
}

func (w *WrappedLogger) WithField(key string, value any) {
	w.data[key] = value
}

func (w *WrappedLogger) dataToString(prefix, suffix string) string {
	extraData := make([]string, 0, len(w.data))
	for k, v := range w.data {
		extraData = append(extraData, fmt.Sprintf("%s=%v", k, v))
	}

	extraDataStr := ""
	if len(extraData) > 0 {
		extraDataStr = prefix + strings.Join(extraData, " ") + suffix
	}

	return extraDataStr
}

func (w *WrappedLogger) Debug(args ...interface{}) {
	w.baseLogger.Debug(w.basic(args...)...)
}

func (w *WrappedLogger) Debugf(format string, values ...interface{}) {
	w.baseLogger.Debugf(format+w.dataToString(" ", ""), values...)
}

func (w *WrappedLogger) Debugw(msg string, keysAndValues ...interface{}) {
	w.baseLogger.Debugw(msg+w.dataToString(" ", ""), keysAndValues...)
}

func (w *WrappedLogger) basic(args ...interface{}) []interface{} {
	return append(args, []interface{}{w.dataToString(" ", "")}...)
}
