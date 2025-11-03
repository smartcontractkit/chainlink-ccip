package sequences_test

import (
	"testing"

	sequencescommon "github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/stretchr/testify/require"
)

type mockCfg struct {
	chainSel uint64
}

func (c mockCfg) ChainSelector() uint64 {
	return c.chainSel
}

func TestResolveEVMChainDep(t *testing.T) {
	tests := []struct {
		desc        string
		expectedErr string
		chains      []chain.BlockChain
		cfg         sequencescommon.WithChainSelector
	}{
		{
			desc: "happy path",
			chains: []chain.BlockChain{
				evm.Chain{Selector: 1},
			},
			cfg: mockCfg{chainSel: 1},
		},
		{
			desc:        "no matching chain selector",
			expectedErr: "no EVM chain with selector 2 found in environment",
			chains: []chain.BlockChain{
				evm.Chain{Selector: 1},
			},
			cfg: mockCfg{chainSel: 2},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			e := deployment.Environment{}
			e.BlockChains = chain.NewBlockChainsFromSlice(test.chains)

			chain, err := sequences.ResolveEVMChainDep(e, test.cfg)
			if test.expectedErr == "" {
				require.NoError(t, err)
				require.Equal(t, uint64(1), chain.Selector)
			} else {
				require.Error(t, err)
				require.ErrorContains(t, err, test.expectedErr)
			}
		})
	}
}
