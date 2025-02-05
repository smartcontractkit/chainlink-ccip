package config

import (
	"os"

	"github.com/spf13/viper"
)

func InitViper() {
	MustHaveEnv("DEVSPACE_NAMESPACE")
	MustHaveEnv("PROVIDER")
	MustHaveEnv("DON_BOOT_NODE_COUNT")
	MustHaveEnv("DON_NODE_COUNT")
	MustHaveEnv("DEVSPACE_INGRESS_BASE_DOMAIN")
	MustHaveEnv("TMP_DIR")
	OptionalEnv("CRIB_CI_ENV")
	OptionalEnv("CHAINS_COUNT")
}

func OptionalEnv(key string) {
	err := viper.BindEnv(key)
	if err != nil {
		logger.Fatal("unable to bind env", "err", err)
		os.Exit(1)
	}
}

func MustHaveEnv(key string) string {
	err := viper.BindEnv(key)
	if err != nil {
		logger.Fatal("unable to bind env", "err", err)
		os.Exit(1)
	}
	value := viper.GetString(key)
	if value == "" {
		logger.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
