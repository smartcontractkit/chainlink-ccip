package cmd

import (
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/scripts"
	"github.com/spf13/cobra"
)

var generateInitialDONOverrides = &cobra.Command{
	Use:   "generate-initial-don-overrides",
	Short: "Generate initial don-overrides",
	Long:  `Generate initial don-overrides are supplied to DON during the first DON deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		config.BindEnvVars()
		devspaceEnv := config.NewDevspaceEnvFromEnv()
		scripts.GenerateInitialNodeOverrides(devspaceEnv)
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(generateInitialDONOverrides)
}
