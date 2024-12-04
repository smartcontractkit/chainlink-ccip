package utils

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"k8s.io/client-go/dynamic"
)

type RegistryLoginAttempt struct {
	RegistryType string // "docker" or "helm"
	RegistryHost string // e.g. "123456789012.dkr.ecr.us-west-2.amazonaws.com"
	LoginErr     error
}
type RefreshRegistriesECRCredentialsOutput struct {
	RegistryLoginAttempts         *[]RegistryLoginAttempt
	ECRGetAuthorizationTokenError error
}

func GetGitTopLevelDir(dir string) (string, error) {
	if stat, err := os.Stat(dir); os.IsNotExist(err) || !stat.IsDir() {
		return "", git.ErrRepositoryNotExists
	}

	// normalize the path
	path, err := filepath.Abs(dir)
	if err != nil {
		return "", err
	}

	for {
		repo, err := git.PlainOpen(path)
		if err == nil {
			// Get the working directory of the repository
			worktree, err := repo.Worktree()
			if err != nil {
				return "", fmt.Errorf("failed to get worktree for repo %s: %v", path, err)
			}

			return worktree.Filesystem.Root(), nil
		}

		parentDir := filepath.Dir(path)
		if parentDir == path {
			return "", git.ErrRepositoryNotExists
		}

		// Move up one directory level
		path = parentDir
	}
}

func ListFiles(dirPath string) ([]string, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Readdirnames(-1)
}

func PromptForInput(key string, defaultValue string) (string, error) {
	userInput := ""
	prompt := fmt.Sprintf("Please enter a value for %s", key)
	if defaultValue != "" {
		prompt = fmt.Sprintf("%s (default is '%s')", prompt, defaultValue)
	}
	prompt = fmt.Sprintf("%s: ", prompt)
	fmt.Print(prompt)
	_, err := fmt.Scanln(&userInput)
	userInput = strings.Trim(userInput, " ")
	if userInput == "" {
		return defaultValue, nil
	}
	return userInput, err
}

// PresentPrompt presents a prompt to the user with possible choices and waits for valid input
func PresentPrompt(prompt string, choices []string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display the prompt and valid choices
		fmt.Printf("%s (%s): ", prompt, strings.Join(choices, "/"))

		// Read the user's input
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		// Clean up the input (remove newline and extra spaces)
		input = strings.TrimSpace(input)

		// Check if the input is one of the valid choices
		for _, choice := range choices {
			if input == choice {
				return input // Return valid input
			}
		}

		// Inform the user that the input is invalid
		fmt.Printf("Invalid choice. Please choose one of %v.\n", choices)
	}
}

func CopyFile(srcPath string, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}

	// Copy source to destination
	_, err = io.Copy(dst, src)
	return err
}

// WriteConfig writes the configuration settings to a file by searching and replacing the values of the keys in kv
func WriteConfig(path string, kv map[string]string) error {
	tempFile, err := os.CreateTemp(filepath.Dir(path), ".env.temp")
	if err != nil {
		return fmt.Errorf("error creating temporary file: %w", err)
	}
	defer tempFile.Close()

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening .env file: %w", err)
	}
	defer file.Close()

	linesToWrite := []string{}
	updatedLines := map[string]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for key, newValue := range kv {
			if strings.HasPrefix(line, key) {
				line = fmt.Sprintf("%s=%s", key, newValue)
				updatedLines[key] = newValue
				break
			}
		}
		linesToWrite = append(linesToWrite, line)
	}

	cliMarker := "# Added by CRIB CLI"
	if len(updatedLines) < len(kv) {
		if !slices.Contains(linesToWrite, cliMarker) {
			linesToWrite = append(linesToWrite, cliMarker)
		}
		for key, newValue := range kv {
			if _, ok := updatedLines[key]; !ok {
				linesToWrite = append(linesToWrite, fmt.Sprintf("%s=%s", key, newValue))
			}
		}
	}

	for _, line := range linesToWrite {
		if _, err := tempFile.WriteString(line + "\n"); err != nil {
			return fmt.Errorf("error writing to temporary file: %w", err)
		}
	}

	// Move the temporary file to replace the original .env file
	if err := os.Rename(tempFile.Name(), path); err != nil {
		return fmt.Errorf("error moving temporary file to original file: %w", err)
	}

	return nil
}

func ExtractHostFromUrl(input string) (string, error) {
	parsedUrl, err := url.Parse(input)
	if err != nil {
		return "", err
	}
	return parsedUrl.Host, nil
}

