package changesets_test

import (
	"strconv"
	"testing"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	execcontract "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/executor"
	offrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/offramp"
	onrampoperations "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/rmn_remote"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
)

const (
	testCommittee         = "test-committee"
	testDefaultQualifier  = "default"
	testAggregatorName    = "instance-1"
	testAggregatorAddress = "aggregator-1:443"
	testIndexerAddress    = "http://indexer:8100"
)

var (
	defaultSelectors = []uint64{
		chainsel.TEST_90000001.Selector,
		chainsel.TEST_90000002.Selector,
	}

	testContractAddresses = struct {
		CommitteeVerifier1 common.Address
		CommitteeVerifier2 common.Address
		OnRamp1            common.Address
		OnRamp2            common.Address
		Executor1          common.Address
		Executor2          common.Address
		OffRamp1           common.Address
		OffRamp2           common.Address
		RMN1               common.Address
		RMN2               common.Address
	}{
		CommitteeVerifier1: common.HexToAddress("0x1111111111111111111111111111111111111111"),
		CommitteeVerifier2: common.HexToAddress("0x2222222222222222222222222222222222222222"),
		OnRamp1:            common.HexToAddress("0x3333333333333333333333333333333333333333"),
		OnRamp2:            common.HexToAddress("0x4444444444444444444444444444444444444444"),
		Executor1:          common.HexToAddress("0x5555555555555555555555555555555555555555"),
		Executor2:          common.HexToAddress("0x6666666666666666666666666666666666666666"),
		OffRamp1:           common.HexToAddress("0x7777777777777777777777777777777777777777"),
		OffRamp2:           common.HexToAddress("0x8888888888888888888888888888888888888888"),
		RMN1:               common.HexToAddress("0x9999999999999999999999999999999999999999"),
		RMN2:               common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
	}
)

type TopologyOption func(*topology.EnvironmentTopology)

func WithPyroscopeURL(url string) TopologyOption {
	return func(t *topology.EnvironmentTopology) {
		t.PyroscopeURL = url
	}
}

func WithMonitoring(cfg topology.MonitoringConfig) TopologyOption {
	return func(t *topology.EnvironmentTopology) {
		t.Monitoring = cfg
	}
}

func WithNOPs(nops []topology.NOPConfig) TopologyOption {
	return func(t *topology.EnvironmentTopology) {
		t.NOPTopology.NOPs = nops
	}
}

func WithCommittee(name string, cfg topology.CommitteeConfig) TopologyOption {
	return func(t *topology.EnvironmentTopology) {
		if t.NOPTopology.Committees == nil {
			t.NOPTopology.Committees = make(map[string]topology.CommitteeConfig)
		}
		t.NOPTopology.Committees[name] = cfg
	}
}

func WithExecutorPool(name string, cfg topology.ExecutorPoolConfig) TopologyOption {
	return func(t *topology.EnvironmentTopology) {
		if t.ExecutorPools == nil {
			t.ExecutorPools = make(map[string]topology.ExecutorPoolConfig)
		}
		t.ExecutorPools[name] = cfg
	}
}

func WithIndexerAddress(address []string) TopologyOption {
	return func(t *topology.EnvironmentTopology) {
		t.IndexerAddress = address
	}
}

func newTestTopology(opts ...TopologyOption) *topology.EnvironmentTopology {
	sel1Str := strconv.FormatUint(chainsel.TEST_90000001.Selector, 10)
	sel2Str := strconv.FormatUint(chainsel.TEST_90000002.Selector, 10)

	topology := &topology.EnvironmentTopology{
		IndexerAddress: []string{testIndexerAddress},
		PyroscopeURL:   "",
		NOPTopology: &topology.NOPTopology{
			NOPs: []topology.NOPConfig{
				{
					Alias:                 "nop-1",
					Name:                  "nop-1",
					SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xABCDEF1234567890ABCDEF1234567890ABCDEF12"},
					Mode:                  shared.NOPModeStandalone,
				},
				{
					Alias:                 "nop-2",
					Name:                  "nop-2",
					SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0x1234567890ABCDEF1234567890ABCDEF12345678"},
					Mode:                  shared.NOPModeStandalone,
				},
			},
			Committees: map[string]topology.CommitteeConfig{
				testCommittee: {
					Qualifier:       testCommittee,
					VerifierVersion: semver.MustParse("1.7.0"),
					Aggregators: []topology.AggregatorConfig{
						{Name: testAggregatorName, Address: testAggregatorAddress, InsecureAggregatorConnection: true},
					},
					ChainConfigs: map[string]topology.ChainCommitteeConfig{
						sel1Str: {
							NOPAliases:    []string{"nop-1", "nop-2"},
							Threshold:     2,
							FeeAggregator: "0x0000000000000000000000000000000000000001",
						},
						sel2Str: {
							NOPAliases:    []string{"nop-1", "nop-2"},
							Threshold:     2,
							FeeAggregator: "0x0000000000000000000000000000000000000001",
						},
					},
				},
			},
		},
		ExecutorPools: map[string]topology.ExecutorPoolConfig{
			testDefaultQualifier: {
				NOPAliases:        []string{"nop-1", "nop-2"},
				ExecutionInterval: 15 * time.Second,
			},
		},
	}

	for _, opt := range opts {
		opt(topology)
	}

	return topology
}

