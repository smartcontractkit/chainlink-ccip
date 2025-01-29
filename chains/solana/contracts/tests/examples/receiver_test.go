package examples

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	ccip_receiver "github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/example_ccip_receiver"
)

func TestCcipReceiver(t *testing.T) {
	t.Parallel()
	ctx := tests.Context(t)

	ccip_receiver.SetProgramID(config.CcipBaseReceiver)

	owner, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	user, err := solana.NewRandomPrivateKey()
	require.NoError(t, err)

	solClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, owner)

	t.Run("setup", func(t *testing.T) {
		t.Run("funding", func(t *testing.T) {
			testutils.FundAccounts(ctx, []solana.PrivateKey{owner, user}, solClient, t)
		})

		t.Run("initialize receiver", func(t *testing.T) {

		})
	})
}
