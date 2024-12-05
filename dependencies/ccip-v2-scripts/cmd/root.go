package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccip-v2-scripts",
	Short: "Deploy OnChain resources required to fully configure CCIP environments",
	Long:  `CRIB wrapper utilizing chainlink/deployment package for deploying fully configured CCIP DONs`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		initLogger()

		MustHaveEnv("DEVSPACE_NAMESPACE")
		MustHaveEnv("PROVIDER")
		MustHaveEnv("DON_BOOT_NODE_COUNT")
		MustHaveEnv("DON_NODE_COUNT")
		MustHaveEnv("DEVSPACE_INGRESS_BASE_DOMAIN")
		MustHaveEnv("TMP_DIR")
	},
}

func MustHaveEnv(key string) string {
	err := viper.BindEnv(key)
	if err != nil {
		panic(err)
	}
	value := viper.GetString(key)
	if value == "" {
		logger.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

//nolint:gochecknoinits
func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initLogger() {
	zapLogger := NewConsoleLogger()

	sugaredLogger := zapLogger.Sugar()
	logger = sugaredLogger
}

// NewConsoleLogger creates a Zap logger with human-readable console output
func NewConsoleLogger() *zap.Logger {
	// Define encoder config for human-readable output
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                           // Timestamp field
		LevelKey:       "level",                          // Log level field
		NameKey:        "logger",                         // Logger name field
		CallerKey:      "caller",                         // Caller field
		MessageKey:     "msg",                            // Log message field
		StacktraceKey:  "stacktrace",                     // Stack trace field
		LineEnding:     zapcore.DefaultLineEnding,        // Default line ending
		EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Level with color
		EncodeDuration: zapcore.StringDurationEncoder,    // Duration as string
		EncodeCaller:   zapcore.ShortCallerEncoder,       // Caller in short format
	}

	// Create core
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zap.DebugLevel)

	// Return the logger
	return zap.New(core, zap.AddCaller())
}
