package executor_config_test

import (
	"strconv"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	execcontract "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	offrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/executor_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	testChainSelector   = uint64(16015286601757825753)
	executorQualifier   = "default"
	offRamp1_6_0Address = "0xCB91b03C1F0C802581B6B65411643cE1283121A4"
	offRamp1_7_0Address = "0xA7D86119ADB5F20Eb08AD4c486904945241f7438"
	rmn1_6_0Address     = "0x1f1d4D673d345A39904E4444810bE9cdF8F15dB6"
	rmn1_7_0Address     = "0x9999999999999999999999999999999999999999"
	executor1_6_0Addr   = "0x5555555555555555555555555555555555555555"
	executor1_7_0Addr   = "0x0526a85B3b7c78cBF53614CA575d85e978c3C21A"
)

var otherVersion = semver.MustParse("10.0.0")

func TestBuildConfig_ResolvesOffRamp1_7_0WhenDatastoreHasBoth1_6_0And1_7_0(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSelector

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(offrampoperations.ContractType),
		Qualifier:     "",
		Address:       offRamp1_6_0Address,
		Version:       otherVersion,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(offrampoperations.ContractType),
		Qualifier:     "",
		Address:       offRamp1_7_0Address,
		Version:       offrampoperations.Version,
	}))

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Address,
		Version:       rmn_remote.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(execcontract.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       execcontract.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, executor_config.BuildConfig, executor_config.BuildConfigDeps{Env: env}, executor_config.BuildConfigInput{
		ExecutorQualifier: executorQualifier,
		ChainSelectors:    []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	chainCfg, ok := report.Output.Config.ChainConfigs[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, offRamp1_7_0Address, chainCfg.OffRampAddress, "BuildConfig must resolve 1.7.0 OffRamp when both 1.6.0 and 1.7.0 exist")
}

func TestBuildConfig_ResolvesRMNRemote1_6_0WhenDatastoreHasBoth1_6_0And1_7_0(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSelector

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(offrampoperations.ContractType),
		Qualifier:     "",
		Address:       offRamp1_7_0Address,
		Version:       offrampoperations.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_7_0Address,
		Version:       otherVersion,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Address,
		Version:       rmn_remote.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(execcontract.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       execcontract.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, executor_config.BuildConfig, executor_config.BuildConfigDeps{Env: env}, executor_config.BuildConfigInput{
		ExecutorQualifier: executorQualifier,
		ChainSelectors:    []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	chainCfg, ok := report.Output.Config.ChainConfigs[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, rmn1_6_0Address, chainCfg.RmnAddress, "BuildConfig must resolve RMNRemote matching rmn_remote.Version (1.6.0) when multiple versions exist")
}

func TestBuildConfig_ResolvesExecutorProxy1_7_0WhenDatastoreHasMultipleVersions(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSelector

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(offrampoperations.ContractType),
		Qualifier:     "",
		Address:       offRamp1_7_0Address,
		Version:       offrampoperations.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Address,
		Version:       rmn_remote.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(execcontract.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_6_0Addr,
		Version:       otherVersion,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(execcontract.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       execcontract.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, executor_config.BuildConfig, executor_config.BuildConfigDeps{Env: env}, executor_config.BuildConfigInput{
		ExecutorQualifier: executorQualifier,
		ChainSelectors:    []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	chainCfg, ok := report.Output.Config.ChainConfigs[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, executor1_7_0Addr, chainCfg.ExecutorProxyAddress, "BuildConfig must resolve ExecutorProxy 1.7.0 when multiple versions exist for same qualifier")
}

func TestBuildConfig_SucceedsWithSingleVersionPerType(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSelector

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(offrampoperations.ContractType),
		Qualifier:     "",
		Address:       offRamp1_7_0Address,
		Version:       offrampoperations.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Address,
		Version:       rmn_remote.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(execcontract.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       execcontract.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, executor_config.BuildConfig, executor_config.BuildConfigDeps{Env: env}, executor_config.BuildConfigInput{
		ExecutorQualifier: executorQualifier,
		ChainSelectors:    []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	chainCfg, ok := report.Output.Config.ChainConfigs[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, offRamp1_7_0Address, chainCfg.OffRampAddress)
	assert.Equal(t, rmn1_6_0Address, chainCfg.RmnAddress)
	assert.Equal(t, executor1_7_0Addr, chainCfg.ExecutorProxyAddress)
}
