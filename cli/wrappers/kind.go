package wrappers

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"slices"

	containertypes "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/go-connections/nat"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
	"sigs.k8s.io/kind/pkg/cluster"
	"sigs.k8s.io/kind/pkg/cluster/nodes"
)

type registryConfig struct {
	containerName   string
	containerConfig *containertypes.Config
	hostConfig      *containertypes.HostConfig
}

// KindProvider is used to perform cluster operations
type KindProvider interface {
	Create(name string, options ...cluster.CreateOption) error
	Delete(name, explicitKubeconfigPath string) error
	ExportKubeConfig(name string, explicitPath string, internal bool) error
	List() ([]string, error)
	ListNodes(name string) ([]nodes.Node, error)
}

type KindCluster struct {
	name           string
	config         *v1alpha4.Cluster
	dockerCli      DockerCLI
	k8sClient      K8sCLI
	provider       KindProvider
	kubeconfigPath string
	registryConfig *registryConfig
}

const (
	DefaultClusterName  = "crib-cluster"
	DefaultRegistryName = "kind-registry"
	DefaultRegistryPort = "5001"
)

var DefaultClusterConfig = &v1alpha4.Cluster{
	ContainerdConfigPatches: []string{
		`[plugins."io.containerd.grpc.v1.cri".registry]
config_path = "/etc/containerd/certs.d"`,
	},
	Nodes: []v1alpha4.Node{
		{
			Role: "control-plane",
			KubeadmConfigPatches: []string{
				`kind: InitConfiguration
nodeRegistration:
kubeletExtraArgs:
node-labels: "ingress-ready=true"`,
			},
			ExtraPortMappings: []v1alpha4.PortMapping{
				{
					ContainerPort: 80,
					HostPort:      80,
					Protocol:      v1alpha4.PortMappingProtocolTCP,
				},
				{
					ContainerPort: 443,
					HostPort:      443,
					Protocol:      v1alpha4.PortMappingProtocolTCP,
				},
			},
		},
		{
			Role: "worker",
		},
	},
}

func NewKindCluster(name string, config *v1alpha4.Cluster, dockerCli DockerCLI, provider KindProvider, kubeconfigPath string, registryName string, registryContainerConfig *containertypes.Config, registryContainerHostConfig *containertypes.HostConfig) *KindCluster {
	if name == "" {
		name = DefaultClusterName
	}

	if config == nil {
		config = DefaultClusterConfig
	}

	if provider == nil {
		// TODO: cluster.ProviderWithLogger(logger)
		provider = cluster.NewProvider(cluster.ProviderWithDocker())
	}

	if kubeconfigPath == "" {
		val, found := os.LookupEnv("KUBECONFIG")
		kubeconfigPath = val
		if !found {
			userHomeDir, _ := os.UserHomeDir()
			kubeconfigPath = filepath.Join(userHomeDir, ".kube", "config")
		}
	}

	if registryName == "" {
		registryName = DefaultRegistryName
	}

	if registryContainerConfig == nil {
		registryContainerConfig = &containertypes.Config{
			Image: "registry:2",
			ExposedPorts: nat.PortSet{
				"5000/tcp": struct{}{},
			},
		}
	}

	if registryContainerHostConfig == nil {
		containerPort, _ := nat.NewPort("tcp", "5000")
		portBindings := nat.PortMap{
			containerPort: []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: DefaultRegistryPort}},
		}
		registryContainerHostConfig = &containertypes.HostConfig{
			PortBindings: portBindings,
			RestartPolicy: containertypes.RestartPolicy{
				Name: "always",
			},
			NetworkMode: "bridge",
		}
	}

	return &KindCluster{
		name:           name,
		config:         config,
		dockerCli:      dockerCli,
		provider:       provider,
		kubeconfigPath: kubeconfigPath,
		registryConfig: &registryConfig{
			containerName:   registryName,
			containerConfig: registryContainerConfig,
			hostConfig:      registryContainerHostConfig,
		},
	}
}

