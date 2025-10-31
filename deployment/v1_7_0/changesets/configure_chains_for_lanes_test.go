package changesets_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	v1_7_0_changesets "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type lanesTest_MockReader struct{}

const LANES_OP_COUNT = 42

func (m *lanesTest_MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, input mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{
		MCMAddress:      input.MCMSAddressRef.Address,
		StartingOpCount: LANES_OP_COUNT,
	}, nil
}

type lanesTest_MockChainFamily struct {
	errorMsg         string
	sequenceErrorMsg string
}

func (ma *lanesTest_MockChainFamily) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if ma.errorMsg != "" {
		return nil, errors.New(ma.errorMsg)
	}
	return []byte(ref.Address), nil
}

func (ma *lanesTest_MockChainFamily) ConfigureChainForLanes() *cldf_ops.Sequence[adapters.ConfigureChainForLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-chain-for-lanes-sequence",
		semver.MustParse("1.0.0"),
		"Mock sequence for testing chain configuration",
		func(bundle cldf_ops.Bundle, deps cldf_chain.BlockChains, input adapters.ConfigureChainForLanesInput) (sequences.OnChainOutput, error) {
			if ma.sequenceErrorMsg != "" {
				return sequences.OnChainOutput{}, errors.New(ma.sequenceErrorMsg)
			}
			batchOps := make([]mcms_types.BatchOperation, 0, len(input.RemoteChains))
			for remoteChainSel, remoteChain := range input.RemoteChains {
				batchOps = append(batchOps, mcms_types.BatchOperation{
					ChainSelector: mcms_types.ChainSelector(remoteChainSel),
					Transactions: []mcms_types.Transaction{
						{
							To:               input.Router,
							Data:             remoteChain.OnRamp,
							AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
						},
						{
							To:               input.OnRamp,
							Data:             remoteChain.OffRamp,
							AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
						},
					},
				})
			}
			return sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{
					{
						ChainSelector: input.ChainSelector,
						Address:       input.Router,
						Type:          datastore.ContractType("Router"),
						Version:       semver.MustParse("1.0.0"),
					},
					{
						ChainSelector: input.ChainSelector,
						Address:       input.OnRamp,
						Type:          datastore.ContractType("OnRamp"),
						Version:       semver.MustParse("1.0.0"),
					},
					{
						ChainSelector: input.ChainSelector,
						Address:       input.FeeQuoter,
						Type:          datastore.ContractType("FeeQuoter"),
						Version:       semver.MustParse("1.0.0"),
					},
					{
						ChainSelector: input.ChainSelector,
						Address:       input.OffRamp,
						Type:          datastore.ContractType("OffRamp"),
						Version:       semver.MustParse("1.0.0"),
					},
				},
				BatchOps: batchOps,
			}, nil
		},
	)
}

var lanesTest_BasicMCMSInput = mcms.Input{
	OverridePreviousRoot: true,
	ValidUntil:           3759765795,
	TimelockDelay:        mcms_types.MustParseDuration("1h"),
	TimelockAction:       mcms_types.TimelockActionSchedule,
	MCMSAddressRef: datastore.AddressRef{
		Type:    "MCM",
		Version: semver.MustParse("1.0.0"),
	},
	TimelockAddressRef: datastore.AddressRef{
		Type:    "Timelock",
		Version: semver.MustParse("1.0.0"),
	},
}

