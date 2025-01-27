package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/smartcontractkit/crib/cli/wrappers"
)

type StartCmdConfig struct {
	RetryAttempts            int
	RetryDelay               time.Duration
	LokiEndpoint             string
	SkipSetup                bool
	Clean                    bool
	TestDelay                time.Duration
	ChainlinkCodeDir         string
	GoreleaserKey            string
	ChainlinkHelmRegistryURI string
}

type DevspaceEnvConfig struct {
	ScriptsDir               string
	ComponentsDir            string
	ChainlinkCodeDir         string
	GoreleaserKey            string
	Provider                 string
	DevspaceNamespace        string
	DevspaceImage            string
	DevspaceIngressDomain    string
	KeystoneAccountKey       string
	ChainlinkHelmRegistryURI string
}

func NewDevspaceEnvConfig(config StartCmdConfig) DevspaceEnvConfig {
	return DevspaceEnvConfig{
		ScriptsDir:               "../../scripts",
		ComponentsDir:            "../../components",
		ChainlinkCodeDir:         config.ChainlinkCodeDir,
		GoreleaserKey:            config.GoreleaserKey,
		Provider:                 "kind",
		DevspaceNamespace:        "crib-local",
		DevspaceImage:            "localhost:5001/chainlink-devspace",
		DevspaceIngressDomain:    "localhost",
		KeystoneAccountKey:       "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		ChainlinkHelmRegistryURI: config.ChainlinkHelmRegistryURI,
	}
}

func SetupKeystoneKindCrib(ctx context.Context, config StartCmdConfig) error {
	err, gitRoot := gitRoot()
	if err != nil {
		return err
	}

	if os.Getenv("IN_NIX_SHELL") == "" {
		return fmt.Errorf("this script must be run within a Nix shell")
	}

	env := getStartCmdEnvVars(config)

	dockerCli, err := wrappers.NewDockerCli()
	if err != nil {
		return fmt.Errorf("failed to initialize Docker CLI: %w", err)
	}
	// NOTE: because this subcommand doesn't load the config through viper, it relies on the bare environment variables + defaults
	kindCluster := wrappers.NewKindCluster("", nil, dockerCli, nil, "", "", nil, nil)

	if config.Clean {
		slog.Info("purging kind cluster")
		// NOTE: This uses a different set of envionment variables compared to the PurgeKindCluster function
		// This is because the PurgeKindCluster function is meant to be used as a standalone function
		// and shouldn't depend on the environment variables set by the StartCmdConfig
		if err := kindCluster.Delete(); err != nil {
			return fmt.Errorf("failed to purge kind cluster: %w", err)
		}
	}

	if err := kindCluster.CreateOrReuse("", nil); err != nil {
		return fmt.Errorf("failed to setup kind cluster: %w", err)
	}

	deployPath := filepath.Join(gitRoot, "deployments", "chainlink")
	slog.Info("setting namespace to crib-local")
	err = executeCommand(ctx, deployPath, env, "devspace", "use", "namespace", "crib-local")
	if err != nil {
		return fmt.Errorf("failed to set devspace namespace: %w", err)
	}

	slog.Info("setting up keystone-kind crib environment")
	deploymentType := deploymentType(config)
	err = executeCommand(ctx, deployPath, env, "devspace", "run", deploymentType, "--var=DEVSPACE_ENV_FILE=idonotexist")
	if err != nil {
		return fmt.Errorf("failed to setup keystone: %w", err)
	}

	return nil
}

func PrintDevspaceChainlinkClusterDep(ctx context.Context, config StartCmdConfig) error {
	slog.Info("Printing configuration for devspace dependency crib-chainlink-cluster")
	err, gitRoot := gitRoot()
	if err != nil {
		return err
	}

	env := getStartCmdEnvVars(config)
	// pretty print the env vars
	for _, e := range env {
		slog.Info(e)
	}

	deployPath := filepath.Join(gitRoot, "deployments", "chainlink")
	deploymentType := deploymentType(config)
	err = executeCommand(
		ctx,
		deployPath,
		env,
		"devspace",
		"--var=DEVSPACE_ENV_FILE=idonotexist",
		"print",
		"-p="+deploymentType,
		"--dependency=crib-chainlink-cluster",
	)
	if err != nil {
		return fmt.Errorf("failed to print devspace configuration: %w", err)
	}

	return nil
}

