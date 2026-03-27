package deployment

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	evmsequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	evmfq "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_3/fee_quoter"
	solanasequences "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/v0_1_1/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	fdeployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	evmfqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
)

// convertOpsConfigToGobinding converts operations type to gobinding type for test assertions.
// This helper keeps gobinding imports in tests, not production code.
func convertOpsConfigToGobinding(cfg evmfqops.DestChainConfig) evmfq.FeeQuoterDestChainConfig {
	return evmfq.FeeQuoterDestChainConfig{
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

func getFQOverrides() lanesapi.FeeQuoterDestChainConfigOverride {
	override := lanesapi.FeeQuoterDestChainConfigOverride(func(c *lanesapi.FeeQuoterDestChainConfig) {
		c.MaxDataBytes = 60_000
		if c.V1Params != nil {
			c.V1Params.EnforceOutOfOrder = true
		}
	})
	return override
}

func checkBidirectionalLaneConnectivity(
	t *testing.T,
	e *fdeployment.Environment,
	solanaChain lanesapi.ChainDefinition,
	evmChain lanesapi.ChainDefinition,
	solanaAdapter lanesapi.LaneAdapter,
	evmAdapter lanesapi.LaneAdapter,
	testRouter bool,
	disable bool,
) {
	// Solana Validation
	var offRampSourceChain ccip_offramp.SourceChain
	var destChainStateAccount ccip_router.DestChain
	var destChainFqAccount fee_quoter.DestChain
	var offRampEvmSourceChainPDA solana.PublicKey
	var evmDestChainStatePDA solana.PublicKey
	var fqEvmDestChainPDA solana.PublicKey
	feeQuoterOnSrcAddr, err := solanaAdapter.GetFQAddress(e.DataStore, solanaChain.Selector)
	require.NoError(t, err, "must get feeQuoter from srcAdapter")
	routerOnSrcAddr, err := solanaAdapter.GetRouterAddress(e.DataStore, solanaChain.Selector)
	require.NoError(t, err, "must get router from srcAdapter")
	offRampOnSrcAddr, err := solanaAdapter.GetOffRampAddress(e.DataStore, solanaChain.Selector)
	require.NoError(t, err, "must get offRamp from srcAdapter")

	// EVM Validation
	feeQuoterOnDestAddr, err := evmAdapter.GetFQAddress(e.DataStore, evmChain.Selector)
	require.NoError(t, err, "must get feeQuoter from evmAdapter")
	feeQuoterOnDest, err := evmfq.NewFeeQuoter(common.BytesToAddress(feeQuoterOnDestAddr), e.BlockChains.EVMChains()[evmChain.Selector].Client)
	require.NoError(t, err, "must instantiate feeQuoter")

	onRampDestAddr, err := evmAdapter.GetOnRampAddress(e.DataStore, evmChain.Selector)
	require.NoError(t, err, "must get onRamp from destAdapter")
	onRampDest, err := onramp.NewOnRamp(common.BytesToAddress(onRampDestAddr), e.BlockChains.EVMChains()[evmChain.Selector].Client)
	require.NoError(t, err, "must instantiate onRamp")

	offRampDestAddr, err := evmAdapter.GetOffRampAddress(e.DataStore, evmChain.Selector)
	require.NoError(t, err, "must get offRamp from destAdapter")
	offRampDest, err := offramp.NewOffRamp(common.BytesToAddress(offRampDestAddr), e.BlockChains.EVMChains()[evmChain.Selector].Client)
	require.NoError(t, err, "must instantiate offRamp")

	routerOnDestAddr, err := evmAdapter.GetRouterAddress(e.DataStore, evmChain.Selector)
	require.NoError(t, err, "must get router from destAdapter")
	routerOnDest, err := router.NewRouter(common.BytesToAddress(routerOnDestAddr), e.BlockChains.EVMChains()[evmChain.Selector].Client)
	require.NoError(t, err, "must instantiate router")

	// Validate EVM PDAs are set
	offRampEvmSourceChainPDA, _, _ = state.FindOfframpSourceChainPDA(evmChain.Selector, solana.PublicKeyFromBytes(offRampOnSrcAddr))
	err = e.BlockChains.SolanaChains()[solanaChain.Selector].GetAccountDataBorshInto(e.GetContext(), offRampEvmSourceChainPDA, &offRampSourceChain)
	require.NoError(t, err)
	require.Equal(t, !disable, offRampSourceChain.Config.IsEnabled)
	require.Equal(t, common.BytesToAddress(onRampDestAddr).Hex(), common.BytesToAddress(offRampSourceChain.Config.OnRamp.Bytes[:offRampSourceChain.Config.OnRamp.Len]).Hex())

	evmDestChainStatePDA, _ = state.FindDestChainStatePDA(evmChain.Selector, solana.PublicKeyFromBytes(routerOnSrcAddr))
	err = e.BlockChains.SolanaChains()[solanaChain.Selector].GetAccountDataBorshInto(e.GetContext(), evmDestChainStatePDA, &destChainStateAccount)
	require.NoError(t, err)
	require.Equal(t, solanaChain.AllowListEnabled, destChainStateAccount.Config.AllowListEnabled)

	fqEvmDestChainPDA, _, _ = state.FindFqDestChainPDA(evmChain.Selector, solana.PublicKeyFromBytes(feeQuoterOnSrcAddr))
	err = e.BlockChains.SolanaChains()[solanaChain.Selector].GetAccountDataBorshInto(e.GetContext(), fqEvmDestChainPDA, &destChainFqAccount)
	require.NoError(t, err, "failed to get account info")
	require.Equal(t, !disable, destChainFqAccount.Config.IsEnabled)
	ea := evmsequences.EVMAdapter{}
	require.Equal(t, solanasequences.TranslateFQ(ea.GetFeeQuoterDestChainConfig()), destChainFqAccount.Config)

	// Validate EVM onRamp/OffRamp/Router/FeeQuoter state
	destChainConfig, err := onRampDest.GetDestChainConfig(nil, solanaChain.Selector)
	require.NoError(t, err, "must get dest chain config from onRamp")
	require.Equal(t, routerOnDestAddr, destChainConfig.Router.Bytes(), "onRamp dest config router must always be the real router")
	require.Equal(t, evmChain.AllowListEnabled, destChainConfig.AllowlistEnabled, "allowListEnabled must equal expected")

	srcChainConfig, err := offRampDest.GetSourceChainConfig(nil, solanaChain.Selector)
	require.NoError(t, err, "must get src chain config from offRamp")
	require.Equal(t, !disable, srcChainConfig.IsEnabled, "isEnabled must be expected")
	require.Equal(t, !solanaChain.RMNVerificationEnabled, srcChainConfig.IsRMNVerificationDisabled, "rmnVerificationDisabled must equal expected")
	require.Equal(t, routerOnSrcAddr, srcChainConfig.OnRamp, "remote onRamp must be set on offRamp")
	// require.Equal(t, routerOnSrcAddr, srcChainConfig.Router.Bytes(), "router must equal expected")

	isOffRamp, err := routerOnDest.IsOffRamp(nil, solanaChain.Selector, common.Address(offRampDestAddr))
	require.NoError(t, err, "must check if router has offRamp")
	require.Equal(t, !disable, isOffRamp, "isOffRamp result must equal expected")
	// onRampOnRouter, err := routerOnDest.GetOnRamp(nil, solanaChain.Selector)
	// require.NoError(t, err, "must get onRamp from router")
	// onRampAddr := routerOnSrcAddr
	// if disable {
	// 	onRampAddr = common.HexToAddress("0x0").Bytes()
	// }
	// require.Equal(t, onRampAddr, onRampOnRouter.Bytes(), "onRamp must equal expected")

	feeQuoterDestConfig, err := feeQuoterOnDest.GetDestChainConfig(nil, solanaChain.Selector)
	require.NoError(t, err, "must get dest chain config from feeQuoter")
	sa := solanasequences.SolanaAdapter{}
	safq := sa.GetFeeQuoterDestChainConfig()
	override := getFQOverrides()
	override(&safq)
	expectedConfig := convertOpsConfigToGobinding(evmsequences.TranslateFQ(safq))
	require.Equal(t, expectedConfig, feeQuoterDestConfig, "feeQuoter dest chain config must equal expected")

	price, err := feeQuoterOnDest.GetDestinationChainGasPrice(nil, solanaChain.Selector)
	require.NoError(t, err, "must get price from feeQuoter")
	require.Equal(t, solanaChain.GasPrice, price.Value, "price must equal expected")
}

func checkLaneDisabled(
	t *testing.T,
	e *fdeployment.Environment,
	solanaChain lanesapi.ChainDefinition,
	evmChain lanesapi.ChainDefinition,
	solanaAdapter lanesapi.LaneAdapter,
	evmAdapter lanesapi.LaneAdapter,
) {
	t.Helper()

	// EVM Router: OnRamp must be zeroed for the remote chain
	routerOnDestAddr, err := evmAdapter.GetRouterAddress(e.DataStore, evmChain.Selector)
	require.NoError(t, err, "must get router from evmAdapter")
	routerOnDest, err := router.NewRouter(common.BytesToAddress(routerOnDestAddr), e.BlockChains.EVMChains()[evmChain.Selector].Client)
	require.NoError(t, err, "must instantiate router")

	onRampOnRouter, err := routerOnDest.GetOnRamp(nil, solanaChain.Selector)
	require.NoError(t, err, "must get onRamp from router")
	require.Equal(t, common.Address{}, onRampOnRouter, "onRamp must be zeroed after disable")

	// EVM Router: OffRamp must be removed for the remote chain
	offRampDestAddr, err := evmAdapter.GetOffRampAddress(e.DataStore, evmChain.Selector)
	require.NoError(t, err, "must get offRamp from evmAdapter")
	isOffRamp, err := routerOnDest.IsOffRamp(nil, solanaChain.Selector, common.Address(offRampDestAddr))
	require.NoError(t, err, "must check if router has offRamp")
	require.False(t, isOffRamp, "offRamp must be removed from router after disable")

	// Solana FeeQuoter: DestChain must be disabled for the remote chain
	feeQuoterOnSrcAddr, err := solanaAdapter.GetFQAddress(e.DataStore, solanaChain.Selector)
	require.NoError(t, err, "must get feeQuoter from solanaAdapter")
	var destChainFqAccount fee_quoter.DestChain
	fqEvmDestChainPDA, _, _ := state.FindFqDestChainPDA(evmChain.Selector, solana.PublicKeyFromBytes(feeQuoterOnSrcAddr))
	err = e.BlockChains.SolanaChains()[solanaChain.Selector].GetAccountDataBorshInto(e.GetContext(), fqEvmDestChainPDA, &destChainFqAccount)
	require.NoError(t, err, "must get FeeQuoter dest chain account")
	require.False(t, destChainFqAccount.Config.IsEnabled, "Solana FeeQuoter dest chain must be disabled")

	// Solana OffRamp: SourceChain must be disabled for the remote chain
	offRampOnSrcAddr, err := solanaAdapter.GetOffRampAddress(e.DataStore, solanaChain.Selector)
	require.NoError(t, err, "must get offRamp from solanaAdapter")
	var offRampSourceChain ccip_offramp.SourceChain
	offRampEvmSourceChainPDA, _, _ := state.FindOfframpSourceChainPDA(evmChain.Selector, solana.PublicKeyFromBytes(offRampOnSrcAddr))
	err = e.BlockChains.SolanaChains()[solanaChain.Selector].GetAccountDataBorshInto(e.GetContext(), offRampEvmSourceChainPDA, &offRampSourceChain)
	require.NoError(t, err, "must get OffRamp source chain account")
	require.False(t, offRampSourceChain.Config.IsEnabled, "Solana OffRamp source chain must be disabled")
}

// setupEVM2SVMForConnectChains deploys contracts on one EVM and one Solana chain and returns env, chain defs, and adapters.
func setupEVM2SVMForConnectChains(t *testing.T) (
	e *fdeployment.Environment,
	chain1, chain2 lanesapi.ChainDefinition,
	srcAdapter, destAdapter lanesapi.LaneAdapter,
	version *semver.Version,
) {
	t.Helper()
	programsPath, ds, err := PreloadSolanaEnvironment(t, chain_selectors.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, ds, "Datastore should be created")

	evmChains := []uint64{chain_selectors.ETHEREUM_MAINNET.Selector}
	solanaChains := []uint64{chain_selectors.SOLANA_MAINNET.Selector}
	allChains := append(evmChains, solanaChains...)
	e, err = environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
		environment.WithSolanaContainer(t, solanaChains, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	e.DataStore = ds.Seal()

	dReg := deployops.GetRegistry()
	version = semver.MustParse("1.6.0")
	for _, chainSel := range allChains {
		mint, _ := solana.NewRandomPrivateKey()
		out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
			MCMS: mcms.Input{},
			Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
				chainSel: {
					Version:                                 version,
					TokenPrivKey:                            mint.String(),
					TokenDecimals:                           9,
					MaxFeeJuelsPerMsg:                       big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
					TokenPriceStalenessThreshold:            uint32(24 * 60 * 60),
					LinkPremiumMultiplier:                   9e17,
					NativeTokenPremiumMultiplier:            1e18,
					PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
					GasForCallExactCheck:                    uint16(5000),
				},
			},
		})
		require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
		out.DataStore.Merge(e.DataStore)
		e.DataStore = out.DataStore.Seal()
	}
	DeployMCMS(t, e, chain_selectors.SOLANA_MAINNET.Selector, []string{cciputils.CLLQualifier})
	SolanaTransferOwnership(t, e, chain_selectors.SOLANA_MAINNET.Selector)

	override := getFQOverrides()
	chain1 = lanesapi.ChainDefinition{
		Selector:                          chain_selectors.SOLANA_MAINNET.Selector,
		GasPrice:                          big.NewInt(1e17),
		FeeQuoterDestChainConfigOverrides: &override,
	}
	chain2 = lanesapi.ChainDefinition{
		Selector: chain_selectors.ETHEREUM_MAINNET.Selector,
	}

	laneRegistry := lanesapi.GetLaneAdapterRegistry()
	srcFamily, err := chain_selectors.GetSelectorFamily(chain1.Selector)
	require.NoError(t, err, "must get selector family for src")
	srcAdapter, exists := laneRegistry.GetLaneAdapter(srcFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for src chain family")
	destFamily, err := chain_selectors.GetSelectorFamily(chain2.Selector)
	require.NoError(t, err, "must get selector family for dest")
	destAdapter, exists = laneRegistry.GetLaneAdapter(destFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for dest chain family")

	return e, chain1, chain2, srcAdapter, destAdapter, version
}

// TestConnectChains_EVM2SVM_Lifecycle exercises the full lifecycle of a bidirectional
// EVM↔Solana lane: connect → verify connectivity → disable → verify disabled state.
// A single Solana container setup covers all code paths efficiently.
func TestConnectChains_EVM2SVM_Lifecycle(t *testing.T) {
	t.Parallel()
	e, chain1, chain2, srcAdapter, destAdapter, version := setupEVM2SVMForConnectChains(t)
	mcmsRegistry := cs_core.GetRegistry()

	// ── Phase 1: Connect ─────────────────────────────────────────────────
	connectOut, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{
				Version: version,
				ChainA:  chain1,
				ChainB:  chain2,
			},
		},
		MCMS: NewDefaultInputForMCMS("Connect Chains"),
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")
	testhelpers.ProcessTimelockProposals(t, *e, connectOut.MCMSTimelockProposals, false)
	checkBidirectionalLaneConnectivity(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)
	require.Equal(t, 2, len(connectOut.DataStore.Addresses().Filter()))

	// ── Phase 2: Disable ─────────────────────────────────────────────────
	disableOut, err := lanesapi.DisableLane(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.DisableLaneConfig{
		Lanes: []lanesapi.DisableLanePair{
			{
				ChainA:  chain_selectors.SOLANA_MAINNET.Selector,
				ChainB:  chain_selectors.ETHEREUM_MAINNET.Selector,
				Version: version,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: true,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            cciputils.CLLQualifier,
			Description:          "Disable Lane",
		},
	})
	require.NoError(t, err, "Failed to apply DisableLane changeset")
	testhelpers.ProcessTimelockProposals(t, *e, disableOut.MCMSTimelockProposals, false)
	checkLaneDisabled(t, e, chain1, chain2, srcAdapter, destAdapter)
}

// setupEVM2EVMForConnectChains deploys 1.6 contracts on two EVM chains and returns env, chain defs, and adapters.
// Used by both TestConnectChains_EVM2EVM_NoMCMS and TestConnectChains_EVM2EVM_UpgradeFeeQuoter_ThenLaneExpansion.
func setupEVM2EVMForConnectChains(t *testing.T, chains []uint64) (
	e *fdeployment.Environment,
	chain1, chain2 lanesapi.ChainDefinition,
	srcAdapter, destAdapter lanesapi.LaneAdapter,
	version *semver.Version,
) {
	t.Helper()
	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, chains))
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")

	dReg := deployops.GetRegistry()
	version = semver.MustParse("1.6.0")
	deployCfg := deployops.ContractDeploymentConfigPerChain{
		Version:                                 version,
		MaxFeeJuelsPerMsg:                       big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
		TokenPriceStalenessThreshold:            uint32(24 * 60 * 60),
		LinkPremiumMultiplier:                   9e17,
		NativeTokenPremiumMultiplier:            1e18,
		PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
		GasForCallExactCheck:                    uint16(5000),
	}
	for _, chainSel := range chains {
		out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
			MCMS:   mcms.Input{},
			Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{chainSel: deployCfg},
		})
		require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
		out.DataStore.Merge(e.DataStore)
		e.DataStore = out.DataStore.Seal()
	}

	chain1 = lanesapi.ChainDefinition{
		Selector: chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice: big.NewInt(1e17),
	}
	chain2 = lanesapi.ChainDefinition{
		Selector: chain_selectors.POLYGON_MAINNET.Selector,
		GasPrice: big.NewInt(1e9),
	}

	laneRegistry := lanesapi.GetLaneAdapterRegistry()
	srcFamily, err := chain_selectors.GetSelectorFamily(chains[0])
	require.NoError(t, err)
	destFamily, err := chain_selectors.GetSelectorFamily(chains[1])
	require.NoError(t, err)
	srcAdapter, ok := laneRegistry.GetLaneAdapter(srcFamily, version)
	require.True(t, ok, "must have ChainAdapter for src chain family")
	destAdapter, ok = laneRegistry.GetLaneAdapter(destFamily, version)
	require.True(t, ok, "must have ChainAdapter for dest chain family")
	return e, chain1, chain2, srcAdapter, destAdapter, version
}

