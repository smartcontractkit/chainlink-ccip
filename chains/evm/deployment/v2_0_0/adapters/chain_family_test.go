package adapters_test

import (
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/engine/test/environment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	evmadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/adapters"
	evmchangesets "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/create2_factory"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/testsetup"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	v2changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
)

const committeeQualifierDefault = "default"

// evmFamilySelector is bytes4(keccak256("CCIP ChainFamilySelector EVM")) = 0x2812d52c.
var evmFamilySelector = [4]byte{0x28, 0x12, 0xd5, 0x2c}

type laneDeployedContracts struct {
	router                    string
	onRamp                    string
	feeQuoter                 string
	offRamp                   string
	committeeVerifier         string
	committeeVerifierResolver string
	executor                  string
}

// contractParamsForLaneTopologyTest uses committee qualifier "default" to match
// ConfigureChainsForLanesFromTopology lane expansion.
func contractParamsForLaneTopologyTest() sequences.ContractParams {
	params := testsetup.CreateBasicContractParams()
	if len(params.CommitteeVerifiers) > 0 {
		params.CommitteeVerifiers[0].Qualifier = committeeQualifierDefault
	}
	if len(params.MockReceivers) > 0 && len(params.MockReceivers[0].RequiredVerifiers) > 0 {
		params.MockReceivers[0].RequiredVerifiers[0].Qualifier = committeeQualifierDefault
	}
	return params
}

func requireExecutorProxyTargetInitialized(
	t *testing.T,
	b operations.Bundle,
	evmChain cldfevm.Chain,
	executorProxyAddr string,
) {
	t.Helper()
	targetReport, err := operations.ExecuteOperation(b, proxy.GetTarget, evmChain, contract_utils.FunctionInput[struct{}]{
		ChainSelector: evmChain.Selector,
		Address:       common.HexToAddress(executorProxyAddr),
	})
	require.NoError(t, err)
	require.NotEqual(t, common.Address{1}, targetReport.Output, "executor proxy target must be initialized to the implementation, not the deploy placeholder 0x01")
}

func deployLaneContractsToDatastore(
	t *testing.T,
	e *deployment.Environment,
	chainSelector uint64,
	ds datastore.MutableDataStore,
) laneDeployedContracts {
	t.Helper()
	evmChain := e.BlockChains.EVMChains()[chainSelector]

	create2Ref, err := contract_utils.MaybeDeployContract(
		e.OperationsBundle, create2_factory.Deploy, evmChain,
		contract_utils.DeployInput[create2_factory.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
			ChainSelector:  chainSelector,
			Args:           create2_factory.ConstructorArgs{AllowList: []common.Address{evmChain.DeployerKey.From}},
		}, nil,
	)
	require.NoError(t, err)

	deployOut, err := evmchangesets.DeployChainContracts(changesets.GetRegistry()).Apply(*e, changesets.WithMCMS[evmchangesets.DeployChainContractsCfg]{
		Cfg: evmchangesets.DeployChainContractsCfg{
			ChainSel:         chainSelector,
			CREATE2Factory:   common.HexToAddress(create2Ref.Address),
			Params:           contractParamsForLaneTopologyTest(),
			DeployerKeyOwned: true,
		},
		MCMS: mcms.Input{},
	})
	require.NoError(t, err)
	require.NoError(t, ds.Merge(deployOut.DataStore.Seal()))

	var out laneDeployedContracts
	for _, addr := range deployOut.DataStore.Addresses().Filter(datastore.AddressRefByChainSelector(chainSelector)) {
		switch addr.Type {
		case datastore.ContractType(router.ContractType):
			out.router = addr.Address
		case datastore.ContractType(onramp.ContractType):
			out.onRamp = addr.Address
		case datastore.ContractType(fee_quoter.ContractType):
			out.feeQuoter = addr.Address
		case datastore.ContractType(offramp.ContractType):
			out.offRamp = addr.Address
		case datastore.ContractType(committee_verifier.ContractType):
			if addr.Qualifier == committeeQualifierDefault {
				out.committeeVerifier = addr.Address
			}
		case datastore.ContractType(sequences.ExecutorProxyType):
			if addr.Qualifier == committeeQualifierDefault && out.executor == "" {
				out.executor = addr.Address
			}
		case datastore.ContractType(sequences.CommitteeVerifierResolverType):
			if addr.Qualifier == committeeQualifierDefault {
				out.committeeVerifierResolver = addr.Address
			}
		}
	}
	require.NotEmpty(t, out.committeeVerifier, "committee verifier with qualifier %q must be deployed", committeeQualifierDefault)
	require.NotEmpty(t, out.committeeVerifierResolver, "committee verifier resolver with qualifier %q must be deployed", committeeQualifierDefault)
	require.NotEmpty(t, out.executor, "executor with qualifier %q must be deployed", committeeQualifierDefault)

	adapter := &evmadapters.ChainFamilyAdapter{}
	out.executor, err = adapter.ResolveExecutor(ds.Seal(), chainSelector, committeeQualifierDefault)
	require.NoError(t, err)

	requireExecutorProxyTargetInitialized(t, e.OperationsBundle, evmChain, out.executor)
	return out
}

