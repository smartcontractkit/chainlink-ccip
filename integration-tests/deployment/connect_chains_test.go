package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	mcms_types "github.com/smartcontractkit/mcms/types"
	"github.com/stretchr/testify/require"

	evmsequences "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
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
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	fdeployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"

	evmfqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
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
	require.NoError(t, err, "must get feeQuoter from srcAdapter")
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
	require.Equal(t, solanasequences.TranslateFQ(evmChain.FeeQuoterDestChainConfig), destChainFqAccount.Config)

	// Validate EVM onRamp/OffRamp/Router/FeeQuoter state
	destChainConfig, err := onRampDest.GetDestChainConfig(nil, solanaChain.Selector)
	require.NoError(t, err, "must get dest chain config from onRamp")
	routerAddr := routerOnDestAddr
	if disable {
		routerAddr = common.HexToAddress("0x0").Bytes()
	}
	require.Equal(t, routerAddr, destChainConfig.Router.Bytes(), "router must equal expected")
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
	expectedConfig := convertOpsConfigToGobinding(evmsequences.TranslateFQ(solanaChain.FeeQuoterDestChainConfig))
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

func TestConnectChains_EVM2SVM_NoMCMS(t *testing.T) {
	t.Parallel()
	programsPath, ds, err := PreloadSolanaEnvironment(t, chain_selectors.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, ds, "Datastore should be created")

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}
	solanaChains := []uint64{
		chain_selectors.SOLANA_MAINNET.Selector,
	}
	allChains := append(evmChains, solanaChains...)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
		environment.WithSolanaContainer(t, solanaChains, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	e.DataStore = ds.Seal() // Add preloaded contracts to env datastore

	mcmsRegistry := cs_core.GetRegistry()
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	for _, chainSel := range allChains {
		mint, _ := solana.NewRandomPrivateKey()
		out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
			MCMS: mcms.Input{},
			Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
				chainSel: {
					Version: version,
					// LINK TOKEN CONFIG
					// token private key used to deploy the LINK token. Solana: base58 encoded private key
					TokenPrivKey: mint.String(),
					// token decimals used to deploy the LINK token
					TokenDecimals: 9,
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
	DeployMCMS(t, e, chain_selectors.SOLANA_MAINNET.Selector, []string{cciputils.CLLQualifier})
	SolanaTransferOwnership(t, e, chain_selectors.SOLANA_MAINNET.Selector)
	// TODO: EVM doesn't work with a non-zero timelock delay
	// DeployMCMS(t, e, chain_selectors.ETHEREUM_MAINNET.Selector)
	// EVMTransferOwnership(t, e, chain_selectors.ETHEREUM_MAINNET.Selector)
	chain1 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.SOLANA_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e17),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.SOLANA_MAINNET.Selector),
	}
	chain2 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.ETHEREUM_MAINNET.Selector),
	}

	connectOut, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{
				Version: version,
				ChainA:  chain1,
				ChainB:  chain2,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            cciputils.CLLQualifier,
			Description:          "Connect Chains",
		},
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")
	testhelpers.ProcessTimelockProposals(t, *e, connectOut.MCMSTimelockProposals, false)
	laneRegistry := lanesapi.GetLaneAdapterRegistry()
	srcFamily, err := chain_selectors.GetSelectorFamily(chain1.Selector)
	require.NoError(t, err, "must get selector family for src")
	srcAdapter, exists := laneRegistry.GetLaneAdapter(srcFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for src chain family")
	destFamily, err := chain_selectors.GetSelectorFamily(chain2.Selector)
	require.NoError(t, err, "must get selector family for dest")
	destAdapter, exists := laneRegistry.GetLaneAdapter(destFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for dest chain family")
	checkBidirectionalLaneConnectivity(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)
	// should add a new entry for remote source and remote dest in solana
	require.Equal(t, 2, len(connectOut.DataStore.Addresses().Filter()))
}

func TestDisableLane_EVM2SVM(t *testing.T) {
	t.Parallel()
	programsPath, ds, err := PreloadSolanaEnvironment(t, chain_selectors.SOLANA_MAINNET.Selector)
	require.NoError(t, err, "Failed to set up Solana environment")
	require.NotNil(t, ds, "Datastore should be created")

	evmChains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
	}
	solanaChains := []uint64{
		chain_selectors.SOLANA_MAINNET.Selector,
	}
	allChains := append(evmChains, solanaChains...)
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, evmChains),
		environment.WithSolanaContainer(t, solanaChains, programsPath, solanaProgramIDs),
	)
	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	e.DataStore = ds.Seal()

	mcmsRegistry := cs_core.GetRegistry()
	dReg := deployops.GetRegistry()
	version := semver.MustParse("1.6.0")
	for _, chainSel := range allChains {
		mint, _ := solana.NewRandomPrivateKey()
		out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
			MCMS: mcms.Input{},
			Chains: map[uint64]deployops.ContractDeploymentConfigPerChain{
				chainSel: {
					Version: version,
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

	chain1 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.SOLANA_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e17),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.SOLANA_MAINNET.Selector),
	}
	chain2 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.ETHEREUM_MAINNET.Selector),
	}

	connectOut, err := lanesapi.ConnectChains(lanesapi.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.ConnectChainsConfig{
		Lanes: []lanesapi.LaneConfig{
			{
				Version: version,
				ChainA:  chain1,
				ChainB:  chain2,
			},
		},
		MCMS: mcms.Input{
			OverridePreviousRoot: false,
			ValidUntil:           3759765795,
			TimelockDelay:        mcms_types.MustParseDuration("1s"),
			TimelockAction:       mcms_types.TimelockActionSchedule,
			Qualifier:            cciputils.CLLQualifier,
			Description:          "Connect Chains",
		},
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")
	testhelpers.ProcessTimelockProposals(t, *e, connectOut.MCMSTimelockProposals, false)

	laneRegistry := lanesapi.GetLaneAdapterRegistry()
	srcFamily, err := chain_selectors.GetSelectorFamily(chain1.Selector)
	require.NoError(t, err, "must get selector family for src")
	srcAdapter, exists := laneRegistry.GetLaneAdapter(srcFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for src chain family")
	destFamily, err := chain_selectors.GetSelectorFamily(chain2.Selector)
	require.NoError(t, err, "must get selector family for dest")
	destAdapter, exists := laneRegistry.GetLaneAdapter(destFamily, version)
	require.True(t, exists, "must have ChainAdapter registered for dest chain family")
	checkBidirectionalLaneConnectivity(t, e, chain1, chain2, srcAdapter, destAdapter, false, false)

	disableOut, err := lanesapi.DisableLane(lanesapi.GetDisableLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanesapi.DisableLaneConfig{
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