// TestConnectChains_EVM2EVM_NoMCMS is the EVM↔EVM connect-chains e2e (moved from
// chains/evm/deployment/v1_6_0/changesets/connect_chains_test.go). It deploys 1.6 contracts
// on two EVM chains, runs ConnectChains without deploying MCMS, and verifies bidirectional lane connectivity.
func TestConnectChains_EVM2EVM_NoMCMS(t *testing.T) {
	t.Parallel()
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.POLYGON_MAINNET.Selector,
	}
	e, chain1, chain2, srcAdapter, destAdapter, version := setupEVM2EVMForConnectChains(t, chains)
	mcmsRegistry := cs_core.GetRegistry()

	_, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{Version: version, ChainA: chain1, ChainB: chain2},
		},
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")

	checkBidirectionalLaneConnectivityEVM2EVM(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)
}

// checkBidirectionalLaneConnectivityEVM2EVM verifies lane state for EVM↔EVM. When useFeeQuoterV2
// is true, the FeeQuoter on the source chain is expected to be v2.0 (e.g. after upgrade); otherwise 1.6.x.
func checkBidirectionalLaneConnectivityEVM2EVM(
	t *testing.T,
	e *fdeployment.Environment,
	chainOne lanesapi.ChainDefinition,
	chainTwo lanesapi.ChainDefinition,
	srcAdapter lanesapi.LaneAdapter,
	destAdapter lanesapi.LaneAdapter,
	useFeeQuoterV2 bool,
	disable bool,
) {
	type laneDef struct {
		Source lanesapi.ChainDefinition
		Dest   lanesapi.ChainDefinition
	}
	lanes := []laneDef{
		{Source: chainOne, Dest: chainTwo},
		{Source: chainTwo, Dest: chainOne},
	}
	for _, lane := range lanes {
		// The adapter for the destination chain provides the expected FQ config.
		// In this EVM↔EVM test both adapters are identical, but we pick the right one.
		destAdapterForLane := destAdapter
		if lane.Dest.Selector == chainOne.Selector {
			destAdapterForLane = srcAdapter
		}
		expectedFQCfg := destAdapterForLane.GetFeeQuoterDestChainConfig()

		chain := e.BlockChains.EVMChains()[lane.Source.Selector]
		onRampSrcAddr, err := srcAdapter.GetOnRampAddress(e.DataStore, lane.Source.Selector)
		require.NoError(t, err, "must get onRamp from srcAdapter")
		onRampSrc, err := onramp.NewOnRamp(common.BytesToAddress(onRampSrcAddr), chain.Client)
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

		routerOnSrcAddr, err := srcAdapter.GetRouterAddress(e.DataStore, lane.Source.Selector)
		require.NoError(t, err, "must get router from srcAdapter")
		routerOnSrc, err := router.NewRouter(common.BytesToAddress(routerOnSrcAddr), chain.Client)
		require.NoError(t, err, "must instantiate router")

		routerOnDestAddr, err := destAdapter.GetRouterAddress(e.DataStore, lane.Dest.Selector)
		require.NoError(t, err, "must get router from destAdapter")
		routerOnDest, err := router.NewRouter(common.BytesToAddress(routerOnDestAddr), e.BlockChains.EVMChains()[lane.Dest.Selector].Client)
		require.NoError(t, err, "must instantiate router")

		destChainConfig, err := onRampSrc.GetDestChainConfig(nil, lane.Dest.Selector)
		require.NoError(t, err, "must get dest chain config from onRamp")
		require.Equal(t, routerOnSrc.Address().Hex(), destChainConfig.Router.Hex(), "onRamp dest config router must always be the real router")
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
		expOnRampSrc := onRampSrcAddr
		if disable {
			expOnRampSrc = common.HexToAddress("0x0").Bytes()
		}
		require.Equal(t, expOnRampSrc, onRampOnRouter.Bytes(), "onRamp must equal expected")

		isOffRamp, err = routerOnDest.IsOffRamp(nil, lane.Source.Selector, common.Address(offRampDestAddr))
		require.NoError(t, err, "must check if router has offRamp")
		require.Equal(t, !disable, isOffRamp, "isOffRamp result must equal expected")
		onRampOnRouter, err = routerOnDest.GetOnRamp(nil, lane.Source.Selector)
		require.NoError(t, err, "must get onRamp from router")
		expOnRampDest := onRampDestAddr
		if disable {
			expOnRampDest = common.HexToAddress("0x0").Bytes()
		}
		require.Equal(t, expOnRampDest, onRampOnRouter.Bytes(), "onRamp must equal expected")

		if useFeeQuoterV2 {
			fqContract, err := fqops.NewFeeQuoterContract(common.BytesToAddress(feeQuoterOnSrcAddr), chain.Client)
			require.NoError(t, err, "must instantiate FeeQuoter v2")
			destCfg, err := fqContract.GetDestChainConfig(nil, lane.Dest.Selector)
			require.NoError(t, err, "must get dest chain config from FeeQuoter v2")
			require.Equal(t, expectedFQCfg.IsEnabled, destCfg.IsEnabled, "feeQuoter v2 dest chain config IsEnabled must equal expected")
			require.Equal(t, expectedFQCfg.DefaultTxGasLimit, destCfg.DefaultTxGasLimit, "feeQuoter v2 dest chain config DefaultTxGasLimit must equal expected")
			price, err := fqContract.GetDestinationChainGasPrice(nil, lane.Dest.Selector)
			require.NoError(t, err, "must get gas price from FeeQuoter v2")
			require.Equal(t, lane.Dest.GasPrice, price.Value, "feeQuoter v2 gas price must equal expected")
		} else {
			feeQuoterOnSrc, err := evmfq.NewFeeQuoter(common.BytesToAddress(feeQuoterOnSrcAddr), chain.Client)
			require.NoError(t, err, "must instantiate feeQuoter 1.6")
			feeQuoterDestConfig, err := feeQuoterOnSrc.GetDestChainConfig(nil, lane.Dest.Selector)
			require.NoError(t, err, "must get dest chain config from feeQuoter")
			expectedConfig := convertOpsConfigToGobinding(evmsequences.TranslateFQ(expectedFQCfg))
			require.Equal(t, expectedConfig, feeQuoterDestConfig, "feeQuoter dest chain config must equal expected")
			price, err := feeQuoterOnSrc.GetDestinationChainGasPrice(nil, lane.Dest.Selector)
			require.NoError(t, err, "must get price from feeQuoter")
			require.Equal(t, lane.Dest.GasPrice, price.Value, "price must equal expected")
		}
	}
}

