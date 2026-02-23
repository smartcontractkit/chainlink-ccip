package changesets_test

import (
	"context"
	"errors"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	v1_7_0_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
)

type cctpTest_MockReader struct{}

func (m *cctpTest_MockReader) GetMCMSRef(e deployment.Environment, selector uint64, input mcms.Input) (datastore.AddressRef, error) {
	return datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType("MCM"),
		Version:       semver.MustParse("1.0.0"),
	}, selector, datastore_utils.FullRef)
}

func (m *cctpTest_MockReader) GetChainMetadata(e deployment.Environment, selector uint64, input mcms.Input) (mcms_types.ChainMetadata, error) {
	mcmsRef, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: selector,
		Type:          datastore.ContractType("MCM"),
		Version:       semver.MustParse("1.0.0"),
	}, selector, datastore_utils.FullRef)
	if err != nil {
		return mcms_types.ChainMetadata{}, err
	}
	return mcms_types.ChainMetadata{
		StartingOpCount: 10,
		MCMAddress:      mcmsRef.Address,
	}, nil
}

func (m *cctpTest_MockReader) GetTimelockRef(e deployment.Environment, selector uint64, input mcms.Input) (datastore.AddressRef, error) {
	return datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
		ChainSelector: selector,
		Type:          "Timelock",
		Version:       semver.MustParse("1.0.0"),
	}, selector, datastore_utils.FullRef)
}

type cctpTest_MockCCTPChain struct {
	sequenceErrorMsg string
}

// DeployCCTPChain returns a sequence that accepts resolved adapter input (with string addresses)
func (m *cctpTest_MockCCTPChain) DeployCCTPChain() *cldf_ops.Sequence[adapters.DeployCCTPInput, sequences.OnChainOutput, adapters.DeployCCTPChainDeps] {
	return cldf_ops.NewSequence(
		"mock-deploy-cctp-chain-sequence",
		semver.MustParse("1.0.0"),
		"Mock sequence for testing CCTP deployment",
		func(bundle cldf_ops.Bundle, deps adapters.DeployCCTPChainDeps, input adapters.DeployCCTPInput) (sequences.OnChainOutput, error) {
			if m.sequenceErrorMsg != "" {
				return sequences.OnChainOutput{}, errors.New(m.sequenceErrorMsg)
			}
			// In a real implementation, the adapter would convert adapter input to sequence input
			// and execute the actual DeployCCTPChain sequence. For testing, we just return mock output.
			return sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{
					{
						ChainSelector: input.ChainSelector,
						Address:       "0x6666666666666666666666666666666666666666",
						Type:          datastore.ContractType("USDCTokenPoolProxy"),
						Version:       semver.MustParse("1.7.0"),
					},
				},
				BatchOps: []mcms_types.BatchOperation{},
			}, nil
		},
	)
}

// ConfigureCCTPChainForLanes returns a sequence that configures CCTP for lanes
func (m *cctpTest_MockCCTPChain) ConfigureCCTPChainForLanes() *cldf_ops.Sequence[adapters.ConfigureCCTPChainForLanesInput, sequences.OnChainOutput, adapters.ConfigureCCTPChainForLanesDeps] {
	return cldf_ops.NewSequence(
		"mock-configure-cctp-chain-for-lanes-sequence",
		semver.MustParse("1.0.0"),
		"Mock sequence for testing CCTP configuration",
		func(bundle cldf_ops.Bundle, deps adapters.ConfigureCCTPChainForLanesDeps, input adapters.ConfigureCCTPChainForLanesInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{},
				BatchOps:  []mcms_types.BatchOperation{},
			}, nil
		},
	)
}

// PoolAddress returns the address of the token pool on the remote chain in bytes
func (m *cctpTest_MockCCTPChain) PoolAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64, registeredPoolRef datastore.AddressRef) ([]byte, error) {
	return []byte("pool-address"), nil
}

// TokenAddress returns the address of the token on the remote chain in bytes
func (m *cctpTest_MockCCTPChain) TokenAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte("token-address"), nil
}

// CCTPV1AllowedCallerOnDest returns the address allowed to trigger message reception on the remote domain for CCTP V1.
func (m *cctpTest_MockCCTPChain) CCTPV1AllowedCallerOnDest(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte("allowed-caller-dest-v1"), nil
}

// CCTPV2AllowedCallerOnDest returns the address allowed to trigger message reception on the remote domain for CCTP V2.
func (m *cctpTest_MockCCTPChain) CCTPV2AllowedCallerOnDest(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte("allowed-caller-dest-v2"), nil
}

// AllowedCallerOnSource returns the address allowed to deposit tokens for burn on the remote chain
func (m *cctpTest_MockCCTPChain) AllowedCallerOnSource(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte("allowed-caller-source"), nil
}

