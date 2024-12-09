package logger

import (
	cc "github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/libocr/commontypes"
)

// conforms to the logger interface in chainlink-common
var _ cc.Logger = Logger{}

func appendArgs(lggr cc.Logger, newArgs ...interface{}) Logger {
	var args []interface{}

	switch t := lggr.(type) {
	case Logger:
		args = t.args
		lggr = t.Logger
	}

	args = append(args, newArgs...)

	newLogger := cc.With(lggr, args...)

	return Logger{
		Logger: newLogger,
		args:   args,
	}
}

func NewProcessorLogWrapper(lggr cc.Logger, processor string) Logger {
	return appendArgs(lggr, "processor", processor)
}

func NewPluginLogWrapper(lggr cc.Logger, plugin string, donID uint32, oracleID commontypes.OracleID) Logger {
	return appendArgs(lggr, "plugin", plugin, "oracleID", oracleID, "donID", donID)
}

// Named functionality from chainlink-common.
func Named(lggr cc.Logger, n string) Logger {
	var args []interface{}

	switch t := lggr.(type) {
	case Logger:
		lggr = t.Logger
		args = t.args
	}

	return Logger{
		Logger: cc.Named(lggr, n),
		args:   args,
	}
}

// Logger embeds a pointer for copy safety.
type Logger struct {
	cc.Logger
	args []interface{}
}

// WithSeqNr replaces the sequence number. This is probably a bit janky, but
// I mainly want to add something to the plugin logger so that it is a superset
// of the common logger.
func (l *Logger) WithSeqNr(seqNr int) {
	l.Logger = cc.With(l.Logger, "seqNr", seqNr)
}
