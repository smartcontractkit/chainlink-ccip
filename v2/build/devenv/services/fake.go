package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"
)

const (
	DefaultFakeName  = "fake"
	DefaultFakeImage = "ccv-fakes:dev"
	DefaultFakePort  = 9111
)

type FakeInput struct {
	Image          string      `toml:"image"`
	Port           int         `toml:"port"`
	SourceCodePath string      `toml:"source_code_path"`
	ContainerName  string      `toml:"container_name"`
	UseCache       bool        `toml:"use_cache"`
	Out            *FakeOutput `toml:"-"`
}

type FakeOutput struct {
	UseCache        bool   `toml:"use_cache"`
	ContainerName   string `toml:"container_name"`
	ExternalHTTPURL string `toml:"http_url"`
	InternalHTTPURL string `toml:"internal_http_url"`
}

func fakeDefaults(in *FakeInput) {
	if in.Image == "" {
		in.Image = DefaultFakeImage
	}
	if in.ContainerName == "" {
		in.ContainerName = DefaultFakeName
	}
}

func NewFake(in *FakeInput) (*FakeOutput, error) {
	if in.Out != nil && in.Out.UseCache {
		return in.Out, nil
	}
	ctx := context.Background()

	fakeDefaults(in)
	p, err := CwdSourcePath(in.SourceCodePath)
	if err != nil {
		return in.Out, nil
	}

	/* Service */
	req := testcontainers.ContainerRequest{
		Image:    in.Image,
		Name:     in.ContainerName,
		Labels:   framework.DefaultTCLabels(),
		Networks: []string{framework.DefaultNetworkName},
		NetworkAliases: map[string][]string{
			framework.DefaultNetworkName: {in.ContainerName},
		},
		ExposedPorts: []string{"9111/tcp"},
		HostConfigModifier: func(h *container.HostConfig) {
			h.PortBindings = nat.PortMap{
				"9111/tcp": []nat.PortBinding{
					{HostPort: strconv.Itoa(in.Port)},
				},
			}
		},
	}

	if in.SourceCodePath != "" {
		req.Mounts = GoSourcePathMounts(p, AppPathInsideContainer)
		framework.L.Info().
			Str("Service", in.ContainerName).
			Str("Source", p).Msg("Using source code path, hot-reload mode")
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start container: %w", err)
	}
	host, err := c.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get container host: %w", err)
	}

	return &FakeOutput{
		ContainerName:   in.ContainerName,
		ExternalHTTPURL: fmt.Sprintf("http://%s:%d", host, in.Port),
		InternalHTTPURL: fmt.Sprintf("http://%s:%d", in.ContainerName, in.Port),
	}, nil
}
