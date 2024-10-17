package utils

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	ecrtypes "github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/ini.v1"
)

// AwsProfile represent an AWS profile inside ~/.aws/config
type AwsProfile struct {
	name    string
	entries map[string]string
}

// MockAwsConfigFile is a helper function that returns the path to a tempfile
// containing the desired content and permissions
func MockAwsConfigFile(content []byte, perm fs.FileMode) *os.File {
	tempFile, err := os.CreateTemp("", "config")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}

	if err := os.WriteFile(tempFile.Name(), content, perm); err != nil {
		log.Fatal(err)
	}

	return tempFile
}

// LoadTestAwsConfig is a convenience wrapper for config.LoadDefaultConfig from AWS Go SDK
func LoadTestAwsConfig(awsConfigFile string, profileName string) (aws.Config, error) {
	os.Unsetenv("AWS_DEFAULT_REGION")
	return config.LoadDefaultConfig(
		context.TODO(),
		config.WithSharedConfigFiles([]string{awsConfigFile}),
		config.WithSharedConfigProfile(profileName))
}

// InMemIniWithAwsProfiles takes a list of AwsProfile and returns an in-memory *ini.File instance
func InMemIniWithAwsProfiles(awsProfiles []AwsProfile) (*ini.File, error) {
	inMemIni := ini.Empty()

	for _, awsProfile := range awsProfiles {
		section, err := inMemIni.NewSection(fmt.Sprintf("profile %s", awsProfile.name))
		if err != nil {
			return inMemIni, err
		}

		for k, v := range awsProfile.entries {
			_, _ = section.NewKey(k, v)
		}
	}

	return inMemIni, nil
}

// AssertEqualIniSection compares two ini.Section objects based on its
// keys and values only, as direct equality doesn't work well when
// working with temporary files
func AssertEqualIniSections(t *testing.T, a *ini.Section, b *ini.Section) {
	t.Helper()

	// strictly comparing ini.File, ini.Section or ini.Key instances
	// fails because each of them carry a pointer to its respective parent
	// which contains the dataSources attribute, which carries the path
	// for the file, which is incomparable when dealing with temp files
	// see: https://github.com/go-ini/ini/blob/b2f570e5b5b844226bbefe6fb521d891f529a951/file.go#L31
	for _, k := range a.Keys() {
		assert.Equal(t, k.Value(), b.Key(k.Name()).Value())
	}
	assert.Equal(t, len(a.Keys()), len(b.Keys()))
}

func TestSetupAwsProfileNonExisting(t *testing.T) {
	t.Parallel()

	mockedAwsConfig := MockAwsConfigFile([]byte(""), 0o666)
	defer os.Remove(mockedAwsConfig.Name())

	require.NoError(t, SetupAwsProfile(mockedAwsConfig.Name(), "test-profile", "12345678909", "ap-southeast-1", "test-role-name", "https://sso.start.url"))

	// inspect the new config with the AWS SDK
	cfg, err := LoadTestAwsConfig(mockedAwsConfig.Name(), "test-profile")
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, "ap-southeast-1", cfg.Region)

	// aws.Config doesn't directly expose sso_ prefixed profile parameters
	// so our only choice is to inspect the file and assert the content is what we expect
	want, err := InMemIniWithAwsProfiles([]AwsProfile{
		{
			name: "test-profile",
			entries: map[string]string{
				"region":         "ap-southeast-1",
				"sso_start_url":  "https://sso.start.url",
				"sso_region":     "ap-southeast-1",
				"sso_account_id": "12345678909",
				"sso_role_name":  "test-role-name",
			},
		},
	})
	require.NoError(t, err)
	require.NotNil(t, want)

	got, err := ini.Load(mockedAwsConfig.Name())
	require.NoError(t, err)
	require.NotNil(t, got)

	wantSectionName := "profile test-profile"
	AssertEqualIniSections(t, want.Section(wantSectionName), got.Section(wantSectionName))
}