// TestConnectChains_EVM2EVM_Lifecycle exercises the full lifecycle of a bidirectional
// EVM lane: connect with allowlists → idempotent reconnect (prices not overwritten) →
// disconnect → idempotent disconnect (no revert). A single environment setup covers
// all code paths efficiently.
func TestConnectChains_EVM2EVM_Lifecycle(t *testing.T) {
	t.Parallel()
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.POLYGON_MAINNET.Selector,
	}
	e, chain1, chain2, srcAdapter, destAdapter, version := setupEVM2EVMForConnectChains(t, chains)
	mcmsRegistry := cs_core.GetRegistry()

	allowedSender1 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	allowedSender2 := common.HexToAddress("0x2222222222222222222222222222222222222222")
	chain1.AllowListEnabled = true
	chain1.AllowList = []string{allowedSender1.Hex(), allowedSender2.Hex()}
	chain2.AllowListEnabled = true
	chain2.AllowList = []string{allowedSender1.Hex()}
	initialGasPrice1 := new(big.Int).Set(chain1.GasPrice)
	initialGasPrice2 := new(big.Int).Set(chain2.GasPrice)

	connect := func(isDisabled bool) {
		t.Helper()
		_, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
			Lanes: []lanesapi.LaneConfig{
				{Version: version, ChainA: chain1, ChainB: chain2, IsDisabled: isDisabled},
			},
		})
		require.NoError(t, err)
	}
	freshBundle := func() {
		t.Helper()
		e.OperationsBundle = operations.NewBundle(t.Context, e.OperationsBundle.Logger, operations.NewMemoryReporter())
	}

	// ── Phase 1: Connect ─────────────────────────────────────────────────
	connect(false)
	checkBidirectionalLaneConnectivityEVM2EVM(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)

	// Verify allowlists: each chain's OnRamp reflects its own AllowList.
	onRampAddr1, err := srcAdapter.GetOnRampAddress(e.DataStore, chain1.Selector)
	require.NoError(t, err)
	onRampOn1, err := onramp.NewOnRamp(common.BytesToAddress(onRampAddr1), e.BlockChains.EVMChains()[chain1.Selector].Client)
	require.NoError(t, err)
	al1, err := onRampOn1.GetAllowedSendersList(nil, chain2.Selector)
	require.NoError(t, err)
	require.True(t, al1.IsEnabled)
	require.ElementsMatch(t, []common.Address{allowedSender1, allowedSender2}, al1.ConfiguredAddresses)

	onRampAddr2, err := srcAdapter.GetOnRampAddress(e.DataStore, chain2.Selector)
	require.NoError(t, err)
	onRampOn2, err := onramp.NewOnRamp(common.BytesToAddress(onRampAddr2), e.BlockChains.EVMChains()[chain2.Selector].Client)
	require.NoError(t, err)
	al2, err := onRampOn2.GetAllowedSendersList(nil, chain1.Selector)
	require.NoError(t, err)
	require.True(t, al2.IsEnabled)
	require.ElementsMatch(t, []common.Address{allowedSender1}, al2.ConfiguredAddresses)

	// ── Phase 2: Reconnect with different prices (must NOT overwrite) ────
	freshBundle()
	chain1.GasPrice = new(big.Int).Mul(big.NewInt(999), big.NewInt(1e17))
	chain2.GasPrice = big.NewInt(999_000_000_000)
	connect(false)

	fqAddr1, err := srcAdapter.GetFQAddress(e.DataStore, chain1.Selector)
	require.NoError(t, err)
	fqOn1, err := evmfq.NewFeeQuoter(common.BytesToAddress(fqAddr1), e.BlockChains.EVMChains()[chain1.Selector].Client)
	require.NoError(t, err)
	price1, err := fqOn1.GetDestinationChainGasPrice(nil, chain2.Selector)
	require.NoError(t, err)
	require.Equal(t, initialGasPrice2, price1.Value, "gas price must not be overwritten")

	fqAddr2, err := srcAdapter.GetFQAddress(e.DataStore, chain2.Selector)
	require.NoError(t, err)
	fqOn2, err := evmfq.NewFeeQuoter(common.BytesToAddress(fqAddr2), e.BlockChains.EVMChains()[chain2.Selector].Client)
	require.NoError(t, err)
	price2, err := fqOn2.GetDestinationChainGasPrice(nil, chain1.Selector)
	require.NoError(t, err)
	require.Equal(t, initialGasPrice1, price2.Value, "gas price must not be overwritten")

	// ── Phase 3: Disconnect ──────────────────────────────────────────────
	freshBundle()
	chain1.GasPrice = initialGasPrice1
	chain2.GasPrice = initialGasPrice2
	connect(true)
	checkBidirectionalLaneConnectivityEVM2EVM(t, e, chain1, chain2, srcAdapter, destAdapter, false, true)

	// ── Phase 4: Disconnect again (must not revert) ──────────────────────
	freshBundle()
	connect(true)
	checkBidirectionalLaneConnectivityEVM2EVM(t, e, chain1, chain2, srcAdapter, destAdapter, false, true)
}

