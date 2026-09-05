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
	bnmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	usdcpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	tpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/glamsterdam"
	ccipdeploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

const (
	tpGasCfgTargetChainSel   = uint64(3379446385462418246)
	tpGasCfgBaselineChain    = uint64(4949039107694359620)
	tpGasCfgMismatchedChain  = uint64(5548718428018410741)
	tpGasCfgLombardFieldName = "TokenPool.TokenTransferFeeConfig.DestGasOverhead (Lombard)"
	tpGasCfgUSDCFieldName    = "TokenPool.TokenTransferFeeConfig.DestGasOverhead (USDC)"
)

var (
	tpGasCfgRmnProxyAddr = common.HexToAddress("0x3333333333333333333333333333333333333333")
	tpGasCfgRouterAddr   = common.HexToAddress("0x4444444444444444444444444444444444444444")
	tpGasCfgRemoteToken  = common.HexToAddress("0x5555555555555555555555555555555555555555").Bytes()
)

// deployTokenPoolFixture deploys a SiloedUSDCTokenPool as a stand-in fixture for both the Lombard
// and USDC "slots" in this test: the sequence under test only calls the shared TokenPool-ABI
// methods (applyChainUpdates/applyTokenTransferFeeConfigUpdates/getTokenTransferFeeConfig), which
// both concrete pool types inherit unmodified from the same TokenPool base contract, so a single
// concrete deployable stand-in is sufficient to exercise the sequence's logic for both lanes.
func deployTokenPoolFixture(t *testing.T, e *cldf.Environment, chainSel uint64, qualifier string, destGasOverhead uint32) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	tokenOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, bnmops.Deploy, chain, contract.DeployInput[bnmops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(bnmops.ContractType, *ccipdeploymentutils.Version_1_0_0),
		Qualifier:      &qualifier,
		Args: bnmops.ConstructorArgs{
			Name:      "TEST",
			Symbol:    "TEST",
			Decimals:  6,
			MaxSupply: big.NewInt(0),
			PreMint:   big.NewInt(0),
		},
	})
	require.NoError(t, err)
	tokenAddr := common.HexToAddress(tokenOut.Output.Address)

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, usdcpoolops.Deploy, chain, contract.DeployInput[usdcpoolops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(usdcpoolops.ContractType, *usdcpoolops.Version),
		Qualifier:      &qualifier,
		Args: usdcpoolops.ConstructorArgs{
			Token:              tokenAddr,
			LocalTokenDecimals: 6,
			AdvancedPoolHooks:  common.Address{},
			RmnProxy:           tpGasCfgRmnProxyAddr,
			Router:             tpGasCfgRouterAddr,
		},
	})
	require.NoError(t, err)
	addr := common.HexToAddress(out.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, tpops.ApplyChainUpdates, chain, contract.FunctionInput[tpops.ApplyChainUpdatesArgs]{
		ChainSelector: chainSel,
		Address:       addr,
		Args: tpops.ApplyChainUpdatesArgs{
			ChainsToAdd: []tpops.ChainUpdate{
				{
					RemoteChainSelector: tpGasCfgTargetChainSel,
					RemoteTokenAddress:  tpGasCfgRemoteToken,
					OutboundRateLimiterConfig: tpops.Config{
						IsEnabled: false,
						Capacity:  big.NewInt(0),
						Rate:      big.NewInt(0),
					},
					InboundRateLimiterConfig: tpops.Config{
						IsEnabled: false,
						Capacity:  big.NewInt(0),
						Rate:      big.NewInt(0),
					},
				},
			},
		},
	})
	require.NoError(t, err)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, tpops.ApplyTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[tpops.ApplyTokenTransferFeeConfigUpdatesArgs]{
		ChainSelector: chainSel,
		Address:       addr,
		Args: tpops.ApplyTokenTransferFeeConfigUpdatesArgs{
			TokenTransferFeeConfigArgs: []tpops.TokenTransferFeeConfigArgs{
				{
					DestChainSelector: tpGasCfgTargetChainSel,
					TokenTransferFeeConfig: tpops.TokenTransferFeeConfig{
						DestGasOverhead:   destGasOverhead,
						DestBytesOverhead: 100,
						IsEnabled:         true,
					},
				},
			},
		},
	})
	require.NoError(t, err)

	return addr
}

func decodeApplyTokenTransferFeeConfigUpdatesArgs(t *testing.T, data []byte) tpops.ApplyTokenTransferFeeConfigUpdatesArgs {
	t.Helper()
	parsed, err := ethabi.JSON(strings.NewReader(tpops.TokenPoolABI))
	require.NoError(t, err)
	vals, err := parsed.Methods["applyTokenTransferFeeConfigUpdates"].Inputs.Unpack(data[4:])
	require.NoError(t, err)
	feeConfigArgs := *ethabi.ConvertType(vals[0], new([]tpops.TokenTransferFeeConfigArgs)).(*[]tpops.TokenTransferFeeConfigArgs)
	disableArgs := *ethabi.ConvertType(vals[1], new([]uint64)).(*[]uint64)
	return tpops.ApplyTokenTransferFeeConfigUpdatesArgs{
		TokenTransferFeeConfigArgs:     feeConfigArgs,
		DisableTokenTransferFeeConfigs: disableArgs,
	}
}

