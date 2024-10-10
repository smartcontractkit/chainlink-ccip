package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// MockKubeConfigFile is a helper function that returns the path to a tempfile
// containing the desired content and permissions
func MockKubeConfigFile(content []byte, perm fs.FileMode) *os.File {
	tempFile, err := os.CreateTemp("", "config")
	if err != nil {
		log.Fatalf("Failed to create temp file: %v", err)
	}

	if err := os.WriteFile(tempFile.Name(), content, perm); err != nil {
		log.Fatal(err)
	}

	return tempFile
}

// AssertEqualKubeConfigs compares the fields we modify from two clientcmdapi.Config instances
// TODO: use golang reflect to compare the entire object and recurse into subtypes whilst ignoring
// fields such as LocationOfOrigin (which varies when dealing with tmpfiles)
func AssertEqualKubeConfigs(t *testing.T, want *clientcmdapi.Config, got *clientcmdapi.Config) {
	t.Helper()

	assert.Len(t, got.Clusters, len(want.Clusters))
	for name, wantCluster := range want.Clusters {
		gotCluster := got.Clusters[name]
		assert.Equal(t, wantCluster.Server, gotCluster.Server)
		assert.Equal(t, wantCluster.CertificateAuthorityData, gotCluster.CertificateAuthorityData)
	}

	assert.Len(t, got.Contexts, len(want.Contexts))
	for name, wantContext := range want.Contexts {
		gotContext := got.Contexts[name]
		assert.Equal(t, wantContext.Cluster, gotContext.Cluster)
		assert.Equal(t, wantContext.AuthInfo, gotContext.AuthInfo)
	}

	assert.Len(t, got.AuthInfos, len(want.AuthInfos))
	for name, wantAuthInfo := range want.AuthInfos {
		gotAuthInfo := got.AuthInfos[name]
		if gotAuthInfo.Exec != nil {
			assert.Equal(t, wantAuthInfo.Exec, gotAuthInfo.Exec)
		}
		if gotAuthInfo.ClientKey != "" {
			assert.Equal(t, wantAuthInfo.ClientKey, gotAuthInfo.ClientKey)
		}
		if gotAuthInfo.ClientCertificate != "" {
			assert.Equal(t, wantAuthInfo.ClientCertificate, gotAuthInfo.ClientCertificate)
		}
	}

	assert.Equal(t, want.CurrentContext, got.CurrentContext)
}

func TestSetupKubeConfigNonExisting(t *testing.T) {
	t.Parallel()

	nonExistingKubeConfig := filepath.Join(t.TempDir(), "non-existing")

	// mocking return of eks.DescribeCluster
	eksClusterName := "test-eks-cluster"
	eksClusterAlias := "test-eks-cluster-alias"
	eksClusterArn := "arn:aws:eks:ap-southeast-1:123456789000:cluster/test-eks-cluster"
	eksClusterEndpoint := "https://cluster.endpoint"
	eksEncodedCAData := base64.StdEncoding.EncodeToString([]byte("cadata"))
	mockEksClient := wrappermocks.NewEKSAPI(t)
	mockEksClient.EXPECT().
		DescribeCluster(
			context.TODO(), &eks.DescribeClusterInput{Name: &eksClusterName},
		).Return(
		&eks.DescribeClusterOutput{
			Cluster: &ekstypes.Cluster{
				Arn:                  &eksClusterArn,
				Endpoint:             &eksClusterEndpoint,
				CertificateAuthority: &ekstypes.Certificate{Data: &eksEncodedCAData},
			},
		}, nil,
	)

	setupKubeConfigInput := &SetupKubeConfigInput{
		EksClient:            mockEksClient,
		KubeconfigPath:       nonExistingKubeConfig,
		EksClusterName:       eksClusterName,
		EksAliasName:         eksClusterAlias,
		CribNamespace:        "crib-test",
		AwsProfile:           "profile-test",
		AwsRegion:            "ap-southeast-1",
		ChangeDefaultContext: true,
	}
	require.NoError(t, SetupKubeConfig(setupKubeConfigInput))
	require.FileExists(t, nonExistingKubeConfig)

	got, err := clientcmd.LoadFromFile(nonExistingKubeConfig)
	require.NoError(t, err)

	want := &clientcmdapi.Config{
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			eksClusterArn: {
				Exec: &clientcmdapi.ExecConfig{
					Command: "aws",
					Args:    []string{"--region", "ap-southeast-1", "eks", "get-token", "--cluster-name", eksClusterName, "--output", "json"},
					Env: []clientcmdapi.ExecEnvVar{
						{Name: "AWS_PROFILE", Value: "profile-test"},
					},
					APIVersion:      "client.authentication.k8s.io/v1beta1",
					InteractiveMode: "IfAvailable",
				},
			},
		},
		Clusters: map[string]*clientcmdapi.Cluster{
			eksClusterArn: {
				Server:                   eksClusterEndpoint,
				CertificateAuthorityData: []byte("cadata"),
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			eksClusterAlias: {
				Cluster:  eksClusterArn,
				AuthInfo: eksClusterArn,
			},
		},
		CurrentContext: eksClusterAlias,
	}
	AssertEqualKubeConfigs(t, want, got)
}

