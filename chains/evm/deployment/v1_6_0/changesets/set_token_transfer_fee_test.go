package changesets_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	adaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"
)

func TestSetTokenTransferFeeV1_6_0(t *testing.T) {
	// Define alias for v1.6.0
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)

	// Setup test environment
	src := chainsel.TEST_90000001.Selector
	dst := chainsel.TEST_90000002.Selector
	env, err := environment.New(t.Context(), environment.WithEVMSimulated(t, []uint64{src, dst}))
	require.NoError(t, err)

	// Initialize an EVM v1.6.0 adapter
	evmAdapter := evmseqV1_6_0.EVMAdapter{}

	// Configure deployment registry
	deployRegistry := deploy.GetRegistry()
	deployAdapter := &adapters.EVMDeployer{}
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, deployAdapter)

	// Configure fees registry
	feesRegistry := fees.GetRegistry()
	feesAdapter := adaptersV1_6_0.NewFeesAdapter(&evmAdapter)
	feesRegistry.RegisterFeeAdapter(chainsel.FamilyEVM, v1_6_0, feesAdapter)

	// Configure MCMS registry
	mcmsRegistry := changesets.GetRegistry()
	mcmsAdapter := &adapters.EVMMCMSReader{}
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, mcmsAdapter)

	// Use the same deployment config for all selectors
	dummyDeployConfig := deploy.ContractDeploymentConfigPerChain{
		Version:                                 v1_6_0,
		MaxFeeJuelsPerMsg:                       big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
		TokenPriceStalenessThreshold:            uint32(24 * 60 * 60),
		NativeTokenPremiumMultiplier:            1e18, // 1.0 ETH
		LinkPremiumMultiplier:                   9e17, // 0.9 ETH
		PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
		GasForCallExactCheck:                    uint16(5000),
	}

	// Deploy FeeQuoter + other contracts
	output, err := deploy.DeployContracts(deployRegistry).Apply(*env, deploy.ContractDeploymentConfig{
		MCMS: mcms.Input{},
		Chains: map[uint64]deploy.ContractDeploymentConfigPerChain{
			src: dummyDeployConfig,
			dst: dummyDeployConfig,
		},
	})
	require.NoError(t, err)

	// Get the address of the LINK token on the source chain
	srcLinkRef, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(src,
			datastore.ContractType(link.ContractType),
			semver.MustParse("1.0.0"),
			"", // no qualifier is needed for EVM LINK token
		),
	)
	require.NoError(t, err)

	// Get the address of the LINK token on the destination chain
	dstLinkRef, err := output.DataStore.Addresses().Get(
		datastore.NewAddressRefKey(dst,
			datastore.ContractType(link.ContractType),
			semver.MustParse("1.0.0"),
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
	srcCfg, err := feesAdapter.GetOnchainTokenTransferFeeConfig(*env, src, dst, srcLinkRef.Address)
	require.NoError(t, err)
	srcSensibleDefaults := feesAdapter.GetDefaultTokenTransferFeeConfig(src, dst)
	require.Equal(t, srcCfg.DestBytesOverhead, srcSensibleDefaults.DestBytesOverhead)
	require.Equal(t, srcCfg.DestGasOverhead, uint32(120_000))
	require.Equal(t, srcCfg.MinFeeUSDCents, srcSensibleDefaults.MinFeeUSDCents)
	require.Equal(t, srcCfg.MaxFeeUSDCents, srcSensibleDefaults.MaxFeeUSDCents)
	require.Equal(t, srcCfg.DeciBps, srcSensibleDefaults.DeciBps)
	require.True(t, srcCfg.IsEnabled)

	// Confirm that the config was correctly set on the destination
	dstCfg, err := feesAdapter.GetOnchainTokenTransferFeeConfig(*env, dst, src, dstLinkRef.Address)
	require.NoError(t, err)
	dstSensibleDefaults := feesAdapter.GetDefaultTokenTransferFeeConfig(dst, src)
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
	srcCfg, err = feesAdapter.GetOnchainTokenTransferFeeConfig(*env, src, dst, srcLinkRef.Address)
	require.NoError(t, err)
	require.False(t, srcCfg.IsEnabled)

	// Confirm that the config was disabled on the destination
	dstCfg, err = feesAdapter.GetOnchainTokenTransferFeeConfig(*env, dst, src, dstLinkRef.Address)
	require.NoError(t, err)
	require.False(t, dstCfg.IsEnabled)
}
