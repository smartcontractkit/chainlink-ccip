package tokens_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
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

func TestPartialTokenTransferFeeConfig_ResolveForAutoMigrate(t *testing.T) {
	legacyEnabled := tokens.TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: 10,
		CustomFinalityTransferFeeBps:  20,
		DefaultFinalityFeeUSDCents:    17,
		CustomFinalityFeeUSDCents:     30,
		DestBytesOverhead:             150_000,
		DestGasOverhead:               50_000,
		IsEnabled:                     true,
	}
	legacyDisabled := tokens.TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: 10,
		DefaultFinalityFeeUSDCents:    17,
		DestBytesOverhead:             150_000,
		DestGasOverhead:               50_000,
		IsEnabled:                     false,
	}
	overrides := tokens.PartialTokenTransferFeeConfig{
		DefaultFinalityFeeUSDCents: cciputils.NewOptional(uint32(99)),
		IsEnabled:                  cciputils.NewOptional(true),
	}

	wantLegacyEnabledMerged := tokens.PartialTokenTransferFeeConfig{}.Populate(overrides.MergeWith(legacyEnabled))
	wantLegacyEnabled := tokens.PartialTokenTransferFeeConfig{}.Populate(legacyEnabled)
	tests := []struct {
		name   string
		errStr string
		inputs *tokens.PartialTokenTransferFeeConfig
		legacy tokens.TokenTransferFeeConfig
		expect *tokens.PartialTokenTransferFeeConfig
	}{
		{
			name:   "legacy enabled without yaml imports legacy fees",
			legacy: legacyEnabled,
			inputs: nil,
			expect: &wantLegacyEnabled,
		},
		{
			name:   "legacy enabled with yaml merges when isEnabled is set",
			legacy: legacyEnabled,
			inputs: &overrides,
			expect: &wantLegacyEnabledMerged,
		},
		{
			name:   "legacy enabled with yaml requires isEnabled",
			legacy: legacyEnabled,
			inputs: &tokens.PartialTokenTransferFeeConfig{
				DefaultFinalityFeeUSDCents: cciputils.NewOptional(uint32(99)),
			},
			errStr: "tokenTransferFeeConfig must set isEnabled",
		},
		{
			name:   "legacy disabled without yaml leaves fees unset",
			legacy: legacyDisabled,
			inputs: nil,
			expect: nil,
		},
		{
			name:   "legacy disabled with yaml passes through when isEnabled is set",
			legacy: legacyDisabled,
			inputs: &overrides,
			expect: &overrides,
		},
		{
			name:   "legacy disabled with yaml requires isEnabled",
			legacy: legacyDisabled,
			inputs: &tokens.PartialTokenTransferFeeConfig{
				DefaultFinalityFeeUSDCents: cciputils.NewOptional(uint32(99)),
			},
			errStr: "tokenTransferFeeConfig must set isEnabled",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.inputs.ResolveForAutoMigrate(tt.legacy)
			if tt.errStr != "" {
				require.EqualError(t, err, tt.errStr)
				require.Nil(t, got)
				return
			}
			require.NoError(t, err)
			if tt.expect == nil {
				require.Nil(t, got)
			} else {
				require.NotNil(t, got)
				require.Equal(t, *tt.expect, *got)
			}
		})
	}
}
