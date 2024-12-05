package cmd

import (
	"fmt"

	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/scripts"
	"github.com/spf13/cobra"
)

var (
	deployCCIPInputDirPath  string
	deployCCIPOutputDirPath string
)

// deployCcipCmd represents the deployCcip command
var deployCcipCmd = &cobra.Command{
	Use:   "deploy-ccip",
	Short: "Deploy CCIP Contracts and add Lanes",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		devspaceEnv := config.NewDevspaceEnvFromEnv()

		fmt.Printf("deployCCIP called with %v\n", devspaceEnv)
		scripts.DeployCCIPAndAddLanes(logger, devspaceEnv, deployCCIPInputDirPath)
	},
}

//nolint:gochecknoinits
func init() {
	deployCcipCmd.Flags().StringVar(&deployCCIPInputDirPath, "deploy-ccip-out", "/tmp", "Specify the output dir path")
	deployCcipCmd.Flags().StringVar(&deployCCIPOutputDirPath, "deploy-ccip-in", "/tmp", "Specify the input dir path")

	rootCmd.AddCommand(deployCcipCmd)
}
