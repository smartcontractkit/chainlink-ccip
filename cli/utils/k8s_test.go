package utils

import (
	"context"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
// TODO: use golang reflect to compare the entire object and recurse into subtypes
func AssertEqualKubeConfigs(t *testing.T, want *clientcmdapi.Config, got *clientcmdapi.Config) {
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
	mockedKubeConfig := MockKubeConfigFile([]byte(""), 0666)
	defer os.Remove(mockedKubeConfig.Name())

	// mocking return of eks.DescribeCluster
	eksClusterName := "test-eks-cluster"
	eksClusterAlias := "test-eks-cluster-alias"
	eksClusterArn := "arn:aws:eks:ap-southeast-1:123456789000:cluster/test-eks-cluster"
	eksClusterEndpoint := "https://cluster.endpoint"
	mockEksClient := wrappermocks.NewEKSAPI(t)
	mockEksClient.EXPECT().
		DescribeCluster(
			context.TODO(), &eks.DescribeClusterInput{Name: &eksClusterName},
		).Return(
		&eks.DescribeClusterOutput{
			Cluster: &ekstypes.Cluster{
				Arn:                  &eksClusterArn,
				Endpoint:             &eksClusterEndpoint,
				CertificateAuthority: &ekstypes.Certificate{Data: &eksClusterArn},
			},
		}, nil,
	)
	require.NoError(t, SetupKubeConfig(mockEksClient, mockedKubeConfig.Name(), eksClusterName, eksClusterAlias, "ap-southeast-1", true))

	got, err := clientcmd.LoadFromFile(mockedKubeConfig.Name())
	require.NoError(t, err)

	want := &clientcmdapi.Config{
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			eksClusterArn: &clientcmdapi.AuthInfo{
				Exec: &clientcmdapi.ExecConfig{
					Command:         "aws",
					Args:            []string{"eks", "get-token", "--cluster-name", eksClusterName},
					APIVersion:      "client.authentication.k8s.io/v1beta1",
					InteractiveMode: "IfAvailable",
				},
			},
		},
		Clusters: map[string]*clientcmdapi.Cluster{
			eksClusterArn: &clientcmdapi.Cluster{
				Server:                   eksClusterEndpoint,
				CertificateAuthorityData: []byte(eksClusterArn),
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			eksClusterAlias: &clientcmdapi.Context{
				Cluster:  eksClusterArn,
				AuthInfo: eksClusterArn,
			},
		},
		CurrentContext: eksClusterAlias,
	}
	AssertEqualKubeConfigs(t, want, got)
}

func TestSetupKubeConfigExistsButDiverges(t *testing.T) {
	mockedKubeConfig := MockKubeConfigFile([]byte(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: c29tZXRoaW5nCg==  # base64-encoded string "something"
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
      provideClusterInfo: false`), 0666)
	defer os.Remove(mockedKubeConfig.Name())

	// mocking return of eks.DescribeCluster
	eksClusterName := "test-eks-cluster"
	eksClusterAlias := "test-eks-cluster-alias"
	eksClusterArn := "arn:aws:eks:ap-southeast-1:123456789000:cluster/test-eks-cluster"
	eksClusterEndpoint := "https://cluster.endpoint"
	mockEksClient := wrappermocks.NewEKSAPI(t)
	mockEksClient.EXPECT().
		DescribeCluster(
			context.TODO(), &eks.DescribeClusterInput{Name: &eksClusterName},
		).Return(
		&eks.DescribeClusterOutput{
			Cluster: &ekstypes.Cluster{
				Arn:                  &eksClusterArn,
				Endpoint:             &eksClusterEndpoint,
				CertificateAuthority: &ekstypes.Certificate{Data: &eksClusterArn},
			},
		}, nil,
	)
	require.NoError(t, SetupKubeConfig(mockEksClient, mockedKubeConfig.Name(), eksClusterName, eksClusterAlias, "ap-southeast-1", true))

	got, err := clientcmd.LoadFromFile(mockedKubeConfig.Name())
	require.NoError(t, err)

	want := &clientcmdapi.Config{
		AuthInfos: map[string]*clientcmdapi.AuthInfo{
			eksClusterArn: &clientcmdapi.AuthInfo{
				Exec: &clientcmdapi.ExecConfig{
					Command:         "aws",
					Args:            []string{"eks", "get-token", "--cluster-name", eksClusterName},
					APIVersion:      "client.authentication.k8s.io/v1beta1",
					InteractiveMode: "IfAvailable",
				},
			},
			"arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch": &clientcmdapi.AuthInfo{
				ClientCertificate: "path/to/some/client/cert",
				ClientKey:         "path/to/some/client/key",
			},
		},
		Clusters: map[string]*clientcmdapi.Cluster{
			eksClusterArn: &clientcmdapi.Cluster{
				Server:                   eksClusterEndpoint,
				CertificateAuthorityData: []byte(eksClusterArn),
			},
			"arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch": &clientcmdapi.Cluster{
				Server:                   "https://unrelated.endpoint",
				CertificateAuthorityData: []byte("something"),
			},
		},
		Contexts: map[string]*clientcmdapi.Context{
			eksClusterAlias: &clientcmdapi.Context{
				Cluster:  eksClusterArn,
				AuthInfo: eksClusterArn,
			},
			"context-we-should-not-touch": &clientcmdapi.Context{
				Cluster:  "arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch",
				AuthInfo: "arn:aws:eks:us-east-1:123456789000:cluster/cluster-we-should-not-touch",
			},
		},
		CurrentContext: eksClusterAlias,
	}
	AssertEqualKubeConfigs(t, want, got)

	// reading file content
	content, err := ioutil.ReadFile(mockedKubeConfig.Name()) // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content)) // This is some content
	t.Fail()
}
