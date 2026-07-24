package changesets_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	bnmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	cvops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	rmpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	usdcpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	tpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	ccipdeploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

const (
	ggTargetChainSel  = uint64(3379446385462418246)
	ggSrcAChainSel    = uint64(4949039107694359620)
	ggSrcBChainSel    = uint64(5548718428018410741)
	ggSrcSkipChainSel = uint64(6433500567565415381)
	ggSrcNoLaneSel    = uint64(4356164186791070119)
)

var (
	ggRouterAddr          = common.HexToAddress("0x1111111111111111111111111111111111111111")
	ggDefaultExecutorAddr = common.HexToAddress("0x2222222222222222222222222222222222222222")
	ggDefaultCCVAddr      = common.HexToAddress("0x3333333333333333333333333333333333333333")
	ggOffRampBytesAddr    = common.HexToAddress("0x4444444444444444444444444444444444444444")
	ggRmnAddr             = common.HexToAddress("0x5555555555555555555555555555555555555555")
	ggTokenAdminRegistry  = common.HexToAddress("0x6666666666666666666666666666666666666666")
	ggLinkToken           = common.HexToAddress("0x7777777777777777777777777777777777777777")
	ggEVMFamilySelector   = [4]byte{0x28, 0x12, 0xd5, 0x2c}
)

