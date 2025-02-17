package cmd

import (
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/scripts"
	"github.com/spf13/cobra"
)

var configureOCR = &cobra.Command{
	Use:   "configure-ocr",
	Short: "Configure CCIP OCR",
	Long:  `TBD`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitViper()
		devspaceEnv := config.NewDevspaceEnvFromEnv()
		scripts.ConfigureOCR(logger, devspaceEnv)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(configureOCR)
}
