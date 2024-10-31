package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-git/go-git/v5"
	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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

		stsClient := wrappers.NewSTSClientWrapper(awsSdkConfig)
		if err := utils.EnsureValidAwsSession(stsClient, viper.GetString("AWS_CONFIG_FILE"), viper.GetString("AWS_PROFILE"), !viper.GetBool("CRIB_CI_ENV")); err != nil {
			logger.Error("failed to get a valid AWS session", slog.Any("error", err))
			os.Exit(1)
		}
		logger.Info("AWS credentials working.")

		var dockerCli wrappers.DockerCLI
		var helmRegistryClient wrappers.HelmRegistryAPI

		if viper.GetBool("docker") && !viper.GetBool("CRIB_SKIP_DOCKER_ECR_LOGIN") {
			logger.Info("refreshing ECR credentials for docker")
			dockerCli, err = utils.InitializeDockerCLI()
			if err != nil {
				logger.Error("failed to initialize Docker CLI", slog.Any("error", err))
				os.Exit(1)
			}
		} else {
			logger.Info("Skipping Docker ECR login")
		}

		if viper.GetBool("helm") && !viper.GetBool("CRIB_SKIP_HELM_ECR_LOGIN") {
			logger.Info("refreshing ECR credentials for helm registry")
			helmRegistryClient, err = utils.InitializeHelmRegistryClient(nil)
			if err != nil {
				logger.Error("failed to initialize Helm Registry Client", slog.Any("error", err))
				os.Exit(1)
			}
		} else {
			logger.Info("Skipping Helm Registry ECR login")
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

// ensureNamespaceCmd represents the devspace ensure-namespace subcommand
var ensureNamespaceCmd = &cobra.Command{
	Use:   "ensure-namespace",
	Short: "Ensure the k8s namespace exists and - when PROVIDER=aws - the kyverno-generated power user rolebinding is in place",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.IsValidCribNamespace(viper.GetString("DEVSPACE_NAMESPACE"), viper.GetBool("CRIB_IGNORE_NAMESPACE_PREFIX")); err != nil {
			logger.Error("invalid namespace for CRIB", slog.Any("error", err))
			os.Exit(1)
		}

		kubeconfig, err := clientcmd.BuildConfigFromFlags("", viper.GetString("KUBECONFIG"))
		if err != nil {
			logger.Error("failed to initialize kubeconfig", slog.Any("error", err))
			os.Exit(1)
		}

		kubeClientset, err := kubernetes.NewForConfig(kubeconfig)
		if err != nil {
			logger.Error("failed to initialize kube clientset", slog.Any("error", err))
			os.Exit(1)
		}

		if err := utils.CheckK8sAccess(kubeClientset.CoreV1()); err != nil {
			msg := "k8s access not working."
			if !viper.GetBool("CRIB_CI_ENV") {
				msg = fmt.Sprintf("%s Make sure you're connected to the VPN and try again.", msg)
			}
			logger.Error(
				msg,
				slog.String("kubeconfig", viper.GetString("KUBECONFIG")),
				slog.String("kubecontext", viper.GetString("CRIB_EKS_ALIAS_NAME")),
				slog.Any("error", err),
			)
			os.Exit(1)
		}

		logger.Info(
			"k8s access working",
			slog.String("kubeconfig", viper.GetString("KUBECONFIG")),
			slog.String("kubecontext", viper.GetString("CRIB_EKS_ALIAS_NAME")),
		)

		created, err := utils.EnsureNamespaceExists(context.TODO(), kubeClientset.CoreV1().Namespaces(), viper.GetString("DEVSPACE_NAMESPACE"))
		if err != nil {
			logger.Error("failed to ensure namespace existence", slog.String("name", viper.GetString("DEVSPACE_NAMESPACE")), slog.Bool("already_exists", !created), slog.Any("error", err))
			os.Exit(1)
		}
		logger.Debug("k8s namespace in place", slog.String("name", viper.GetString("DEVSPACE_NAMESPACE")), slog.Bool("already_exists", !created))

		if viper.GetString("PROVIDER") == "aws" {
			dynamicClient, err := dynamic.NewForConfig(kubeconfig)
			if err != nil {
				logger.Error("failed to initialize kube dynamic client", slog.Any("error", err))
				os.Exit(1)
			}

			rolebindingClient := dynamicClient.Resource(schema.GroupVersionResource{
				Group:    "rbac.authorization.k8s.io",
				Version:  "v1",
				Resource: "rolebindings",
			}).Namespace(viper.GetString("DEVSPACE_NAMESPACE"))

			// Use the WaitForResource function to check if a RoleBinding exists
			roleBindingName := fmt.Sprintf("%s-crib-poweruser", viper.GetString("DEVSPACE_NAMESPACE"))
			timeout := 20 * time.Second
			if err := utils.WaitForResource(context.TODO(), rolebindingClient, roleBindingName, 2*time.Second, timeout); err != nil {
				logger.Error("timed out waiting for role binding to be created", slog.String("role_binding_name", roleBindingName), slog.Duration("timeout", timeout), slog.String("namespace", viper.GetString("DEVSPACE_NAMESPACE")), slog.Any("error", err))
				os.Exit(1)
			}
			logger.Info("role binding found", slog.String("role_binding_name", roleBindingName), slog.String("namespace", viper.GetString("DEVSPACE_NAMESPACE")))
		}
		logger.Info("k8s namespace ready", slog.String("name", viper.GetString("DEVSPACE_NAMESPACE")))
	},
}

