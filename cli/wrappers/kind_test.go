package wrappers_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	kindmocks "github.com/smartcontractkit/crib/cli/mocks/external/kind"
	testingutils "github.com/smartcontractkit/crib/cli/testing/utils"
	"github.com/smartcontractkit/crib/cli/wrappers"
	wrappermocks "github.com/smartcontractkit/crib/cli/wrappers/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd/api"
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
	"sigs.k8s.io/kind/pkg/cluster/nodes"
)

func TestNewKindCluster_CreateOrReuse(t *testing.T) {
	t.Parallel()

	clusterName := "test-cluster"
	registryName := "test-registry"
	namespaceName := "test-namespace"
	registryContainerConfig := &containertypes.Config{
		Image: "someregistry:sometag",
	}
	containerPort, _ := nat.NewPort("tcp", "5000")
	registryContainerHostConfig := &containertypes.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: "5001"}},
		},
		RestartPolicy: containertypes.RestartPolicy{
			Name: "unless-stopped",
		},
		NetworkMode: "bridge",
	}

	kubeConfigContent := testingutils.MockKubeConfigFile([]byte(`apiVersion: v1
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
kind: Config`), 0o600)
	kubeConfigFile := kubeConfigContent.Name()
	t.Cleanup(func() { os.Remove(kubeConfigFile) })

	testCases := []struct {
		name                       string
		applyKindProviderMockCalls func(m *wrappermocks.KindProvider)
		applyDockerCLIMockCalls    func(m *wrappermocks.DockerCLI)
		applyK8sCLIMockCalls       func(m *wrappermocks.K8sCLI)
		applyKubeConfigMockCalls   func(m *wrappermocks.KubeConfigInterface)
		expectedErr                string
	}{
		{
			name: "CreateSucceeds",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {
				kindExecCmd1 := kindmocks.NewKindExecCmd(t)
				kindExecCmd1.EXPECT().Run().Return(nil)
				kindExecCmd2 := kindmocks.NewKindExecCmd(t)
				kindExecCmd2.EXPECT().Run().Return(nil)

				kindNode := kindmocks.NewKindNode(t)
				kindNode.EXPECT().Command("mkdir", "-p", "/etc/containerd/certs.d/localhost:5001").Return(kindExecCmd1)
				kindNode.EXPECT().Command("sh", "-c", `echo '[host."http://test-registry:5000"]' > /etc/containerd/certs.d/localhost:5001/hosts.toml`).Return(kindExecCmd2)

				m.EXPECT().List().Return([]string{}, nil)
				m.EXPECT().Create(
					clusterName, mock.AnythingOfType("cluster.createOptionAdapter"), mock.AnythingOfType("cluster.createOptionAdapter"),
				).Return(nil)
				m.EXPECT().ExportKubeConfig(clusterName, mock.AnythingOfType("string"), false).Return(nil)
				m.EXPECT().
					ListNodes(clusterName).Return([]nodes.Node{kindNode}, nil)
			},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().
					RunContainer(
						context.Background(), registryContainerConfig, registryContainerHostConfig, &network.NetworkingConfig{}, registryName, false, mock.Anything,
					).Return(false, nil)
				m.EXPECT().
					ConnectContainerToNetwork(
						context.Background(), registryName, "kind",
					).Return(false, nil)
			},
			applyK8sCLIMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					ApplyConfigMap(context.TODO(), &corev1.ConfigMap{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "local-registry-hosting",
							Namespace: "kube-public",
						},
						Data: map[string]string{
							"localRegistryHosting.v1": `host: "localhost:5001"\nhelp: "https://kind.sigs.k8s.io/docs/user/local-registry/"`,
						},
					}).Return(false, nil)
			},
			applyKubeConfigMockCalls: func(m *wrappermocks.KubeConfigInterface) {
				m.EXPECT().Path().Return(kubeConfigFile)
				m.EXPECT().LoadConfig().Return(nil)
				// NOTE: here the kubeconfig is updated
				kindContext := fmt.Sprintf("kind-%s", clusterName)
				m.EXPECT().CurrentContext().Return(kindContext)
				m.EXPECT().Contexts().Return(map[string]*api.Context{
					"context-some-cluster": {
						Cluster:  "arn:aws:eks:us-east-1:123456789000:cluster/some-cluster",
						AuthInfo: "arn:aws:eks:us-east-1:123456789000:cluster/some-cluster",
					},
					"context-some-other-cluster": {
						Cluster:   "arn:aws:eks:us-east-1:123456789000:cluster/some-other-cluster",
						Namespace: "crib-test",
						AuthInfo:  "arn:aws:eks:us-east-1:123456789000:cluster/some-other-cluster",
					},
					kindContext: { // NOTE: in the new kind context, cluster and user have the same name as the cluster as per kind convention
						Cluster:  kindContext,
						AuthInfo: kindContext,
					},
				})
				m.EXPECT().SetNamespaceForContext(kindContext, namespaceName).Return(nil)
			},
			expectedErr: "",
		},
		{
			name: "ReuseSucceeds",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {
				kindExecCmd1 := kindmocks.NewKindExecCmd(t)
				kindExecCmd1.EXPECT().Run().Return(nil)
				kindExecCmd2 := kindmocks.NewKindExecCmd(t)
				kindExecCmd2.EXPECT().Run().Return(nil)

				kindNodeControlPlane := kindmocks.NewKindNode(t)
				kindNodeControlPlane.EXPECT().Role().Return("control-plane", nil)
				kindNodeControlPlane.EXPECT().String().Return(fmt.Sprintf("%s-%s", clusterName, "control-plane"))
				kindNodeControlPlane.EXPECT().Command("mkdir", "-p", "/etc/containerd/certs.d/localhost:5001").Return(kindExecCmd1)
				kindNodeControlPlane.EXPECT().Command("sh", "-c", `echo '[host."http://test-registry:5000"]' > /etc/containerd/certs.d/localhost:5001/hosts.toml`).Return(kindExecCmd2)

				kindNodeWorker := kindmocks.NewKindNode(t)
				kindNodeWorker.EXPECT().Role().Return("worker", nil)
				kindNodeWorker.EXPECT().String().Return(fmt.Sprintf("%s-%s", clusterName, "worker"))
				kindNodeWorker.EXPECT().Command("mkdir", "-p", "/etc/containerd/certs.d/localhost:5001").Return(kindExecCmd1)
				kindNodeWorker.EXPECT().Command("sh", "-c", `echo '[host."http://test-registry:5000"]' > /etc/containerd/certs.d/localhost:5001/hosts.toml`).Return(kindExecCmd2)

				m.EXPECT().List().Return([]string{clusterName}, nil)
				m.EXPECT().ListNodes(clusterName).Return([]nodes.Node{kindNodeControlPlane, kindNodeWorker}, nil)
				m.EXPECT().ExportKubeConfig(clusterName, mock.AnythingOfType("string"), false).Return(nil)
			},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().
					RunContainer(
						context.Background(), registryContainerConfig, registryContainerHostConfig, &network.NetworkingConfig{}, registryName, false, mock.Anything,
					).Return(true, nil)
				m.EXPECT().
					RunContainer(
						context.Background(), mock.Anything, mock.Anything, mock.Anything, fmt.Sprintf("%s-control-plane", clusterName), false, mock.Anything,
					).Return(false, nil)
				m.EXPECT().
					RunContainer(
						context.Background(), mock.Anything, mock.Anything, mock.Anything, fmt.Sprintf("%s-worker", clusterName), false, mock.Anything,
					).Return(false, nil)
				m.EXPECT().
					ConnectContainerToNetwork(
						context.Background(), registryName, "kind",
					).Return(true, nil)
			},
			applyK8sCLIMockCalls: func(m *wrappermocks.K8sCLI) {
				m.EXPECT().
					ApplyConfigMap(context.TODO(), &corev1.ConfigMap{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "local-registry-hosting",
							Namespace: "kube-public",
						},
						Data: map[string]string{
							"localRegistryHosting.v1": `host: "localhost:5001"\nhelp: "https://kind.sigs.k8s.io/docs/user/local-registry/"`,
						},
					}).Return(false, nil)
			},
			applyKubeConfigMockCalls: func(m *wrappermocks.KubeConfigInterface) {
				m.EXPECT().Path().Return(kubeConfigFile)
				m.EXPECT().LoadConfig().Return(nil)
				// NOTE: here the kubeconfig is updated
				kindContext := fmt.Sprintf("kind-%s", clusterName)
				m.EXPECT().CurrentContext().Return(kindContext)
				m.EXPECT().Contexts().Return(map[string]*api.Context{
					"context-some-cluster": {
						Cluster:  "arn:aws:eks:us-east-1:123456789000:cluster/some-cluster",
						AuthInfo: "arn:aws:eks:us-east-1:123456789000:cluster/some-cluster",
					},
					"context-some-other-cluster": {
						Cluster:   "arn:aws:eks:us-east-1:123456789000:cluster/some-other-cluster",
						Namespace: "crib-test",
						AuthInfo:  "arn:aws:eks:us-east-1:123456789000:cluster/some-other-cluster",
					},
					kindContext: { // NOTE: in the new kind context, cluster and user have the same name as the cluster as per kind convention
						Cluster:  kindContext,
						AuthInfo: kindContext,
					},
				})
				m.EXPECT().SetNamespaceForContext(kindContext, namespaceName).Return(nil)
			},
			expectedErr: "",
		},
		{
			name: "CreateFails",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {
				m.EXPECT().List().Return([]string{}, nil)
				m.EXPECT().Create(
					clusterName, mock.AnythingOfType("cluster.createOptionAdapter"), mock.AnythingOfType("cluster.createOptionAdapter"),
				).Return(errors.New("error creating cluster"))
			},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().
					RunContainer(
						context.Background(), registryContainerConfig, registryContainerHostConfig, &network.NetworkingConfig{}, registryName, false, mock.Anything,
					).Return(false, nil)
			},
			applyK8sCLIMockCalls: func(m *wrappermocks.K8sCLI) {},
			applyKubeConfigMockCalls: func(m *wrappermocks.KubeConfigInterface) {
				m.EXPECT().Path().Return(kubeConfigFile)
			},
			expectedErr: "error creating cluster",
		},
		{
			name: "ReuseFails",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {
				kindNodeControlPlane := kindmocks.NewKindNode(t)
				kindNodeControlPlane.EXPECT().Role().Return("control-plane", nil)
				kindNodeControlPlane.EXPECT().String().Return(fmt.Sprintf("%s-%s", clusterName, "control-plane"))

				kindNodeWorker := kindmocks.NewKindNode(t)
				kindNodeWorker.EXPECT().Role().Return("worker", nil)

				m.EXPECT().List().Return([]string{clusterName}, nil)
				m.EXPECT().ListNodes(clusterName).Return([]nodes.Node{kindNodeControlPlane, kindNodeWorker}, nil)
			},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().
					RunContainer(
						context.Background(), registryContainerConfig, registryContainerHostConfig, &network.NetworkingConfig{}, registryName, false, mock.Anything,
					).Return(true, nil)
				m.EXPECT().
					RunContainer(
						context.Background(), mock.Anything, mock.Anything, mock.Anything, fmt.Sprintf("%s-control-plane", clusterName), false, mock.Anything,
					).Return(false, errors.New("error running container"))
			},
			applyK8sCLIMockCalls:     func(m *wrappermocks.K8sCLI) {},
			applyKubeConfigMockCalls: func(m *wrappermocks.KubeConfigInterface) {},
			expectedErr:              "error running container",
		},
		{
			name: "FailsToCreateRegistry",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {
			},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().
					RunContainer(
						context.Background(), registryContainerConfig, registryContainerHostConfig, &network.NetworkingConfig{}, registryName, false, mock.Anything,
					).Return(false, errors.New("error creating registry"))
			},
			applyK8sCLIMockCalls:     func(m *wrappermocks.K8sCLI) {},
			applyKubeConfigMockCalls: func(m *wrappermocks.KubeConfigInterface) {},
			expectedErr:              "error creating registry",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockKindProvider := wrappermocks.NewKindProvider(t)
			tt.applyKindProviderMockCalls(mockKindProvider)
			mockDockerCLI := wrappermocks.NewDockerCLI(t)
			tt.applyDockerCLIMockCalls(mockDockerCLI)
			mockK8sCLI := wrappermocks.NewK8sCLI(t)
			tt.applyK8sCLIMockCalls(mockK8sCLI)
			mockKubeConfig := wrappermocks.NewKubeConfigInterface(t)
			tt.applyKubeConfigMockCalls(mockKubeConfig)

			kindCluster := wrappers.NewKindCluster(clusterName, &v1alpha4.Cluster{}, mockDockerCLI, mockKindProvider, kubeConfigFile, registryName, registryContainerConfig, registryContainerHostConfig)
			kindCluster.SetKubeConfig(mockKubeConfig)

			err := kindCluster.CreateOrReuse(namespaceName, mockK8sCLI)
			if tt.expectedErr == "" {
				require.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.expectedErr)
			}
		})
	}
}

