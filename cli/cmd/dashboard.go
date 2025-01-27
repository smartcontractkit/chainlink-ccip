package cmd

import (
	"fmt"
	"os"

	"github.com/smartcontractkit/crib/cli/dashboard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// dashboardCmd represents the dashboard command
var dashboardCmd = &cobra.Command{
	Use:   "dashboard",
	Short: "Manage grafana dashboards",
	Long:  `Manage grafana dashboards created for CRIB environments`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dashboard called")
	},
}

var deployDashboardCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy grafana dashboard",
	Long:  `Deploy grafana dashboard`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("product") == "" {
			logger.Error("product flag is required")
			os.Exit(1)
		}

		if viper.GetString("product") == "ccip" {
			dashboard.DeployCCIP()
		} else if viper.GetString("product") == "core" {
			dashboard.DeployCore()
		}
	},
}

var deleteDashboardCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete grafana dashboard",
	Long:  `Delete grafana dashboard`,
	Run: func(cmd *cobra.Command, args []string) {
		dashboard.Delete()
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(dashboardCmd)

	deployDashboardCmd.Flags().String("product", "chainlink", "Required, Name of the product, in the same form as product dir")
	_ = viper.BindPFlags(deployDashboardCmd.Flags())

	dashboardCmd.AddCommand(deployDashboardCmd)
	dashboardCmd.AddCommand(deleteDashboardCmd)
}
