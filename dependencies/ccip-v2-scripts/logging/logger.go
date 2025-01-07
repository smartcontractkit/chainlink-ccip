package logging

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewConsoleLogger creates a Zap logger with human-readable console output
func NewConsoleLogger() *zap.SugaredLogger {
	// Define encoder config for human-readable output
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:          "time",                           // Timestamp field
		LevelKey:         "level",                          // Log level field
		NameKey:          "logger",                         // Logger name field
		CallerKey:        "caller",                         // Caller field
		MessageKey:       "msg",                            // Log message field
		StacktraceKey:    "stacktrace",                     // Stack trace field
		LineEnding:       zapcore.DefaultLineEnding,        // Default line ending
		EncodeLevel:      zapcore.CapitalColorLevelEncoder, // Level with color
		EncodeDuration:   zapcore.StringDurationEncoder,    // Duration as string
		EncodeCaller:     zapcore.ShortCallerEncoder,       // Caller in short format
		ConsoleSeparator: " | ",                            // Add a custom separator
	}

	// Create core
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zap.DebugLevel)

	// Return the logger
	zapLogger := zap.New(core, zap.AddCaller())
	return zapLogger.Sugar()
}
