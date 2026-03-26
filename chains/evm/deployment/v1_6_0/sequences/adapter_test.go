package sequences

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
)

func fqAddrRef(chainSelector uint64, address, version string) datastore.AddressRef {
	return datastore.AddressRef{
		Address:       address,
		ChainSelector: chainSelector,
		Type:          datastore.ContractType(fee_quoter.ContractType),
		Version:       semver.MustParse(version),
	}
}

func TestGetFeeQuoterAddress(t *testing.T) {
	t.Parallel()

	const chain = uint64(5009297550715157269)

	tests := []struct {
		name           string
		addresses      []datastore.AddressRef
		chainSelector  uint64
		tooHighVersion *semver.Version
		wantAddress    string
		wantErr        string
	}{
		{
			name: "returns sole 1_6 fee quoter",
			addresses: []datastore.AddressRef{
				fqAddrRef(chain, "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "1.6.0"),
			},
			chainSelector: chain,
			wantAddress:   "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
		},
		{
			name: "picks latest among multiple 1_6_plus versions",
			addresses: []datastore.AddressRef{
				fqAddrRef(chain, "0x1111111111111111111111111111111111111111", "1.6.0"),
				fqAddrRef(chain, "0x2222222222222222222222222222222222222222", "1.6.3"),
			},
			chainSelector: chain,
			wantAddress:   "0x2222222222222222222222222222222222222222",
		},
		{
			name: "ignores other contract types on same chain",
			addresses: []datastore.AddressRef{
				{
					Address:       "0x3333333333333333333333333333333333333333",
					ChainSelector: chain,
					Type:          datastore.ContractType("EVM2EVMOnRamp"),
					Version:       semver.MustParse("1.5.0"),
				},
				fqAddrRef(chain, "0x4444444444444444444444444444444444444444", "1.6.0"),
			},
			chainSelector: chain,
			wantAddress:   "0x4444444444444444444444444444444444444444",
		},
		{
			name: "tooHighVersion excludes newer deployments",
			addresses: []datastore.AddressRef{
				fqAddrRef(chain, "0x5555555555555555555555555555555555555555", "1.6.0"),
				fqAddrRef(chain, "0x6666666666666666666666666666666666666666", "1.6.3"),
				fqAddrRef(chain, "0x7777777777777777777777777777777777777777", "2.0.0"),
			},
			chainSelector:  chain,
			tooHighVersion: semver.MustParse("2.0.0"),
			wantAddress:    "0x6666666666666666666666666666666666666666",
		},
		{
			name: "error when no fee quoter for chain",
			addresses: []datastore.AddressRef{
				fqAddrRef(chain+1, "0x8888888888888888888888888888888888888888", "1.6.0"),
			},
			chainSelector: chain,
			wantErr:       "no fee quoter address found for chain selector",
		},
		{
			name: "error when only pre_1_6 semver",
			addresses: []datastore.AddressRef{
				fqAddrRef(chain, "0x9999999999999999999999999999999999999999", "1.5.0"),
			},
			chainSelector: chain,
			wantErr:       "no fee quoter address found for chain selector",
		},
		{
			name: "error when tooHighVersion excludes all candidates",
			addresses: []datastore.AddressRef{
				fqAddrRef(chain, "0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", "2.0.0"),
			},
			chainSelector:  chain,
			tooHighVersion: semver.MustParse("2.0.0"),
			wantErr:        "no fee quoter address found for chain selector",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := GetFeeQuoterAddress(tt.addresses, tt.chainSelector, tt.tooHighVersion)
			if tt.wantErr != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.wantErr)
				require.Equal(t, datastore.AddressRef{}, got)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantAddress, got.Address)
		})
	}
}
