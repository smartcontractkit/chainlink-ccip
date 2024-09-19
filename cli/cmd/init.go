package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new CRIB installation",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")

		// TODO: call SetupAwsProfile when applicable
		// e.g. utils.SetupAwsProfile("$HOME/.aws/config", viper.GetString("crib.awsProfileName"), viper.GetString("aws.accountId"), viper.GetString("aws.region"), viper.GetString("aws.ssoRoleName"), viper.GetString("aws.ssoStartUrl"))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
