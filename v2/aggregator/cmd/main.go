// Package main provides the entry point for the aggregator service.
package main

import (
	"context"
	"net"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	aggregator "github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
)

func main() {
	lvlStr := os.Getenv("AGGREGATOR_LOG_LEVEL")
	if lvlStr == "" {
		lvlStr = "info"
	}
	lvl, err := zerolog.ParseLevel(lvlStr)
	if err != nil {
		panic(err)
	}
	l := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(lvl)

	config := model.AggregatorConfig{
		Server: model.ServerConfig{
			Address: ":50051",
		},
		Storage: model.StorageConfig{
			StorageType: "memory",
		},
		Aggregation: model.AggregationConfig{
			AggregationStrategy: "stub",
		},
	}

	server := aggregator.NewServer(l, config)

	address := config.Server.Address
	lc := &net.ListenConfig{}
	lis, err := lc.Listen(context.Background(), "tcp", address)
	if err != nil {
		l.Fatal().Err(err).Msg("failed to listen")
	}
	defer func() {
		if err := lis.Close(); err != nil {
			l.Error().Err(err).Msg("failed to close listener")
		}
	}()

	stop, err := server.Start(lis)
	if err != nil {
		l.Fatal().Err(err).Msg("failed to start server")
	}
	defer stop()
}
