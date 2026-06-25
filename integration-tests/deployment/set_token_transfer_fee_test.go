package deployment

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	datastore_utils_evm "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	evmadaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	bnmERC20ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	evmadaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	onrampV1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	evmadaptersV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	fqV2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	tpopsV2_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_0_0/adapters"
	soladaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/router"
	solseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
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

func TestSetTokenPoolTokenTransferFeeV2_0_0(t *testing.T) {
	// Define EVM chains
	evmChainSelA := chainsel.TEST_90000001.Selector
	evmChainSelB := chainsel.TEST_90000002.Selector
	evmChainSelC := chainsel.TEST_90000003.Selector

	// Setup test environment
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{evmChainSelA, evmChainSelB, evmChainSelC}))
	require.NoError(t, err)

	// Get chains
	evmChainA, ok := env.BlockChains.EVMChains()[evmChainSelA]
	require.True(t, ok)
	evmChainB, ok := env.BlockChains.EVMChains()[evmChainSelB]
	require.True(t, ok)
	evmChainC, ok := env.BlockChains.EVMChains()[evmChainSelC]
	require.True(t, ok)

	// Configure deployment registry
	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, &evmadaptersV1_0_0.EVMDeployer{})
	chainsToDeploy := map[uint64]deploy.ContractDeploymentConfigPerChain{
		evmChainSelA: NewDefaultDeploymentConfigForEVM(utils.Version_1_6_0),
		evmChainSelB: NewDefaultDeploymentConfigForEVM(utils.Version_1_6_0),
		evmChainSelC: NewDefaultDeploymentConfigForEVM(utils.Version_1_6_0),
	}

	// Deploy contracts
	output, err := deploy.DeployContracts(deployRegistry).Apply(*env, deploy.ContractDeploymentConfig{Chains: chainsToDeploy, MCMS: mcms.Input{}})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	// Deploy MCMS and transfer ownership
	DeployMCMS(t, env, evmChainSelA, []string{utils.CLLQualifier})
	DeployMCMS(t, env, evmChainSelB, []string{utils.CLLQualifier})
	DeployMCMS(t, env, evmChainSelC, []string{utils.CLLQualifier})
	EVMTransferOwnership(t, env, evmChainSelA)
	EVMTransferOwnership(t, env, evmChainSelB)
	EVMTransferOwnership(t, env, evmChainSelC)

	// Define a full pool mesh between all EVM chains
	disabledOutboundRL := tokens.RateLimiterConfigFloatInput{IsEnabled: false}
	tokenExpansionInput := map[uint64]tokens.TokenExpansionInputPerChain{
		evmChainSelA: {
			TokenPoolVersion: utils.Version_2_0_0,
			DeployTokenInput: &tokens.DeployTokenInput{
				Type:     bnmERC20ops.ContractType,
				Decimals: 18,
				Symbol:   "TEST_TOKEN_B",
				Name:     "Test Token B",
				PreMint:  nil, // no pre-mint
				Supply:   nil, // unlimited
			},
			DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
				PoolType: utils.BurnMintTokenPool.String(),
			},
			TokenTransferConfig: &tokens.TokenTransferConfig{
				RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					evmChainSelB: {OutboundRateLimiterConfig: &disabledOutboundRL},
					evmChainSelC: {OutboundRateLimiterConfig: &disabledOutboundRL},
				},
			},
		},
		evmChainSelB: {
			TokenPoolVersion: utils.Version_2_0_0,
			DeployTokenInput: &tokens.DeployTokenInput{
				Type:     bnmERC20ops.ContractType,
				Decimals: 18,
				Symbol:   "TEST_TOKEN_B",
				Name:     "Test Token B",
				PreMint:  nil, // no pre-mint
				Supply:   nil, // unlimited
			},
			DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
				PoolType: utils.BurnMintTokenPool.String(),
			},
			TokenTransferConfig: &tokens.TokenTransferConfig{
				RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					evmChainSelA: {OutboundRateLimiterConfig: &disabledOutboundRL},
					evmChainSelC: {OutboundRateLimiterConfig: &disabledOutboundRL},
				},
			},
		},
		evmChainSelC: {
			TokenPoolVersion: utils.Version_2_0_0,
			DeployTokenInput: &tokens.DeployTokenInput{
				Type:     bnmERC20ops.ContractType,
				Decimals: 18,
				Symbol:   "TEST_TOKEN_C",
				Name:     "Test Token C",
				PreMint:  nil, // no pre-mint
				Supply:   nil, // unlimited
			},
			DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
				PoolType: utils.BurnMintTokenPool.String(),
			},
			TokenTransferConfig: &tokens.TokenTransferConfig{
				RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
					evmChainSelA: {OutboundRateLimiterConfig: &disabledOutboundRL},
					evmChainSelB: {OutboundRateLimiterConfig: &disabledOutboundRL},
				},
			},
		},
	}

	// Deploy the pool mesh
	output, err = tokens.TokenExpansion().Apply(*env, tokens.TokenExpansionInput{
		TokenExpansionInputPerChain: tokenExpansionInput,
		ChainAdapterVersion:         utils.Version_2_0_0,
		MCMS:                        NewDefaultInputForMCMS("Token Expansion"),
	})
	require.NoError(t, err)
	MergeAddresses(t, env, output.DataStore)

	// Fetch the pool addresses from the datastore
	fltrA := datastore.AddressRef{Type: datastore.ContractType(tokenExpansionInput[evmChainSelA].DeployTokenPoolInput.PoolType), Version: tokenExpansionInput[evmChainSelA].DeployTokenPoolInput.TokenPoolVersion}
	poolA, err := datastore_utils.FindAndFormatRef(env.DataStore, fltrA, evmChainSelA, datastore_utils_evm.ToEVMAddress)
	require.NoError(t, err)
	fltrB := datastore.AddressRef{Type: datastore.ContractType(tokenExpansionInput[evmChainSelB].DeployTokenPoolInput.PoolType), Version: tokenExpansionInput[evmChainSelB].DeployTokenPoolInput.TokenPoolVersion}
	poolB, err := datastore_utils.FindAndFormatRef(env.DataStore, fltrB, evmChainSelB, datastore_utils_evm.ToEVMAddress)
	require.NoError(t, err)
	fltrC := datastore.AddressRef{Type: datastore.ContractType(tokenExpansionInput[evmChainSelC].DeployTokenPoolInput.PoolType), Version: tokenExpansionInput[evmChainSelC].DeployTokenPoolInput.TokenPoolVersion}
	poolC, err := datastore_utils.FindAndFormatRef(env.DataStore, fltrC, evmChainSelC, datastore_utils_evm.ToEVMAddress)
	require.NoError(t, err)

	// Set the token pool token transfer fee config on all pools
	output, err = tokens.
		SetTokenTransferFee().
		Apply(*env, tokens.SetTokenTransferFeeInput{
			Version: utils.Version_2_0_0,
			MCMS:    NewDefaultInputForMCMS("Set token transfer fee"),
			Args: []tokens.TokenTransferFeeForSrc{
				{
					Selector: evmChainSelA,
					TokenPools: []tokens.TokenTransferFeeForPool{
						{
							PoolAddress:           poolA.Hex(),
							AllowedFinalityConfig: finality.Config{BlockDepth: 12},
							Destinations: []tokens.TokenTransferFeeForDst{
								{
									IsReset:  false,
									Selector: evmChainSelB,
									Settings: tokens.UnresolvedTokenTransferFeeArgs{
										DefaultFinalityTransferFeeBps: utils.NewOptional(uint16(100)),
									},
								},
								{
									IsReset:  false,
									Selector: evmChainSelC,
									Settings: tokens.UnresolvedTokenTransferFeeArgs{
										CustomFinalityTransferFeeBps: utils.NewOptional(uint16(10)),
									},
								},
							},
						},
					},
				},
				{
					Selector: evmChainSelB,
					TokenPools: []tokens.TokenTransferFeeForPool{
						{
							PoolAddress:           poolB.Hex(),
							AllowedFinalityConfig: finality.Config{},
							Destinations: []tokens.TokenTransferFeeForDst{
								{
									IsReset:  false,
									Selector: evmChainSelA,
									Settings: tokens.UnresolvedTokenTransferFeeArgs{
										DefaultFinalityFeeUSDCents: utils.NewOptional(uint32(200)),
									},
								},
								{
									IsReset:  false,
									Selector: evmChainSelC,
									Settings: tokens.UnresolvedTokenTransferFeeArgs{
										CustomFinalityFeeUSDCents: utils.NewOptional(uint32(20)),
									},
								},
							},
						},
					},
				},
				{
					Selector: evmChainSelC,
					TokenPools: []tokens.TokenTransferFeeForPool{
						{
							PoolAddress:           poolC.Hex(),
							AllowedFinalityConfig: finality.Config{WaitForFinality: true},
							Destinations: []tokens.TokenTransferFeeForDst{
								{
									IsReset:  false,
									Selector: evmChainSelA,
									Settings: tokens.UnresolvedTokenTransferFeeArgs{
										DestBytesOverhead: utils.NewOptional(uint32(100_000)),
										DestGasOverhead:   utils.NewOptional(uint32(10_000)),
									},
								},
								{
									IsReset:  false,
									Selector: evmChainSelB,
									Settings: tokens.UnresolvedTokenTransferFeeArgs{
										DestBytesOverhead: utils.NewOptional(uint32(200_000)),
										DestGasOverhead:   utils.NewOptional(uint32(20_000)),
									},
								},
							},
						},
					},
				},
			},
		})
	require.NoError(t, err)
	testhelpers.ProcessTimelockProposals(t, *env, output.MCMSTimelockProposals, false)

	// Get the v2.0 token adapter
	tokensV2 := evmadaptersV2_0_0.TokenAdapter{}

	// Reset the operation cache
	env.OperationsBundle = operations.NewBundle(t.Context, env.OperationsBundle.Logger, operations.NewMemoryReporter())

	// Check allowed finality config for pool A (BlockDepth=12 was set)
	reportA, err := operations.ExecuteOperation(
		env.OperationsBundle, tpopsV2_0_0.GetAllowedFinalityConfig, evmChainA,
		contract.FunctionInput[struct{}]{ChainSelector: evmChainSelA, Address: poolA, Args: struct{}{}},
	)
	require.NoError(t, err)
	require.Equal(t, finality.Config{BlockDepth: 12}.Raw(), reportA.Output)

	// Check allowed finality config for pool B (not set, defaults to WaitForFinality)
	reportB, err := operations.ExecuteOperation(
		env.OperationsBundle, tpopsV2_0_0.GetAllowedFinalityConfig, evmChainB,
		contract.FunctionInput[struct{}]{ChainSelector: evmChainSelB, Address: poolB, Args: struct{}{}},
	)
	require.NoError(t, err)
	require.Equal(t, finality.RawWaitForFinality, reportB.Output)

	// Check allowed finality config for pool C (explicitly set to 0, i.e. WaitForFinality)
	reportC, err := operations.ExecuteOperation(
		env.OperationsBundle, tpopsV2_0_0.GetAllowedFinalityConfig, evmChainC,
		contract.FunctionInput[struct{}]{ChainSelector: evmChainSelC, Address: poolC, Args: struct{}{}},
	)
	require.NoError(t, err)
	require.Equal(t, finality.RawWaitForFinality, reportC.Output)

	// Query the on-chain config for both pools
	cfgAB, err := tokensV2.GetOnchainTokenTransferFeeConfig(*env, poolA.Hex(), evmChainSelA, evmChainSelB)
	require.NoError(t, err)
	cfgAC, err := tokensV2.GetOnchainTokenTransferFeeConfig(*env, poolA.Hex(), evmChainSelA, evmChainSelC)
	require.NoError(t, err)
	cfgBA, err := tokensV2.GetOnchainTokenTransferFeeConfig(*env, poolB.Hex(), evmChainSelB, evmChainSelA)
	require.NoError(t, err)
	cfgBC, err := tokensV2.GetOnchainTokenTransferFeeConfig(*env, poolB.Hex(), evmChainSelB, evmChainSelC)
	require.NoError(t, err)
	cfgCA, err := tokensV2.GetOnchainTokenTransferFeeConfig(*env, poolC.Hex(), evmChainSelC, evmChainSelA)
	require.NoError(t, err)
	cfgCB, err := tokensV2.GetOnchainTokenTransferFeeConfig(*env, poolC.Hex(), evmChainSelC, evmChainSelB)
	require.NoError(t, err)

	// Confirm that the configs for A->B match what was set, and that sensible defaults are applied for fields that were not set
	defaultsAB := tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(evmChainSelA, evmChainSelB)
	require.Equal(t, uint16(100), cfgAB.DefaultFinalityTransferFeeBps)
	require.Equal(t, defaultsAB.CustomFinalityTransferFeeBps, cfgAB.CustomFinalityTransferFeeBps)
	require.Equal(t, defaultsAB.DefaultFinalityFeeUSDCents, cfgAB.DefaultFinalityFeeUSDCents)
	require.Equal(t, defaultsAB.CustomFinalityFeeUSDCents, cfgAB.CustomFinalityFeeUSDCents)
	require.Equal(t, defaultsAB.DestBytesOverhead, cfgAB.DestBytesOverhead)
	require.Equal(t, defaultsAB.DestGasOverhead, cfgAB.DestGasOverhead)
	require.True(t, cfgAB.IsEnabled)

	// Confirm that the configs for A->C match what was set, and that sensible defaults are applied for fields that were not set
	defaultsAC := tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(evmChainSelA, evmChainSelC)
	require.Equal(t, defaultsAC.DefaultFinalityTransferFeeBps, cfgAC.DefaultFinalityTransferFeeBps)
	require.Equal(t, uint16(10), cfgAC.CustomFinalityTransferFeeBps)
	require.Equal(t, defaultsAC.DefaultFinalityFeeUSDCents, cfgAC.DefaultFinalityFeeUSDCents)
	require.Equal(t, defaultsAC.CustomFinalityFeeUSDCents, cfgAC.CustomFinalityFeeUSDCents)
	require.Equal(t, defaultsAC.DestBytesOverhead, cfgAC.DestBytesOverhead)
	require.Equal(t, defaultsAC.DestGasOverhead, cfgAC.DestGasOverhead)
	require.True(t, cfgAC.IsEnabled)

	// Confirm that the configs for B->A match what was set, and that sensible defaults are applied for fields that were not set
	defaultsBA := tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(evmChainSelB, evmChainSelA)
	require.Equal(t, defaultsBA.DefaultFinalityTransferFeeBps, cfgBA.DefaultFinalityTransferFeeBps)
	require.Equal(t, defaultsBA.CustomFinalityTransferFeeBps, cfgBA.CustomFinalityTransferFeeBps)
	require.Equal(t, uint32(200), cfgBA.DefaultFinalityFeeUSDCents)
	require.Equal(t, defaultsBA.CustomFinalityFeeUSDCents, cfgBA.CustomFinalityFeeUSDCents)
	require.Equal(t, defaultsBA.DestBytesOverhead, cfgBA.DestBytesOverhead)
	require.Equal(t, defaultsBA.DestGasOverhead, cfgBA.DestGasOverhead)
	require.True(t, cfgBA.IsEnabled)

	// Confirm that the configs for B->C match what was set, and that sensible defaults are applied for fields that were not set
	defaultsBC := tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(evmChainSelB, evmChainSelC)
	require.Equal(t, defaultsBC.DefaultFinalityTransferFeeBps, cfgBC.DefaultFinalityTransferFeeBps)
	require.Equal(t, defaultsBC.CustomFinalityTransferFeeBps, cfgBC.CustomFinalityTransferFeeBps)
	require.Equal(t, defaultsBC.DefaultFinalityFeeUSDCents, cfgBC.DefaultFinalityFeeUSDCents)
	require.Equal(t, uint32(20), cfgBC.CustomFinalityFeeUSDCents)
	require.Equal(t, defaultsBC.DestBytesOverhead, cfgBC.DestBytesOverhead)
	require.Equal(t, defaultsBC.DestGasOverhead, cfgBC.DestGasOverhead)
	require.True(t, cfgBC.IsEnabled)

	// Confirm that the configs for C->A match what was set, and that sensible defaults are applied for fields that were not set
	defaultsCA := tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(evmChainSelC, evmChainSelA)
	require.Equal(t, defaultsCA.DefaultFinalityTransferFeeBps, cfgCA.DefaultFinalityTransferFeeBps)
	require.Equal(t, defaultsCA.CustomFinalityTransferFeeBps, cfgCA.CustomFinalityTransferFeeBps)
	require.Equal(t, defaultsCA.DefaultFinalityFeeUSDCents, cfgCA.DefaultFinalityFeeUSDCents)
	require.Equal(t, defaultsCA.CustomFinalityFeeUSDCents, cfgCA.CustomFinalityFeeUSDCents)
	require.Equal(t, uint32(100_000), cfgCA.DestBytesOverhead)
	require.Equal(t, uint32(10_000), cfgCA.DestGasOverhead)
	require.True(t, cfgCA.IsEnabled)

	// Confirm that the configs for C->B match what was set, and that sensible defaults are applied for fields that were not set
	defaultsCB := tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(evmChainSelC, evmChainSelB)
	require.Equal(t, defaultsCB.DefaultFinalityTransferFeeBps, cfgCB.DefaultFinalityTransferFeeBps)
	require.Equal(t, defaultsCB.CustomFinalityTransferFeeBps, cfgCB.CustomFinalityTransferFeeBps)
	require.Equal(t, defaultsCB.DefaultFinalityFeeUSDCents, cfgCB.DefaultFinalityFeeUSDCents)
	require.Equal(t, defaultsCB.CustomFinalityFeeUSDCents, cfgCB.CustomFinalityFeeUSDCents)
	require.Equal(t, uint32(200_000), cfgCB.DestBytesOverhead)
	require.Equal(t, uint32(20_000), cfgCB.DestGasOverhead)
	require.True(t, cfgCB.IsEnabled)
}
