package e2e

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccv/protocol"
	"github.com/smartcontractkit/chainlink-testing-framework/framework"

	ccipEVM "github.com/smartcontractkit/chainlink-ccip/ccip-evm"
	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
)

const (
	// See Internal.sol for the full enum values.
	MessageExecutionStateSuccess uint8 = 2
	MessageExecutionStateFailed  uint8 = 3

	defaultSentTimeout = 10 * time.Second
	defaultExecTimeout = 40 * time.Second
)

func TestE2ESmoke(t *testing.T) {
	in, err := ccip.LoadOutput[ccip.Cfg]("../../env-out.toml")
	require.NoError(t, err)
	ctx := context.Background()
	l := zerolog.Ctx(ctx)

	chainIDs, wsURLs := make([]string, 0), make([]string, 0)
	for _, bc := range in.Blockchains {
		chainIDs = append(chainIDs, bc.ChainID)
		wsURLs = append(wsURLs, bc.Out.Nodes[0].ExternalWSUrl)
	}

	selectors, e, err := ccip.NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	require.NoError(t, err)

	c, err := ccipEVM.NewCCIP16EVM(ctx, *l, e, chainIDs, wsURLs)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, err := framework.SaveContainerLogs(fmt.Sprintf("%s-%s", framework.DefaultCTFLogsDir, t.Name()))
		require.NoError(t, err)
	})

	t.Run("test extra args v2 messages", func(t *testing.T) {
		type testcase struct {
			name                     string
			fromSelector             uint64
			toSelector               uint64
			receiver                 protocol.UnknownAddress
			expectFail               bool
			numExpectedVerifications int
		}

		tcs := []testcase{
			{
				name:                     "src->dst msg execution eoa receiver",
				fromSelector:             selectors[0],
				toSelector:               selectors[1],
				receiver:                 mustGetEOAReceiverAddress(t, ctx, c, selectors[1]),
				expectFail:               false,
				numExpectedVerifications: 1,
			},
			{
				name:                     "dst->src msg execution eoa receiver",
				fromSelector:             selectors[1],
				toSelector:               selectors[0],
				receiver:                 mustGetEOAReceiverAddress(t, ctx, c, selectors[0]),
				expectFail:               false,
				numExpectedVerifications: 1,
			},
		}
		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				_ = c
			})
		}
	})
}

func mustGetEOAReceiverAddress(t *testing.T, ctx context.Context, c *ccipEVM.CCIP16EVM, chainSelector uint64) protocol.UnknownAddress {
	receiver, err := c.GetEOAReceiverAddress(ctx, chainSelector)
	require.NoError(t, err)
	return receiver
}
