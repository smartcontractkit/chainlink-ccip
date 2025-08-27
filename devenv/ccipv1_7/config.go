package ccipv1_7

/*
This file provides a simple boilerplate for TOML configuration with overrides
It has 3 functions: Load[T], Store[T] and LoadCache[T]

To configure the environment we use a set of files we read from the env var CTF_CONFIGS=env.toml,overrides.toml (can be more than 2) in Load[T]
To store infra or product component outputs we use Store[T] that creates env-cache.toml file.
This file can be used in tests or in any other code that integrated with dev environment.
LoadCache[T] is used if you need to write outputs the second time.
*/

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	// DefaultConfigDir is the default directory we are expecting TOML config to be.
	DefaultConfigDir = "."
	// EnvVarTestConfigs is the environment variable name to read config paths from, ex.: CTF_CONFIGS=env.toml,overrides.toml.
	EnvVarTestConfigs = "CTF_CONFIGS"
)

var L = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.InfoLevel)

// Load loads TOML configurations from environment variable, ex.: CTF_CONFIGS=env.toml,overrides.toml
// and unmarshalls the files from left to right overriding keys.
func Load[T any]() (*T, error) {
	var config T
	paths := strings.Split(os.Getenv(EnvVarTestConfigs), ",")
	for _, path := range paths {
		L.Info().Str("Path", path).Msg("Loading configuration input")
		data, err := os.ReadFile(filepath.Join(DefaultConfigDir, path)) //nolint:gosec
		if err != nil {
			return nil, fmt.Errorf("error reading config file %s: %w", path, err)
		}
		if L.GetLevel() == zerolog.TraceLevel {
			fmt.Println(string(data)) //nolint:forbidigo
		}

		decoder := toml.NewDecoder(strings.NewReader(string(data)))
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&config); err != nil {
			var details *toml.StrictMissingError
			if errors.As(err, &details) {
				fmt.Println(details.String()) //nolint:forbidigo
			}
			return nil, fmt.Errorf("failed to decode TOML config, strict mode: %s", err)
		}
	}
	if L.GetLevel() == zerolog.TraceLevel {
		L.Trace().Msg("Merged inputs")
		spew.Dump(config) //nolint:forbidigo
	}
	return &config, nil
}

// Store writes config to a file, adds -cache.toml suffix if it's an initial configuration.
func Store[T any](cfg *T) error {
	baseConfigPath, err := BaseConfigPath()
	if err != nil {
		return err
	}
	newCacheName := strings.ReplaceAll(baseConfigPath, ".toml", "")
	var outCacheName string
	if strings.Contains(newCacheName, "cache") {
		L.Info().Str("Cache", baseConfigPath).Msg("Cache file already exists, overriding")
		outCacheName = baseConfigPath
	} else {
		outCacheName = fmt.Sprintf("%s-out.toml", strings.ReplaceAll(baseConfigPath, ".toml", ""))
	}
	L.Info().Str("OutputFile", outCacheName).Msg("Storing configuration output")
	d, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(DefaultConfigDir, outCacheName), d, 0o600)
}

// LoadOutput loads config output file from path.
func LoadOutput[T any](path string) (*T, error) {
	_ = os.Setenv(EnvVarTestConfigs, path)
	return Load[T]()
}

// BaseConfigPath returns base config path, ex. env.toml,overrides.toml -> env.toml.
func BaseConfigPath() (string, error) {
	configs := os.Getenv(EnvVarTestConfigs)
	if configs == "" {
		return "", fmt.Errorf("no %s env var is provided, you should provide at least one test config in TOML", EnvVarTestConfigs)
	}
	L.Debug().Str("Configs", configs).Msg("Getting base config path")
	return strings.Split(configs, ",")[0], nil
}

func getNetworkPrivateKey() string {
	pk := os.Getenv("PRIVATE_KEY")
	if pk == "" {
		// that's the first Anvil and Geth private key, serves as a fallback for local testing if not overridden
		return "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	}
	return pk
}
