package logutil

import (
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"
)

// WithProcessor adds the "processor" field. Do not call multiple times.
func WithProcessor(lggr logger.Logger, processor string) logger.Logger {
	return logger.With(lggr, "processor", processor)
}

// WithPluginConstants adds the plugin name, donID and oracleID. Maybe more in the future.
// Do not call multiple times.
func WithPluginConstants(lggr logger.Logger, plugin string, donID uint32, oracleID commontypes.OracleID) logger.Logger {
	return logger.With(lggr, "plugin", plugin, "oracleID", oracleID, "donID", donID)
}
