package tokens_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type productTest_MockTokenAdapter struct{}

func (ma *productTest_MockTokenAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return []byte{}, nil
}

func (ma *productTest_MockTokenAdapter) ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) (string, error) {
	return "", nil
}

func (ma *productTest_MockTokenAdapter) DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error) {
	return 18, nil
}

func (ma *productTest_MockTokenAdapter) DeriveTokenPoolCounterpart(e deployment.Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error) {
	return []byte{}, nil
}

func (ma *productTest_MockTokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokens.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) SetTokenPoolRateLimits() *cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) DeployToken() *cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) DeployTokenVerify(e deployment.Environment, in tokens.DeployTokenInput) error {
	return nil
}

func (ma *productTest_MockTokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment] {
	return &cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment]{}
}

func (ma *productTest_MockTokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokens.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func TestRegisterTokenAdapter(t *testing.T) {
	tests := []struct {
		desc         string
		chainFamily1 string
		version1     *semver.Version
		chainFamily2 string
		version2     *semver.Version
		expectedErr  string
	}{
		{
			desc:         "registering two adapters with different chain families succeeds",
			chainFamily1: "evm",
			version1:     semver.MustParse("1.0.1"),
			chainFamily2: "solana",
			version2:     semver.MustParse("1.0.1"),
			expectedErr:  "",
		},
		{
			desc:         "registering two adapters with different versions succeeds",
			chainFamily1: "evm",
			version1:     semver.MustParse("1.0.2"),
			chainFamily2: "evm",
			version2:     semver.MustParse("2.0.2"),
			expectedErr:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			registry := tokens.GetTokenAdapterRegistry()

			// First registration should always succeed
			require.NotPanics(t, func() {
				registry.RegisterTokenAdapter(tt.chainFamily1, tt.version1, &productTest_MockTokenAdapter{})
			})

			if tt.expectedErr != "" {
				require.PanicsWithError(t, tt.expectedErr, func() {
					registry.RegisterTokenAdapter(tt.chainFamily2, tt.version2, &productTest_MockTokenAdapter{})
				})
			} else {
				require.NotPanics(t, func() {
					registry.RegisterTokenAdapter(tt.chainFamily2, tt.version2, &productTest_MockTokenAdapter{})
				})
			}
		})
	}
}

// productTest_MockTARReader is a reader-only TokenAdminRegistry; productTest_MockTARManager can also
// unregister.
type productTest_MockTARReader struct{}

func (m *productTest_MockTARReader) GetActivePool(deployment.Environment, uint64, datastore.AddressRef, ...datastore.AddressRef) ([]byte, error) {
	return nil, nil
}

func (m *productTest_MockTARReader) GetTokenAdminRegistryRef(deployment.Environment, uint64) (datastore.AddressRef, error) {
	return datastore.AddressRef{}, nil
}

type productTest_MockTARManager struct{ productTest_MockTARReader }

func (m *productTest_MockTARManager) UnregisterToken() *cldf_ops.Sequence[tokens.UnregisterTokenSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return nil
}

func TestTokenAdminRegistryRegistration(t *testing.T) {
	// The registry is global and registrations are first-write-wins, so each case uses its own family.
	registry := tokens.GetTokenAdapterRegistry()

	t.Run("a manager is also the family's reader", func(t *testing.T) {
		family := "productTest-tar-manager-is-reader"
		manager := &productTest_MockTARManager{}
		registry.RegisterTokenAdminRegistryManager(family, manager)

		reader, ok := registry.GetTokenAdminRegistryReader(family)
		require.True(t, ok)
		require.Same(t, manager, reader)
	})

	t.Run("a manager replaces an earlier reader-only registration", func(t *testing.T) {
		family := "productTest-tar-manager-replaces-reader"
		registry.RegisterTokenAdminRegistryReader(family, &productTest_MockTARReader{})
		manager := &productTest_MockTARManager{}
		registry.RegisterTokenAdminRegistryManager(family, manager)

		reader, ok := registry.GetTokenAdminRegistryReader(family)
		require.True(t, ok)
		require.Same(t, manager, reader)
	})

	t.Run("a later reader-only registration does not replace a manager", func(t *testing.T) {
		family := "productTest-tar-reader-keeps-manager"
		manager := &productTest_MockTARManager{}
		registry.RegisterTokenAdminRegistryManager(family, manager)
		registry.RegisterTokenAdminRegistryReader(family, &productTest_MockTARReader{})

		reader, ok := registry.GetTokenAdminRegistryReader(family)
		require.True(t, ok)
		require.Same(t, manager, reader)
	})

	t.Run("a reader-only family has no manager", func(t *testing.T) {
		family := "productTest-tar-reader-only"
		readerOnly := &productTest_MockTARReader{}
		registry.RegisterTokenAdminRegistryReader(family, readerOnly)

		reader, ok := registry.GetTokenAdminRegistryReader(family)
		require.True(t, ok)
		require.Same(t, readerOnly, reader)
		_, ok = registry.GetTokenAdminRegistryManager(family)
		require.False(t, ok)
	})
}

