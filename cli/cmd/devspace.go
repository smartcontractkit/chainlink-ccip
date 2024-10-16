package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-git/go-git/v5"
	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// devspaceCmd represents the devspace command
var devspaceCmd = &cobra.Command{
	Use:   "devspace",
	Short: "CRIB logic for Devspace",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("devspace called")
	},
}

// beforeBuildChecksCmd represents the devspace before-build-checks subcommand
var beforeBuildChecksCmd = &cobra.Command{
	Use:   "before-build-checks",
	Short: "CRIB logic for Devspace hooks",
	Run: func(cmd *cobra.Command, args []string) {
		if utils.IsCustomImage(viper.GetString("DEVSPACE_NAMESPACE"), viper.GetString("DEVSPACE_IMAGE")) {
			logger.Error("DEVSPACE_IMAGE var was set to a non-standard image, use --skip-build and -o <image_tag> options to skip the build entirely", slog.String("DEVSPACE_IMAGE", viper.GetString("DEVSPACE_IMAGE")))
			os.Exit(1)
		}

		mandatoryEnvVars := []string{"CHAINLINK_CODE_DIR", "CHAINLINK_REPO_DIR"}
		for _, envVar := range mandatoryEnvVars {
			if os.Getenv(envVar) == "" {
				logger.Error("required environment variable is not set.", slog.String("envVar", envVar))
			}
		}

		gitRepo, err := git.PlainOpen(os.Getenv("CHAINLINK_REPO_DIR"))
		if err != nil {
			logger.Error("failed to open git repository", slog.String("CHAINLINK_REPO_DIR", os.Getenv("CHAINLINK_REPO_DIR")), slog.Any("error", err))
			os.Exit(1)
		}
		currentRef, err := gitRepo.Head()
		if err != nil {
			logger.Error("failed to get current git ref", slog.String("CHAINLINK_REPO_DIR", os.Getenv("CHAINLINK_REPO_DIR")), slog.Any("error", err))
			os.Exit(1)
		}
		logger.Info("Repository exists at", slog.String("CHAINLINK_REPO_DIR", os.Getenv("CHAINLINK_REPO_DIR")), slog.Any("ref", currentRef))
	},
}

// refreshEcrCredentialsCmd represents the devspace refresh-ecr-credentials subcommand
var refreshEcrCredentialsCmd = &cobra.Command{
	Use:   "refresh-ecr-credentials",
	Short: "Refresh ECR credentials for docker and helm registry",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		awsSdkConfig, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithSharedConfigFiles([]string{viper.GetString("AWS_CONFIG_FILE")}),
			config.WithSharedConfigProfile(viper.GetString("AWS_PROFILE")),
		)
		if err != nil {
			logger.Error("failed to load AWS config", slog.Any("error", err))
			os.Exit(1)
		}

		var dockerCli wrappers.DockerCLI
		var helmRegistryClient wrappers.HelmRegistryAPI

		if viper.GetBool("docker") {
			logger.Info("refreshing ECR credentials for docker")
			dockerCli, err = utils.InitializeDockerCLI()
			if err != nil {
				logger.Error("failed to initialize Docker CLI", slog.Any("error", err))
				os.Exit(1)
			}
		}
		if viper.GetBool("helm") {
			logger.Info("refreshing ECR credentials for helm registry")
			helmRegistryClient, err = utils.InitializeHelmRegistryClient(nil)
			if err != nil {
				logger.Error("failed to initialize Helm Registry Client", slog.Any("error", err))
				os.Exit(1)
			}
		}

		ecrClient := wrappers.NewECRClientWrapper(awsSdkConfig)
		refreshRegistriesOutput := utils.RefreshRegistriesECRCredentials(ecrClient, dockerCli, helmRegistryClient, viper.GetString("CHAINLINK_HELM_REGISTRY_URI"))
		if refreshRegistriesOutput.ECRGetAuthorizationTokenError != nil {
			logger.Error("failed to refresh ECR credentials", slog.Any("error", refreshRegistriesOutput.ECRGetAuthorizationTokenError))
			os.Exit(1)
		}

		for _, attempt := range *refreshRegistriesOutput.RegistryLoginAttempts {
			if attempt.LoginErr != nil {
				logger.Error("failed to refresh ECR credentials for registry", slog.String("registry_type", attempt.RegistryType), slog.String("registry_host", attempt.RegistryHost), slog.Any("error", attempt.LoginErr))
				os.Exit(1)
			} else {
				logger.Info("Registry login successful", slog.String("registry_type", attempt.RegistryType), slog.String("registry_host", attempt.RegistryHost))
			}
		}

		logger.Info("ECR credentials refreshed")
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(devspaceCmd)

	devspaceCmd.AddCommand(beforeBuildChecksCmd)
	devspaceCmd.AddCommand(refreshEcrCredentialsCmd)

	refreshEcrCredentialsCmd.Flags().Bool("docker", false, "Refresh docker ECR credentials")
	refreshEcrCredentialsCmd.Flags().Bool("helm", false, "Refresh helm ECR credentials")
	refreshEcrCredentialsCmd.MarkFlagsOneRequired("docker", "helm")

	_ = viper.BindPFlags(refreshEcrCredentialsCmd.Flags())
}
