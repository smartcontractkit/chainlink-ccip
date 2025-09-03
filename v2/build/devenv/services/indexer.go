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
	DefaultIndexerName     = "indexer"
	DefaultIndexerDBName   = "indexer-db"
	DefaultIndexerImage    = "indexer:dev"
	DefaultIndexerHTTPPort = 8102
	DefaultIndexerDBPort   = 6432
	DefaultIndexerSQLInit  = "init.sql"

	DefaultIndexerDBImage = "postgres:16-alpine"
)

var (
	DefaultIndexerDBConnectionString = fmt.Sprintf("postgresql://%s:%s@localhost:%d/%s?sslmode=disable",
		DefaultIndexerName, DefaultIndexerName, DefaultIndexerDBPort, DefaultIndexerName)
)

type DBInput struct {
	Image string `toml:"image"`
}

type IndexerInput struct {
	Image          string         `toml:"image"`
	Port           int            `toml:"port"`
	SourceCodePath string         `toml:"source_code_path"`
	DB             *DBInput       `toml:"db"`
	ContainerName  string         `toml:"container_name"`
	UseCache       bool           `toml:"use_cache"`
	Out            *IndexerOutput `toml:"-"`
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
		in.Image = DefaultIndexerImage
	}
	if in.Port == 0 {
		in.Port = DefaultIndexerHTTPPort
	}
	if in.ContainerName == "" {
		in.ContainerName = DefaultIndexerName
	}
	if in.DB == nil {
		in.DB = &DBInput{
			Image: DefaultIndexerDBImage,
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
	p, err := CwdSourcePath(in.SourceCodePath)
	if err != nil {
		return in.Out, err
	}

	/* Database */

	_, err = postgres.Run(ctx,
		in.DB.Image,
		testcontainers.WithName(DefaultIndexerDBName),
		testcontainers.WithExposedPorts("5432/tcp"),
		testcontainers.WithHostConfigModifier(func(h *container.HostConfig) {
			h.PortBindings = nat.PortMap{
				"5432/tcp": []nat.PortBinding{
					{HostPort: strconv.Itoa(DefaultIndexerDBPort)},
				},
			}
		}),
		testcontainers.WithLabels(framework.DefaultTCLabels()),
		postgres.WithDatabase(DefaultIndexerName),
		postgres.WithUsername(DefaultIndexerName),
		postgres.WithPassword(DefaultIndexerName),
		postgres.WithInitScripts(filepath.Join(p, DefaultIndexerSQLInit)),
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

	return &IndexerOutput{
		ContainerName:      in.ContainerName,
		ExternalHTTPURL:    fmt.Sprintf("http://%s:%d", host, in.Port),
		InternalHTTPURL:    fmt.Sprintf("http://%s:%d", in.ContainerName, in.Port),
		DBConnectionString: DefaultIndexerDBConnectionString,
	}, nil
}
