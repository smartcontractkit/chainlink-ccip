package shared

import (
	"fmt"
	"sort"
	"strings"

	chainsel "github.com/smartcontractkit/chain-selectors"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
)

// canonicalEVMAddress lowercases addr and gives it an 0x prefix, leaving an empty address empty.
//
// JD stores an OCR key bundle's OnchainSigningAddress as bare hex. A Chainlink node decodes the
// signer_address of a ccvcommitteeverifier job spec with hexutil.Decode, which rejects a string
// without the prefix, so an address that reaches a job spec unprefixed fails job creation.
//
// It is idempotent, so an already-canonical address passes through unchanged.
func canonicalEVMAddress(addr string) string {
	if addr == "" {
		return ""
	}
	lower := strings.ToLower(addr)
	if !strings.HasPrefix(lower, "0x") {
		return "0x" + lower
	}
	return lower
}

// SigningIdentityReader returns the family-appropriate signer identity from a JD
// OCRKeyBundle. Families that need OnchainSigningPubKey instead of the default
// OnchainSigningAddress register a reader at init time.
type SigningIdentityReader interface {
	FromBundle(bundle *nodev1.OCR2Config_OCRKeyBundle) (string, error)
}

// EVMSigningIdentityReader reads OnchainSigningAddress — the 20-byte EVM-derived
// address. This is the default; families with different identity formats register
// their own reader.
type EVMSigningIdentityReader struct{}

func (EVMSigningIdentityReader) FromBundle(b *nodev1.OCR2Config_OCRKeyBundle) (string, error) {
	if b == nil {
		return "", fmt.Errorf("nil OCR key bundle")
	}
	if b.OnchainSigningAddress == "" {
		return "", fmt.Errorf("OnchainSigningAddress is empty")
	}
	return canonicalEVMAddress(b.OnchainSigningAddress), nil
}

var signingIdentityReaders = map[string]SigningIdentityReader{
	chainsel.FamilyEVM: EVMSigningIdentityReader{},
}

// RegisterSigningIdentityReader associates a chain family with its SigningIdentityReader.
// Called from init() in chain-specific adapter packages.
func RegisterSigningIdentityReader(family string, reader SigningIdentityReader) {
	signingIdentityReaders[family] = reader
}

// SigningIdentityFromBundle returns the signer identity for the given chain family
// from a JD OCRKeyBundle. If the family has registered a reader, that reader is used;
// otherwise the default — OnchainSigningAddress (the EVM-derived address) — is returned.
func SigningIdentityFromBundle(family string, bundle *nodev1.OCR2Config_OCRKeyBundle) (string, error) {
	if reader, ok := signingIdentityReaders[family]; ok {
		return reader.FromBundle(bundle)
	}
	if bundle == nil {
		return "", fmt.Errorf("nil OCR key bundle")
	}
	if bundle.OnchainSigningAddress == "" {
		return "", fmt.Errorf("OnchainSigningAddress is empty")
	}
	// Deliberately not canonicalised: this branch is reached only by a family that registered no
	// reader, and 0x-prefixing an identity that is not an EVM address would corrupt it. EVM has a
	// reader registered by default, so the fix above covers it.
	return bundle.OnchainSigningAddress, nil
}

// RegisteredSigningIdentityFamilies returns the families that have registered a
// SigningIdentityReader, in sorted order.
func RegisteredSigningIdentityFamilies() []string {
	families := make([]string, 0, len(signingIdentityReaders))
	for f := range signingIdentityReaders {
		families = append(families, f)
	}
	sort.Strings(families)
	return families
}
