package logutil

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// Type to use as context keys, regular string type can't be used.
type contextKey string

const (
	// Standardized logger field keys.
	// These should be sufficiently different from the ephemeral log
	// keys that are added inline to log calls, in order to avoid
	// collisions.
	ocrSeqNrLoggerKey     = "ocrSeqNr"
	componentLoggerKey    = "component"
	pluginLoggerKey       = "plugin"
	oracleIDLoggerKey     = "oracleID"
	donIDLoggerKey        = "donID"
	configDigestLoggerKey = "configDigest"

	// ocrSeqNrKey refers to the OCR sequence number.
	// Linter complains if this is a plain "string" type.
	ocrSeqNrKey = contextKey(ocrSeqNrLoggerKey)
)

// WithComponent returns a logger with the component log field set.
// Components can be any object that uses a logger.
// Do not call multiple times.
func WithComponent(lggr logger.Logger, processor string) logger.Logger {
	return logger.With(lggr, componentLoggerKey, processor)
}

// WithPluginConstants adds the plugin name, donID and oracleID. Maybe more in the future.
// Do not call multiple times.
func WithPluginConstants(
	lggr logger.Logger,
	plugin string,
	donID uint32,
	oracleID commontypes.OracleID,
	configDigest types.ConfigDigest,
) logger.Logger {
	return logger.With(
		lggr,
		pluginLoggerKey, plugin,
		oracleIDLoggerKey, oracleID,
		donIDLoggerKey, donID,
		configDigestLoggerKey, configDigest,
	)
}

// WithOCRSeqNr returns a context.Context and logger with the OCR sequence number set
// in both the context and the logger fields.
func WithOCRSeqNr(
	ctx context.Context,
	lggr logger.Logger,
	ocrSeqNr uint64,
) (context.Context, logger.Logger) {
	newCtx := ctxWithOCRSeqNr(ctx, ocrSeqNr)
	newLggr := WithContextValues(newCtx, lggr)
	return newCtx, newLggr
}

// ctxWithOCRSeqNr returns a context.Context with the OCR sequence number appropriately set.
func ctxWithOCRSeqNr(ctx context.Context, seqNr uint64) context.Context {
	return context.WithValue(ctx, ocrSeqNrKey, seqNr)
}

// WithContextValues returns a logger with the OCR sequence number set to the correct log field key.
// It assumes that the given ctx has the OCR sequence number set, if not the field will be set to
// 0 in the log field.
func WithContextValues(
	ctx context.Context,
	lggr logger.Logger,
) logger.Logger {
	return logger.With(
		lggr,
		ocrSeqNrLoggerKey, GetSeqNr(ctx),
	)
}

// GetSeqNr returns the sequence number from the context.
// 0 is returned if the sequence number is not set.
// Note that this isn't confusing because 0 is an invalid sequence number.
func GetSeqNr(ctx context.Context) uint64 {
	seqNr, ok := ctx.Value(ocrSeqNrKey).(uint64)
	if !ok {
		return 0
	}
	return seqNr
}