func bidirectionalLaneTopology(signer string, chainSelectors ...uint64) *offchain.EnvironmentTopology {
	chainConfigs := make(map[string]offchain.ChainCommitteeConfig, len(chainSelectors))
	for _, sel := range chainSelectors {
		chainConfigs[fmt.Sprintf("%d", sel)] = offchain.ChainCommitteeConfig{
			NOPAliases: []string{"nop-1"},
			Threshold:  1,
		}
	}
	return &offchain.EnvironmentTopology{
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop-1", SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: signer}},
			},
			Committees: map[string]offchain.CommitteeConfig{
				committeeQualifierDefault: {
					Qualifier:    committeeQualifierDefault,
					ChainConfigs: chainConfigs,
				},
			},
		},
	}
}

func assertAdapterDefaultsOnChain(
	t *testing.T,
	b operations.Bundle,
	adapter *evmadapters.ChainFamilyAdapter,
	evmChain cldfevm.Chain,
	local laneDeployedContracts,
	remote laneDeployedContracts,
	localSelector, remoteSelector uint64,
) {
	t.Helper()

	familySel := adapter.GetChainFamilySelector()
	fqDefaults := adapter.GetDefaultFeeQuoterDestChainConfig(localSelector, remoteSelector, familySel)
	remoteDefaults := adapter.GetDefaultRemoteChainConfig(localSelector, remoteSelector)
	cvDefaults := adapter.GetDefaultCommitteeVerifierRemoteChainConfig()
	finalityDefaults := adapter.GetDefaultFinalityConfig()

	srcCfg, err := operations.ExecuteOperation(b, offramp.GetSourceChainConfig, evmChain, contract_utils.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.offRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, remoteDefaults.AllowTrafficFrom, srcCfg.Output.IsEnabled)
	assert.Equal(t, local.router, srcCfg.Output.Router.Hex())
	assert.Len(t, srcCfg.Output.OnRamps, 1)
	assert.Equal(t, common.LeftPadBytes(common.HexToAddress(remote.onRamp).Bytes(), 32), srcCfg.Output.OnRamps[0])
	assert.Len(t, srcCfg.Output.DefaultCCVs, 1)
	// Empty lane config auto-resolves default inbound/outbound CCVs to the committee verifier resolver.
	assert.Equal(t, local.committeeVerifierResolver, srcCfg.Output.DefaultCCVs[0].Hex())

	destCfg, err := operations.ExecuteOperation(b, onramp.GetDestChainConfig, evmChain, contract_utils.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.onRamp), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, local.router, destCfg.Output.Router.Hex())
	assert.Equal(t, common.HexToAddress(remote.offRamp).Bytes(), destCfg.Output.OffRamp)
	assert.Equal(t, local.executor, destCfg.Output.DefaultExecutor.Hex())
	assert.Len(t, destCfg.Output.DefaultCCVs, 1)
	assert.Equal(t, local.committeeVerifierResolver, destCfg.Output.DefaultCCVs[0].Hex())
	assert.Equal(t, adapter.GetAddressBytesLength(), destCfg.Output.AddressBytesLength)
	assert.Equal(t, remoteDefaults.BaseExecutionGasCost, destCfg.Output.BaseExecutionGasCost)
	assert.Equal(t, remoteDefaults.MessageNetworkFeeUSDCents, destCfg.Output.MessageNetworkFeeUSDCents)
	assert.Equal(t, remoteDefaults.TokenNetworkFeeUSDCents, destCfg.Output.TokenNetworkFeeUSDCents)
	assert.Equal(t, remoteDefaults.TokenReceiverAllowed, destCfg.Output.TokenReceiverAllowed)

	executorDestChains, err := operations.ExecuteOperation(b, executor.GetDestChains, evmChain, contract_utils.FunctionInput[struct{}]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.executor),
	})
	require.NoError(t, err)
	require.Len(t, executorDestChains.Output, 1)
	assert.Equal(t, remoteSelector, executorDestChains.Output[0].DestChainSelector)
	assert.Equal(t, remoteDefaults.ExecutorDestChainConfig.USDCentsFee, executorDestChains.Output[0].Config.UsdCentsFee)
	assert.Equal(t, remoteDefaults.ExecutorDestChainConfig.Enabled, executorDestChains.Output[0].Config.Enabled)

	fqDestCfg, err := operations.ExecuteOperation(b, fee_quoter.GetDestChainConfig, evmChain, contract_utils.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.feeQuoter), Args: remoteSelector,
	})
	require.NoError(t, err)
	require.NotNil(t, fqDefaults.IsEnabled)
	require.NotNil(t, fqDefaults.MaxDataBytes)
	require.NotNil(t, fqDefaults.MaxPerMsgGasLimit)
	require.NotNil(t, fqDefaults.DestGasOverhead)
	require.NotNil(t, fqDefaults.DestGasPerPayloadByteBase)
	require.NotNil(t, fqDefaults.DefaultTokenFeeUSDCents)
	require.NotNil(t, fqDefaults.DefaultTokenDestGasOverhead)
	require.NotNil(t, fqDefaults.DefaultTxGasLimit)
	require.NotNil(t, fqDefaults.NetworkFeeUSDCents)
	require.NotNil(t, fqDefaults.LinkFeeMultiplierPercent)
	assert.Equal(t, *fqDefaults.IsEnabled, fqDestCfg.Output.IsEnabled)
	assert.Equal(t, *fqDefaults.MaxDataBytes, fqDestCfg.Output.MaxDataBytes)
	assert.Equal(t, *fqDefaults.MaxPerMsgGasLimit, fqDestCfg.Output.MaxPerMsgGasLimit)
	assert.Equal(t, *fqDefaults.DestGasOverhead, fqDestCfg.Output.DestGasOverhead)
	assert.Equal(t, *fqDefaults.DestGasPerPayloadByteBase, fqDestCfg.Output.DestGasPerPayloadByteBase)
	assert.Equal(t, fqDefaults.ChainFamilySelector, fqDestCfg.Output.ChainFamilySelector)
	assert.Equal(t, *fqDefaults.DefaultTokenFeeUSDCents, fqDestCfg.Output.DefaultTokenFeeUSDCents)
	assert.Equal(t, *fqDefaults.DefaultTokenDestGasOverhead, fqDestCfg.Output.DefaultTokenDestGasOverhead)
	assert.Equal(t, *fqDefaults.DefaultTxGasLimit, fqDestCfg.Output.DefaultTxGasLimit)
	assert.Equal(t, *fqDefaults.NetworkFeeUSDCents, fqDestCfg.Output.NetworkFeeUSDCents)
	assert.Equal(t, *fqDefaults.LinkFeeMultiplierPercent, fqDestCfg.Output.LinkFeeMultiplierPercent)

	verifierRemoteCfg, err := operations.ExecuteOperation(b, committee_verifier.GetRemoteChainConfig, evmChain, contract_utils.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier), Args: remoteSelector,
	})
	require.NoError(t, err)
	assert.Equal(t, local.router, verifierRemoteCfg.Output.RemoteChainConfig.Router.Hex())
	assert.Equal(t, cvDefaults.AllowlistEnabled, verifierRemoteCfg.Output.RemoteChainConfig.AllowlistEnabled)

	verifierFee, err := operations.ExecuteOperation(b, committee_verifier.GetFee, evmChain, contract_utils.FunctionInput[committee_verifier.GetFeeArgs]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier),
		Args: committee_verifier.GetFeeArgs{DestChainSelector: remoteSelector},
	})
	require.NoError(t, err)
	assert.Equal(t, cvDefaults.FeeUSDCents, verifierFee.Output.FeeUSDCents)
	assert.Equal(t, cvDefaults.GasForVerification, verifierFee.Output.GasForVerification)
	assert.Equal(t, uint32(cvDefaults.PayloadSizeBytes), verifierFee.Output.PayloadSizeBytes)

	finalityCfg, err := operations.ExecuteOperation(b, committee_verifier.GetAllowedFinalityConfig, evmChain, contract_utils.FunctionInput[struct{}]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier),
	})
	require.NoError(t, err)
	assert.Equal(t, finalityDefaults.Raw(), finalityCfg.Output)

	sigCfg, err := operations.ExecuteOperation(b, committee_verifier.GetSignatureConfig, evmChain, contract_utils.FunctionInput[uint64]{
		ChainSelector: evmChain.Selector, Address: common.HexToAddress(local.committeeVerifier), Args: remoteSelector,
	})
	require.NoError(t, err)
	require.Len(t, sigCfg.Output.Signers, 1)
	assert.Equal(t, evmChain.DeployerKey.From, sigCfg.Output.Signers[0])
	assert.Equal(t, uint8(1), sigCfg.Output.Threshold)
}