// CreateOrReuse creates a new kind cluster or reuses an existing one if available
// It also sets up a local docker registry and configures the cluster to use it
// You don't need to pass k8sClient normally, but it's useful for testing
func (k *KindCluster) CreateOrReuse(k8sClient K8sCLI) error {
	registryAlreadyRunning, err := k.createDockerRegistry(false)
	if err != nil {
		return err
	}
	slog.Info("docker registry in place for kind", slog.String("name", k.registryConfig.containerName), slog.Bool("already_existed", registryAlreadyRunning))

	clusterAlreadyExists, err := k.clusterAlreadyExists()
	if err != nil {
		return err
	}
	if clusterAlreadyExists {
		if err := k.ensureExistingClusterStarted(); err != nil {
			return err
		}
	} else {
		if err := k.provider.Create(k.name, cluster.CreateWithKubeconfigPath(k.kubeconfigPath), cluster.CreateWithV1Alpha4Config(k.config)); err != nil {
			return err
		}
	}

	kindContext := fmt.Sprintf("kind-%s", k.name)
	if err := k.configureKubectlContext(); err != nil {
		return err
	}
	slog.Info("kubeconfig updated for kind", slog.String("context", kindContext), slog.String("kubeconfig", k.kubeconfigPath))

	if k8sClient == nil {
		// instantiate k8sClient after the cluster is in place
		k8sConfigFlags := genericclioptions.NewConfigFlags(true)
		k8sConfigFlags.KubeConfig = &k.kubeconfigPath
		k8sClient, err = NewK8sClient(k8sConfigFlags, nil)
		if err != nil {
			return err
		}
	}
	k.k8sClient = k8sClient
	slog.Info("kind cluster in place", slog.String("name", k.name), slog.Bool("already_existed", clusterAlreadyExists))

	if err := k.configureRegistryOnNodes(); err != nil {
		return err
	}
	slog.Info("registry configured on the kind cluster nodes successfully", slog.String("cluster", k.name), slog.String("registry", k.registryConfig.containerName))

	// TODO: move the network name to a proper config
	networkName := "kind"
	alreadyConnectedToNetwork, err := k.connectRegistryToNetwork(networkName)
	if err != nil {
		return err
	}
	slog.Info("registry connected to the network", slog.String("registry", k.registryConfig.containerName), slog.String("network", networkName), slog.Bool("already_connected", alreadyConnectedToNetwork))

	alreadyExistingConfigMap, err := k.documentLocalRegistry()
	if err != nil {
		return err
	}
	slog.Info("local registry documented in Kubernetes successfully", slog.String("cluster", k.name), slog.String("registry", k.registryConfig.containerName), slog.Bool("already_existed", alreadyExistingConfigMap))

	return nil
}

func (k *KindCluster) Delete() error {
	if err := k.deleteDockerRegistry(true, true); err != nil {
		return err
	}
	slog.Info("docker registry deleted", slog.String("name", k.registryConfig.containerName))
	return k.provider.Delete(k.name, k.kubeconfigPath)
}

func (k *KindCluster) Name() string {
	return k.name
}

func (k *KindCluster) K8sClient() K8sCLI {
	return k.k8sClient
}

func (k *KindCluster) createDockerRegistry(forceRecreate bool) (bool, error) {
	registryAlreadyRunning, err := k.dockerCli.RunContainer(context.Background(), k.registryConfig.containerConfig, k.registryConfig.hostConfig, &network.NetworkingConfig{}, k.registryConfig.containerName, forceRecreate, nil)
	if err != nil {
		return registryAlreadyRunning, fmt.Errorf("error running container %s: %w", k.registryConfig.containerName, err)
	}
	return registryAlreadyRunning, nil
}

func (k *KindCluster) clusterAlreadyExists() (bool, error) {
	clusters, err := k.provider.List()
	if err != nil {
		return false, err
	}
	return slices.Contains(clusters, k.name), nil
}

