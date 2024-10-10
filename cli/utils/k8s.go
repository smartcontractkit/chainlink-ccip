package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/smartcontractkit/crib/cli/wrappers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

// getKubeConfigFromFile tries to read a kubeconfig file and if it can't, returns an error. Missing files result in empty configs, not an error
func getKubeConfigFromFile(filename string) (*clientcmdapi.Config, error) {
	config, err := clientcmd.LoadFromFile(filename)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if config == nil {
		config = clientcmdapi.NewConfig()
	}
	return config, nil
}

type SetupKubeConfigInput struct {
	EksClient            wrappers.EKSAPI
	KubeconfigPath       string
	EksClusterName       string
	EksAliasName         string
	CribNamespace        string
	AwsProfile           string
	AwsRegion            string
	ChangeDefaultContext bool
}

// SetupKubeConfig produces a new kubeconfig for accessing a given EKS cluster under the context named after eksAliasName.
// If kubeconfigPath points at a non-existing file, it'll get created. If the file exists, it'll attempt to parse it
// and modify the respective cluster, context and user entries.
func SetupKubeConfig(input *SetupKubeConfigInput) error {
	eksCluster, err := input.EksClient.DescribeCluster(context.TODO(), &eks.DescribeClusterInput{
		Name: &input.EksClusterName,
	})
	if err != nil {
		return fmt.Errorf("unable to fetch EKS cluster info, %v", err)
	}

	newConfig, err := getKubeConfigFromFile(input.KubeconfigPath)
	if err != nil {
		return fmt.Errorf("unable to parse kube config, %v", err)
	}

	// modify kubeconfig
	eksClusterArn := *eksCluster.Cluster.Arn
	eksClusterEndpoint := *eksCluster.Cluster.Endpoint

	// eks.DescribeCluster's output returns base64-encoded
	// CAData, but clientcmdapi.Cluster expects it decoded
	// see: https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/eks@v1.49.2/types#Certificate
	decodedEksClusterCAData, err := base64.StdEncoding.DecodeString(*eksCluster.Cluster.CertificateAuthority.Data)
	if err != nil {
		return fmt.Errorf("failed to decode base64 data from eks.DescribeCluster output, %v", err)
	}

	newConfig.Clusters[eksClusterArn] = &clientcmdapi.Cluster{
		Server:                   eksClusterEndpoint,
		CertificateAuthorityData: decodedEksClusterCAData,
	}

	// clientcmdapi.ExecConfig based on current state of aws eks update-kubeconfig
	// see: https://github.com/aws/aws-cli/blob/497a62cd38df982eb8dd3c06db447fb534cea009/awscli/customizations/eks/update_kubeconfig.py#L308-L327
	newConfig.AuthInfos[eksClusterArn] = &clientcmdapi.AuthInfo{
		Exec: &clientcmdapi.ExecConfig{
			APIVersion: "client.authentication.k8s.io/v1beta1",
			Env: []clientcmdapi.ExecEnvVar{
				{Name: "AWS_PROFILE", Value: input.AwsProfile},
			},
			Command:            "aws",
			Args:               []string{"--region", input.AwsRegion, "eks", "get-token", "--cluster-name", input.EksClusterName, "--output", "json"},
			InteractiveMode:    "IfAvailable",
			ProvideClusterInfo: false,
		},
	}

	newConfig.Contexts[input.EksAliasName] = &clientcmdapi.Context{
		Cluster:   eksClusterArn,
		AuthInfo:  eksClusterArn,
		Namespace: input.CribNamespace,
	}

	if input.ChangeDefaultContext {
		// Set the current context to the one we just created
		newConfig.CurrentContext = input.EksAliasName
	}

	// Write the produced config into kubeconfigPath
	pathOptions := clientcmd.NewDefaultPathOptions()
	pathOptions.GlobalFile = input.KubeconfigPath
	return clientcmd.ModifyConfig(pathOptions, *newConfig, true)
}

func CheckEksAccess(kubeCoreV1Client corev1.CoreV1Interface) error {
	_, err := kubeCoreV1Client.Namespaces().List(context.TODO(), metav1.ListOptions{})
	return err
}