func makeChainConfig(chainSelector uint64, remoteChainSelector uint64) lanes.ChainDefinition {
	return lanes.ChainDefinition{
		Selector: chainSelector,
		CommitteeVerifiers: []lanes.CommitteeVerifierConfig[datastore.AddressRef]{
			{
				CommitteeVerifier: []datastore.AddressRef{
					{
						ChainSelector: chainSelector,
						Type:          datastore.ContractType(committee_verifier.ContractType),
						Version:       committee_verifier.Version,
					},
					{
						ChainSelector: chainSelector,
						Type:          datastore.ContractType(sequences.CommitteeVerifierResolverType),
						Version:       semver.MustParse("2.0.0"),
					},
				},
				RemoteChains: map[uint64]lanes.CommitteeVerifierRemoteChainConfig{
					remoteChainSelector: testsetup.CreateBasicCommitteeVerifierRemoteChainConfig(),
				},
			},
		},
		DefaultInboundCCVs: []datastore.AddressRef{
			{
				ChainSelector: chainSelector,
				Type:          datastore.ContractType(committee_verifier.ContractType),
				Version:       committee_verifier.Version,
			},
		},
		DefaultOutboundCCVs: []datastore.AddressRef{
			{
				ChainSelector: chainSelector,
				Type:          datastore.ContractType(committee_verifier.ContractType),
				Version:       committee_verifier.Version,
			},
		},
		DefaultExecutor: datastore.AddressRef{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(sequences.ExecutorProxyType),
			Version:       executor.Version,
			Qualifier:     "default",
		},
		FeeQuoterDestChainConfig: testsetup.CreateBasicFeeQuoterDestChainConfig(),
		ExecutorDestChainConfig:  testsetup.CreateBasicExecutorDestChainConfig(),
		AddressBytesLength:       20,
		BaseExecutionGasCost:     80_000,
	}
}

