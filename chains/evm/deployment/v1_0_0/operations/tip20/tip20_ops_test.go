package tip20_test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func TestDeploy_RejectsNonTempoChain(t *testing.T) {
	t.Parallel()

	evmSel := chainsel.ETHEREUM_MAINNET.Selector
	chains := []uint64{evmSel}

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, chains))
	require.NoError(t, err)

	chain, ok := e.BlockChains.EVMChains()[evmSel]
	require.True(t, ok)

	_, err = operations.ExecuteSequence(e.OperationsBundle, tip20.Deploy, chain, tip20.FactoryDeployArgs{
		QuoteToken: common.Address{}, // defaults to sensible value
		Currency:   "",               // defaults to sensible value
		Salt:       [32]byte{},       // generate random salt
		Name:       "MyTestToken",
		Symbol:     "MTK",
		Admin:      chain.DeployerKey.From,
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "only supported on Tempo testnet and mainnet")
}
