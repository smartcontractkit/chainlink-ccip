package changesets_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

var _ adapters.TokenVerifierConfigAdapter = (*mockTokenVerifierConfigAdapter)(nil)

type mockTokenVerifierConfigAdapter struct {
	addressesByChain map[uint64]*adapters.TokenVerifierChainAddresses
	err              error
}

func (m *mockTokenVerifierConfigAdapter) ResolveTokenVerifierAddresses(
	_ datastore.DataStore, sel uint64, _ string, _ string,
) (*adapters.TokenVerifierChainAddresses, error) {
	if m.err != nil {
		return nil, m.err
	}
	addrs, ok := m.addressesByChain[sel]
	if !ok {
		return nil, fmt.Errorf("no addresses for chain %d", sel)
	}
	return addrs, nil
}

func TestGenerateTokenVerifierConfig(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	tests := []struct {
		name              string
		mock              *mockTokenVerifierConfigAdapter
		input             changesets.GenerateTokenVerifierConfigInput
		selectors         []uint64
		envName           string
		wantErr           string
		validateOnErr     bool
		expectedVerifiers int
	}{
		{
			name: "validates service identifier is required",
			mock: &mockTokenVerifierConfigAdapter{},
			input: changesets.GenerateTokenVerifierConfigInput{
				ServiceIdentifier: "",
			},
			selectors:     []uint64{sel1},
			wantErr:       "service identifier is required",
			validateOnErr: true,
		},
		{
			name: "validates chain selectors exist in environment",
			mock: &mockTokenVerifierConfigAdapter{},
			input: changesets.GenerateTokenVerifierConfigInput{
				ServiceIdentifier: "test",
				ChainSelectors:    []uint64{9999},
			},
			selectors:     []uint64{sel1},
			wantErr:       "selector 9999 is not available in environment",
			validateOnErr: true,
		},
		{
			name: "returns error when adapter fails",
			mock: &mockTokenVerifierConfigAdapter{
				err: fmt.Errorf("adapter failure"),
			},
			input: changesets.GenerateTokenVerifierConfigInput{
				ServiceIdentifier: "test",
				ChainSelectors:    []uint64{sel1},
			},
			selectors: []uint64{sel1},
			wantErr:   "adapter failure",
		},
		{
			name: "generates config with both CCTP and Lombard verifiers",
			mock: &mockTokenVerifierConfigAdapter{
				addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
					sel1: {
						OnRampAddress:                  "0xOnRamp1",
						RMNRemoteAddress:               "0xRMN1",
						CCTPVerifierAddress:            "0xCCTP1",
						CCTPVerifierResolverAddress:    "0xCCTPResolver1",
						LombardVerifierResolverAddress: "0xLombardResolver1",
					},
					sel2: {
						OnRampAddress:                  "0xOnRamp2",
						RMNRemoteAddress:               "0xRMN2",
						CCTPVerifierAddress:            "0xCCTP2",
						CCTPVerifierResolverAddress:    "0xCCTPResolver2",
						LombardVerifierResolverAddress: "0xLombardResolver2",
					},
				},
			},
			input: changesets.GenerateTokenVerifierConfigInput{
				ServiceIdentifier: "token-verifier",
				ChainSelectors:    []uint64{sel1, sel2},
				PyroscopeURL:      "http://pyroscope:4040",
				CCTP: changesets.CCTPConfigInput{
					Qualifier:  "default",
					VerifierID: "CCTPVerifier",
				},
				Lombard: changesets.LombardConfigInput{
					Qualifier:  "default",
					VerifierID: "LombardVerifier",
				},
				Monitoring: shared.MonitoringInput{
					Enabled: true,
					Type:    "beholder",
				},
			},
			selectors:         []uint64{sel1, sel2},
			expectedVerifiers: 2,
		},
		{
			name: "generates config with only CCTP verifier",
			mock: &mockTokenVerifierConfigAdapter{
				addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
					sel1: {
						OnRampAddress:               "0xOnRamp1",
						RMNRemoteAddress:            "0xRMN1",
						CCTPVerifierAddress:         "0xCCTP1",
						CCTPVerifierResolverAddress: "0xCCTPResolver1",
					},
				},
			},
			input: changesets.GenerateTokenVerifierConfigInput{
				ServiceIdentifier: "token-verifier",
				ChainSelectors:    []uint64{sel1},
				CCTP: changesets.CCTPConfigInput{
					Qualifier:  "default",
					VerifierID: "CCTPVerifier",
				},
			},
			selectors:         []uint64{sel1},
			expectedVerifiers: 1,
		},
		{
			name: "generates config with only Lombard verifier",
			mock: &mockTokenVerifierConfigAdapter{
				addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
					sel1: {
						OnRampAddress:                  "0xOnRamp1",
						RMNRemoteAddress:               "0xRMN1",
						LombardVerifierResolverAddress: "0xLombardResolver1",
					},
				},
			},
			input: changesets.GenerateTokenVerifierConfigInput{
				ServiceIdentifier: "token-verifier",
				ChainSelectors:    []uint64{sel1},
				Lombard: changesets.LombardConfigInput{
					Qualifier:  "default",
					VerifierID: "LombardVerifier",
				},
			},
			selectors:         []uint64{sel1},
			expectedVerifiers: 1,
		},
		{
			name: "generates config with no token verifiers",
			mock: &mockTokenVerifierConfigAdapter{
				addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
					sel1: {
						OnRampAddress:    "0xOnRamp1",
						RMNRemoteAddress: "0xRMN1",
					},
				},
			},
			input: changesets.GenerateTokenVerifierConfigInput{
				ServiceIdentifier: "token-verifier",
				ChainSelectors:    []uint64{sel1},
			},
			selectors:         []uint64{sel1},
			expectedVerifiers: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			registry := adapters.NewTokenVerifierConfigRegistry()
			registry.Register(chainsel.FamilyEVM, tc.mock)

			ds := datastore.NewMemoryDataStore()
			env := deployment.Environment{
				Name:        tc.envName,
				DataStore:   ds.Seal(),
				BlockChains: newTestBlockChains(tc.selectors),
			}

			cs := changesets.GenerateTokenVerifierConfig(registry)

			if tc.validateOnErr {
				err := cs.VerifyPreconditions(env, tc.input)
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
				return
			}

			output, err := cs.Apply(env, tc.input)
			if tc.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tc.wantErr)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, output.DataStore)

			cfg, err := offchain.GetTokenVerifierConfig(output.DataStore.Seal(), tc.input.ServiceIdentifier)
			require.NoError(t, err)
			require.NotNil(t, cfg)

			assert.Equal(t, tc.input.PyroscopeURL, cfg.PyroscopeURL)
			assert.Len(t, cfg.TokenVerifiers, tc.expectedVerifiers)

			for _, sel := range tc.input.ChainSelectors {
				selStr := strconv.FormatUint(sel, 10)
				expectedAddrs, ok := tc.mock.addressesByChain[sel]
				require.True(t, ok, "test setup: missing addresses for chain %d", sel)

				assert.Equal(t, expectedAddrs.OnRampAddress, cfg.OnRampAddresses[selStr])
				assert.Equal(t, expectedAddrs.RMNRemoteAddress, cfg.RMNRemoteAddresses[selStr])
			}
		})
	}

	t.Run("both CCTP and Lombard verifiers have correct structure", func(t *testing.T) {
		sel1Str := strconv.FormatUint(sel1, 10)
		sel2Str := strconv.FormatUint(sel2, 10)

		mock := &mockTokenVerifierConfigAdapter{
			addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
				sel1: {
					OnRampAddress:                  "0xOnRamp1",
					RMNRemoteAddress:               "0xRMN1",
					CCTPVerifierAddress:            "0xCCTP1",
					CCTPVerifierResolverAddress:    "0xCCTPResolver1",
					LombardVerifierResolverAddress: "0xLombardResolver1",
				},
				sel2: {
					OnRampAddress:                  "0xOnRamp2",
					RMNRemoteAddress:               "0xRMN2",
					CCTPVerifierAddress:            "0xCCTP2",
					CCTPVerifierResolverAddress:    "0xCCTPResolver2",
					LombardVerifierResolverAddress: "0xLombardResolver2",
				},
			},
		}

		registry := adapters.NewTokenVerifierConfigRegistry()
		registry.Register(chainsel.FamilyEVM, mock)

		ds := datastore.NewMemoryDataStore()
		env := deployment.Environment{
			DataStore:   ds.Seal(),
			BlockChains: newTestBlockChains([]uint64{sel1, sel2}),
		}

		cs := changesets.GenerateTokenVerifierConfig(registry)
		output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigInput{
			ServiceIdentifier: "token-verifier",
			ChainSelectors:    []uint64{sel1, sel2},
			PyroscopeURL:      "http://pyroscope:4040",
			CCTP: changesets.CCTPConfigInput{
				Qualifier:       "default",
				VerifierID:      "CCTPVerifier",
				AttestationAPI:  "https://cctp.example.com",
				VerifierVersion: []byte{0x8e, 0x1d, 0x1a, 0x9d},
			},
			Lombard: changesets.LombardConfigInput{
				Qualifier:       "default",
				VerifierID:      "LombardVerifier",
				AttestationAPI:  "https://lombard.example.com",
				VerifierVersion: []byte{0xf0, 0xf3, 0xa1, 0x35},
			},
			Monitoring: shared.MonitoringInput{
				Enabled: true,
				Type:    "beholder",
				Beholder: shared.BeholderInput{
					InsecureConnection: true,
				},
			},
		})
		require.NoError(t, err)

		cfg, err := offchain.GetTokenVerifierConfig(output.DataStore.Seal(), "token-verifier")
		require.NoError(t, err)

		assert.Equal(t, "http://pyroscope:4040", cfg.PyroscopeURL)
		assert.Equal(t, "0xOnRamp1", cfg.OnRampAddresses[sel1Str])
		assert.Equal(t, "0xOnRamp2", cfg.OnRampAddresses[sel2Str])
		assert.Equal(t, "0xRMN1", cfg.RMNRemoteAddresses[sel1Str])
		assert.Equal(t, "0xRMN2", cfg.RMNRemoteAddresses[sel2Str])
		assert.True(t, cfg.Monitoring.Enabled)
		assert.Equal(t, "beholder", cfg.Monitoring.Type)
		assert.True(t, cfg.Monitoring.Beholder.InsecureConnection)

		require.Len(t, cfg.TokenVerifiers, 2)

		cctpEntry := cfg.TokenVerifiers[0]
		assert.Equal(t, "CCTPVerifier", cctpEntry.VerifierID)
		assert.Equal(t, "cctp", cctpEntry.Type)
		assert.Equal(t, "2.0", cctpEntry.Version)
		require.NotNil(t, cctpEntry.CCTP)
		assert.Equal(t, "https://cctp.example.com", cctpEntry.CCTP.AttestationAPI)
		assert.Equal(t, []byte{0x8e, 0x1d, 0x1a, 0x9d}, cctpEntry.CCTP.VerifierVersion)
		assert.Equal(t, "0xCCTP1", cctpEntry.CCTP.Verifiers[sel1Str])
		assert.Equal(t, "0xCCTP2", cctpEntry.CCTP.Verifiers[sel2Str])
		assert.Equal(t, "0xCCTPResolver1", cctpEntry.CCTP.VerifierResolvers[sel1Str])
		assert.Equal(t, "0xCCTPResolver2", cctpEntry.CCTP.VerifierResolvers[sel2Str])
		assert.Equal(t, 1*time.Second, cctpEntry.CCTP.AttestationAPITimeout)
		assert.Equal(t, 100*time.Millisecond, cctpEntry.CCTP.AttestationAPIInterval)
		assert.Equal(t, 5*time.Minute, cctpEntry.CCTP.AttestationAPICooldown)

		lombardEntry := cfg.TokenVerifiers[1]
		assert.Equal(t, "LombardVerifier", lombardEntry.VerifierID)
		assert.Equal(t, "lombard", lombardEntry.Type)
		assert.Equal(t, "1.0", lombardEntry.Version)
		require.NotNil(t, lombardEntry.Lombard)
		assert.Equal(t, "https://lombard.example.com", lombardEntry.Lombard.AttestationAPI)
		assert.Equal(t, []byte{0xf0, 0xf3, 0xa1, 0x35}, lombardEntry.Lombard.VerifierVersion)
		assert.Equal(t, "0xLombardResolver1", lombardEntry.Lombard.VerifierResolvers[sel1Str])
		assert.Equal(t, "0xLombardResolver2", lombardEntry.Lombard.VerifierResolvers[sel2Str])
		assert.Equal(t, 1*time.Second, lombardEntry.Lombard.AttestationAPITimeout)
		assert.Equal(t, 100*time.Millisecond, lombardEntry.Lombard.AttestationAPIInterval)
		assert.Equal(t, 20, lombardEntry.Lombard.AttestationAPIBatchSize)
	})

	t.Run("applies testnet defaults when attestation API not specified", func(t *testing.T) {
		mock := &mockTokenVerifierConfigAdapter{
			addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
				sel1: {
					OnRampAddress:                  "0xOnRamp1",
					RMNRemoteAddress:               "0xRMN1",
					CCTPVerifierAddress:            "0xCCTP1",
					CCTPVerifierResolverAddress:    "0xCCTPResolver1",
					LombardVerifierResolverAddress: "0xLombardResolver1",
				},
			},
		}

		registry := adapters.NewTokenVerifierConfigRegistry()
		registry.Register(chainsel.FamilyEVM, mock)

		env := deployment.Environment{
			Name:        "testnet",
			DataStore:   datastore.NewMemoryDataStore().Seal(),
			BlockChains: newTestBlockChains([]uint64{sel1}),
		}

		cs := changesets.GenerateTokenVerifierConfig(registry)
		output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigInput{
			ServiceIdentifier: "test",
			ChainSelectors:    []uint64{sel1},
			CCTP:              changesets.CCTPConfigInput{Qualifier: "default", VerifierID: "CCTP"},
			Lombard:           changesets.LombardConfigInput{Qualifier: "default", VerifierID: "Lombard"},
		})
		require.NoError(t, err)

		cfg, err := offchain.GetTokenVerifierConfig(output.DataStore.Seal(), "test")
		require.NoError(t, err)

		require.Len(t, cfg.TokenVerifiers, 2)
		assert.Equal(t, changesets.TestnetCCTPAttestationAPI, cfg.TokenVerifiers[0].CCTP.AttestationAPI)
		assert.Equal(t, changesets.TestnetLombardAttestationAPI, cfg.TokenVerifiers[1].Lombard.AttestationAPI)
	})

	t.Run("applies mainnet defaults when environment is mainnet", func(t *testing.T) {
		mock := &mockTokenVerifierConfigAdapter{
			addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
				sel1: {
					OnRampAddress:                  "0xOnRamp1",
					RMNRemoteAddress:               "0xRMN1",
					CCTPVerifierAddress:            "0xCCTP1",
					CCTPVerifierResolverAddress:    "0xCCTPResolver1",
					LombardVerifierResolverAddress: "0xLombardResolver1",
				},
			},
		}

		registry := adapters.NewTokenVerifierConfigRegistry()
		registry.Register(chainsel.FamilyEVM, mock)

		env := deployment.Environment{
			Name:        "mainnet",
			DataStore:   datastore.NewMemoryDataStore().Seal(),
			BlockChains: newTestBlockChains([]uint64{sel1}),
		}

		cs := changesets.GenerateTokenVerifierConfig(registry)
		output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigInput{
			ServiceIdentifier: "test",
			ChainSelectors:    []uint64{sel1},
			CCTP:              changesets.CCTPConfigInput{Qualifier: "default", VerifierID: "CCTP"},
			Lombard:           changesets.LombardConfigInput{Qualifier: "default", VerifierID: "Lombard"},
		})
		require.NoError(t, err)

		cfg, err := offchain.GetTokenVerifierConfig(output.DataStore.Seal(), "test")
		require.NoError(t, err)

		require.Len(t, cfg.TokenVerifiers, 2)
		assert.Equal(t, changesets.MainnetCCTPAttestationAPI, cfg.TokenVerifiers[0].CCTP.AttestationAPI)
		assert.Equal(t, changesets.MainnetLombardAttestationAPI, cfg.TokenVerifiers[1].Lombard.AttestationAPI)
	})

	t.Run("missing adapter for chain family returns error", func(t *testing.T) {
		registry := adapters.NewTokenVerifierConfigRegistry()

		env := deployment.Environment{
			DataStore:   datastore.NewMemoryDataStore().Seal(),
			BlockChains: newTestBlockChains([]uint64{sel1}),
		}

		cs := changesets.GenerateTokenVerifierConfig(registry)
		_, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigInput{
			ServiceIdentifier: "test",
			ChainSelectors:    []uint64{sel1},
		})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "no token verifier config adapter registered")
	})

	t.Run("default verifier ID fallback uses qualifier", func(t *testing.T) {
		mock := &mockTokenVerifierConfigAdapter{
			addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
				sel1: {
					OnRampAddress:                  "0xOnRamp1",
					RMNRemoteAddress:               "0xRMN1",
					CCTPVerifierAddress:            "0xCCTP1",
					CCTPVerifierResolverAddress:    "0xCCTPResolver1",
					LombardVerifierResolverAddress: "0xLombardResolver1",
				},
			},
		}

		registry := adapters.NewTokenVerifierConfigRegistry()
		registry.Register(chainsel.FamilyEVM, mock)

		env := deployment.Environment{
			DataStore:   datastore.NewMemoryDataStore().Seal(),
			BlockChains: newTestBlockChains([]uint64{sel1}),
		}

		cs := changesets.GenerateTokenVerifierConfig(registry)
		output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigInput{
			ServiceIdentifier: "test",
			ChainSelectors:    []uint64{sel1},
			CCTP:              changesets.CCTPConfigInput{Qualifier: "mycctp"},
			Lombard:           changesets.LombardConfigInput{Qualifier: "mylombard"},
		})
		require.NoError(t, err)

		cfg, err := offchain.GetTokenVerifierConfig(output.DataStore.Seal(), "test")
		require.NoError(t, err)

		require.Len(t, cfg.TokenVerifiers, 2)
		assert.Equal(t, "cctp-mycctp", cfg.TokenVerifiers[0].VerifierID)
		assert.Equal(t, "lombard-mylombard", cfg.TokenVerifiers[1].VerifierID)
	})

	t.Run("empty chain selectors defaults to all environment chains", func(t *testing.T) {
		mock := &mockTokenVerifierConfigAdapter{
			addressesByChain: map[uint64]*adapters.TokenVerifierChainAddresses{
				sel1: {
					OnRampAddress:    "0xOnRamp1",
					RMNRemoteAddress: "0xRMN1",
				},
				sel2: {
					OnRampAddress:    "0xOnRamp2",
					RMNRemoteAddress: "0xRMN2",
				},
			},
		}

		registry := adapters.NewTokenVerifierConfigRegistry()
		registry.Register(chainsel.FamilyEVM, mock)

		env := deployment.Environment{
			DataStore:   datastore.NewMemoryDataStore().Seal(),
			BlockChains: newTestBlockChains([]uint64{sel1, sel2}),
		}

		cs := changesets.GenerateTokenVerifierConfig(registry)
		output, err := cs.Apply(env, changesets.GenerateTokenVerifierConfigInput{
			ServiceIdentifier: "test",
		})
		require.NoError(t, err)

		cfg, err := offchain.GetTokenVerifierConfig(output.DataStore.Seal(), "test")
		require.NoError(t, err)

		assert.Len(t, cfg.OnRampAddresses, 2)
		assert.Len(t, cfg.RMNRemoteAddresses, 2)
	})
}
