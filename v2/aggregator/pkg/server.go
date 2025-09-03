// Package aggregator provides the main gRPC server implementation for the aggregator service.
package aggregator

import (
	"context"
	"fmt"
	"net"

	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pb/aggregator"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/aggregation"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/common"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/handlers"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Server represents a gRPC server for the aggregator service.
type Server struct {
	aggregator.UnimplementedAggregatorServer

	l                                    zerolog.Logger
	readCommitVerificationRecordHandler  *handlers.ReadCommitVerificationRecordHandler
	writeCommitVerificationRecordHandler *handlers.WriteCommitVerificationRecordHandler
}

// WriteCommitVerification handles requests to write commit verification records.
func (s *Server) WriteCommitVerification(ctx context.Context, req *aggregator.WriteCommitVerificationRequest) (*aggregator.WriteCommitVerificationResponse, error) {
	return s.writeCommitVerificationRecordHandler.Handle(ctx, req)
}

// ReadCommitVerification handles requests to read commit verification records.
func (s *Server) ReadCommitVerification(ctx context.Context, req *aggregator.ReadCommitVerificationRequest) (*aggregator.ReadCommitVerificationResponse, error) {
	return s.readCommitVerificationRecordHandler.Handle(ctx, req)
}

// Start starts the gRPC server on the provided listener.
func (s *Server) Start(lis net.Listener) (func(), error) {
	grpcServer := grpc.NewServer()
	aggregator.RegisterAggregatorServer(grpcServer, s)
	reflection.Register(grpcServer)

	s.l.Info().Msg("Aggregator gRPC server started")
	if err := grpcServer.Serve(lis); err != nil {
		return func() {}, err
	}

	return grpcServer.Stop, nil
}

func createAggregatorConfig(storage common.CommitVerificationStore, config model.AggregatorConfig) (handlers.AggregationTriggerer, error) {
	if config.Aggregation.AggregationStrategy == "stub" {
		aggregator := aggregation.NewCommitReportAggregator(storage, &aggregation.AggregatorSinkStub{}, config)
		aggregator.StartBackground(context.Background())
		return aggregator, nil
	}

	return nil, fmt.Errorf("unknown aggregation strategy: %s", config.Aggregation.AggregationStrategy)
}

func createStorage(config model.AggregatorConfig) (common.CommitVerificationStore, error) {
	if config.Storage.StorageType == "memory" {
		return storage.NewInMemoryStorage(), nil
	}

	return nil, fmt.Errorf("unknown storage type: %s", config.Storage.StorageType)
}

// NewServer creates a new aggregator server with the specified logger and configuration.
func NewServer(l zerolog.Logger, config model.AggregatorConfig) *Server {
	store, err := createStorage(config)
	if err != nil {
		l.Error().Err(err).Msg("failed to create storage")
		return nil
	}

	aggregator, err := createAggregatorConfig(store, config)
	if err != nil {
		l.Error().Err(err).Msg("failed to create aggregator")
		return nil
	}

	readCommitVerificationRecordHandler := handlers.NewReadCommitVerificationRecordHandler(store)
	writeCommitVerificationRecordHandler := handlers.NewWriteCommitVerificationRecordHandler(store, aggregator)

	return &Server{
		l:                                    l,
		readCommitVerificationRecordHandler:  readCommitVerificationRecordHandler,
		writeCommitVerificationRecordHandler: writeCommitVerificationRecordHandler,
	}
}
