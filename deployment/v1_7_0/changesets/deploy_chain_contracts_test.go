package changesets_test

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/mock_receiver"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/rmn_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/weth"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/link_token"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/registry_module_owner_custom"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	cs_core "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
)

func testContractParamsPerChain(t *testing.T, create2Factory common.Address) changesets.DeployChainContractsFromTopologyCfgPerChain {
	t.Helper()
	usdPerLink, ok := new(big.Int).SetString("15000000000000000000", 10)
	require.True(t, ok, "usdPerLink must parse")
	usdPerWeth, ok := new(big.Int).SetString("2000000000000000000000", 10)
	require.True(t, ok, "usdPerWeth must parse")

	return changesets.DeployChainContractsFromTopologyCfgPerChain{
		CREATE2Factory: create2Factory,
		RMNRemote: sequences.RMNRemoteParams{
			Version: rmn_remote.Version,
		},
		OffRamp: sequences.OffRampParams{
			Version:                   offramp.Version,
			GasForCallExactCheck:      5_000,
			MaxGasBufferToUpdateState: 12_000,
		},
		OnRamp: sequences.OnRampParams{
			Version:               onramp.Version,
			FeeAggregator:         common.HexToAddress("0x01"),
			MaxUSDCentsPerMessage: 100_00,
		},
		FeeQuoter: sequences.FeeQuoterParams{
			Version:                        fee_quoter.Version,
			MaxFeeJuelsPerMsg:              big.NewInt(1e18),
			LINKPremiumMultiplierWeiPerEth: 9e17,
			WETHPremiumMultiplierWeiPerEth: 1e18,
			USDPerLINK:                     usdPerLink,
			USDPerWETH:                     usdPerWeth,
		},
		Executors: []sequences.ExecutorParams{
			{
				Version:       executor.Version,
				MaxCCVsPerMsg: 10,
				DynamicConfig: executor.SetDynamicConfigArgs{
					FeeAggregator:         common.HexToAddress("0x01"),
					MinBlockConfirmations: 1,
				},
				Qualifier: "default",
			},
		},
		MockReceivers: []sequences.MockReceiverParams{
			{
				Version:   semver.MustParse("1.7.0"),
				Qualifier: "default",
			},
		},
	}
}

func deployCreate2Factory(t *testing.T, e deployment.Environment, chain cldfevm.Chain) common.Address {
	t.Helper()
	ref, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, chain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("1.7.0")),
			ChainSelector:  chain.Selector,
			Args: create2_factory.ConstructorArgs{
				AllowList: []common.Address{chain.DeployerKey.From},
			},
		}, nil,
	)
	require.NoError(t, err)
	return common.HexToAddress(ref.Address)
}

func TestDeployChainContractsFromTopology_Validate(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	validTopology := newTestTopology()

	tests := []struct {
		name        string
		cfg         changesets.DeployChainContractsFromTopologyCfg
		expectedErr string
	}{
		{
			name: "rejects nil topology",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				ChainSelectors: []uint64{sel1},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
			},
			expectedErr: "topology is required",
		},
		{
			name: "rejects empty committees",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology: &topology.EnvironmentTopology{
					NOPTopology: &topology.NOPTopology{
						Committees: map[string]topology.CommitteeConfig{},
					},
				},
				ChainSelectors: []uint64{sel1},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
			},
			expectedErr: "no committees defined in topology",
		},
		{
			name: "rejects empty chain selectors",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
			},
			expectedErr: "at least one chain selector is required",
		},
		{
			name: "rejects unknown chain selector",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{99999},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
			},
			expectedErr: "chain selector 99999 is not available in environment",
		},
		{
			name: "rejects zero CREATE2Factory in default config",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.Address{},
				},
			},
			expectedErr: "CREATE2Factory address is required",
		},
		{
			name: "rejects zero CREATE2Factory in chain override",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
				ChainCfgs: map[uint64]changesets.DeployChainContractsFromTopologyCfgPerChain{
					sel1: {CREATE2Factory: common.Address{}},
				},
			},
			expectedErr: "CREATE2Factory address is required",
		},
		{
			name: "rejects duplicate chain selectors",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1, sel1},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
			},
			expectedErr: "duplicate chain selector",
		},
		{
			name: "rejects ChainCfgs key not in ChainSelectors",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
				ChainCfgs: map[uint64]changesets.DeployChainContractsFromTopologyCfgPerChain{
					77777: {CREATE2Factory: common.HexToAddress("0x5678")},
				},
			},
			expectedErr: "ChainCfgs contains selector 77777 which is not in ChainSelectors",
		},
		{
			name: "accepts valid config",
			cfg: changesets.DeployChainContractsFromTopologyCfg{
				Topology:       validTopology,
				ChainSelectors: []uint64{sel1},
				DefaultCfg: changesets.DeployChainContractsFromTopologyCfgPerChain{
					CREATE2Factory: common.HexToAddress("0x1234"),
				},
			},
		},
	}

	env := newTestEnvironment(t, defaultSelectors)
	mcmsRegistry := cs_core.GetRegistry()
	cs := changesets.DeployChainContractsFromTopology(mcmsRegistry)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(env, cs_core.WithMCMS[changesets.DeployChainContractsFromTopologyCfg]{
				MCMS: mcms.Input{},
				Cfg:  tc.cfg,
			})
			if tc.expectedErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var allExpectedContractTypes = []datastore.ContractType{
	datastore.ContractType(weth.ContractType),
	datastore.ContractType(link_token.ContractType),
	datastore.ContractType(rmn_remote.ContractType),
	datastore.ContractType(rmn_proxy.ContractType),
	datastore.ContractType(router.ContractType),
	datastore.ContractType(token_admin_registry.ContractType),
	datastore.ContractType(registry_module_owner_custom.ContractType),
	datastore.ContractType(fee_quoter.ContractType),
	datastore.ContractType(offramp.ContractType),
	datastore.ContractType(onramp.ContractType),
	datastore.ContractType(committee_verifier.ContractType),
	datastore.ContractType(executor.ContractType),
	datastore.ContractType(executor.ProxyType),
	datastore.ContractType(mock_receiver.ContractType),
}

