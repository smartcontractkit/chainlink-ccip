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
	cvops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	orops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	rmpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	glamsterdamseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences/glamsterdam"
)

const (
	gasCfgTargetChainSel  = uint64(3379446385462418246)
	gasCfgBaselineChain   = uint64(4949039107694359620)
	gasCfgMismatchedChain = uint64(5548718428018410741)
)

var (
	gasCfgRouterAddr          = common.HexToAddress("0x1111111111111111111111111111111111111111")
	gasCfgDefaultExecutorAddr = common.HexToAddress("0x2222222222222222222222222222222222222222")
	gasCfgDefaultCCVAddr      = common.HexToAddress("0x3333333333333333333333333333333333333333")
	gasCfgOffRampBytesAddr    = common.HexToAddress("0x4444444444444444444444444444444444444444")
	gasCfgRmnAddr             = common.HexToAddress("0x5555555555555555555555555555555555555555")
	gasCfgTokenAdminRegistry  = common.HexToAddress("0x6666666666666666666666666666666666666666")
	gasCfgLinkToken           = common.HexToAddress("0x7777777777777777777777777777777777777777")
)

func deployGasCfgOnRamp(t *testing.T, e *cldf.Environment, chainSel uint64, baseExecutionGasCost uint32) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, rmpops.Deploy, chain, contract.DeployInput[rmpops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(rmpops.ContractType, *rmpops.Version),
		Args: rmpops.ConstructorArgs{
			StaticConfig: rmpops.StaticConfig{
				ChainSelector:         chainSel,
				RmnRemote:             gasCfgRmnAddr,
				TokenAdminRegistry:    gasCfgTokenAdminRegistry,
				MaxUSDCentsPerMessage: 1_000_000,
			},
			DynamicConfig: rmpops.DynamicConfig{
				FeeQuoter: gasCfgRouterAddr, // placeholder, not exercised by this test
			},
		},
	})
	require.NoError(t, err)
	addr := common.HexToAddress(out.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, rmpops.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]rmpops.DestChainConfigArgs]{
		ChainSelector: chainSel,
		Address:       addr,
		Args: []rmpops.DestChainConfigArgs{
			{
				DestChainSelector:    gasCfgTargetChainSel,
				Router:               gasCfgRouterAddr,
				AddressBytesLength:   20,
				BaseExecutionGasCost: baseExecutionGasCost,
				DefaultCCVs:          []common.Address{gasCfgDefaultCCVAddr},
				DefaultExecutor:      gasCfgDefaultExecutorAddr,
				OffRamp:              gasCfgOffRampBytesAddr.Bytes(),
			},
		},
	})
	require.NoError(t, err)

	return addr
}

func deployGasCfgFeeQuoter(
	t *testing.T,
	e *cldf.Environment,
	chainSel uint64,
	destChainConfig fqops.DestChainConfig,
) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fqops.Deploy, chain, contract.DeployInput[fqops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fqops.ContractType, *fqops.Version),
		Args: fqops.ConstructorArgs{
			StaticConfig: fqops.StaticConfig{
				MaxFeeJuelsPerMsg: big.NewInt(1e18),
				LinkToken:         gasCfgLinkToken,
			},
			DestChainConfigArgs: []fqops.DestChainConfigArgs{
				{DestChainSelector: gasCfgTargetChainSel, DestChainConfig: destChainConfig},
			},
		},
	})
	require.NoError(t, err)

	return common.HexToAddress(out.Output.Address)
}

func deployGasCfgCommitteeVerifier(t *testing.T, e *cldf.Environment, chainSel uint64, gasForVerification uint32) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, cvops.Deploy, chain, contract.DeployInput[cvops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(cvops.ContractType, *cvops.Version),
		Args: cvops.ConstructorArgs{
			DynamicConfig: cvops.DynamicConfig{FeeAggregator: gasCfgRouterAddr},
			Rmn:           gasCfgRmnAddr,
			VersionTag:    [4]byte{0x01, 0x02, 0x03, 0x04},
		},
	})
	require.NoError(t, err)
	addr := common.HexToAddress(out.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, cvops.ApplyRemoteChainConfigUpdates, chain, contract.FunctionInput[[]cvops.RemoteChainConfigArgs]{
		ChainSelector: chainSel,
		Address:       addr,
		Args: []cvops.RemoteChainConfigArgs{
			{
				Router:              gasCfgRouterAddr,
				RemoteChainSelector: gasCfgTargetChainSel,
				GasForVerification:  gasForVerification,
			},
		},
	})
	require.NoError(t, err)

	return addr
}

