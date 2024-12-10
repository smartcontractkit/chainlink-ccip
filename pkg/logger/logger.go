package logger

import (
	cc "github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/libocr/commontypes"
)

// conforms to the logger interface in chainlink-common
var _ cc.Logger = Logger{}

func baseLogger(lggr cc.Logger) cc.Logger {
	switch t := lggr.(type) {
	case Logger:
		return t.Logger
	case *Logger:
		return t.Logger
	}
	return lggr
}

func appendArgs(lggr cc.Logger, newArgs ...interface{}) Logger {
	newLogger := cc.With(baseLogger(lggr), newArgs...)
	return Logger{
		Logger: newLogger,
	}
}

// WithProcessor adds the "processor" field. Do not call multiple times.
func WithProcessor(lggr cc.Logger, processor string) Logger {
	return appendArgs(lggr, "processor", processor)
}

// WithPluginConstants adds the plugin name, donID and oracleID. Maybe more in the future.
// Do not call multiple times.
func WithPluginConstants(lggr cc.Logger, plugin string, donID uint32, oracleID commontypes.OracleID) Logger {
	return appendArgs(lggr, "plugin", plugin, "oracleID", oracleID, "donID", donID)
}

// Named functionality from chainlink-common.
func Named(lggr cc.Logger, n string) Logger {
	return Logger{
		Logger: cc.Named(baseLogger(lggr), n),
	}
}

// Logger embeds a pointer for copy safety.
type Logger struct {
	cc.Logger
}

// TODO: I don't think it replaces the field, it appends repeatedly.
/*
// WithSeqNr replaces the sequence number. This is probably a bit janky, but
// I mainly want to add something to the plugin logger so that it is a superset
// of the common logger.
func (l *Logger) WithSeqNr(seqNr int) {
	l.Logger = cc.With(baseLogger(l), "seqNr", seqNr)
}
*/
