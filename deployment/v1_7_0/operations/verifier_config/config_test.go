package verifier_config_test

import (
	"strconv"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	onrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/verifier_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

const (
	testChainSel       = uint64(16015286601757825753)
	committeeQualifier = "default"
	executorQualifier  = "default"
	verifier1_7_0Addr  = "0x958F44bbA928E294D5199870e330c4f30E5E5Ed4"
	verifier1_6_0Addr  = "0x1111111111111111111111111111111111111111"
	onRamp1_7_0Addr    = "0x7eaAA8FF09bc1f2e53AdC09970E63660e9cc1f36"
	onRamp1_6_0Addr    = "0xDe5aC0AB395e34c383D9cBa8e08D8F3d66Ec5808"
	executor1_7_0Addr  = "0x0526a85B3b7c78cBF53614CA575d85e978c3C21A"
	executor1_6_0Addr  = "0x5555555555555555555555555555555555555555"
	rmn1_6_0Addr       = "0x1f1d4D673d345A39904E4444810bE9cdF8F15dB6"
	rmn1_7_0Addr       = "0x9999999999999999999999999999999999999999"
)

var otherVersion = semver.MustParse("10.0.0")

func TestBuildConfig_ResolvesCommitteeVerifierWhenSingleRefExists(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSel

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Qualifier:     committeeQualifier,
		Address:       verifier1_7_0Addr,
		Version:       committee_verifier.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(committee_verifier.ContractType),
		Qualifier:     committeeQualifier,
		Address:       verifier1_7_0Addr,
		Version:       committee_verifier.Version,
	}))

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(onrampoperations.ContractType),
		Qualifier:     "",
		Address:       onRamp1_7_0Addr,
		Version:       onrampoperations.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(executor.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       executor.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Addr,
		Version:       rmn_remote.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, verifier_config.BuildConfig, verifier_config.BuildConfigDeps{Env: env}, verifier_config.BuildConfigInput{
		CommitteeQualifier: committeeQualifier,
		ExecutorQualifier:  executorQualifier,
		ChainSelectors:     []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	addr, ok := report.Output.Config.CommitteeVerifierAddresses[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, verifier1_7_0Addr, addr)
}

func TestBuildConfig_ResolvesOnRamp1_7_0WhenDatastoreHasBothVersions(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSel

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Qualifier:     committeeQualifier,
		Address:       verifier1_7_0Addr,
		Version:       committee_verifier.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(onrampoperations.ContractType),
		Qualifier:     "",
		Address:       onRamp1_6_0Addr,
		Version:       otherVersion,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(onrampoperations.ContractType),
		Qualifier:     "",
		Address:       onRamp1_7_0Addr,
		Version:       onrampoperations.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(executor.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       executor.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Addr,
		Version:       rmn_remote.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, verifier_config.BuildConfig, verifier_config.BuildConfigDeps{Env: env}, verifier_config.BuildConfigInput{
		CommitteeQualifier: committeeQualifier,
		ExecutorQualifier:  executorQualifier,
		ChainSelectors:     []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	addr, ok := report.Output.Config.OnRampAddresses[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, onRamp1_7_0Addr, addr, "BuildConfig must resolve OnRamp 1.7.0 when both versions exist")
}

func TestBuildConfig_ResolvesExecutor1_7_0WhenDatastoreHasBothVersions(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSel

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Qualifier:     committeeQualifier,
		Address:       verifier1_7_0Addr,
		Version:       committee_verifier.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(onrampoperations.ContractType),
		Qualifier:     "",
		Address:       onRamp1_7_0Addr,
		Version:       onrampoperations.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(executor.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_6_0Addr,
		Version:       otherVersion,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(executor.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       executor.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Addr,
		Version:       rmn_remote.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, verifier_config.BuildConfig, verifier_config.BuildConfigDeps{Env: env}, verifier_config.BuildConfigInput{
		CommitteeQualifier: committeeQualifier,
		ExecutorQualifier:  executorQualifier,
		ChainSelectors:     []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	addr, ok := report.Output.Config.DefaultExecutorOnRampAddresses[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, executor1_7_0Addr, addr, "BuildConfig must resolve Executor 1.7.0 when both versions exist")
}

func TestBuildConfig_ResolvesRMNRemote1_6_0WhenDatastoreHasBothVersions(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	chainSel := testChainSel

	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(committee_verifier.ResolverType),
		Qualifier:     committeeQualifier,
		Address:       verifier1_7_0Addr,
		Version:       committee_verifier.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(onrampoperations.ContractType),
		Qualifier:     "",
		Address:       onRamp1_7_0Addr,
		Version:       onrampoperations.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(executor.ProxyType),
		Qualifier:     executorQualifier,
		Address:       executor1_7_0Addr,
		Version:       executor.Version,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_7_0Addr,
		Version:       otherVersion,
	}))
	require.NoError(t, ds.Addresses().Add(datastore.AddressRef{
		ChainSelector: chainSel,
		Type:          datastore.ContractType(rmn_remote.ContractType),
		Qualifier:     "",
		Address:       rmn1_6_0Addr,
		Version:       rmn_remote.Version,
	}))

	env := deployment.Environment{
		OperationsBundle: testutils.NewTestBundle(),
		DataStore:        ds.Seal(),
	}
	report, err := operations.ExecuteOperation(env.OperationsBundle, verifier_config.BuildConfig, verifier_config.BuildConfigDeps{Env: env}, verifier_config.BuildConfigInput{
		CommitteeQualifier: committeeQualifier,
		ExecutorQualifier:  executorQualifier,
		ChainSelectors:     []uint64{chainSel},
	})
	require.NoError(t, err)
	require.NotNil(t, report.Output.Config)

	addr, ok := report.Output.Config.RMNRemoteAddresses[strconv.FormatUint(chainSel, 10)]
	require.True(t, ok)
	assert.Equal(t, rmn1_6_0Addr, addr, "BuildConfig must resolve RMNRemote matching rmn_remote.Version (1.6.0) when both versions exist")
}
