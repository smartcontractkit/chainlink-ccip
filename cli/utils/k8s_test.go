package utils_test

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	ekstypes "github.com/aws/aws-sdk-go-v2/service/eks/types"
	k8smocks "github.com/smartcontractkit/crib/cli/mocks/external/kubernetes"
	testingutils "github.com/smartcontractkit/crib/cli/testing/utils"
	"github.com/smartcontractkit/crib/cli/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	networkingv1api "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

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

	setupKubeConfigInput := &utils.SetupKubeConfigInput{
		EksClient:            mockEksClient,
		KubeconfigPath:       nonExistingKubeConfig,
		EksClusterName:       eksClusterName,
		EksAliasName:         eksClusterAlias,
		CribNamespace:        "crib-test",
		AwsProfile:           "profile-test",
		AwsRegion:            "ap-southeast-1",
		ChangeDefaultContext: true,
	}
	require.NoError(t, utils.SetupKubeConfig(setupKubeConfigInput))
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

	mockedKubeConfig := testingutils.MockKubeConfigFile([]byte(`apiVersion: v1
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

	setupKubeConfigInput := &utils.SetupKubeConfigInput{
		EksClient:            mockEksClient,
		KubeconfigPath:       mockedKubeConfig.Name(),
		EksClusterName:       eksClusterName,
		EksAliasName:         eksClusterAlias,
		CribNamespace:        "crib-test",
		AwsProfile:           "profile-test",
		AwsRegion:            "ap-southeast-1",
		ChangeDefaultContext: true,
	}
	require.NoError(t, utils.SetupKubeConfig(setupKubeConfigInput))

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

func TestLabelNamespace(t *testing.T) {
	t.Parallel()

	namespace := "test-namespace"
	labelKey := "test-key"
	labelValue := "test-value"

	tests := []struct {
		name                     string
		namespace                string
		labelKey                 string
		labelValue               string
		applyNamespacesMockCalls func(m *k8smocks.NamespaceInterface)
		expectErr                string
	}{
		{
			name: "Success",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().
					Patch(
						context.TODO(),
						namespace,
						types.MergePatchType,
						[]byte(`{"metadata":{"labels":{"test-key":"test-value"}}}`),
						metav1.PatchOptions{},
					).Return(&v1.Namespace{}, nil).Times(1)
			},
			expectErr: "",
		},
		{
			name: "Error",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().
					Patch(
						context.TODO(),
						namespace,
						types.MergePatchType,
						[]byte(`{"metadata":{"labels":{"test-key":"test-value"}}}`),
						metav1.PatchOptions{},
					).Return(nil, fmt.Errorf("some error")).Times(1)
			},
			expectErr: "some error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockNamespaces := k8smocks.NewNamespaceInterface(t)
			tt.applyNamespacesMockCalls(mockNamespaces)

			mockCoreV1 := k8smocks.NewCoreV1Interface(t)
			mockCoreV1.EXPECT().Namespaces().Return(mockNamespaces)

			mockClientset := wrappermocks.NewK8sClientset(t)
			mockClientset.EXPECT().CoreV1().Return(mockCoreV1)

			configFlags := &genericclioptions.ConfigFlags{}
			k8sClient, err := wrappers.NewK8sClient(configFlags, mockClientset)
			require.NoError(t, err)

			err = k8sClient.LabelNamespace(context.TODO(), namespace, labelKey, labelValue)
			if tt.expectErr != "" {
				require.Error(t, err)
				assert.Equal(t, tt.expectErr, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestGetIngress(t *testing.T) {
	t.Parallel()

	namespace := "test-namespace"
	ingressName := "test-ingress"

	tests := []struct {
		name                    string
		applyIngressesMockCalls func(m *k8smocks.IngressInterface)
		expectedIngress         *networkingv1api.Ingress
		expectErr               string
	}{
		{
			name: "Success",
			applyIngressesMockCalls: func(m *k8smocks.IngressInterface) {
				m.EXPECT().
					Get(
						context.TODO(),
						ingressName,
						metav1.GetOptions{},
					).Return(&networkingv1api.Ingress{
					ObjectMeta: metav1.ObjectMeta{
						Name: ingressName,
					},
				}, nil).Times(1)
			},
			expectedIngress: &networkingv1api.Ingress{
				ObjectMeta: metav1.ObjectMeta{
					Name: ingressName,
				},
			},
			expectErr: "",
		},
		{
			name: "Error",
			applyIngressesMockCalls: func(m *k8smocks.IngressInterface) {
				m.EXPECT().
					Get(
						context.TODO(),
						ingressName,
						metav1.GetOptions{},
					).Return(nil, fmt.Errorf("some error")).Times(1)
			},
			expectedIngress: nil,
			expectErr:       "some error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockIngresses := k8smocks.NewIngressInterface(t)
			tt.applyIngressesMockCalls(mockIngresses)

			mockNetworkingV1 := k8smocks.NewNetworkingV1Interface(t)
			mockNetworkingV1.EXPECT().Ingresses(namespace).Return(mockIngresses)

			mockClientset := wrappermocks.NewK8sClientset(t)
			mockClientset.EXPECT().NetworkingV1().Return(mockNetworkingV1)

			configFlags := &genericclioptions.ConfigFlags{}
			k8sClient, err := wrappers.NewK8sClient(configFlags, mockClientset)
			require.NoError(t, err)

			ingress, err := k8sClient.GetIngress(context.TODO(), namespace, ingressName)
			if tt.expectErr != "" {
				require.Error(t, err)
				assert.Equal(t, tt.expectErr, err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedIngress, ingress)
			}
		})
	}
}

func TestListIngresses(t *testing.T) {
	t.Parallel()

	namespace := "test-namespace"

	tests := []struct {
		name                    string
		applyIngressesMockCalls func(m *k8smocks.IngressInterface)
		expectedIngressList     *networkingv1api.IngressList
		expectErr               string
	}{
		{
			name: "Success",
			applyIngressesMockCalls: func(m *k8smocks.IngressInterface) {
				m.EXPECT().
					List(
						context.TODO(),
						metav1.ListOptions{},
					).Return(&networkingv1api.IngressList{
					Items: []networkingv1api.Ingress{
						{
							ObjectMeta: metav1.ObjectMeta{
								Name: "test-ingress",
							},
						},
					},
				}, nil).Times(1)
			},
			expectedIngressList: &networkingv1api.IngressList{
				Items: []networkingv1api.Ingress{
					{
						ObjectMeta: metav1.ObjectMeta{
							Name: "test-ingress",
						},
					},
				},
			},
			expectErr: "",
		},
		{
			name: "Error",
			applyIngressesMockCalls: func(m *k8smocks.IngressInterface) {
				m.EXPECT().
					List(
						context.TODO(),
						metav1.ListOptions{},
					).Return(nil, fmt.Errorf("some error")).Times(1)
			},
			expectedIngressList: nil,
			expectErr:           "some error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockIngresses := k8smocks.NewIngressInterface(t)
			tt.applyIngressesMockCalls(mockIngresses)

			mockNetworkingV1 := k8smocks.NewNetworkingV1Interface(t)
			mockNetworkingV1.EXPECT().Ingresses(namespace).Return(mockIngresses)

			mockClientset := wrappermocks.NewK8sClientset(t)
			mockClientset.EXPECT().NetworkingV1().Return(mockNetworkingV1)

			configFlags := &genericclioptions.ConfigFlags{}
			k8sClient, err := wrappers.NewK8sClient(configFlags, mockClientset)
			require.NoError(t, err)

			ingressList, err := k8sClient.ListIngresses(context.TODO(), namespace)
			if tt.expectErr != "" {
				require.Error(t, err)
				assert.Equal(t, tt.expectErr, err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedIngressList, ingressList)
			}
		})
	}
}
