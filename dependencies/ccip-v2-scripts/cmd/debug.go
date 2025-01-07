package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/smartcontractkit/crib/dependencies/ccip-v2-scripts/gap"
	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("debug called")
		token, err := gap.FetchJWTTokenForGAP(context.Background())
		if err != nil {
			logger.Fatalf("failed to get token: %s", err)
			os.Exit(1)
		}
		_, _, err = gap.CheckToken(token)
		if err != nil {
			logger.Fatalf("failed to check token: %s", err)
			os.Exit(1)
		}
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(debugCmd)
}