// TestConnectChains_EVM2EVM_UpgradeFeeQuoter_ThenLaneExpansion runs an e2e with 1.6 FeeQuoter, upgrades to 2.0,
// then runs ConnectChains again (lane expansion / re-apply) and verifies connectivity using the 2.0 FeeQuoter path.
func TestConnectChains_EVM2EVM_UpgradeFeeQuoter_ThenLaneExpansion(t *testing.T) {
	t.Parallel()
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.POLYGON_MAINNET.Selector,
	}
	e, chain1, chain2, srcAdapter, destAdapter, version := setupEVM2EVMForConnectChains(t, chains)
	mcmsRegistry := cs_core.GetRegistry()

	_, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{Version: version, ChainA: chain1, ChainB: chain2},
		},
		MCMS: NewDefaultInputForMCMS("Connect Chains"),
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")

	checkBidirectionalLaneConnectivityEVM2EVM(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)

	// Deploy MCMS so we can run the FeeQuoter 2.0 upgrade (timelock).
	DeployMCMS(t, e, chain_selectors.ETHEREUM_MAINNET.Selector, []string{cciputils.CLLQualifier})
	DeployMCMS(t, e, chain_selectors.POLYGON_MAINNET.Selector, []string{cciputils.CLLQualifier})
	// Reset bundle so second ConnectChains runs without cached executions.
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle
	fqInput := map[uint64]deployops.UpdateFeeQuoterInputPerChain{
		chain_selectors.ETHEREUM_MAINNET.Selector: {
			FeeQuoterVersion: semver.MustParse("2.0.0"),
			RampsVersion:     version,
		},
		chain_selectors.POLYGON_MAINNET.Selector: {
			FeeQuoterVersion: semver.MustParse("2.0.0"),
			RampsVersion:     version,
		},
	}
	fqUpdateChangeset := deployops.UpdateFeeQuoterChangeset()
	out, err := fqUpdateChangeset.Apply(*e, deployops.UpdateFeeQuoterInput{
		Chains: fqInput,
		MCMS:   NewDefaultInputForMCMS("Transfer ownership FQ2"),
	})
	require.NoError(t, err, "Failed to apply UpdateFeeQuoterChangeset")
	require.Greater(t, len(out.Reports), 0)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
	require.NoError(t, out.DataStore.Merge(e.DataStore), "Failed to merge changeset output datastore")
	e.DataStore = out.DataStore.Seal()

	// Reset bundle so second ConnectChains runs without cached executions.
	bundle = operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle

	// Run ConnectChains again (lane expansion / configure-as-source with 2.0 FeeQuoter).
	connectOut2, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{Version: version, ChainA: chain1, ChainB: chain2},
		},
		MCMS: NewDefaultInputForMCMS("Connect Chains after FQ upgrade"),
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset after FQ upgrade")
	testhelpers.ProcessTimelockProposals(t, *e, connectOut2.MCMSTimelockProposals, false)

	// Verify connectivity using 2.0 FeeQuoter (adapter now returns 2.0 address from datastore).
	checkBidirectionalLaneConnectivityEVM2EVM(t, e, chain1, chain2, srcAdapter, destAdapter, true, false)
}

