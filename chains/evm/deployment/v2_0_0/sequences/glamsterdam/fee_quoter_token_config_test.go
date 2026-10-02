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
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	tpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/glamsterdam"
)

var (
	fqTokCfgGenericToken    = common.HexToAddress("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fqTokCfgNoOverrideToken = common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
)

func fqTokCfgOverride(token common.Address, destGasOverhead uint32) fqops.TokenTransferFeeConfigSingleTokenArgs {
	return fqops.TokenTransferFeeConfigSingleTokenArgs{
		Token: token,
		TokenTransferFeeConfig: fqops.TokenTransferFeeConfig{
			FeeUSDCents:       150,
			DestGasOverhead:   destGasOverhead,
			DestBytesOverhead: 32, // must be >= Pool.CCIP_LOCK_OR_BURN_V1_RET_BYTES
			IsEnabled:         true,
		},
	}
}

// deployFQWithTokenOverrides deploys a FeeQuoter 2.0.0 with the target destination enabled and the
// given per-token overrides already set for it.
func deployFQWithTokenOverrides(t *testing.T, e *cldf.Environment, chainSel uint64, overrides []fqops.TokenTransferFeeConfigSingleTokenArgs) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]
	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fqops.Deploy, chain, contract.DeployInput[fqops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fqops.ContractType, *fqops.Version),
		Args: fqops.ConstructorArgs{
			StaticConfig: fqops.StaticConfig{MaxFeeJuelsPerMsg: big.NewInt(1e18), LinkToken: gasCfgLinkToken},
			DestChainConfigArgs: []fqops.DestChainConfigArgs{{
				DestChainSelector: tpGasCfgTargetChainSel,
				DestChainConfig: fqops.DestChainConfig{
					IsEnabled:                   true,
					MaxDataBytes:                1_000,
					MaxPerMsgGasLimit:           15_000_000,
					DestGasOverhead:             300_000,
					DestGasPerPayloadByteBase:   20,
					ChainFamilySelector:         [4]byte{0x28, 0x12, 0xd5, 0x2c},
					DefaultTokenDestGasOverhead: 90_000,
					DefaultTxGasLimit:           200_000,
					LinkFeeMultiplierPercent:    100,
				},
			}},
			TokenTransferFeeConfigArgs: []fqops.TokenTransferFeeConfigArgs{{
				DestChainSelector:       tpGasCfgTargetChainSel,
				TokenTransferFeeConfigs: overrides,
			}},
		},
	})
	require.NoError(t, err)
	return common.HexToAddress(out.Output.Address)
}

func poolToken(t *testing.T, e *cldf.Environment, chainSel uint64, pool common.Address) common.Address {
	t.Helper()
	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, tpops.GetToken, e.BlockChains.EVMChains()[chainSel], contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       pool,
	})
	require.NoError(t, err)
	return out.Output
}

func decodeApplyFQTokenTransferFeeConfigUpdates(t *testing.T, data []byte) []fqops.TokenTransferFeeConfigArgs {
	t.Helper()
	parsed, err := ethabi.JSON(strings.NewReader(fqops.FeeQuoterABI))
	require.NoError(t, err)
	vals, err := parsed.Methods["applyTokenTransferFeeConfigUpdates"].Inputs.Unpack(data[4:])
	require.NoError(t, err)
	return *ethabi.ConvertType(vals[0], new([]fqops.TokenTransferFeeConfigArgs)).(*[]fqops.TokenTransferFeeConfigArgs)
}

