package tip20_test

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	"github.com/stretchr/testify/require"
)

func TestTIP20TokenABI_Parses(t *testing.T) {
	t.Parallel()

	_, err := abi.JSON(strings.NewReader(tip20.TIP20TokenABI))
	require.NoError(t, err, "TIP20TokenABI must remain valid JSON ABI")
}