func TestSetupKubeConfigExistsButDiverges(t *testing.T) {
	t.Parallel()

	mockedKubeConfig := MockKubeConfigFile([]byte(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: c29tZXRoaW5n  # base64-encoded string "something"
    server: https://unrelated.endpoint
  name: arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch
- cluster:
    certificate-authority-data: c29tZXRoaW5nd3JvbmcK  # base64-encoded string "somethingwrong"
    server: https://wrong.endpoint
  name: arn:aws:eks:ap-southeast-1:123456789000:cluster/test-eks-cluster
contexts:
- context:
    cluster: arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch
    user: arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch
  name: context-we-should-not-touch
- context:
    cluster: arn:aws:eks:us-east-1:123456789000:cluster/test-eks-cluster
    namespace: crib-test
    user: arn:aws:eks:us-east-1:123456789000:cluster/test-eks-cluster
  name: test-eks-cluster-alias
current-context: context-we-should-not-touch
kind: Config
preferences:
  colors: true
users:
- name: arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch
  user:
    client-certificate: path/to/some/client/cert
    client-key: path/to/some/client/key
- name: arn:aws:eks:ap-southeast-1:123456789000:cluster/test-eks-cluster
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1beta1
      args:
      - eks
      - get-token
      - --cluster-name
      - wrong-eks-cluster
      command: aws
      env: null
      provideClusterInfo: false`), 0o666)
	defer os.Remove(mockedKubeConfig.Name())

	// mocking return of eks.DescribeCluster
	eksClusterName := "test-eks-cluster"
	eksClusterAlias := "test-eks-cluster-alias"
	eksClusterArn := "arn:aws:eks:ap-southeast-1:123456789000:cluster/test-eks-cluster"
	eksClusterEndpoint := "https://cluster.endpoint"
	eksEncodedCAData := base64.StdEncoding.EncodeToString([]byte("cadata"))
	mockEksClient := wrappermocks.NewEKSAPI(t)
	mockEksClient.EXPECT().
		DescribeCluster(
			context.TODO(), &eks.DescribeClusterInput{Name: &eksClusterName},
		).Return(
		&eks.DescribeClusterOutput{
			Cluster: &ekstypes.Cluster{
				Arn:                  &eksClusterArn,
				Endpoint:             &eksClusterEndpoint,
				CertificateAuthority: &ekstypes.Certificate{Data: &eksEncodedCAData},
			},
		}, nil,
	)

	setupKubeConfigInput := &SetupKubeConfigInput{
		EksClient:            mockEksClient,
		KubeconfigPath:       mockedKubeConfig.Name(),
		EksClusterName:       eksClusterName,
		EksAliasName:         eksClusterAlias,
		CribNamespace:        "crib-test",
		AwsProfile:           "profile-test",
		AwsRegion:            "ap-southeast-1",
		ChangeDefaultContext: true,
	}
	require.NoError(t, SetupKubeConfig(setupKubeConfigInput))

	got, err := clientcmd.LoadFromFile(mockedKubeConfig.Name())
	require.NoError(t, err)

	want := &clientcmdapi.Config{
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			eksClusterArn: {
				Exec: &clientcmdapi.ExecConfig{
					Command: "aws",
					Args:    []string{"--region", "ap-southeast-1", "eks", "get-token", "--cluster-name", eksClusterName, "--output", "json"},
					Env: []clientcmdapi.ExecEnvVar{
						{Name: "AWS_PROFILE", Value: "profile-test"},
					},
					APIVersion:      "client.authentication.k8s.io/v1beta1",
					InteractiveMode: "IfAvailable",
				},
			},
			"arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch": {
				ClientCertificate: "path/to/some/client/cert",
				ClientKey:         "path/to/some/client/key",
			},
		},
		Clusters: map[string]*clientcmdapi.Cluster{
			eksClusterArn: {
				Server:                   eksClusterEndpoint,
				CertificateAuthorityData: []byte("cadata"),
			},
			"arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch": {
				Server:                   "https://unrelated.endpoint",
				CertificateAuthorityData: []byte("something"),
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			eksClusterAlias: {
				Cluster:  eksClusterArn,
				AuthInfo: eksClusterArn,
			},
			"context-we-should-not-touch": {
				Cluster:  "arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch",
				AuthInfo: "arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch",
			},
		},
		CurrentContext: eksClusterAlias,
	}
	AssertEqualKubeConfigs(t, want, got)
}

func TestCheckEksAccess(t *testing.T) {
	t.Parallel()

	// mocking a successful call to CoreV1().Namespaces().List()
	mockedCoreV1NamespacesWorking := wrappermocks.NewNamespaceInterface(t)
	mockedCoreV1NamespacesWorking.EXPECT().
		List(
			context.TODO(), metav1.ListOptions{},
		).Return(
		&v1.NamespaceList{
			Items: []v1.Namespace{
				{ObjectMeta: metav1.ObjectMeta{Name: "some"}},
				{ObjectMeta: metav1.ObjectMeta{Name: "namespaces"}},
			},
		}, nil)

	mockedCoreV1ClientWorking := wrappermocks.NewCoreV1Interface(t)
	mockedCoreV1ClientWorking.EXPECT().Namespaces().Return(mockedCoreV1NamespacesWorking)

	// mocking a failed call to CoreV1().Namespaces().List()
	mockedCoreV1NamespacesNotWorking := wrappermocks.NewNamespaceInterface(t)
	mockedCoreV1NamespacesNotWorking.EXPECT().
		List(
			context.TODO(), metav1.ListOptions{},
		).Return(nil, fmt.Errorf("some error"))

	mockedCoreV1ClientNotWorking := wrappermocks.NewCoreV1Interface(t)
	mockedCoreV1ClientNotWorking.EXPECT().Namespaces().Return(mockedCoreV1NamespacesNotWorking)

	testCases := []struct {
		name        string
		corev1      corev1.CoreV1Interface
		listErr     error
		expectedErr string
	}{
		{
			name:        "Success",
			corev1:      mockedCoreV1ClientWorking,
			expectedErr: "",
		},
		{
			name:        "Error",
			corev1:      mockedCoreV1ClientNotWorking,
			expectedErr: "some error",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := CheckEksAccess(tt.corev1)
			if tt.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}
