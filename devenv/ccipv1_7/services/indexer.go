package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"
)

const (
	DefaultName = "example-service"
	// TODO: insert real image here
	DefaultImage    = "f4hrenh9it/df-fakes:latest"
	DefaultHTTPPort = 8331

	DefaultDBImage            = "postgres:16-alpine"
	DefaultDBConnectionString = "postgresql://indexer:indexer@localhost:6432/indexer?sslmode=disable"
)

type DBInput struct {
	Image string `toml:"image"`
}

type IndexerInput struct {
	Image         string         `toml:"image"`
	Port          int            `toml:"port"`
	DB            *DBInput       `toml:"db"`
	ExposedPorts  []string       `toml:"exposed_ports"`
	ContainerName string         `toml:"container_name"`
	UseCache      bool           `toml:"use_cache"`
	Out           *IndexerOutput `toml:"-"`
}

type IndexerOutput struct {
	UseCache           bool   `toml:"use_cache"`
	ContainerName      string `toml:"container_name"`
	ExternalHTTPURL    string `toml:"http_url"`
	InternalHTTPURL    string `toml:"internal_http_url"`
	DBURL              string `toml:"db_url"`
	DBConnectionString string `toml:"db_connection_string"`
}

func defaults(in *IndexerInput) {
	if in.Image == "" {
		in.Image = DefaultImage
	}
	if in.Port == 0 {
		in.Port = DefaultHTTPPort
	}
	if in.ContainerName == "" {
		in.ContainerName = DefaultName
	}
	if in.DB == nil {
		in.DB = &DBInput{
			Image: DefaultDBImage,
		}
	}
}

// NewIndexer creates and starts a new Service container using testcontainers
func NewIndexer(in *IndexerInput) (*IndexerOutput, error) {
	if in.Out != nil && in.Out.UseCache {
		return in.Out, nil
	}
	ctx := context.Background()

	defaults(in)

	/* Database */

	_, err := postgres.Run(ctx,
		in.DB.Image,
		testcontainers.WithExposedPorts("5432/tcp"),
		testcontainers.WithHostConfigModifier(func(h *container.HostConfig) {
			h.PortBindings = nat.PortMap{
				"5432/tcp": []nat.PortBinding{
					{HostPort: "6432"},
				},
			}
		}),
		testcontainers.WithLabels(framework.DefaultTCLabels()),
		postgres.WithDatabase("indexer"),
		postgres.WithUsername("indexer"),
		postgres.WithPassword("indexer"),
		postgres.WithInitScripts("services/init.sql"),
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
		ExposedPorts: []string{"9111/tcp"},
		HostConfigModifier: func(h *container.HostConfig) {
			h.PortBindings = nat.PortMap{
				// add more internal/external pairs here, ex.: 9222/tcp as a key and HostPort is the exposed port (no /tcp prefix!)
				"9111/tcp": []nat.PortBinding{
					{HostPort: strconv.Itoa(in.Port)},
				},
			}
		},
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

	return &IndexerOutput{
		ContainerName:      in.ContainerName,
		ExternalHTTPURL:    fmt.Sprintf("http://%s:%d", host, in.Port),
		InternalHTTPURL:    fmt.Sprintf("http://%s:%d", in.ContainerName, in.Port),
		DBConnectionString: DefaultDBConnectionString,
	}, nil
}
