package utils

import (
	"os"
	"strings"

	"github.com/smartcontractkit/crib/cli/wrappers"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/registry"
)

// InitializeHelmRegistryClient performs a lightweight version of helm's CLI initialization and returns a wrapper for interacting with helm registries
func InitializeHelmRegistryClient(envSettings *cli.EnvSettings) (wrappers.HelmRegistryAPI, error) {
	localHelmEnvSettings := envSettings
	if localHelmEnvSettings == nil {
		localHelmEnvSettings = cli.New()
	}

	opts := []registry.ClientOption{
		registry.ClientOptDebug(localHelmEnvSettings.Debug),
		registry.ClientOptEnableCache(true),
		registry.ClientOptWriter(os.Stderr),
		registry.ClientOptCredentialsFile(localHelmEnvSettings.RegistryConfig),
	}

	return wrappers.NewHelmRegistryClientWrapper(opts...)
}

// HelmRegistryLogin performs a registry login using the provided info
func HelmRegistryLogin(helmRegistryClient wrappers.HelmRegistryAPI, username string, password string, registryURL string) error {
	return helmRegistryClient.Login(strings.TrimPrefix(registryURL, "https://"), registry.LoginOptBasicAuth(username, password))
}
