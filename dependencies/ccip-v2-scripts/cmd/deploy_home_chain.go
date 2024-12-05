package cmd

import (
	"fmt"

	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/scripts"
	"github.com/spf13/cobra"
)

var deployHomeOutputDirPath string

// deployHomeChainCmd represents the deployHomeChain command
var deployHomeChainCmd = &cobra.Command{
	Use:   "deploy-home-chain",
	Short: "Deploy home chain contracts",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		devspaceEnv := config.NewDevspaceEnvFromEnv()

		fmt.Printf("deployHomeChain called with %vs\n", devspaceEnv)

		scripts.DeployHomeChain(logger, devspaceEnv)
	},
}

//nolint:gochecknoinits
func init() {
	deployHomeChainCmd.Flags().StringVar(&deployHomeOutputDirPath, "deploy-home-out", "/tmp", "Specify the output dir path")

	rootCmd.AddCommand(deployHomeChainCmd)
}