func TestSetupAwsProfileExistsButDiverges(t *testing.T) {
	t.Parallel()

	mockedAwsConfig := MockAwsConfigFile([]byte(`[default]
region = us-west-2
output = text

[profile some-other-profile]
region = ap-southeast-1
output = text

[profile test-profile]
region = us-east-1
output = json
sso_start_url = https://some.outdated.url
sso_region = us-east-1
sso_account_id = 123
sso_role_name = outdated-role`), 0o666)
	defer os.Remove(mockedAwsConfig.Name())

	require.NoError(t, SetupAwsProfile(mockedAwsConfig.Name(), "test-profile", "12345678909", "ap-southeast-1", "test-role-name", "https://sso.start.url"))

	// inspect the new config with the AWS SDK
	cfg, err := LoadTestAwsConfig(mockedAwsConfig.Name(), "test-profile")
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, "ap-southeast-1", cfg.Region)

	// aws.Config doesn't directly expose sso_ prefixed profile parameters
	// so our only choice is to inspect the file and assert the content is what we expect
	wantAwsProfiles := []AwsProfile{
		{
			name: "some-other-profile",
			entries: map[string]string{
				"region": "ap-southeast-1",
				"output": "text",
			},
		},
		{
			name: "test-profile",
			entries: map[string]string{
				"region":         "ap-southeast-1",
				"sso_start_url":  "https://sso.start.url",
				"sso_region":     "ap-southeast-1",
				"sso_account_id": "12345678909",
				"sso_role_name":  "test-role-name",
			},
		},
	}

	want, err := InMemIniWithAwsProfiles(wantAwsProfiles)
	require.NoError(t, err)
	require.NotNil(t, want)

	got, err := ini.Load(mockedAwsConfig.Name())
	require.NoError(t, err)
	require.NotNil(t, got)

	for _, wantAwsProfile := range wantAwsProfiles {
		wantSectionName := fmt.Sprintf("profile %s", wantAwsProfile.name)
		AssertEqualIniSections(t, want.Section(wantSectionName), got.Section(wantSectionName))
	}

	// test if the default section is also untouched
	want = ini.Empty()
	section, err := want.NewSection("default")
	require.NoError(t, err)
	_, _ = section.NewKey("region", "us-west-2")
	_, _ = section.NewKey("output", "text")

	AssertEqualIniSections(t, want.Section("default"), got.Section("default"))
}

func TestGetDecodedECRAuthorizationTokenSingle(t *testing.T) {
	t.Parallel()

	authToken := base64.StdEncoding.EncodeToString([]byte("user:password"))
	proxyEndpoint := "https://012345678910.dkr.ecr.us-east-1.amazonaws.com"
	mockEcrClient := wrappermocks.NewECRAPI(t)
	mockEcrClient.EXPECT().
		GetAuthorizationToken(
			context.TODO(), &ecr.GetAuthorizationTokenInput{},
		).Return(
		&ecr.GetAuthorizationTokenOutput{
			AuthorizationData: []ecrtypes.AuthorizationData{
				{AuthorizationToken: &authToken, ProxyEndpoint: &proxyEndpoint},
			},
		}, nil,
	)

	want := []*GetDecodedECRAuthorizationTokenOutput{
		{Username: "user", Password: "password", RegistryURL: proxyEndpoint},
	}
	got, err := GetDecodedECRAuthorizationToken(mockEcrClient)
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestGetDecodedECRAuthorizationTokenMultiple(t *testing.T) {
	t.Parallel()

	authToken1 := base64.StdEncoding.EncodeToString([]byte("user:password"))
	proxyEndpoint1 := "https://012345678910.dkr.ecr.us-east-1.amazonaws.com"
	authToken2 := base64.StdEncoding.EncodeToString([]byte("anotheruser:anotherpassword"))
	proxyEndpoint2 := "https://012345678910.dkr.ecr.us-west-2.amazonaws.com"
	mockEcrClient := wrappermocks.NewECRAPI(t)
	mockEcrClient.EXPECT().
		GetAuthorizationToken(
			context.TODO(), &ecr.GetAuthorizationTokenInput{},
		).Return(
		&ecr.GetAuthorizationTokenOutput{
			AuthorizationData: []ecrtypes.AuthorizationData{
				{AuthorizationToken: &authToken1, ProxyEndpoint: &proxyEndpoint1},
				{AuthorizationToken: &authToken2, ProxyEndpoint: &proxyEndpoint2},
			},
		}, nil,
	)

	want := []*GetDecodedECRAuthorizationTokenOutput{
		{Username: "user", Password: "password", RegistryURL: proxyEndpoint1},
		{Username: "anotheruser", Password: "anotherpassword", RegistryURL: proxyEndpoint2},
	}
	got, err := GetDecodedECRAuthorizationToken(mockEcrClient)
	require.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestGetDecodedECRAuthorizationTokenErrors(t *testing.T) {
	t.Parallel()

	type errorTestCases struct {
		description   string
		mockEcrClient wrappers.ECRAPI
		wantError     string
	}

	mockEcrClientLackOfIAMPermissions := wrappermocks.NewECRAPI(t)
	mockEcrClientLackOfIAMPermissions.EXPECT().
		GetAuthorizationToken(
			context.TODO(), &ecr.GetAuthorizationTokenInput{},
		).Return(nil, fmt.Errorf("unable to fetch ECR authorization token, %v", errors.New("some problem in AWS")))

	mockEcrClientNoAuthData := wrappermocks.NewECRAPI(t)
	mockEcrClientNoAuthData.EXPECT().
		GetAuthorizationToken(
			context.TODO(), &ecr.GetAuthorizationTokenInput{},
		).Return(&ecr.GetAuthorizationTokenOutput{AuthorizationData: []ecrtypes.AuthorizationData{}}, nil)

	mockEcrClientInvalidTokenFormat := wrappermocks.NewECRAPI(t)
	faultyToken := base64.StdEncoding.EncodeToString([]byte("user:password:"))
	proxyEndpoint := "https://some.ecr.url"
	mockEcrClientInvalidTokenFormat.EXPECT().
		GetAuthorizationToken(
			context.TODO(), &ecr.GetAuthorizationTokenInput{},
		).Return(
		&ecr.GetAuthorizationTokenOutput{AuthorizationData: []ecrtypes.AuthorizationData{
			{AuthorizationToken: &faultyToken, ProxyEndpoint: &proxyEndpoint},
		}}, nil)

	for _, scenario := range []errorTestCases{
		{
			description:   "lack of permission",
			mockEcrClient: mockEcrClientLackOfIAMPermissions,
			wantError:     "unable to fetch ECR authorization token",
		},
		{
			description:   "empty response from GetAuthorizationToken",
			mockEcrClient: mockEcrClientNoAuthData,
			wantError:     "no authorization data returned",
		},
		{
			description:   "invalid token format",
			mockEcrClient: mockEcrClientInvalidTokenFormat,
			wantError:     "unexpected token format",
		},
	} {
		t.Run(scenario.description, func(t *testing.T) {
			t.Parallel()

			got, err := GetDecodedECRAuthorizationToken(scenario.mockEcrClient)
			assert.ErrorContains(t, err, scenario.wantError)
			assert.Nil(t, got)
		})
	}
}

func TestHasValidAwsSession(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description   string
		mockStsClient wrappers.STSAPI
		expectedValid bool
	}

	mockStsClientValidSession := wrappermocks.NewSTSAPI(t)
	mockStsClientValidSession.EXPECT().
		GetCallerIdentity(
			context.TODO(), &sts.GetCallerIdentityInput{},
		).Return(&sts.GetCallerIdentityOutput{}, nil)

	mockStsClientInvalidSession := wrappermocks.NewSTSAPI(t)
	mockStsClientInvalidSession.EXPECT().
		GetCallerIdentity(
			context.TODO(), &sts.GetCallerIdentityInput{},
		).Return(nil, fmt.Errorf("some error"))

	testCases := []testCase{
		{
			description:   "valid session",
			mockStsClient: mockStsClientValidSession,
			expectedValid: true,
		},
		{
			description:   "invalid session",
			mockStsClient: mockStsClientInvalidSession,
			expectedValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			valid := HasValidAwsSession(tc.mockStsClient)
			assert.Equal(t, tc.expectedValid, valid)
		})
	}
}

