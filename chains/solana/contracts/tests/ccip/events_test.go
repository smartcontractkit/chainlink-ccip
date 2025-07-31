package contracts

import (
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/test_event_emitter"
	solccip "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/ccip"
	solcommon "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"
)

func TestEventEmitter(t *testing.T) {
	t.Parallel()

	ctx := tests.Context(t)
	test_event_emitter.SetProgramID(config.EventEmitter)

	deployer, err := solana.PrivateKeyFromBase58("4whgxZhpxcArYWzM1iTmokruAzws9YVi2f9M7pWwchQniaFXBr1WGSGXgadeqHtiRooxNiPosdLj2g2ohbtkWtu5")
	require.NoError(t, err)

	rpcClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, deployer)
	testutils.FundAccounts(ctx, []solana.PrivateKey{deployer}, rpcClient, t)

	t.Run("emits CcipCctpMessageSentEvent", func(t *testing.T) {
		args := test_event_emitter.CcipCctpMessageSentEventArgs{
			RemoteChainSelector: config.EvmChainSelector,
			OriginalSender:      deployer.PublicKey(),
			EventAddress:        deployer.PublicKey(),
			MessageSentBytes:    nil,
			MsgTotalNonce:       0,
			SourceDomain:        5,
			CctpNonce:           0,
		}

		ix, err := test_event_emitter.NewEmitCcipCctpMsgSentInstruction(args, solana.SysVarClockPubkey).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, rpcClient, []solana.Instruction{ix}, deployer, config.DefaultCommitment)
		require.NotNil(t, result)

		event := solccip.EventCcipCctpMessageSent{}
		print := false

		err = solcommon.ParseEvent(result.Meta.LogMessages, "CcipCctpMessageSentEvent", &event, print)
		require.NoError(t, err)

		require.Equal(t, args.RemoteChainSelector, event.RemoteChainSelector)
		require.Equal(t, args.MessageSentBytes, event.MessageSentBytes)
		require.Equal(t, args.OriginalSender, event.OriginalSender)
		require.Equal(t, args.MsgTotalNonce, event.MsgTotalNonce)
		require.Equal(t, args.EventAddress, event.EventAddress)
		require.Equal(t, args.SourceDomain, event.SourceDomain)
		require.Equal(t, args.CctpNonce, event.CctpNonce)
	})
}
