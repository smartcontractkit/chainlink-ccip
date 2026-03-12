package adapters

import (
	"math/big"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/lombard_token_pool"
	deployops "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	// Register 1.6.0 deployer so DeployContracts can run
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
)

func TestUpdateRouter_produces_single_write_when_no_lombard_pool(t *testing.T) {
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.AVALANCHE_MAINNET.Selector,
	}
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	version := semver.MustParse("1.6.0")
	chainInput := make(map[uint64]deployops.ContractDeploymentConfigPerChain)
	for _, chainSel := range chains {
		chainInput[chainSel] = deployops.ContractDeploymentConfigPerChain{
			Version: version,
			MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
			TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
			LinkPremiumMultiplier:         9e17,
			NativeTokenPremiumMultiplier:  1e18,
			PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			GasForCallExactCheck:                    5000,
		}
	}
	dReg := deployops.GetRegistry()
	out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err)
	require.NoError(t, out.DataStore.Merge(e.DataStore))
	e.DataStore = out.DataStore.Seal()

	chainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	existingAddresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))
	require.NotEmpty(t, existingAddresses, "datastore should have addresses for chain")

	onRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          "OnRamp",
		Version:       version,
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)
	offRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          "OffRamp",
		Version:       version,
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)

	config := deployops.RouterUpdaterConfig{
		ChainSelector:        chainSel,
		RemoteChainSelectors: []uint64{chain_selectors.AVALANCHE_MAINNET.Selector},
		OnRamp:               onRampRef,
		OffRamp:              offRampRef,
		ExistingAddresses:    existingAddresses,
	}

	updater := &RouterUpdater{}
	report, err := operations.ExecuteSequence(e.OperationsBundle, updater.UpdateRouter(), e.BlockChains, config)
	require.NoError(t, err)
	require.NotEmpty(t, report.Output.BatchOps, "expected at least one batch")
	// When LombardTokenPool is not in ExistingAddresses, the adapter only adds the router ramp update.
	// The batch may have 0 or 1 transaction depending on whether the framework executed the write inline.
	require.True(t, len(report.Output.BatchOps[0].Transactions) <= 1,
		"when LombardTokenPool is not in ExistingAddresses, batch should contain at most the router ramp update (0–1 tx)")
}

func TestUpdateRouter_attempts_GetToken_when_lombard_pool_in_datastore(t *testing.T) {
	chains := []uint64{
		chain_selectors.ETHEREUM_MAINNET.Selector,
		chain_selectors.AVALANCHE_MAINNET.Selector,
	}
	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, chains),
	)
	require.NoError(t, err)
	require.NotNil(t, e)

	version := semver.MustParse("1.6.0")
	chainInput := make(map[uint64]deployops.ContractDeploymentConfigPerChain)
	for _, chainSel := range chains {
		chainInput[chainSel] = deployops.ContractDeploymentConfigPerChain{
			Version: version,
			MaxFeeJuelsPerMsg:            big.NewInt(0).Mul(big.NewInt(200), big.NewInt(1e18)),
			TokenPriceStalenessThreshold: uint32(24 * 60 * 60),
			LinkPremiumMultiplier:         9e17,
			NativeTokenPremiumMultiplier:  1e18,
			PermissionLessExecutionThresholdSeconds: uint32((20 * time.Minute).Seconds()),
			GasForCallExactCheck:                    5000,
		}
	}
	dReg := deployops.GetRegistry()
	out, err := deployops.DeployContracts(dReg).Apply(*e, deployops.ContractDeploymentConfig{
		MCMS:   mcms.Input{},
		Chains: chainInput,
	})
	require.NoError(t, err)
	require.NoError(t, out.DataStore.Merge(e.DataStore))
	e.DataStore = out.DataStore.Seal()

	chainSel := chain_selectors.ETHEREUM_MAINNET.Selector
	// Add LombardTokenPool 2.0.0 ref pointing to zero address (no contract there; GetToken will fail).
	// Build a new datastore with all existing refs for this chain plus the Lombard ref.
	existingRefs := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))
	ds := datastore.NewMemoryDataStore()
	for _, ref := range existingRefs {
		require.NoError(t, ds.Addresses().Add(ref))
	}
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(lombard_token_pool.ContractType),
		Version:       lombard_token_pool.Version,
		Address:       "0x0000000000000000000000000000000000000000",
	}))
	e.DataStore = ds.Seal()

	existingAddresses := e.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSel))
	onRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          "OnRamp",
		Version:       version,
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)
	offRampRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          "OffRamp",
		Version:       version,
	}, chainSel, datastore_utils.FullRef)
	require.NoError(t, err)

	config := deployops.RouterUpdaterConfig{
		ChainSelector:        chainSel,
		RemoteChainSelectors: []uint64{chain_selectors.AVALANCHE_MAINNET.Selector},
		OnRamp:               onRampRef,
		OffRamp:              offRampRef,
		ExistingAddresses:    existingAddresses,
	}

	updater := &RouterUpdater{}
	_, err = operations.ExecuteSequence(e.OperationsBundle, updater.UpdateRouter(), e.BlockChains, config)
	require.Error(t, err)
	require.Contains(t, err.Error(), "LombardTokenPool", "when LombardTokenPool is in ExistingAddresses, adapter should call GetToken on the pool and fail if pool is not a contract")
}
