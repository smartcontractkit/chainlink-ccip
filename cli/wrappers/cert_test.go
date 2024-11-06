package wrappers_test

import (
	"context"
	"testing"

	certmanagerv1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	k8smocks "github.com/smartcontractkit/crib/cli/mocks/external/kubernetes"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

func TestNewCACert(t *testing.T) {
	t.Parallel()

	mockCrk8sClient := k8smocks.NewControllerRuntimeClient(t)
	mockK8sClient := wrappermocks.NewK8sCLI(t)
	mockK8sClient.EXPECT().ControllerRuntimeClient(mock.Anything).Return(mockCrk8sClient, nil)

	caCert, err := wrappers.NewCACert(mockK8sClient)
	require.NoError(t, err)
	assert.NotNil(t, caCert)
}

// nolint: paralleltest,nolintlint
func TestCACert_EnsureCertManagerSecret(t *testing.T) {
	secretName := "test-secret"
	secretNamespace := "default"

	testCases := []struct {
		name                      string
		applyCrk8sClientMockCalls func(m *k8smocks.ControllerRuntimeClient)
		expectedErr               string
		expectedExists            bool
	}{
		{
			name: "SecretDoesNotExistCreateSucceeds",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: secretNamespace, Name: secretName}, &corev1.Secret{}).Return(k8serrors.NewNotFound(corev1.Resource("secrets"), "test-secret"))
				m.EXPECT().Create(context.TODO(), mock.Anything).Return(nil)
			},
			expectedErr:    "",
			expectedExists: false,
		},
		{
			name: "SecretDoesNotExistCreateFails",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: secretNamespace, Name: secretName}, &corev1.Secret{}).Return(k8serrors.NewNotFound(corev1.Resource("secrets"), "test-secret"))
				m.EXPECT().Create(context.TODO(), mock.Anything).Return(k8serrors.NewServiceUnavailable("failed creating secret"))
			},
			expectedErr:    "failed creating secret",
			expectedExists: false,
		},
		{
			name: "SecretAlreadyExists",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: secretNamespace, Name: secretName}, &corev1.Secret{}).Return(k8serrors.NewAlreadyExists(corev1.Resource("secrets"), "test-secret"))
			},
			expectedErr:    "",
			expectedExists: true,
		},
		{
			name: "GetSecretError",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: secretNamespace, Name: secretName}, &corev1.Secret{}).Return(k8serrors.NewServiceUnavailable("get secret error"))
			},
			expectedErr:    "get secret error",
			expectedExists: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockCrk8sClient := k8smocks.NewControllerRuntimeClient(t)
			tt.applyCrk8sClientMockCalls(mockCrk8sClient)

			mockK8sClient := wrappermocks.NewK8sCLI(t)
			mockK8sClient.EXPECT().ControllerRuntimeClient(mock.Anything).Return(mockCrk8sClient, nil)

			caCert, err := wrappers.NewCACert(mockK8sClient)
			require.NoError(t, err)
			require.NotNil(t, caCert)

			exists, err := caCert.EnsureCertManagerSecret(context.TODO(), secretName, secretNamespace)
			assert.Equal(t, tt.expectedExists, exists)
			if tt.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// nolint: paralleltest,nolintlint
func TestCACert_EnsureCAClusterIssuer(t *testing.T) {
	secretName := "test-secret"
	clusterIssuerName := "test-cluster-issuer"
	clusterIssuerNamespace := "default"

	testCases := []struct {
		name                      string
		applyCrk8sClientMockCalls func(m *k8smocks.ControllerRuntimeClient)
		expectedErr               string
	}{
		{
			name: "ClusterIssuerDoesNotExistCreateSucceeds",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: "", Name: clusterIssuerName}, &certmanagerv1.ClusterIssuer{}).Return(k8serrors.NewNotFound(corev1.Resource("cluster-issuer"), clusterIssuerName))
				m.EXPECT().Create(context.TODO(), mock.Anything).Return(nil)
			},
			expectedErr: "",
		},
		{
			name: "ClusterIssuerDoesNotExistCreateFails",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: "", Name: clusterIssuerName}, &certmanagerv1.ClusterIssuer{}).Return(k8serrors.NewNotFound(corev1.Resource("cluster-issuer"), clusterIssuerName))
				m.EXPECT().Create(context.TODO(), mock.Anything).Return(k8serrors.NewServiceUnavailable("error creating cluster issuer"))
			},
			expectedErr: "error creating cluster issuer",
		},
		{
			name: "ClusterIssuerAlreadyExistsUpdateSucceeds",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: "", Name: clusterIssuerName}, &certmanagerv1.ClusterIssuer{}).Return(k8serrors.NewAlreadyExists(corev1.Resource("cluster-issuer"), clusterIssuerName))
				m.EXPECT().Update(context.TODO(), mock.Anything).Return(nil)
			},
			expectedErr: "",
		},
		{
			name: "ClusterIssuerAlreadyExistsUpdateFails",
			applyCrk8sClientMockCalls: func(m *k8smocks.ControllerRuntimeClient) {
				m.EXPECT().Get(context.TODO(), types.NamespacedName{Namespace: "", Name: clusterIssuerName}, &certmanagerv1.ClusterIssuer{}).Return(nil)
				m.EXPECT().Update(context.TODO(), mock.Anything).Return(k8serrors.NewServiceUnavailable("error updating cluster issuer"))
			},
			expectedErr: "error updating cluster issuer",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			mockCrk8sClient := k8smocks.NewControllerRuntimeClient(t)
			tt.applyCrk8sClientMockCalls(mockCrk8sClient)

			mockK8sClient := wrappermocks.NewK8sCLI(t)
			mockK8sClient.EXPECT().ControllerRuntimeClient(mock.Anything).Return(mockCrk8sClient, nil)

			caCert, err := wrappers.NewCACert(mockK8sClient)
			require.NoError(t, err)
			require.NotNil(t, caCert)

			err = caCert.EnsureCAClusterIssuer(context.TODO(), secretName, clusterIssuerName, clusterIssuerNamespace)
			if tt.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

// nolint: paralleltest,nolintlint
func TestMkcertInstall(t *testing.T) {
	m := wrappers.Mkcert{}
	err := m.Install()
	assert.NoError(t, err)
	assert.FileExists(t, m.CertPath)
	assert.FileExists(t, m.KeyPath)
}

// nolint: paralleltest,nolintlint
func TestMkcertCaRoot(t *testing.T) {
	m := wrappers.Mkcert{}
	caRoot, err := m.CARoot()
	assert.NoError(t, err)
	assert.DirExists(t, caRoot)
}