func deploymentType(config StartCmdConfig) string {
	if config.ChainlinkHelmRegistryURI == "" {
		return "keystone-kind-git-charts"
	} else {
		return "keystone-kind"
	}
}

func gitRoot() (error, string) {
	slog.Info("determining git root directory")
	gitRootCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	gitRootOutput, err := gitRootCmd.Output()
	gitRoot := string(gitRootOutput[:len(gitRootOutput)-1])
	if err != nil {
		return fmt.Errorf("failed to get git root: %w", err), ""
	}

	return nil, gitRoot
}

func getBaseEnvVars() []string {
	env := []string{
		fmt.Sprintf("HOME=%s", os.Getenv("HOME")),
		fmt.Sprintf("PATH=%s", os.Getenv("PATH")),
		fmt.Sprintf("USER=%s", os.Getenv("USER")),
		fmt.Sprintf("CRIB_CI_ENV=%t", true),
		fmt.Sprintf("IS_CRIB=%t", true),
		fmt.Sprintf("CRIB_SKIP_DOCKER_ECR_LOGIN=%t", true),
	}

	// include any AWS_ prefixed env vars
	for _, e := range os.Environ() {
		if strings.HasPrefix(e, "AWS_") {
			env = append(env, e)
		}
	}

	// include any NIX contained env vars
	for _, e := range os.Environ() {
		if strings.Contains(e, "NIX") {
			env = append(env, e)
		}
	}

	// include any git related env vars, case insensitive
	for _, e := range os.Environ() {
		if strings.Contains(strings.ToLower(e), "git") {
			env = append(env, e)
		}
	}

	return env
}

// We want to make sure that we're not loading in any env vars other than
// what we're explicitly setting here.
func getStartCmdEnvVars(config StartCmdConfig) []string {
	envConfig := NewDevspaceEnvConfig(config)
	baseEnv := getBaseEnvVars()

	env := []string{
		fmt.Sprintf("DEVSPACE_INGRESS_BASE_DOMAIN=%s", envConfig.DevspaceIngressDomain),
		fmt.Sprintf("SCRIPTS_DIR=%s", envConfig.ScriptsDir),
		fmt.Sprintf("COMPONENTS_DIR=%s", envConfig.ComponentsDir),
		fmt.Sprintf("CHAINLINK_CODE_DIR=%s", envConfig.ChainlinkCodeDir),
		fmt.Sprintf("GORELEASER_KEY=%s", envConfig.GoreleaserKey),
		fmt.Sprintf("PROVIDER=%s", envConfig.Provider),
		fmt.Sprintf("DEVSPACE_NAMESPACE=%s", envConfig.DevspaceNamespace),
		fmt.Sprintf("DEVSPACE_IMAGE=%s", envConfig.DevspaceImage),
		fmt.Sprintf("KEYSTONE_ACCOUNT_KEY=%s", envConfig.KeystoneAccountKey),
		fmt.Sprintf("CHAINLINK_HELM_REGISTRY_URI=%s", envConfig.ChainlinkHelmRegistryURI),
	}

	return append(baseEnv, env...)
}

