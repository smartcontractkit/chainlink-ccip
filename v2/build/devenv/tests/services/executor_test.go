package services_test

import (
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/v2/devenv/services"
	"github.com/stretchr/testify/require"
)

func TestServiceExecutor(t *testing.T) {
	out, err := services.NewExecutor(&services.ExecutorInput{SourceCodePath: "../../../executor"})
	require.NoError(t, err)
	t.Run("test #1", func(t *testing.T) {
		_ = out
		// use the data from output, connect HTTP, gRPC clients etc and do the testing
	})
}
