package utils

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"gopkg.in/ini.v1"
)

func SetupAwsProfile(configPath string, profileName string, accountId string, region string, ssoRoleName string, ssoStartUrl string) error {
	// setting up go-ini configs so it matches the ~/.aws/config format closer
	// see: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html#cli-configure-files-format
	ini.PrettyFormat = false
	ini.PrettyEqual = true

	configFile, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}

	cfg, err := ini.Load(configFile)
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

type GetDecodedECRAuthorizationTokenOutput struct {
	Username    string
	Password    string
	RegistryURL string
}

// GetDecodedECRAuthorizationToken performs ecr.GetAuthorizationToken, handle possible errors and return the
// base64-decoded token
func GetDecodedECRAuthorizationToken(ecrClient wrappers.ECRAPI) ([]*GetDecodedECRAuthorizationTokenOutput, error) {
	output := []*GetDecodedECRAuthorizationTokenOutput{}

	ecrAuthToken, err := ecrClient.GetAuthorizationToken(context.TODO(), &ecr.GetAuthorizationTokenInput{})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch ECR authorization token, %v", err)
	}
	if len(ecrAuthToken.AuthorizationData) == 0 {
		return nil, fmt.Errorf("no authorization data returned")
	}

	for i := 0; i < len(ecrAuthToken.AuthorizationData); i++ {
		// Decode each Base64-encoded authorization token
		authData := ecrAuthToken.AuthorizationData[i]
		authToken := *authData.AuthorizationToken
		decodedToken, err := base64.StdEncoding.DecodeString(authToken)
		if err != nil {
			return nil, err
		}

		// token format is username:password
		parts := strings.Split(string(decodedToken), ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("unexpected token format")
		}

		output = append(output,
			&GetDecodedECRAuthorizationTokenOutput{
				Username:    parts[0],
				Password:    parts[1],
				RegistryURL: *authData.ProxyEndpoint,
			})
	}

	return output, nil
}

// TODO: evaluate inspecting credentials expiry, so we can refresh the token beforehand
func HasValidAwsSession(stsClient wrappers.STSAPI) bool {
	_, err := stsClient.GetCallerIdentity(context.TODO(), &sts.GetCallerIdentityInput{})
	return err == nil
}

func AwsSsoLogin(awsConfigFile string, awsProfile string) error {
	cmd := exec.Command("aws", "sso", "login")
	cmd.Env = []string{
		"PATH=" + os.Getenv("PATH"),
		"AWS_CONFIG_FILE=" + awsConfigFile,
		"AWS_PROFILE=" + awsProfile,
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func EnsureValidAwsSession(stsClient wrappers.STSAPI, awsConfigFile string, awsProfile string, shouldTryAwsSso bool) error {
	if HasValidAwsSession(stsClient) {
		return nil
	}

	msg := "No valid AWS session found."
	if !shouldTryAwsSso {
		return errors.New(msg)
	}

	slog.Warn(fmt.Sprintf("%s Attempting to login via AWS SSO", msg))
	if err := AwsSsoLogin(awsConfigFile, awsProfile); err != nil {
		return fmt.Errorf("failed to aws sso login, %v", err)
	}

	return nil
}