func TestUpdateFeeQuoterTokenTransferFeeConfig(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{tpGasCfgBaselineChain, tpGasCfgMismatchedChain}))
	require.NoError(t, err)

	// Baseline chain: USDC (250k), Lombard (410k), an unrelated token (150k) and a token with no
	// override at all.
	usdcPool := deployTokenPoolFixture(t, e, tpGasCfgBaselineChain, "usdc", 250_000)
	lombardPool := deployTokenPoolFixture(t, e, tpGasCfgBaselineChain, "lombard", 410_000)
	usdcToken := poolToken(t, e, tpGasCfgBaselineChain, usdcPool)
	lombardToken := poolToken(t, e, tpGasCfgBaselineChain, lombardPool)
	fqBaseline := deployFQWithTokenOverrides(t, e, tpGasCfgBaselineChain, []fqops.TokenTransferFeeConfigSingleTokenArgs{
		fqTokCfgOverride(usdcToken, 250_000),
		fqTokCfgOverride(lombardToken, 410_000),
		fqTokCfgOverride(fqTokCfgGenericToken, 150_000),
	})

	// Mismatched chain: USDC override (180k) doesn't match the 250k baseline -> ratio fallback.
	usdcPoolMismatched := deployTokenPoolFixture(t, e, tpGasCfgMismatchedChain, "usdc", 250_000)
	usdcTokenMismatched := poolToken(t, e, tpGasCfgMismatchedChain, usdcPoolMismatched)
	fqMismatched := deployFQWithTokenOverrides(t, e, tpGasCfgMismatchedChain, []fqops.TokenTransferFeeConfigSingleTokenArgs{
		fqTokCfgOverride(usdcTokenMismatched, 180_000),
	})

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateFeeQuoterTokenTransferFeeConfig, e.BlockChains, glamsterdamseq.UpdateFeeQuoterTokenTransferFeeConfigInput{
		TargetChainSelector: tpGasCfgTargetChainSel,
		Lanes: []glamsterdamseq.FeeQuoterTokenConfigLane{
			{
				ChainSelector:        tpGasCfgBaselineChain,
				FeeQuoterAddress:     fqBaseline,
				ExtraCandidateTokens: []common.Address{fqTokCfgGenericToken, fqTokCfgNoOverrideToken},
				USDCPoolAddresses:    []common.Address{usdcPool},
				LombardPoolAddresses: []common.Address{lombardPool},
			},
			{
				ChainSelector:     tpGasCfgMismatchedChain,
				FeeQuoterAddress:  fqMismatched,
				USDCPoolAddresses: []common.Address{usdcPoolMismatched},
			},
		},
	})
	require.NoError(t, err)

	require.Len(t, report.Output.BatchOps, 2, "one FeeQuoter write per chain")
	reportStr := report.Output.Report.String()

	t.Run("USDC, Lombard and generic tokens are migrated by their own rule", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 4949039107694359620: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (USDC) matched expected Prague value 250000, applying Glamsterdam value 750000")
		require.Contains(t, reportStr, "chain 4949039107694359620: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (Lombard) matched expected Prague value 410000, applying Glamsterdam value 1200000")
		require.Contains(t, reportStr, "(token "+fqTokCfgGenericToken.Hex()+") generic scaling 150000 -> 450000")
		require.Contains(t, reportStr, "checked 4 candidate tokens, 3 with an enabled override for dst 3379446385462418246, 3 updated")

		args := decodeApplyFQTokenTransferFeeConfigUpdates(t, report.Output.BatchOps[0].Transactions[0].Data)
		require.Len(t, args, 1)
		require.Equal(t, tpGasCfgTargetChainSel, args[0].DestChainSelector)
		got := map[common.Address]fqops.TokenTransferFeeConfig{}
		for _, c := range args[0].TokenTransferFeeConfigs {
			got[c.Token] = c.TokenTransferFeeConfig
		}
		require.Len(t, got, 3, "the token with no override must not be written")
		require.NotContains(t, got, fqTokCfgNoOverrideToken)
		require.Equal(t, uint32(750_000), got[usdcToken].DestGasOverhead)
		require.Equal(t, uint32(1_200_000), got[lombardToken].DestGasOverhead)
		require.Equal(t, uint32(450_000), got[fqTokCfgGenericToken].DestGasOverhead)
		for tok, c := range got { // every other field is preserved
			require.Equal(t, uint32(150), c.FeeUSDCents, tok.Hex())
			require.Equal(t, uint32(32), c.DestBytesOverhead, tok.Hex())
			require.True(t, c.IsEnabled, tok.Hex())
		}
	})

	t.Run("mismatched USDC override takes the ratio fallback", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 5548718428018410741: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (USDC) MISMATCH - current value 180000")
		require.Contains(t, reportStr, "applying fallback value 540000 instead of literal Glamsterdam value 750000")
		args := decodeApplyFQTokenTransferFeeConfigUpdates(t, report.Output.BatchOps[1].Transactions[0].Data)
		require.Len(t, args[0].TokenTransferFeeConfigs, 1)
		require.Equal(t, uint32(540_000), args[0].TokenTransferFeeConfigs[0].TokenTransferFeeConfig.DestGasOverhead)
	})

	t.Run("nothing is written directly, always routed through MCMS", func(t *testing.T) {
		cur, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fqops.GetTokenTransferFeeConfig, e.BlockChains.EVMChains()[tpGasCfgBaselineChain], contract.FunctionInput[fqops.GetTokenTransferFeeConfigArgs]{
			ChainSelector: tpGasCfgBaselineChain,
			Address:       fqBaseline,
			Args:          fqops.GetTokenTransferFeeConfigArgs{DestChainSelector: tpGasCfgTargetChainSel, Token: usdcToken},
		})
		require.NoError(t, err)
		require.Equal(t, uint32(250_000), cur.Output.DestGasOverhead)
	})
}

