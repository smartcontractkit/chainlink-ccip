package contracts

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/tokens"
)

//go:embed devnet.config.yaml
var devnetInfoBuffer []byte

func TestDevnet(t *testing.T) {
	ctx := tests.Context(t)

	type DevnetInfo struct {
		Router    string `yaml:"router"`
		UserPrivK []byte `yaml:"user_privk"`
	}
	var devnetInfo DevnetInfo
	err := yaml.Unmarshal(devnetInfoBuffer, &devnetInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config: %+v\n", devnetInfo)

	client := rpc.New("https://api.devnet.solana.com")

	routerAddress, err := solana.PublicKeyFromBase58(devnetInfo.Router)
	require.NoError(t, err)

	user := solana.PrivateKey(devnetInfo.UserPrivK)
	require.True(t, user.IsValid())

	configPDA, _, err := state.FindConfigPDA(routerAddress)
	require.NoError(t, err)

	var routerConfig ccip_router.Config

	t.Run("Read Config PDA", func(t *testing.T) {
		common.GetAccountDataBorshInto(ctx, client, configPDA, rpc.CommitmentConfirmed, &routerConfig)
		fmt.Printf("Router Config: %+v\n", routerConfig)
	})

	t.Run("OnRamp: CCIP Send", func(t *testing.T) {
		destinationChainSelector := uint64(16423721717087811551) // "svm"

		destinationChainStatePDA, err := state.FindDestChainStatePDA(destinationChainSelector, routerAddress)
		require.NoError(t, err)

		nonceEvmPDA, err := state.FindNoncePDA(destinationChainSelector, user.PublicKey(), routerAddress)
		require.NoError(t, err)

		billingSignerPDA, _, err := state.FindFeeBillingSignerPDA(routerAddress)
		require.NoError(t, err)

		// TODO: create in FeeQuoter the wsol receiver --> Register WSOL as a valid fee token
		wsolReceiver, _, err := tokens.FindAssociatedTokenAddress(solana.TokenProgramID, solana.SolMint, billingSignerPDA)
		require.NoError(t, err)

		fqConfigPDA, _, err := state.FindFqConfigPDA(routerConfig.FeeQuoter)
		require.NoError(t, err)

		fqDestChainPDA, _, err := state.FindFqDestChainPDA(destinationChainSelector, routerConfig.FeeQuoter)
		require.NoError(t, err)

		fqBillingTokenConfigPDA, _, err := state.FindFqBillingTokenConfigPDA(solana.SolMint, routerConfig.FeeQuoter)
		require.NoError(t, err)

		fqLinkBillingConfigPDA, _, err := state.FindFqBillingTokenConfigPDA(routerConfig.LinkTokenMint, routerConfig.FeeQuoter)
		require.NoError(t, err)

		externalTpSignerPDA, _, err := state.FindExternalTokenPoolsSignerPDA(routerAddress)
		require.NoError(t, err)

		message := ccip_router.SVM2AnyMessage{
			FeeToken:  solana.SolMint,
			Receiver:  []byte{1, 2, 3}, // TODO
			Data:      []byte{4, 5, 6},
			ExtraArgs: []byte{},
		}

		raw := ccip_router.NewCcipSendInstruction(
			destinationChainSelector,
			message,
			[]byte{},
			configPDA,
			destinationChainStatePDA,
			nonceEvmPDA,
			user.PublicKey(),
			solana.SystemProgramID,
			solana.TokenProgramID,
			solana.SolMint,
			solana.PublicKey{},
			wsolReceiver,
			billingSignerPDA,
			routerConfig.FeeQuoter,
			fqConfigPDA,
			fqDestChainPDA,
			fqBillingTokenConfigPDA,
			fqLinkBillingConfigPDA,
			externalTpSignerPDA,
		)
		raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
		instruction, err := raw.ValidateAndBuild()
		require.NoError(t, err)
		result := testutils.SendAndConfirm(ctx, t, client, []solana.Instruction{instruction}, user, rpc.CommitmentConfirmed)
		require.NotNil(t, result)
	})
}

/*
	destinationChainSelector := config.EvmChainSelector
	destinationChainStatePDA := config.EvmDestChainStatePDA
	message := ccip_router.SVM2AnyMessage{
		FeeToken:  wsol.mint,
		Receiver:  validReceiverAddress[:],
		Data:      []byte{4, 5, 6},
		ExtraArgs: emptyEVMExtraArgsV2,
	}

	raw := ccip_router.NewCcipSendInstruction(
		destinationChainSelector,
		message,
		[]byte{},
		config.RouterConfigPDA,
		destinationChainStatePDA,
		nonceEvmPDA,
		user.PublicKey(),
		solana.SystemProgramID,
		wsol.program,
		wsol.mint,
		wsol.userATA,
		wsol.billingATA,
		config.BillingSignerPDA,
		config.FeeQuoterProgram,
		config.FqConfigPDA,
		config.FqEvmDestChainPDA,
		wsol.fqBillingConfigPDA,
		link22.fqBillingConfigPDA,
		config.ExternalTokenPoolsSignerPDA,
	)
	raw.GetFeeTokenUserAssociatedAccountAccount().WRITE()
	instruction, err := raw.ValidateAndBuild()
	require.NoError(t, err)
	result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, user, config.DefaultCommitment)
	require.NotNil(t, result)
*/
