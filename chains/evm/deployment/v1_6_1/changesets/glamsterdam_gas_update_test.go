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
	usdcerc20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	rmnproxyops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	orops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/offramp"
	usdcpoolops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/burn_mint_with_lock_release_flag_token_pool"
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
	gg16FeeQuoterFamily    = [4]byte{0x28, 0x12, 0xd5, 0x2c}
)

// gg16FeeQuoterFixture deploys a v1.6 FeeQuoter with a lane to gg16TargetChainSel and registers
// it in ds. usdcToken is only used (and must be non-zero) when usdcDestGasOverhead > 0.
func gg16FeeQuoterFixture(t *testing.T, e *cldf.Environment, ds datastore.MutableDataStore, chainSel uint64, hasLane bool, destGasOverhead uint32, usdcDestGasOverhead uint32, usdcToken common.Address) common.Address {
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
						Token: usdcToken,
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

// gg16USDCTokenPoolFixture deploys a real BurnMintERC20 token plus a
// BurnMintWithLockReleaseFlagTokenPool wrapping it, and registers both in ds. Per v1.6.1's
// non-canonical-USDC convention, this pool type is what the changeset resolves per chain (via
// resolveUSDCTokenPoolRef) to discover "the" USDC token address by calling getToken() on it. It
// returns the deployed token's address, which callers must also use when configuring the
// FeeQuoter's TokenTransferFeeConfig override for this chain, so the two agree on which token is
// "USDC".
func gg16USDCTokenPoolFixture(t *testing.T, e *cldf.Environment, ds datastore.MutableDataStore, chainSel uint64) common.Address {
	t.Helper()
	chain := e.BlockChains.EVMChains()[chainSel]

	// The pool constructor needs a real token, RMNProxy, and Router contract deployed on-chain
	// (not just bare addresses), mirroring sequences.setupNonCanonicalTestEnvironment in
	// v1_6_1/sequences/non_canonical_usdc_test.go.
	tokenOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, usdcerc20ops.Deploy, chain, contract.DeployInput[usdcerc20ops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(usdcerc20ops.ContractType, *ccipdeploymentutils.Version_1_0_0),
		Args: usdcerc20ops.ConstructorArgs{
			Name:      "USD Coin",
			Symbol:    "USDC",
			Decimals:  6,
			MaxSupply: big.NewInt(0),
			PreMint:   big.NewInt(0),
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(tokenOut.Output))
	usdcTokenAddr := common.HexToAddress(tokenOut.Output.Address)

	rmnOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, rmnproxyops.Deploy, chain, contract.DeployInput[rmnproxyops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(rmnproxyops.ContractType, *rmnproxyops.Version),
		Args:           rmnproxyops.ConstructorArgs{RMN: chain.DeployerKey.From},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(rmnOut.Output))
	rmnProxyAddr := common.HexToAddress(rmnOut.Output.Address)

	routerOut, err := cldf_ops.ExecuteOperation(e.OperationsBundle, routerops.Deploy, chain, contract.DeployInput[routerops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(routerops.ContractType, *routerops.Version),
		Args: routerops.ConstructorArgs{
			WrappedNative: common.Address{},
			RMNProxy:      rmnProxyAddr,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(routerOut.Output))
	routerAddr := common.HexToAddress(routerOut.Output.Address)

	out, err := cldf_ops.ExecuteOperation(e.OperationsBundle, usdcpoolops.Deploy, chain, contract.DeployInput[usdcpoolops.ConstructorArgs]{
		ChainSelector:  chainSel,
		TypeAndVersion: cldf.NewTypeAndVersion(usdcpoolops.ContractType, *usdcpoolops.Version),
		Args: usdcpoolops.ConstructorArgs{
			Token:              usdcTokenAddr,
			LocalTokenDecimals: 6,
			RmnProxy:           rmnProxyAddr,
			Router:             routerAddr,
		},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Addresses().Add(out.Output))

	return usdcTokenAddr
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
	// The USDC pool must be deployed first so its token address can be used to configure the
	// FeeQuoter's TokenTransferFeeConfig override for the same token.
	usdcTokenA := gg16USDCTokenPoolFixture(t, e, ds, gg16SrcAChainSel)
	gg16FeeQuoterFixture(t, e, ds, gg16SrcAChainSel, true, 300_000, 180_000, usdcTokenA)
	gg16OffRampFixture(t, e, ds, gg16SrcAChainSel, 5_000)
	// srcB: lane to target, FeeQuoter.DestGasOverhead mismatched -> exercises fallback. No USDC
	// token pool on this chain, so the token-transfer-fee lane is skipped for it.
	gg16FeeQuoterFixture(t, e, ds, gg16SrcBChainSel, true, 240_000, 0, common.Address{})
	// srcSkip: has a lane, but is passed in SkipChainSelectors -> must be excluded entirely.
	gg16FeeQuoterFixture(t, e, ds, gg16SrcSkipChainSel, true, 300_000, 0, common.Address{})
	// srcNoLane: FeeQuoter deployed, but no dest chain config for target -> discovered, no lane.
	gg16FeeQuoterFixture(t, e, ds, gg16SrcNoLaneSel, false, 300_000, 0, common.Address{})

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
	require.NotContains(t, proposal.Description, "TokenTransferFeeConfig.DestGasOverhead (USDC)) matched")
}