func TestEnsureValidAwsSession(t *testing.T) {
	t.Parallel()

	type testCase struct {
		description     string
		mockStsClient   wrappers.STSAPI
		awsConfigFile   string
		awsProfile      string
		shouldTryAwsSso bool
		wantError       string
	}

	mockStsClientValidSession := wrappermocks.NewSTSAPI(t)
	mockStsClientValidSession.EXPECT().
		GetCallerIdentity(
			context.TODO(), &sts.GetCallerIdentityInput{},
		).Return(&sts.GetCallerIdentityOutput{}, nil)

	mockStsClientInvalidSession := wrappermocks.NewSTSAPI(t)
	mockStsClientInvalidSession.EXPECT().
		GetCallerIdentity(
			context.TODO(), &sts.GetCallerIdentityInput{},
		).Return(nil, fmt.Errorf("some error"))

	testCases := []testCase{
		{
			description:     "valid session",
			mockStsClient:   mockStsClientValidSession,
			awsConfigFile:   "mockConfigFile",
			awsProfile:      "mockProfile",
			shouldTryAwsSso: false,
			wantError:       "",
		},
		{
			description:     "invalid session, should not try SSO",
			mockStsClient:   mockStsClientInvalidSession,
			awsConfigFile:   "mockConfigFile",
			awsProfile:      "mockProfile",
			shouldTryAwsSso: false,
			wantError:       "No valid AWS session found.",
		},
		// TODO: can't easily test SSO login here, that will require refactoring the EnsureValidAwsSession to receive an interface that forces the AwsSsoLogin function
		// or, we can take care of it in future integration tests
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Parallel()

			err := EnsureValidAwsSession(tc.mockStsClient, tc.awsConfigFile, tc.awsProfile, tc.shouldTryAwsSso)
			if tc.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tc.wantError)
			}
		})
	}
}
