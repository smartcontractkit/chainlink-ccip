package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-git/go-git/v5"
	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
)

const (
	CleanupLabelKey   = "cleanup.kyverno.io/ttl"
	CleanupLabelValue = "72h"
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
		skipDocker := viper.GetBool("CRIB_SKIP_DOCKER_ECR_LOGIN")
		skipHelm := viper.GetBool("CRIB_SKIP_HELM_ECR_LOGIN")

		if viper.GetBool("docker") && skipDocker {
			logger.Info("Skipping Docker ECR login")
			if skipHelm {
				logger.Info("Skipping Helm Registry ECR login. Reason: Helm login dependency on Docker login is skipped")
			}

			return
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

		var dockerCli wrappers.DockerCLI
		var helmRegistryClient wrappers.HelmRegistryAPI

		logger.Info("refreshing ECR credentials for docker")
		dockerCli, err = wrappers.NewDockerCli()
		if err != nil {
			logger.Error("failed to initialize Docker CLI", slog.Any("error", err))
			os.Exit(1)
		}

		if viper.GetBool("helm") && !skipHelm {
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
		dockerRegistries := utils.GetChainlinkDockerRegistries(viper.GetString("PROVIDER"))
		refreshRegistriesOutput := utils.RefreshRegistriesECRCredentials(ecrClient, dockerCli, helmRegistryClient, viper.GetString("CHAINLINK_HELM_REGISTRY_URI"), dockerRegistries)
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
		if err := utils.IsValidCribNamespace(viper.GetString("DEVSPACE_NAMESPACE"), viper.GetString("PROVIDER"), viper.GetBool("CRIB_IGNORE_NAMESPACE_PREFIX")); err != nil {
			logger.Error("invalid namespace for CRIB", slog.Any("error", err))
			os.Exit(1)
		}

		configFlags := genericclioptions.NewConfigFlags(true)
		kubeconfig := viper.GetString("KUBECONFIG")
		configFlags.KubeConfig = &kubeconfig
		k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
		if err != nil {
			logger.Error("failed to initialize kube client", slog.Any("error", err))
			os.Exit(1)
		}

		if err := k8sClient.CheckAccess(context.TODO()); err != nil {
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

		var rolebindingClient dynamic.ResourceInterface
		if viper.GetString("PROVIDER") == "aws" {
			dynamicClient, err := dynamic.NewForConfig(k8sClient.RestConfig())
			if err != nil {
				logger.Error("failed to initialize kube dynamic client", slog.Any("error", err))
				os.Exit(1)
			}

			rolebindingClient = dynamicClient.Resource(schema.GroupVersionResource{
				Group:    "rbac.authorization.k8s.io",
				Version:  "v1",
				Resource: "rolebindings",
			}).Namespace(viper.GetString("DEVSPACE_NAMESPACE"))
		}

		if err := utils.EnsureCribNamespaceReady(context.TODO(), k8sClient, rolebindingClient, viper.GetString("DEVSPACE_NAMESPACE"), viper.GetString("PROVIDER"), nil, nil); err != nil {
			logger.Error("failed to ensure crib namespace ready", slog.Any("error", err))
			os.Exit(1)
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

		if viper.GetString("DEVSPACE_PROFILE") == "keystone" {
			requiredEnvVars = append(requiredEnvVars, "KEYSTONE_ETH_WS_URL")
		}

		missingEnvVars := []string{}
		for _, name := range requiredEnvVars {
			value := os.Getenv(name)
			logger.Debug("reading env var", slog.String("name", name), slog.String("value", value))
			if value == "" {
				missingEnvVars = append(missingEnvVars, name)
			}
		}

		if viper.GetString("PROVIDER") == "aws" {
			for _, name := range []string{"DEVSPACE_INGRESS_CERT_ARN", "CHAINLINK_PRODUCT", "CHAINLINK_TEAM"} {
				value := os.Getenv(name)
				logger.Debug("reading env var", slog.String("name", name), slog.String("value", value))
				if value == "" {
					missingEnvVars = append(missingEnvVars, name)
				}
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

// configureCertProvisioningCmd represents the devspace configure-cert-provisioning subcommand
var configureCertProvisioningCmd = &cobra.Command{
	Use:   "configure-cert-provisioning",
	Short: "Provision the certificate for usage with cert-manager in k8s",
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetString("PROVIDER") != "kind" {
			logger.Error("this action can only be ran when provider is kind", slog.String("provider", viper.GetString("PROVIDER")))
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Running with the following parameters", "config", viper.AllSettings())

		configFlags := genericclioptions.NewConfigFlags(true)
		kubeconfig := viper.GetString("KUBECONFIG")
		configFlags.KubeConfig = &kubeconfig
		k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
		if err != nil {
			logger.Error("failed to initialize kube client", slog.Any("error", err))
			os.Exit(1)
		}

		caCert, err := wrappers.NewCACert(k8sClient, &wrappers.Mkcert{})
		if err != nil {
			logger.Error("failed to initialize CA cert", slog.Any("error", err))
			os.Exit(1)
		}

		//nolint: gosec
		secretName := "mkcert-ca-key-pair"
		secretNamespace := "cert-manager"
		alreadyExistingSecret, err := caCert.EnsureCertManagerSecret(context.TODO(), secretName, secretNamespace)
		if err != nil {
			logger.Error("failed to ensure cert-manager secret", slog.Any("error", err))
			os.Exit(1)
		}
		logger.Info("cert-manager secret in place",
			slog.String("name", secretName),
			slog.String("namespace", secretNamespace),
			slog.Bool("already_existing", alreadyExistingSecret),
		)

		clusterIssuerName := "mkcert-issuer"
		clusterIssuerNamespace := "cert-manager"
		if err := caCert.EnsureCAClusterIssuer(context.TODO(), secretName, clusterIssuerName, clusterIssuerNamespace); err != nil {
			logger.Error("failed to ensure cert-manager ClusterIssuer", slog.Any("error", err))
			os.Exit(1)
		}
		logger.Info("cert-manager ClusterIssuer in place",
			slog.String("name", clusterIssuerName),
			slog.String("namespace", clusterIssuerNamespace),
		)
	},
}

// purgeKindCmd represents the devspace purge-kind subcommand
var purgeKindCmd = &cobra.Command{
	Use:   "purge-kind",
	Short: "Delete the kind cluster and its associated resources",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetString("PROVIDER") != "kind" {
			logger.Error("this action can only be ran when provider is kind", slog.String("provider", viper.GetString("PROVIDER")))
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Running with the following parameters", "config", viper.AllSettings())

		dockerCli, err := wrappers.NewDockerCli()
		if err != nil {
			logger.Error("failed to initialize Docker CLI", slog.Any("error", err))
			os.Exit(1)
		}

		kindClusterName := wrappers.DefaultClusterName
		if len(args) > 0 {
			kindClusterName = args[0]
		}

		kindCluster := wrappers.NewKindCluster(kindClusterName, nil, dockerCli, nil, viper.GetString("KUBECONFIG"), wrappers.DefaultRegistryName, nil, nil)
		if err := kindCluster.Delete(); err != nil {
			logger.Error("failed to delete kind cluster", slog.Any("error", err))
			os.Exit(1)
		}
		logger.Info("kind cluster deleted", slog.String("name", kindClusterName))
	},
}

// labelNamespaceCmd represents the devspace label-namespace subcommand
var labelNamespaceCmd = &cobra.Command{
	Use:   "label-namespace [key] [value]",
	Short: "Add/update a label to the namespace",
	Long:  fmt.Sprintf("Add/update a label to the namespace. Values for known keys are validated.\nIf only the key is provided and it's %s, the value will default to '%s'. ", CleanupLabelKey, CleanupLabelValue),
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(2)),
	PreRun: func(cmd *cobra.Command, args []string) {
		if viper.GetString("PROVIDER") != "aws" {
			logger.Error("this action can only be ran when provider is aws", slog.String("provider", viper.GetString("PROVIDER")))
			os.Exit(1)
		}

		if viper.GetString("DEVSPACE_NAMESPACE") == "" {
			logger.Error("DEVSPACE_NAMESPACE must be set")
			os.Exit(1)
		}

		if len(args) == 2 {
			// attempts to validate values for known label keys
			switch args[0] {
			case CleanupLabelKey:
				// validation according to https://kyverno.io/docs/writing-policies/cleanup/#cleanup-label
				_, err := time.ParseDuration(args[1])
				if err != nil {
					_, err := time.Parse(time.RFC3339, args[1])
					if err != nil {
						logger.Error("invalid time format, expected ISO 8601 date or a duration e.g. 72h", slog.String("input", args[1]))
						os.Exit(1)
					}
				}
			}
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Running with the following parameters", "config", viper.AllSettings(), "args", args)

		if args[0] == CleanupLabelKey && len(args) == 1 {
			args = append(args, CleanupLabelValue)
		}

		configFlags := genericclioptions.NewConfigFlags(true)
		kubeconfig := viper.GetString("KUBECONFIG")
		configFlags.KubeConfig = &kubeconfig
		k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
		if err != nil {
			logger.Error("failed to initialize kube client", slog.Any("error", err))
			os.Exit(1)
		}

		// Generate labels
		labels := map[string]string{
			args[0]: args[1],
		}

		if err := k8sClient.LabelNamespace(context.TODO(), viper.GetString("DEVSPACE_NAMESPACE"), labels); err != nil {
			logger.Error("failed to label namespace", slog.Any("error", err))
			os.Exit(1)
		}
		logger.Info("namespace labeled", slog.String("namespace", viper.GetString("DEVSPACE_NAMESPACE")), slog.String("label_key", args[0]), slog.String("label_value", args[1]))
	},
}

// ingressCheckCmd represents the devspace ingress-check subcommand
var ingressCheckCmd = &cobra.Command{
	Use:   "ingress-check",
	Short: "Verify that all ingresses hostnames are resolvable up to a certain timeout and print them",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Running with the following parameters", "config", viper.AllSettings())

		configFlags := genericclioptions.NewConfigFlags(true)
		kubeconfig := viper.GetString("KUBECONFIG")
		configFlags.KubeConfig = &kubeconfig
		k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
		if err != nil {
			logger.Error("failed to initialize kube client", slog.Any("error", err))
			os.Exit(1)
		}

		ingressList, err := k8sClient.ListIngresses(context.TODO(), viper.GetString("DEVSPACE_NAMESPACE"))
		if err != nil {
			logger.Error("failed to list ingresses", slog.Any("error", err))
			os.Exit(1)
		}

		start := time.Now()
		i := 0
		ingressHosts := []string{}
		for i < len(ingressList.Items) {
			if time.Since(start) > viper.GetDuration("timeout") {
				logger.Error("timeout reached", slog.Any("timeout", viper.GetDuration("timeout")))
				os.Exit(1)
			}

			ingress := &ingressList.Items[i]
			logger.Info("validating ingress", slog.String("name", ingress.Name), slog.String("namespace", ingress.Namespace))
			logger.Debug("ingress object contents", slog.Any("status", ingress.Status))
			var ingressSuffix string
			switch *ingress.Spec.IngressClassName {
			case "alb":
				ingressSuffix = ".elb.amazonaws.com"
			case "nginx":
				ingressSuffix = "localhost"
			default:
				logger.Error("unsupported ingress class", slog.String("ingress_class", *ingress.Spec.IngressClassName), slog.Any("supported_classes", []string{"alb", "nginx"}))
				os.Exit(1)
			}

			for {
				if time.Since(start) > viper.GetDuration("timeout") {
					logger.Error("timeout reached", slog.Any("timeout", viper.GetDuration("timeout")))
					os.Exit(1)
				}

				// GetIngress makes sure we're refreshing the object's status every time
				ingress, err := k8sClient.GetIngress(context.TODO(), ingress.Namespace, ingress.Name)
				if err != nil {
					logger.Error("failed to refresh ingress status", slog.Any("ingress", ingress), slog.Any("error", err))
					os.Exit(1)
				}
				logger.Debug("ingress object contents", slog.Any("status", ingress.Status))

				if len(ingress.Status.LoadBalancer.Ingress) > 0 && strings.HasSuffix(ingress.Status.LoadBalancer.Ingress[0].Hostname, ingressSuffix) {
					logger.Info("ingress ready", slog.String("hostname", ingress.Status.LoadBalancer.Ingress[0].Hostname), slog.Duration("elapsed", time.Since(start)))
					for _, rule := range ingress.Spec.Rules {
						if rule.Host != "" {
							ingressHosts = append(ingressHosts, rule.Host)
						}
					}
					i++
					break
				} else {
					logger.Info("waiting for ingress creation", slog.String("name", ingress.Name), slog.Duration("internal", viper.GetDuration("interval")))
					time.Sleep(viper.GetDuration("interval"))
				}
			}
		}

		if len(ingressHosts) == 0 {
			logger.Error("no ingress hostnames found in namespace", slog.String("namespace", viper.GetString("DEVSPACE_NAMESPACE")))
			os.Exit(1)
		}

		if viper.GetString("PROVIDER") == "kind" {
			ingressIP := "127.0.0.1"
			for _, host := range ingressHosts {
				entry := fmt.Sprintf("%s %s", ingressIP, host)
				already_exists, err := utils.AddToEtcHosts(entry, "")
				if err != nil {
					logger.Error("failed to add entry to /etc/hosts", slog.String("entry", entry), slog.Any("error", err))
					os.Exit(1)
				}
				logger.Info("entry in /etc/hosts in place", slog.String("entry", entry), slog.Bool("already_exists", already_exists))
			}
			logger.Info("hosts file configured successfully")
		}

		for _, host := range ingressHosts {
			if time.Since(start) > viper.GetDuration("timeout") {
				logger.Error("timeout reached", slog.Any("timeout", viper.GetDuration("timeout")))
				os.Exit(1)
			}

			elapsed, err := utils.CheckHostnameResolution(host, viper.GetDuration("nsTimeout"), viper.GetDuration("interval"), nil)
			if err != nil {
				logger.Error("failed to resolve hostname", slog.String("hostname", host), slog.Any("error", err))
				os.Exit(1)
			}
			logger.Info("hostname resolved", slog.String("hostname", host), slog.Duration("elapsed", elapsed))
		}

		logger.Info("all ingress hostnames resolved", slog.String("namespace", viper.GetString("DEVSPACE_NAMESPACE")))

		if err := utils.PrintIngressHosts(context.TODO(), k8sClient, viper.GetString("DEVSPACE_NAMESPACE"), viper.GetString("PROVIDER")); err != nil {
			logger.Error("failed to ingress hosts", slog.Any("error", err))
			os.Exit(1)
		}
	},
}

// printIngressHostsCmd represents the devspace print-ingress-hosts subcommand
var printIngressHostsCmd = &cobra.Command{
	Use:   "print-ingress-hosts",
	Short: "Print all ingress hostnames in the namespace",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug("Running with the following parameters", "config", viper.AllSettings())

		configFlags := genericclioptions.NewConfigFlags(true)
		kubeconfig := viper.GetString("KUBECONFIG")
		configFlags.KubeConfig = &kubeconfig
		k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
		if err != nil {
			logger.Error("failed to initialize kube client", slog.Any("error", err))
			os.Exit(1)
		}

		if err := utils.PrintIngressHosts(context.TODO(), k8sClient, viper.GetString("DEVSPACE_NAMESPACE"), viper.GetString("PROVIDER")); err != nil {
			logger.Error("failed to ingress hosts", slog.Any("error", err))
			os.Exit(1)
		}
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
	devspaceCmd.AddCommand(configureCertProvisioningCmd)
	devspaceCmd.AddCommand(purgeKindCmd)

	devspaceCmd.AddCommand(labelNamespaceCmd)

	ingressCheckCmd.Flags().Duration("timeout", 2*time.Minute, "Timeout for the entire ingress check")
	ingressCheckCmd.Flags().Duration("nsTimeout", 1*time.Minute, "Timeout for each DNS lookup attempt")
	ingressCheckCmd.Flags().Duration("interval", 10*time.Second, "Time between retries")
	devspaceCmd.AddCommand(ingressCheckCmd)

	_ = viper.BindPFlags(ingressCheckCmd.Flags())

	devspaceCmd.AddCommand(printIngressHostsCmd)
}