func assertAllExpectedContractsDeployed(t *testing.T, addrs []datastore.AddressRef, sel uint64) {
	t.Helper()

	deployed := map[datastore.ContractType]bool{}
	for _, ref := range addrs {
		if ref.ChainSelector == sel {
			deployed[ref.Type] = true
		}
	}

	for _, ct := range allExpectedContractTypes {
		assert.True(t, deployed[ct], "expected contract type %s on chain %d", ct, sel)
	}
}

func TestDeployChainContractsFromTopology_Apply_SingleChain(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	env, evmChains := testutils.NewSimulatedEVMEnvironment(t, []uint64{sel})
	create2Addr := deployCreate2Factory(t, env, evmChains[0])

	topology := newTestTopology()
	mcmsRegistry := cs_core.GetRegistry()
	cs := changesets.DeployChainContractsFromTopology(mcmsRegistry)

	out, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsFromTopologyCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsFromTopologyCfg{
			Topology:       topology,
			ChainSelectors: []uint64{sel},
			DefaultCfg:     testContractParamsPerChain(t, create2Addr),
		},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	assertAllExpectedContractsDeployed(t, addrs, sel)
}

func TestDeployChainContractsFromTopology_Apply_MultipleChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	selectors := []uint64{sel1, sel2}
	env, evmChains := testutils.NewSimulatedEVMEnvironment(t, selectors)

	create2Addr1 := deployCreate2Factory(t, env, evmChains[0])
	create2Addr2 := deployCreate2Factory(t, env, evmChains[1])

	topology := newTestTopology()
	mcmsRegistry := cs_core.GetRegistry()
	cs := changesets.DeployChainContractsFromTopology(mcmsRegistry)

	out, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsFromTopologyCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsFromTopologyCfg{
			Topology:       topology,
			ChainSelectors: selectors,
			DefaultCfg:     testContractParamsPerChain(t, create2Addr1),
			ChainCfgs: map[uint64]changesets.DeployChainContractsFromTopologyCfgPerChain{
				sel2: testContractParamsPerChain(t, create2Addr2),
			},
		},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	assertAllExpectedContractsDeployed(t, addrs, sel1)
	assertAllExpectedContractsDeployed(t, addrs, sel2)
}

func TestDeployChainContractsFromTopology_Apply_ExistingAddressesArePreserved(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	env, evmChains := testutils.NewSimulatedEVMEnvironment(t, []uint64{sel})
	create2Addr := deployCreate2Factory(t, env, evmChains[0])

	ds := datastore.NewMemoryDataStore()
	wethAddr := common.HexToAddress("0xDEAD")
	linkAddr := common.HexToAddress("0xBEEF")
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: sel,
		Type:          datastore.ContractType(weth.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       wethAddr.Hex(),
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: sel,
		Type:          datastore.ContractType(link_token.ContractType),
		Version:       semver.MustParse("1.0.0"),
		Address:       linkAddr.Hex(),
	}))
	env.DataStore = ds.Seal()

	topology := newTestTopology()
	mcmsRegistry := cs_core.GetRegistry()
	cs := changesets.DeployChainContractsFromTopology(mcmsRegistry)

	out, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsFromTopologyCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsFromTopologyCfg{
			Topology:       topology,
			ChainSelectors: []uint64{sel},
			DefaultCfg:     testContractParamsPerChain(t, create2Addr),
		},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	assertAllExpectedContractsDeployed(t, addrs, sel)

	for _, ref := range addrs {
		if ref.ChainSelector == sel && ref.Type == datastore.ContractType(weth.ContractType) {
			assert.Equal(t, wethAddr.Hex(), ref.Address, "existing WETH address should be preserved")
		}
		if ref.ChainSelector == sel && ref.Type == datastore.ContractType(link_token.ContractType) {
			assert.Equal(t, linkAddr.Hex(), ref.Address, "existing LINK address should be preserved")
		}
	}
}

