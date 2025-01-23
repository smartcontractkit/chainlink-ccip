package contracts

import (
	"testing"

	bin "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/testutils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"
)

func TestE2ECCIPRouter(t *testing.T) {

	t.Parallel()

	// Deploy Router Program
	ccip_router.SetProgramID(config.CcipRouterProgram)
	ctx := tests.Context(t)

	// Generate CCIP Admin
	ccipAdmin, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)
	// Address of Fee Aggregator, the account that will be used for the ATA of billing fees
	feeAggregator, gerr := solana.NewRandomPrivateKey()
	require.NoError(t, gerr)

	// Deploy CCIP Router by CCIP Admin
	solanaGoClient := testutils.DeployAllPrograms(t, testutils.PathToAnchorConfig, ccipAdmin)

	t.Run("CONFIG: fund CCIP Admin account with SOL", func(t *testing.T) {
		testutils.FundAccounts(ctx, []solana.PrivateKey{ccipAdmin, feeAggregator}, solanaGoClient, t)
	})

	t.Run("CONFIG: Initialize CCIP Router Program", func(t *testing.T) {

		// get program data account
		data, err := solanaGoClient.GetAccountInfoWithOpts(ctx, config.CcipRouterProgram, &rpc.GetAccountInfoOpts{
			Commitment: config.DefaultCommitment,
		})
		require.NoError(t, err)

		// Decode program data
		var programData struct {
			DataType uint32
			Address  solana.PublicKey
		}
		require.NoError(t, bin.UnmarshalBorsh(&programData, data.Bytes()))

		instruction, err := ccip_router.NewInitializeInstruction(
			config.SVMChainSelector,
			config.DefaultGasLimit,
			config.DefaultAllowOutOfOrderExecution,
			config.DefaultEnableExecutionAfter,
			feeAggregator.PublicKey(),
			// We use token2022 as the LINK address, which will be used as a base
			// for fees. It could be any other token mint address, but we use this
			// one for simplicity.
			linkTokenMint.PublicKey(), // TODO: Calculate this
			config.DefaultMaxFeeJuelsPerMsg,
			config.RouterConfigPDA,
			config.RouterStatePDA,
			ccipAdmin.PublicKey(),
			solana.SystemProgramID,
			config.CcipRouterProgram,
			programData.Address,
			config.ExternalExecutionConfigPDA,
			config.ExternalTokenPoolsSignerPDA,
		).ValidateAndBuild()
		require.NoError(t, err)

		result := testutils.SendAndConfirm(ctx, t, solanaGoClient, []solana.Instruction{instruction}, ccipAdmin, config.DefaultCommitment)
		require.NotNil(t, result)

		// Fetch account data
		var configAccount ccip_router.Config
		err = common.GetAccountDataBorshInto(ctx, solanaGoClient, config.RouterConfigPDA, config.DefaultCommitment, &configAccount)
		if err != nil {
			require.NoError(t, err, "failed to get account info")
		}
		require.Equal(t, config.SVMChainSelector, configAccount.SvmChainSelector)
		require.Equal(t, config.DefaultGasLimit, configAccount.DefaultGasLimit)
		require.Equal(t, uint8(0x0), configAccount.DefaultAllowOutOfOrderExecution) // false
	})
}