// checkEnvVarsCmd represents the devspace check-env-vars subcommand
var checkEnvVarsCmd = &cobra.Command{
	Use:       "check-env-vars [core|ccip]",
	Short:     "Ensure all required environment variables are set for the given product (core|ccip). core is the default product",
	Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.MaximumNArgs(1)),
	ValidArgs: []string{"", "core", "ccip"},
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Running with the following parameters", "config", viper.AllSettings())

		product := "core"
		if len(args) == 1 && args[0] != "" {
			product = args[0]
		}

		logger.Info("checking required env vars", slog.String("product", product))
		requiredEnvVars := []string{
			"IS_CRIB",
			"CHAINLINK_CODE_DIR",
			"DEVSPACE_IMAGE",
			"DEVSPACE_INGRESS_CIDRS",
			"DEVSPACE_INGRESS_BASE_DOMAIN",
			"DEVSPACE_INGRESS_CERT_ARN",
			"DEVSPACE_K8S_POD_WAIT_TIMEOUT",
		}

		switch product {
		case "core":
			requiredEnvVars = append(requiredEnvVars, "CHAINLINK_CLUSTER_HELM_CHART_URI")
		case "ccip":
			requiredEnvVars = append(requiredEnvVars, "CHAINLINK_HELM_REGISTRY_URI")
		}

		if viper.GetString("DEVSPACE_PROFILE") == "keystone" {
			requiredEnvVars = append(requiredEnvVars, "KEYSTONE_ETH_WS_URL", "KEYSTONE_ETH_HTTP_URL", "KEYSTONE_ACCOUNT_KEY")
		}

		missingEnvVars := []string{}
		for _, name := range requiredEnvVars {
			value := os.Getenv(name)
			logger.Debug("reading env var", slog.String("name", name), slog.String("value", value))
			if value == "" {
				missingEnvVars = append(missingEnvVars, name)
			}
		}

		missingEnvVarsCount := len(missingEnvVars)
		if missingEnvVarsCount > 0 {
			logger.Error("missing required environment variables, make sure they're all added ot the config file and try again (check '.env.example' for reference)", slog.Int("count", missingEnvVarsCount), slog.Any("env_vars", missingEnvVars), slog.String("config_file", viper.ConfigFileUsed()), slog.String("product", product))
			os.Exit(1)
		}

		logger.Info("all required environment variables are set", slog.String("product", product), slog.String("config_file", viper.ConfigFileUsed()))
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

	devspaceCmd.AddCommand(ensureNamespaceCmd)
	devspaceCmd.AddCommand(checkEnvVarsCmd)
}
