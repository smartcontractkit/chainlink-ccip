package cmd

import (
	"os"

	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/logging"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ccip-v2-scripts",
	Short: "Deploy OnChain resources required to fully configure CCIP environments",
	Long:  `CRIB wrapper utilizing chainlink/deployment package for deploying fully configured CCIP DONs`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = logging.NewConsoleLogger()
	},
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