func TestTokenExpansion_VerifyPreconditions(t *testing.T) {
	const chainSel = uint64(5009297550715157269)

	// Register a mock adapter for the chain family/version the cases below resolve against.
	registry := tokens.GetTokenAdapterRegistry()
	registry.RegisterTokenAdapter("evm", semver.MustParse("9.9.9"), &productTest_MockTokenAdapter{})

	baseInput := func() tokens.TokenExpansionInput {
		return tokens.TokenExpansionInput{
			ChainAdapterVersion: semver.MustParse("9.9.9"),
			TokenExpansionInputPerChain: map[uint64]tokens.TokenExpansionInputPerChain{
				chainSel: {
					TokenPoolVersion: semver.MustParse("9.9.9"),
				},
			},
		}
	}

	tests := []struct {
		name        string
		mutate      func(cfg *tokens.TokenExpansionInput)
		expectedErr string
	}{
		{
			name:   "Success - valid input with no deploy token pool",
			mutate: func(cfg *tokens.TokenExpansionInput) {},
		},
		{
			name: "Failure - invalid chain selector",
			mutate: func(cfg *tokens.TokenExpansionInput) {
				cfg.TokenExpansionInputPerChain = map[uint64]tokens.TokenExpansionInputPerChain{
					0: {TokenPoolVersion: semver.MustParse("9.9.9")},
				}
			},
			expectedErr: "not a valid selector",
		},
		{
			name: "Failure - no adapter registered for version",
			mutate: func(cfg *tokens.TokenExpansionInput) {
				cfg.TokenExpansionInputPerChain[chainSel] = tokens.TokenExpansionInputPerChain{
					TokenPoolVersion: semver.MustParse("8.8.8"),
				}
			},
			expectedErr: "no TokenPoolAdapter registered",
		},
		{
			name: "Failure - invalid liquidity migration amount",
			mutate: func(cfg *tokens.TokenExpansionInput) {
				cfg.TokenExpansionInputPerChain[chainSel] = tokens.TokenExpansionInputPerChain{
					TokenPoolVersion: semver.MustParse("9.9.9"),
					DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
						PoolType: "LockReleaseTokenPool",
						LiquidityMigrationAmount: &tokens.LockReleasePoolLiquidityMigrationAmount{
							Format: tokens.LiquidityMigrationAmountFormatBPS,
							Value:  "10001",
						},
					},
				}
			},
			expectedErr: "invalid liquidity migration amount",
		},
		{
			name: "Failure - liquidity migration on a non lock-release pool",
			mutate: func(cfg *tokens.TokenExpansionInput) {
				cfg.TokenExpansionInputPerChain[chainSel] = tokens.TokenExpansionInputPerChain{
					TokenPoolVersion: semver.MustParse("9.9.9"),
					DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
						PoolType: "BurnMintTokenPool",
						LiquidityMigrationAmount: &tokens.LockReleasePoolLiquidityMigrationAmount{
							Format: tokens.LiquidityMigrationAmountFormatBPS,
							Value:  "5000",
						},
					},
				}
			},
			expectedErr: "only supported for lock-release pools",
		},
		{
			name: "Failure - unsiloed lockbox selector without lockbox groups",
			mutate: func(cfg *tokens.TokenExpansionInput) {
				selector := chainSel
				cfg.TokenExpansionInputPerChain[chainSel] = tokens.TokenExpansionInputPerChain{
					TokenPoolVersion: semver.MustParse("9.9.9"),
					DeployTokenPoolInput: &tokens.DeployTokenPoolInput{
						PoolType:                     "SiloedLockReleaseTokenPool",
						UnsiloedLockBoxChainSelector: &selector,
					},
				}
			},
			expectedErr: "unsiloedLockBoxChainSelector requires lockBoxGroups",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := baseInput()
			tc.mutate(&cfg)

			changeset := tokens.TokenExpansion()
			err := changeset.VerifyPreconditions(deployment.Environment{}, cfg)

			if tc.expectedErr == "" {
				require.NoError(t, err)
				return
			}
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.expectedErr)
		})
	}
}