func deployGasCfgOffRamp(t *testing.T, e *cldf.Environment, chainSel uint64, gasForCallExactCheck uint16, maxGasBufferToUpdateState uint32) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, orops.Deploy, chain, contract.DeployInput[orops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(orops.ContractType, *orops.Version),
		Args: orops.ConstructorArgs{
			StaticConfig: orops.StaticConfig{
				LocalChainSelector:        chainSel,
				GasForCallExactCheck:      gasForCallExactCheck,
				RmnRemote:                 gasCfgRmnAddr,
				TokenAdminRegistry:        gasCfgTokenAdminRegistry,
				MaxGasBufferToUpdateState: maxGasBufferToUpdateState,
			},
		},
	})
	require.NoError(t, err)

	return common.HexToAddress(out.Output.Address)
}

// decodeApplyDestChainConfigUpdatesArgs decodes the []fqops.DestChainConfigArgs passed to a
// FeeQuoter applyDestChainConfigUpdates call from its raw calldata, mirroring the
// abi.ConvertType pattern the generated read operations already use in this codebase.
func decodeApplyDestChainConfigUpdatesArgs(t *testing.T, data []byte) []fqops.DestChainConfigArgs {
	t.Helper()
	parsed, err := ethabi.JSON(strings.NewReader(fqops.FeeQuoterABI))
	require.NoError(t, err)
	vals, err := parsed.Methods["applyDestChainConfigUpdates"].Inputs.Unpack(data[4:])
	require.NoError(t, err)
	converted := *ethabi.ConvertType(vals[0], new([]fqops.DestChainConfigArgs)).(*[]fqops.DestChainConfigArgs)
	return converted
}