func TestChainFamilyAdapter(t *testing.T) {
	tests := []struct {
		desc string
	}{
		{
			desc: "happy path",
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			chainA := uint64(5009297550715157269)
			chainB := uint64(4949039107694359620)
			e, err := environment.New(t.Context(),
				environment.WithEVMSimulated(t, []uint64{chainA, chainB}),
			)
			require.NoError(t, err, "Failed to create test environment")
			require.NotNil(t, e, "Environment should be created")

			chainFamilyRegistry := lanes.GetLaneAdapterRegistry()
			mcmsRegistry := changesets.GetRegistry()

			// On each chain, deploy chain contracts
			ds := datastore.NewMemoryDataStore()
			for _, chainSel := range []uint64{chainA, chainB} {
				create2FactoryRef, err := contract_utils.MaybeDeployContract(e.OperationsBundle, create2_factory.Deploy, e.BlockChains.EVMChains()[chainSel], contract_utils.DeployInput[create2_factory.ConstructorArgs]{
					TypeAndVersion: deployment.NewTypeAndVersion(create2_factory.ContractType, *semver.MustParse("2.0.0")),
					ChainSelector:  chainSel,
					Args: create2_factory.ConstructorArgs{
						AllowList: []common.Address{e.BlockChains.EVMChains()[chainSel].DeployerKey.From},
					},
				}, nil)
				require.NoError(t, err, "Failed to deploy CREATE2Factory")

				deployChainOut, err := evmchangesets.DeployChainContracts(mcmsRegistry).Apply(*e, changesets.WithMCMS[evmchangesets.DeployChainContractsCfg]{
					Cfg: evmchangesets.DeployChainContractsCfg{
						ChainSel:         chainSel,
						CREATE2Factory:   common.HexToAddress(create2FactoryRef.Address),
						Params:           testsetup.CreateBasicContractParams(),
						DeployerKeyOwned: true,
					},
				})
				require.NoError(t, err, "Failed to apply DeployChainContracts changeset")
				err = ds.Merge(deployChainOut.DataStore.Seal())
				require.NoError(t, err, "Failed to merge datastore from DeployChainContracts changeset")
			}

			// Overwrite datastore in the environment
			e.DataStore = ds.Seal()

			// Configure chains for lanes
			e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
			_, err = lanes.ConnectChains(chainFamilyRegistry, mcmsRegistry).Apply(*e, lanes.ConnectChainsConfig{
				Lanes: []lanes.LaneConfig{
					{
						ChainA:  makeChainConfig(chainA, chainB),
						ChainB:  makeChainConfig(chainB, chainA),
						Version: semver.MustParse("2.0.0"),
					},
				},
			})
			require.NoError(t, err, "Failed to apply ConnectChains changeset")
		})
	}
}