func TestDeployChainContractsFromTopology_Apply_PerChainOverrideIsUsed(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	selectors := []uint64{sel1, sel2}
	env, evmChains := testutils.NewSimulatedEVMEnvironment(t, selectors)

	create2Addr1 := deployCreate2Factory(t, env, evmChains[0])
	create2Addr2 := deployCreate2Factory(t, env, evmChains[1])

	topology := newTestTopology()
	mcmsRegistry := cs_core.GetRegistry()
	cs := changesets.DeployChainContractsFromTopology(mcmsRegistry)

	overrideCfg := testContractParamsPerChain(t, create2Addr2)
	overrideCfg.OnRamp.MaxUSDCentsPerMessage = 999_99

	out, err := cs.Apply(env, cs_core.WithMCMS[changesets.DeployChainContractsFromTopologyCfg]{
		MCMS: mcms.Input{},
		Cfg: changesets.DeployChainContractsFromTopologyCfg{
			Topology:       topology,
			ChainSelectors: selectors,
			DefaultCfg:     testContractParamsPerChain(t, create2Addr1),
			ChainCfgs: map[uint64]changesets.DeployChainContractsFromTopologyCfgPerChain{
				sel2: overrideCfg,
			},
		},
	})
	require.NoError(t, err)

	addrs, err := out.DataStore.Addresses().Fetch()
	require.NoError(t, err)

	assertAllExpectedContractsDeployed(t, addrs, sel1)
	assertAllExpectedContractsDeployed(t, addrs, sel2)
}

func TestBuildCommitteeVerifierParams_MapsAllCommittees(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)
	topology := &topology.EnvironmentTopology{
		NOPTopology: &topology.NOPTopology{
			Committees: map[string]topology.CommitteeConfig{
				"alpha": {
					Qualifier:        "alpha",
					VerifierVersion:  semver.MustParse("1.7.0"),
					StorageLocations: []string{"https://store1.test"},
					ChainConfigs: map[string]topology.ChainCommitteeConfig{
						selStr: {
							NOPAliases:     []string{"nop-1"},
							Threshold:      1,
							FeeAggregator:  "0x0000000000000000000000000000000000000001",
							AllowlistAdmin: "0x0000000000000000000000000000000000000002",
						},
					},
				},
				"beta": {
					Qualifier:        "beta",
					VerifierVersion:  semver.MustParse("1.7.0"),
					StorageLocations: []string{"https://store2.test"},
					ChainConfigs: map[string]topology.ChainCommitteeConfig{
						selStr: {
							NOPAliases:     []string{"nop-1"},
							Threshold:      1,
							FeeAggregator:  "0x0000000000000000000000000000000000000003",
							AllowlistAdmin: "0x0000000000000000000000000000000000000004",
						},
					},
				},
			},
		},
	}

	params, err := changesets.BuildCommitteeVerifierParams(topology, sel)
	require.NoError(t, err)
	assert.Len(t, params, 2)

	qualifiers := map[string]sequences.CommitteeVerifierParams{}
	for _, p := range params {
		qualifiers[p.Qualifier] = p
	}

	alpha, ok := qualifiers["alpha"]
	require.True(t, ok)
	assert.Equal(t, common.HexToAddress("0x01"), alpha.FeeAggregator)
	assert.Equal(t, common.HexToAddress("0x02"), alpha.AllowlistAdmin)
	assert.Equal(t, []string{"https://store1.test"}, alpha.StorageLocations)

	beta, ok := qualifiers["beta"]
	require.True(t, ok)
	assert.Equal(t, common.HexToAddress("0x03"), beta.FeeAggregator)
	assert.Equal(t, common.HexToAddress("0x04"), beta.AllowlistAdmin)
	assert.Equal(t, []string{"https://store2.test"}, beta.StorageLocations)
}

