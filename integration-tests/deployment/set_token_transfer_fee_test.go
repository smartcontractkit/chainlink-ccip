package deployment

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	evmadaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evmadaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	onrampV1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	evmadaptersV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	fqV2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_0_0/adapters"

	soladaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestSetTokenTransferFeeV1_6_0(t *testing.T) {
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

	// Ensure chains exist
	_, ok := env.BlockChains.SolanaChains()[src]
	require.True(t, ok, "Source chain not found in environment")
	_, ok = env.BlockChains.EVMChains()[dst]
	require.True(t, ok, "Destination chain not found in environment")

	// Initialize v1.6.0 adapters
	solAdapter := solseqV1_6_0.SolanaAdapter{}
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	// Configure deployment registry
	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, &evmadaptersV1_0_0.EVMDeployer{})
	deployRegistry.RegisterDeployer(chainsel.FamilySolana, deploy.MCMSVersion, &solAdapter)

	// Adapters registered via init() in adapter packages
	evmFeesAdapter := evmadaptersV1_6_0.NewFeesAdapter(&evmAdapter)
	solFeesAdapter := soladaptersV1_6_0.NewFeesAdapter(&solAdapter)

	// Configure MCMS registry
	mcmsRegistry := changesets.GetRegistry()
	mcmsAdapter := &evmadaptersV1_0_0.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, mcmsAdapter)

	// Deploy FeeQuoter + other contracts
	SeedUltraFastCurseMCMS(t, env)
	output, err := deploy.DeployContracts(deployRegistry).Apply(*env, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			src: NewDefaultDeploymentConfigForSolana(utils.Version_1_6_0),
			dst: NewDefaultDeploymentConfigForEVM(utils.Version_1_6_0),
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	// Connect the chains so that srcRouter.getOnRamp(dst) works
	output, err = lanes.ConnectChains(lanes.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*env, lanes.ConnectChainsConfig{
		Lanes: []lanes.LaneConfig{
			{
				Version: utils.Version_1_6_0,
				ChainA:  lanes.ChainDefinition{Selector: src},
				ChainB:  lanes.ChainDefinition{Selector: dst},
			},
		},
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	// Any valid address can be used when setting token transfer fees. The
	// contracts do *not* validate that the addresses are actually tokens.
	srcTokenAddress := "GcqdKBdgcJNdBeC1TnZvJTaWuRXXg8WotC5qw1BNBSEp"
	dstTokenAddress := "0x2222222222222222222222222222222222222222"

	// Get the FQ on the source
	srcOnRampRef := datastore.AddressRef{ChainSelector: src, Type: datastore.ContractType(router.ContractType), Version: utils.Version_1_6_0}
	srcOnRampRef, err = datastore_utils.FindAndFormatRef(env.DataStore, srcOnRampRef, src, datastore_utils.FullRef)
	require.NoError(t, err)
	srcFQ, err := solFeesAdapter.GetFeeContractRef(env.OperationsBundle, env.BlockChains, env.DataStore, srcOnRampRef, src, dst)
	require.NoError(t, err)

	// Get the FQ on the destination
	dstOnRampRef := datastore.AddressRef{ChainSelector: dst, Type: datastore.ContractType(onrampV1_6.ContractType), Version: utils.Version_1_6_0}
	dstOnRampRef, err = datastore_utils.FindAndFormatRef(env.DataStore, dstOnRampRef, dst, datastore_utils.FullRef)
	require.NoError(t, err)
	dstFQ, err := evmFeesAdapter.GetFeeContractRef(env.OperationsBundle, env.BlockChains, env.DataStore, dstOnRampRef, dst, src)
	require.NoError(t, err)

	// Set the token transfer fee config for LINK
	_, err = fees.
		SetTokenTransferFee().
		Apply(*env, fees.SetTokenTransferFeeInput{
			Version: utils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcTokenAddress,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: utils.NewOptional(uint32(120_000)),
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
									Address: dstTokenAddress,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: utils.NewOptional(uint32(150_000)),
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
	srcCfg, err := solFeesAdapter.GetOnchainTokenTransferFeeConfig(env.OperationsBundle, env.BlockChains, srcFQ, src, dst, srcTokenAddress)
	require.NoError(t, err)
	srcSensibleDefaults := solFeesAdapter.GetDefaultTokenTransferFeeConfig(src, dst)
	require.Equal(t, srcCfg.DestBytesOverhead, srcSensibleDefaults.DestBytesOverhead)
	require.Equal(t, srcCfg.DestGasOverhead, uint32(120_000))
	require.Equal(t, srcCfg.MinFeeUSDCents, srcSensibleDefaults.MinFeeUSDCents)
	require.Equal(t, srcCfg.MaxFeeUSDCents, srcSensibleDefaults.MaxFeeUSDCents)
	require.Equal(t, srcCfg.DeciBps, srcSensibleDefaults.DeciBps)
	require.True(t, srcCfg.IsEnabled)

	// Confirm that the config was correctly set on the destination
	dstCfg, err := evmFeesAdapter.GetOnchainTokenTransferFeeConfig(env.OperationsBundle, env.BlockChains, dstFQ, dst, src, dstTokenAddress)
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
		SetTokenTransferFee().
		Apply(*env, fees.SetTokenTransferFeeInput{
			Version: utils.Version_1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcTokenAddress,
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
									Address: dstTokenAddress,
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
	srcCfg, err = solFeesAdapter.GetOnchainTokenTransferFeeConfig(env.OperationsBundle, env.BlockChains, srcFQ, src, dst, srcTokenAddress)
	require.NoError(t, err)
	require.False(t, srcCfg.IsEnabled)

	// Confirm that the config was disabled on the destination
	dstCfg, err = evmFeesAdapter.GetOnchainTokenTransferFeeConfig(env.OperationsBundle, env.BlockChains, dstFQ, dst, src, dstTokenAddress)
	require.NoError(t, err)
	require.False(t, dstCfg.IsEnabled)
}

func TestSetTokenTransferFeeV2_0_0(t *testing.T) {
	src := chainsel.TEST_90000002.Selector
	dst := chainsel.TEST_90000001.Selector

	chains := []uint64{
		src,
		dst,
	}

	e, err := environment.New(t.Context(), environment.WithEVMSimulated(t, chains))
	require.NoError(t, err)

	mcmsRegistry := changesets.GetRegistry()
	dplyRegistry := deploy.GetRegistry()

	chainInput := make(map[uint64]deploy.ContractDeploymentConfigPerChain)
	fqInput := make(map[uint64]deploy.UpdateFeeQuoterInputPerChain)
	for _, chainSel := range chains {
		chainInput[chainSel] = NewDefaultDeploymentConfigForEVM(utils.Version_1_6_0)
		fqInput[chainSel] = deploy.UpdateFeeQuoterInputPerChain{
			FeeQuoterVersion: utils.Version_2_0_0,
			RampsVersion:     utils.Version_1_6_0,
		}
	}

	// Deploy FeeQuoter + other contracts
	SeedUltraFastCurseMCMS(t, e)
	out, err := deploy.DeployContracts(dplyRegistry).Apply(*e, deploy.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
	MergeAddresses(t, e, out.DataStore)

	// Connect the chains so that srcRouter.getOnRamp(dst) works
	connectOut, err := lanes.ConnectChains(lanes.GetLaneAdapterRegistry(), mcmsRegistry).Apply(*e, lanes.ConnectChainsConfig{
		Lanes: []lanes.LaneConfig{
			{
				Version: utils.Version_1_6_0,
				ChainA:  lanes.ChainDefinition{Selector: src},
				ChainB:  lanes.ChainDefinition{Selector: dst},
			},
		},
	})
	require.NoError(t, err, "Failed to apply ConnectChains changeset")
	MergeAddresses(t, e, connectOut.DataStore)

	// Deploy MCMS
	DeployMCMS(t, e, src, []string{utils.CLLQualifier})
	DeployMCMS(t, e, dst, []string{utils.CLLQualifier})

	// Reset bundle so second ConnectChains runs without cached executions.
	e.OperationsBundle = operations.NewBundle(e.GetContext, e.Logger, operations.NewMemoryReporter())

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
	MergeAddresses(t, e, out.DataStore)
	for _, chainSel := range chains {
		fqUpgradeValidation(t, e, chainSel, chains, true, true)
	}

	// Ensure new FQ v2.0 contracts exist in datastore
	refs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByType(datastore.ContractType(fqV2_0.ContractType)),
		datastore.AddressRefByVersion(utils.Version_2_0_0),
	)
	require.Len(t, refs, len(chains))

	// Any valid address can be used when setting token transfer fees. The
	// contracts do *not* validate that the addresses are actually tokens.
	evmFeesAdapterV2_0_0 := evmadaptersV2_0_0.NewFeesAdapter(&evmseqV1_6_0.EVMAdapter{})
	evmFeesAdapterV1_6_0 := evmadaptersV1_6_0.NewFeesAdapter(&evmseqV1_6_0.EVMAdapter{})
	srcTokenAddress := "0x1111111111111111111111111111111111111111"
	dstTokenAddress := "0x2222222222222222222222222222222222222222"

	// Reset bundle so second ConnectChains runs without cached executions.
	e.OperationsBundle = operations.NewBundle(e.GetContext, e.Logger, operations.NewMemoryReporter())

	// Get the FQ on the source
	srcOnRampRef := datastore.AddressRef{ChainSelector: src, Type: datastore.ContractType(onrampV1_6.ContractType), Version: utils.Version_1_6_0}
	srcOnRampRef, err = datastore_utils.FindAndFormatRef(e.DataStore, srcOnRampRef, src, datastore_utils.FullRef)
	require.NoError(t, err)
	srcFQ, err := evmFeesAdapterV1_6_0.GetFeeContractRef(e.OperationsBundle, e.BlockChains, e.DataStore, srcOnRampRef, src, dst)
	require.NoError(t, err)
	require.True(t, srcFQ.Version.Equal(utils.Version_2_0_0), "Expected v1.6 OnRamp to be connected to v2.0 FeeQuoter after upgrade, but got version %s", srcFQ.Version.String())

	// Get the FQ on the destination
	dstOnRampRef := datastore.AddressRef{ChainSelector: dst, Type: datastore.ContractType(onrampV1_6.ContractType), Version: utils.Version_1_6_0}
	dstOnRampRef, err = datastore_utils.FindAndFormatRef(e.DataStore, dstOnRampRef, dst, datastore_utils.FullRef)
	require.NoError(t, err)
	dstFQ, err := evmFeesAdapterV1_6_0.GetFeeContractRef(e.OperationsBundle, e.BlockChains, e.DataStore, dstOnRampRef, dst, src)
	require.NoError(t, err)
	require.True(t, dstFQ.Version.Equal(utils.Version_2_0_0), "Expected v1.6 OnRamp to be connected to v2.0 FeeQuoter after upgrade, but got version %s", dstFQ.Version.String())

	// Set the token transfer fee config for LINK
	out, err = fees.
		SetTokenTransferFee().
		Apply(*e, fees.SetTokenTransferFeeInput{
			Version: utils.Version_2_0_0,
			MCMS:    NewDefaultInputForMCMS("Set token transfer fee"),
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcTokenAddress,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: utils.NewOptional(uint32(150_000)),
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
									Address: dstTokenAddress,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{
										DestGasOverhead: utils.NewOptional(uint32(150_000)),
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

	// Confirm that the config was correctly set on the source
	srcCfg, err := evmFeesAdapterV2_0_0.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, srcFQ, src, dst, srcTokenAddress)
	require.NoError(t, err)
	srcSensibleDefaults := evmFeesAdapterV2_0_0.GetDefaultTokenTransferFeeConfig(src, dst)
	require.Equal(t, srcCfg.DestBytesOverhead, srcSensibleDefaults.DestBytesOverhead)
	require.Equal(t, srcCfg.DestGasOverhead, uint32(150_000))
	require.Equal(t, srcCfg.MinFeeUSDCents, srcSensibleDefaults.MinFeeUSDCents)
	require.True(t, srcCfg.IsEnabled)

	// Confirm that the config was correctly set on the destination
	dstCfg, err := evmFeesAdapterV2_0_0.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, dstFQ, dst, src, dstTokenAddress)
	require.NoError(t, err)
	dstSensibleDefaults := evmFeesAdapterV2_0_0.GetDefaultTokenTransferFeeConfig(dst, src)
	require.Equal(t, dstCfg.DestBytesOverhead, dstSensibleDefaults.DestBytesOverhead)
	require.Equal(t, dstCfg.DestGasOverhead, uint32(150_000))
	require.Equal(t, dstCfg.MinFeeUSDCents, dstSensibleDefaults.MinFeeUSDCents)
	require.True(t, dstCfg.IsEnabled)

	out, err = fees.
		SetTokenTransferFee().
		Apply(*e, fees.SetTokenTransferFeeInput{
			Version: utils.Version_2_0_0,
			MCMS:    NewDefaultInputForMCMS("Set token transfer fee"),
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcTokenAddress,
									IsReset: false,
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
									Address: dstTokenAddress,
									IsReset: false,
									FeeArgs: fees.UnresolvedTokenTransferFeeArgs{},
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	require.Empty(t, out.MCMSTimelockProposals)

	// Now reset the configs
	out, err = fees.
		SetTokenTransferFee().
		Apply(*e, fees.SetTokenTransferFeeInput{
			Version: utils.Version_2_0_0,
			MCMS:    NewDefaultInputForMCMS("Set token transfer fee"),
			Args: []fees.TokenTransferFeeForSrc{
				{
					Selector: src,
					Settings: []fees.TokenTransferFeeForDst{
						{
							Selector: dst,
							Settings: []fees.TokenTransferFee{
								{
									Address: srcTokenAddress,
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
									Address: dstTokenAddress,
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
	srcCfg, err = evmFeesAdapterV2_0_0.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, srcFQ, src, dst, srcTokenAddress)
	require.NoError(t, err)
	require.False(t, srcCfg.IsEnabled)

	// Confirm that the config was disabled on the destination
	dstCfg, err = evmFeesAdapterV2_0_0.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, dstFQ, dst, src, dstTokenAddress)
	require.NoError(t, err)
	require.False(t, dstCfg.IsEnabled)
}
