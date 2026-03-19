package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	evmadaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	evmadaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	evmadaptersV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	soladaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	tokensops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/tokens"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestSetTokenTransferFeeV1_6_0(t *testing.T) {
	// Define alias for v1.6.0
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)

	// Define source and destination chain selectors
	src := chainsel.SOLANA_DEVNET.Selector
	dst := chainsel.TEST_90000002.Selector

	// Preload Solana programs
	programsPath, ds, err := PreloadSolanaEnvironment(t, src)
	require.NoError(t, err)

	// Setup test environment
	env, err := environment.New(t.Context(),
		environment.WithSolanaContainer(t, []uint64{src}, programsPath, solanaProgramIDs),
		environment.WithEVMSimulated(t, []uint64{dst}),
	)
	require.NoError(t, err)
	env.DataStore = ds.Seal()

	// Initialize v1.6.0 adapters
	solAdapter := solseqV1_6_0.SolanaAdapter{}
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	// Configure deployment registry
	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, &evmadaptersV1_0_0.EVMDeployer{})
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deploy.MCMSVersion, &solAdapter)

	// Configure fees registry
	feesRegistry := fees.GetRegistry()
	evmFeesAdapter := evmadaptersV1_6_0.NewFeesAdapter(&evmAdapter)
	solFeesAdapter := soladaptersV1_6_0.NewFeesAdapter(&solAdapter)
	feesRegistry.RegisterFeeAdapter(chainsel.FamilySolana, v1_6_0, solFeesAdapter)
	feesRegistry.RegisterFeeAdapter(chainsel.FamilyEVM, v1_6_0, evmFeesAdapter)

	// Configure MCMS registry
	mcmsRegistry := changesets.GetRegistry()
	mcmsAdapter := &evmadaptersV1_0_0.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, mcmsAdapter)

	// Deploy FeeQuoter + other contracts
	output, err := deploy.DeployContracts(deployRegistry).Apply(*env, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			src: NewDefaultDeploymentConfigForSolana(v1_6_0),
			dst: NewDefaultDeploymentConfigForEVM(v1_6_0),
		},
	})
	require.NoError(t, err)

	// Get the address of the LINK token on the source chain
	srcLinkRef, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(src,
			datastore.ContractType(tokensops.LinkContractType),
			semver.MustParse("1.6.0"), // version 1.6.0 for Solana LINK token
			"",                        // no qualifier is needed for Solana LINK token
		),
	)
	require.NoError(t, err)

	// Get the address of the LINK token on the destination chain
	dstLinkRef, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(dst,
			datastore.ContractType(link.ContractType),
			link.Version,
			"", // no qualifier is needed for EVM LINK token
		),
	)
	require.NoError(t, err)

	// Ensure environment has the updated datastore
	env.DataStore = output.DataStore.Seal()

	// Set the token transfer fee config for LINK
	_, err = fees.
		SetTokenTransferFee(feesRegistry, mcmsRegistry).
		Apply(*env, fees.SetTokenTransferFeeInput{
			Version: v1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcLinkRef.Address,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: fees.TokenTransferFeeValue[uint32]{
											Value: 120_000,
											Valid: true,
										},
									},
								},
							},
						},
					},
				},
				{
					Selector: dst,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: src,
							Settings: []fees.TokenTransferFee{
								{
									Address: dstLinkRef.Address,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: fees.TokenTransferFeeValue[uint32]{
											Value: 150_000,
											Valid: true,
										},
									},
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)

	// Confirm that the config was correctly set on the source
	srcCfg, err := solFeesAdapter.GetOnchainTokenTransferFeeConfig(*env, src, dst, srcLinkRef.Address)
	require.NoError(t, err)
	srcSensibleDefaults := solFeesAdapter.GetDefaultTokenTransferFeeConfig(src, dst)
	require.Equal(t, srcCfg.DestBytesOverhead, srcSensibleDefaults.DestBytesOverhead)
	require.Equal(t, srcCfg.DestGasOverhead, uint32(120_000))
	require.Equal(t, srcCfg.MinFeeUSDCents, srcSensibleDefaults.MinFeeUSDCents)
	require.Equal(t, srcCfg.MaxFeeUSDCents, srcSensibleDefaults.MaxFeeUSDCents)
	require.Equal(t, srcCfg.DeciBps, srcSensibleDefaults.DeciBps)
	require.True(t, srcCfg.IsEnabled)

	// Confirm that the config was correctly set on the destination
	dstCfg, err := evmFeesAdapter.GetOnchainTokenTransferFeeConfig(*env, dst, src, dstLinkRef.Address)
	require.NoError(t, err)
	dstSensibleDefaults := evmFeesAdapter.GetDefaultTokenTransferFeeConfig(dst, src)
	require.Equal(t, dstCfg.DestBytesOverhead, dstSensibleDefaults.DestBytesOverhead)
	require.Equal(t, dstCfg.DestGasOverhead, uint32(150_000))
	require.Equal(t, dstCfg.MinFeeUSDCents, dstSensibleDefaults.MinFeeUSDCents)
	require.Equal(t, dstCfg.MaxFeeUSDCents, dstSensibleDefaults.MaxFeeUSDCents)
	require.Equal(t, dstCfg.DeciBps, dstSensibleDefaults.DeciBps)
	require.True(t, dstCfg.IsEnabled)

	// Now reset the configs
	_, err = fees.
		SetTokenTransferFee(feesRegistry, mcmsRegistry).
		Apply(*env, fees.SetTokenTransferFeeInput{
			Version: v1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcLinkRef.Address,
									IsReset: true,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{},
								},
							},
						},
					},
				},
				{
					Selector: dst,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: src,
							Settings: []fees.TokenTransferFee{
								{
									Address: dstLinkRef.Address,
									IsReset: true,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{},
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)

	// Confirm that the config was disabled on the source
	srcCfg, err = solFeesAdapter.GetOnchainTokenTransferFeeConfig(*env, src, dst, srcLinkRef.Address)
	require.NoError(t, err)
	require.False(t, srcCfg.IsEnabled)

	// Confirm that the config was disabled on the destination
	dstCfg, err = evmFeesAdapter.GetOnchainTokenTransferFeeConfig(*env, dst, src, dstLinkRef.Address)
	require.NoError(t, err)
	require.False(t, dstCfg.IsEnabled)
}

func TestSetTokenTransferFeeV2_0_0(t *testing.T) {
	// Define alias for v1.6.0 & v2.0.0
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)
	v2_0_0, err := semver.NewVersion("2.0.0")
	require.NoError(t, err)

	src := chainsel.TEST_90000002.Selector
	dst := chainsel.TEST_90000001.Selector

	chains := []uint64{
		src,
		dst,
	}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)

	require.NoError(t, err, "Failed to create test environment")
	require.NotNil(t, e, "Environment should be created")
	mcmsRegistry := changesets.GetRegistry()
	dReg := deploy.GetRegistry()
	chainInput := make(map[uint64]deploy.ContractDeploymentConfigPerChain)
	fqInput := make(map[uint64]deploy.UpdateFeeQuoterInputPerChain)

	for _, chainSel := range chains {
		chainInput[chainSel] = deploy.ContractDeploymentConfigPerChain{
			Version: v1_6_0,
			// FEE QUOTER CONFIG
			MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
			TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
			LinkPremiumMultiplier:        9e17, // 0.9 ETH
			NativeTokenPremiumMultiplier: 1e18, // 1.0 ETH
			// OFFRAMP CONFIG
			PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			GasForCallExactCheck:                    uint16(5000),
		}
		fqInput[chainSel] = deploy.UpdateFeeQuoterInputPerChain{
			FeeQuoterVersion: v2_0_0,
			RampsVersion:     v1_6_0,
		}
	}
	out, err := deploy.DeployContracts(dReg).Apply(*e, deploy.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
	MergeAddresses(t, e, out.DataStore)

	chain1 := lanes.ChainDefinition{
		Selector:                 src,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanes.DefaultFeeQuoterDestChainConfig(true, src),
	}
	chain2 := lanes.ChainDefinition{
		Selector:                 dst,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanes.DefaultFeeQuoterDestChainConfig(true, dst),
	}
	_, err = lanes.ConnectChains(lanes.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanes.ConnectChainsConfig{
		Lanes: []lanes.LaneConfig{
			{
				Version: v1_6_0,
				ChainA:  chain1,
				ChainB:  chain2,
			},
		},
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")
	// Deploy MCMS
	DeployMCMS(t, e, src, []string{common_utils.CLLQualifier})
	DeployMCMS(t, e, dst, []string{common_utils.CLLQualifier})

	// now update to FeeQuoter 2.0.0
	fqUpdateChangeset := deploy.UpdateFeeQuoterChangeset()
	out, err = fqUpdateChangeset.Apply(*e, deploy.UpdateFeeQuoterInput{
		Chains: fqInput,
		MCMS:   NewDefaultInputForMCMS("Transfer ownership FQ2"),
	})
	require.NoError(t, err, "Failed to apply UpdateFeeQuoterChangeset changeset")
	require.Greater(t, len(out.Reports), 0)
	require.Equal(t, 1, len(out.MCMSTimelockProposals))

	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
	// update datastore with changeset output
	MergeAddresses(t, e, out.DataStore)

	for _, chainSel := range chains {
		fqUpgradeValidation(t, e, chainSel, chains, true)
	}

	// Configure fees registry
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	feesRegistry := fees.GetRegistry()
	evmFeesAdapterV2_0 := evmadaptersV2_0_0.NewFeesAdapter(&evmAdapter)

	// Get the address of the LINK token on the source chain
	srcLinkRef, err := out.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(src,
			datastore.ContractType(link.ContractType),
			semver.MustParse("1.0.0"),
			"", // no qualifier is needed for EVM LINK token
		),
	)
	require.NoError(t, err)

	// Get the address of the LINK token on the destination chain
	dstLinkRef, err := out.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(dst,
			datastore.ContractType(link.ContractType),
			semver.MustParse("1.0.0"),
			"", // no qualifier is needed for EVM LINK token
		),
	)
	require.NoError(t, err)

	// Ensure environment has the updated datastore
	e.DataStore = out.DataStore.Seal()

	// Set the token transfer fee config for LINK
	out, err = fees.
		SetTokenTransferFee(feesRegistry, mcmsRegistry).
		Apply(*e, fees.SetTokenTransferFeeInput{
			Version: v2_0_0,
			MCMS:    NewDefaultInputForMCMS("Set token transfer fee"),
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcLinkRef.Address,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: fees.TokenTransferFeeValue[uint32]{
											Value: 150_000,
											Valid: true,
										},
									},
								},
							},
						},
					},
				},
				{
					Selector: dst,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: src,
							Settings: []fees.TokenTransferFee{
								{
									Address: dstLinkRef.Address,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: fees.TokenTransferFeeValue[uint32]{
											Value: 150_000,
											Valid: true,
										},
									},
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)
	// require.NoError(t, out.DataStore.Merge(e.DataStore), "Failed to merge changeset output datastore")

	// Confirm that the config was correctly set on the source
	srcCfg, err := evmFeesAdapterV2_0.GetOnchainTokenTransferFeeConfig(*e, src, dst, srcLinkRef.Address)
	require.NoError(t, err)
	srcSensibleDefaults := evmFeesAdapterV2_0.GetDefaultTokenTransferFeeConfig(src, dst)
	require.Equal(t, srcCfg.DestBytesOverhead, srcSensibleDefaults.DestBytesOverhead)
	require.Equal(t, srcCfg.DestGasOverhead, uint32(150_000))
	require.Equal(t, srcCfg.MinFeeUSDCents, srcSensibleDefaults.MinFeeUSDCents)
	require.True(t, srcCfg.IsEnabled)

	// Confirm that the config was correctly set on the destination
	dstCfg, err := evmFeesAdapterV2_0.GetOnchainTokenTransferFeeConfig(*e, dst, src, dstLinkRef.Address)
	require.NoError(t, err)
	dstSensibleDefaults := evmFeesAdapterV2_0.GetDefaultTokenTransferFeeConfig(dst, src)
	require.Equal(t, dstCfg.DestBytesOverhead, dstSensibleDefaults.DestBytesOverhead)
	require.Equal(t, dstCfg.DestGasOverhead, uint32(150_000))
	require.Equal(t, dstCfg.MinFeeUSDCents, dstSensibleDefaults.MinFeeUSDCents)
	require.True(t, dstCfg.IsEnabled)

	// Now reset the configs
	out, err = fees.
		SetTokenTransferFee(feesRegistry, mcmsRegistry).
		Apply(*e, fees.SetTokenTransferFeeInput{
			Version: v2_0_0,
			MCMS:    NewDefaultInputForMCMS("Set token transfer fee"),
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcLinkRef.Address,
									IsReset: true,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{},
								},
							},
						},
					},
				},
				{
					Selector: dst,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: src,
							Settings: []fees.TokenTransferFee{
								{
									Address: dstLinkRef.Address,
									IsReset: true,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{},
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err, "Failed to apply UpdateFeeQuoterChangeset changeset")
	require.Greater(t, len(out.Reports), 0)
	require.Equal(t, 1, len(out.MCMSTimelockProposals))
	testhelpers.ProcessTimelockProposals(t, *e, out.MCMSTimelockProposals, false)

	// Confirm that the config was disabled on the source
	srcCfg, err = evmFeesAdapterV2_0.GetOnchainTokenTransferFeeConfig(*e, src, dst, srcLinkRef.Address)
	require.NoError(t, err)
	require.False(t, srcCfg.IsEnabled)

	// Confirm that the config was disabled on the destination
	dstCfg, err = evmFeesAdapterV2_0.GetOnchainTokenTransferFeeConfig(*e, dst, src, dstLinkRef.Address)
	require.NoError(t, err)
	require.False(t, dstCfg.IsEnabled)
}
