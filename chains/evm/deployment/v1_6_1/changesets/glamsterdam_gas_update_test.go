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
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	orops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/changesets"
	ccipdeploymentutils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

const (
	gg16TargetChainSel  = uint64(3379446385462418246)
	gg16SrcAChainSel    = uint64(4949039107694359620)
	gg16SrcBChainSel    = uint64(5548718428018410741)
	gg16SrcSkipChainSel = uint64(6433500567565415381)
	gg16SrcNoLaneSel    = uint64(4356164186791070119)
)

var (
	gg16RmnAddr            = common.HexToAddress("0x5555555555555555555555555555555555555555")
	gg16TokenAdminRegistry = common.HexToAddress("0x6666666666666666666666666666666666666666")
	gg16NonceManagerAddr   = common.HexToAddress("0x8888888888888888888888888888888888888888")
	gg16USDCToken          = common.HexToAddress("0x9999999999999999999999999999999999999999")
	gg16FeeQuoterFamily    = [4]byte{0x28, 0x12, 0xd5, 0x2c}
)

// gg16FeeQuoterFixture deploys a v1.6 FeeQuoter with a lane to gg16TargetChainSel and registers
// it in ds.
func gg16FeeQuoterFixture(t *testing.T, e *cldf.Environment, ds datastore.MutableDataStore, chainSel uint64, hasLane bool, destGasOverhead uint32, usdcDestGasOverhead uint32) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	var destChainConfigArgs []fq16ops.DestChainConfigArgs
	if hasLane {
		destChainConfigArgs = []fq16ops.DestChainConfigArgs{
			{
				DestChainSelector: gg16TargetChainSel,
				DestChainConfig: fq16ops.DestChainConfig{
					IsEnabled:                   true,
					MaxNumberOfTokensPerMsg:     5,
					MaxPerMsgGasLimit:           3_000_000,
					DestGasOverhead:             destGasOverhead,
					ChainFamilySelector:         gg16FeeQuoterFamily,
					DefaultTokenDestGasOverhead: 90_000,
					DefaultTxGasLimit:           200_000,
					GasMultiplierWeiPerEth:      1e18,
				},
			},
		}
	}

	var tokenTransferFeeConfigArgs []fq16ops.TokenTransferFeeConfigArgs
	if usdcDestGasOverhead > 0 {
		tokenTransferFeeConfigArgs = []fq16ops.TokenTransferFeeConfigArgs{
			{
				DestChainSelector: gg16TargetChainSel,
				TokenTransferFeeConfigs: []fq16ops.TokenTransferFeeConfigSingleTokenArgs{
					{
						Token: gg16USDCToken,
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
		}
	}

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, fq16ops.Deploy, chain, contract.DeployInput[fq16ops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(fq16ops.ContractType, *fq16ops.Version),
		Args: fq16ops.ConstructorArgs{
			StaticConfig: fq16ops.StaticConfig{
				MaxFeeJuelsPerMsg:            big.NewInt(1e18),
				LinkToken:                    common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"),
				TokenPriceStalenessThreshold: 3600,
			},
			DestChainConfigArgs:        destChainConfigArgs,
			TokenTransferFeeConfigArgs: tokenTransferFeeConfigArgs,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(out.Output))

	return common.HexToAddress(out.Output.Address)
}

func gg16OffRampFixture(t *testing.T, e *cldf.Environment, ds datastore.MutableDataStore, chainSel uint64, gasForCallExactCheck uint16) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, orops.Deploy, chain, contract.DeployInput[orops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(orops.ContractType, *orops.Version),
		Args: orops.ConstructorArgs{
			StaticConfig: orops.StaticConfig{
				ChainSelector:        chainSel,
				GasForCallExactCheck: gasForCallExactCheck,
				RmnRemote:            gg16RmnAddr,
				TokenAdminRegistry:   gg16TokenAdminRegistry,
				NonceManager:         gg16NonceManagerAddr,
			},
			DynamicConfig: orops.DynamicConfig{FeeQuoter: gg16RmnAddr},
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(out.Output))
}

// gg16TokenAdminRegistryFixture deploys a TokenAdminRegistry, registers gg16USDCToken as a
// configured token (via proposeAdministrator, which is enough for getAllConfiguredTokens to
// return it), and registers the address in ds.
func gg16TokenAdminRegistryFixture(t *testing.T, e *cldf.Environment, ds datastore.MutableDataStore, chainSel uint64) {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, tarops.Deploy, chain, contract.DeployInput[tarops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(tarops.ContractType, *tarops.Version),
		Args:           tarops.ConstructorArgs{},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(out.Output))
	tarAddr := common.HexToAddress(out.Output.Address)

	_, err = cldf_ops.ExecuteOperation(e.OperationsBundle, tarops.ProposeAdministrator, chain, contract.FunctionInput[tarops.ProposeAdministratorArgs]{
		ChainSelector: chainSel,
		Address:       tarAddr,
		Args: tarops.ProposeAdministratorArgs{
			TokenAddress:  gg16USDCToken,
			Administrator: chain.DeployerKey.From,
		},
	})
	require.NoError(t, err)
}

func TestUpdateGasConfigForGlamsterdamV16(t *testing.T) {
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{
		gg16SrcAChainSel,
		gg16SrcBChainSel,
		gg16SrcSkipChainSel,
		gg16SrcNoLaneSel,
	}))
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	// srcA: lane to target, Prague baseline matches exactly, plus a configured USDC override.
	gg16FeeQuoterFixture(t, e, ds, gg16SrcAChainSel, true, 300_000, 180_000)
	gg16TokenAdminRegistryFixture(t, e, ds, gg16SrcAChainSel)
	gg16OffRampFixture(t, e, ds, gg16SrcAChainSel, 5_000)
	// srcB: lane to target, FeeQuoter.DestGasOverhead mismatched -> exercises fallback. No
	// TokenAdminRegistry on this chain, so the token-transfer-fee lane is skipped for it.
	gg16FeeQuoterFixture(t, e, ds, gg16SrcBChainSel, true, 240_000, 0)
	// srcSkip: has a lane, but is passed in SkipChainSelectors -> must be excluded entirely.
	gg16FeeQuoterFixture(t, e, ds, gg16SrcSkipChainSel, true, 300_000, 0)
	// srcNoLane: FeeQuoter deployed, but no dest chain config for target -> discovered, no lane.
	gg16FeeQuoterFixture(t, e, ds, gg16SrcNoLaneSel, false, 300_000, 0)

	// Only srcA and srcB will end up with batch ops, so only they need a real MCMS+Timelock
	// deployment for the OutputBuilder to resolve chain metadata against.
	for _, chainSel := range []uint64{gg16SrcAChainSel, gg16SrcBChainSel} {
		chain := e.BlockChains.EVMChains()[chainSel]
		_, mcmsAddrs := deployMCMSInstanceForTest(t, e.OperationsBundle, chain, chain.DeployerKey.From, ccipdeploymentutils.CLLQualifier)
		for _, ref := range mcmsAddrs {
			require.NoError(t, ds.Addresses().Add(ref))
		}
	}
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	out, err := changesets.UpdateGasConfigForGlamsterdamV16(mcmsRegistry).Apply(*e, cs_core.WithMCMS[changesets.GlamsterdamGasUpdateV16Cfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.GlamsterdamGasUpdateV16Cfg{
			TargetChainSelector: gg16TargetChainSel,
			SkipChainSelectors:  []uint64{gg16SrcSkipChainSel},
		},
	})
	require.NoError(t, err)

	require.Len(t, out.MCMSTimelockProposals, 1)
	proposal := out.MCMSTimelockProposals[0]

	batchOpChains := make(map[uint64]int)
	for _, op := range proposal.Operations {
		batchOpChains[uint64(op.ChainSelector)]++
	}
	// srcA: 1 FeeQuoter gas-config write + 1 FeeQuoter token-transfer-fee write.
	require.Equal(t, 2, batchOpChains[gg16SrcAChainSel])
	// srcB: 1 FeeQuoter gas-config write only (no TokenAdminRegistry -> no token lane).
	require.Equal(t, 1, batchOpChains[gg16SrcBChainSel])
	require.NotContains(t, batchOpChains, gg16SrcSkipChainSel)
	require.NotContains(t, batchOpChains, gg16SrcNoLaneSel)

	require.Contains(t, proposal.Description, "chain 6433500567565415381: skipped (explicit SkipChainSelectors entry)")
	require.Contains(t, proposal.Description, "chain 4356164186791070119: no lane to target chain, skipped")
	require.Contains(t, proposal.Description, "chain 4949039107694359620: FeeQuoter.DestChainConfig.DestGasOverhead matched expected Prague value 300000, applying Glamsterdam value 500000")
	require.Contains(t, proposal.Description, "chain 4949039107694359620: FeeQuoter.TokenTransferFeeConfig.DestGasOverhead (USDC) matched expected Prague value 180000, applying Glamsterdam value 540000")
	require.Contains(t, proposal.Description, "chain 5548718428018410741: FeeQuoter.DestChainConfig.DestGasOverhead MISMATCH")
	require.Contains(t, proposal.Description, "applying fallback value 400000 instead of literal Glamsterdam value 500000")
	require.Contains(t, proposal.Description, "chain 5548718428018410741: ERROR - could not resolve TokenAdminRegistry address, skipping this chain")
}
