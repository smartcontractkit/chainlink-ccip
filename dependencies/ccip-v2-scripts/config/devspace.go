package config

import (
	"github.com/spf13/viper"
)

type DevspaceEnv struct {
	Namespace         string
	Provider          string
	DonBootNodeCount  int
	DonNodeCount      int
	IngressBaseDomain string
	TmpDir            string
}

func NewDevspaceEnvFromEnv() DevspaceEnv {
	return DevspaceEnv{
		Namespace:         viper.GetString("DEVSPACE_NAMESPACE"),
		Provider:          viper.GetString("PROVIDER"),
		DonBootNodeCount:  viper.GetInt("DON_BOOT_NODE_COUNT"),
		DonNodeCount:      viper.GetInt("DON_NODE_COUNT"),
		IngressBaseDomain: viper.GetString("DEVSPACE_INGRESS_BASE_DOMAIN"),
		TmpDir:            viper.GetString("TMP_DIR"),
	}
}