// MintRecipientOnDest returns the address that will receive tokens on the remote domain
func (m *cctpTest_MockCCTPChain) MintRecipientOnDest(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error) {
	return []byte("mint-recipient"), nil
}

// USDCType returns the type of the USDC on the remote chain
func (m *cctpTest_MockCCTPChain) USDCType() adapters.USDCType {
	return adapters.Canonical
}

var cctpTest_BasicMCMSInput = mcms.Input{
	OverridePreviousRoot: true,
	ValidUntil:           3759765795,
	TimelockDelay:        mcms_types.MustParseDuration("1h"),
	TimelockAction:       mcms_types.TimelockActionSchedule,
}

func TestDeployCCTPChains_Apply(t *testing.T) {
	tests := []struct {
		desc                     string
		makeDataStore            func(t *testing.T) *datastore.MemoryDataStore
		cfg                      v1_7_0_changesets.DeployCCTPChainsConfig
		expectedSequenceErrorMsg string
		expectedError            string
	}{
		{
			desc: "success - basic CCTP deployment",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				chainSelector := uint64(5009297550715157269)
				// Add required addresses to datastore
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x1111111111111111111111111111111111111111",
					Type:          datastore.ContractType("Router"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x2222222222222222222222222222222222222222",
					Type:          datastore.ContractType("RMN"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x3333333333333333333333333333333333333333",
					Type:          datastore.ContractType("TokenAdminRegistry"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x4444444444444444444444444444444444444444",
					Type:          datastore.ContractType("MCM"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x5555555555555555555555555555555555555555",
					Type:          datastore.ContractType("Timelock"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				// Add CCTP-specific addresses
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x6666666666666666666666666666666666666666",
					Type:          datastore.ContractType("USDCTokenPoolProxy"),
					Version:       semver.MustParse("1.7.0"),
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x7777777777777777777777777777777777777777",
					Type:          datastore.ContractType("CCTPVerifier"),
					Version:       semver.MustParse("1.7.0"),
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: chainSelector,
					Address:       "0x8888888888888888888888888888888888888888",
					Type:          datastore.ContractType("MessageTransmitterProxy"),
					Version:       semver.MustParse("1.7.0"),
				})
				require.NoError(t, err)
				return ds
			},
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x9999999999999999999999999999999999999999",
						TokenMessengerV2: "0x9999999999999999999999999999999999999999",
						USDCToken:        "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
						StorageLocations: []string{"storage1", "storage2"},
						FeeAggregator:    "0xCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC",
						FastFinalityBps:  100,
						RemoteChains:     make(map[uint64]adapters.RemoteCCTPChainConfig),
					},
				},
				MCMS: &cctpTest_BasicMCMSInput,
			},
		},
		{
			desc: "error - no adapter registered",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return datastore.NewMemoryDataStore()
			},
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType: adapters.Canonical,
					},
				},
				MCMS: &cctpTest_BasicMCMSInput,
			},
			expectedError: "no CCTP adapter registered",
		},
		{
			desc: "error - sequence execution fails",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return datastore.NewMemoryDataStore()
			},
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType: adapters.Canonical,
					},
				},
				MCMS: &cctpTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "sequence execution failed",
			expectedError:            "failed to deploy CCTP on chain",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			// Create environment with datastore
			ds := tt.makeDataStore(t)
			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")
			bundle := cldf_ops.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				cldf_ops.NewMemoryReporter(),
			)
			e := deployment.Environment{
				OperationsBundle: bundle,
				DataStore:        ds.Seal(),
			}

			cctpChainRegistry := adapters.NewCCTPChainRegistry()
			mcmsRegistry := changesets.GetRegistry()
			mcmsRegistry.RegisterMCMSReader("evm", &cctpTest_MockReader{})

			// Register mock adapter for successful tests
			if tt.expectedError == "" || tt.expectedSequenceErrorMsg != "" {
				mockAdapter := &cctpTest_MockCCTPChain{
					sequenceErrorMsg: tt.expectedSequenceErrorMsg,
				}
				cctpChainRegistry.RegisterCCTPChain("evm", mockAdapter)
			}

			changeset := v1_7_0_changesets.DeployCCTPChains(cctpChainRegistry, mcmsRegistry)
			output, err := changeset.Apply(e, tt.cfg)

			if tt.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.expectedError)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, output)
		})
	}
}

