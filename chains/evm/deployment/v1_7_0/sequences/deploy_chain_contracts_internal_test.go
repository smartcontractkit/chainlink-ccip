package sequences

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/link"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/committee_verifier"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/require"
)

func TestGetMockReceiverVerifiers(t *testing.T) {
	chainSelector := uint64(12345)
	v1_7_0 := semver.MustParse("1.7.0")

	// Note that address is not specified in the refs below - this is intentional.
	requiredVerifierRef := datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(committee_verifier.ContractType),
		Version:       v1_7_0,
		Qualifier:     "alpha",
	}
	optionalVerifierRef := datastore.AddressRef{
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(committee_verifier.ContractType),
		Version:       v1_7_0,
		Qualifier:     "beta",
	}

	requiredAddr := common.HexToAddress("0x01")
	optionalAddr := common.HexToAddress("0x02")

	// Create address refs with addresses for the datastore
	addresses := []datastore.AddressRef{
		{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(committee_verifier.ContractType),
			Version:       v1_7_0,
			Qualifier:     "alpha",
			Address:       requiredAddr.Hex(),
		},
		{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(router.ContractType),
			Version:       semver.MustParse("1.2.0"),
			Address:       common.HexToAddress("0xAA").Hex(),
		},
		{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(committee_verifier.ContractType),
			Version:       v1_7_0,
			Qualifier:     "gamma",
			Address:       common.HexToAddress("0xCC").Hex(),
		},
	}
	existingAddresses := []datastore.AddressRef{
		{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(committee_verifier.ContractType),
			Version:       v1_7_0,
			Qualifier:     "beta",
			Address:       optionalAddr.Hex(),
		},
		{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(link.ContractType),
			Version:       semver.MustParse("1.0.0"),
			Address:       common.HexToAddress("0xBB").Hex(),
		},
		{
			ChainSelector: chainSelector,
			Type:          datastore.ContractType(committee_verifier.ContractType),
			Version:       v1_7_0,
			Qualifier:     "delta",
			Address:       common.HexToAddress("0xDD").Hex(),
		},
	}

	tests := []struct {
		name               string
		mockReceiverParams MockReceiverParams
		addresses          []datastore.AddressRef
		existingAddresses  []datastore.AddressRef
		expectedRequired   []common.Address
		expectedOptional   []common.Address
		expectErr          bool
		errContains        string
	}{
		{
			name: "happy path - all verifiers found",
			mockReceiverParams: MockReceiverParams{
				RequiredVerifiers: []datastore.AddressRef{requiredVerifierRef},
				OptionalVerifiers: []datastore.AddressRef{optionalVerifierRef},
			},
			addresses:         addresses,
			existingAddresses: existingAddresses,
			expectedRequired:  []common.Address{requiredAddr},
			expectedOptional:  []common.Address{optionalAddr},
			expectErr:         false,
		},
		{
			name: "error - required verifier not found",
			mockReceiverParams: MockReceiverParams{
				RequiredVerifiers: []datastore.AddressRef{requiredVerifierRef},
			},
			addresses:         []datastore.AddressRef{},
			existingAddresses: []datastore.AddressRef{},
			expectErr:         true,
			errContains:       "failed to find required verifier",
		},
		{
			name: "error - optional verifier not found",
			mockReceiverParams: MockReceiverParams{
				OptionalVerifiers: []datastore.AddressRef{optionalVerifierRef},
			},
			addresses:         []datastore.AddressRef{},
			existingAddresses: []datastore.AddressRef{},
			expectErr:         true,
			errContains:       "failed to find optional verifier",
		},
		{
			name: "no verifiers requested",
			mockReceiverParams: MockReceiverParams{
				RequiredVerifiers: []datastore.AddressRef{},
				OptionalVerifiers: []datastore.AddressRef{},
			},
			addresses:         addresses,
			existingAddresses: existingAddresses,
			expectedRequired:  nil,
			expectedOptional:  nil,
			expectErr:         false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			required, optional, err := getMockReceiverVerifiers(tc.mockReceiverParams, tc.addresses, tc.existingAddresses)

			if tc.expectErr {
				require.Error(t, err)
				if tc.errContains != "" {
					require.Contains(t, err.Error(), tc.errContains)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedRequired, required)
				require.Equal(t, tc.expectedOptional, optional)
			}
		})
	}
}
