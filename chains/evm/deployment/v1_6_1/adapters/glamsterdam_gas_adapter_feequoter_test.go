package adapters

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/fee_quoter"
)

// TestClassifyUsedFeeQuoter is a regression test for resolving the FeeQuoter a v1.6 OnRamp is
// actually wired to (not necessarily the datastore's FeeQuoter 1.6.0): a chain whose FeeQuoter was
// upgraded in place to 2.0.0 must be recognized as "handled by the v2.0 changeset instead", not
// silently read/written with the v1.6 ABI.
func TestClassifyUsedFeeQuoter(t *testing.T) {
	const sel = uint64(42)
	fq160 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	fq163 := common.HexToAddress("0x2222222222222222222222222222222222222222")
	fq200 := common.HexToAddress("0x3333333333333333333333333333333333333333")
	unknown := common.HexToAddress("0x4444444444444444444444444444444444444444")

	ref := func(chain uint64, typ string, addr common.Address, version string) datastore.AddressRef {
		return datastore.AddressRef{ChainSelector: chain, Type: datastore.ContractType(typ), Address: addr.Hex(), Version: semver.MustParse(version)}
	}
	addrs := []datastore.AddressRef{
		ref(sel, string(fee_quoter.ContractType), fq160, "1.6.0"),
		ref(sel, string(fee_quoter.ContractType), fq163, "1.6.3"),
		ref(sel, string(fee_quoter.ContractType), fq200, "2.0.0"),
		ref(sel+1, string(fee_quoter.ContractType), unknown, "1.6.0"), // other chain: must not match
		ref(sel, "OnRamp", unknown, "1.6.0"),                          // right chain, wrong type: must not match
	}

	tests := []struct {
		name        string
		used        common.Address
		wantKind    feeQuoterUse
		wantVersion string
	}{
		{"1.6.0 is usable", fq160, feeQuoterV16, "1.6.0"},
		{"1.6.3 is usable", fq163, feeQuoterV16, "1.6.3"},
		{"2.0.0 is another version (handled by the v2.0 changeset)", fq200, feeQuoterOtherVersion, "2.0.0"},
		{"address on another chain does not match", unknown, feeQuoterUnknown, ""},
		{"address not in the datastore", common.HexToAddress("0x5555555555555555555555555555555555555555"), feeQuoterUnknown, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			kind, version := classifyUsedFeeQuoter(addrs, sel, tc.used)
			require.Equal(t, tc.wantKind, kind)
			require.Equal(t, tc.wantVersion, version)
		})
	}
}
