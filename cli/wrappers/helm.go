package wrappers

import (
	"helm.sh/helm/v3/pkg/registry"
)

type HelmRegistryAPI interface {
	Login(host string, options ...registry.LoginOption) error
}

type HelmRegistryClient struct {
	cli *registry.Client
}

func NewHelmRegistryClientWrapper(opts ...registry.ClientOption) (*HelmRegistryClient, error) {
	helmRegistryCli, err := registry.NewClient(opts...)
	if err != nil {
		return nil, err
	}
	return &HelmRegistryClient{
		cli: helmRegistryCli,
	}, nil
}

func (c *HelmRegistryClient) Login(host string, options ...registry.LoginOption) error {
	return c.cli.Login(host, options...)
}