func (k *KindCluster) ensureExistingClusterStarted() error {
	nodeList, err := k.provider.ListNodes(k.name)
	if err != nil {
		return fmt.Errorf("failed to list kind nodes for cluster %s: %w", k.name, err)
	}

	var controlPlane nodes.Node
	var workers []nodes.Node

	for _, node := range nodeList {
		role, err := node.Role()
		if err != nil {
			return fmt.Errorf("failed to get role for node %s: %w", node.String(), err)
		}
		if role == "control-plane" {
			controlPlane = node
		} else if role == "worker" {
			workers = append(workers, node)
		}
	}

	if controlPlane != nil {
		if err := k.ensureNodeRunning(controlPlane); err != nil {
			return fmt.Errorf("failed to start control plane node %s: %w", controlPlane.String(), err)
		}
	} else {
		return fmt.Errorf("no control plane node found; cannot start cluster")
	}

	for _, worker := range workers {
		if err := k.ensureNodeRunning(worker); err != nil {
			return fmt.Errorf("failed to start worker node %s: %w", worker.String(), err)
		}
	}

	return nil
}

func (k *KindCluster) ensureNodeRunning(node nodes.Node) error {
	_, err := k.dockerCli.RunContainer(context.Background(), nil, nil, nil, node.String(), false, nil)
	if err != nil {
		return err
	}

	return nil
}

func (k *KindCluster) configureKubectlContext() error {
	return k.provider.ExportKubeConfig(k.name, k.kubeconfigPath, false)
}

func (k *KindCluster) configureRegistryOnNodes() error {
	nodes, err := k.provider.ListNodes(k.name)
	if err != nil {
		return err
	}

	baseCertsDir := "/etc/containerd/certs.d"
	for port, bindings := range k.registryConfig.hostConfig.PortBindings {
		for _, binding := range bindings {
			hostIP := binding.HostIP
			if hostIP == "127.0.0.1" {
				hostIP = "localhost"
			}

			registryCertsDir := fmt.Sprintf("%s/%s:%s", baseCertsDir, hostIP, binding.HostPort)
			for _, node := range nodes {
				if err := node.Command("mkdir", "-p", registryCertsDir).Run(); err != nil {
					return err
				}
				hostsTomlContent := fmt.Sprintf("[host.\"http://%s:%s\"]\n", k.registryConfig.containerName, port.Port())
				hostsTomlPathInsideNode := fmt.Sprintf("%s/hosts.toml", registryCertsDir)
				if err := node.Command("echo", hostsTomlContent, ">", hostsTomlPathInsideNode).Run(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (k *KindCluster) connectRegistryToNetwork(networkName string) (bool, error) {
	if networkName == "" {
		networkName = "kind"
	}

	alreadyConnectedToNetwork, err := k.dockerCli.ConnectContainerToNetwork(context.Background(), k.registryConfig.containerName, networkName)
	if err != nil {
		return alreadyConnectedToNetwork, err
	}

	return alreadyConnectedToNetwork, nil
}

func (k *KindCluster) documentLocalRegistry() (bool, error) {
	// TODO: read the actual port bindings in order to generate this
	registryHost := fmt.Sprintf("localhost:%s", DefaultRegistryPort)

	configMap := &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "local-registry-hosting",
			Namespace: "kube-public",
		},
		Data: map[string]string{
			"localRegistryHosting.v1": fmt.Sprintf(`host: "%s"\nhelp: "https://kind.sigs.k8s.io/docs/user/local-registry/"`, registryHost),
		},
	}

	alreadyExistingConfigMap, err := k.k8sClient.ApplyConfigMap(context.TODO(), configMap)
	if err != nil {
		return alreadyExistingConfigMap, err
	}

	return alreadyExistingConfigMap, nil
}

func (k *KindCluster) deleteDockerRegistry(force bool, removeVolumes bool) error {
	return k.dockerCli.DeleteContainer(context.TODO(), k.registryConfig.containerName, containertypes.RemoveOptions{
		Force:         force,
		RemoveVolumes: removeVolumes,
	}, nil)
}