func TestChainFamilyAdapter_DefaultsAppliedOnChainViaConfigureChainsForLanesFromTopology(t *testing.T) {
	localSelector := chainsel.TEST_90000001.Selector
	remoteSelector := chainsel.TEST_90000002.Selector

	e, err := environment.New(t.Context(),
		environment.WithEVMSimulated(t, []uint64{localSelector, remoteSelector}),
	)
	require.NoError(t, err)

	ds := datastore.NewMemoryDataStore()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	local := deployLaneContractsToDatastore(t, e, localSelector, ds)
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	remote := deployLaneContractsToDatastore(t, e, remoteSelector, ds)
	e.DataStore = ds.Seal()

	deployer := e.BlockChains.EVMChains()[localSelector].DeployerKey.From.Hex()
	e.OperationsBundle = testsetup.BundleWithFreshReporter(e.OperationsBundle)
	cs := v2changesets.ConfigureChainsForLanesFromTopology(
		ccvadapters.GetCommitteeVerifierContractRegistry(),
		ccvadapters.GetChainFamilyRegistry(),
		changesets.GetRegistry(),
	)
	_, err = cs.Apply(*e, v2changesets.ConfigureChainsForLanesFromTopologyConfig{
		Topology: bidirectionalLaneTopology(deployer, localSelector, remoteSelector),
		BuildLanesCrossFamilyConfig: v2changesets.BuildLanesCrossFamilyConfig{
			Lanes: []v2changesets.CrossFamilyLanePair{
				{ChainA: localSelector, ChainB: remoteSelector},
			},
			MCMS: mcms.Input{},
		},
	})
	require.NoError(t, err)

	adapter := &evmadapters.ChainFamilyAdapter{}
	evmChain := e.BlockChains.EVMChains()[localSelector]
	assertAdapterDefaultsOnChain(
		t,
		testsetup.BundleWithFreshReporter(e.OperationsBundle),
		adapter,
		evmChain,
		local,
		remote,
		localSelector,
		remoteSelector,
	)
}

func TestChainFamilyAdapter_ValidateNOPsTopology(t *testing.T) {
	adapter := &evmadapters.ChainFamilyAdapter{}
	const chainSelector = "5009297550715157269"

	tests := []struct {
		name     string
		nopCount int
		wantErr  string
	}{
		{
			name:     "fewer than the minimum is rejected",
			nopCount: 8,
			wantErr:  `chain "5009297550715157269" requires at least 9 unique NOPs for production environments, got 8`,
		},
		{
			name:     "exactly the minimum is allowed",
			nopCount: 9,
		},
		{
			name:     "more than the minimum is allowed",
			nopCount: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := adapter.ValidateNOPsTopology(chainSelector, tt.nopCount)
			if tt.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Equal(t, tt.wantErr, err.Error())
		})
	}
}