// ggLaneFixture deploys FeeQuoter, OnRamp, and CommitteeVerifier on chainSel with a lane pointed
// at ggTargetChainSel, registers them in ds, and returns whether a lane was actually configured
// (hasLane=false leaves the FeeQuoter's dest chain config for the target disabled, so discovery
// finds no lane).
func ggLaneFixture(t *testing.T, e *cldf.Environment, ds datastore.MutableDataStore, chainSel uint64, hasLane bool, onRampBaseExecGasCost uint32) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	fqDestChainConfigArgs := []fqops.DestChainConfigArgs{}
	if hasLane {
		fqDestChainConfigArgs = []fqops.DestChainConfigArgs{
			{
				DestChainSelector: ggTargetChainSel,
				DestChainConfig: fqops.DestChainConfig{
					IsEnabled:                   true,
					MaxPerMsgGasLimit:           15_000_000,
					DestGasOverhead:             300_000,
					DestGasPerPayloadByteBase:   20,
					ChainFamilySelector:         ggEVMFamilySelector,
					DefaultTokenDestGasOverhead: 90_000,
					DefaultTxGasLimit:           200_000,
					LinkFeeMultiplierPercent:    100,
				},
			},
		}
	}

	fqOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fqops.Deploy, chain, contract.DeployInput[fqops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fqops.ContractType, *fqops.Version),
		Args: fqops.ConstructorArgs{
			StaticConfig:        fqops.StaticConfig{MaxFeeJuelsPerMsg: big.NewInt(1e18), LinkToken: ggLinkToken},
			DestChainConfigArgs: fqDestChainConfigArgs,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(fqOut.Output))
	fqAddr := common.HexToAddress(fqOut.Output.Address)

	onRampOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, rmpops.Deploy, chain, contract.DeployInput[rmpops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(rmpops.ContractType, *rmpops.Version),
		Args: rmpops.ConstructorArgs{
			StaticConfig: rmpops.StaticConfig{
				ChainSelector:         chainSel,
				RmnRemote:             ggRmnAddr,
				TokenAdminRegistry:    ggTokenAdminRegistry,
				MaxUSDCentsPerMessage: 1_000_000,
			},
			DynamicConfig: rmpops.DynamicConfig{FeeQuoter: fqAddr},
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(onRampOut.Output))
	onRampAddr := common.HexToAddress(onRampOut.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, rmpops.ApplyDestChainConfigUpdates, chain, contract.FunctionInput[[]rmpops.DestChainConfigArgs]{
		ChainSelector: chainSel,
		Address:       onRampAddr,
		Args: []rmpops.DestChainConfigArgs{
			{
				DestChainSelector:    ggTargetChainSel,
				Router:               ggRouterAddr,
				AddressBytesLength:   20,
				BaseExecutionGasCost: onRampBaseExecGasCost,
				DefaultCCVs:          []common.Address{ggDefaultCCVAddr},
				DefaultExecutor:      ggDefaultExecutorAddr,
				OffRamp:              ggOffRampBytesAddr.Bytes(),
			},
		},
	})
	require.NoError(t, err)

	cvOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, cvops.Deploy, chain, contract.DeployInput[cvops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(cvops.ContractType, *cvops.Version),
		Args: cvops.ConstructorArgs{
			DynamicConfig: cvops.DynamicConfig{FeeAggregator: ggRouterAddr},
			Rmn:           ggRmnAddr,
			VersionTag:    [4]byte{0x01, 0x02, 0x03, 0x04},
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(cvOut.Output))
	cvAddr := common.HexToAddress(cvOut.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, cvops.ApplyRemoteChainConfigUpdates, chain, contract.FunctionInput[[]cvops.RemoteChainConfigArgs]{
		ChainSelector: chainSel,
		Address:       cvAddr,
		Args: []cvops.RemoteChainConfigArgs{
			{Router: ggRouterAddr, RemoteChainSelector: ggTargetChainSel, GasForVerification: 75_000},
		},
	})
	require.NoError(t, err)
}

// ggUSDCPoolFixture deploys a SiloedUSDCTokenPool with a lane to ggTargetChainSel and registers
// it in ds, standing in for the "USDCTokenPool" table row.
func ggUSDCPoolFixture(t *testing.T, e *cldf.Environment, ds datastore.MutableDataStore, chainSel uint64) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	tokenOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, bnmops.Deploy, chain, contract.DeployInput[bnmops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(bnmops.ContractType, *ccipdeploymentutils.Version_1_0_0),
		Args: bnmops.ConstructorArgs{
			Name: "TEST", Symbol: "TEST", Decimals: 6, MaxSupply: big.NewInt(0), PreMint: big.NewInt(0),
		},
	})
	require.NoError(t, err)
	tokenAddr := common.HexToAddress(tokenOut.Output.Address)

	poolOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, usdcpoolops.Deploy, chain, contract.DeployInput[usdcpoolops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(usdcpoolops.ContractType, *usdcpoolops.Version),
		Args: usdcpoolops.ConstructorArgs{
			Token:              tokenAddr,
			LocalTokenDecimals: 6,
			AdvancedPoolHooks:  common.Address{},
			RmnProxy:           ggRmnAddr,
			Router:             ggRouterAddr,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(poolOut.Output))
	poolAddr := common.HexToAddress(poolOut.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, tpops.ApplyChainUpdates, chain, contract.FunctionInput[tpops.ApplyChainUpdatesArgs]{
		ChainSelector: chainSel,
		Address:       poolAddr,
		Args: tpops.ApplyChainUpdatesArgs{
			ChainsToAdd: []tpops.ChainUpdate{
				{
					RemoteChainSelector:       ggTargetChainSel,
					RemoteTokenAddress:        ggRmnAddr.Bytes(),
					OutboundRateLimiterConfig: tpops.Config{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
					InboundRateLimiterConfig:  tpops.Config{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
				},
			},
		},
	})
	require.NoError(t, err)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, tpops.ApplyTokenTransferFeeConfigUpdates, chain, contract.FunctionInput[tpops.ApplyTokenTransferFeeConfigUpdatesArgs]{
		ChainSelector: chainSel,
		Address:       poolAddr,
		Args: tpops.ApplyTokenTransferFeeConfigUpdatesArgs{
			TokenTransferFeeConfigArgs: []tpops.TokenTransferFeeConfigArgs{
				{
					DestChainSelector:      ggTargetChainSel,
					TokenTransferFeeConfig: tpops.TokenTransferFeeConfig{DestGasOverhead: 250_000, IsEnabled: true},
				},
			},
		},
	})
	require.NoError(t, err)
}

func TestUpdateGasConfigForGlamsterdamV2(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{
		ggSrcAChainSel,
		ggSrcBChainSel,
		ggSrcSkipChainSel,
		ggSrcNoLaneSel,
	}))
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	// srcA: lane to target, Prague baseline matches exactly.
	ggLaneFixture(t, e, ds, ggSrcAChainSel, true, 200_000)
	ggUSDCPoolFixture(t, e, ds, ggSrcAChainSel)
	// srcB: lane to target, OnRamp.BaseExecutionGasCost mismatched -> exercises fallback.
	ggLaneFixture(t, e, ds, ggSrcBChainSel, true, 150_000)
	// srcSkip: has a lane, but is passed in SkipChainSelectors -> must be excluded entirely.
	ggLaneFixture(t, e, ds, ggSrcSkipChainSel, true, 200_000)
	// srcNoLane: FeeQuoter deployed, but no dest chain config for target -> discovered, no lane.
	ggLaneFixture(t, e, ds, ggSrcNoLaneSel, false, 200_000)

	// Only srcA and srcB will end up with batch ops, so only they need a real MCMS+Timelock
	// deployment for the OutputBuilder to resolve chain metadata against.
	for _, chainSel := range []uint64{ggSrcAChainSel, ggSrcBChainSel} {
		chain := e.BlockChains.EVMChains()[chainSel]
		_, mcmsAddrs := deployMCMSInstanceForTest(t, e.OperationsBundle, chain, chain.DeployerKey.From, ccipdeploymentutils.CLLQualifier)
		for _, ref := range mcmsAddrs {
			require.NoError(t, ds.Addresses().Add(ref))
		}
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.UpdateGasConfigForGlamsterdamV2(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.GlamsterdamGasUpdateCfg{
			TargetChainSelector: ggTargetChainSel,
			SkipChainSelectors:  []uint64{ggSrcSkipChainSel},
		},
	})
	require.NoError(t, err)

	require.Len(t, out.MCMSTimelockProposals, 1)
	proposal := out.MCMSTimelockProposals[0]

	batchOpChains := make(map[uint64]int)
	for _, op := range proposal.Operations {
		batchOpChains[uint64(op.ChainSelector)]++
	}
	// srcA: 3 gas-config writes (OnRamp+FeeQuoter+CommitteeVerifier) + 1 USDC token-pool write.
	require.Equal(t, 2, batchOpChains[ggSrcAChainSel])
	// srcB: 3 gas-config writes, no token pool.
	require.Equal(t, 1, batchOpChains[ggSrcBChainSel])
	// srcSkip must never appear: it was unconditionally skipped.
	require.NotContains(t, batchOpChains, ggSrcSkipChainSel)
	// srcNoLane must never appear: discovery found no lane to target.
	require.NotContains(t, batchOpChains, ggSrcNoLaneSel)

	require.Contains(t, proposal.Description, "chain 6433500567565415381: skipped (explicit SkipChainSelectors entry)")
	require.Contains(t, proposal.Description, "chain 4356164186791070119: no lane to target chain, skipped")
	require.Contains(t, proposal.Description, "chain 4949039107694359620: OnRamp.DestChainConfig.BaseExecutionGasCost matched expected Prague value 200000, applying Glamsterdam value 400000")
	require.Contains(t, proposal.Description, "chain 5548718428018410741: OnRamp.DestChainConfig.BaseExecutionGasCost MISMATCH")
	require.Contains(t, proposal.Description, "applying fallback value 300000 instead of literal Glamsterdam value 400000")
	require.Contains(t, proposal.Description, "chain 4949039107694359620: TokenPool.TokenTransferFeeConfig.DestGasOverhead (USDC) matched expected Prague value 250000, applying Glamsterdam value 750000")
}
