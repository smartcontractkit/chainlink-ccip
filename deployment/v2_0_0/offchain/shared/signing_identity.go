package shared

import (
	"fmt"
	"sort"

	chainsel "github.com/smartcontractkit/chain-selectors"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"
)

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
	return b.OnchainSigningAddress, nil
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
