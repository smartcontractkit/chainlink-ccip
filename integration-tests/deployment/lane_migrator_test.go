package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestMigrateTo1_6_0(t *testing.T) {
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
	out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
	require.NoError(t, out.DataStore.Merge(e.DataStore))
	e.DataStore = out.DataStore.Seal()
	// find TestRouter
	testRouterRefChain1, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			Type:    datastore.ContractType(routerops.TestRouterContractType),
			Version: routerops.Version,
		}, chain_selectors.ETHEREUM_MAINNET.Selector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	require.NotEmpty(t, testRouterRefChain1)
	chain1 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.ETHEREUM_MAINNET.Selector),
		Router:                   testRouterRefChain1.Bytes(),
	}
	testRouterRefChain2, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			Type:    datastore.ContractType(routerops.TestRouterContractType),
			Version: routerops.Version,
		}, chain_selectors.AVALANCHE_MAINNET.Selector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	require.NotEmpty(t, testRouterRefChain2)
	chain2 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.AVALANCHE_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.AVALANCHE_MAINNET.Selector),
		Router:                   testRouterRefChain2.Bytes(),
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

	// migrate to prod router
	migratorReg := deployops.GetLaneMigratorRegistry()
	cs := deployops.LaneMigrateToNewVersionChangeset(migratorReg, mcmsRegistry)
	_, err = cs.Apply(*e, deployops.LaneMigratorConfig{
		Input: map[uint64]deployops.LaneMigratorConfigPerChain{
			chain_selectors.ETHEREUM_MAINNET.Selector: {
				RouterVersion: semver.MustParse("1.2.0"),
				RampVersion:   version,
				RemoteChains: []uint64{
					chain_selectors.AVALANCHE_MAINNET.Selector,
				},
			},
			chain_selectors.AVALANCHE_MAINNET.Selector: {
				RouterVersion: semver.MustParse("1.2.0"),
				RampVersion:   version,
				RemoteChains: []uint64{
					chain_selectors.ETHEREUM_MAINNET.Selector,
				},
			},
		},
	})
	require.NoError(t, err, "Failed to apply LaneMigrateToNewVersionChangeset")
	// query the router to ensure it has the correct onRamp and offRamp addresses
	evmChain1 := e.BlockChains.EVMChains()[chain_selectors.ETHEREUM_MAINNET.Selector]
	routerAddr, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			Type:    "Router",
			Version: semver.MustParse("1.2.0"),
		}, chain_selectors.ETHEREUM_MAINNET.Selector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	onRampAddr, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			Type:    "OnRamp",
			Version: semver.MustParse("1.6.0"),
		}, chain_selectors.ETHEREUM_MAINNET.Selector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	offRampAddr, err := datastore_utils.FindAndFormatRef(
		e.DataStore,
		datastore.AddressRef{
			Type:    "OffRamp",
			Version: semver.MustParse("1.6.0"),
		}, chain_selectors.ETHEREUM_MAINNET.Selector, evm_datastore_utils.ToEVMAddress)
	require.NoError(t, err)
	// query router
	routerC, err := router.NewRouter(routerAddr, evmChain1.Client)
	require.NoError(t, err)
	onRamp, err := routerC.GetOnRamp(nil, chain_selectors.AVALANCHE_MAINNET.Selector)
	require.NoError(t, err)
	require.Equal(t, onRampAddr, onRamp)
	offRamps, err := routerC.GetOffRamps(nil)
	require.NoError(t, err)
	require.Contains(t, offRamps, router.RouterOffRamp{
		SourceChainSelector: chain_selectors.AVALANCHE_MAINNET.Selector,
		OffRamp:             offRampAddr,
	})
}
