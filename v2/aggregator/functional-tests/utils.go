package functionaltests

import (
	"context"
	"net"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pb/aggregator"
	agg "github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg"
	"github.com/smartcontractkit/chainlink-ccip/v2/aggregator/pkg/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

// CreateServerAndClient creates a test server and client for functional testing.
func CreateServerAndClient(t *testing.T) (aggregator.AggregatorClient, func(), error) {
	lis := bufconn.Listen(bufSize)
	l := log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	s := agg.NewServer(l, model.AggregatorConfig{
		Server: model.ServerConfig{
			Address: ":50051",
		},
		Storage: model.StorageConfig{
			StorageType: "memory",
		},
		Aggregation: model.AggregationConfig{
			AggregationStrategy: "stub",
		},
	})
	go func() {
		_, _ = s.Start(lis)
	}()

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	bufDialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	//nolint:staticcheck // grpc.WithInsecure is deprecated but needed for test setup
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to dial bufnet: %v", err)
	}

	client := aggregator.NewAggregatorClient(conn)
	return client, func() {
		if err := conn.Close(); err != nil {
			t.Errorf("failed to close connection: %v", err)
		}
		if err := lis.Close(); err != nil {
			t.Errorf("failed to close listener: %v", err)
		}
	}, nil
}
