package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/smartcontractkit/crib/cli/wrappers"
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