func RefreshRegistriesECRCredentials(ecrClient wrappers.ECRAPI, dockerCli wrappers.DockerCLI,
	helmRegistryClient wrappers.HelmRegistryAPI, chainlinkHelmRegistryUri string, chainlinkDockerRegistriesMap map[string]string,
) *RefreshRegistriesECRCredentialsOutput {
	registryLoginAttempts := []RegistryLoginAttempt{}
	output := &RefreshRegistriesECRCredentialsOutput{
		RegistryLoginAttempts:         &registryLoginAttempts,
		ECRGetAuthorizationTokenError: nil,
	}
	if dockerCli == nil && helmRegistryClient == nil {
		return output
	}

	ecrAuthToken, err := GetDecodedECRAuthorizationToken(ecrClient)
	if err != nil {
		output.ECRGetAuthorizationTokenError = err
		return output
	}

	if len(ecrAuthToken) == 0 {
		output.ECRGetAuthorizationTokenError = fmt.Errorf("no authorization data returned")
		return output
	}

	if len(ecrAuthToken) > 1 {
		slog.Warn("got multiple ECR auth tokens back from ecr.GetAuthorizationToken. Only the first one will be used")
	}

	if dockerCli != nil {
		for env, dockerRegistryHost := range chainlinkDockerRegistriesMap {
			dockerRegistryLoginAttempt := RegistryLoginAttempt{RegistryType: "docker"}
			slog.Info("login to docker registry", "env", env)
			if err == nil {
				dockerRegistryLoginAttempt.RegistryHost = dockerRegistryHost
				if _, loginErr := dockerCli.Login(ecrAuthToken[0].Username, ecrAuthToken[0].Password, dockerRegistryHost); loginErr != nil {
					dockerRegistryLoginAttempt.LoginErr = loginErr
				}
			} else {
				dockerRegistryLoginAttempt.LoginErr = err
			}
			registryLoginAttempts = append(registryLoginAttempts, dockerRegistryLoginAttempt)
		}
	}

	if helmRegistryClient != nil {
		helmRegistryLoginAttempt := RegistryLoginAttempt{RegistryType: "helm"}
		helmRegistryHost, err := ExtractHostFromUrl(chainlinkHelmRegistryUri)
		if err == nil {
			helmRegistryLoginAttempt.RegistryHost = helmRegistryHost
			if loginErr := HelmRegistryLogin(helmRegistryClient, ecrAuthToken[0].Username, ecrAuthToken[0].Password, helmRegistryHost); loginErr != nil {
				helmRegistryLoginAttempt.LoginErr = loginErr
			}
		} else {
			helmRegistryLoginAttempt.LoginErr = err
		}
		registryLoginAttempts = append(registryLoginAttempts, helmRegistryLoginAttempt)
	}

	return output
}

func IsValidCribNamespace(namespace string, provider string, skipPrefixCheck bool) error {
	if provider == "kind" && namespace != "crib-local" {
		return fmt.Errorf("DEVSPACE_NAMESPACE must be set to 'crib-local' when using kind provider")
	}

	if !skipPrefixCheck && !strings.HasPrefix(namespace, "crib-") {
		return fmt.Errorf("DEVSPACE_NAMESPACE must begin with 'crib-' prefix")
	}
	return nil
}

func GetChainlinkDockerRegistries(provider string) map[string]string {
	if provider == "kind" {
		// In kind, we need to login to all 3 registries, so pods can pull images via image pull secrets
		return map[string]string{
			"sdlc":  "795953128386.dkr.ecr.us-west-2.amazonaws.com",
			"stage": "323150190480.dkr.ecr.us-west-2.amazonaws.com",
			"prod":  "804282218731.dkr.ecr.us-west-2.amazonaws.com",
		}
	}
	// If we're deploying to remote EKS we just need to login to stage, so we can push dev images
	return map[string]string{
		"stage": "323150190480.dkr.ecr.us-west-2.amazonaws.com",
	}
}

func EnsureCribNamespaceReady(ctx context.Context, k8sClient wrappers.K8sCLI, rolebindingClient dynamic.ResourceInterface, namespace string, provider string, waitTimeout *time.Duration, sleepBetweenAttempts *time.Duration) error {
	defaultWaitTimeout := 20 * time.Second
	defaultSleepBetweenAttempts := 500 * time.Millisecond

	if waitTimeout == nil {
		waitTimeout = &defaultWaitTimeout
	}
	if sleepBetweenAttempts == nil {
		sleepBetweenAttempts = &defaultSleepBetweenAttempts
	}

	alreadyExists, err := k8sClient.EnsureNamespaceExists(ctx, namespace)
	if err != nil {
		return fmt.Errorf("failed to ensure namespace existence: %w", err)
	}
	slog.Debug("k8s namespace in place", slog.String("name", namespace), slog.Bool("already_exists", alreadyExists))

	if provider == "aws" {
		roleBindingName := fmt.Sprintf("%s-crib-poweruser", namespace)
		slog.Info("waiting for rolebinding creation", slog.String("role_binding_name", roleBindingName), slog.String("namespace", namespace))
		startTime := time.Now()
		if err := k8sClient.WaitForResource(ctx, rolebindingClient, roleBindingName, *sleepBetweenAttempts, *waitTimeout); err != nil {
			return fmt.Errorf("failed to wait for crib-power-user role binding to be created: %w", err)
		}
		slog.Info("role binding found",
			slog.String("role_binding_name", roleBindingName),
			slog.String("namespace", namespace),
			slog.Float64("elapsed_seconds", time.Since(startTime).Seconds()),
		)
	}
	return nil
}
