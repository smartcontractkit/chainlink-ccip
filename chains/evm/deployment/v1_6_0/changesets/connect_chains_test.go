package changesets_test

import (
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	fdeployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

func checkBidirectionalLaneConnectivity(
	t *testing.T,
	e *fdeployment.Environment,
	chainOne lanesapi.ChainDefinition,
	chainTwo lanesapi.ChainDefinition,
	srcAdapter lanesapi.LaneAdapter,
	destAdapter lanesapi.LaneAdapter,
	testRouter bool,
	disable bool,
) {
	type laneDefinition struct {
		Source lanesapi.ChainDefinition
		Dest   lanesapi.ChainDefinition
	}
	lanes := []laneDefinition{
		{
			Source: chainOne,
			Dest:   chainTwo,
		},
		{
			Source: chainTwo,
			Dest:   chainOne,
		},
	}
	for _, lane := range lanes {
		onRampSrcAddr, err := srcAdapter.GetOnRampAddress(e.DataStore, lane.Source.Selector)
		require.NoError(t, err, "must get onRamp from srcAdapter")
		onRampSrc, err := onramp.NewOnRamp(common.BytesToAddress(onRampSrcAddr), e.BlockChains.EVMChains()[lane.Source.Selector].Client)
		require.NoError(t, err, "must instantiate onRamp")

		onRampDestAddr, err := destAdapter.GetOnRampAddress(e.DataStore, lane.Dest.Selector)
		require.NoError(t, err, "must get onRamp from destAdapter")

		offRampDestAddr, err := destAdapter.GetOffRampAddress(e.DataStore, lane.Dest.Selector)
		require.NoError(t, err, "must get offRamp from destAdapter")
		offRampDest, err := offramp.NewOffRamp(common.BytesToAddress(offRampDestAddr), e.BlockChains.EVMChains()[lane.Dest.Selector].Client)
		require.NoError(t, err, "must instantiate offRamp")

		offRampSrcAddr, err := srcAdapter.GetOffRampAddress(e.DataStore, lane.Source.Selector)
		require.NoError(t, err, "must get offRamp from srcAdapter")

		feeQuoterOnSrcAddr, err := srcAdapter.GetFQAddress(e.DataStore, lane.Source.Selector)
		require.NoError(t, err, "must get feeQuoter from srcAdapter")
		feeQuoterOnSrc, err := fee_quoter.NewFeeQuoter(common.BytesToAddress(feeQuoterOnSrcAddr), e.BlockChains.EVMChains()[lane.Source.Selector].Client)
		require.NoError(t, err, "must instantiate feeQuoter")

		routerOnSrcAddr, err := srcAdapter.GetRouterAddress(e.DataStore, lane.Source.Selector)
		require.NoError(t, err, "must get router from srcAdapter")
		routerOnSrc, err := router.NewRouter(common.BytesToAddress(routerOnSrcAddr), e.BlockChains.EVMChains()[lane.Source.Selector].Client)
		require.NoError(t, err, "must instantiate router")

		routerOnDestAddr, err := destAdapter.GetRouterAddress(e.DataStore, lane.Dest.Selector)
		require.NoError(t, err, "must get router from destAdapter")
		routerOnDest, err := router.NewRouter(common.BytesToAddress(routerOnDestAddr), e.BlockChains.EVMChains()[lane.Dest.Selector].Client)
		require.NoError(t, err, "must instantiate router")

		destChainConfig, err := onRampSrc.GetDestChainConfig(nil, lane.Dest.Selector)
		require.NoError(t, err, "must get dest chain config from onRamp")
		routerAddr := routerOnSrc.Address().Hex()
		if disable {
			routerAddr = common.HexToAddress("0x0").Hex()
		}
		require.Equal(t, routerAddr, destChainConfig.Router.Hex(), "router must equal expected")
		require.Equal(t, lane.Dest.AllowListEnabled, destChainConfig.AllowlistEnabled, "allowListEnabled must equal expected")

		srcChainConfig, err := offRampDest.GetSourceChainConfig(nil, lane.Source.Selector)
		require.NoError(t, err, "must get src chain config from offRamp")
		require.Equal(t, !disable, srcChainConfig.IsEnabled, "isEnabled must be expected")
		require.Equal(t, !lane.Source.RMNVerificationEnabled, srcChainConfig.IsRMNVerificationDisabled, "rmnVerificationDisabled must equal expected")
		require.Equal(t, common.LeftPadBytes(onRampSrcAddr, 32), srcChainConfig.OnRamp, "remote onRamp must be set on offRamp")
		require.Equal(t, routerOnDest.Address().Hex(), srcChainConfig.Router.Hex(), "router must equal expected")

		isOffRamp, err := routerOnSrc.IsOffRamp(nil, lane.Dest.Selector, common.Address(offRampSrcAddr))
		require.NoError(t, err, "must check if router has offRamp")
		require.Equal(t, !disable, isOffRamp, "isOffRamp result must equal expected")
		onRampOnRouter, err := routerOnSrc.GetOnRamp(nil, lane.Dest.Selector)
		require.NoError(t, err, "must get onRamp from router")
		if disable {
			onRampSrcAddr = common.HexToAddress("0x0").Bytes()
		}
		require.Equal(t, onRampSrcAddr, onRampOnRouter.Bytes(), "onRamp must equal expected")

		isOffRamp, err = routerOnDest.IsOffRamp(nil, lane.Source.Selector, common.Address(offRampDestAddr))
		require.NoError(t, err, "must check if router has offRamp")
		require.Equal(t, !disable, isOffRamp, "isOffRamp result must equal expected")
		onRampOnRouter, err = routerOnDest.GetOnRamp(nil, lane.Source.Selector)
		require.NoError(t, err, "must get onRamp from router")
		if disable {
			onRampDestAddr = common.HexToAddress("0x0").Bytes()
		}
		require.Equal(t, onRampDestAddr, onRampOnRouter.Bytes(), "onRamp must equal expected")

		feeQuoterDestConfig, err := feeQuoterOnSrc.GetDestChainConfig(nil, lane.Dest.Selector)
		require.NoError(t, err, "must get dest chain config from feeQuoter")
		require.Equal(t, sequences.TranslateFQ(lane.Dest.FeeQuoterDestChainConfig), feeQuoterDestConfig, "feeQuoter dest chain config must equal expected")

		price, err := feeQuoterOnSrc.GetDestinationChainGasPrice(nil, lane.Dest.Selector)
		require.NoError(t, err, "must get price from feeQuoter")
		require.Equal(t, lane.Dest.GasPrice, price.Value, "price must equal expected")
	}
}

func TestConnectChains_EVM2EVM_NoMCMS(t *testing.T) {
	t.Parallel()
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.POLYGON_MAINNET.Selector,
	}
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	mcmsRegistry := cs_core.GetRegistry()
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	for _, chainSel := range chains {
		out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
			MCMS: mcms.Input{},
			Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
				chainSel: {
					Version: version,
					// FEE QUOTER CONFIG
					MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
					TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
					LinkPremiumMultiplier:        9e17, // 0.9 ETH
					NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
					// OFFRAMP CONFIG
					PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
					GasForCallExactCheck:                    uint16(5000),
				},
			},
		})
		require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
		out.DataStore.Merge(e.DataStore)
		e.DataStore = out.DataStore.Seal()
	}
	evmEncoded, err := hex.DecodeString(cciputils.EVMFamilySelector)
	require.NoError(t, err, "Failed to decode EVM family selector")
	chain1 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e17),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, evmEncoded),
	}
	chain2 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.POLYGON_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, evmEncoded),
	}

	_, err = lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{
				Version: version,
				ChainA:  chain1,
				ChainB:  chain2,
			},
		},
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")
	laneRegistry := lanesapi.GetLaneAdapterRegistry()
	src, dest := chains[0], chains[1]
	srcFamily, err := chain_selectors.GetSelectorFamily(src)
	require.NoError(t, err, "must get selector family for src")
	srcAdapter, exists := laneRegistry.GetLaneAdapter(srcFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for src chain family")
	destFamily, err := chain_selectors.GetSelectorFamily(dest)
	require.NoError(t, err, "must get selector family for dest")
	destAdapter, exists := laneRegistry.GetLaneAdapter(destFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for dest chain family")
	checkBidirectionalLaneConnectivity(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)
}