func TestUpdateTokenPoolGasConfig(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{
		tpGasCfgBaselineChain,
		tpGasCfgMismatchedChain,
	}))
	require.NoError(t, err)

	lombardBaselineAddr := deployTokenPoolFixture(t, e, tpGasCfgBaselineChain, "lombard", 410_000)
	usdcBaselineAddr := deployTokenPoolFixture(t, e, tpGasCfgBaselineChain, "usdc", 250_000)

	// Mismatched chain: USDC pool's current DestGasOverhead doesn't match the expected Prague
	// baseline (250,000) — exercises the fallback path instead of the literal Glamsterdam value.
	lombardMismatchedAddr := deployTokenPoolFixture(t, e, tpGasCfgMismatchedChain, "lombard", 410_000)
	usdcMismatchedAddr := deployTokenPoolFixture(t, e, tpGasCfgMismatchedChain, "usdc", 200_000)

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateTokenPoolGasConfig, e.BlockChains, glamsterdamseq.UpdateTokenPoolGasConfigInput{
		TargetChainSelector: tpGasCfgTargetChainSel,
		LombardPools: []glamsterdamseq.TokenPoolLane{
			{ChainSelector: tpGasCfgBaselineChain, PoolAddress: lombardBaselineAddr},
			{ChainSelector: tpGasCfgMismatchedChain, PoolAddress: lombardMismatchedAddr},
		},
		USDCPools: []glamsterdamseq.TokenPoolLane{
			{ChainSelector: tpGasCfgBaselineChain, PoolAddress: usdcBaselineAddr},
			{ChainSelector: tpGasCfgMismatchedChain, PoolAddress: usdcMismatchedAddr},
		},
	})
	require.NoError(t, err)

	require.Len(t, report.Output.BatchOps, 4)
	for _, batchOp := range report.Output.BatchOps {
		require.Len(t, batchOp.Transactions, 1)
	}

	reportStr := report.Output.Report.String()

	t.Run("baseline chain applies literal Glamsterdam values", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 4949039107694359620: "+tpGasCfgLombardFieldName+" matched expected Prague value 410000, applying Glamsterdam value 1200000")
		require.Contains(t, reportStr, "chain 4949039107694359620: "+tpGasCfgUSDCFieldName+" matched expected Prague value 250000, applying Glamsterdam value 750000")

		lombardArgs := decodeApplyTokenTransferFeeConfigUpdatesArgs(t, report.Output.BatchOps[0].Transactions[0].Data)
		require.Len(t, lombardArgs.TokenTransferFeeConfigArgs, 1)
		require.Equal(t, uint32(1_200_000), lombardArgs.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfig.DestGasOverhead)
		require.Equal(t, uint32(100), lombardArgs.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfig.DestBytesOverhead) // preserved untouched field
		require.True(t, lombardArgs.TokenTransferFeeConfigArgs[0].TokenTransferFeeConfig.IsEnabled)
	})

	t.Run("mismatched chain applies fallback instead of literal value", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 5548718428018410741: "+tpGasCfgUSDCFieldName+" MISMATCH")
		require.Contains(t, reportStr, "current value 200000")
		require.Contains(t, reportStr, "applying fallback value 600000 instead of literal Glamsterdam value 750000")
	})

	t.Run("no writes executed directly, always routed through MCMS", func(t *testing.T) {
		for _, addr := range []common.Address{lombardBaselineAddr, usdcBaselineAddr, lombardMismatchedAddr, usdcMismatchedAddr} {
			chainSel := tpGasCfgBaselineChain
			if addr == lombardMismatchedAddr || addr == usdcMismatchedAddr {
				chainSel = tpGasCfgMismatchedChain
			}
			chain := e.BlockChains.EVMChains()[chainSel]
			cur, err := cldf_ops.ExecuteOperation(e.OperationsBundle, tpops.GetTokenTransferFeeConfig, chain, contract.FunctionInput[tpops.GetTokenTransferFeeConfigArgs]{
				ChainSelector: chainSel,
				Address:       addr,
				Args:          tpops.GetTokenTransferFeeConfigArgs{DestChainSelector: tpGasCfgTargetChainSel},
			})
			require.NoError(t, err)
			require.NotEqual(t, uint32(1_200_000), cur.Output.DestGasOverhead)
			require.NotEqual(t, uint32(750_000), cur.Output.DestGasOverhead)
		}
	})
}
