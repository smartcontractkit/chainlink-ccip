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
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	"github.com/stretchr/testify/require"
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
