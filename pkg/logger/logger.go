package logger

import (
	cc "github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/libocr/commontypes"
)

// conforms to the logger interface in chainlink-common
var _ cc.Logger = Logger{}

func NewProcessorLogWrapper(lggr cc.Logger, processor string) Logger {
	return Logger{
		Logger: cc.With(lggr, "processor", processor),
	}
}

func NewPluginLogWrapper(lggr cc.Logger, plugin string, donID uint32, oracleID commontypes.OracleID) Logger {
	return Logger{
		Logger: cc.With(lggr, "plugin", plugin, "oracleID", oracleID, "donID", donID),
	}
}

// Named functionality from chainlink-common.
func Named(l Logger, n string) Logger {
	return Logger{
		Logger: cc.Named(l, n),
	}
}

// Logger embeds a pointer for copy safety.
type Logger struct {
	cc.Logger
}

// WithSeqNr replaces the sequence number. This is probably a bit janky, but
// I mainly want to add something to the plugin logger so that it is a superset
// of the common logger.
func (l *Logger) WithSeqNr(seqNr int) {
	l.Logger = cc.With(l.Logger, "seqNr", seqNr)
}
