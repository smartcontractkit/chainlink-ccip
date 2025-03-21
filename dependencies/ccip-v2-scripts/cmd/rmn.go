package cmd

import (
	"fmt"

	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/config"
	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/scripts"
	"github.com/spf13/cobra"
)

// deployCCIPInputDirPath string
var nodeCount int

var rmnCommand = &cobra.Command{
	Use:   "rmn",
	Short: "Manage RMN deployment in CRIB",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmn called")
	},
}

// setupRMNOnChain represents the generate-rmn-identities command
var setupRMNOnChain = &cobra.Command{
	Use:   "setup-rmn-onchain",
	Short: "Setup RMN on chain",
	Long:  `Setup RMN On Chain. Requires CCIP Home chain and infra config to be generated`,
	Run: func(cmd *cobra.Command, args []string) {
		config.BindEnvVars()
		devspaceEnv := config.NewDevspaceEnvFromEnv()

		fmt.Printf("deploy-rmn-contracts called with %v\n", devspaceEnv)

		configurer := scripts.NewRMNConfigurer(devspaceEnv, nodeCount)

		configurer.SetupRMNOnChain()
	},
}

// generateInfraConfig represents the generate-rmn-identities command
var generateInfraConfig = &cobra.Command{
	Use:   "generate-infra-config",
	Short: "Generate RMN Configuration input for deploying charts",
	Long:  `Generates RMN node identities, shared-toml and local-toml based on the data from Env State`,
	Run: func(cmd *cobra.Command, args []string) {
		config.BindEnvVars()
		devspaceEnv := config.NewDevspaceEnvFromEnv()

		fmt.Printf("generate-config called with %v\n", devspaceEnv)

		configurer := scripts.NewRMNConfigurer(devspaceEnv, nodeCount)

		configurer.GenerateNodeIdentities()
		configurer.GenerateTOMLConfigs()
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(rmnCommand)
	rmnCommand.Flags().IntVar(&nodeCount, "node-count", 3, "Specify the RMN node count")

	rmnCommand.AddCommand(setupRMNOnChain)
	rmnCommand.AddCommand(generateInfraConfig)
}
