package wrappers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type ECRAPI interface {
	GetAuthorizationToken(ctx context.Context, params *ecr.GetAuthorizationTokenInput, optFns ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error)
}

type EKSAPI interface {
	DescribeCluster(ctx context.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error)
}

type ECRClient struct {
	client *ecr.Client
}

type EKSClient struct {
	client *eks.Client
}

func NewECRClientWrapper(config aws.Config) *ECRClient {
	return &ECRClient{
		client: ecr.NewFromConfig(config),
	}
}

func NewEKSClientWrapper(config aws.Config) *EKSClient {
	return &EKSClient{
		client: eks.NewFromConfig(config),
	}
}

// Implement DescribeCluster method, fulfilling the EKSAPI interface.
func (c *EKSClient) DescribeCluster(ctx context.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error) {
	return c.client.DescribeCluster(ctx, input)
}

// Implement GetAuthorizationToken method, fulfilling the ECRAPI interface.
func (c *ECRClient) GetAuthorizationToken(ctx context.Context, params *ecr.GetAuthorizationTokenInput, optFns ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error) {
	return c.client.GetAuthorizationToken(ctx, params, optFns...)
}
