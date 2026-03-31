package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	evmadaptersV1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evmadaptersV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	evmseqV1_6_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestUpdateFeeQuoterDestsV1_6_0(t *testing.T) {
	v1_6_0, err := semver.NewVersion("1.6.0")
	require.NoError(t, err)

	src := chainsel.TEST_90000001.Selector
	dst := chainsel.TEST_90000002.Selector
	chains := []uint64{src, dst}

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err)

	// Register adapters
	evmAdapter := evmseqV1_6_0.EVMAdapter{}
	deployRegistry := deploy.GetRegistry()
	deployRegistry.RegisterDeployer(chainsel.FamilyEVM, deploy.MCMSVersion, &evmadaptersV1_0_0.EVMDeployer{})

	feesRegistry := fees.GetRegistry()
	evmFeesAdapter := evmadaptersV1_6_0.NewFeesAdapter(&evmAdapter)
	feesRegistry.RegisterFeeAdapter(chainsel.FamilyEVM, v1_6_0, evmFeesAdapter)

	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader(chainsel.FamilyEVM, &evmadaptersV1_0_0.EVMMCMSReader{})

	// Deploy contracts
	chainInput := make(map[uint64]deploy.ContractDeploymentConfigPerChain)
	for _, sel := range chains {
		chainInput[sel] = deploy.ContractDeploymentConfigPerChain{
			Version:                                 v1_6_0,
			MaxFeeJuelsPerMsg:                       big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
			TokenPriceStalenessThreshold:            uint32(24 * 60 * 60),
			LinkPremiumMultiplier:                   9e17,
			NativeTokenPremiumMultiplier:            1e18,
			PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			GasForCallExactCheck:                    uint16(5000),
		}
	}
	out, err := deploy.DeployContracts(deployRegistry).Apply(*e, deploy.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err)
	MergeAddresses(t, e, out.DataStore)

	// Connect chains to set initial FQ dest config
	chain1 := lanes.ChainDefinition{
		Selector: src,
		GasPrice: big.NewInt(1e9),
	}
	chain2 := lanes.ChainDefinition{
		Selector: dst,
		GasPrice: big.NewInt(1e9),
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
	require.NoError(t, err)

	// Read the initial on-chain config set by ConnectChains
	initialCfg, err := evmFeesAdapter.GetOnchainDestChainConfig(*e, src, dst)
	require.NoError(t, err)
	require.True(t, initialCfg.IsEnabled)

	originalMaxDataBytes := initialCfg.MaxDataBytes
	originalDestGasOverhead := initialCfg.DestGasOverhead

	// Now update only MaxDataBytes via override (upsert semantics)
	newMaxDataBytes := uint32(99_999)
	require.NotEqual(t, originalMaxDataBytes, newMaxDataBytes)

	override := lanes.FeeQuoterDestChainConfigOverride(func(c *lanes.FeeQuoterDestChainConfig) {
		c.MaxDataBytes = newMaxDataBytes
	})

	_, err = fees.
		UpdateFeeQuoterDests(feesRegistry, mcmsRegistry).
		Apply(*e, fees.UpdateFeeQuoterDestsInput{
			Version: v1_6_0,
			MCMS:    mcms.Input{},
			Args: []fees.DestChainConfigForSrc{
				{
					Selector: src,
					Settings: []fees.DestChainConfigForDst{
						{
							Selector: dst,
							Override: &override,
						},
					},
				},
			},
		})
	require.NoError(t, err)

	// Read back the config and verify the upsert
	updatedCfg, err := evmFeesAdapter.GetOnchainDestChainConfig(*e, src, dst)
	require.NoError(t, err)
	require.True(t, updatedCfg.IsEnabled)
	require.Equal(t, newMaxDataBytes, updatedCfg.MaxDataBytes, "MaxDataBytes should be updated")
	require.Equal(t, originalDestGasOverhead, updatedCfg.DestGasOverhead, "DestGasOverhead should be unchanged")
}
