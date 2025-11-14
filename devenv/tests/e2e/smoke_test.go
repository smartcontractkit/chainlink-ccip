package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"

	ccip "github.com/smartcontractkit/chainlink-ccip/devenv"
)

func TestE2ESmoke(t *testing.T) {
	in, err := ccip.LoadOutput[ccip.Cfg]("../../env-out.toml")
	require.NoError(t, err)

	chainIDs, wsURLs := make([]string, 0), make([]string, 0)
	for _, bc := range in.Blockchains {
		chainIDs = append(chainIDs, bc.ChainID)
		wsURLs = append(wsURLs, bc.Out.Nodes[0].ExternalWSUrl)
	}

	selectors, e, err := ccip.NewCLDFOperationsEnvironment(in.Blockchains, in.CLDF.DataStore)
	require.NoError(t, err)

	impls := make([]ccip.CCIP16ProductConfiguration, 0)
	for _, bc := range in.Blockchains {
		i, err := ccip.NewCCIPImplFromNetwork(bc.Out.Type)
		require.NoError(t, err)
		i.SetCLDF(e)
		impls = append(impls, i)
	}

	t.Cleanup(func() {
		_, err := framework.SaveContainerLogs(fmt.Sprintf("%s-%s", framework.DefaultCTFLogsDir, t.Name()))
		require.NoError(t, err)
	})

	t.Run("EVM<>EVM test CCIP trasfers", func(t *testing.T) {
		type testcase struct {
			name         string
			fromSelector uint64
			toSelector   uint64
			implOne      ccip.CCIP16ProductConfiguration
			implTwo      ccip.CCIP16ProductConfiguration
		}

		tcs := []testcase{
			{
				name:         "src->dst msg execution eoa receiver",
				fromSelector: selectors[0],
				toSelector:   selectors[1],
				implOne:      impls[0],
				implTwo:      impls[1],
			},
			{
				name:         "dst->src msg execution eoa receiver",
				fromSelector: selectors[1],
				toSelector:   selectors[0],
				implOne:      impls[0],
				implTwo:      impls[1],
			},
		}
		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				t.Logf("Testing CCIP message from chain %d to chain %d", tc.fromSelector, tc.toSelector)
				tc.implOne.SendMessage(t.Context(), tc.fromSelector, tc.toSelector, nil, nil)
				tc.implOne.WaitOneSentEventBySeqNo(t.Context(), tc.fromSelector, tc.toSelector, 0, 5*time.Minute)
				// tc.implOne.WaitOneExecEventBySeqNo()
				// tc.implTwo.SendMessage(..)
			})
		}
	})
}
