package glamsterdam

import (
	"testing"

	"github.com/stretchr/testify/require"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

func fakeWrite(chainSelector uint64, to string) contract.WriteOutput {
	return contract.WriteOutput{ChainSelector: chainSelector, Tx: mcms_types.Transaction{To: to}}
}

func TestBuildLaneBatchOps(t *testing.T) {
	const sel = uint64(1)
	core := []contract.WriteOutput{fakeWrite(sel, "onramp"), fakeWrite(sel, "feequoter"), fakeWrite(sel, "committeeverifier")}
	isolated := []contract.WriteOutput{fakeWrite(sel, "lombardverifier"), fakeWrite(sel, "cctpverifier")}

	ops, err := buildLaneBatchOps(core, isolated)
	require.NoError(t, err)
	require.Len(t, ops, 3, "one core batch + one batch per isolated write")

	require.Len(t, ops[0].Transactions, 3)
	require.Equal(t, []string{"onramp", "feequoter", "committeeverifier"}, []string{ops[0].Transactions[0].To, ops[0].Transactions[1].To, ops[0].Transactions[2].To})
	for i, want := range []string{"lombardverifier", "cctpverifier"} {
		require.Len(t, ops[i+1].Transactions, 1, "isolated batch must hold exactly one tx")
		require.Equal(t, want, ops[i+1].Transactions[0].To)
	}
	for _, op := range ops {
		require.Equal(t, mcms_types.ChainSelector(sel), op.ChainSelector)
	}

	t.Run("no isolated writes keeps a single core batch", func(t *testing.T) {
		ops, err := buildLaneBatchOps(core, nil)
		require.NoError(t, err)
		require.Len(t, ops, 1)
		require.Len(t, ops[0].Transactions, 3)
	})
}
