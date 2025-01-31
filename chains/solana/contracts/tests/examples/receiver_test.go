package examples

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
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

	receiverState, _, err := solana.FindProgramAddress([][]byte{[]byte("state")}, config.CcipBaseReceiver)
	require.NoError(t, err)

	solClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, owner)

	t.Run("setup", func(t *testing.T) {
		t.Run("funding", func(t *testing.T) {
			testutils.FundAccounts(ctx, []solana.PrivateKey{owner, user}, solClient, t)
		})

		t.Run("initialize receiver", func(t *testing.T) {
			ix, err := ccip_receiver.NewInitializeInstruction(
				config.CcipRouterProgram,
				receiverState,
				owner.PublicKey(),
				solana.SystemProgramID,
			).ValidateAndBuild()
			require.NoError(t, err)

			testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ix}, owner, rpc.CommitmentConfirmed)
		})
	})

	t.Run("enable/disable chains", func(t *testing.T) {
		ixAllow, err := ccip_receiver.NewAddChainToInstruction(ccip_receiver.Allow_ListType, config.EvmChainSelector, receiverState, owner.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
		require.NoError(t, err)
		ixDeny, err := ccip_receiver.NewAddChainToInstruction(ccip_receiver.Deny_ListType, config.SVMChainSelector, receiverState, owner.PublicKey(), solana.SystemProgramID).ValidateAndBuild()
		require.NoError(t, err)
		ixEnableAllow, err := ccip_receiver.NewEnableListInstruction(ccip_receiver.Allow_ListType, true, receiverState, owner.PublicKey()).ValidateAndBuild()
		require.NoError(t, err)
		ixEnableDeny, err := ccip_receiver.NewEnableListInstruction(ccip_receiver.Deny_ListType, true, receiverState, owner.PublicKey()).ValidateAndBuild()
		require.NoError(t, err)

		testutils.SendAndConfirm(ctx, t, solClient, []solana.Instruction{ixAllow, ixDeny, ixEnableAllow, ixEnableDeny}, owner, rpc.CommitmentConfirmed)
	})

	t.Run("check ccip_receiver constraints", func(t *testing.T) {
		t.Run("invalid chain", func(t *testing.T) {
			t.Parallel()
			ix, err := ccip_receiver.NewCcipReceiveInstruction(ccip_receiver.Any2SVMMessage{SourceChainSelector: config.SVMChainSelector}, user.PublicKey(), receiverState).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solClient, []solana.Instruction{ix}, user, rpc.CommitmentConfirmed, []string{"Invalid chain"})
		})
		t.Run("invalid sender", func(t *testing.T) {
			t.Parallel()
			ix, err := ccip_receiver.NewCcipReceiveInstruction(ccip_receiver.Any2SVMMessage{SourceChainSelector: config.EvmChainSelector}, user.PublicKey(), receiverState).ValidateAndBuild()
			require.NoError(t, err)
			testutils.SendAndFailWith(ctx, t, solClient, []solana.Instruction{ix}, user, rpc.CommitmentConfirmed, []string{"Address is not router external execution PDA"})
		})
	})
}
