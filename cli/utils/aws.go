package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

func SetupAwsProfile(configPath string, profileName string, accountId string, region string, ssoRoleName string, ssoStartUrl string) error {
	// setting up go-ini configs so it matches the ~/.aws/config format closer
	// see: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-format
	ini.PrettyFormat = false
	ini.PrettyEqual = true

	cfg, err := ini.Load(configPath)
	if err != nil {
		return err
	}

	// write values under the respective profile section
	awsProfileIniSectionName := fmt.Sprintf("profile %s", profileName)
	if cfg.HasSection(awsProfileIniSectionName) {
		cfg.DeleteSection(awsProfileIniSectionName)
	}
	cfg.Section(awsProfileIniSectionName).Key("region").SetValue(region)
	cfg.Section(awsProfileIniSectionName).Key("sso_start_url").SetValue(ssoStartUrl)
	cfg.Section(awsProfileIniSectionName).Key("sso_region").SetValue(region)
	cfg.Section(awsProfileIniSectionName).Key("sso_account_id").SetValue(accountId)
	cfg.Section(awsProfileIniSectionName).Key("sso_role_name").SetValue(ssoRoleName)
	return cfg.SaveTo(configPath)
}
