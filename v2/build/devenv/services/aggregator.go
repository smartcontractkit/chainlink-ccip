package services

import (
	"context"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"
)

const (
	DefaultAggregatorName    = "aggregator"
	DefaultAggregatorDBName  = "aggregator-db"
	DefaultAggregatorImage   = "aggregator:dev"
	DefaultAggregatorPort    = 8103
	DefaultAggregatorDBPort  = 7432
	DefaultAggregatorSQLInit = "init.sql"

	DefaultAggregatorDBImage = "postgres:16-alpine"
)

var (
	DefaultAggregatorDBConnectionString = fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable",
		DefaultAggregatorName, DefaultAggregatorName, DefaultAggregatorDBPort, DefaultAggregatorName)
)

type AggregatorDBInput struct {
	Image string `toml:"image"`
}

type AggregatorInput struct {
	Image          string            `toml:"image"`
	Port           int               `toml:"port"`
	SourceCodePath string            `toml:"source_code_path"`
	DB             *DBInput          `toml:"db"`
	ContainerName  string            `toml:"container_name"`
	UseCache       bool              `toml:"use_cache"`
	Out            *AggregatorOutput `toml:"-"`
}

type AggregatorOutput struct {
	UseCache           bool   `toml:"use_cache"`
	ContainerName      string `toml:"container_name"`
	ExternalHTTPURL    string `toml:"http_url"`
	InternalHTTPURL    string `toml:"internal_http_url"`
	DBURL              string `toml:"db_url"`
	DBConnectionString string `toml:"db_connection_string"`
}

func aggregatorDefaults(in *AggregatorInput) {
	if in.Image == "" {
		in.Image = DefaultAggregatorImage
	}
	if in.Port == 0 {
		in.Port = DefaultAggregatorPort
	}
	if in.ContainerName == "" {
		in.ContainerName = DefaultAggregatorName
	}
	if in.DB == nil {
		in.DB = &DBInput{
			Image: DefaultAggregatorDBImage,
		}
	}
}

func NewAggregator(in *AggregatorInput) (*AggregatorOutput, error) {
	if in.Out != nil && in.Out.UseCache {
		return in.Out, nil
	}
	ctx := context.Background()
	aggregatorDefaults(in)
	p, err := CwdSourcePath(in.SourceCodePath)
	if err != nil {
		return in.Out, err
	}

	/* Database */
	_, err = postgres.Run(ctx,
		in.DB.Image,
		testcontainers.WithName(DefaultAggregatorDBName),
		testcontainers.WithExposedPorts("5432/tcp"),
		testcontainers.WithHostConfigModifier(func(h *container.HostConfig) {
			h.PortBindings = nat.PortMap{
				"5432/tcp": []nat.PortBinding{
					{HostPort: strconv.Itoa(DefaultAggregatorDBPort)},
				},
			}
		}),
		testcontainers.WithLabels(framework.DefaultTCLabels()),
		postgres.WithDatabase(DefaultAggregatorName),
		postgres.WithUsername(DefaultAggregatorName),
		postgres.WithPassword(DefaultAggregatorName),
		postgres.WithInitScripts(filepath.Join(p, DefaultAggregatorSQLInit)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
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
		// add more internal ports here with /tcp suffix, ex.: 9222/tcp
		ExposedPorts: []string{"8100/tcp"},
		HostConfigModifier: func(h *container.HostConfig) {
			h.PortBindings = nat.PortMap{
				// add more internal/external pairs here, ex.: 9222/tcp as a key and HostPort is the exposed port (no /tcp prefix!)
				"8100/tcp": []nat.PortBinding{
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

	return &AggregatorOutput{
		ContainerName:      in.ContainerName,
		ExternalHTTPURL:    fmt.Sprintf("http://%s:%d", host, in.Port),
		InternalHTTPURL:    fmt.Sprintf("http://%s:%d", in.ContainerName, in.Port),
		DBConnectionString: DefaultAggregatorDBConnectionString,
	}, nil
}
