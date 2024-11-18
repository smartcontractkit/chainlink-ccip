package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
)

// Allowed values for the "provider" flag
var supportedProviders = []string{"aws", "kind"}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new CRIB installation",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			viper.Set("DEVSPACE_NAMESPACE", args[0])
		}
		logger.Debug("input params for PreRun", "config", viper.AllSettings())

		userWasPrompted := false
		if !viper.GetBool("CRIB_CI_ENV") {
			// DEVSPACE_NAMESPACE and PROVIDER are the only parameters that can be set interactively so
			// the CLI flow doesn't require an extra step
			for key, defaultValue := range map[string]string{"DEVSPACE_NAMESPACE": "", "PROVIDER": "aws"} {
				if viper.GetString(key) == "" {
					userWasPrompted = true
					userInput, err := utils.PromptForInput(key, defaultValue)
					if err != nil {
						logger.Error("failed to prompt for input", slog.Any("error", err))
						os.Exit(1)
					}
					viper.Set(key, userInput)
				}
			}
		} else {
			// CI environment, PROVIDER is always "aws"
			viper.Set("PROVIDER", "aws")
		}

		if !slices.Contains(supportedProviders, viper.GetString("PROVIDER")) {
			logger.Error("unsupported provider", "supportedProviders", supportedProviders)
			os.Exit(1)
		}

		if err := utils.IsValidCribNamespace(viper.GetString("DEVSPACE_NAMESPACE"), viper.GetBool("CRIB_IGNORE_NAMESPACE_PREFIX")); err != nil {
			logger.Error("invalid namespace for CRIB", slog.Any("error", err))
			os.Exit(1)
		}

		if userWasPrompted && viper.GetBool("WRITE_CONFIG") {
			promptedKvs := map[string]string{
				"DEVSPACE_NAMESPACE": viper.GetString("DEVSPACE_NAMESPACE"),
				"PROVIDER":           viper.GetString("PROVIDER"),
			}
			if err := utils.WriteConfig(viper.ConfigFileUsed(), promptedKvs); err != nil {
				logger.Error("failed to write config file", slog.Any("error", err))
				os.Exit(1)
			}
			logger.Info("prompted configs written to .env file", "config_file", viper.ConfigFileUsed())

			// reload configs from .env after writing to it
			if err := viper.ReadInConfig(); err != nil {
				logger.Error("failed to reload the .env config file", slog.Any("error", err))
				os.Exit(1)
			}
		}

		// making sure we've got everything loaded by viper, whilst not enforcing optional params
		// TODO: think about a better way to tell mandatory from optional flags apart
		optionalKeys := []string{"GRAFANA_TOKEN", "DASHBOARD_NAME"}
		missingRequiredFlags := []string{}
		for _, key := range viper.AllKeys() {
			if (strings.HasPrefix(key, "KEYSTONE_") || slices.Contains(optionalKeys, key)) && viper.Get(key) == "" {
				missingRequiredFlags = append(missingRequiredFlags, key)
			}
		}
		if len(missingRequiredFlags) > 0 {
			logger.Error("missing required flags", "flags", missingRequiredFlags)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Running with the following parameters", "config", viper.AllSettings())

		if !viper.GetBool("CRIB_CI_ENV") {
			if err := utils.SetupAwsProfile(
				viper.GetString("AWS_CONFIG_FILE"),
				viper.GetString("AWS_PROFILE"),
				viper.GetString("AWS_ACCOUNT_ID"),
				viper.GetString("AWS_REGION"),
				viper.GetString("AWS_SSO_ROLE_NAME"),
				viper.GetString("AWS_SSO_START_URL"),
			); err != nil {
				logger.Error("failed to setup AWS Profile for CRIB", slog.Any("error", err))
				os.Exit(1)
			}

			logger.Info("AWS config modified",
				"config_file", viper.GetString("AWS_CONFIG_FILE"),
				"profile", viper.GetString("AWS_PROFILE"),
			)
		}

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

		if !viper.GetBool("CRIB_CI_ENV") && !utils.HasValidAwsSession(stsClient) {
			logger.Error("AWS credentials still not detected. Exiting.")
			os.Exit(1)
		}

		if !viper.GetBool("CRIB_CI_ENV") {
			if err := utils.SetupKubeConfig(&utils.SetupKubeConfigInput{
				EksClient:            wrappers.NewEKSClientWrapper(awsSdkConfig),
				KubeconfigPath:       viper.GetString("KUBECONFIG"),
				EksClusterName:       viper.GetString("CRIB_EKS_CLUSTER_NAME"),
				EksAliasName:         viper.GetString("CRIB_EKS_ALIAS_NAME"),
				CribNamespace:        viper.GetString("DEVSPACE_NAMESPACE"),
				AwsProfile:           viper.GetString("AWS_PROFILE"),
				AwsRegion:            viper.GetString("AWS_REGION"),
				ChangeDefaultContext: true,
			}); err != nil {
				logger.Error("failed to setup kubeconfig",
					slog.String("kubeconfig", viper.GetString("KUBECONFIG")),
					slog.Any("error", err),
				)
				os.Exit(1)
			}
			logger.Info("kubeconfig setup complete", "kubeconfig", viper.GetString("KUBECONFIG"))
		}

		switch viper.GetString("PROVIDER") {
		case "kind":
			logger.Info("Skipped k8s access check for provider Kind (make sure you run ./cribbit.sh crib-local for now, kind support for the CLI coming soon)")
		case "aws":
			configFlags := genericclioptions.NewConfigFlags(true)
			kubeconfig := viper.GetString("KUBECONFIG")
			configFlags.KubeConfig = &kubeconfig
			k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
			if err != nil {
				logger.Error("failed to initialize kube client", slog.Any("error", err))
				os.Exit(1)
			}

			if err := k8sClient.CheckAccess(context.TODO()); err != nil {
				msg := "EKS access not working."
				if !viper.GetBool("CRIB_CI_ENV") {
					msg = fmt.Sprintf("%s Make sure you're connected to the VPN and try again.", msg)
				}
				logger.Error(msg, slog.Any("error", err))
				os.Exit(1)
			}

			logger.Info("EKS access working",
				"kubeconfig", viper.GetString("KUBECONFIG"),
				"kubecontext", viper.GetString("CRIB_EKS_ALIAS_NAME"),
			)

			dynamicClient, err := dynamic.NewForConfig(k8sClient.RestConfig())
			if err != nil {
				logger.Error("failed to initialize kube dynamic client", slog.Any("error", err))
				os.Exit(1)
			}

			rolebindingClient := dynamicClient.Resource(schema.GroupVersionResource{
				Group:    "rbac.authorization.k8s.io",
				Version:  "v1",
				Resource: "rolebindings",
			}).Namespace(viper.GetString("DEVSPACE_NAMESPACE"))

			if err := utils.EnsureCribNamespaceReady(context.TODO(), k8sClient, rolebindingClient, viper.GetString("DEVSPACE_NAMESPACE"), viper.GetString("PROVIDER"), nil, nil); err != nil {
				logger.Error("failed to ensure crib namespace ready", slog.Any("error", err))
				os.Exit(1)
			}
			logger.Info("k8s namespace ready", slog.String("name", viper.GetString("DEVSPACE_NAMESPACE")))
		}

		var dockerCli wrappers.DockerCLI
		var helmRegistryClient wrappers.HelmRegistryAPI

		if !viper.GetBool("CRIB_SKIP_DOCKER_ECR_LOGIN") {
			dockerCli, err = utils.InitializeDockerCLI()
			if err != nil {
				logger.Error("failed to initialize Docker CLI", slog.Any("error", err))
				os.Exit(1)
			}
		} else {
			logger.Info("Skipping Docker ECR login")
		}

		if !viper.GetBool("CRIB_SKIP_HELM_ECR_LOGIN") {
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

		logger.Info("CRIB initialization complete")
	},
}

