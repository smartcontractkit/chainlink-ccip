package wrappers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/eks"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

type ECRAPI interface {
	GetAuthorizationToken(ctx context.Context, params *ecr.GetAuthorizationTokenInput, optFns ...func(*ecr.Options)) (*ecr.GetAuthorizationTokenOutput, error)
}

type EKSAPI interface {
	DescribeCluster(ctx context.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error)
}

type STSAPI interface {
	GetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error)
}

type ECRClient struct {
	client *ecr.Client
}

type EKSClient struct {
	client *eks.Client
}

type STSClient struct {
	client *sts.Client
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

func NewSTSClientWrapper(config aws.Config) *STSClient {
	return &STSClient{
		client: sts.NewFromConfig(config),
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

// Implement GetCallerIdentity method, fulfilling the STSAPI interface.
func (c *STSClient) GetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error) {
	return c.client.GetCallerIdentity(ctx, params, optFns...)
}