func TestDeployCCTPChains_VerifyPreconditions(t *testing.T) {
	tests := []struct {
		desc          string
		cfg           v1_7_0_changesets.DeployCCTPChainsConfig
		expectedError string
	}{
		{
			desc: "success - valid configuration",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x9999999999999999999999999999999999999999",
						TokenMessengerV2: "0x8888888888888888888888888888888888888888",
						RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
							15971525489660198786: {},
						},
					},
					15971525489660198786: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x7777777777777777777777777777777777777777",
						TokenMessengerV2: "0x6666666666666666666666666666666666666666",
						RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
							5009297550715157269: {},
						},
					},
				},
				MCMS: &cctpTest_BasicMCMSInput,
			},
		},
		{
			desc: "success - no MCMS config",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x9999999999999999999999999999999999999999",
						TokenMessengerV2: "0x8888888888888888888888888888888888888888",
					},
				},
			},
		},
		{
			desc: "success - empty CCTP v1 token messenger for V2-only chain",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "",
						TokenMessengerV2: "0x8888888888888888888888888888888888888888",
					},
				},
			},
		},
		{
			desc: "success - empty CCTP v1 token messenger with CCTP_V1 lane in config",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "",
						TokenMessengerV2: "0x8888888888888888888888888888888888888888",
						RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
							15971525489660198786: {
								LockOrBurnMechanism: "CCTP_V1",
							},
						},
					},
					15971525489660198786: {
						USDCType:         adapters.Canonical,
						TokenMessengerV2: "0x6666666666666666666666666666666666666666",
					},
				},
			},
		},
		{
			desc: "failure - invalid CCTP type",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType: "invalid-type",
					},
				},
			},
			expectedError: "invalid CCTP type",
		},
		{
			desc: "failure - invalid CCTP v1 token messenger",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "not-an-address",
						TokenMessengerV2: "0x8888888888888888888888888888888888888888",
					},
				},
			},
			expectedError: "invalid TokenMessengerV1",
		},
		{
			desc: "failure - invalid CCTP v2 token messenger",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x9999999999999999999999999999999999999999",
						TokenMessengerV2: "not-an-address",
					},
				},
			},
			expectedError: "invalid TokenMessengerV2",
		},
		{
			desc: "failure - unknown chain selector",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					0: {
						USDCType: adapters.Canonical,
					}, // Invalid chain selector
				},
			},
			expectedError: "unknown chain selector",
		},
		{
			desc: "failure - unknown remote chain selector",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x9999999999999999999999999999999999999999",
						TokenMessengerV2: "0x8888888888888888888888888888888888888888",
						RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
							0: {}, // Invalid remote chain selector
						},
					},
				},
			},
			expectedError: "unknown chain selector",
		},
		{
			desc: "failure - invalid MCMS timelock action",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType: adapters.Canonical,
					},
				},
				MCMS: &mcms.Input{
					OverridePreviousRoot: true,
					ValidUntil:           3759765795,
					TimelockDelay:        mcms_types.MustParseDuration("1h"),
					TimelockAction:       "InvalidAction", // Invalid action
				},
			},
			expectedError: "failed to validate MCMS input",
		},
		{
			desc: "failure - multiple remote chains with one having unknown selector",
			cfg: v1_7_0_changesets.DeployCCTPChainsConfig{
				Chains: map[uint64]v1_7_0_changesets.CCTPChainConfig{
					5009297550715157269: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x9999999999999999999999999999999999999999",
						TokenMessengerV2: "0x8888888888888888888888888888888888888888",
						RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
							15971525489660198786: {},
							0:                    {}, // Invalid remote chain selector
						},
					},
					15971525489660198786: {
						USDCType:         adapters.Canonical,
						TokenMessengerV1: "0x7777777777777777777777777777777777777777",
						TokenMessengerV2: "0x6666666666666666666666666666666666666666",
						RemoteChains: map[uint64]adapters.RemoteCCTPChainConfig{
							5009297550715157269: {},
						},
					},
				},
			},
			expectedError: "unknown chain selector",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			// Create CCTP chain registry and MCMS registry (not used in verify, but required for changeset creation)
			cctpChainRegistry := adapters.NewCCTPChainRegistry()
			mcmsRegistry := changesets.GetRegistry()

			// Create changeset
			changeset := v1_7_0_changesets.DeployCCTPChains(cctpChainRegistry, mcmsRegistry)

			// Create minimal environment (verify doesn't use datastore or other fields)
			lggr, err := logger.New()
			require.NoError(t, err, "Failed to create logger")
			bundle := cldf_ops.NewBundle(
				func() context.Context { return context.Background() },
				lggr,
				cldf_ops.NewMemoryReporter(),
			)
			e := deployment.Environment{
				OperationsBundle: bundle,
				DataStore:        datastore.NewMemoryDataStore().Seal(),
			}

			// Run verify
			err = changeset.VerifyPreconditions(e, tt.cfg)
			if tt.expectedError != "" {
				require.Error(t, err)
				require.ErrorContains(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
