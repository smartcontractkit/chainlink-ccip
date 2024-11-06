package wrappers_test

import (
	// "fmt"
	"context"
	"errors"
	"os"
	"testing"
	"time"

	k8smocks "github.com/smartcontractkit/crib/cli/mocks/external/kubernetes"
	testingutils "github.com/smartcontractkit/crib/cli/testing/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/kubernetes/fake"
)

func TestNewK8sClient(t *testing.T) {
	t.Parallel()

	configFlags := &genericclioptions.ConfigFlags{}
	k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
	require.NoError(t, err)
	require.NotNil(t, k8sClient)

	assert.NotNil(t, k8sClient.Clientset())
	assert.NotNil(t, k8sClient.ClientConfig())
	assert.NotNil(t, k8sClient.RestConfig())

	fakeClientset := fake.NewSimpleClientset()
	k8sClient, err = wrappers.NewK8sClient(configFlags, fakeClientset)
	require.NoError(t, err)
	require.NotNil(t, k8sClient)
	assert.Equal(t, fakeClientset, k8sClient.Clientset())
}

func TestK8sClient_CurrentContext(t *testing.T) {
	t.Parallel()

	mockKubeConfig := testingutils.MockKubeConfigFile([]byte(`apiVersion: v1
clusters:
- cluster:
    server: https://some.endpoint
  name: arn:aws:eks:us-east-1:123456789000:cluster/some-cluster
- cluster:
    server: https://other.endpoint
  name: arn:aws:eks:ap-southeast-1:123456789000:cluster/some-other-cluster
contexts:
- context:
    cluster: arn:aws:eks:us-east-1:123456789000:cluster/some-cluster
    user: arn:aws:eks:us-east-1:123456789000:cluster/some-cluster
  name: context-some-cluster
- context:
    cluster: arn:aws:eks:us-east-1:123456789000:cluster/some-other-cluster
    namespace: crib-test
    user: arn:aws:eks:us-east-1:123456789000:cluster/some-other-cluster
  name: context-some-other-cluster
current-context: context-some-cluster
kind: Config`), 0o644)
	mockKubeConfigFile := mockKubeConfig.Name()
	defer os.Remove(mockKubeConfigFile)

	k8sClient, err := wrappers.NewK8sClient(&genericclioptions.ConfigFlags{KubeConfig: &mockKubeConfigFile}, nil)
	require.NoError(t, err)
	require.NotNil(t, k8sClient)

	assert.Equal(t, "context-some-cluster", k8sClient.CurrentContext())
}

func TestK8sClient_CheckAccess(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name                     string
		applyNamespacesMockCalls func(*k8smocks.NamespaceInterface)
		expectedErr              error
	}

	for _, tt := range []testCase{
		{
			name: "Success",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().List(context.TODO(), metav1.ListOptions{}).Return(&v1.NamespaceList{}, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().List(context.TODO(), metav1.ListOptions{}).Return(nil, errors.New("some error listing namespaces"))
			},
			expectedErr: errors.New("some error listing namespaces"),
		},
	} {
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

			err = k8sClient.CheckAccess(context.TODO())
			if tt.expectedErr == nil {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err)
			}
		})
	}
}

func TestK8sClient_EnsureNamespaceExists(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                     string
		namespaceName            string
		applyNamespacesMockCalls func(*k8smocks.NamespaceInterface)
		expectedExists           bool
		expectedErr              string
	}{
		{
			name:          "NamespaceExists",
			namespaceName: "existing-namespace",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().Get(context.TODO(), "existing-namespace", metav1.GetOptions{}).Return(&v1.Namespace{}, nil)
			},
			expectedExists: true,
			expectedErr:    "",
		},
		{
			name:          "NamespaceExistsWithK8sError",
			namespaceName: "existing-namespace",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().Get(context.TODO(), "existing-namespace", metav1.GetOptions{}).Return(&v1.Namespace{}, k8serrors.NewAlreadyExists(v1.Resource("namespace"), "existing-namespace"))
			},
			expectedExists: true,
			expectedErr:    "",
		},
		{
			name:          "NamespaceDoesNotExistAndCreateSucceeds",
			namespaceName: "new-namespace",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().Get(context.TODO(), "new-namespace", metav1.GetOptions{}).Return(&v1.Namespace{}, k8serrors.NewNotFound(v1.Resource("namespace"), "new-namespace"))
				m.EXPECT().Create(context.TODO(), &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "new-namespace"}}, metav1.CreateOptions{}).Return(&v1.Namespace{}, nil)
			},
			expectedExists: false,
			expectedErr:    "",
		},
		{
			name:          "NamespaceDoesNotExistAndCreateFails",
			namespaceName: "new-namespace",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().Get(context.TODO(), "new-namespace", metav1.GetOptions{}).Return(&v1.Namespace{}, k8serrors.NewNotFound(v1.Resource("namespaces"), "new-namespace"))
				m.EXPECT().Create(context.TODO(), &v1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "new-namespace"}}, metav1.CreateOptions{}).Return(&v1.Namespace{}, k8serrors.NewServiceUnavailable("server says no"))
			},
			expectedExists: false,
			expectedErr:    "failed to create namespace new-namespace: server says no",
		},
		{
			name:          "GetNamespaceError",
			namespaceName: "some-namespace",
			applyNamespacesMockCalls: func(m *k8smocks.NamespaceInterface) {
				m.EXPECT().Get(context.TODO(), "some-namespace", metav1.GetOptions{}).Return(nil, k8serrors.NewUnauthorized("get error"))
			},
			expectedExists: false,
			expectedErr:    "failed to get namespace some-namespace: get error",
		},
	}

	for _, tt := range testCases {
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

			exists, err := k8sClient.EnsureNamespaceExists(context.TODO(), tt.namespaceName)
			assert.Equal(t, tt.expectedExists, exists)
			if tt.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}

