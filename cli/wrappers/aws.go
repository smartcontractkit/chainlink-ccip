package wrappers

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

type EKSAPI interface {
	DescribeCluster(ctx context.Context, input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error)
}

type EKSClient struct {
	client *eks.Client
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