func TestBuildCommitteeVerifierParams_SkipsCommitteesWithoutChainConfig(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)
	otherSel := chainsel.TEST_90000002.Selector
	topology := &topology.EnvironmentTopology{
		NOPTopology: &topology.NOPTopology{
			Committees: map[string]topology.CommitteeConfig{
				"present": {
					Qualifier:       "present",
					VerifierVersion: semver.MustParse("1.7.0"),
					ChainConfigs: map[string]topology.ChainCommitteeConfig{
						selStr: {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "0x0000000000000000000000000000000000000001",
						},
					},
				},
				"absent": {
					Qualifier:       "absent",
					VerifierVersion: semver.MustParse("1.7.0"),
					ChainConfigs: map[string]topology.ChainCommitteeConfig{
						strconv.FormatUint(otherSel, 10): {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "0x0000000000000000000000000000000000000002",
						},
					},
				},
			},
		},
	}

	params, err := changesets.BuildCommitteeVerifierParams(topology, sel)
	require.NoError(t, err)
	assert.Len(t, params, 1)
	assert.Equal(t, "present", params[0].Qualifier)
}

func TestBuildCommitteeVerifierParams_RejectsNilNOPTopology(t *testing.T) {
	topology := &topology.EnvironmentTopology{}
	_, err := changesets.BuildCommitteeVerifierParams(topology, 1)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "NOPTopology is nil")
}

func TestBuildCommitteeVerifierParams_RejectsNilVerifierVersion(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)
	topology := &topology.EnvironmentTopology{
		NOPTopology: &topology.NOPTopology{
			Committees: map[string]topology.CommitteeConfig{
				"broken": {
					VerifierVersion: nil,
					ChainConfigs: map[string]topology.ChainCommitteeConfig{
						selStr: {
							NOPAliases:    []string{"nop-1"},
							Threshold:     1,
							FeeAggregator: "0x0000000000000000000000000000000000000001",
						},
					},
				},
			},
		},
	}
	_, err := changesets.BuildCommitteeVerifierParams(topology, sel)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nil VerifierVersion")
}

func TestBuildCommitteeVerifierParams_ReturnsEmptyForEmptyCommittees(t *testing.T) {
	topology := &topology.EnvironmentTopology{
		NOPTopology: &topology.NOPTopology{
			Committees: map[string]topology.CommitteeConfig{},
		},
	}
	params, err := changesets.BuildCommitteeVerifierParams(topology, 1)
	require.NoError(t, err)
	assert.Empty(t, params)
}

func TestBuildCommitteeVerifierParams_AddressValidation(t *testing.T) {
	sel := chainsel.TEST_90000001.Selector
	selStr := strconv.FormatUint(sel, 10)

	tests := []struct {
		name           string
		feeAggregator  string
		allowlistAdmin string
		errContains    string
	}{
		{
			name:          "empty FeeAggregator is rejected",
			feeAggregator: "",
			errContains:   "FeeAggregator is required",
		},
		{
			name:          "invalid hex FeeAggregator is rejected",
			feeAggregator: "not-hex",
			errContains:   "not a valid hex address",
		},
		{
			name:          "zero address FeeAggregator is rejected",
			feeAggregator: "0x0000000000000000000000000000000000000000",
			errContains:   "cannot be zero address",
		},
		{
			name:           "invalid hex AllowlistAdmin is rejected",
			feeAggregator:  "0x0000000000000000000000000000000000000001",
			allowlistAdmin: "not-hex",
			errContains:    "AllowlistAdmin",
		},
		{
			name:           "valid FeeAggregator and empty AllowlistAdmin succeeds",
			feeAggregator:  "0x0000000000000000000000000000000000000001",
			allowlistAdmin: "",
		},
		{
			name:           "valid FeeAggregator and valid AllowlistAdmin succeeds",
			feeAggregator:  "0x0000000000000000000000000000000000000001",
			allowlistAdmin: "0x0000000000000000000000000000000000000002",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			topology := &topology.EnvironmentTopology{
				NOPTopology: &topology.NOPTopology{
					Committees: map[string]topology.CommitteeConfig{
						"test": {
							Qualifier:       "test",
							VerifierVersion: semver.MustParse("1.7.0"),
							ChainConfigs: map[string]topology.ChainCommitteeConfig{
								selStr: {
									NOPAliases:     []string{"nop-1"},
									Threshold:      1,
									FeeAggregator:  tt.feeAggregator,
									AllowlistAdmin: tt.allowlistAdmin,
								},
							},
						},
					},
				},
			}
			params, err := changesets.BuildCommitteeVerifierParams(topology, sel)
			if tt.errContains != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
			} else {
				require.NoError(t, err)
				assert.Len(t, params, 1)
			}
		})
	}
}
