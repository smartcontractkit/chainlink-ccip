package deployment

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	fq16ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fq163ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	fqops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
	offrampops "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/operations/offramp"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	lanesapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/testhelpers"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

func TestUpdateToFeeQuoter_2_0(t *testing.T) {
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
			FeeQuoterVersion: semver.MustParse("2.0.0"),
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
	}
	chain2 := lanesapi.ChainDefinition{
		Selector:                 chain_selectors.AVALANCHE_MAINNET.Selector,
		GasPrice:                 big.NewInt(1e9),
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
	// now update to FeeQuoter 2.0.0
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
	for _, chainSel := range chains {
		fqUpgradeValidation(t, e, chainSel, chains, true)
	}
	// do this to reset cached executions
	bundle := operations.NewBundle(
		func() context.Context { return context.Background() },
		e.Logger,
		operations.NewMemoryReporter(),
	)
	e.OperationsBundle = bundle
	// downgrade back to 1.6.3 to make sure the changeset is reversible
	for _, chainSel := range chains {
		fqInput[chainSel] = deployops.UpdateFeeQuoterInputPerChain{
			FeeQuoterVersion: fq163ops.Version,
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

	for _, chainSel := range chains {
		fqUpgradeValidation(t, e, chainSel, chains, false)
	}
}

func fqUpgradeValidation(t *testing.T, e *cldf.Environment, chainSel uint64, chains []uint64, expected17fq bool) {
	chain := e.BlockChains.EVMChains()[chainSel]

	// get fee quoter address and check config
	fq17AddrRefs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSel),
		datastore.AddressRefByType(datastore.ContractType(fqops.ContractType)),
		datastore.AddressRefByVersion(fqops.Version),
	)
	require.Len(t, fq17AddrRefs, 1, "Expected exactly 1 FeeQuoter address ref for chain selector %d", chainSel)
	fq17Addr := common.HexToAddress(fq17AddrRefs[0].Address)
	fq17Contract, err := fqops.NewFeeQuoterContract(fq17Addr, chain.Client)
	require.NoError(t, err, "Failed to instantiate FeeQuoter 2.0 contract for chain selector %d", chainSel)
	fq16AddrRefs := e.DataStore.Addresses().Filter(
		datastore.AddressRefByChainSelector(chainSel),
		datastore.AddressRefByType(datastore.ContractType(fq16ops.ContractType)),
		datastore.AddressRefByVersion(fq163ops.Version),
	)
	require.Len(t, fq16AddrRefs, 1, "Expected exactly 1 FeeQuoter address ref for version 1.6.0 and chain selector %d", chainSel)
	fq16Addr := common.HexToAddress(fq16AddrRefs[0].Address)
	// check that the new fee quoter has the same config as the old fee quoter
	fq16Contract, err := fq163ops.NewFeeQuoterContract(fq16Addr, chain.Client)
	require.NoError(t, err, "Failed to instantiate old FeeQuoter 1.6.0 contract for chain selector %d", chainSel)
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
	for _, caller := range updaters16 {
		require.Contains(t, updaters17, caller, "FQ 2.0 should contain all authorized callers from FQ 1.6 for chain selector %d", chainSel)
	}

	if expected17fq {
		timelockRef := datastore_utils.GetAddressRef(
			e.DataStore.Addresses().Filter(),
			chainSel,
			common_utils.RBACTimelock,
			semver.MustParse("1.0.0"),
			common_utils.CLLQualifier,
		)
		require.NotEmpty(t, timelockRef.Address, "Expected timelock address for chain selector %d", chainSel)
		timelockAddr := common.HexToAddress(timelockRef.Address)
		require.Contains(t, updaters17, timelockAddr, "FQ 2.0 should have timelock as a price updater for chain selector %d", chainSel)
		fq17Owner, err := fq17Contract.Owner(nil)
		require.NoError(t, err, "Failed to get FeeQuoter 2.0 owner for chain selector %d", chainSel)
		require.Equal(t, timelockAddr, fq17Owner, "FeeQuoter 2.0 should be owned by timelock after ownership transfer for chain selector %d", chainSel)
	}

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
	expectedFQAddr := fq17Addr
	if !expected17fq {
		expectedFQAddr = fq16Addr
	}
	require.Equal(t, onRampDCfg.FeeQuoter, expectedFQAddr, "OnRamp should point to expected FeeQuoter after update for chain selector %d", chainSel)
	require.Equal(t, offRampDCfg.FeeQuoter, expectedFQAddr, "OffRamp should point to expected FeeQuoter after update for chain selector %d", chainSel)
}
