package tokens_test

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type MockReader struct{}

const OP_COUNT = 42

func (m *MockReader) GetChainMetadata(_ deployment.Environment, _ uint64, input mcms.Input) (mcms_types.ChainMetadata, error) {
	return mcms_types.ChainMetadata{
		StartingOpCount: OP_COUNT,
	}, nil
}

func (m *MockReader) GetTimelockRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0x01",
		Type:          datastore.ContractType("Timelock"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

func (m *MockReader) GetMCMSRef(_ deployment.Environment, selector uint64, _ mcms.Input) (datastore.AddressRef, error) {
	return datastore.AddressRef{
		ChainSelector: selector,
		Address:       "0x02",
		Type:          datastore.ContractType("MCM"),
		Version:       semver.MustParse("1.0.0"),
	}, nil
}

var transfersTest_NewTokenAdapterRegistry = tokens.NewTokenAdapterRegistry()

type transfersTest_MockTokenAdapter struct {
	errorMsg            string
	deriveTokenErrorMsg string
	sequenceErrorMsg    string
}

func (ma *transfersTest_MockTokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return []byte(ref.Address), nil
}

func (ma *transfersTest_MockTokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-configure-token-sequence",
		semver.MustParse("1.0.0"),
		"Mock sequence for testing token configuration",
		func(bundle cldf_ops.Bundle, deps cldf_chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (sequences.OnChainOutput, error) {
			if ma.sequenceErrorMsg != "" {
				return sequences.OnChainOutput{}, errors.New(ma.sequenceErrorMsg)
			}
			batchOps := make([]mcms_types.BatchOperation, 0, len(input.RemoteChains))
			for remoteChainSel, remoteChain := range input.RemoteChains {
				batchOps = append(batchOps, mcms_types.BatchOperation{
					ChainSelector: mcms_types.ChainSelector(remoteChainSel),
					Transactions: []mcms_types.Transaction{
						{
							To:               string(remoteChain.RemotePool),
							Data:             remoteChain.RemoteToken,
							AdditionalFields: []byte{0x7B, 0x7D}, // "{}" in bytes
						},
					},
				})
			}
			return sequences.OnChainOutput{
				Addresses: []datastore.AddressRef{
					{
						ChainSelector: input.ChainSelector,
						Address:       input.RegistryAddress,
						Type:          datastore.ContractType("Registry"),
						Version:       semver.MustParse("1.0.0"),
					},
					{
						ChainSelector: input.ChainSelector,
						Address:       input.TokenPoolAddress,
						Type:          datastore.ContractType("TokenPool"),
						Version:       semver.MustParse("1.0.0"),
					},
					{
						ChainSelector: input.ChainSelector,
						Address:       input.ExternalAdmin,
						Type:          datastore.ContractType("EOA"),
						Version:       semver.MustParse("1.0.0"),
					},
				},
				BatchOps: batchOps,
			}, nil
		},
	)
}

func (ma *transfersTest_MockTokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	if poolRef.Address == "" {
		// Address in pool ref MUST be populated before this adapter method gets called.
		// Most implementations will require the pool address in order to fetch the token address.
		return nil, errors.New("pool ref address is empty")
	}
	if ma.deriveTokenErrorMsg != "" {
		return nil, errors.New(ma.deriveTokenErrorMsg)
	}
	return []byte("mocked-remote-token-address"), nil
}

var basicMCMSInput = mcms.Input{
	OverridePreviousRoot: true,
	ValidUntil:           3759765795,
	TimelockDelay:        mcms_types.MustParseDuration("1h"),
	TimelockAction:       mcms_types.TimelockActionSchedule,
}

func makeBaseDataStore(t *testing.T, chains []uint64) *datastore.MemoryDataStore {
	ds := datastore.NewMemoryDataStore()

	for _, chain := range chains {
		err := ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-token-pool", chain),
			Type:          datastore.ContractType("TokenPool"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)
		ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-registry", chain),
			Type:          datastore.ContractType("Registry"),
			Version:       semver.MustParse("1.0.0"),
		})
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-token", chain),
			Type:          datastore.ContractType("Token"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)
		err = ds.Addresses().Add(datastore.AddressRef{
			ChainSelector: chain,
			Address:       fmt.Sprintf("%d-mcm", chain),
			Type:          datastore.ContractType("MCM"),
			Version:       semver.MustParse("1.0.0"),
		})
		require.NoError(t, err)
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

func TestConfigureTokensForTransfers_Apply(t *testing.T) {
	tests := []struct {
		desc                            string
		makeDataStore                   func(t *testing.T) *datastore.MemoryDataStore
		cfg                             tokens.ConfigureTokensForTransfersConfig
		shouldDeriveToken               bool
		expectedSequenceErrorMsg        string
		expectedTokenDerivationErrorMsg string
	}{
		{
			desc: "success - inputted remote token",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								InboundRateLimiterConfig: tokens.RateLimiterConfig{
									IsEnabled: true,
									Capacity:  big.NewInt(1000),
									Rate:      big.NewInt(100),
								},
								OutboundRateLimiterConfig: tokens.RateLimiterConfig{
									IsEnabled: true,
									Capacity:  big.NewInt(1000),
									Rate:      big.NewInt(100),
								},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
		},
		{
			desc: "success - derived remote token",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: nil, // This will trigger derivation
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								InboundRateLimiterConfig: tokens.RateLimiterConfig{
									IsEnabled: true,
									Capacity:  big.NewInt(1000),
									Rate:      big.NewInt(100),
								},
								OutboundRateLimiterConfig: tokens.RateLimiterConfig{
									IsEnabled: true,
									Capacity:  big.NewInt(1000),
									Rate:      big.NewInt(100),
								},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			shouldDeriveToken: true,
		},
		{
			desc: "failure to resolve token pool ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				// Don't add token pool address - only add registry
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       "5009297550715157269-registry",
					Type:          datastore.ContractType("Registry"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve token pool ref",
		},
		{
			desc: "failure to resolve registry ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := datastore.NewMemoryDataStore()
				// Only add token pool, not registry
				err := ds.Addresses().Add(datastore.AddressRef{
					ChainSelector: 5009297550715157269,
					Address:       "5009297550715157269-token-pool",
					Type:          datastore.ContractType("TokenPool"),
					Version:       semver.MustParse("1.0.0"),
				})
				require.NoError(t, err)
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve registry ref",
		},
		{
			desc: "invalid chain selector",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 0, // Invalid chain selector
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "unknown chain selector",
		},
		{
			desc: "failure to resolve remote pool ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseDataStore(t, []uint64{5009297550715157269})
				// Don't add remote chain data
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								InboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false},
								OutboundRateLimiterConfig: tokens.RateLimiterConfig{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve remote pool ref",
		},
		{
			desc: "failure to resolve remote token ref",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				ds := makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
				// Remove remote token
				err := ds.Addresses().Delete(datastore.NewAddressRefKey(15971525489660198786, "Token", semver.MustParse("1.0.0"), ""))
				require.NoError(t, err)
				return ds
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								InboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false},
								OutboundRateLimiterConfig: tokens.RateLimiterConfig{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "failed to resolve remote token ref",
		},
		{
			desc: "failure to derive remote token address",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: nil, // Will trigger derivation which should fail
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								InboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false},
								OutboundRateLimiterConfig: tokens.RateLimiterConfig{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedTokenDerivationErrorMsg: "failed to derive remote token address",
		},
		{
			desc: "failure to execute sequence",
			makeDataStore: func(t *testing.T) *datastore.MemoryDataStore {
				return makeBaseDataStore(t, []uint64{5009297550715157269, 15971525489660198786})
			},
			cfg: tokens.ConfigureTokensForTransfersConfig{
				Tokens: []tokens.TokenTransferConfig{
					{
						ChainSelector: 5009297550715157269,
						TokenPoolRef: datastore.AddressRef{
							Type:          "TokenPool",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RegistryRef: datastore.AddressRef{
							Type:          "Registry",
							Version:       semver.MustParse("1.0.0"),
							ChainSelector: 5009297550715157269,
						},
						RemoteChains: map[uint64]tokens.RemoteChainConfig[*datastore.AddressRef, datastore.AddressRef]{
							15971525489660198786: {
								RemoteToken: &datastore.AddressRef{
									Type:          "Token",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								RemotePool: &datastore.AddressRef{
									Type:          "TokenPool",
									Version:       semver.MustParse("1.0.0"),
									ChainSelector: 15971525489660198786,
								},
								InboundRateLimiterConfig:  tokens.RateLimiterConfig{IsEnabled: false},
								OutboundRateLimiterConfig: tokens.RateLimiterConfig{IsEnabled: false},
							},
						},
					},
				},
				MCMS: basicMCMSInput,
			},
			expectedSequenceErrorMsg: "sequence execution failed",
		},
	}
	// Register MCMS reader
	mcmsRegistry := changesets.GetRegistry()
	mcmsRegistry.RegisterMCMSReader("evm", &MockReader{})

	for _, tt := range tests {
		tt := tt
		t.Run(tt.desc, func(t *testing.T) {
			// Register token adapter with appropriate error condition
			tokenAdapterRegistry := tokens.NewTokenAdapterRegistry()
			mockAdapter := &transfersTest_MockTokenAdapter{}

			// Set error conditions for specific test cases
			mockAdapter.deriveTokenErrorMsg = tt.expectedTokenDerivationErrorMsg
			mockAdapter.sequenceErrorMsg = tt.expectedSequenceErrorMsg

			tokenAdapterRegistry.RegisterTokenAdapter("evm", semver.MustParse("1.0.0"), mockAdapter)

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
			changeset := tokens.ConfigureTokensForTransfers(tokenAdapterRegistry, mcmsRegistry)
			out, err := changeset.Apply(e, tt.cfg)
			if tt.expectedSequenceErrorMsg != "" {
				require.ErrorContains(t, err, tt.expectedSequenceErrorMsg)
				return
			}
			if tt.expectedTokenDerivationErrorMsg != "" {
				require.ErrorContains(t, err, tt.expectedTokenDerivationErrorMsg)
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

				require.Len(t, proposal.Operations, len(tt.cfg.Tokens))
				for _, op := range proposal.Operations {
					// For derived remote token test, expect mocked address
					if tt.shouldDeriveToken {
						require.Equal(t, []byte("mocked-remote-token-address"), op.Transactions[0].Data)
					} else {
						require.Equal(t, []byte(fmt.Sprintf("%d-token", op.ChainSelector)), op.Transactions[0].Data)
					}
					require.Equal(t, fmt.Sprintf("%d-token-pool", op.ChainSelector), op.Transactions[0].To)
				}
			}
		})
	}
}