func TestK8sClient_WaitForResource(t *testing.T) {
	t.Parallel()

	interval := 100 * time.Millisecond
	timeout := 500 * time.Millisecond

	testCases := []struct {
		name                   string
		applyResourceMockCalls func(*k8smocks.ResourceInterface)
		expectedErr            string
		resourceFound          bool
	}{
		{
			name: "ResourceFound",
			applyResourceMockCalls: func(m *k8smocks.ResourceInterface) {
				m.EXPECT().Get(context.TODO(), "test-resource", metav1.GetOptions{}).Return(&unstructured.Unstructured{}, nil).Times(1)
			},
			expectedErr:   "",
			resourceFound: true,
		},
		{
			name: "ResourceNotFound",
			applyResourceMockCalls: func(m *k8smocks.ResourceInterface) {
				m.EXPECT().Get(context.TODO(), "test-resource", metav1.GetOptions{}).Return(nil, errors.New("not found"))
			},
			expectedErr:   "timed out waiting for resource test-resource",
			resourceFound: false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockResourceClient := k8smocks.NewResourceInterface(t)
			tt.applyResourceMockCalls(mockResourceClient)

			configFlags := &genericclioptions.ConfigFlags{}
			k8sClient, err := wrappers.NewK8sClient(configFlags, nil)
			require.NoError(t, err)

			err = k8sClient.WaitForResource(context.TODO(), mockResourceClient, "test-resource", interval, timeout)
			if tt.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}

func TestK8sClient_ApplyConfigMap(t *testing.T) {
	t.Parallel()

	name := "test-configmap"
	namespace := "test-namespace"
	configMap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Data: map[string]string{
			"key": "value",
		},
	}

	testCases := []struct {
		name                     string
		applyConfigMapsMockCalls func(*k8smocks.ConfigMapInterface)
		expectedExists           bool
		expectedErr              string
	}{
		{
			name: "ConfigMapExists",
			applyConfigMapsMockCalls: func(m *k8smocks.ConfigMapInterface) {
				m.EXPECT().Get(context.TODO(), name, metav1.GetOptions{}).Return(&v1.ConfigMap{}, nil)
				m.EXPECT().Update(context.TODO(), configMap, metav1.UpdateOptions{}).Return(&v1.ConfigMap{}, nil)
			},
			expectedExists: true,
			expectedErr:    "",
		},
		{
			name: "ConfigMapWithK8sError",
			applyConfigMapsMockCalls: func(m *k8smocks.ConfigMapInterface) {
				m.EXPECT().Get(context.TODO(), name, metav1.GetOptions{}).Return(&v1.ConfigMap{}, k8serrors.NewAlreadyExists(v1.Resource("configmap"), name))
				m.EXPECT().Update(context.TODO(), configMap, metav1.UpdateOptions{}).Return(&v1.ConfigMap{}, errors.New("update error"))
			},
			expectedExists: true,
			expectedErr:    "failed to update configmap test-configmap in namespace test-namespace: update error",
		},
		{
			name: "ConfigMapDoesNotExistAndCreateSucceeds",
			applyConfigMapsMockCalls: func(m *k8smocks.ConfigMapInterface) {
				m.EXPECT().Get(context.TODO(), name, metav1.GetOptions{}).Return(&v1.ConfigMap{}, k8serrors.NewNotFound(v1.Resource("configmap"), name))
				m.EXPECT().Create(context.TODO(), configMap, metav1.CreateOptions{}).Return(&v1.ConfigMap{}, nil)
			},
			expectedExists: false,
			expectedErr:    "",
		},
		{
			name: "ConfigMapDoesNotExistAndCreateFails",
			applyConfigMapsMockCalls: func(m *k8smocks.ConfigMapInterface) {
				m.EXPECT().Get(context.TODO(), name, metav1.GetOptions{}).Return(&v1.ConfigMap{}, k8serrors.NewNotFound(v1.Resource("configmap"), name))
				m.EXPECT().Create(context.TODO(), configMap, metav1.CreateOptions{}).Return(&v1.ConfigMap{}, k8serrors.NewServiceUnavailable("server says no"))
			},
			expectedExists: false,
			expectedErr:    "failed to create configmap test-configmap in namespace test-namespace: server says no",
		},
		{
			name: "GetConfigMapError",
			applyConfigMapsMockCalls: func(m *k8smocks.ConfigMapInterface) {
				m.EXPECT().Get(context.TODO(), name, metav1.GetOptions{}).Return(&v1.ConfigMap{}, errors.New("get error"))
			},
			expectedExists: false,
			expectedErr:    "failed to get configmap test-configmap in namespace test-namespace: get error",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockConfigMaps := k8smocks.NewConfigMapInterface(t)
			tt.applyConfigMapsMockCalls(mockConfigMaps)

			mockCoreV1 := k8smocks.NewCoreV1Interface(t)
			mockCoreV1.EXPECT().ConfigMaps(namespace).Return(mockConfigMaps)

			mockClientset := wrappermocks.NewK8sClientset(t)
			mockClientset.EXPECT().CoreV1().Return(mockCoreV1)

			configFlags := &genericclioptions.ConfigFlags{}
			k8sClient, err := wrappers.NewK8sClient(configFlags, mockClientset)
			require.NoError(t, err)

			exists, err := k8sClient.ApplyConfigMap(context.TODO(), configMap)
			assert.Equal(t, tt.expectedExists, exists)
			if tt.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			}
		})
	}
}