func TestNewKindCluster_Delete(t *testing.T) {
	t.Parallel()

	clusterName := "test-cluster"
	registryName := "test-registry"

	testCases := []struct {
		name                       string
		applyKindProviderMockCalls func(m *wrappermocks.KindProvider)
		applyDockerCLIMockCalls    func(m *wrappermocks.DockerCLI)
		expectedErr                string
	}{
		{
			name: "DeleteSucceeds",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {
				m.EXPECT().Delete(clusterName, mock.AnythingOfType("string")).Return(nil)
			},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().DeleteContainer(context.TODO(), registryName, containertypes.RemoveOptions{RemoveVolumes: true, Force: true}, mock.Anything).Return(nil)
			},
			expectedErr: "",
		},
		{
			name: "DeleteFails",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {
				m.EXPECT().Delete(clusterName, mock.AnythingOfType("string")).Return(errors.New("error deleting cluster"))
			},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().DeleteContainer(context.TODO(), registryName, containertypes.RemoveOptions{RemoveVolumes: true, Force: true}, mock.Anything).Return(nil)
			},
			expectedErr: "error deleting cluster",
		},
		{
			name:                       "DeleteFailsToRemoveRegistry",
			applyKindProviderMockCalls: func(m *wrappermocks.KindProvider) {},
			applyDockerCLIMockCalls: func(m *wrappermocks.DockerCLI) {
				m.EXPECT().DeleteContainer(context.TODO(), registryName, containertypes.RemoveOptions{RemoveVolumes: true, Force: true}, mock.Anything).Return(errors.New("error removing registry"))
			},
			expectedErr: "error removing registry",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockKindProvider := wrappermocks.NewKindProvider(t)
			tt.applyKindProviderMockCalls(mockKindProvider)
			mockDockerCLI := wrappermocks.NewDockerCLI(t)
			tt.applyDockerCLIMockCalls(mockDockerCLI)

			kindCluster := wrappers.NewKindCluster(clusterName, nil, mockDockerCLI, mockKindProvider, string(mock.AnythingOfType("string")), registryName, nil, nil)
			err := kindCluster.Delete()
			if tt.expectedErr == "" {
				require.NoError(t, err)
			} else {
				assert.ErrorContains(t, err, tt.expectedErr)
			}
		})
	}
}