func makeBaseChainDataStore(t *testing.T, chains []uint64) *datastore.MemoryDataStore {
	ds := datastore.NewMemoryDataStore()

	for _, chain := range chains {
		// Router
		err := ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-router", chain),
			Type:          datastore.ContractType("Router"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// OnRamp
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-onramp", chain),
			Type:          datastore.ContractType("OnRamp"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// OffRamp
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-offramp", chain),
			Type:          datastore.ContractType("OffRamp"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// FeeQuoter
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-fee-quoter", chain),
			Type:          datastore.ContractType("FeeQuoter"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// CommitteeVerifier
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-committee-verifier", chain),
			Type:          datastore.ContractType("CommitteeVerifier"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// Executor
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-executor", chain),
			Type:          datastore.ContractType("Executor"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// MCMS
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-mcm", chain),
			Type:          datastore.ContractType("MCM"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)

		// Timelock
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-timelock", chain),
			Type:          datastore.ContractType("Timelock"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)
	}

	return ds
}

func TestConfigureChainsForLanes_Apply(t *testing.T) {
	tests := []struct {
		desc                     string
		makeDataStore            func(t *testing.T) *datastore.MemoryDataStore
		cfg                      v1_7_0_changesets.ConfigureChainsForLanesConfig
		expectedSequenceErrorMsg string
	}{
		{
			desc: "success - basic lane configuration",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseChainDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						CommitteeVerifiers: []datastore.AddressRef{
							{
								Type:          "CommitteeVerifier",
								Version:       semver.MustParse("1.0.0"),
								ChainSelector: 5009297550715157269,
							},
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								AllowTrafficFrom: true,
								OnRamp: datastore.AddressRef{
									Type:          "OnRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OffRamp: datastore.AddressRef{
									Type:          "OffRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								DefaultExecutor: datastore.AddressRef{
									Type:          "Executor",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
								CommitteeVerifierDestChainConfig: adapters.CommitteeVerifierDestChainConfig{
									AllowlistEnabled:   false,
									FeeUSDCents:        10,
									GasForVerification: 100000,
									PayloadSizeBytes:   256,
								},
								FeeQuoterDestChainConfig: adapters.FeeQuoterDestChainConfig{
									IsEnabled:                   true,
									MaxDataBytes:                1000,
									MaxPerMsgGasLimit:           3000000,
									DestGasOverhead:             50000,
									DestGasPerPayloadByteBase:   16,
									ChainFamilySelector:         [4]byte{0x1, 0x2, 0x3, 0x4},
									DefaultTokenFeeUSDCents:     5,
									DefaultTokenDestGasOverhead: 10000,
									DefaultTxGasLimit:           200000,
									NetworkFeeUSDCents:          100,
								},
								ExecutorDestChainConfig: adapters.ExecutorDestChainConfig{
									USDCentsFee:            20,
									BaseExecGas:            100000,
									DestAddressLengthBytes: 20,
									Enabled:                true,
								},
							},
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
		},
		{
			desc: "success - multiple committee verifiers",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseChainDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
				// Update the first committee verifier to have a qualifier
				err := ds.Addresses().Delete(datastore.NewAddressRefKey(5009297550715157269, "CommitteeVerifier", semver.MustParse("1.0.0"), ""))
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       fmt.Sprintf("%d-committee-verifier", 5009297550715157269),
					Type:          datastore.ContractType("CommitteeVerifier"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "primary",
				})
				require.NoError(t, err)
				// Add second committee verifier
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       fmt.Sprintf("%d-committee-verifier-2", 5009297550715157269),
					Type:          datastore.ContractType("CommitteeVerifier"),
					Version:       semver.MustParse("1.0.0"),
					Qualifier:     "secondary",
				})
				require.NoError(t, err)
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						CommitteeVerifiers: []datastore.AddressRef{
							{
								Type:          "CommitteeVerifier",
								Version:       semver.MustParse("1.0.0"),
								ChainSelector: 5009297550715157269,
								Qualifier:     "primary",
							},
							{
								Type:          "CommitteeVerifier",
								Version:       semver.MustParse("1.0.0"),
								ChainSelector: 5009297550715157269,
								Qualifier:     "secondary",
							},
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								AllowTrafficFrom: true,
								OnRamp: datastore.AddressRef{
									Type:          "OnRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OffRamp: datastore.AddressRef{
									Type:          "OffRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								DefaultExecutor: datastore.AddressRef{
									Type:          "Executor",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
							},
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
		},
		{
			desc: "failure to resolve router ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				// Don't add router address
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve router ref",
		},
		{
			desc: "failure to resolve onramp ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				// Only add router, not onramp
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       "5009297550715157269-router",
					Type:          datastore.ContractType("Router"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve onRamp ref",
		},
		{
			desc: "failure to resolve fee quoter ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				// Add router and onramp, but not fee quoter
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       "5009297550715157269-router",
					Type:          datastore.ContractType("Router"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				err = ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       "5009297550715157269-onramp",
					Type:          datastore.ContractType("OnRamp"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve feeQuoter ref",
		},
		{
			desc: "failure to resolve offramp ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseChainDataStore(t, []uint64{5009297550715157269})
				// Remove offramp
				err := ds.Addresses().Delete(datastore.NewAddressRefKey(5009297550715157269, "OffRamp", semver.MustParse("1.0.0"), ""))
				require.NoError(t, err)
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve offRamp ref",
		},
		{
			desc: "failure to resolve committee verifier ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseChainDataStore(t, []uint64{5009297550715157269})
				// Remove committee verifier
				err := ds.Addresses().Delete(datastore.NewAddressRefKey(5009297550715157269, "CommitteeVerifier", semver.MustParse("1.0.0"), ""))
				require.NoError(t, err)
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						CommitteeVerifiers: []datastore.AddressRef{
							{
								Type:          "CommitteeVerifier",
								Version:       semver.MustParse("1.0.0"),
								ChainSelector: 5009297550715157269,
							},
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve committeeVerifier ref",
		},
		{
			desc: "invalid chain selector",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseChainDataStore(t, []uint64{5009297550715157269})
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 0, // Invalid chain selector
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to get chain family",
		},
		{
			desc: "failure to resolve remote onramp ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseChainDataStore(t, []uint64{5009297550715157269})
				// Don't add remote chain data
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								AllowTrafficFrom: true,
								OnRamp: datastore.AddressRef{
									Type:          "OnRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OffRamp: datastore.AddressRef{
									Type:          "OffRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								DefaultExecutor: datastore.AddressRef{
									Type:          "Executor",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
							},
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve onRamp ref on remote chain",
		},
		{
			desc: "failure to resolve executor ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseChainDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
				// Remove executor
				err := ds.Addresses().Delete(datastore.NewAddressRefKey(5009297550715157269, "Executor", semver.MustParse("1.0.0"), ""))
				require.NoError(t, err)
				return ds
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								AllowTrafficFrom: true,
								OnRamp: datastore.AddressRef{
									Type:          "OnRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OffRamp: datastore.AddressRef{
									Type:          "OffRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								DefaultExecutor: datastore.AddressRef{
									Type:          "Executor",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
							},
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve executor ref",
		},
		{
			desc: "failure to execute sequence",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseChainDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: v1_7_0_changesets.ConfigureChainsForLanesConfig{
				Chains: []v1_7_0_changesets.ChainConfig{
					{
						ChainSelector: 5009297550715157269,
						Router: datastore.AddressRef{
							Type:          "Router",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OnRamp: datastore.AddressRef{
							Type:          "OnRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						FeeQuoter: datastore.AddressRef{
							Type:          "FeeQuoter",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						OffRamp: datastore.AddressRef{
							Type:          "OffRamp",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]adapters.RemoteChainConfig[datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								AllowTrafficFrom: true,
								OnRamp: datastore.AddressRef{
									Type:          "OnRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								OffRamp: datastore.AddressRef{
									Type:          "OffRamp",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								DefaultExecutor: datastore.AddressRef{
									Type:          "Executor",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 5009297550715157269,
								},
							},
						},
					},
				},
				MCMS: lanesTest_BasicMCMSInput,
			},
			expectedSequenceErrorMsg: "sequence execution failed",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			// Register MCMS reader
			mcmsRegistry := changesets.NewMCMSReaderRegistry()
			mcmsRegistry.RegisterMCMSReader("evm", &lanesTest_MockReader{})

			// Register chain family adapter with appropriate error condition
			chainFamilyRegistry := adapters.NewChainFamilyRegistry()
			mockAdapter := &lanesTest_MockChainFamily{}

			// Set error conditions for specific test cases
			mockAdapter.sequenceErrorMsg = tt.expectedSequenceErrorMsg

			chainFamilyRegistry.RegisterChainFamily("evm", mockAdapter)

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

			// Run changeset
			changeset := v1_7_0_changesets.ConfigureChainsForLanes(chainFamilyRegistry, mcmsRegistry)
			out, err := changeset.Apply(e, tt.cfg)
			if tt.expectedSequenceErrorMsg != "" {
				require.ErrorContains(t, err, tt.expectedSequenceErrorMsg)
				return
			}
			require.NoError(t, err)

			// Validate output
			require.Len(t, out.MCMSTimelockProposals, 1)
			for _, proposal := range out.MCMSTimelockProposals {
				require.Equal(t, tt.cfg.MCMS.Description, proposal.Description)
				require.Equal(t, tt.cfg.MCMS.OverridePreviousRoot, proposal.OverridePreviousRoot)
				require.Equal(t, tt.cfg.MCMS.ValidUntil, proposal.ValidUntil)
				require.Equal(t, tt.cfg.MCMS.TimelockDelay, proposal.Delay)
				require.Equal(t, tt.cfg.MCMS.TimelockAction, proposal.Action)

				// Validate that batch operations were created for each chain
				require.Len(t, proposal.Operations, len(tt.cfg.Chains))
				for _, op := range proposal.Operations {
					// Validate transactions were created for remote chains
					require.Greater(t, len(op.Transactions), 0)
				}
			}
		})
	}
}