func TestDowngradeLane_ConnectChains_EVM2EVM(t *testing.T) {
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.AVALANCHE_MAINNET.Selector,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	mcmsRegistry := cs_core.GetRegistry()
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	chainInput := make(map[uint64]deployops.ContractDeploymentConfigPerChain)
	fqInput := make(map[uint64]deployops.UpdateFeeQuoterInputPerChain)

	for _, chainSel := range chains {
		chainInput[chainSel] = deployops.ContractDeploymentConfigPerChain{
			Version: version,
			// FEE QUOTER CONFIG
			MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
			TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
			LinkPremiumMultiplier:        9e17, // 0.9 ETH
			NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
			// OFFRAMP CONFIG
			PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			GasForCallExactCheck:                    uint16(5000),
		}
	}

	fqInput[chain_selectors.ETHEREUM_MAINNET.Selector] = deployops.UpdateFeeQuoterInputPerChain{
		FeeQuoterVersion:     semver.MustParse("2.0.0"),
		RampsVersion:         semver.MustParse("1.6.0"),
		RemoteChainSelectors: []uint64{chain_selectors.AVALANCHE_MAINNET.Selector},
	}
	fqInput[chain_selectors.AVALANCHE_MAINNET.Selector] = deployops.UpdateFeeQuoterInputPerChain{
		FeeQuoterVersion:     semver.MustParse("2.0.0"),
		RampsVersion:         semver.MustParse("1.6.0"),
		RemoteChainSelectors: []uint64{chain_selectors.ETHEREUM_MAINNET.Selector},
	}

	out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
	require.NoError(t, out.DataStore.Merge(e.DataStore))
	e.DataStore = out.DataStore.Seal()
	chain1 := lanesapi.ChainDefinition{
		Selector: chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice: big.NewInt(1e9),
	}
	chain2 := lanesapi.ChainDefinition{
		Selector: chain_selectors.AVALANCHE_MAINNET.Selector,
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
	// Deploy MCMS
	DeployMCMS(t, e, chain_selectors.ETHEREUM_MAINNET.Selector, []string{common_utils.CLLQualifier})
	DeployMCMS(t, e, chain_selectors.AVALANCHE_MAINNET.Selector, []string{common_utils.CLLQualifier})
	// do this to reset cached executions
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle
	// now update to FeeQuoter 2.0.0
	t.Log(out.DataStore.Addresses())
	fqUpdateChangeset := deployops.UpdateFeeQuoterChangeset()
	out, err = fqUpdateChangeset.Apply(*e, deployops.UpdateFeeQuoterInput{
		Chains: fqInput,
		MCMS:   NewDefaultInputForMCMS("Transfer ownership FQ2"),
	})
	require.NoError(t, err, "Failed to apply UpdateFeeQuoterChangeset changeset")
	require.Greater(t, len(out.Reports), 0)
	require.Equal(t, 1, len(out.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
	// update datastore with changeset output
	require.NoError(t, out.DataStore.Merge(e.DataStore), "Failed to merge changeset output datastore")
	e.DataStore = out.DataStore.Seal()
	t.Log(out.DataStore.Addresses())
	for _, chainSel := range chains {
		fqUpgradeValidation(t, e, chainSel, chains, true, true)
	}
	// do this to reset cached executions
	bundle = operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle
	// downgrade back to 1.6.3 to make sure the changeset is reversible
	for _, chainSel := range chains {
		fqInput[chainSel] = deployops.UpdateFeeQuoterInputPerChain{
			FeeQuoterVersion: semver.MustParse("1.6.3"),
			RampsVersion:     semver.MustParse("1.6.0"),
		}
	}
	// now downgrade back to FeeQuoter 1.6.3
	fqUpdateChangeset = deployops.UpdateFeeQuoterChangeset()
	out, err = fqUpdateChangeset.Apply(*e, deployops.UpdateFeeQuoterInput{
		Chains: fqInput,
	})
	require.NoError(t, err, "Failed to apply UpdateFeeQuoterChangeset changeset")
	require.Len(t, out.DataStore.Addresses().Filter(), 0, "new addresses found on downgrade")
	require.NoError(t, out.DataStore.Merge(e.DataStore), "Failed to merge changeset output datastore")

	for _, chainSel := range chains {
		fqUpgradeValidation(t, e, chainSel, chains, false, true)
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

	// Verify the version in Datastore is still 1.6.0 after downgrade
	chain1Family, err := chain_selectors.GetSelectorFamily(chain1.Selector)

	chain1Adapter, exists := lanesapi.GetLaneAdapterRegistry().GetLaneAdapter(chain1Family, version)
	require.Equal(t, exists, true, "adapter should exist for chain1")

	chain2Family, err := chain_selectors.GetSelectorFamily(chain2.Selector)

	chain2Adapter, exists := lanesapi.GetLaneAdapterRegistry().GetLaneAdapter(chain2Family, version)
	require.Equal(t, exists, true, "adapter should exist for chain2")

	// Verify the version of the FQ version has been downgraded
	fqAddressChain1 := common.Address{}
	if dfq, ok := chain1Adapter.(lanesapi.DynamicFeeQuoter); ok {
		fqAddressChain1Bytes, err := dfq.GetFQAddressDynamic(e.DataStore, chain1.Selector, e.BlockChains)
		require.NoError(t, err, "error fetching fee quoter address %d", chain1.Selector)
		fqAddressChain1 = common.BytesToAddress(fqAddressChain1Bytes)
	}
	chain1Ref := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chain1.Selector),
		datastore.AddressRefByType(datastore.ContractType(cciputils.FeeQuoter)),
		datastore.AddressRefByAddress(fqAddressChain1.Hex()),
	)
	require.Len(t, chain1Ref, 1, "expected exactly 1 fee quoter address for chain1 after downgrade")

	fqAddressChain2 := common.Address{}
	if dfq, ok := chain2Adapter.(lanesapi.DynamicFeeQuoter); ok {
		fqAddressChain2Bytes, err := dfq.GetFQAddressDynamic(e.DataStore, chain2.Selector, e.BlockChains)
		require.NoError(t, err, "error fetching fee quoter address %d", chain2.Selector)
		fqAddressChain2 = common.BytesToAddress(fqAddressChain2Bytes)
	}
	chain2Ref := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chain2.Selector),
		datastore.AddressRefByType(datastore.ContractType(cciputils.FeeQuoter)),
		datastore.AddressRefByAddress(fqAddressChain2.Hex()),
	)
	require.Len(t, chain2Ref, 1, "expected exactly 1 fee quoter address for chain2 after downgrade")

	require.Equal(t, chain1Ref[0].Version, semver.MustParse("1.6.3"))
	require.Equal(t, chain2Ref[0].Version, semver.MustParse("1.6.3"))

	require.NoError(t, err, "Failed to apply ConnectChains changeset")
}
