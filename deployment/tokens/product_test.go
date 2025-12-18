package tokens_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
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

func (ma *productTest_MockTokenAdapter) DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error) {
	return []byte{}, nil
}

func (ma *productTest_MockTokenAdapter) ManualRegistration() *cldf_ops.Sequence[tokens.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) DeployToken() *cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) DeployTokenVerify(in any) error {
	return nil
}

func (ma *productTest_MockTokenAdapter) DeployTokenPoolForToken() *cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) RegisterToken() *cldf_ops.Sequence[tokens.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.RegisterTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) SetPool() *cldf_ops.Sequence[tokens.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.SetPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
}

func (ma *productTest_MockTokenAdapter) UpdateAuthorities() *cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return &cldf_ops.Sequence[tokens.UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains]{}
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
		{
			desc:         "registering two adapters with same chain family and version fails",
			chainFamily1: "evm",
			version1:     semver.MustParse("1.0.3"),
			chainFamily2: "evm",
			version2:     semver.MustParse("1.0.3"),
			expectedErr:  "TokenAdapter 'evm 1.0.3' already registered",
		},
	}

	for _, tt := range tests {
		tt := tt
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
