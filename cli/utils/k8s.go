package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/smartcontractkit/crib/cli/wrappers"
	"k8s.io/client-go/tools/clientcmd"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// getKubeConfigFromFile tries to read a kubeconfig file and if it can't, returns an error. Missing files result in empty configs, not an error
func getKubeConfigFromFile(filename string) (*clientcmdapi.Config, error) {
	config, err := clientcmd.LoadFromFile(filename)
	// TODO: os.IsPermission()?
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if config == nil {
		config = clientcmdapi.NewConfig()
	}
	return config, nil
}

// SetupKubeConfig produces a new kubeconfig for accessing a given EKS cluster under the context named after eksAliasName.
// If kubeconfigPath points at a non-existing file, it'll get created. If the file exists, it'll attempt to parse it
// and modify the respective cluster, context and user entries.
func SetupKubeConfig(eksClient wrappers.EKSAPI, kubeconfigPath string, eksClusterName string, eksAliasName string, awsRegion string, changeDefaultContext bool) error {
	eksCluster, err := eksClient.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{
		Name: &eksClusterName,
	})
	if err != nil {
		return fmt.Errorf("unable to fetch EKS cluster info, %v", err)
	}

	newConfig, err := getKubeConfigFromFile(kubeconfigPath)
	if err != nil {
		return fmt.Errorf("unable to retrieve kube config, %v", err)
	}

	// modify kubeconfig
	eksClusterArn := *eksCluster.Cluster.Arn
	eksClusterEndpoint := *eksCluster.Cluster.Endpoint
	eksClusterCAData := []byte(*eksCluster.Cluster.CertificateAuthority.Data)

	newConfig.Clusters[eksClusterArn] = &clientcmdapi.Cluster{
		Server:                   eksClusterEndpoint,
		CertificateAuthorityData: eksClusterCAData,
	}

	newConfig.AuthInfos[eksClusterArn] = &clientcmdapi.AuthInfo{
		Exec: &clientcmdapi.ExecConfig{
			APIVersion: "client.authentication.k8s.io/v1beta1",
			Command:    "aws",
			Args:       []string{"eks", "get-token", "--cluster-name", eksClusterName},
		},
	}

	newConfig.Contexts[eksAliasName] = &clientcmdapi.Context{
		Cluster:  eksClusterArn,
		AuthInfo: eksClusterArn,
	}

	if changeDefaultContext {
		// Set the current context to the one we just created
		newConfig.CurrentContext = eksAliasName
	}

	// Write the produced config into kubeconfigPath
	pathOptions := clientcmd.NewDefaultPathOptions()
	pathOptions.GlobalFile = kubeconfigPath
	return clientcmd.ModifyConfig(pathOptions, *newConfig, true)
}
