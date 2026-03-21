package changesets_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	fdeployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"

	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
)

// convertOpsConfigToGobinding converts operations type to gobinding type for test assertions.
// This helper keeps gobinding imports in tests, not production code.
func convertOpsConfigToGobinding(cfg fqops.DestChainConfig) fee_quoter.FeeQuoterDestChainConfig {
	return fee_quoter.FeeQuoterDestChainConfig{
		IsEnabled:                         cfg.IsEnabled,
		MaxNumberOfTokensPerMsg:           cfg.MaxNumberOfTokensPerMsg,
		MaxDataBytes:                      cfg.MaxDataBytes,
		MaxPerMsgGasLimit:                 cfg.MaxPerMsgGasLimit,
		DestGasOverhead:                   cfg.DestGasOverhead,
		DestGasPerPayloadByteBase:         cfg.DestGasPerPayloadByteBase,
		DestGasPerPayloadByteHigh:         cfg.DestGasPerPayloadByteHigh,
		DestGasPerPayloadByteThreshold:    cfg.DestGasPerPayloadByteThreshold,
		DestDataAvailabilityOverheadGas:   cfg.DestDataAvailabilityOverheadGas,
		DestGasPerDataAvailabilityByte:    cfg.DestGasPerDataAvailabilityByte,
		DestDataAvailabilityMultiplierBps: cfg.DestDataAvailabilityMultiplierBps,
		ChainFamilySelector:               cfg.ChainFamilySelector,
		EnforceOutOfOrder:                 cfg.EnforceOutOfOrder,
		DefaultTokenFeeUSDCents:           cfg.DefaultTokenFeeUSDCents,
		DefaultTokenDestGasOverhead:       cfg.DefaultTokenDestGasOverhead,
		DefaultTxGasLimit:                 cfg.DefaultTxGasLimit,
		GasMultiplierWeiPerEth:            cfg.GasMultiplierWeiPerEth,
		GasPriceStalenessThreshold:        cfg.GasPriceStalenessThreshold,
		NetworkFeeUSDCents:                cfg.NetworkFeeUSDCents,
	}
}

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
		require.Equal(t, lane.Source.AllowListEnabled, destChainConfig.AllowlistEnabled, "allowListEnabled must equal expected")

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
		a := &sequences.EVMAdapter{}
		expectedConfig := convertOpsConfigToGobinding(sequences.TranslateFQ(a.GetFeeQuoterDestChainConfig()))
		require.Equal(t, expectedConfig, feeQuoterDestConfig, "feeQuoter dest chain config must equal expected")

		price, err := feeQuoterOnSrc.GetDestinationChainGasPrice(nil, lane.Dest.Selector)
		require.NoError(t, err, "must get price from feeQuoter")
		require.Equal(t, lane.Dest.GasPrice, price.Value, "price must equal expected")
	}
}

