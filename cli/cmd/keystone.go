package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	skipSetupKey                = "skip-setup"
	cleanKey                    = "clean"
	printConfigKey              = "print-config"
	retryAttemptsKey            = "retry-attempts"
	retryDelayKey               = "retry-delay"
	lokiEndpointKey             = "loki-endpoint"
	testDelayKey                = "test-delay"
	chainlinkCodeDirKey         = "chainlink-code-dir"
	goreleaserKeyKey            = "goreleaser-key"
	chainlinkHelmRegistryURIKey = "chainlink-helm-registry-uri"
)

var keystoneCmd = &cobra.Command{
	Use:   "keystone",
	Short: "Manage keystone operations",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// check if goreleaser key is set
		if viper.GetString(goreleaserKeyKey) == "" {
			return fmt.Errorf("goreleaser key is required")
		}

		return nil
	},
	Run: startCmd.Run, // Default to startCmd when no subcommand is provided
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start keystone and run a smoke test",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		// The user's mental model is that they will be executing this command from the root of the crib repository
		clCodeDir := filepath.Join("../..", viper.GetString(chainlinkCodeDirKey))
		config := utils.StartCmdConfig{
			SkipSetup:                viper.GetBool(skipSetupKey),
			Clean:                    viper.GetBool(cleanKey),
			RetryAttempts:            viper.GetInt(retryAttemptsKey),
			RetryDelay:               viper.GetDuration(retryDelayKey),
			LokiEndpoint:             viper.GetString(lokiEndpointKey),
			TestDelay:                viper.GetDuration(testDelayKey),
			ChainlinkCodeDir:         clCodeDir,
			GoreleaserKey:            viper.GetString(goreleaserKeyKey),
			ChainlinkHelmRegistryURI: viper.GetString(chainlinkHelmRegistryURIKey),
		}
		logger.Info("starting keystone with smoke test check",
			slog.Int("retry_attempts", config.RetryAttempts),
			slog.Duration("retry_delay", config.RetryDelay),
			slog.String("loki_endpoint", config.LokiEndpoint),
			slog.Bool("skip_setup", config.SkipSetup),
			slog.Bool("clean", config.Clean),
			slog.Duration("test_delay", config.TestDelay),
			slog.String("chainlink_code_dir", config.ChainlinkCodeDir),
			slog.String("chainlink_helm_registry_uri", config.ChainlinkHelmRegistryURI))

		if viper.GetBool(printConfigKey) {
			if err := utils.PrintDevspaceChainlinkClusterDep(ctx, config); err != nil {
				logger.Error("failed to print config", slog.Any("error", err))
				os.Exit(1)
			}
			return
		}

		err := utils.RunKeystoneKind(ctx, config)
		if err != nil {
			logger.Error("keystone smoke test failed",
				slog.Any("error", err))

			if os.Getenv("CRIB_ALERT_SLACK_WEBHOOK") != "" && os.Getenv("SEND_ALERTS") == "true" {
				logger.Info("sending slack notification for failure")
				if notifyErr := utils.SendSlackNotification(os.Getenv("CRIB_ALERT_SLACK_WEBHOOK"), "queryLoki"); notifyErr != nil {
					logger.Error("failed to send slack notification",
						slog.Any("error", notifyErr))
				}
			}
			os.Exit(1)
		}

		logger.Info("keystone smoke test completed successfully")
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop keystone by purging the kind cluster",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.PurgeKindCluster()
		if err != nil {
			logger.Error("failed to stop keystone",
				slog.Any("error", err))
			os.Exit(1)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(keystoneCmd)
	// Start is the default command
	keystoneCmd.AddCommand(startCmd)

	flags := keystoneCmd.Flags()
	flags.BoolP(skipSetupKey, "s", false, "Skip setup steps, and immediately run the smoke test instead")
	flags.BoolP(cleanKey, "c", false, "Start with a clean environment")
	flags.IntP(retryAttemptsKey, "r", 10, "Number of retry attempts")
	flags.DurationP(retryDelayKey, "d", 5*time.Second, "Delay between retries")
	flags.StringP(lokiEndpointKey, "l", "https://loki.localhost", "Loki endpoint URL")
	flags.DurationP(testDelayKey, "t", 2*time.Minute, "Delay before testing")
	flags.BoolP(printConfigKey, "p", false, "Print the configuration and exit")
	flags.StringP(chainlinkCodeDirKey, "C", "..", "Directory that contains the /chainlink repository")
	flags.StringP(goreleaserKeyKey, "g", "", "Goreleaser key")
	flags.StringP(chainlinkHelmRegistryURIKey, "H", "", "The URI of the Helm registry for infra-charts, if left empty, charts will be pulled from git")
	_ = viper.BindPFlags(keystoneCmd.Flags())
	startCmd.Flags().AddFlagSet(flags)

	keystoneCmd.AddCommand(stopCmd)

	// Set a key replacer to handle hyphens in flag names for environment variables
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}
