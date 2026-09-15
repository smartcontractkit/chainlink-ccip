package glamsterdam_test

import (
	"math/big"
	"strings"
	"testing"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/glamsterdam"
)

const (
	ttfcTargetChainSel  = uint64(3379446385462418246)
	ttfcBaselineChain   = uint64(4949039107694359620)
	ttfcMismatchedChain = uint64(5548718428018410741)
)

var (
	ttfcUSDCToken       = common.HexToAddress("0x1111111111111111111111111111111111111111")
	ttfcUnconfiguredTok = common.HexToAddress("0x2222222222222222222222222222222222222222")
)

func deployTTFCFeeQuoter(t *testing.T, e *cldf.Environment, chainSel uint64, usdcDestGasOverhead uint32) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq16ops.Deploy, chain, contract.DeployInput[fq16ops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fq16ops.ContractType, *fq16ops.Version),
		Args: fq16ops.ConstructorArgs{
			StaticConfig: fq16ops.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1e18),
				LinkToken:                    common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
				TokenPriceStalenessThreshold: 3600,
			},
			TokenTransferFeeConfigArgs: []fq16ops.TokenTransferFeeConfigArgs{
				{
					DestChainSelector: ttfcTargetChainSel,
					TokenTransferFeeConfigs: []fq16ops.TokenTransferFeeConfigSingleTokenArgs{
						{
							Token: ttfcUSDCToken,
							TokenTransferFeeConfig: fq16ops.TokenTransferFeeConfig{
								MinFeeUSDCents:    50,
								MaxFeeUSDCents:    500,
								DestGasOverhead:   usdcDestGasOverhead,
								DestBytesOverhead: 100,
								IsEnabled:         true,
							},
						},
					},
				},
			},
		},
	})
	require.NoError(t, err)

	return common.HexToAddress(out.Output.Address)
}

func decodeApplyTokenTransferFeeConfigUpdatesArgs(t *testing.T, data []byte) fq16ops.ApplyTokenTransferFeeConfigUpdatesArgs {
	t.Helper()
	parsed, err := ethabi.JSON(strings.NewReader(fq16ops.FeeQuoterABI))
	require.NoError(t, err)
	vals, err := parsed.Methods["applyTokenTransferFeeConfigUpdates"].Inputs.Unpack(data[4:])
	require.NoError(t, err)
	feeConfigArgs := *ethabi.ConvertType(vals[0], new([]fq16ops.TokenTransferFeeConfigArgs)).(*[]fq16ops.TokenTransferFeeConfigArgs)
	removeArgs := *ethabi.ConvertType(vals[1], new([]fq16ops.TokenTransferFeeConfigRemoveArgs)).(*[]fq16ops.TokenTransferFeeConfigRemoveArgs)
	return fq16ops.ApplyTokenTransferFeeConfigUpdatesArgs{
		TokenTransferFeeConfigArgs:   feeConfigArgs,
		TokensToUseDefaultFeeConfigs: removeArgs,
	}
}

func TestUpdateTokenTransferFeeConfig(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{
		ttfcBaselineChain,
		ttfcMismatchedChain,
	}))
	require.NoError(t, err)

	fqBaselineAddr := deployTTFCFeeQuoter(t, e, ttfcBaselineChain, 180_000)
	// Mismatched chain: USDC's current DestGasOverhead doesn't match the expected Prague baseline
	// (180,000) — exercises the fallback path instead of the literal Glamsterdam value.
	fqMismatchedAddr := deployTTFCFeeQuoter(t, e, ttfcMismatchedChain, 150_000)

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateTokenTransferFeeConfig, e.BlockChains, glamsterdamseq.UpdateTokenTransferFeeConfigInput{
		TargetChainSelector: ttfcTargetChainSel,
		Lanes: []glamsterdamseq.TokenTransferFeeConfigLane{
			{
				ChainSelector:    ttfcBaselineChain,
				FeeQuoterAddress: fqBaselineAddr,
				// ttfcUnconfiguredTok has no override on this lane -> must be silently skipped.
				CandidateTokens: []common.Address{ttfcUSDCToken, ttfcUnconfiguredTok},
			},
			{
				ChainSelector:    ttfcMismatchedChain,
				FeeQuoterAddress: fqMismatchedAddr,
				CandidateTokens:  []common.Address{ttfcUSDCToken},
			},
		},
	})
	require.NoError(t, err)

	require.Len(t, report.Output.BatchOps, 2, "chains with an enabled override should each produce exactly one batch op")
	for _, batchOp := range report.Output.BatchOps {
		require.Len(t, batchOp.Transactions, 1)
	}

	reportStr := report.Output.Report.String()

	t.Run("baseline chain applies literal Glamsterdam value and preserves untouched fields", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 4949039107694359620: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (USDC) matched expected Prague value 180000, applying Glamsterdam value 540000")

		args := decodeApplyTokenTransferFeeConfigUpdatesArgs(t, report.Output.BatchOps[0].Transactions[0].Data)
		require.Len(t, args.TokenTransferFeeConfigArgs, 1)
		require.Len(t, args.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfigs, 1, "unconfigured token must not appear in the write")
		single := args.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfigs[0]
		require.Equal(t, ttfcUSDCToken, single.Token)
		require.Equal(t, uint32(540_000), single.TokenTransferFeeConfig.DestGasOverhead)
		require.Equal(t, uint32(100), single.TokenTransferFeeConfig.DestBytesOverhead) // preserved untouched field
		require.True(t, single.TokenTransferFeeConfig.IsEnabled)
	})

	t.Run("mismatched chain applies fallback instead of literal value", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 5548718428018410741: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (USDC) MISMATCH")
		require.Contains(t, reportStr, "current value 150000")
		require.Contains(t, reportStr, "applying fallback value 450000 instead of literal Glamsterdam value 540000")
	})

	t.Run("no writes executed directly, always routed through MCMS", func(t *testing.T) {
		for _, lane := range []struct {
			feeQuoter common.Address
			chainSel  uint64
		}{
			{fqBaselineAddr, ttfcBaselineChain},
			{fqMismatchedAddr, ttfcMismatchedChain},
		} {
			chain := e.BlockChains.EVMChains()[lane.chainSel]
			cur, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq16ops.GetTokenTransferFeeConfig, chain, contract.FunctionInput[fq16ops.GetTokenTransferFeeConfigArgs]{
				ChainSelector: lane.chainSel,
				Address:       lane.feeQuoter,
				Args:          fq16ops.GetTokenTransferFeeConfigArgs{DestChainSelector: ttfcTargetChainSel, Token: ttfcUSDCToken},
			})
			require.NoError(t, err)
			require.NotEqual(t, uint32(540_000), cur.Output.DestGasOverhead)
		}
	})
}