func TestConnectChains_AllowlistAndIdempotentPrices(t *testing.T) {
	t.Parallel()
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.POLYGON_MAINNET.Selector,
	}
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err)

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
					LinkPremiumMultiplier:        9e17,
					NativeTokenPremiumMultiplier: 1e18,
					// OFFRAMP CONFIG
					PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
					GasForCallExactCheck:                    uint16(5000),
				},
			},
		})
		require.NoError(t, err)
		out.DataStore.Merge(e.DataStore)
		e.DataStore = out.DataStore.Seal()
	}

	allowedSender1 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	allowedSender2 := common.HexToAddress("0x2222222222222222222222222222222222222222")

	initialGasPrice1 := big.NewInt(1e17)
	initialGasPrice2 := big.NewInt(1e9)

	chain1 := lanesapi.ChainDefinition{
		Selector:         chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice:         initialGasPrice1,
		AllowListEnabled: true,
		AllowList:        []string{allowedSender1.Hex(), allowedSender2.Hex()},
	}
	chain2 := lanesapi.ChainDefinition{
		Selector:         chain_selectors.POLYGON_MAINNET.Selector,
		GasPrice:         initialGasPrice2,
		AllowListEnabled: true,
		AllowList:        []string{allowedSender1.Hex()},
	}

	// --- First ConnectChains: seeds prices and applies allowlists ---
	_, err = lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{Version: version, ChainA: chain1, ChainB: chain2},
		},
	})
	require.NoError(t, err, "first ConnectChains should succeed")

	laneRegistry := lanesapi.GetLaneAdapterRegistry()
	srcAdapter, exists := laneRegistry.GetLaneAdapter(chain_selectors.FamilyEVM, version)
	require.True(t, exists)
	destAdapter, exists := laneRegistry.GetLaneAdapter(chain_selectors.FamilyEVM, version)
	require.True(t, exists)

	// Verify basic lane connectivity (including gas prices)
	checkBidirectionalLaneConnectivity(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)

	// The allowlist controls who on a given chain can send outbound messages.
	// Each chain's OnRamp reflects its own AllowList.
	// chain1's OnRamp (dest=chain2) → chain1's AllowList = [sender1, sender2]
	// chain2's OnRamp (dest=chain1) → chain2's AllowList = [sender1]

	onRampAddr1, err := srcAdapter.GetOnRampAddress(e.DataStore, chain1.Selector)
	require.NoError(t, err)
	onRampOn1, err := onramp.NewOnRamp(common.BytesToAddress(onRampAddr1), e.BlockChains.EVMChains()[chain1.Selector].Client)
	require.NoError(t, err)
	allowlistResult, err := onRampOn1.GetAllowedSendersList(nil, chain2.Selector)
	require.NoError(t, err)
	require.True(t, allowlistResult.IsEnabled, "allowlist must be enabled on chain1 OnRamp for dest=chain2")
	require.ElementsMatch(t, []common.Address{allowedSender1, allowedSender2}, allowlistResult.ConfiguredAddresses,
		"chain1 OnRamp for dest=chain2 must reflect chain1's AllowList")

	onRampAddr2, err := srcAdapter.GetOnRampAddress(e.DataStore, chain2.Selector)
	require.NoError(t, err)
	onRampOn2, err := onramp.NewOnRamp(common.BytesToAddress(onRampAddr2), e.BlockChains.EVMChains()[chain2.Selector].Client)
	require.NoError(t, err)
	allowlistResult, err = onRampOn2.GetAllowedSendersList(nil, chain1.Selector)
	require.NoError(t, err)
	require.True(t, allowlistResult.IsEnabled, "allowlist must be enabled on chain2 OnRamp for dest=chain1")
	require.ElementsMatch(t, []common.Address{allowedSender1}, allowlistResult.ConfiguredAddresses,
		"chain2 OnRamp for dest=chain1 must reflect chain2's AllowList")

	// Verify initial gas prices were seeded
	fqAddr1, err := srcAdapter.GetFQAddress(e.DataStore, chain1.Selector)
	require.NoError(t, err)
	fqOn1, err := fee_quoter.NewFeeQuoter(common.BytesToAddress(fqAddr1), e.BlockChains.EVMChains()[chain1.Selector].Client)
	require.NoError(t, err)
	price1, err := fqOn1.GetDestinationChainGasPrice(nil, chain2.Selector)
	require.NoError(t, err)
	require.Equal(t, initialGasPrice2, price1.Value, "chain1 FeeQuoter gas price for dest=chain2 must equal initial value")
	fqAddr2, err := srcAdapter.GetFQAddress(e.DataStore, chain2.Selector)
	require.NoError(t, err)
	fqOn2, err := fee_quoter.NewFeeQuoter(common.BytesToAddress(fqAddr2), e.BlockChains.EVMChains()[chain2.Selector].Client)
	require.NoError(t, err)
	price2, err := fqOn2.GetDestinationChainGasPrice(nil, chain1.Selector)
	require.NoError(t, err)
	require.Equal(t, initialGasPrice1, price2.Value, "chain2 FeeQuoter gas price for dest=chain1 must equal initial value")

	// --- Second ConnectChains with different gas prices: must NOT overwrite ---
	// Fresh operations bundle simulates a separate production invocation (no cached reads).
	e.OperationsBundle = operations.NewBundle(t.Context, e.OperationsBundle.Logger, operations.NewMemoryReporter())
	chain1.GasPrice = new(big.Int).Mul(big.NewInt(999), big.NewInt(1e17)) // changed
	chain2.GasPrice = big.NewInt(999_000_000_000)                        // changed
	_, err = lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{Version: version, ChainA: chain1, ChainB: chain2},
		},
	})
	require.NoError(t, err, "second ConnectChains should succeed")

	// Verify gas prices are still the ORIGINAL values (idempotency)
	price1After, err := fqOn1.GetDestinationChainGasPrice(nil, chain2.Selector)
	require.NoError(t, err)
	require.Equal(t, initialGasPrice2, price1After.Value,
		"chain1 FeeQuoter gas price must not be overwritten on second ConnectChains")

	price2After, err := fqOn2.GetDestinationChainGasPrice(nil, chain1.Selector)
	require.NoError(t, err)
	require.Equal(t, initialGasPrice1, price2After.Value,
		"chain2 FeeQuoter gas price must not be overwritten on second ConnectChains")
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
	chain1 := lanesapi.ChainDefinition{
		Selector: chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice: big.NewInt(1e17),
	}
	chain2 := lanesapi.ChainDefinition{
		Selector: chain_selectors.POLYGON_MAINNET.Selector,
		GasPrice: big.NewInt(1e9),
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