func queryLoki(ctx context.Context, config StartCmdConfig) error {
	query := `{instance="0-ks-wf-node2"} |= "write_geth-test" |= "WriteTarget" |= "Transaction submitted"`

	for attempt := 0; attempt < config.RetryAttempts; attempt++ {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			slog.Info("attempting loki query",
				slog.Int("attempt", attempt+1),
				slog.String("query", query))

			sanitizedLokiEndpoint, err := url.Parse(config.LokiEndpoint)
			if err != nil {
				return fmt.Errorf("invalid LokiEndpoint URL: %w", err)
			}
			if sanitizedLokiEndpoint.Scheme != "https" && sanitizedLokiEndpoint.Scheme != "http" {
				return fmt.Errorf("LokiEndpoint must use http(s) scheme")
			}

			// #nosec G204
			cmd := exec.CommandContext(ctx, "logcli",
				"query",
				"--tls-skip-verify",
				"--since", "1m",
				"--limit", "1",
				"--addr", sanitizedLokiEndpoint.String(),
				query,
				"--output", "jsonl",
				"--quiet")

			output, err := cmd.Output()
			if err != nil {
				slog.Warn("attempt failed",
					slog.Int("attempt", attempt+1),
					slog.Any("error", err))
				time.Sleep(config.RetryDelay)
				continue
			}

			if len(bytes.TrimSpace(output)) > 0 {
				slog.Info("found matching log line", slog.String("line", string(output)))
				return nil
			}

			slog.Warn("no results found", slog.Int("attempt", attempt+1))
			time.Sleep(config.RetryDelay)
		}
	}

	return fmt.Errorf("no results found after %d attempts", config.RetryAttempts)
}

func SendSlackNotification(webhookURL string, command string) error {
	slog.Info("preparing slack notification",
		slog.String("command", command))

	// Get environment variables
	githubServerURL := os.Getenv("GITHUB_SERVER_URL")
	githubRepo := os.Getenv("GITHUB_REPOSITORY")
	githubRunID := os.Getenv("GITHUB_RUN_ID")
	githubActor := os.Getenv("GITHUB_ACTOR")

	workflowRunURL := fmt.Sprintf("%s/%s/actions/runs/%s",
		githubServerURL, githubRepo, githubRunID)

	// Construct Slack payload
	payload := map[string]interface{}{
		"blocks": []map[string]interface{}{
			{
				"type": "section",
				"text": map[string]interface{}{
					"type": "mrkdwn",
					"text": fmt.Sprintf(":red_circle: *Failed to provision CRIB environment for %s*", command),
				},
			},
			{
				"type": "section",
				"text": map[string]interface{}{
					"type": "mrkdwn",
					"text": fmt.Sprintf("<%s|Workflow Run>\nGit Repo: %s\nGithub Actor: %s",
						workflowRunURL, githubRepo, githubActor),
				},
			},
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	slog.Info("sending slack notification")
	ctx := context.TODO()
	// #nosec G107
	req, err := http.NewRequestWithContext(ctx, "POST", webhookURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending webhook: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("non-OK response from Slack: %s", resp.Status)
	}

	slog.Info("slack notification sent successfully")
	return nil
}

func RunKeystoneKind(ctx context.Context, config StartCmdConfig) error {
	slog.Info("starting keystone kind smoke test execution")

	if config.SkipSetup {
		slog.Info("skipping setup")
	} else {
		slog.Info("setting up test environment")
		if err := SetupKeystoneKindCrib(ctx, config); err != nil {
			return fmt.Errorf("command execution failed: %w", err)
		}
		slog.Info("commands executed successfully")

		if config.TestDelay > 0 {
			slog.Info("waiting after setup before performing smoke test", slog.Duration("delay", config.TestDelay))
			time.Sleep(config.TestDelay)
		}

		slog.Info("starting loki query")
	}

	if err := queryLoki(ctx, config); err != nil {
		return fmt.Errorf("loki query failed: %w", err)
	}

	slog.Info("keystone kind smoke test completed successfully")
	return nil
}

func executeCommand(
	ctx context.Context,
	dir string,
	env []string,
	name string,
	args ...string,
) error {
	slog.Info("executing command", slog.String("name", name), slog.String("args", strings.Join(args, " ")))
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Dir = dir
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func PurgeKindCluster() error {
	dockerCli, err := wrappers.NewDockerCli()
	if err != nil {
		return fmt.Errorf("failed to initialize Docker CLI: %w", err)
	}

	// NOTE: because this subcommand doesn't load the config through viper, it relies on the bare environment variables + defaults
	kindCluster := wrappers.NewKindCluster("", nil, dockerCli, nil, "", "", nil, nil)
	if err := kindCluster.Delete(); err != nil {
		return fmt.Errorf("failed to purge kind cluster: %w", err)
	}
	slog.Info("kind cluster deleted")
	return nil
}
