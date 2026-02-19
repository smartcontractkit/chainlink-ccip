package deployment

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/stretchr/testify/require"

	_ "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/adapters"
	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/fee_quoter"
	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fq16 "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestUpdateToFeeQuoter_1_7(t *testing.T) {
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
		fqInput[chainSel] = deployops.UpdateFeeQuoterInputPerChain{
			ImportFeeQuoterConfigFromVersions: []*semver.Version{
				semver.MustParse("1.6.0"),
			},
			FeeQuoterVersion: semver.MustParse("1.7.0"),
			RampsVersion:     semver.MustParse("1.6.0"),
		}
	}
	out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
	require.NoError(t, out.DataStore.Merge(e.DataStore))
	e.DataStore = out.DataStore.Seal()
	chain1 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.ETHEREUM_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.ETHEREUM_MAINNET.Selector),
	}
	chain2 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.AVALANCHE_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
		FeeQuoterDestChainConfig: lanesapi.DefaultFeeQuoterDestChainConfig(true, chain_selectors.AVALANCHE_MAINNET.Selector),
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
	fqReg := deployops.GetFQAndRampUpdaterRegistry()
	// now update to FeeQuoter 1.7.0
	fqUpdateChangeset := deployops.UpdateFeeQuoterChangeset(fqReg, nil)
	out, err = fqUpdateChangeset.Apply(*e, deployops.UpdateFeeQuoterInput{
		Chains: fqInput,
	})
	require.NoError(t, err, "Failed to apply UpdateFeeQuoterChangeset changeset")
	// update datastore with changeset output
	require.NoError(t, out.DataStore.Merge(e.DataStore), "Failed to merge changeset output datastore")
	e.DataStore = out.DataStore.Seal()
	for _, chainSel := range chains {
		chain := e.BlockChains.EVMChains()[chainSel]

		// get fee quoter address and check config
		fq17AddrRefs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chainSel),
			datastore.AddressRefByType(datastore.ContractType(fqops.ContractType)),
			datastore.AddressRefByVersion(fqops.Version),
		)
		require.Len(t, fq17AddrRefs, 1, "Expected exactly 1 FeeQuoter address ref for chain selector %d", chainSel)
		fq17Addr := common.HexToAddress(fq17AddrRefs[0].Address)
		fq17Contract, err := fee_quoter.NewFeeQuoter(fq17Addr, chain.Client)
		require.NoError(t, err, "Failed to instantiate FeeQuoter 1.7 contract for chain selector %d", chainSel)
		fq16AddrRefs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chainSel),
			datastore.AddressRefByType(datastore.ContractType(fq16ops.ContractType)),
			datastore.AddressRefByVersion(semver.MustParse("1.6.3")),
		)
		require.Len(t, fq16AddrRefs, 1, "Expected exactly 1 FeeQuoter address ref for version 1.6.3 and chain selector %d", chainSel)
		fq16Addr := common.HexToAddress(fq16AddrRefs[0].Address)
		// check that the new fee quoter has the same config as the old fee quoter
		fq16Contract, err := fq16.NewFeeQuoter(fq16Addr, chain.Client)
		require.NoError(t, err, "Failed to instantiate old FeeQuoter 1.6.3 contract for chain selector %d", chainSel)
		staticConfig16, err := fq16Contract.GetStaticConfig(nil)
		require.NoError(t, err, "Failed to get FeeQuoter config for old contract for chain selector %d", chainSel)
		staticConfig17, err := fq17Contract.GetStaticConfig(nil)
		require.NoError(t, err, "Failed to get FeeQuoter config for chain selector %d", chainSel)
		require.Equal(t, staticConfig16.MaxFeeJuelsPerMsg, staticConfig17.MaxFeeJuelsPerMsg, "MaxFeeJuelsPerMsg should be the same after update for chain selector %d", chainSel)
		require.Equal(t, staticConfig16.LinkToken, staticConfig17.LinkToken, "LinkToken address should be the same after update for chain selector %d", chainSel)
		updaters16, err := fq16Contract.GetAllAuthorizedCallers(nil)
		require.NoError(t, err, "Failed to get FeeQuoter dynamic config for old contract for chain selector %d", chainSel)
		updaters17, err := fq17Contract.GetAllAuthorizedCallers(nil)
		require.NoError(t, err, "Failed to get FeeQuoter dynamic config for chain selector %d", chainSel)
		require.ElementsMatch(t, updaters16, updaters17)

		var remoteChainSelector uint64
		for _, sel := range chains {
			if sel != chainSel {
				remoteChainSelector = sel
				break
			}
		}
		// check the destination chain config for the lane for few elements to make sure it was copied correctly
		destChainConfig17, err := fq17Contract.GetDestChainConfig(nil, remoteChainSelector)
		require.NoError(t, err, "Failed to get FeeQuoter dest chain config for chain selector %d", chainSel)
		destChainConfig16, err := fq16Contract.GetDestChainConfig(nil, remoteChainSelector)
		require.NoError(t, err, "Failed to get FeeQuoter dest chain config for old contract for chain selector %d", chainSel)
		require.Equal(t, destChainConfig16.IsEnabled, destChainConfig17.IsEnabled, "IsEnabled in dest chain config should be the same after update for chain selector %d", chainSel)
		require.Equal(t, destChainConfig16.DefaultTxGasLimit, destChainConfig17.DefaultTxGasLimit, "DefaultTxGasLimit in dest chain config should be the same after update for chain selector %d", chainSel)

		// get the onramp and offramp
		onRampAddrRefs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chainSel),
			datastore.AddressRefByType(datastore.ContractType(onrampops.ContractType)),
			datastore.AddressRefByVersion(onrampops.Version),
		)
		require.Len(t, onRampAddrRefs, 1, "Expected exactly 1 OnRamp address ref for chain selector %d", chainSel)
		onRampAddr := common.HexToAddress(onRampAddrRefs[0].Address)
		offRampAddrRefs := e.DataStore.Addresses().Filter(
			datastore.AddressRefByChainSelector(chainSel),
			datastore.AddressRefByType(datastore.ContractType(offrampops.ContractType)),
			datastore.AddressRefByVersion(offrampops.Version),
		)
		require.Len(t, offRampAddrRefs, 1, "Expected exactly 1 OffRamp address ref for chain selector %d", chainSel)
		offRampAddr := common.HexToAddress(offRampAddrRefs[0].Address)
		onRampContract, err := onramp.NewOnRamp(onRampAddr, chain.Client)
		require.NoError(t, err, "Failed to instantiate OnRamp contract for chain selector %d", chainSel)
		offRampContract, err := offramp.NewOffRamp(offRampAddr, chain.Client)
		require.NoError(t, err, "Failed to instantiate OffRamp contract for chain selector %d", chainSel)
		onRampDCfg, err := onRampContract.GetDynamicConfig(nil)
		require.NoError(t, err, "Failed to get OnRamp static config for chain selector %d", chainSel)
		offRampDCfg, err := offRampContract.GetDynamicConfig(nil)
		require.NoError(t, err, "Failed to get OffRamp static config for chain selector %d", chainSel)
		require.Equal(t, onRampDCfg.FeeQuoter, fq17Addr, "OnRamp should point to new FeeQuoter after update for chain selector %d", chainSel)
		require.Equal(t, offRampDCfg.FeeQuoter, fq17Addr, "OffRamp should point to new FeeQuoter after update for chain selector %d", chainSel)
	}
}
