package wrappers

import (
	"context"
	"fmt"
	"time"

	corev1api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	corev1typed "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd/api"
)

type K8sCLI interface {
	Clientset() K8sClientset
	ClientConfig() *api.Config
	RestConfig() *rest.Config
	CurrentContext() string
	CheckAccess(ctx context.Context) error
	EnsureNamespaceExists(ctx context.Context, name string) (bool, error)
	WaitForResource(ctx context.Context, resourceClient dynamic.ResourceInterface, resourceName string, interval, timeout time.Duration) error
}

type K8sClientset interface {
	CoreV1() corev1typed.CoreV1Interface
}

type K8sClient struct {
	clientGetter genericclioptions.RESTClientGetter
	clientset    K8sClientset
	clientConfig *api.Config
	restConfig   *rest.Config
}

// NewK8sClient creates a new K8sClient instance. If clientset is nil, a new clientset is created using the provided clientGetter.
// clientset doesn't have to be passed normally, but it's useful for testing.
func NewK8sClient(clientGetter genericclioptions.RESTClientGetter, clientset K8sClientset) (*K8sClient, error) {
	clientConfig, err := clientGetter.ToRawKubeConfigLoader().RawConfig()
	if err != nil {
		return nil, err
	}

	restConfig, err := clientGetter.ToRESTConfig()
	if err != nil {
		return nil, err
	}

	if clientset == nil {
		clientset, err = kubernetes.NewForConfig(restConfig)
		if err != nil {
			return nil, err
		}
	}

	return &K8sClient{
		clientGetter: clientGetter,
		clientset:    clientset,
		clientConfig: &clientConfig,
		restConfig:   restConfig,
	}, nil
}

func (k *K8sClient) Clientset() K8sClientset {
	return k.clientset
}

func (k *K8sClient) ClientConfig() *api.Config {
	return k.clientConfig
}

func (k *K8sClient) RestConfig() *rest.Config {
	return k.restConfig
}

func (k *K8sClient) CurrentContext() string {
	return k.clientConfig.CurrentContext
}

func (k *K8sClient) CheckAccess(ctx context.Context) error {
	_, err := k.clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	return err
}

// EnsureNamespaceExists creates a namespace if it doesn't already exist.
// Returns a boolean indicating if the namespace already existed.
func (k *K8sClient) EnsureNamespaceExists(ctx context.Context, name string) (bool, error) {
	nsClient := k.clientset.CoreV1().Namespaces()
	_, err := nsClient.Get(ctx, name, metav1.GetOptions{})
	if err == nil || errors.IsAlreadyExists(err) {
		return true, nil
	}
	if err != nil && !errors.IsNotFound(err) {
		return false, fmt.Errorf("failed to get namespace %s: %w", name, err)
	}

	newNamespace := &corev1api.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	_, err = nsClient.Create(ctx, newNamespace, metav1.CreateOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to create namespace %s: %w", name, err)
	}

	return false, nil
}

// WaitForResource waits for a specified resource to be created in a namespace.
// It periodically checks for the resource until the timeout is reached.
func (k *K8sClient) WaitForResource(ctx context.Context, resourceClient dynamic.ResourceInterface, resourceName string, interval, timeout time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	timeoutCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			return fmt.Errorf("timed out waiting for resource %s", resourceName)

		case <-ticker.C:
			_, err := resourceClient.Get(ctx, resourceName, metav1.GetOptions{})
			if err == nil {
				return nil
			}
		}
	}
}