func newTestEnvironment(t *testing.T, selectors []uint64) deployment.Environment {
	t.Helper()
	env, _ := testutils.NewSimulatedEVMEnvironment(t, selectors)
	return env
}

func setupVerifierDatastore(t *testing.T, ds datastore.MutableDataStore, selectors []uint64, committeeQualifier, executorQualifier string) {
	t.Helper()
	addrs := testContractAddresses

	addContractToDatastoreWithVersion(t, ds, selectors[0], committeeQualifier, committee_verifier.ResolverType, committee_verifier.Version, addrs.CommitteeVerifier1)
	addContractToDatastoreWithVersion(t, ds, selectors[0], "", onrampoperations.ContractType, onrampoperations.Version, addrs.OnRamp1)
	addContractToDatastoreWithVersion(t, ds, selectors[0], executorQualifier, execcontract.ProxyType, execcontract.Version, addrs.Executor1)
	addContractToDatastoreWithVersion(t, ds, selectors[0], "", rmn_remote.ContractType, rmn_remote.Version, addrs.RMN1)

	if len(selectors) > 1 {
		addContractToDatastoreWithVersion(t, ds, selectors[1], committeeQualifier, committee_verifier.ResolverType, committee_verifier.Version, addrs.CommitteeVerifier2)
		addContractToDatastoreWithVersion(t, ds, selectors[1], "", onrampoperations.ContractType, onrampoperations.Version, addrs.OnRamp2)
		addContractToDatastoreWithVersion(t, ds, selectors[1], executorQualifier, execcontract.ProxyType, execcontract.Version, addrs.Executor2)
		addContractToDatastoreWithVersion(t, ds, selectors[1], "", rmn_remote.ContractType, rmn_remote.Version, addrs.RMN2)
	}
}

func setupExecutorDatastore(t *testing.T, ds datastore.MutableDataStore, selectors []uint64, executorQualifier string) {
	t.Helper()
	addrs := testContractAddresses

	addContractToDatastoreWithVersion(t, ds, selectors[0], executorQualifier, execcontract.ProxyType, execcontract.Version, addrs.Executor1)
	addContractToDatastoreWithVersion(t, ds, selectors[0], "", offrampoperations.ContractType, offrampoperations.Version, addrs.OffRamp1)
	addContractToDatastoreWithVersion(t, ds, selectors[0], "", rmn_remote.ContractType, rmn_remote.Version, addrs.RMN1)

	if len(selectors) > 1 {
		addContractToDatastoreWithVersion(t, ds, selectors[1], executorQualifier, execcontract.ProxyType, execcontract.Version, addrs.Executor2)
		addContractToDatastoreWithVersion(t, ds, selectors[1], "", offrampoperations.ContractType, offrampoperations.Version, addrs.OffRamp2)
		addContractToDatastoreWithVersion(t, ds, selectors[1], "", rmn_remote.ContractType, rmn_remote.Version, addrs.RMN2)
	}
}

func addContractToDatastoreWithVersion(t *testing.T, ds datastore.MutableDataStore, chainSelector uint64, qualifier string, contractType deployment.ContractType, version *semver.Version, addr common.Address) {
	t.Helper()
	ref := datastore.AddressRef{
		ChainSelector: chainSelector,
		Qualifier:     qualifier,
		Type:          datastore.ContractType(contractType),
		Address:       addr.Hex(),
	}
	if version != nil {
		ref.Version = version
	}
	err := ds.Addresses().Add(ref)
	require.NoError(t, err)
}

func defaultMonitoringConfig() topology.MonitoringConfig {
	return topology.MonitoringConfig{
		Enabled: true,
		Type:    "beholder",
		Beholder: topology.BeholderConfig{
			InsecureConnection:       true,
			OtelExporterHTTPEndpoint: "otel:4318",
			MetricReaderInterval:     5,
			TraceSampleRatio:         1.0,
			TraceBatchTimeout:        10,
		},
	}
}
