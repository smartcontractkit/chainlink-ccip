package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
)

// JD stores an OCR key bundle's OnchainSigningAddress as bare hex. That address becomes
// signer_address in a ccvcommitteeverifier job spec, which a Chainlink node decodes with
// hexutil.Decode — it rejects a string without the 0x prefix and fails to create the job.
func TestEVMSigningIdentityReaderCanonicalisesAddress(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		addr string
		want string
	}{
		{"bare hex as JD stores it", "47a5eed5c86da7dd9bb75488cd3832dd6782252e", "0x47a5eed5c86da7dd9bb75488cd3832dd6782252e"},
		{"already prefixed", "0x47a5eed5c86da7dd9bb75488cd3832dd6782252e", "0x47a5eed5c86da7dd9bb75488cd3832dd6782252e"},
		{"mixed case", "0x47A5EED5C86DA7DD9BB75488CD3832DD6782252E", "0x47a5eed5c86da7dd9bb75488cd3832dd6782252e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := EVMSigningIdentityReader{}.FromBundle(
				&nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: tt.addr})
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEVMSigningIdentityReaderErrors(t *testing.T) {
	t.Parallel()

	_, err := EVMSigningIdentityReader{}.FromBundle(nil)
	require.ErrorContains(t, err, "nil OCR key bundle")

	_, err = EVMSigningIdentityReader{}.FromBundle(&nodev1.OCR2Config_OCRKeyBundle{})
	require.ErrorContains(t, err, "OnchainSigningAddress is empty")
}

// An empty address must stay empty rather than becoming "0x", which is non-empty and so survives
// every emptiness check a caller might make.
func TestCanonicalEVMAddressPreservesEmpty(t *testing.T) {
	t.Parallel()

	assert.Empty(t, canonicalEVMAddress(""))
}
