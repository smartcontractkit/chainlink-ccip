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
	orops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/sequences/glamsterdam"
)

const (
	gasCfgTargetChainSel  = uint64(3379446385462418246)
	gasCfgBaselineChain   = uint64(4949039107694359620)
	gasCfgMismatchedChain = uint64(5548718428018410741)
)

var (
	gasCfgRmnAddr            = common.HexToAddress("0x5555555555555555555555555555555555555555")
	gasCfgTokenAdminRegistry = common.HexToAddress("0x6666666666666666666666666666666666666666")
	gasCfgNonceManagerAddr   = common.HexToAddress("0x8888888888888888888888888888888888888888")
)

func deployGasCfgFeeQuoter(
	t *testing.T,
	e *cldf.Environment,
	chainSel uint64,
	destChainConfig fq16ops.DestChainConfig,
) common.Address {
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
			DestChainConfigArgs: []fq16ops.DestChainConfigArgs{
				{DestChainSelector: gasCfgTargetChainSel, DestChainConfig: destChainConfig},
			},
		},
	})
	require.NoError(t, err)

	return common.HexToAddress(out.Output.Address)
}

func deployGasCfgOffRamp(t *testing.T, e *cldf.Environment, chainSel uint64, gasForCallExactCheck uint16) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, orops.Deploy, chain, contract.DeployInput[orops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(orops.ContractType, *orops.Version),
		Args: orops.ConstructorArgs{
			StaticConfig: orops.StaticConfig{
				ChainSelector:        chainSel,
				GasForCallExactCheck: gasForCallExactCheck,
				RmnRemote:            gasCfgRmnAddr,
				TokenAdminRegistry:   gasCfgTokenAdminRegistry,
				NonceManager:         gasCfgNonceManagerAddr,
			},
			DynamicConfig: orops.DynamicConfig{
				FeeQuoter: gasCfgRmnAddr, // placeholder, not exercised by this test
			},
		},
	})
	require.NoError(t, err)

	return common.HexToAddress(out.Output.Address)
}

// decodeApplyDestChainConfigUpdatesArgs decodes the []fq16ops.DestChainConfigArgs passed to a
// FeeQuoter applyDestChainConfigUpdates call from its raw calldata, mirroring the abi.ConvertType
// pattern the generated read operations already use in this codebase.
func decodeApplyDestChainConfigUpdatesArgs(t *testing.T, data []byte) []fq16ops.DestChainConfigArgs {
	t.Helper()
	parsed, err := ethabi.JSON(strings.NewReader(fq16ops.FeeQuoterABI))
	require.NoError(t, err)
	vals, err := parsed.Methods["applyDestChainConfigUpdates"].Inputs.Unpack(data[4:])
	require.NoError(t, err)
	converted := *ethabi.ConvertType(vals[0], new([]fq16ops.DestChainConfigArgs)).(*[]fq16ops.DestChainConfigArgs)
	return converted
}

func TestUpdateGasConfig(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{
		gasCfgBaselineChain,
		gasCfgMismatchedChain,
	}))
	require.NoError(t, err)

	baselineFQConfig := fq16ops.DestChainConfig{
		IsEnabled:                   true,
		MaxNumberOfTokensPerMsg:     5,
		MaxPerMsgGasLimit:           3_000_000,
		DestGasOverhead:             300_000,
		ChainFamilySelector:         evmFamilySelector,
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		GasMultiplierWeiPerEth:      1e18,
	}

	fqBaselineAddr := deployGasCfgFeeQuoter(t, e, gasCfgBaselineChain, baselineFQConfig)
	offRampBaselineAddr := deployGasCfgOffRamp(t, e, gasCfgBaselineChain, 5_000)

	// Mismatched chain: FeeQuoter.DestGasOverhead doesn't match the expected Prague baseline
	// (300,000) — exercises the fallback path instead of the literal Glamsterdam value.
	mismatchedFQConfig := baselineFQConfig
	mismatchedFQConfig.DestGasOverhead = 240_000
	fqMismatchedAddr := deployGasCfgFeeQuoter(t, e, gasCfgMismatchedChain, mismatchedFQConfig)

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateGasConfig, e.BlockChains, glamsterdamseq.UpdateGasConfigInput{
		TargetChainSelector: gasCfgTargetChainSel,
		Lanes: []glamsterdamseq.LaneAddresses{
			{
				ChainSelector:    gasCfgBaselineChain,
				FeeQuoterAddress: fqBaselineAddr,
				OffRampAddress:   offRampBaselineAddr,
			},
			{
				ChainSelector:    gasCfgMismatchedChain,
				FeeQuoterAddress: fqMismatchedAddr,
			},
		},
	})
	require.NoError(t, err)

	require.Len(t, report.Output.BatchOps, 2)
	for _, batchOp := range report.Output.BatchOps {
		require.Len(t, batchOp.Transactions, 1, "expected a single FeeQuoter write per chain")
	}

	reportStr := report.Output.Report.String()

	t.Run("baseline chain applies literal Glamsterdam values and preserves untouched fields", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 4949039107694359620: FeeQuoter.DestChainConfig.DestGasOverhead matched expected Prague value 300000, applying Glamsterdam value 500000")
		require.Contains(t, reportStr, "chain 4949039107694359620: FeeQuoter.DestChainConfig.DefaultTokenDestGasOverhead matched expected Prague value 90000, applying Glamsterdam value 270000")

		batchOp := report.Output.BatchOps[0]
		fqArgs := decodeApplyDestChainConfigUpdatesArgs(t, batchOp.Transactions[0].Data)
		require.Len(t, fqArgs, 1)
		require.Equal(t, uint32(500_000), fqArgs[0].DestChainConfig.DestGasOverhead)
		require.Equal(t, uint32(270_000), fqArgs[0].DestChainConfig.DefaultTokenDestGasOverhead)
		// untouched fields preserved across the merge
		require.Equal(t, uint32(3_000_000), fqArgs[0].DestChainConfig.MaxPerMsgGasLimit)
		require.True(t, fqArgs[0].DestChainConfig.IsEnabled)
		require.Equal(t, uint16(5), fqArgs[0].DestChainConfig.MaxNumberOfTokensPerMsg)
	})

	t.Run("mismatched chain applies fallback instead of literal value", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 5548718428018410741: FeeQuoter.DestChainConfig.DestGasOverhead MISMATCH")
		require.Contains(t, reportStr, "current value 240000")
		require.Contains(t, reportStr, "applying fallback value 400000 instead of literal Glamsterdam value 500000")
	})

	t.Run("no writes executed directly, always routed through MCMS", func(t *testing.T) {
		for _, lane := range []struct {
			feeQuoter common.Address
			chainSel  uint64
		}{
			{fqBaselineAddr, gasCfgBaselineChain},
			{fqMismatchedAddr, gasCfgMismatchedChain},
		} {
			chain := e.BlockChains.EVMChains()[lane.chainSel]
			cur, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq16ops.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: lane.chainSel,
				Address:       lane.feeQuoter,
				Args:          gasCfgTargetChainSel,
			})
			require.NoError(t, err)
			require.NotEqual(t, uint32(500_000), cur.Output.DestGasOverhead)
		}
	})
}
