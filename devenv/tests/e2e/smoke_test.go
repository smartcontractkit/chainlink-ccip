package e2e

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-testing-framework/framework"

	chainsel "github.com/smartcontractkit/chain-selectors"
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
	selectorsToImpl := make(map[uint64]ccip.CCIP16ProductConfiguration)

	for _, bc := range in.Blockchains {
		i, err := ccip.NewCCIPImplFromNetwork(bc.Out.Type)
		require.NoError(t, err)
		i.SetCLDF(e)
		var family string
		switch bc.Type {
		case "anvil", "geth":
			family = chainsel.FamilyEVM
		case "solana":
			family = chainsel.FamilySolana
		default:
			panic("unsupported blockchain type")
		}
		networkInfo, err := chainsel.GetChainDetailsByChainIDAndFamily(bc.ChainID, family)
		require.NoError(t, err)
		selectorsToImpl[networkInfo.ChainSelector] = i
	}

	t.Cleanup(func() {
		_, err := framework.SaveContainerLogs(fmt.Sprintf("%s-%s", framework.DefaultCTFLogsDir, t.Name()))
		require.NoError(t, err)
	})

	t.Run("Test CCIP transfers", func(t *testing.T) {
		type testcase struct {
			name         string
			fromSelector uint64
			toSelector   uint64
		}

		tcs := []testcase{
			{
				name:         "evm->evm msg execution eoa receiver",
				fromSelector: selectors[0],
				toSelector:   selectors[1],
			},
			{
				name:         "evm->evm msg execution eoa receiver",
				fromSelector: selectors[1],
				toSelector:   selectors[0],
			},
			{
				name:         "evm->svm msg execution eoa receiver",
				fromSelector: selectors[0],
				toSelector:   selectors[2],
			},
			{
				name:         "svm->evm msg execution eoa receiver",
				fromSelector: selectors[2],
				toSelector:   selectors[0],
			},
		}
		for _, tc := range tcs {
			t.Run(tc.name, func(t *testing.T) {
				t.Logf("Testing CCIP message from chain %d to chain %d", tc.fromSelector, tc.toSelector)
				fromImpl := selectorsToImpl[tc.fromSelector]
				toImpl := selectorsToImpl[tc.toSelector]
				err := fromImpl.SendMessage(t.Context(), tc.fromSelector, tc.toSelector, nil, nil)
				require.NoError(t, err)
				seq, err := fromImpl.GetExpectedNextSequenceNumber(t.Context(), tc.fromSelector, tc.toSelector)
				require.NoError(t, err)
				_, err = toImpl.WaitOneSentEventBySeqNo(t.Context(), tc.fromSelector, tc.toSelector, seq, 2*time.Minute)
				require.NoError(t, err)
				_, err = toImpl.WaitOneExecEventBySeqNo(t.Context(), tc.fromSelector, tc.toSelector, seq, 2*time.Minute)
				require.NoError(t, err)
			})
		}
	})
}
