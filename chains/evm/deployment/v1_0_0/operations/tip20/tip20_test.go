package tip20_test

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func TestTIP20FactoryABI_Parses(t *testing.T) {
	t.Parallel()

	_, err := abi.JSON(strings.NewReader(tip20.TIP20FactoryABI))
	require.NoError(t, err, "TIP20FactoryABI must remain valid JSON ABI")
}

func TestNewTIP20Factory_Constructs(t *testing.T) {
	t.Parallel()

	key, err := crypto.GenerateKey()
	require.NoError(t, err)

	address := crypto.PubkeyToAddress(key.PublicKey)
	backend := simulated.NewBackend(types.GenesisAlloc{
		address: {Balance: big.NewInt(1_000_000_000_000_000_000)},
	})
	t.Cleanup(func() { backend.Close() })

	factory, err := tip20.NewTIP20Factory(common.HexToAddress(tip20.TokenFactoryAddress), backend.Client())
	require.NoError(t, err)

	require.Equal(t, common.HexToAddress(tip20.TokenFactoryAddress), factory.Address())
}

func TestDeploy_RejectsNonTempoChain(t *testing.T) {
	t.Parallel()

	evmSel := chainsel.ETHEREUM_MAINNET.Selector
	chains := []uint64{evmSel}

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, chains))
	require.NoError(t, err)

	chain, ok := e.BlockChains.EVMChains()[evmSel]
	require.True(t, ok)

	_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, tip20.Deploy, chain, tip20.FactoryDeployArgs{
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