//nolint:gochecknoinits
func init() {
	rootCmd.AddCommand(initCmd)

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		logger.Error("failed to determine user's home dir", slog.Any("error", err))
	}

	// AWS Flags
	initCmd.Flags().String("aws-config-file", filepath.Join(userHomeDir, ".aws", "config"), "Path to AWS config")
	initCmd.Flags().String("aws-profile", "", "AWS Profile name to setup")
	initCmd.Flags().String("aws-account-id", "", "AWS Account ID for the profile")
	initCmd.Flags().String("aws-region", "", "AWS Region")
	initCmd.Flags().String("aws-sso-role-name", "", "AWS SSO Role Name")
	initCmd.Flags().String("aws-sso-start-url", "", "AWS SSO Start URL")

	// K8S flags
	initCmd.Flags().String("kubeconfig", filepath.Join(userHomeDir, ".kube", "config"), "Path to kube config file")
	initCmd.Flags().String("crib-eks-cluster-name", "", "EKS cluster name to point kubeconfig at")
	initCmd.Flags().String("crib-eks-alias-name", "", "EKS alias name (will be used as the name of the kube context)")

	// devspace flags
	initCmd.Flags().String("devspace-namespace", "", "CRIB K8S Namespace (to be used by devspace)")
	initCmd.Flags().String("provider", "", fmt.Sprintf("Provider to initialize (should be one of: %v)", supportedProviders))

	// flow control flags
	initCmd.Flags().Bool("write-config", false, "Persists config acquired interactively back to .env passed via --config (WARNING: comments will be lost!)")

	// bind to viper (we can safely ignore the errors here, as the flags are guaranteed to exist)
	_ = viper.BindPFlag("AWS_CONFIG_FILE", initCmd.Flags().Lookup("aws-config-file"))
	_ = viper.BindPFlag("AWS_PROFILE", initCmd.Flags().Lookup("aws-profile"))
	_ = viper.BindPFlag("AWS_ACCOUNT_ID", initCmd.Flags().Lookup("aws-account-id"))
	_ = viper.BindPFlag("AWS_REGION", initCmd.Flags().Lookup("aws-region"))
	_ = viper.BindPFlag("AWS_SSO_ROLE_NAME", initCmd.Flags().Lookup("aws-sso-role-name"))
	_ = viper.BindPFlag("AWS_SSO_START_URL", initCmd.Flags().Lookup("aws-sso-start-url"))
	_ = viper.BindPFlag("KUBECONFIG", initCmd.Flags().Lookup("kubeconfig"))
	_ = viper.BindPFlag("CRIB_EKS_CLUSTER_NAME", initCmd.Flags().Lookup("crib-eks-cluster-name"))
	_ = viper.BindPFlag("CRIB_EKS_ALIAS_NAME", initCmd.Flags().Lookup("crib-eks-alias-name"))
	_ = viper.BindPFlag("DEVSPACE_NAMESPACE", initCmd.Flags().Lookup("devspace-namespace"))
	_ = viper.BindPFlag("WRITE_CONFIG", initCmd.Flags().Lookup("write-config"))
	_ = viper.BindPFlag("PROVIDER", initCmd.Flags().Lookup("provider"))

	// set defaults
	viper.SetDefault("AWS_CONFIG_FILE", initCmd.Flags().Lookup("aws-config-file").DefValue)
	viper.SetDefault("KUBECONFIG", initCmd.Flags().Lookup("kubeconfig").DefValue)

	// defaults that came from cribbit.sh
	viper.SetDefault("CRIB_EKS_CLUSTER_NAME", "main-stage-cluster")
	viper.SetDefault("CRIB_EKS_ALIAS_NAME", "main-stage-cluster-crib")
}