// A re-run after a prior proposal executed must not touch USDC/Lombard overrides that already sit at
// the Glamsterdam value.
func TestUpdateFeeQuoterTokenTransferFeeConfig_AlreadyAppliedIsNoOp(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{tpGasCfgBaselineChain}))
	require.NoError(t, err)

	usdcPool := deployTokenPoolFixture(t, e, tpGasCfgBaselineChain, "usdc", 750_000)
	usdcToken := poolToken(t, e, tpGasCfgBaselineChain, usdcPool)
	fq := deployFQWithTokenOverrides(t, e, tpGasCfgBaselineChain, []fqops.TokenTransferFeeConfigSingleTokenArgs{fqTokCfgOverride(usdcToken, 750_000)})

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateFeeQuoterTokenTransferFeeConfig, e.BlockChains, glamsterdamseq.UpdateFeeQuoterTokenTransferFeeConfigInput{
		TargetChainSelector: tpGasCfgTargetChainSel,
		Lanes: []glamsterdamseq.FeeQuoterTokenConfigLane{{
			ChainSelector:     tpGasCfgBaselineChain,
			FeeQuoterAddress:  fq,
			USDCPoolAddresses: []common.Address{usdcPool},
		}},
	})
	require.NoError(t, err)
	require.Empty(t, report.Output.BatchOps)
	require.Contains(t, report.Output.Report.String(), "(USDC) already matches Glamsterdam value 750000")
}

// A pool that can't be read (e.g. a stale datastore entry with no contract code) must not abort the
// batch: it is only used to recognise a token as USDC, so warn and keep processing the other tokens.
func TestUpdateFeeQuoterTokenTransferFeeConfig_UnreadablePoolIsNotFatal(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{tpGasCfgBaselineChain}))
	require.NoError(t, err)

	stalePool := common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddeaddead") // no code at this address
	fq := deployFQWithTokenOverrides(t, e, tpGasCfgBaselineChain, []fqops.TokenTransferFeeConfigSingleTokenArgs{fqTokCfgOverride(fqTokCfgGenericToken, 150_000)})

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateFeeQuoterTokenTransferFeeConfig, e.BlockChains, glamsterdamseq.UpdateFeeQuoterTokenTransferFeeConfigInput{
		TargetChainSelector: tpGasCfgTargetChainSel,
		Lanes: []glamsterdamseq.FeeQuoterTokenConfigLane{{
			ChainSelector:        tpGasCfgBaselineChain,
			FeeQuoterAddress:     fq,
			ExtraCandidateTokens: []common.Address{fqTokCfgGenericToken},
			USDCPoolAddresses:    []common.Address{stalePool},
			LombardPoolAddresses: []common.Address{stalePool},
		}},
	})
	require.NoError(t, err)
	reportStr := report.Output.Report.String()
	require.Contains(t, reportStr, "its token will not be recognised as USDC")
	require.Contains(t, reportStr, "its token will not be recognised as Lombard")
	require.Len(t, report.Output.BatchOps, 1, "the other token is still migrated")
	require.Contains(t, reportStr, "(token "+fqTokCfgGenericToken.Hex()+") generic scaling 150000 -> 450000")
}