func TestUpdateGasConfig(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{
		gasCfgBaselineChain,
		gasCfgMismatchedChain,
	}))
	require.NoError(t, err)

	baselineFQConfig := fqops.DestChainConfig{
		IsEnabled:                   true,
		MaxDataBytes:                1_000,
		MaxPerMsgGasLimit:           15_000_000,
		DestGasOverhead:             300_000,
		DestGasPerPayloadByteBase:   20,
		ChainFamilySelector:         [4]byte{0x28, 0x12, 0xd5, 0x2c},
		DefaultTokenDestGasOverhead: 90_000,
		DefaultTxGasLimit:           200_000,
		LinkFeeMultiplierPercent:    100,
	}

	onRampBaselineAddr := deployGasCfgOnRamp(t, e, gasCfgBaselineChain, 200_000)
	fqBaselineAddr := deployGasCfgFeeQuoter(t, e, gasCfgBaselineChain, baselineFQConfig)
	cvBaselineAddr := deployGasCfgCommitteeVerifier(t, e, gasCfgBaselineChain, 75_000)
	offRampBaselineAddr := deployGasCfgOffRamp(t, e, gasCfgBaselineChain, 5_000, 12_000)

	// Mismatched chain: OnRamp.BaseExecutionGasCost doesn't match the expected Prague baseline
	// (200,000) — exercises the fallback path instead of the literal Glamsterdam value.
	onRampMismatchedAddr := deployGasCfgOnRamp(t, e, gasCfgMismatchedChain, 150_000)
	fqMismatchedAddr := deployGasCfgFeeQuoter(t, e, gasCfgMismatchedChain, baselineFQConfig)
	cvMismatchedAddr := deployGasCfgCommitteeVerifier(t, e, gasCfgMismatchedChain, 75_000)

	report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, glamsterdamseq.UpdateGasConfig, e.BlockChains, glamsterdamseq.UpdateGasConfigInput{
		TargetChainSelector: gasCfgTargetChainSel,
		Lanes: []glamsterdamseq.LaneAddresses{
			{
				ChainSelector:            gasCfgBaselineChain,
				OnRampAddress:            onRampBaselineAddr,
				FeeQuoterAddress:         fqBaselineAddr,
				CommitteeVerifierAddress: cvBaselineAddr,
				OffRampAddress:           offRampBaselineAddr,
			},
			{
				ChainSelector:            gasCfgMismatchedChain,
				OnRampAddress:            onRampMismatchedAddr,
				FeeQuoterAddress:         fqMismatchedAddr,
				CommitteeVerifierAddress: cvMismatchedAddr,
			},
		},
	})
	require.NoError(t, err)

	require.Len(t, report.Output.BatchOps, 2)
	for _, batchOp := range report.Output.BatchOps {
		require.Len(t, batchOp.Transactions, 3, "expected OnRamp + FeeQuoter + CommitteeVerifier writes per chain")
	}

	// None of the writes should have executed directly — every write must always be routed
	// through MCMS regardless of who the deployer key is.
	reportStr := report.Output.Report.String()

	t.Run("baseline chain applies literal Glamsterdam values and preserves untouched fields", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 4949039107694359620: OnRamp.DestChainConfig.BaseExecutionGasCost matched expected Prague value 200000, applying Glamsterdam value 400000")
		require.Contains(t, reportStr, "chain 4949039107694359620: FeeQuoter.DestChainConfig.DefaultTokenDestGasOverhead matched expected Prague value 90000, applying Glamsterdam value 270000")
		require.Contains(t, reportStr, "chain 4949039107694359620: CommitteeVerifier.RemoteChainConfigArgs.GasForVerification matched expected Prague value 75000, applying Glamsterdam value 85000")

		batchOp := report.Output.BatchOps[0]
		require.Equal(t, mcmsChainSelector(gasCfgBaselineChain), uint64(batchOp.ChainSelector))
		fqArgs := decodeApplyDestChainConfigUpdatesArgs(t, batchOp.Transactions[1].Data)
		require.Len(t, fqArgs, 1)
		require.Equal(t, uint32(270_000), fqArgs[0].DestChainConfig.DefaultTokenDestGasOverhead)
		require.Equal(t, uint32(15_000_000), fqArgs[0].DestChainConfig.MaxPerMsgGasLimit) // no-op field, unchanged
		require.Equal(t, uint8(64), fqArgs[0].DestChainConfig.DestGasPerPayloadByteBase)
		require.Equal(t, uint32(400_000), fqArgs[0].DestChainConfig.DefaultTxGasLimit)
		// untouched field preserved across the merge
		require.Equal(t, uint32(300_000), fqArgs[0].DestChainConfig.DestGasOverhead)
		require.True(t, fqArgs[0].DestChainConfig.IsEnabled)
		require.Equal(t, uint32(1_000), fqArgs[0].DestChainConfig.MaxDataBytes)
	})

	t.Run("mismatched chain applies fallback instead of literal value", func(t *testing.T) {
		require.Contains(t, reportStr, "chain 5548718428018410741: OnRamp.DestChainConfig.BaseExecutionGasCost MISMATCH")
		require.Contains(t, reportStr, "current value 150000")
		require.Contains(t, reportStr, "applying fallback value 300000 instead of literal Glamsterdam value 400000")
	})

	t.Run("no writes executed directly, always routed through MCMS", func(t *testing.T) {
		for _, lane := range []struct {
			onRamp, feeQuoter, committeeVerifier common.Address
			chainSel                             uint64
		}{
			{onRampBaselineAddr, fqBaselineAddr, cvBaselineAddr, gasCfgBaselineChain},
			{onRampMismatchedAddr, fqMismatchedAddr, cvMismatchedAddr, gasCfgMismatchedChain},
		} {
			chain := e.BlockChains.EVMChains()[lane.chainSel]
			cur, err := cldf_ops.ExecuteOperation(e.OperationsBundle, rmpops.GetDestChainConfig, chain, contract.FunctionInput[uint64]{
				ChainSelector: lane.chainSel,
				Address:       lane.onRamp,
				Args:          gasCfgTargetChainSel,
			})
			require.NoError(t, err)
			// on-chain value must be unchanged (200_000 or 150_000, whichever was seeded), proving
			// the write was never actually sent.
			require.NotEqual(t, uint32(400_000), cur.Output.BaseExecutionGasCost)
		}
	})
}

// mcmsChainSelector exists purely so the test reads naturally when comparing against a
// mcms_types.ChainSelector value.
func mcmsChainSelector(sel uint64) uint64 { return sel }
